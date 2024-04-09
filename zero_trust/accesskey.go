// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// Updates the Access key rotation settings for an account.
func (r *AccessKeyService) Update(ctx context.Context, identifier string, body AccessKeyUpdateParams, opts ...option.RequestOption) (res *AccessKeyUpdateResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccessKeyService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Perfoms a key rotation for an account.
func (r *AccessKeyService) Rotate(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyRotateResponseUnion, err error) {
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

// Union satisfied by [zero_trust.AccessKeyUpdateResponseUnknown] or
// [shared.UnionString].
type AccessKeyUpdateResponseUnion interface {
	ImplementsZeroTrustAccessKeyUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [zero_trust.AccessKeyGetResponseUnknown] or
// [shared.UnionString].
type AccessKeyGetResponseUnion interface {
	ImplementsZeroTrustAccessKeyGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [zero_trust.AccessKeyRotateResponseUnknown] or
// [shared.UnionString].
type AccessKeyRotateResponseUnion interface {
	ImplementsZeroTrustAccessKeyRotateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyRotateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessKeyUpdateParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccessKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessKeyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   AccessKeyUpdateResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success AccessKeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                           `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyUpdateResponseEnvelopeJSON `json:"-"`
}

// accessKeyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyUpdateResponseEnvelope]
type accessKeyUpdateResponseEnvelopeJSON struct {
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

func (r *AccessKeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessKeyUpdateResponseEnvelopeSuccess bool

const (
	AccessKeyUpdateResponseEnvelopeSuccessTrue AccessKeyUpdateResponseEnvelopeSuccess = true
)

func (r AccessKeyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessKeyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   AccessKeyGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success AccessKeyGetResponseEnvelopeSuccess `json:"success,required"`
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                        `json:"last_key_rotation_at" format:"date-time"`
	JSON              accessKeyGetResponseEnvelopeJSON `json:"-"`
}

// accessKeyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessKeyGetResponseEnvelope]
type accessKeyGetResponseEnvelopeJSON struct {
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

func (r *AccessKeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessKeyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessKeyGetResponseEnvelopeSuccess bool

const (
	AccessKeyGetResponseEnvelopeSuccessTrue AccessKeyGetResponseEnvelopeSuccess = true
)

func (r AccessKeyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessKeyRotateResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   AccessKeyRotateResponseUnion `json:"result,required"`
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

func (r accessKeyRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessKeyRotateResponseEnvelopeSuccess bool

const (
	AccessKeyRotateResponseEnvelopeSuccessTrue AccessKeyRotateResponseEnvelopeSuccess = true
)

func (r AccessKeyRotateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessKeyRotateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
