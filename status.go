// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// StatusService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewStatusService] method instead.
type StatusService struct {
	Options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r *StatusService) {
	r = &StatusService{}
	r.Options = opts
	return
}

// API status check
func (r *StatusService) Get(ctx context.Context, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StatusGetResponse struct {
	Message string                `json:"message"`
	JSON    statusGetResponseJSON `json:"-"`
}

// statusGetResponseJSON contains the JSON metadata for the struct
// [StatusGetResponse]
type statusGetResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
