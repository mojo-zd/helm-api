package repository

import (
	"testing"
)

var (
	repoURL     = "https://charts.bitnami.com/bitnami"
	mojoRepoURL = "https://mojo-zd.github.io/helm-charts"
	repo        = NewRepoOption(WithName("bitnami"), WithURL(mojoRepoURL))
)

func TestRepo(t *testing.T) {
	_, _, err := repo.GetRepo()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestChart(t *testing.T) {
	indexFile, _, err := repo.GetRepo()
	if err != nil {
		t.Error(err)
		return
	}
	charts := repo.ChartsFromIndex(indexFile)
	for _, c := range charts {
		t.Log(c)
	}
}
