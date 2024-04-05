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
func (r *CaptionService) Update(ctx context.Context, identifier string, language string, params CaptionUpdateParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
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
func (r *CaptionService) Delete(ctx context.Context, identifier string, language string, params CaptionDeleteParams, opts ...option.RequestOption) (res *CaptionDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env CaptionDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", params.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *CaptionService) Get(ctx context.Context, identifier string, query CaptionGetParams, opts ...option.RequestOption) (res *[]Caption, err error) {
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

type Caption struct {
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string      `json:"language"`
	JSON     captionJSON `json:"-"`
}

// captionJSON contains the JSON metadata for the struct [Caption]
type captionJSON struct {
	Label       apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Caption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [stream.CaptionDeleteResponseUnknown],
// [stream.CaptionDeleteResponseArray] or [shared.UnionString].
type CaptionDeleteResponseUnion interface {
	ImplementsStreamCaptionDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CaptionDeleteResponseUnion)(nil)).Elem(),
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

func (r CaptionDeleteResponseArray) ImplementsStreamCaptionDeleteResponseUnion() {}

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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
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

// Whether the API call was successful
type CaptionUpdateResponseEnvelopeSuccess bool

const (
	CaptionUpdateResponseEnvelopeSuccessTrue CaptionUpdateResponseEnvelopeSuccess = true
)

func (r CaptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CaptionDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r CaptionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CaptionDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   CaptionDeleteResponseUnion                                `json:"result,required"`
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

// Whether the API call was successful
type CaptionDeleteResponseEnvelopeSuccess bool

const (
	CaptionDeleteResponseEnvelopeSuccessTrue CaptionDeleteResponseEnvelopeSuccess = true
)

func (r CaptionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CaptionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []Caption                                                 `json:"result,required"`
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

// Whether the API call was successful
type CaptionGetResponseEnvelopeSuccess bool

const (
	CaptionGetResponseEnvelopeSuccessTrue CaptionGetResponseEnvelopeSuccess = true
)

func (r CaptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
