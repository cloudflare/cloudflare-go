// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *LogControlRetentionFlagService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *LogControlRetentionFlagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlRetentionFlagGetResponseEnvelope
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

func (r logControlRetentionFlagNewResponseJSON) RawJSON() string {
	return r.raw
}

type LogControlRetentionFlagGetResponse struct {
	Flag bool                                   `json:"flag"`
	JSON logControlRetentionFlagGetResponseJSON `json:"-"`
}

// logControlRetentionFlagGetResponseJSON contains the JSON metadata for the struct
// [LogControlRetentionFlagGetResponse]
type logControlRetentionFlagGetResponseJSON struct {
	Flag        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlRetentionFlagGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r logControlRetentionFlagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r logControlRetentionFlagNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r logControlRetentionFlagNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogControlRetentionFlagNewResponseEnvelopeSuccess bool

const (
	LogControlRetentionFlagNewResponseEnvelopeSuccessTrue LogControlRetentionFlagNewResponseEnvelopeSuccess = true
)

type LogControlRetentionFlagGetResponseEnvelope struct {
	Errors   []LogControlRetentionFlagGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlRetentionFlagGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlRetentionFlagGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LogControlRetentionFlagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlRetentionFlagGetResponseEnvelopeJSON    `json:"-"`
}

// logControlRetentionFlagGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [LogControlRetentionFlagGetResponseEnvelope]
type logControlRetentionFlagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlRetentionFlagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogControlRetentionFlagGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    logControlRetentionFlagGetResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlRetentionFlagGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LogControlRetentionFlagGetResponseEnvelopeErrors]
type logControlRetentionFlagGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlRetentionFlagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LogControlRetentionFlagGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    logControlRetentionFlagGetResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlRetentionFlagGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LogControlRetentionFlagGetResponseEnvelopeMessages]
type logControlRetentionFlagGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlRetentionFlagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlRetentionFlagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogControlRetentionFlagGetResponseEnvelopeSuccess bool

const (
	LogControlRetentionFlagGetResponseEnvelopeSuccessTrue LogControlRetentionFlagGetResponseEnvelopeSuccess = true
)
