package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testBucketName = "example-bucket"

func TestR2_CreateBucket(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/r2/buckets", testAccountID), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {}
}`)
	})

	err := client.CreateR2Bucket(context.Background(), AccountIdentifier(""), CreateR2BucketParameters{})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.CreateR2Bucket(context.Background(), AccountIdentifier(testAccountID), CreateR2BucketParameters{Name: ""})
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingBucketName, err)
	}

	err = client.CreateR2Bucket(context.Background(), AccountIdentifier(testAccountID), CreateR2BucketParameters{Name: testBucketName})
	assert.NoError(t, err)
}

func TestR2_DeleteBucket(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/accounts/%s/r2/buckets/%s", testAccountID, testBucketName), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {}
}`)
	})

	err := client.DeleteR2Bucket(context.Background(), AccountIdentifier(""), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingAccountID, err)
	}

	err = client.DeleteR2Bucket(context.Background(), AccountIdentifier(testAccountID), "")
	if assert.Error(t, err) {
		assert.Equal(t, ErrMissingBucketName, err)
	}

	err = client.DeleteR2Bucket(context.Background(), AccountIdentifier(testAccountID), "example-bucket")
	assert.NoError(t, err)
}
