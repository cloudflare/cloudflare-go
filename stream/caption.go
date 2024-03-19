// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// CaptionService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCaptionService] method instead.
type CaptionService struct {
	Options []option.RequestOption
}

// NewCaptionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCaptionService(opts ...option.RequestOption) (r *CaptionService) {
	r = &CaptionService{}
	r.Options = opts
	return
}

// Uploads the caption or subtitle file to the endpoint for a specific BCP47
// language. One caption or subtitle file per language is allowed.
func (r *CaptionService) Update(ctx context.Context, identifier string, language string, params CaptionUpdateParams, opts ...option.RequestOption) (res *CaptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CaptionUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", params.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes the captions or subtitles from a video.
func (r *CaptionService) Delete(ctx context.Context, identifier string, language string, body CaptionDeleteParams, opts ...option.RequestOption) (res *CaptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CaptionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", body.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *CaptionService) Get(ctx context.Context, identifier string, query CaptionGetParams, opts ...option.RequestOption) (res *[]StreamCaptions, err error) {
	opts = append(r.Options[:], opts...)
	var env CaptionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamCaptions struct {
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string             `json:"language"`
	JSON     streamCaptionsJSON `json:"-"`
}

// streamCaptionsJSON contains the JSON metadata for the struct [StreamCaptions]
type streamCaptionsJSON struct {
	Label       apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamCaptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamCaptionsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [stream.CaptionUpdateResponseUnknown] or
// [shared.UnionString].
type CaptionUpdateResponse interface {
	ImplementsStreamCaptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CaptionUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.CaptionDeleteResponseUnknown],
// [stream.CaptionDeleteResponseArray] or [shared.UnionString].
type CaptionDeleteResponse interface {
	ImplementsStreamCaptionDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CaptionDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CaptionDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CaptionDeleteResponseArray []interface{}

func (r CaptionDeleteResponseArray) ImplementsStreamCaptionDeleteResponse() {}

type CaptionUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r CaptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CaptionUpdateResponseEnvelope struct {
	Errors   []CaptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CaptionUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CaptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    captionUpdateResponseEnvelopeJSON    `json:"-"`
}

// captionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionUpdateResponseEnvelope]
type captionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    captionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// captionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CaptionUpdateResponseEnvelopeErrors]
type captionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    captionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// captionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CaptionUpdateResponseEnvelopeMessages]
type captionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CaptionUpdateResponseEnvelopeSuccess bool

const (
	CaptionUpdateResponseEnvelopeSuccessTrue CaptionUpdateResponseEnvelopeSuccess = true
)

type CaptionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionDeleteResponseEnvelope struct {
	Errors   []CaptionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CaptionDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CaptionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    captionDeleteResponseEnvelopeJSON    `json:"-"`
}

// captionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionDeleteResponseEnvelope]
type captionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    captionDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// captionDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CaptionDeleteResponseEnvelopeErrors]
type captionDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    captionDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// captionDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CaptionDeleteResponseEnvelopeMessages]
type captionDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CaptionDeleteResponseEnvelopeSuccess bool

const (
	CaptionDeleteResponseEnvelopeSuccessTrue CaptionDeleteResponseEnvelopeSuccess = true
)

type CaptionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionGetResponseEnvelope struct {
	Errors   []CaptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamCaptions                     `json:"result,required"`
	// Whether the API call was successful
	Success CaptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    captionGetResponseEnvelopeJSON    `json:"-"`
}

// captionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionGetResponseEnvelope]
type captionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    captionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// captionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CaptionGetResponseEnvelopeErrors]
type captionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    captionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// captionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CaptionGetResponseEnvelopeMessages]
type captionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CaptionGetResponseEnvelopeSuccess bool

const (
	CaptionGetResponseEnvelopeSuccessTrue CaptionGetResponseEnvelopeSuccess = true
)
