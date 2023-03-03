package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testChallengeWidgetSiteKey = "0x4AAF00AAAABn0R22HWm-YUc"

var (
	challengeWidgetCreatedOn, _  = time.Parse(time.RFC3339, "2014-01-01T05:20:00.123123Z")
	challengeWidgetModifiedOn, _ = time.Parse(time.RFC3339, "2014-01-01T05:20:00.123123Z")
	expectedChallengeWidget      = ChallengeWidget{
		SiteKey:    "0x4AAF00AAAABn0R22HWm-YUc",
		Secret:     "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
		CreatedOn:  &challengeWidgetCreatedOn,
		ModifiedOn: &challengeWidgetModifiedOn,
		Name:       "blog.cloudflare.com login form",
		Domains: []string{
			"203.0.113.1",
			"cloudflare.com",
			"blog.example.com",
		},
		Type: "invisible",
	}
)

func TestChallengeWidgets_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
			{
			  "success": true,
			  "errors": [],
			  "messages": [],
			  "result": {
				"sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
				"secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
				"created_on": "2014-01-01T05:20:00.123123Z",
				"modified_on": "2014-01-01T05:20:00.123123Z",
				"name": "blog.cloudflare.com login form",
				"domains": [
				  "203.0.113.1",
				  "cloudflare.com",
				  "blog.example.com"
				],
				"type": "invisible"
			  }
			}`)
	})

	// Make sure missing account ID is thrown
	_, err := client.CreateChallengeWidget(context.Background(), AccountIdentifier(""), ChallengeWidget{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, err := client.CreateChallengeWidget(context.Background(), AccountIdentifier(testAccountID), ChallengeWidget{
		Name: "blog.cloudflare.com login form",
		Type: "invisible",
		Domains: []string{
			"203.0.113.1",
			"cloudflare.com",
			"blog.example.com",
		},
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedChallengeWidget, out, "create challenge_widgets structs not equal")
	}
}

func TestChallengeWidgets_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
      "secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
      "created_on": "2014-01-01T05:20:00.123123Z",
      "modified_on": "2014-01-01T05:20:00.123123Z",
      "name": "blog.cloudflare.com login form",
      "domains": [
        "203.0.113.1",
        "cloudflare.com",
        "blog.example.com"
      ],
      "type": "invisible"
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 20,
    "count": 1,
    "total_count": 2000
  }
}`)
	})

	_, _, err := client.ListChallengeWidget(context.Background(), AccountIdentifier(""), ListChallengeWidgetRequest{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	out, results, err := client.ListChallengeWidget(context.Background(), AccountIdentifier(testAccountID), ListChallengeWidgetRequest{})
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(out), "expected 1 challenge_widgets")
		assert.Equal(t, 20, results.PerPage, "expected 20 per page")
		assert.Equal(t, expectedChallengeWidget, out[0], "list challenge_widgets structs not equal")
	}
}

func TestChallengeWidgets_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets/"+testChallengeWidgetSiteKey, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
    "secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
    "created_on": "2014-01-01T05:20:00.123123Z",
    "modified_on": "2014-01-01T05:20:00.123123Z",
    "name": "blog.cloudflare.com login form",
    "domains": [
      "203.0.113.1",
      "cloudflare.com",
      "blog.example.com"
    ],
    "type": "invisible"
  }
}`)
	})

	_, err := client.GetChallengeWidget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.GetChallengeWidget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingSiteKey, err)
	}

	out, err := client.GetChallengeWidget(context.Background(), AccountIdentifier(testAccountID), testChallengeWidgetSiteKey)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedChallengeWidget, out, "get challenge_widgets structs not equal")
	}
}

func TestChallengeWidgets_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets/"+testChallengeWidgetSiteKey, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
    "secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
    "created_on": "2014-01-01T05:20:00.123123Z",
    "modified_on": "2014-01-01T05:20:00.123123Z",
    "name": "blog.cloudflare.com login form",
    "domains": [
      "203.0.113.1",
      "cloudflare.com",
      "blog.example.com"
    ],
    "type": "invisible"
  }
}`)
	})

	_, err := client.UpdateChallengeWidget(context.Background(), AccountIdentifier(""), ChallengeWidget{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.UpdateChallengeWidget(context.Background(), AccountIdentifier(testAccountID), ChallengeWidget{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingSiteKey, err)
	}

	out, err := client.UpdateChallengeWidget(context.Background(), AccountIdentifier(testAccountID), ChallengeWidget{
		SiteKey: testChallengeWidgetSiteKey,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedChallengeWidget, out, "update challenge_widgets structs not equal")
	}
}

func TestChallengeWidgets_RotateSecret(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets/"+testChallengeWidgetSiteKey+"/rotate_secret", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
    "secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
    "created_on": "2014-01-01T05:20:00.123123Z",
    "modified_on": "2014-01-01T05:20:00.123123Z",
    "name": "blog.cloudflare.com login form",
    "domains": [
      "203.0.113.1",
      "cloudflare.com",
      "blog.example.com"
    ],
    "type": "invisible"
  }
}`)
	})

	// Make sure missing account ID is thrown
	_, err := client.RotateChallengeWidget(context.Background(), AccountIdentifier(""), RotateChallengeWidgetRequest{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	_, err = client.RotateChallengeWidget(context.Background(), AccountIdentifier(testAccountID), RotateChallengeWidgetRequest{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingSiteKey, err)
	}

	out, err := client.RotateChallengeWidget(context.Background(), AccountIdentifier(testAccountID), RotateChallengeWidgetRequest{SiteKey: testChallengeWidgetSiteKey})
	if assert.NoError(t, err) {
		assert.Equal(t, expectedChallengeWidget, out, "rotate challenge_widgets structs not equal")
	}
}

func TestChallengeWidgets_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/"+testAccountID+"/challenges/widgets/"+testChallengeWidgetSiteKey, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `
{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "sitekey": "0x4AAF00AAAABn0R22HWm-YUc",
    "secret": "0x4AAF00AAAABn0R22HWm098HVBjhdsYUc",
    "created_on": "2014-01-01T05:20:00.123123Z",
    "modified_on": "2014-01-01T05:20:00.123123Z",
    "name": "blog.cloudflare.com login form",
    "domains": [
      "203.0.113.1",
      "cloudflare.com",
      "blog.example.com"
    ],
    "type": "invisible"
  }
}`)
	})

	// Make sure missing account ID is thrown
	err := client.DeleteChallengeWidget(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.DeleteChallengeWidget(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingSiteKey, err)
	}

	err = client.DeleteChallengeWidget(context.Background(), AccountIdentifier(testAccountID), testChallengeWidgetSiteKey)
	assert.NoError(t, err)
}
