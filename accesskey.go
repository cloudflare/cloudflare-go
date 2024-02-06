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
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccessKeyService) AccessKeyConfigurationUpdateTheAccessKeyConfiguration(ctx context.Context, identifier string, body AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams, opts ...option.RequestOption) (res *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                    `json:"days_until_next_rotation"`
	Errors                []AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                    `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage `json:"messages"`
	Result            AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess `json:"success"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON    `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON contains
// the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON struct {
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

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage]
type accessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResultUnknown]
// or [shared.UnionString].
type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult interface {
	ImplementsAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess bool

const (
	AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccessTrue AccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess = true
)

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                       `json:"days_until_next_rotation"`
	Errors                []AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                       `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage `json:"messages"`
	Result            AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess `json:"success"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON    `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON struct {
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

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON `json:"-"`
}

// accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage]
type accessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResultUnknown]
// or [shared.UnionString].
type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult interface {
	ImplementsAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess bool

const (
	AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccessTrue AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess = true
)

type AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
