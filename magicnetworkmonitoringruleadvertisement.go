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

// MagicNetworkMonitoringRuleAdvertisementService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMagicNetworkMonitoringRuleAdvertisementService] method instead.
type MagicNetworkMonitoringRuleAdvertisementService struct {
	Options []option.RequestOption
}

// NewMagicNetworkMonitoringRuleAdvertisementService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMagicNetworkMonitoringRuleAdvertisementService(opts ...option.RequestOption) (r *MagicNetworkMonitoringRuleAdvertisementService) {
	r = &MagicNetworkMonitoringRuleAdvertisementService{}
	r.Options = opts
	return
}

// Update advertisement for rule.
func (r *MagicNetworkMonitoringRuleAdvertisementService) Edit(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MagicVisibilityMNMRuleAdvertisable, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v/advertisement", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicVisibilityMNMRuleAdvertisable struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                   `json:"automatic_advertisement,required,nullable"`
	JSON                   magicVisibilityMNMRuleAdvertisableJSON `json:"-"`
}

// magicVisibilityMNMRuleAdvertisableJSON contains the JSON metadata for the struct
// [MagicVisibilityMNMRuleAdvertisable]
type magicVisibilityMNMRuleAdvertisableJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MagicVisibilityMNMRuleAdvertisable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicVisibilityMNMRuleAdvertisable                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelope]
type magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrors]
type magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessages]
type magicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleAdvertisementEditResponseEnvelopeSuccess = true
)
