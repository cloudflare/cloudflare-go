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

// MNMRuleAdvertisementService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMNMRuleAdvertisementService]
// method instead.
type MNMRuleAdvertisementService struct {
	Options []option.RequestOption
}

// NewMNMRuleAdvertisementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMNMRuleAdvertisementService(opts ...option.RequestOption) (r *MNMRuleAdvertisementService) {
	r = &MNMRuleAdvertisementService{}
	r.Options = opts
	return
}

// Update advertisement for rule.
func (r *MNMRuleAdvertisementService) Edit(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleAdvertisementEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleAdvertisementEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v/advertisement", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MNMRuleAdvertisementEditResponse struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                 `json:"automatic_advertisement,required,nullable"`
	JSON                   mnmRuleAdvertisementEditResponseJSON `json:"-"`
}

// mnmRuleAdvertisementEditResponseJSON contains the JSON metadata for the struct
// [MNMRuleAdvertisementEditResponse]
type mnmRuleAdvertisementEditResponseJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MNMRuleAdvertisementEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleAdvertisementEditResponseEnvelope struct {
	Errors   []MNMRuleAdvertisementEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleAdvertisementEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleAdvertisementEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleAdvertisementEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleAdvertisementEditResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleAdvertisementEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [MNMRuleAdvertisementEditResponseEnvelope]
type mnmRuleAdvertisementEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleAdvertisementEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleAdvertisementEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    mnmRuleAdvertisementEditResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleAdvertisementEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MNMRuleAdvertisementEditResponseEnvelopeErrors]
type mnmRuleAdvertisementEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleAdvertisementEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleAdvertisementEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    mnmRuleAdvertisementEditResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleAdvertisementEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MNMRuleAdvertisementEditResponseEnvelopeMessages]
type mnmRuleAdvertisementEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleAdvertisementEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleAdvertisementEditResponseEnvelopeSuccess bool

const (
	MNMRuleAdvertisementEditResponseEnvelopeSuccessTrue MNMRuleAdvertisementEditResponseEnvelopeSuccess = true
)
