// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
	Full    *ConfigFullService
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	r.Full = NewConfigFullService(opts...)
	return
}

// Create a new network monitoring configuration.
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfig, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfig, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an existing network monitoring configuration.
func (r *ConfigService) Delete(ctx context.Context, params ConfigDeleteParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfig, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *ConfigService) Edit(ctx context.Context, params ConfigEditParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfig, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mnm/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling and router IPs for account.
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfig, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/mnm/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringConfig struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                           `json:"name,required"`
	RouterIPs []string                         `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigJSON `json:"-"`
}

// magicNetworkMonitoringConfigJSON contains the JSON metadata for the struct
// [MagicNetworkMonitoringConfig]
type magicNetworkMonitoringConfigJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigJSON) RawJSON() string {
	return r.raw
}

type ConfigNewParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConfigNewResponseEnvelope struct {
	Errors   []ConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfig        `json:"result,required"`
	// Whether the API call was successful
	Success ConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    configNewResponseEnvelopeJSON    `json:"-"`
}

// configNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelope]
type configNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    configNewResponseEnvelopeErrorsJSON `json:"-"`
}

// configNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelopeErrors]
type configNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    configNewResponseEnvelopeMessagesJSON `json:"-"`
}

// configNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelopeMessages]
type configNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigNewResponseEnvelopeSuccess bool

const (
	ConfigNewResponseEnvelopeSuccessTrue ConfigNewResponseEnvelopeSuccess = true
)

func (r ConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigUpdateParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []ConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfig           `json:"result,required"`
	// Whether the API call was successful
	Success ConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    configUpdateResponseEnvelopeJSON    `json:"-"`
}

// configUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelope]
type configUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    configUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// configUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigUpdateResponseEnvelopeErrors]
type configUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    configUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// configUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigUpdateResponseEnvelopeMessages]
type configUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigUpdateResponseEnvelopeSuccess bool

const (
	ConfigUpdateResponseEnvelopeSuccessTrue ConfigUpdateResponseEnvelopeSuccess = true
)

func (r ConfigUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ConfigDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []ConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfig           `json:"result,required"`
	// Whether the API call was successful
	Success ConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    configDeleteResponseEnvelopeJSON    `json:"-"`
}

// configDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseEnvelope]
type configDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    configDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// configDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseEnvelopeErrors]
type configDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    configDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// configDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConfigDeleteResponseEnvelopeMessages]
type configDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigDeleteResponseEnvelopeSuccess bool

const (
	ConfigDeleteResponseEnvelopeSuccessTrue ConfigDeleteResponseEnvelopeSuccess = true
)

func (r ConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigEditParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ConfigEditResponseEnvelope struct {
	Errors   []ConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfig         `json:"result,required"`
	// Whether the API call was successful
	Success ConfigEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    configEditResponseEnvelopeJSON    `json:"-"`
}

// configEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigEditResponseEnvelope]
type configEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigEditResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    configEditResponseEnvelopeErrorsJSON `json:"-"`
}

// configEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigEditResponseEnvelopeErrors]
type configEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigEditResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    configEditResponseEnvelopeMessagesJSON `json:"-"`
}

// configEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ConfigEditResponseEnvelopeMessages]
type configEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigEditResponseEnvelopeSuccess bool

const (
	ConfigEditResponseEnvelopeSuccessTrue ConfigEditResponseEnvelopeSuccess = true
)

func (r ConfigEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConfigGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []ConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfig        `json:"result,required"`
	// Whether the API call was successful
	Success ConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configGetResponseEnvelopeJSON    `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    configGetResponseEnvelopeErrorsJSON `json:"-"`
}

// configGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeErrors]
type configGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    configGetResponseEnvelopeMessagesJSON `json:"-"`
}

// configGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelopeMessages]
type configGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigGetResponseEnvelopeSuccess bool

const (
	ConfigGetResponseEnvelopeSuccessTrue ConfigGetResponseEnvelopeSuccess = true
)

func (r ConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
