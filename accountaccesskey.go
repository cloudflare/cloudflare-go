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
func (r *AccountAccessKeyService) AccessKeyConfigurationGetTheAccessKeyConfiguration(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *KeysSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccountAccessKeyService) AccessKeyConfigurationUpdateTheAccessKeyConfiguration(ctx context.Context, identifier interface{}, body AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams, opts ...option.RequestOption) (res *KeysSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/keys", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type KeysSingleResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                   `json:"days_until_next_rotation"`
	Errors                []KeysSingleResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                   `json:"last_key_rotation_at" format:"date-time"`
	Messages          []KeysSingleResponseMessage `json:"messages"`
	Result            KeysSingleResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success KeysSingleResponseSuccess `json:"success"`
	JSON    keysSingleResponseJSON    `json:"-"`
}

// keysSingleResponseJSON contains the JSON metadata for the struct
// [KeysSingleResponse]
type keysSingleResponseJSON struct {
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

func (r *KeysSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeysSingleResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    keysSingleResponseErrorJSON `json:"-"`
}

// keysSingleResponseErrorJSON contains the JSON metadata for the struct
// [KeysSingleResponseError]
type keysSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeysSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type KeysSingleResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    keysSingleResponseMessageJSON `json:"-"`
}

// keysSingleResponseMessageJSON contains the JSON metadata for the struct
// [KeysSingleResponseMessage]
type keysSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeysSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [KeysSingleResponseResultObject] or [shared.UnionString].
type KeysSingleResponseResult interface {
	ImplementsKeysSingleResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeysSingleResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type KeysSingleResponseSuccess bool

const (
	KeysSingleResponseSuccessTrue KeysSingleResponseSuccess = true
)

type AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccountAccessKeyAccessKeyConfigurationUpdateTheAccessKeyConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
