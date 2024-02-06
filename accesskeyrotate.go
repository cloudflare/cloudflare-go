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

// AccessKeyRotateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessKeyRotateService] method
// instead.
type AccessKeyRotateService struct {
	Options []option.RequestOption
}

// NewAccessKeyRotateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessKeyRotateService(opts ...option.RequestOption) (r *AccessKeyRotateService) {
	r = &AccessKeyRotateService{}
	r.Options = opts
	return
}

// Perfoms a key rotation for an account.
func (r *AccessKeyRotateService) AccessKeyConfigurationRotateAccessKeys(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseUnknown] or
// [shared.UnionString].
type AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse interface {
	ImplementsAccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelope struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64                                                                       `json:"days_until_next_rotation"`
	Errors                []AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrors `json:"errors"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                                                                       `json:"last_key_rotation_at" format:"date-time"`
	Messages          []AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessages `json:"messages"`
	Result            AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeSuccess `json:"success"`
	JSON    accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeJSON    `json:"-"`
}

// accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelope]
type accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeJSON struct {
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

func (r *AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrorsJSON `json:"-"`
}

// accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrors]
type accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessagesJSON `json:"-"`
}

// accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessages]
type accessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeSuccess bool

const (
	AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeSuccessTrue AccessKeyRotateAccessKeyConfigurationRotateAccessKeysResponseEnvelopeSuccess = true
)
