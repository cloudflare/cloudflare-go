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
func (r *ZoneLogControlRetentionFlagService) LogsReceivedGetLogRetentionFlag(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates log retention flag for Logpull API.
func (r *ZoneLogControlRetentionFlagService) LogsReceivedUpdateLogRetentionFlag(ctx context.Context, zoneIdentifier string, body ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams, opts ...option.RequestOption) (res *ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse struct {
	Errors   []ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseError   `json:"errors"`
	Messages []ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessage `json:"messages"`
	Result   ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseSuccess `json:"success"`
	JSON    zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON    `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON contains
// the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse]
type zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseErrorJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseError]
type zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessageJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessage]
type zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResult struct {
	Flag bool                                                                         `json:"flag"`
	JSON zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResultJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResult]
type zoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResultJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseSuccess bool

const (
	ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseSuccessTrue ZoneLogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseSuccess = true
)

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse struct {
	Errors   []ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseError   `json:"errors"`
	Messages []ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessage `json:"messages"`
	Result   ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseSuccess `json:"success"`
	JSON    zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON    `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse]
type zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseErrorJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseError]
type zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessageJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessage]
type zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResult struct {
	Flag bool                                                                            `json:"flag"`
	JSON zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResultJSON `json:"-"`
}

// zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResult]
type zoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResultJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseSuccess bool

const (
	ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseSuccessTrue ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseSuccess = true
)

type ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r ZoneLogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
