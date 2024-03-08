// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// StreamDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamDownloadService] method
// instead.
type StreamDownloadService struct {
	Options []option.RequestOption
}

// NewStreamDownloadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamDownloadService(opts ...option.RequestOption) (r *StreamDownloadService) {
	r = &StreamDownloadService{}
	r.Options = opts
	return
}

// Creates a download for a video when a video is ready to view.
func (r *StreamDownloadService) New(ctx context.Context, identifier string, body StreamDownloadNewParams, opts ...option.RequestOption) (res *StreamDownloadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the downloads for a video.
func (r *StreamDownloadService) Delete(ctx context.Context, identifier string, body StreamDownloadDeleteParams, opts ...option.RequestOption) (res *StreamDownloadDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the downloads created for a video.
func (r *StreamDownloadService) Get(ctx context.Context, identifier string, query StreamDownloadGetParams, opts ...option.RequestOption) (res *StreamDownloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamDownloadNewResponseUnknown] or [shared.UnionString].
type StreamDownloadNewResponse interface {
	ImplementsStreamDownloadNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamDownloadDeleteResponseUnknown] or
// [shared.UnionString].
type StreamDownloadDeleteResponse interface {
	ImplementsStreamDownloadDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamDownloadGetResponseUnknown] or [shared.UnionString].
type StreamDownloadGetResponse interface {
	ImplementsStreamDownloadGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamDownloadNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamDownloadNewResponseEnvelope struct {
	Errors   []StreamDownloadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadNewResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDownloadNewResponseEnvelope]
type streamDownloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    streamDownloadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamDownloadNewResponseEnvelopeErrors]
type streamDownloadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamDownloadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamDownloadNewResponseEnvelopeMessages]
type streamDownloadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamDownloadNewResponseEnvelopeSuccess bool

const (
	StreamDownloadNewResponseEnvelopeSuccessTrue StreamDownloadNewResponseEnvelopeSuccess = true
)

type StreamDownloadDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamDownloadDeleteResponseEnvelope struct {
	Errors   []StreamDownloadDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamDownloadDeleteResponseEnvelope]
type streamDownloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeErrors]
type streamDownloadDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeMessages]
type streamDownloadDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamDownloadDeleteResponseEnvelopeSuccess bool

const (
	StreamDownloadDeleteResponseEnvelopeSuccessTrue StreamDownloadDeleteResponseEnvelopeSuccess = true
)

type StreamDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamDownloadGetResponseEnvelope struct {
	Errors   []StreamDownloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadGetResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadGetResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDownloadGetResponseEnvelope]
type streamDownloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    streamDownloadGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamDownloadGetResponseEnvelopeErrors]
type streamDownloadGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamDownloadGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamDownloadGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamDownloadGetResponseEnvelopeMessages]
type streamDownloadGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDownloadGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamDownloadGetResponseEnvelopeSuccess bool

const (
	StreamDownloadGetResponseEnvelopeSuccessTrue StreamDownloadGetResponseEnvelopeSuccess = true
)
