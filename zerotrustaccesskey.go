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

// ZeroTrustAccessKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessKeyService] method
// instead.
type ZeroTrustAccessKeyService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessKeyService(opts ...option.RequestOption) (r *ZeroTrustAccessKeyService) {
	r = &ZeroTrustAccessKeyService{}
	r.Options = opts
	return
}

// Updates the Access key rotation settings for an account.
func (r *ZeroTrustAccessKeyService) Update(ctx context.Context, identifier string, body ZeroTrustAccessKeyUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessKeyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessKeyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the Access key rotation settings for an account.
func (r *ZeroTrustAccessKeyService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZeroTrustAccessKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessKeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Perfoms a key rotation for an account.
func (r *ZeroTrustAccessKeyService) Rotate(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZeroTrustAccessKeyRotateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessKeyRotateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustAccessKeyUpdateResponseUnknown] or
// [shared.UnionString].
type ZeroTrustAccessKeyUpdateResponse interface {
	ImplementsZeroTrustAccessKeyUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessKeyUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [ZeroTrustAccessKeyListResponseUnknown] or
// [shared.UnionString].
type ZeroTrustAccessKeyListResponse interface {
	ImplementsZeroTrustAccessKeyListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessKeyListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [ZeroTrustAccessKeyRotateResponseUnknown] or
// [shared.UnionString].
type ZeroTrustAccessKeyRotateResponse interface {
	ImplementsZeroTrustAccessKeyRotateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessKeyRotateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustAccessKeyUpdateParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r ZeroTrustAccessKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessKeyUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessKeyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessKeyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessKeyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessKeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                    `json:"last_key_rotation_at" format:"date-time"`
	JSON              zeroTrustAccessKeyUpdateResponseEnvelopeJSON `json:"-"`
}

// zeroTrustAccessKeyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessKeyUpdateResponseEnvelope]
type zeroTrustAccessKeyUpdateResponseEnvelopeJSON struct {
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

func (r *ZeroTrustAccessKeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessKeyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessKeyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessKeyUpdateResponseEnvelopeErrors]
type zeroTrustAccessKeyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessKeyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessKeyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessKeyUpdateResponseEnvelopeMessages]
type zeroTrustAccessKeyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessKeyUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessKeyUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessKeyUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessKeyListResponseEnvelope struct {
	Errors   []ZeroTrustAccessKeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessKeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessKeyListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessKeyListResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                  `json:"last_key_rotation_at" format:"date-time"`
	JSON              zeroTrustAccessKeyListResponseEnvelopeJSON `json:"-"`
}

// zeroTrustAccessKeyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessKeyListResponseEnvelope]
type zeroTrustAccessKeyListResponseEnvelopeJSON struct {
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

func (r *ZeroTrustAccessKeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustAccessKeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessKeyListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessKeyListResponseEnvelopeErrors]
type zeroTrustAccessKeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessKeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessKeyListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessKeyListResponseEnvelopeMessages]
type zeroTrustAccessKeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessKeyListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessKeyListResponseEnvelopeSuccessTrue ZeroTrustAccessKeyListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessKeyRotateResponseEnvelope struct {
	Errors   []ZeroTrustAccessKeyRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessKeyRotateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessKeyRotateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessKeyRotateResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                    `json:"last_key_rotation_at" format:"date-time"`
	JSON              zeroTrustAccessKeyRotateResponseEnvelopeJSON `json:"-"`
}

// zeroTrustAccessKeyRotateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessKeyRotateResponseEnvelope]
type zeroTrustAccessKeyRotateResponseEnvelopeJSON struct {
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

func (r *ZeroTrustAccessKeyRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyRotateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessKeyRotateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessKeyRotateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessKeyRotateResponseEnvelopeErrors]
type zeroTrustAccessKeyRotateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessKeyRotateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessKeyRotateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessKeyRotateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessKeyRotateResponseEnvelopeMessages]
type zeroTrustAccessKeyRotateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessKeyRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessKeyRotateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessKeyRotateResponseEnvelopeSuccessTrue ZeroTrustAccessKeyRotateResponseEnvelopeSuccess = true
)
