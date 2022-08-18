package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func buildTestPageShieldScript() PageShieldScript {
	addedAt, _ := time.Parse(time.RFC3339Nano, "2021-08-18T10:51:10.09615Z")
	firstSeenAt, _ := time.Parse(time.RFC3339, "2021-08-18T10:51:08Z")
	lastSeenAt, _ := time.Parse(time.RFC3339, "2021-09-02T09:57:54Z")
	fetchedAt, _ := time.Parse(time.RFC3339, "2021-09-02T10:17:54Z")
	return PageShieldScript{
		ScriptID:                "c9ef84a6bf5e47138c75d95e2f933e8f",
		ScriptURL:               "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
		AddedAt:                 &addedAt,
		FirstSeenAt:             &firstSeenAt,
		LastSeenAt:              &lastSeenAt,
		Host:                    "blog.cloudflare.com",
		DomainReportedMalicious: false,
		SeenOnFirst:             "blog.cloudflare.com/page",
		SeenOn:                  []string{"blog.cloudflare.com/page1", "blog.cloudflare.com/page2"},
		AppearsInCDNCGIPath:     false,
		Hash:                    "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		JsIntegrityScore:        10,
		FetchedAt:               &fetchedAt,
	}
}

func TestGetPageShieldStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/page_shield", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  			"success": true,
			"errors": [],
  			"messages": [],
  			"result": {}
		}
		`)
	})

	err := client.GetPageShieldStatus(context.Background(), ZoneIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	err = client.GetPageShieldStatus(context.Background(), ZoneIdentifier(testZoneID))
	assert.NoError(t, err)
}

func TestUpdatePageShieldStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/page_shield", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  			"success": true,
			"errors": [],
  			"messages": [],
  			"result": {}
		}
		`)
	})

	err := client.UpdatePageShieldStatus(context.Background(), ZoneIdentifier(""))
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	err = client.UpdatePageShieldStatus(context.Background(), ZoneIdentifier(testZoneID))
	assert.NoError(t, err)
}

func TestListPageShieldStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/page_shield/scripts", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    {
      "script_id": "c9ef84a6bf5e47138c75d95e2f933e8f",
      "script_url": "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
      "added_at": "2021-08-18T10:51:10.09615Z",
      "first_seen_at": "2021-08-18T10:51:08Z",
      "last_seen_at": "2021-09-02T09:57:54Z",
      "host": "blog.cloudflare.com",
      "domain_reported_malicious": false,
      "seen_on_first": "blog.cloudflare.com/page",
      "seen_on": [
        "blog.cloudflare.com/page1",
        "blog.cloudflare.com/page2"
      ],
      "appears_in_cdn_cgi_path": false,
      "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
      "js_integrity_score": 10,
      "fetched_at": "2021-09-02T10:17:54Z"
    }
  ],
  "result_info": {
    "page": 1,
    "per_page": 20,
    "count": 1,
    "total_count": 2000
  }
}
		`)
	})

	_, _, err := client.ListPageShieldScripts(context.Background(), ZoneIdentifier(""), ListPageShieldParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}

	result, resinfo, err := client.ListPageShieldScripts(context.Background(), ZoneIdentifier(testZoneID), ListPageShieldParameters{})
	if assert.NoError(t, err) {
		want := buildTestPageShieldScript()
		assert.Len(t, result, 1)
		assert.Equal(t, resinfo.Page, 1)
		assert.Equal(t, want, result[0])
	}
}

func TestGetPageShieldScript(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/zones/"+testZoneID+"/page_shield/scripts/c9ef84a6bf5e47138c75d95e2f933e8f", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "script_id": "c9ef84a6bf5e47138c75d95e2f933e8f",
  "script_url": "https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.6.0/js/bootstrap.min.js",
  "added_at": "2021-08-18T10:51:10.09615Z",
  "first_seen_at": "2021-08-18T10:51:08Z",
  "last_seen_at": "2021-09-02T09:57:54Z",
  "host": "blog.cloudflare.com",
  "domain_reported_malicious": false,
  "seen_on_first": "blog.cloudflare.com/page",
  "seen_on": [
    "blog.cloudflare.com/page1",
    "blog.cloudflare.com/page2"
  ],
  "appears_in_cdn_cgi_path": false,
  "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
  "js_integrity_score": 10,
  "fetched_at": "2021-09-02T10:17:54Z",
  "versions": [
    {
      "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b423",
      "js_integrity_score": 2,
      "fetched_at": "2021-08-18T10:51:08Z"
    }
  ]
}`)
	})

	_, err := client.GetPageShieldScript(context.Background(), ZoneIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingZoneID, err)
	}
	_, err = client.GetPageShieldScript(context.Background(), ZoneIdentifier(testZoneID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingScriptID, err)
	}

	res, err := client.GetPageShieldScript(context.Background(), ZoneIdentifier(testZoneID), "c9ef84a6bf5e47138c75d95e2f933e8f")
	if assert.NoError(t, err) {
		want := buildTestPageShieldScript()
		fetchedAt, _ := time.Parse(time.RFC3339, "2021-08-18T10:51:08Z")
		want.Versions = []PageShieldScriptVersions{
			{
				Hash:             "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b423",
				JsIntegrityScore: 2,
				FetchedAt:        &fetchedAt,
			},
		}
		assert.Equal(t, want, res)
	}
}
