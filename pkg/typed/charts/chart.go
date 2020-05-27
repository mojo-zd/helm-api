package charts

import "time"

// helm chart information
type Chart struct {
	Name          string         `json:"name"`
	RepoName      string         `json:"repoName"`
	RepoURL       string         `json:"repoURL"`
	Description   string         `json:"description"`
	Home          string         `json:"home"`
	Keywords      string         `json:"keywords"`
	Sources       string         `json:"sources"`
	Icon          string         `json:"icon"`
	ChartVersions []ChartVersion `json:"chartVersions"`
}

type ChartVersion struct {
	Version    string    `json:"version"`
	AppVersion string    `json:"appVersion"`
	Created    time.Time `json:"created"`
	Digest     string    `json:"digest"`
	URLs       []string  `json:"urls"`
}
