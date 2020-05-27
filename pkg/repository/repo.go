package repository

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/dghubble/sling"
	"github.com/ghodss/yaml"
	"github.com/jinzhu/copier"
	"github.com/mojo-zd/helm-api/pkg/typed/charts"
	"github.com/rs/zerolog/log"
	helmrepo "helm.sh/helm/v3/pkg/repo"
)

var indexYAML = "index.yaml"

type repoOptions struct {
	name                  string
	url                   string
	username              string
	password              string
	certFile              string
	keyFile               string
	caFile                string
	insecureSkipTLSverify bool
	repoFile              string
	repoCache             string
}

// NewRepoOption create repo option
func NewRepoOption(opts ...repoOption) *repoOptions {
	opt := new(repoOptions)
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// GetRepo get charts from repository
func (repo *repoOptions) GetRepo() (*helmrepo.IndexFile, []byte, error) {
	u, err := parseURL(repo.url)
	if err != nil {
		return nil, nil, err
	}

	resp, err := repo.index(u)
	if err != nil {
		return nil, resp, err
	}

	index, err := repo.toIndexFile(resp)
	return index, resp, err
}

func (repo *repoOptions) ChartsFromIndex(index *helmrepo.IndexFile) []charts.Chart {
	var charts []charts.Chart
	for _, entry := range index.Entries {
		// 如果最新版本都弃用了 则跳过该chart
		if entry[0].Deprecated {
			log.Warn().Str("chart", entry[0].Name).Msg("chart has deprecated!!!")
			continue
		}
		c, err := newChart(entry)
		if err != nil {
			return charts
		}

		c.RepoName = repo.name
		c.RepoURL = repo.url
		charts = append(charts, c)
	}
	return charts
}

func newChart(entry helmrepo.ChartVersions) (charts.Chart, error) {
	var c charts.Chart
	chartVer := entry[0]
	if err := copier.Copy(&c, chartVer); err != nil {
		log.Error().Err(err).Str("chart name", chartVer.Name).Msg("copy chart failed")
		return c, err
	}
	if err := copier.Copy(&c.ChartVersions, entry); err != nil {
		log.Error().Err(err).Str("chart name", chartVer.Name).Msg("copy chart version failed")
	}
	return c, nil
}

func (repo *repoOptions) toIndexFile(body []byte) (*helmrepo.IndexFile, error) {
	var index helmrepo.IndexFile
	err := yaml.Unmarshal(body, &index)
	if err != nil {
		log.Error().Err(err).Msg("byte to yaml failed")
		return nil, err
	}
	index.SortEntries()

	return &index, nil
}

func (repo *repoOptions) index(indexURL *url.URL) ([]byte, error) {
	indexURL.Path = path.Join(indexURL.Path, indexYAML)

	request, err := sling.New().Get(indexURL.String()).Request()
	if err != nil {
		log.Error().Err(err).Str("request url", indexURL.String()).Msg("request index failed")
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Str("request url", indexURL.String()).Msg("do request failed")
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("can't read response body")
		return nil, err
	}
	return body, nil
}

func parseURL(repoURL string) (*url.URL, error) {
	repoURL = strings.TrimSpace(repoURL)
	u, err := url.ParseRequestURI(repoURL)
	if err != nil {
		log.Error().Err(err).Msg("parse url failed")
		return nil, err
	}
	return u, nil
}
