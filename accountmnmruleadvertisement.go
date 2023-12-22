// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMnmRuleAdvertisementService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountMnmRuleAdvertisementService] method instead.
type AccountMnmRuleAdvertisementService struct {
	Options []option.RequestOption
}

// NewAccountMnmRuleAdvertisementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMnmRuleAdvertisementService(opts ...option.RequestOption) (r *AccountMnmRuleAdvertisementService) {
	r = &AccountMnmRuleAdvertisementService{}
	r.Options = opts
	return
}

// Update advertisement for rule.
func (r *AccountMnmRuleAdvertisementService) MagicNetworkMonitoringRulesUpdateAdvertisementForRule(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleAdvertisementSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v/advertisement", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type MnmRuleAdvertisementSingleResponse struct {
	Errors   []MnmRuleAdvertisementSingleResponseError   `json:"errors"`
	Messages []MnmRuleAdvertisementSingleResponseMessage `json:"messages"`
	Result   MnmRuleAdvertisementSingleResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success MnmRuleAdvertisementSingleResponseSuccess `json:"success"`
	JSON    mnmRuleAdvertisementSingleResponseJSON    `json:"-"`
}

// mnmRuleAdvertisementSingleResponseJSON contains the JSON metadata for the struct
// [MnmRuleAdvertisementSingleResponse]
type mnmRuleAdvertisementSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementSingleResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmRuleAdvertisementSingleResponseErrorJSON `json:"-"`
}

// mnmRuleAdvertisementSingleResponseErrorJSON contains the JSON metadata for the
// struct [MnmRuleAdvertisementSingleResponseError]
type mnmRuleAdvertisementSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementSingleResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    mnmRuleAdvertisementSingleResponseMessageJSON `json:"-"`
}

// mnmRuleAdvertisementSingleResponseMessageJSON contains the JSON metadata for the
// struct [MnmRuleAdvertisementSingleResponseMessage]
type mnmRuleAdvertisementSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementSingleResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                         `json:"automatic_advertisement,required,nullable"`
	JSON                   mnmRuleAdvertisementSingleResponseResultJSON `json:"-"`
}

// mnmRuleAdvertisementSingleResponseResultJSON contains the JSON metadata for the
// struct [MnmRuleAdvertisementSingleResponseResult]
type mnmRuleAdvertisementSingleResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MnmRuleAdvertisementSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleAdvertisementSingleResponseSuccess bool

const (
	MnmRuleAdvertisementSingleResponseSuccessTrue MnmRuleAdvertisementSingleResponseSuccess = true
)
