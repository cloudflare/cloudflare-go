// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *StreamKeyService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamKeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *StreamKeyService) StreamSigningKeysNewSigningKeys(ctx context.Context, accountID string, opts ...option.RequestOption) (res *StreamKeyStreamSigningKeysNewSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *StreamKeyService) StreamSigningKeysListSigningKeys(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]StreamKeyStreamSigningKeysListSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyStreamSigningKeysListSigningKeysResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type StreamKeyStreamSigningKeysNewSigningKeysResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string                                               `json:"pem"`
	JSON streamKeyStreamSigningKeysNewSigningKeysResponseJSON `json:"-"`
}

// streamKeyStreamSigningKeysNewSigningKeysResponseJSON contains the JSON metadata
// for the struct [StreamKeyStreamSigningKeysNewSigningKeysResponse]
type streamKeyStreamSigningKeysNewSigningKeysResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysNewSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyStreamSigningKeysListSigningKeysResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                                             `json:"created" format:"date-time"`
	JSON    streamKeyStreamSigningKeysListSigningKeysResponseJSON `json:"-"`
}

// streamKeyStreamSigningKeysListSigningKeysResponseJSON contains the JSON metadata
// for the struct [StreamKeyStreamSigningKeysListSigningKeysResponse]
type streamKeyStreamSigningKeysListSigningKeysResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysListSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful
type StreamKeyDeleteResponseEnvelopeSuccess bool

const (
	StreamKeyDeleteResponseEnvelopeSuccessTrue StreamKeyDeleteResponseEnvelopeSuccess = true
)

type StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelope struct {
	Errors   []StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamKeyStreamSigningKeysNewSigningKeysResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeJSON    `json:"-"`
}

// streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelope]
type streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrors struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrors]
type streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessages struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessages]
type streamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeSuccess bool

const (
	StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeSuccessTrue StreamKeyStreamSigningKeysNewSigningKeysResponseEnvelopeSuccess = true
)

type StreamKeyStreamSigningKeysListSigningKeysResponseEnvelope struct {
	Errors   []StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamKeyStreamSigningKeysListSigningKeysResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeJSON    `json:"-"`
}

// streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [StreamKeyStreamSigningKeysListSigningKeysResponseEnvelope]
type streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysListSigningKeysResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrors]
type streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessages]
type streamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeSuccess bool

const (
	StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeSuccessTrue StreamKeyStreamSigningKeysListSigningKeysResponseEnvelopeSuccess = true
)
