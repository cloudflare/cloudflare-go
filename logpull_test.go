package cloudflare

import (
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLogpullRetentionFlag(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		"errors": [],
		"messages": [],
		"result": {
			"flag": true
		},
		"success": true
	}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logs/control/retention/flag", handler)
	want := &LogpullRetentionConfiguration{Flag: true}

	actual, err := client.GetLogpullRetentionFlag(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestSetLogpullRetentionFlag(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		"errors": [],
		"messages": [],
		"result": {
			"flag": false
		},
		"success": true
	}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logs/control/retention/flag", handler)
	want := &LogpullRetentionConfiguration{Flag: false}

	actual, err := client.SetLogpullRetentionFlag(context.Background(), testZoneID, false)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestLogpullFields(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
		"foo": "bar",
		"field_1": "value_1",
		"field_2": "value_2"
	}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logs/received/fields", handler)

	actual, err := client.LogpullFields(context.Background(), testZoneID)
	if assert.NoError(t, err) {
		assert.Equal(t, map[string]string{
			"foo":     "bar",
			"field_1": "value_1",
			"field_2": "value_2",
		}, actual)
	}
}

func TestLogpullReceived(t *testing.T) {
	setup()
	defer teardown()

	var (
		want = []string{
			`{"field1":"value1", "field2":"value2"}`,
			`{"field11":"value11", "field22":"value22"}`,
		}
		fields       = []string{"field1", "field2"}
		count  int64 = 10
		sample       = 0.1
		start        = time.Unix(60, 0)
		end          = time.Unix(120, 0)
	)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, r.URL.Query().Get("count"), "10")
		assert.Equal(t, r.URL.Query().Get("sample"), "0.1")
		assert.Equal(t, r.URL.Query().Get("fields"), strings.Join(fields, ","))
		assert.Equal(t, r.URL.Query().Get("start"), fmt.Sprintf("%d", start.UnixNano()))
		assert.Equal(t, r.URL.Query().Get("end"), fmt.Sprintf("%d", end.UnixNano()))

		w.Header().Set("Content-Encoding", "gzip")
		wg := gzip.NewWriter(w)
		for _, line := range want {
			fmt.Fprintln(wg, line)
		}
		assert.NoError(t, wg.Close())
	}

	mux.HandleFunc("/zones/"+testZoneID+"/logs/received", handler)

	l, err := client.LogpullReceived(context.Background(), testZoneID, start, end, LogpullReceivedOption{
		Fields: fields,
		Count:  &count,
		Sample: &sample,
	})
	if !assert.NoError(t, err) {
		return
	}
	defer l.Close()
	var actual []string
	for l.Next() {
		actual = append(actual, string(l.Line()))
		assert.NoError(t, l.Err())
		f, err := l.Fields()
		if assert.NoError(t, err) {
			assert.Len(t, f, 2)
		}
	}

	assert.Equal(t, want, actual)
}
