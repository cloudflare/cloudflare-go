// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ThreatEventCronService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventCronService] method instead.
type ThreatEventCronService struct {
	Options []option.RequestOption
}

// NewThreatEventCronService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventCronService(opts ...option.RequestOption) (r *ThreatEventCronService) {
	r = &ThreatEventCronService{}
	r.Options = opts
	return
}

// Reads the last cron update time
func (r *ThreatEventCronService) New(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *ThreatEventCronNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/cron", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reads the last cron update time
func (r *ThreatEventCronService) List(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *ThreatEventCronListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/cron", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventCronNewResponse struct {
	ID     float64                        `json:"id,required"`
	Update string                         `json:"update,required"`
	JSON   threatEventCronNewResponseJSON `json:"-"`
}

// threatEventCronNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventCronNewResponse]
type threatEventCronNewResponseJSON struct {
	ID          apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCronNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCronNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCronListResponse struct {
	Update string                          `json:"update,required"`
	JSON   threatEventCronListResponseJSON `json:"-"`
}

// threatEventCronListResponseJSON contains the JSON metadata for the struct
// [ThreatEventCronListResponse]
type threatEventCronListResponseJSON struct {
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCronListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCronListResponseJSON) RawJSON() string {
	return r.raw
}
