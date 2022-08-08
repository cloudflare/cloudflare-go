package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	testScriptName = "this-is_my_script-01"
	testTailID     = "03dc9f77817b488fb26c5861ec18f791"
)

func TestWorkersTail_StartWorkersTail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails", testAccountID, testScriptName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "03dc9f77817b488fb26c5861ec18f791",
    "url": "wss://tail.developers.workers.dev/03dc9f77817b488fb26c5861ec18f791",
    "expires_at": "2021-08-20T19:15:51Z"
  }
}`) //nolint
	})

	// Make sure missing account ID is thrown
	_, err := client.StartWorkersTail(context.Background(), StartWorkersTailParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	// Make sure missing script ID is thrown
	_, err = client.StartWorkersTail(context.Background(), StartWorkersTailParameters{AccountID: testAccountID})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingScriptName, err)
	}

	res, err := client.StartWorkersTail(context.Background(), StartWorkersTailParameters{AccountID: testAccountID, ScriptName: testScriptName})
	expiresAt, _ := time.Parse(time.RFC3339, "2021-08-20T19:15:51Z")
	want := WorkersTail{
		ID:        "03dc9f77817b488fb26c5861ec18f791",
		URL:       "wss://tail.developers.workers.dev/03dc9f77817b488fb26c5861ec18f791",
		ExpiresAt: &expiresAt,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersTail_ListWorkersTail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails", testAccountID, testScriptName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "03dc9f77817b488fb26c5861ec18f791",
    "url": "wss://tail.developers.workers.dev/03dc9f77817b488fb26c5861ec18f791",
    "expires_at": "2021-08-20T19:15:51Z"
  }
}`) //nolint
	})

	// Make sure missing account ID is thrown
	_, err := client.ListWorkersTail(context.Background(), ListWorkersTailParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	// Make sure missing script ID is thrown
	_, err = client.ListWorkersTail(context.Background(), ListWorkersTailParameters{AccountID: testAccountID})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingScriptName, err)
	}

	res, err := client.ListWorkersTail(context.Background(), ListWorkersTailParameters{AccountID: testAccountID, ScriptName: testScriptName})
	expiresAt, _ := time.Parse(time.RFC3339, "2021-08-20T19:15:51Z")
	want := WorkersTail{
		ID:        "03dc9f77817b488fb26c5861ec18f791",
		URL:       "wss://tail.developers.workers.dev/03dc9f77817b488fb26c5861ec18f791",
		ExpiresAt: &expiresAt,
	}

	if assert.NoError(t, err) {
		assert.Equal(t, want, res)
	}
}

func TestWorkersTail_DeleteWorkersTail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails/%s", testAccountID, testScriptName, testTailID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
}`) //nolint
	})

	// Make sure missing account ID is thrown
	err := client.DeleteWorkersTail(context.Background(), DeleteWorkersTailParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	// Make sure missing script ID is thrown
	err = client.DeleteWorkersTail(context.Background(), DeleteWorkersTailParameters{AccountID: testAccountID})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingScriptName, err)
	}

	// Make sure missing Tail ID is thrown
	err = client.DeleteWorkersTail(context.Background(), DeleteWorkersTailParameters{AccountID: testAccountID, ScriptName: testScriptName})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingTailID, err)
	}

	err = client.DeleteWorkersTail(context.Background(), DeleteWorkersTailParameters{AccountID: testAccountID, ScriptName: testScriptName, TailID: testTailID})
	assert.NoError(t, err)
}
