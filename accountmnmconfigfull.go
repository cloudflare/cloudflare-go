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

// AccountMnmConfigFullService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMnmConfigFullService]
// method instead.
type AccountMnmConfigFullService struct {
	Options []option.RequestOption
}

// NewAccountMnmConfigFullService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmConfigFullService(opts ...option.RequestOption) (r *AccountMnmConfigFullService) {
	r = &AccountMnmConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *AccountMnmConfigFullService) MagicNetworkMonitoringConfigurationListRulesAndAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse struct {
	Errors   []AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON    `json:"-"`
}

// accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse]
type accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseError struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseError]
type accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessage struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessage]
type accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                                    `json:"name,required"`
	RouterIPs []string                                                                                                  `json:"router_ips,required"`
	JSON      accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResultJSON `json:"-"`
}

// accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResult]
type accountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseSuccess bool

const (
	AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseSuccessTrue AccountMnmConfigFullMagicNetworkMonitoringConfigurationListRulesAndAccountConfigurationResponseSuccess = true
)
