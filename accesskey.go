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

// AccessKeyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessKeyService] method instead.
type AccessKeyService struct {
	Options []option.RequestOption
}

// NewAccessKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessKeyService(opts ...option.RequestOption) (r *AccessKeyService) {
	r = &AccessKeyService{}
	r.Options = opts
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccessKeyService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccessKeyService) Replace(ctx context.Context, identifier string, body AccessKeyReplaceParams, opts ...option.RequestOption) (res *AccessKeyReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Perfoms a key rotation for an account.
func (r *AccessKeyService) Rotate(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyRotateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyRotateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessKeyListResponseUnknown] or [shared.UnionString].
type AccessKeyListResponse interface {
	ImplementsAccessKeyListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [AccessKeyReplaceResponseUnknown] or [shared.UnionString].
type AccessKeyReplaceResponse interface {
	ImplementsAccessKeyReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [AccessKeyRotateResponseUnknown] or [shared.UnionString].
type AccessKeyRotateResponse interface {
	ImplementsAccessKeyRotateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyRotateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessKeyListResponseEnvelope struct {
	Errors   []AccessKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessKeyListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessKeyListResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                         `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyListResponseEnvelopeJSON `json:"-"`
}

// accessKeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyListResponseEnvelope]
type accessKeyListResponseEnvelopeJSON struct {
	Errors                  apijson.Field
	Messages                apijson.Field
	Result                  apijson.Field
	Success                 apijson.Field
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessKeyListResponseEnvelopeErrors]
type accessKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyListResponseEnvelopeMessages]
type accessKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyListResponseEnvelopeSuccess bool

const (
	AccessKeyListResponseEnvelopeSuccessTrue AccessKeyListResponseEnvelopeSuccess = true
)

type AccessKeyReplaceParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccessKeyReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessKeyReplaceResponseEnvelope struct {
	Errors   []AccessKeyReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessKeyReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessKeyReplaceResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                            `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyReplaceResponseEnvelopeJSON `json:"-"`
}

// accessKeyReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyReplaceResponseEnvelope]
type accessKeyReplaceResponseEnvelopeJSON struct {
	Errors                  apijson.Field
	Messages                apijson.Field
	Result                  apijson.Field
	Success                 apijson.Field
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyReplaceResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessKeyReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessKeyReplaceResponseEnvelopeErrors]
type accessKeyReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyReplaceResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessKeyReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyReplaceResponseEnvelopeMessages]
type accessKeyReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyReplaceResponseEnvelopeSuccess bool

const (
	AccessKeyReplaceResponseEnvelopeSuccessTrue AccessKeyReplaceResponseEnvelopeSuccess = true
)

type AccessKeyRotateResponseEnvelope struct {
	Errors   []AccessKeyRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessKeyRotateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessKeyRotateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessKeyRotateResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                           `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyRotateResponseEnvelopeJSON `json:"-"`
}

// accessKeyRotateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyRotateResponseEnvelope]
type accessKeyRotateResponseEnvelopeJSON struct {
	Errors                  apijson.Field
	Messages                apijson.Field
	Result                  apijson.Field
	Success                 apijson.Field
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyRotateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessKeyRotateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyRotateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessKeyRotateResponseEnvelopeErrors]
type accessKeyRotateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyRotateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessKeyRotateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyRotateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessKeyRotateResponseEnvelopeMessages]
type accessKeyRotateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyRotateResponseEnvelopeSuccess bool

const (
	AccessKeyRotateResponseEnvelopeSuccessTrue AccessKeyRotateResponseEnvelopeSuccess = true
)
