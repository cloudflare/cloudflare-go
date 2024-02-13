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

// MnmConfigService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMnmConfigService] method instead.
type MnmConfigService struct {
	Options []option.RequestOption
	Fulls   *MnmConfigFullService
}

// NewMnmConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMnmConfigService(opts ...option.RequestOption) (r *MnmConfigService) {
	r = &MnmConfigService{}
	r.Options = opts
	r.Fulls = NewMnmConfigFullService(opts...)
	return
}

// Delete an existing network monitoring configuration.
func (r *MnmConfigService) Delete(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new network monitoring configuration.
func (r *MnmConfigService) MagicNetworkMonitoringConfigurationNewAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling and router IPs for account.
func (r *MnmConfigService) MagicNetworkMonitoringConfigurationListAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *MnmConfigService) MagicNetworkMonitoringConfigurationUpdateAccountConfigurationFields(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *MnmConfigService) MagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmConfigDeleteResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                      `json:"name,required"`
	RouterIPs []string                    `json:"router_ips,required"`
	JSON      mnmConfigDeleteResponseJSON `json:"-"`
}

// mnmConfigDeleteResponseJSON contains the JSON metadata for the struct
// [MnmConfigDeleteResponse]
type mnmConfigDeleteResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                          `json:"name,required"`
	RouterIPs []string                                                                        `json:"router_ips,required"`
	JSON      mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse]
type mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                           `json:"name,required"`
	RouterIPs []string                                                                         `json:"router_ips,required"`
	JSON      mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse]
type mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                   `json:"name,required"`
	RouterIPs []string                                                                                 `json:"router_ips,required"`
	JSON      mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                                                                     `json:"name,required"`
	RouterIPs []string                                                                                   `json:"router_ips,required"`
	JSON      mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigDeleteResponseEnvelope struct {
	Errors   []MnmConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmConfigDeleteResponseEnvelope]
type mnmConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmConfigDeleteResponseEnvelopeErrors]
type mnmConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmConfigDeleteResponseEnvelopeMessages]
type mnmConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigDeleteResponseEnvelopeSuccess bool

const (
	MnmConfigDeleteResponseEnvelopeSuccessTrue MnmConfigDeleteResponseEnvelopeSuccess = true
)

type MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelope struct {
	Errors   []MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelope]
type mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrors]
type mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessages]
type mnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeSuccess bool

const (
	MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeSuccessTrue MnmConfigMagicNetworkMonitoringConfigurationNewAccountConfigurationResponseEnvelopeSuccess = true
)

type MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelope struct {
	Errors   []MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelope]
type mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrors]
type mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessages]
type mnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeSuccess bool

const (
	MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeSuccessTrue MnmConfigMagicNetworkMonitoringConfigurationListAccountConfigurationResponseEnvelopeSuccess = true
)

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelope struct {
	Errors   []MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelope]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrors]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessages]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeSuccess bool

const (
	MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeSuccessTrue MnmConfigMagicNetworkMonitoringConfigurationUpdateAccountConfigurationFieldsResponseEnvelopeSuccess = true
)

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelope struct {
	Errors   []MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelope]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrors]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessages]
type mnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeSuccess bool

const (
	MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeSuccessTrue MnmConfigMagicNetworkMonitoringConfigurationUpdateAnEntireAccountConfigurationResponseEnvelopeSuccess = true
)
