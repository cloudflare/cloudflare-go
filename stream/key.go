// File generated from our OpenAPI spec by Stainless.

package stream

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// KeyService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *KeyService) New(ctx context.Context, body KeyNewParams, opts ...option.RequestOption) (res *StreamKeys, err error) {
	opts = append(r.Options[:], opts...)
	var env KeyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *KeyService) Delete(ctx context.Context, identifier string, body KeyDeleteParams, opts ...option.RequestOption) (res *KeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KeyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *KeyService) Get(ctx context.Context, query KeyGetParams, opts ...option.RequestOption) (res *[]KeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KeyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamKeys struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string         `json:"pem"`
	JSON streamKeysJSON `json:"-"`
}

// streamKeysJSON contains the JSON metadata for the struct [StreamKeys]
type streamKeysJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeys) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamKeysJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [stream.KeyDeleteResponseUnknown] or [shared.UnionString].
type KeyDeleteResponse interface {
	ImplementsStreamKeyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type KeyGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time          `json:"created" format:"date-time"`
	JSON    keyGetResponseJSON `json:"-"`
}

// keyGetResponseJSON contains the JSON metadata for the struct [KeyGetResponse]
type keyGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseJSON) RawJSON() string {
	return r.raw
}

type KeyNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KeyNewResponseEnvelope struct {
	Errors   []KeyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamKeys                       `json:"result,required"`
	// Whether the API call was successful
	Success KeyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    keyNewResponseEnvelopeJSON    `json:"-"`
}

// keyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelope]
type keyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    keyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// keyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelopeErrors]
type keyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    keyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// keyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelopeMessages]
type keyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyNewResponseEnvelopeSuccess bool

const (
	KeyNewResponseEnvelopeSuccessTrue KeyNewResponseEnvelopeSuccess = true
)

type KeyDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KeyDeleteResponseEnvelope struct {
	Errors   []KeyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   KeyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    keyDeleteResponseEnvelopeJSON    `json:"-"`
}

// keyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelope]
type keyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    keyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// keyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelopeErrors]
type keyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    keyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// keyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelopeMessages]
type keyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyDeleteResponseEnvelopeSuccess bool

const (
	KeyDeleteResponseEnvelopeSuccessTrue KeyDeleteResponseEnvelopeSuccess = true
)

type KeyGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type KeyGetResponseEnvelope struct {
	Errors   []KeyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []KeyGetResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success KeyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    keyGetResponseEnvelopeJSON    `json:"-"`
}

// keyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyGetResponseEnvelope]
type keyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeyGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    keyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// keyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [KeyGetResponseEnvelopeErrors]
type keyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeyGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    keyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// keyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [KeyGetResponseEnvelopeMessages]
type keyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyGetResponseEnvelopeSuccess bool

const (
	KeyGetResponseEnvelopeSuccessTrue KeyGetResponseEnvelopeSuccess = true
)
