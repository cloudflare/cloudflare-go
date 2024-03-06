// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// StreamKeyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewStreamKeyService] method instead.
type StreamKeyService struct {
	Options []option.RequestOption
}

// NewStreamKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamKeyService(opts ...option.RequestOption) (r *StreamKeyService) {
	r = &StreamKeyService{}
	r.Options = opts
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *StreamKeyService) New(ctx context.Context, body StreamKeyNewParams, opts ...option.RequestOption) (res *StreamKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *StreamKeyService) Delete(ctx context.Context, identifier string, body StreamKeyDeleteParams, opts ...option.RequestOption) (res *StreamKeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *StreamKeyService) Get(ctx context.Context, query StreamKeyGetParams, opts ...option.RequestOption) (res *[]StreamKeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamKeyNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string                   `json:"pem"`
	JSON streamKeyNewResponseJSON `json:"-"`
}

// streamKeyNewResponseJSON contains the JSON metadata for the struct
// [StreamKeyNewResponse]
type streamKeyNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [StreamKeyDeleteResponseUnknown] or [shared.UnionString].
type StreamKeyDeleteResponse interface {
	ImplementsStreamKeyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamKeyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamKeyGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                `json:"created" format:"date-time"`
	JSON    streamKeyGetResponseJSON `json:"-"`
}

// streamKeyGetResponseJSON contains the JSON metadata for the struct
// [StreamKeyGetResponse]
type streamKeyGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyGetResponseJSON) RawJSON() string {
	return r.raw
}

type StreamKeyNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamKeyNewResponseEnvelope struct {
	Errors   []StreamKeyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamKeyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyNewResponseEnvelopeJSON    `json:"-"`
}

// streamKeyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamKeyNewResponseEnvelope]
type streamKeyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamKeyNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    streamKeyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamKeyNewResponseEnvelopeErrors]
type streamKeyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamKeyNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    streamKeyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamKeyNewResponseEnvelopeMessages]
type streamKeyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamKeyNewResponseEnvelopeSuccess bool

const (
	StreamKeyNewResponseEnvelopeSuccessTrue StreamKeyNewResponseEnvelopeSuccess = true
)

type StreamKeyDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamKeyDeleteResponseEnvelope struct {
	Errors   []StreamKeyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamKeyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamKeyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamKeyDeleteResponseEnvelope]
type streamKeyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamKeyDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    streamKeyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamKeyDeleteResponseEnvelopeErrors]
type streamKeyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamKeyDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    streamKeyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamKeyDeleteResponseEnvelopeMessages]
type streamKeyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamKeyDeleteResponseEnvelopeSuccess bool

const (
	StreamKeyDeleteResponseEnvelopeSuccessTrue StreamKeyDeleteResponseEnvelopeSuccess = true
)

type StreamKeyGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamKeyGetResponseEnvelope struct {
	Errors   []StreamKeyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamKeyGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyGetResponseEnvelopeJSON    `json:"-"`
}

// streamKeyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamKeyGetResponseEnvelope]
type streamKeyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamKeyGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    streamKeyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamKeyGetResponseEnvelopeErrors]
type streamKeyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamKeyGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    streamKeyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamKeyGetResponseEnvelopeMessages]
type streamKeyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamKeyGetResponseEnvelopeSuccess bool

const (
	StreamKeyGetResponseEnvelopeSuccessTrue StreamKeyGetResponseEnvelopeSuccess = true
)
