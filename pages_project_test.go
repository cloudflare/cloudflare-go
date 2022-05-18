package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testPagesProjectResponse = `
    {
        "name": "Test Pages Project",
        "id": "5a321fc7-3162-7d36-adce-1213996a7",
        "created_on": "2021-01-01T00:00:00Z",
        "subdomain": "test.pages.dev",
        "domains": [
          "testdomain.com",
          "testdomain.org"
        ],
        "source": {
          "type": "github",
          "config": {
            "owner": "cloudflare",
            "repo_name": "pages-test",
            "production_branch": "main",
            "pr_comments_enabled": true,
            "deployments_enabled": true
          }
        },
        "build_config": {
          "build_command": "npm run build",
          "destination_dir": "build",
          "root_dir": "/",
          "web_analytics_tag": "0ee1d926cd60d2618a108d4232a75b73",
          "web_analytics_token": "c05bb382259183db3a0a822b64c11459"
        },
        "deployment_configs": {
          "preview": {
            "env_vars": {
              "BUILD_VERSION": {
                "value": "1.2"
              },
              "ENV": {
                "value": "preview"
              }
            }
          },
          "production": {
            "env_vars": {
              "BUILD_VERSION": {
                "value": "1.2"
              },
              "ENV": {
                "value": "production"
              }
            }
          }
        },
        "latest_deployment": {
          "id": "c35216d1-ebac-1a3a-d56c-ad74e54e5",
          "short_id": "c35216d1",
          "project_id": "5a321fc7-3162-7d36-adce-1213996a7",
          "project_name": "pages-test",
          "environment": "preview",
          "url": "https://c35216d1.pages-test.pages.dev",
          "created_on": "2021-03-09T00:55:03.923456Z",
          "modified_on": "2021-03-09T00:58:59.045655Z",
          "aliases": [
            "https://branchname.pages-test.pages.dev"
          ],
          "latest_stage": {
            "name": "deploy",
            "started_on": "2021-03-09T00:55:03.923456Z",
            "ended_on": "2021-03-09T00:58:59.045655Z",
            "status": "success"
          },
          "env_vars": {
            "BUILD_VERSION": {
              "value": "1.2"
            },
            "ENV": {
              "value": "STAGING"
            }
          },
          "deployment_trigger": {
            "type": "ad_hoc",
            "metadata": {
              "branch": "main",
              "commit_hash": "fa26be19de6bff93f70bc2308434e4a440bbad02",
              "commit_message": "Update index.html"
            }
          },
          "stages": [
            {
              "name": "queued",
              "started_on": "2021-06-03T15:38:15.608194Z",
              "ended_on": "2021-06-03T15:39:03.134378Z",
              "status": "active"
            },
            {
              "name": "test_stage_1",
              "started_on": null,
              "ended_on": null,
              "status": "idle"
            }
          ],
          "build_config": {
            "build_command": "npm run build",
            "destination_dir": "build",
            "root_dir": "/",
            "web_analytics_tag": "0ee1d926cd60d2618a108d4232a75b73",
            "web_analytics_token": "c05bb382259183db3a0a822b64c11459"
          },
          "source": {
            "type": "github",
            "config": {
              "owner": "cloudflare",
              "repo_name": "pages-test",
              "production_branch": "main",
              "pr_comments_enabled": true,
              "deployments_enabled": true
            }
          }
        },
        "canonical_deployment": {
          "id": "c35216d1-ebac-1a3a-d56c-ad74e54e5",
          "short_id": "c35216d1",
          "project_id": "5a321fc7-3162-7d36-adce-1213996a7",
          "project_name": "pages-test",
          "environment": "preview",
          "url": "https://c35216d1.pages-test.pages.dev",
          "created_on": "2021-03-09T00:55:03.923456Z",
          "modified_on": "2021-03-09T00:58:59.045655Z",
          "aliases": [
            "https://branchname.pages-test.pages.dev"
          ],
          "latest_stage": {
            "name": "deploy",
            "started_on": "2021-03-09T00:55:03.923456Z",
            "ended_on": "2021-03-09T00:58:59.045655Z",
            "status": "success"
          },
          "env_vars": {
            "BUILD_VERSION": {
              "value": "1.2"
            },
            "ENV": {
              "value": "STAGING"
            }
          },
          "deployment_trigger": {
            "type": "ad_hoc",
            "metadata": {
              "branch": "main",
              "commit_hash": "fa26be19de6bff93f70bc2308434e4a440bbad02",
              "commit_message": "Update index.html"
            }
          },
          "stages": [
            {
              "name": "queued",
              "started_on": "2021-06-03T15:38:15.608194Z",
              "ended_on": "2021-06-03T15:39:03.134378Z",
              "status": "active"
            },
            {
              "name": "test_stage_1",
              "started_on": null,
              "ended_on": null,
              "status": "idle"
            }
          ],
          "build_config": {
            "build_command": "npm run build",
            "destination_dir": "build",
            "root_dir": "/",
            "web_analytics_tag": "0ee1d926cd60d2618a108d4232a75b73",
            "web_analytics_token": "c05bb382259183db3a0a822b64c11459"
          },
          "source": {
            "type": "github",
            "config": {
              "owner": "cloudflare",
              "repo_name": "pages-test",
              "production_branch": "main",
              "pr_comments_enabled": true,
              "deployments_enabled": true
            }
          }
        }
      }`
)

