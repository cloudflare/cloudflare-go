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

// AccountAccessKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessKeyService] method
// instead.
type AccountAccessKeyService struct {
	Options []option.RequestOption
	Rotates *AccountAccessKeyRotateService
}

// NewAccountAccessKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessKeyService(opts ...option.RequestOption) (r *AccountAccessKeyService) {
	r = &AccountAccessKeyService{}
	r.Options = opts
	r.Rotates = NewAccountAccessKeyRotateService(opts...)
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccountAccessKeyService) AccessKeyConfigurationGetTheAccessKeyConfiguration(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccountAccessKeyService) AccessKeyConfigurationUpdateTheAccessKeyConfiguration(ctx context.Context, identifier string, body AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams, opts ...option.RequestOption) (res *AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                           `json:"days_until_next_rotation"`
	Errors                []AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                           `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage `json:"messages"`
	Result            AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess `json:"success"`
	JSON    accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON    `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse]
type accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseJSON struct {
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

func (r *AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError]
type accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage]
type accountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResultUnknown]
// or [shared.UnionString].
type AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult interface {
	ImplementsAccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess bool

const (
	AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccessTrue AccountAccessKeyAccessKeyConfigurationGetTheAccessKeyConfigurationResponseSuccess = true
)

type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                              `json:"days_until_next_rotation"`
	Errors                []AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                              `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage `json:"messages"`
	Result            AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess `json:"success"`
	JSON    accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON    `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse]
type accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseJSON struct {
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

func (r *AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError]
type accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON `json:"-"`
}

// accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage]
type accountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResultUnknown]
// or [shared.UnionString].
type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult interface {
	ImplementsAccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess bool

const (
	AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccessTrue AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationResponseSuccess = true
)

type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
