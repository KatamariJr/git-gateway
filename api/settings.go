package api

import "net/http"

type Settings struct {
	GitHub    bool     `json:"github_enabled"`
	GitLab    bool     `json:"gitlab_enabled"`
	BitBucket bool     `json:"bitbucket_enabled"`
	Roles     []string `json:"roles"`
}

func (a *API) Settings(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	config := getConfig(ctx)

	settings := Settings{
		GitHub:    config.GitHub.Repo != "",
		GitLab:    config.GitLab.Repo != "",
		BitBucket: config.BitBucket.Repo != "",
		Roles:     config.Roles,
	}

	return sendJSON(w, http.StatusOK, &settings)
}

func (a *API) InstanceSettings(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	config := getInstance(ctx)

	settings := Settings{
		GitHub:    config.BaseConfig.GitHub.Repo != "",
		GitLab:    config.BaseConfig.GitLab.Repo != "",
		BitBucket: config.BaseConfig.BitBucket.Repo != "",
		Roles:     config.BaseConfig.Roles,
	}

	return sendJSON(w, http.StatusOK, &settings)
}
