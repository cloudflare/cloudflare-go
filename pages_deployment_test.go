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
	testPagesDeplyomentResponse = `
	{
		"id": "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
		"short_id": "0012e50b",
		"project_id": "80776025-b1bd-4181-993f-8238c27d226f",
		"project_name": "test",
		"environment": "production",
		"url": "https://0012e50b.test.pages.dev",
		"created_on": "2021-01-01T00:00:00Z",
		"modified_on": "2021-01-01T00:00:00Z",
		"latest_stage": {
			"name": "deploy",
			"started_on": "2021-01-01T00:00:00Z",
			"ended_on": "2021-01-01T00:00:00Z",
			"status": "success"
		},
		"deployment_trigger": {
			"type": "ad_hoc",
			"metadata": {
				"branch": "main",
				"commit_hash": "20fb65fa9d7fd2a11f7fa3ebdc44137b263ee835",
				"commit_message": "Test commit"
			}
		},
		"stages": [
			{
				"name": "queued",
				"started_on": "2021-01-01T00:00:00Z",
				"ended_on": "2021-01-01T00:00:00Z",
				"status": "success"
			},
			{
				"name": "initialize",
				"started_on": "2021-01-01T00:00:00Z",
				"ended_on": "2021-01-01T00:00:00Z",
				"status": "success"
			},
			{
				"name": "clone_repo",
				"started_on": "2021-01-01T00:00:00Z",
				"ended_on": "2021-01-01T00:00:00Z",
				"status": "success"
			},
			{
				"name": "build",
				"started_on": "2021-01-01T00:00:00Z",
				"ended_on": "2021-01-01T00:00:00Z",
				"status": "success"
			},
			{
				"name": "deploy",
				"started_on": "2021-01-01T00:00:00Z",
				"ended_on": "2021-01-01T00:00:00Z",
				"status": "success"
			}
		],
		"build_config": {
			"build_command": "bash test.sh",
			"destination_dir": "",
			"root_dir": "",
			"web_analytics_tag": null,
			"web_analytics_token": null
		},
		"source": {
			"type": "github",
			"config": {
				"owner": "coudflare",
				"repo_name": "Test",
				"production_branch": "main",
				"pr_comments_enabled": false
			}
		},
		"env_vars": {
			"NODE_VERSION": {
				"value": "16"
			}
		},
		"aliases": null
	}`

	testPagesDeploymentStageLogsResponse = `
	{
		"name": "build",
		"started_on": "2021-01-01T00:00:00Z",
		"ended_on": "2021-01-01T00:00:00Z",
		"status": "success",
		"start": 0,
		"end": 5,
		"total": 6,
		"data": [
			{
				"id": 0,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Installing dependencies"
			},
			{
				"id": 1,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Verify run directory"
			},
			{
				"id": 2,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Executing user command: bash test.sh"
			},
			{
				"id": 3,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Finished"
			},
			{
				"id": 4,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Building functions..."
			},
			{
				"id": 5,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Validating asset output directory"
			},
			{
				"id": 6,
				"timestamp": "2021-01-01T00:00:00Z",
				"message": "Parsed 2 valid header rules."
			}
		]
	}`

	testPagesDeploymentLogsResponse = `
	{
		"total": 6,
		"includes_container_logs": true,
		"data": [
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Installing dependencies"
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Verify run directory"
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Executing user command: bash test.sh"
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Finished"
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Building functions..."
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Validating asset output directory"
			},
			{
				"ts": "2021-01-01T00:00:00Z",
				"line": "Parsed 2 valid header rules."
			}
		]
	}`
)

