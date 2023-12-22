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

// AccountMnmConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMnmConfigService] method
// instead.
type AccountMnmConfigService struct {
	Options []option.RequestOption
	Fulls   *AccountMnmConfigFullService
}

// NewAccountMnmConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmConfigService(opts ...option.RequestOption) (r *AccountMnmConfigService) {
	r = &AccountMnmConfigService{}
	r.Options = opts
	r.Fulls = NewAccountMnmConfigFullService(opts...)
	return
}

// Delete an existing network monitoring configuration.
func (r *AccountMnmConfigService) Delete(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new network monitoring configuration.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationNewAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists default sampling and router IPs for account.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationListAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update fields in an existing network monitoring configuration.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationUpdateAccountConfigurationFields(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *AccountMnmConfigService) MagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type MnmConfigSingleResponse struct {
	Errors   []MnmConfigSingleResponseError   `json:"errors"`
	Messages []MnmConfigSingleResponseMessage `json:"messages"`
	Result   MnmConfigSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success MnmConfigSingleResponseSuccess `json:"success"`
	JSON    mnmConfigSingleResponseJSON    `json:"-"`
}

// mnmConfigSingleResponseJSON contains the JSON metadata for the struct
// [MnmConfigSingleResponse]
type mnmConfigSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigSingleResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    mnmConfigSingleResponseErrorJSON `json:"-"`
}

// mnmConfigSingleResponseErrorJSON contains the JSON metadata for the struct
// [MnmConfigSingleResponseError]
type mnmConfigSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigSingleResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    mnmConfigSingleResponseMessageJSON `json:"-"`
}

// mnmConfigSingleResponseMessageJSON contains the JSON metadata for the struct
// [MnmConfigSingleResponseMessage]
type mnmConfigSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigSingleResponseResult struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                            `json:"name,required"`
	RouterIPs []string                          `json:"router_ips,required"`
	JSON      mnmConfigSingleResponseResultJSON `json:"-"`
}

// mnmConfigSingleResponseResultJSON contains the JSON metadata for the struct
// [MnmConfigSingleResponseResult]
type mnmConfigSingleResponseResultJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigSingleResponseSuccess bool

const (
	MnmConfigSingleResponseSuccessTrue MnmConfigSingleResponseSuccess = true
)
