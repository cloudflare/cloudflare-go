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

// Updates log retention flag for Logpull API.
func (r *LogControlRetentionFlagService) New(ctx context.Context, zoneIdentifier string, body LogControlRetentionFlagNewParams, opts ...option.RequestOption) (res *LogControlRetentionFlagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlRetentionFlagNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/logs/control/retention/flag", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type LogControlRetentionFlagNewResponse struct {
	Flag bool                                   `json:"flag"`
	JSON logControlRetentionFlagNewResponseJSON `json:"-"`
}

// logControlRetentionFlagNewResponseJSON contains the JSON metadata for the struct
// [LogControlRetentionFlagNewResponse]
type logControlRetentionFlagNewResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type LogControlRetentionFlagNewParams struct {
	// The log retention flag for Logpull API.
	Flag param.Field[bool] `json:"flag,required"`
}

func (r LogControlRetentionFlagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogControlRetentionFlagNewResponseEnvelope struct {
	Errors   []LogControlRetentionFlagNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlRetentionFlagNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlRetentionFlagNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LogControlRetentionFlagNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlRetentionFlagNewResponseEnvelopeJSON    `json:"-"`
}

// logControlRetentionFlagNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [LogControlRetentionFlagNewResponseEnvelope]
type logControlRetentionFlagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    logControlRetentionFlagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlRetentionFlagNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LogControlRetentionFlagNewResponseEnvelopeErrors]
type logControlRetentionFlagNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlRetentionFlagNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    logControlRetentionFlagNewResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlRetentionFlagNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LogControlRetentionFlagNewResponseEnvelopeMessages]
type logControlRetentionFlagNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlRetentionFlagNewResponseEnvelopeSuccess bool

const (
	LogControlRetentionFlagNewResponseEnvelopeSuccessTrue LogControlRetentionFlagNewResponseEnvelopeSuccess = true
)

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