var (
	pagesDeploymentDummyTime, _ = time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	expectedPagesDeployment = &PagesProjectDeployment{
		ID:          "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
		ShortID:     "0012e50b",
		ProjectID:   "80776025-b1bd-4181-993f-8238c27d226f",
		ProjectName: "test",
		Environment: "production",
		URL:         "https://0012e50b.test.pages.dev",
		CreatedOn:   &pagesDeploymentDummyTime,
		ModifiedOn:  &pagesDeploymentDummyTime,
		Aliases:     nil,
		LatestStage: *expectedPagesDeploymentLatestStage,
		EnvVars: map[string]map[string]string{
			"NODE_VERSION": {
				"value": "16",
			},
		},
		DeploymentTrigger: PagesProjectDeploymentTrigger{
			Type: "ad_hoc",
			Metadata: &PagesProjectDeploymentTriggerMetadata{
				Branch:        "main",
				CommitHash:    "20fb65fa9d7fd2a11f7fa3ebdc44137b263ee835",
				CommitMessage: "Test commit",
			},
		},
		Stages: expectedPagesDeploymentStages,
		BuildConfig: PagesProjectBuildConfig{
			BuildCommand:   "bash test.sh",
			DestinationDir: "",
			RootDir:        "",
		},
		Source: PagesProjectSource{
			Type: "github",
			Config: &PagesProjectSourceConfig{
				Owner:             "coudflare",
				RepoName:          "Test",
				ProductionBranch:  "main",
				PRCommentsEnabled: false,
			},
		},
	}

	expectedPagesDeploymentStages = []PagesProjectDeploymentStage{
		{
			Name:      "queued",
			StartedOn: &pagesDeploymentDummyTime,
			EndedOn:   &pagesDeploymentDummyTime,
			Status:    "success",
		},
		{
			Name:      "initialize",
			StartedOn: &pagesDeploymentDummyTime,
			EndedOn:   &pagesDeploymentDummyTime,
			Status:    "success",
		},
		{
			Name:      "clone_repo",
			StartedOn: &pagesDeploymentDummyTime,
			EndedOn:   &pagesDeploymentDummyTime,
			Status:    "success",
		},
		{
			Name:      "build",
			StartedOn: &pagesDeploymentDummyTime,
			EndedOn:   &pagesDeploymentDummyTime,
			Status:    "success",
		},
		*expectedPagesDeploymentLatestStage,
	}

	expectedPagesDeploymentLatestStage = &PagesProjectDeploymentStage{
		Name:      "deploy",
		StartedOn: &pagesDeploymentDummyTime,
		EndedOn:   &pagesDeploymentDummyTime,
		Status:    "success",
	}

	expectedPagesDeploymentStageLogs = &PagesDeploymentStageLogs{
		Name:      "build",
		StartedOn: &pagesDeploymentDummyTime,
		EndedOn:   &pagesDeploymentDummyTime,
		Status:    "success",
		Start:     0,
		End:       5,
		Total:     6,
		Data:      expectedPagesDeploymentStageLogEntries,
	}

	expectedPagesDeploymentStageLogEntries = []PagesDeploymentStageLogEntry{
		{
			ID:        0,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Installing dependencies",
		},
		{
			ID:        1,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Verify run directory",
		},
		{
			ID:        2,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Executing user command: bash test.sh",
		},
		{
			ID:        3,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Finished",
		},
		{
			ID:        4,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Building functions...",
		},
		{
			ID:        5,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Validating asset output directory",
		},
		{
			ID:        6,
			Timestamp: &pagesDeploymentDummyTime,
			Message:   "Parsed 2 valid header rules.",
		},
	}

	expectedPagesDeploymentLogs = &PagesDeploymentLogs{
		Total:                 6,
		IncludesContainerLogs: true,
		Data:                  expectedPagesDeploymentLogEntries,
	}

	expectedPagesDeploymentLogEntries = []PagesDeploymentLogEntry{
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Installing dependencies",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Verify run directory",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Executing user command: bash test.sh",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Finished",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Building functions...",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Validating asset output directory",
		},
		{
			Timestamp: &pagesDeploymentDummyTime,
			Line:      "Parsed 2 valid header rules.",
		},
	}
)

func TestListPagesDeployments(t *testing.T) {
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
		}`, testPagesDeplyomentResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments", handler)

	expectedPagesDeployments := []PagesProjectDeployment{
		*expectedPagesDeployment,
	}
	expectedResultInfo := ResultInfo{
		Page:    1,
		PerPage: 100,
		Count:   1,
		Total:   1,
	}
	actual, resultInfo, err := client.ListPagesDeployments(context.Background(), ListPagesDeploymentsParams{
		AccountID:         testAccountID,
		ProjectName:       "test",
		PaginationOptions: PaginationOptions{},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedPagesDeployments, actual)
		assert.Equal(t, expectedResultInfo, resultInfo)
	}
}

func TestGetPagesDeploymentInfo(t *testing.T) {
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
		}`, testPagesDeplyomentResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe", handler)

	actual, err := client.GetPagesDeploymentInfo(context.Background(), GetPagesDeploymentInfoParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeployment, actual)
	}
}

func TestGetPagesDeploymentStageLogs(t *testing.T) {
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
		}`, testPagesDeploymentStageLogsResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe/history/build/logs", handler)

	actual, err := client.GetPagesDeploymentStageLogs(context.Background(), GetPagesDeploymentStageLogsParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
		StageName:    "build",
		SizeOptions:  SizeOptions{},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeploymentStageLogs, actual)
	}
}

func TestGetPagesDeploymentLogs(t *testing.T) {
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
		}`, testPagesDeploymentLogsResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe/history/logs", handler)

	actual, err := client.GetPagesDeploymentLogs(context.Background(), GetPagesDeploymentLogsParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
		SizeOptions:  SizeOptions{},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeploymentLogs, actual)
	}
}

func TestDeletePagesDeployment(t *testing.T) {
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

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe", handler)

	err := client.DeletePagesDeployment(context.Background(), DeletePagesDeploymentParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
	})
	assert.NoError(t, err)
}

func TestCreatePagesDeployment(t *testing.T) {
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
		}`, testPagesDeplyomentResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments", handler)

	actual, err := client.CreatePagesDeployment(context.Background(), CreatePagesDeploymentParams{
		AccountID:   testAccountID,
		ProjectName: "test",
	})

	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeployment, actual)
	}
}

func TestRetryPagesDeployment(t *testing.T) {
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
		}`, testPagesDeplyomentResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe/retry", handler)

	actual, err := client.RetryPagesDeployment(context.Background(), RetryPagesDeploymentParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeployment, actual)
	}
}

func TestRollbackPagesDeployment(t *testing.T) {
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
		}`, testPagesDeplyomentResponse)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/pages/projects/test/deployments/0012e50b-fa5d-44db-8cb5-1f372785dcbe/rollback", handler)

	actual, err := client.RollbackPagesDeployment(context.Background(), RollbackPagesDeploymentParams{
		AccountID:    testAccountID,
		ProjectName:  "test",
		DeploymentID: "0012e50b-fa5d-44db-8cb5-1f372785dcbe",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, *expectedPagesDeployment, actual)
	}
}
