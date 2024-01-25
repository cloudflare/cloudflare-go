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

// AccountAccessKeyRotateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessKeyRotateService]
// method instead.
type AccountAccessKeyRotateService struct {
	Options []option.RequestOption
}

// NewAccountAccessKeyRotateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessKeyRotateService(opts ...option.RequestOption) (r *AccountAccessKeyRotateService) {
	r = &AccountAccessKeyRotateService{}
	r.Options = opts
	return
}

// Perfoms a key rotation for an account.
func (r *AccountAccessKeyRotateService) AccessKeyConfigurationRotateAccessKeys(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                     `json:"days_until_next_rotation"`
	Errors                []AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseError `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                     `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessage `json:"messages"`
	Result            AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseSuccess `json:"success"`
	JSON    accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseJSON    `json:"-"`
}

// accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse]
type accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseJSON struct {
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

func (r *AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseErrorJSON `json:"-"`
}

// accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseError]
type accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessageJSON `json:"-"`
}

// accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessage]
type accountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseResultUnknown]
// or [shared.UnionString].
type AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseResult interface {
	ImplementsAccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseSuccess bool

const (
	AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseSuccessTrue AccountAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseSuccess = true
)
