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
func (r *MnmRuleAdvertisementService) Update(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleAdvertisementUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleAdvertisementUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v/advertisement", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmRuleAdvertisementUpdateResponse struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                   `json:"automatic_advertisement,required,nullable"`
	JSON                   mnmRuleAdvertisementUpdateResponseJSON `json:"-"`
}

// mnmRuleAdvertisementUpdateResponseJSON contains the JSON metadata for the struct
// [MnmRuleAdvertisementUpdateResponse]
type mnmRuleAdvertisementUpdateResponseJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MnmRuleAdvertisementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementUpdateResponseEnvelope struct {
	Errors   []MnmRuleAdvertisementUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleAdvertisementUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleAdvertisementUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleAdvertisementUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleAdvertisementUpdateResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleAdvertisementUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [MnmRuleAdvertisementUpdateResponseEnvelope]
type mnmRuleAdvertisementUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    mnmRuleAdvertisementUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleAdvertisementUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MnmRuleAdvertisementUpdateResponseEnvelopeErrors]
type mnmRuleAdvertisementUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleAdvertisementUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    mnmRuleAdvertisementUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleAdvertisementUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MnmRuleAdvertisementUpdateResponseEnvelopeMessages]
type mnmRuleAdvertisementUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleAdvertisementUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleAdvertisementUpdateResponseEnvelopeSuccess bool

const (
	MnmRuleAdvertisementUpdateResponseEnvelopeSuccessTrue MnmRuleAdvertisementUpdateResponseEnvelopeSuccess = true
)
