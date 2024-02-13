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

// MnmConfigFullService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMnmConfigFullService] method
// instead.
type MnmConfigFullService struct {
	Options []option.RequestOption
}

// NewMnmConfigFullService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMnmConfigFullService(opts ...option.RequestOption) (r *MnmConfigFullService) {
	r = &MnmConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *MnmConfigFullService) MagicNetworkMonitoringConfigurationListRulesAndAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                       `json:"name,required"`
	RouterIPs []string                                                                                     `json:"router_ips,required"`
	JSON      mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON `json:"-"`
}

// mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse]
type mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelope struct {
	Errors   []MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelope]
type mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrors]
type mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessages]
type mnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeSuccess bool

const (
	MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeSuccessTrue MnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseEnvelopeSuccess = true
)