var (
	pagesProjectCreatedOn, _ = time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	expectedPagesProject = &PagesProject{
		SubDomain: "test.pages.dev",
		Name:      "Test Pages Project",
		Domains: []string{
			"testdomain.com",
			"testdomain.org",
		},
		CanonicalDeployment: *expectedPagesProjectDeployment,
		BuildConfig:         *expectedPagesProjectBuildConfig,
		CreatedOn:           &pagesProjectCreatedOn,
		DeploymentConfigs:   *expectedPagesProjectDeploymentConfigs,
		Source:              *expectedPagesProjectSource,
		ID:                  "5a321fc7-3162-7d36-adce-1213996a7",
		LatestDeployment:    *expectedPagesProjectDeployment,
	}

	deploymentCreatedOn, _  = time.Parse(time.RFC3339, "2021-03-09T00:55:03.923456Z")
	deploymentModifiedOn, _ = time.Parse(time.RFC3339, "2021-03-09T00:58:59.045655Z")

	expectedPagesProjectDeployment = &PagesProjectDeployment{
		ID:          "c35216d1-ebac-1a3a-d56c-ad74e54e5",
		ShortID:     "c35216d1",
		ProjectID:   "5a321fc7-3162-7d36-adce-1213996a7",
		ProjectName: "pages-test",
		Environment: "preview",
		URL:         "https://c35216d1.pages-test.pages.dev",
		CreatedOn:   &deploymentCreatedOn,
		ModifiedOn:  &deploymentModifiedOn,
		Aliases: []string{
			"https://branchname.pages-test.pages.dev",
		},
		LatestStage: *expectedPagesProjectLatestDeploymentStage,
		EnvVars: map[string]map[string]string{
			"BUILD_VERSION": {
				"value": "1.2",
			},
			"ENV": {
				"value": "STAGING",
			},
		},
		DeploymentTrigger: *expectedPagesProjectDeploymentTrigger,
		Stages:            expectedStages,
		BuildConfig:       *expectedPagesProjectBuildConfig,
		Source:            *expectedPagesProjectSource,
	}

	latestDeploymentStageStartedOn, _ = time.Parse(time.RFC3339, "2021-03-09T00:55:03.923456Z")
	latestDeploymentStageEndedOn, _   = time.Parse(time.RFC3339, "2021-03-09T00:58:59.045655Z")

	expectedPagesProjectLatestDeploymentStage = &PagesProjectDeploymentStage{
		Name:      "deploy",
		StartedOn: &latestDeploymentStageStartedOn,
		EndedOn:   &latestDeploymentStageEndedOn,
		Status:    "success",
	}

	expectedPagesProjectDeploymentTrigger = &PagesProjectDeploymentTrigger{
		Type:     "ad_hoc",
		Metadata: expectedPagesProjectDeploymentTriggerMetadata,
	}

	expectedPagesProjectDeploymentTriggerMetadata = &PagesProjectDeploymentTriggerMetadata{
		Branch:        "main",
		CommitHash:    "fa26be19de6bff93f70bc2308434e4a440bbad02",
		CommitMessage: "Update index.html",
	}

	queuedStageStartedOn, _ = time.Parse(time.RFC3339, "2021-06-03T15:38:15.608194Z")
	queuedStageEndedOn, _   = time.Parse(time.RFC3339, "2021-06-03T15:39:03.134378Z")

	expectedStages = []PagesProjectDeploymentStage{
		{
			Name:      "queued",
			StartedOn: &queuedStageStartedOn,
			EndedOn:   &queuedStageEndedOn,
			Status:    "active",
		},
		{
			Name:   "test_stage_1",
			Status: "idle",
		},
	}

	expectedPagesProjectBuildConfig = &PagesProjectBuildConfig{
		BuildCommand:      "npm run build",
		DestinationDir:    "build",
		RootDir:           "/",
		WebAnalyticsTag:   "0ee1d926cd60d2618a108d4232a75b73",
		WebAnalyticsToken: "c05bb382259183db3a0a822b64c11459",
	}

	expectedPagesProjectDeploymentConfigs = &PagesProjectDeploymentConfigs{
		Preview:    *expectedPagesProjectDeploymentConfigPreview,
		Production: *expectedPagesProjectDeploymentConfigProduction,
	}

	expectedPagesProjectDeploymentConfigPreview = &PagesProjectDeploymentConfigEnvironment{
		EnvVars: map[string]PagesProjectDeploymentVar{
			"BUILD_VERSION": {
				Value: "1.2",
			},
			"ENV": {
				Value: "preview",
			},
		},
	}

	expectedPagesProjectDeploymentConfigProduction = &PagesProjectDeploymentConfigEnvironment{
		EnvVars: map[string]PagesProjectDeploymentVar{
			"BUILD_VERSION": {
				Value: "1.2",
			},
			"ENV": {
				Value: "production",
			},
		},
	}

	expectedPagesProjectSource = &PagesProjectSource{
		Type:   "github",
		Config: expectedPagesProjectSourceConfig,
	}

	expectedPagesProjectSourceConfig = &PagesProjectSourceConfig{
		Owner:              "cloudflare",
		RepoName:           "pages-test",
		ProductionBranch:   "main",
		PRCommentsEnabled:  true,
		DeploymentsEnabled: true,
	}
)

