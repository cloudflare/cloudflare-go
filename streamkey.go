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

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *StreamKeyService) New(ctx context.Context, accountID string, opts ...option.RequestOption) (res *StreamKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *StreamKeyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]StreamKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamKeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type StreamKeyListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                 `json:"created" format:"date-time"`
	JSON    streamKeyListResponseJSON `json:"-"`
}

// streamKeyListResponseJSON contains the JSON metadata for the struct
// [StreamKeyListResponse]
type streamKeyListResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful
type StreamKeyNewResponseEnvelopeSuccess bool

const (
	StreamKeyNewResponseEnvelopeSuccessTrue StreamKeyNewResponseEnvelopeSuccess = true
)

type StreamKeyListResponseEnvelope struct {
	Errors   []StreamKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []StreamKeyListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success StreamKeyListResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamKeyListResponseEnvelopeJSON    `json:"-"`
}

// streamKeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamKeyListResponseEnvelope]
type streamKeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    streamKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamKeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamKeyListResponseEnvelopeErrors]
type streamKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamKeyListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    streamKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamKeyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamKeyListResponseEnvelopeMessages]
type streamKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamKeyListResponseEnvelopeSuccess bool

const (
	StreamKeyListResponseEnvelopeSuccessTrue StreamKeyListResponseEnvelopeSuccess = true
)

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
