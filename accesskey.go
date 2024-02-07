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
	Rotates *AccessKeyRotateService
}

// NewAccessKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessKeyService(opts ...option.RequestOption) (r *AccessKeyService) {
	r = &AccessKeyService{}
	r.Options = opts
	r.Rotates = NewAccessKeyRotateService(opts...)
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccessKeyService) AccessKeyConfigurationGetTheAccessKeyConfiguration(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccessKeyService) AccessKeyConfigurationUpdateTheAccessKeyConfiguration(ctx context.Context, identifier string, body AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams, opts ...option.RequestOption) (res *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseUnknown] or
// [shared.UnionString].
type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse interface {
	ImplementsAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseUnknown]
// or [shared.UnionString].
type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse interface {
	ImplementsAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelope struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                             `json:"days_until_next_rotation"`
	Errors                []AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrors `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                             `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessages `json:"messages"`
	Result            AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeSuccess `json:"success"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeJSON    `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelope]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeJSON struct {
	DaysUntilNextRotation   apijson.Field
	Errors                  apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	Messages                apijson.Field
	Result                  apijson.Field
	Success                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrors]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessages]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeSuccess bool

const (
	AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeSuccessTrue AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseEnvelopeSuccess = true
)

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelope struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                                `json:"days_until_next_rotation"`
	Errors                []AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrors `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                                `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessages `json:"messages"`
	Result            AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeSuccess `json:"success"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeJSON    `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelope]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeJSON struct {
	DaysUntilNextRotation   apijson.Field
	Errors                  apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	Messages                apijson.Field
	Result                  apijson.Field
	Success                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrors]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessages]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeSuccess bool

const (
	AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeSuccessTrue AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseEnvelopeSuccess = true
)
