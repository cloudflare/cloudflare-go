// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogControlRetentionFlagService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneLogControlRetentionFlagService] method instead.
type ZoneLogControlRetentionFlagService struct {
	Options []option.RequestOption
}

// NewZoneLogControlRetentionFlagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLogControlRetentionFlagService(opts ...option.RequestOption) (r *ZoneLogControlRetentionFlagService) {
	r = &ZoneLogControlRetentionFlagService{}
	r.Options = opts
	return
}

// Gets log retention flag for Logpull API.
func (r *ZoneLogControlRetentionFlagService) LogsReceivedGetLogRetentionFlag(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *FlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates log retention flag for Logpull API.
func (r *ZoneLogControlRetentionFlagService) LogsReceivedUpdateLogRetentionFlag(ctx context.Context, zoneIdentifier string, body ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams, opts ...option.RequestOption) (res *FlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FlagResponse struct {
	Errors   []FlagResponseError   `json:"errors"`
	Messages []FlagResponseMessage `json:"messages"`
	Result   FlagResponseResult    `json:"result"`
	// Whether the API call was successful
	Success FlagResponseSuccess `json:"success"`
	JSON    flagResponseJSON    `json:"-"`
}

// flagResponseJSON contains the JSON metadata for the struct [FlagResponse]
type flagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FlagResponseError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    flagResponseErrorJSON `json:"-"`
}

// flagResponseErrorJSON contains the JSON metadata for the struct
// [FlagResponseError]
type flagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FlagResponseMessage struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    flagResponseMessageJSON `json:"-"`
}

// flagResponseMessageJSON contains the JSON metadata for the struct
// [FlagResponseMessage]
type flagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FlagResponseResult struct {
	Flag bool                   `json:"flag"`
	JSON flagResponseResultJSON `json:"-"`
}

// flagResponseResultJSON contains the JSON metadata for the struct
// [FlagResponseResult]
type flagResponseResultJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FlagResponseSuccess bool

const (
	FlagResponseSuccessTrue FlagResponseSuccess = true
)

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
