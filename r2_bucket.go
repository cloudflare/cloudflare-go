package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrMissingBucketName = errors.New("require bucket name missing")
)

// R2Bucket defines a container for objects stored in R2 Storage.
type R2Bucket struct {
	Name         string `json:"name"`
	CreationDate string `json:"creation_date,omitempty"`
}

// R2Buckets represents the map of buckets response from
// the R2 buckets endpoint.
type R2Buckets struct {
	Buckets []R2Bucket `json:"buckets"`
}

// R2BucketListResponse represents the response from the list
// R2 buckets endpoint.
type R2BucketListResponse struct {
	Result R2Buckets `json:"result"`
	Response
}

// R2Buckets Lists R2 buckets.
func (api *API) R2Buckets(ctx context.Context, rc *ResourceContainer) ([]R2Bucket, error) {
	if rc.Identifier == "" {
		return []R2Bucket{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []R2Bucket{}, err
	}

	var r2BucketListResponse R2BucketListResponse
	err = json.Unmarshal(res, &r2BucketListResponse)
	if err != nil {
		return []R2Bucket{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r2BucketListResponse.Result.Buckets, nil
}

// CreateR2Bucket Creates a new R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-create-bucket
func (api *API) CreateR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if bucketName == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)

	return err
}

// DeleteR2Bucket Deletes an existing R2 bucket.
//
// API reference: https://api.cloudflare.com/#r2-bucket-delete-bucket
func (api *API) DeleteR2Bucket(ctx context.Context, rc *ResourceContainer, bucketName string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if bucketName == "" {
		return ErrMissingBucketName
	}

	uri := fmt.Sprintf("/accounts/%s/r2/buckets/%s", rc.Identifier, bucketName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	return err
}
