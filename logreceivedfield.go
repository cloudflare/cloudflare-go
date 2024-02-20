// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LogReceivedFieldService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogReceivedFieldService] method
// instead.
type LogReceivedFieldService struct {
	Options []option.RequestOption
}

// NewLogReceivedFieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogReceivedFieldService(opts ...option.RequestOption) (r *LogReceivedFieldService) {
	r = &LogReceivedFieldService{}
	r.Options = opts
	return
}

// Lists all fields available. The response is json object with key-value pairs,
// where keys are field names, and values are descriptions.
func (r *LogReceivedFieldService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *LogReceivedFieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/received/fields", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogReceivedFieldListResponse struct {
	Key  string                           `json:"key"`
	JSON logReceivedFieldListResponseJSON `json:"-"`
}

// logReceivedFieldListResponseJSON contains the JSON metadata for the struct
// [LogReceivedFieldListResponse]
type logReceivedFieldListResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogReceivedFieldListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
