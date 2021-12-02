package aggregate

import (
	"github_api/api"
	"testing"
)

type mockClient struct {
}

func (m mockClient) UserData(username string) (*api.UserData, error) {
	return &api.UserData{
		Login:     "testuser",
		Followers: 5,
	}, nil
}

func (m mockClient) UserRepositories(username string) ([]*api.RepositoryData, error) {
	return []*api.RepositoryData{
		{
			Name: "repo-1",
		},
		{
			Name: "repo-2",
		},
	}, nil
}

func (m mockClient) LanguageData(username, repo string) (api.RepositoryLanguageData, error) {
	data := make(map[string]api.RepositoryLanguageData)

	data["repo-1"] = api.RepositoryLanguageData{
		"Go":   5,
		"Java": 3,
	}

	data["repo-2"] = api.RepositoryLanguageData{
		"C++":  1,
		"Java": 3,
	}

	return data[repo], nil
}

func TestAggregator_AggregateUser(t *testing.T) {
	client := &mockClient{}

	aggregator := New(client)
	actual, err := aggregator.AggregateUser("testuser")
	if err != nil {
		t.Error(err)
	}

	expected, err := client
}
