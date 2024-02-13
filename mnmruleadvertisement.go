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

// MnmRuleAdvertisementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMnmRuleAdvertisementService]
// method instead.
type MnmRuleAdvertisementService struct {
	Options []option.RequestOption
}

// NewMnmRuleAdvertisementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMnmRuleAdvertisementService(opts ...option.RequestOption) (r *MnmRuleAdvertisementService) {
	r = &MnmRuleAdvertisementService{}
	r.Options = opts
	return
}

// Update advertisement for rule.
func (r *MnmRuleAdvertisementService) MagicNetworkMonitoringRulesUpdateAdvertisementForRule(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v/advertisement", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponse struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                                                                  `json:"automatic_advertisement,required,nullable"`
	JSON                   mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseJSON `json:"-"`
}

// mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseJSON
// contains the JSON metadata for the struct
// [MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponse]
type mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelope struct {
	Errors   []MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelope]
type mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrors struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrors]
type mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessages struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessages]
type mnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeSuccess bool

const (
	MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeSuccessTrue MnmRuleAdvertisementMagicNetworkMonitoringRulesUpdateAdvertisementForRuleResponseEnvelopeSuccess = true
)
