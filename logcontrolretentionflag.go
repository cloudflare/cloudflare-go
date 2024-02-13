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

// LogControlRetentionFlagService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLogControlRetentionFlagService] method instead.
type LogControlRetentionFlagService struct {
	Options []option.RequestOption
}

// NewLogControlRetentionFlagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogControlRetentionFlagService(opts ...option.RequestOption) (r *LogControlRetentionFlagService) {
	r = &LogControlRetentionFlagService{}
	r.Options = opts
	return
}

// Gets log retention flag for Logpull API.
func (r *LogControlRetentionFlagService) LogsReceivedGetLogRetentionFlag(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelope
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates log retention flag for Logpull API.
func (r *LogControlRetentionFlagService) LogsReceivedUpdateLogRetentionFlag(ctx context.Context, zoneIdentifier string, body LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams, opts ...option.RequestOption) (res *LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelope
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse struct {
	Flag bool                                                               `json:"flag"`
	JSON logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON contains the
// JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse]
type logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse struct {
	Flag bool                                                                  `json:"flag"`
	JSON logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON contains
// the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse]
type logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelope struct {
	Errors   []LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeJSON    `json:"-"`
}

// logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelope]
type logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrors]
type logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessages]
type logControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeSuccess bool

const (
	LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeSuccessTrue LogControlRetentionFlagLogsReceivedGetLogRetentionFlagResponseEnvelopeSuccess = true
)

type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelope struct {
	Errors   []LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeJSON    `json:"-"`
}

// logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelope]
type logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrors]
type logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessages]
type logControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeSuccess bool

const (
	LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeSuccessTrue LogControlRetentionFlagLogsReceivedUpdateLogRetentionFlagResponseEnvelopeSuccess = true
)
