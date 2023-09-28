package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type R2Object struct {
	Key            string            `json:"key"`
	Etag           string            `json:"etag"`
	LastModified   time.Time         `json:"last_modified"`
	Size           int64             `json:"size"`
	HTTPMetadata   map[string]any    `json:"http_metadata"`
	CustomMetadata map[string]string `json:"custom_metadata"`
}

type ListR2ObjectsParams struct {
	StartAfter string `url:"start_after,omitempty"`
	Prefix     string `url:"prefix,omitempty"`
	Cursor     string `url:"cursor,omitempty"`
	PerPage    int    `url:"per_page,omitempty"`
	Delimiter  string `url:"delimiter,omitempty"`
}

type R2ListObjectsResponse struct {
	Result     []R2Object `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// ListR2Objects lists R2 objects in a given bucket.
func (api *API) ListR2Objects(ctx context.Context, rc *ResourceContainer, bucket string, params ListR2ObjectsParams) ([]R2Object, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/r2/buckets/%s/objects", rc.Identifier, bucket), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	var resp R2ListObjectsResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return resp.Result, &resp.ResultInfo, nil
}