func TestListPagesProjects(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": [
                %s
            ],
            "result_info": {
                "page": 1,
                "per_page": 100,
                "count": 1,
                "total_count": 1
            }
        }`, testPagesProjectResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects", handler)

	expectedPagesProjects := []PagesProject{
		*expectedPagesProject,
	}
	expectedResultInfo := ResultInfo{
		Page:    1,
		PerPage: 100,
		Count:   1,
		Total:   1,
	}
	actual, resultInfo, err := client.ListPagesProjects(context.Background(), testAccountID, PaginationOptions{})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedPagesProjects, actual)
		assert.Equal(t, expectedResultInfo, resultInfo)
	}
}

func TestPagesProject(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": %s
        }`, testPagesProjectResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/Test Pages Project", handler)

	actual, err := client.PagesProject(context.Background(), testAccountID, "Test Pages Project")
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesProject, actual)
	}
}

func TestCreatePagesProject(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": %s
        }`, testPagesProjectResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects", handler)

	actual, err := client.CreatePagesProject(context.Background(), testAccountID, *expectedPagesProject)
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesProject, actual)
	}
}

func TestUpdatePagesProject(t *testing.T) {
	setup()
	defer teardown()

	updateAttributes := &PagesProject{
		Name: "updated-project-name",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method, "Expected method 'PATCH', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": %s
        }`, testPagesProjectResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/Test Pages Project", handler)

	_, err := client.UpdatePagesProject(context.Background(), testAccountID, "Test Pages Project", *updateAttributes)

	t.Log(err)

	assert.NoError(t, err)
}

func TestDeletePagesProject(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
            "success": true,
            "errors": [],
            "messages": [],
            "result": null
        }`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/Test Pages Project", handler)

	err := client.DeletePagesProject(context.Background(), testAccountID, "Test Pages Project")
	assert.NoError(t, err)
}
