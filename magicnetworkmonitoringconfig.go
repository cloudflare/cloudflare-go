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

// MagicNetworkMonitoringConfigService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewMagicNetworkMonitoringConfigService] method instead.
type MagicNetworkMonitoringConfigService struct {
	Options []option.RequestOption
	Full    *MagicNetworkMonitoringConfigFullService
}

// NewMagicNetworkMonitoringConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicNetworkMonitoringConfigService(opts ...option.RequestOption) (r *MagicNetworkMonitoringConfigService) {
	r = &MagicNetworkMonitoringConfigService{}
	r.Options = opts
	r.Full = NewMagicNetworkMonitoringConfigFullService(opts...)
	return
}

// Create a new network monitoring configuration.
func (r *MagicNetworkMonitoringConfigService) New(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *MagicNetworkMonitoringConfigService) Update(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an existing network monitoring configuration.
func (r *MagicNetworkMonitoringConfigService) Delete(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *MagicNetworkMonitoringConfigService) Edit(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling and router IPs for account.
func (r *MagicNetworkMonitoringConfigService) Get(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringConfigNewResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                      `json:"name,required"`
	RouterIPs []string                                    `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigNewResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigNewResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringConfigNewResponse]
type magicNetworkMonitoringConfigNewResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigUpdateResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                         `json:"name,required"`
	RouterIPs []string                                       `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigUpdateResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigUpdateResponseJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringConfigUpdateResponse]
type magicNetworkMonitoringConfigUpdateResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigDeleteResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                         `json:"name,required"`
	RouterIPs []string                                       `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigDeleteResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigDeleteResponseJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringConfigDeleteResponse]
type magicNetworkMonitoringConfigDeleteResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigEditResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                       `json:"name,required"`
	RouterIPs []string                                     `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigEditResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigEditResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringConfigEditResponse]
type magicNetworkMonitoringConfigEditResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigEditResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigGetResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                      `json:"name,required"`
	RouterIPs []string                                    `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigGetResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigGetResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringConfigGetResponse]
type magicNetworkMonitoringConfigGetResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigNewResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigNewResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringConfigNewResponseEnvelope]
type magicNetworkMonitoringConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicNetworkMonitoringConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigNewResponseEnvelopeErrors]
type magicNetworkMonitoringConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    magicNetworkMonitoringConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigNewResponseEnvelopeMessages]
type magicNetworkMonitoringConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigNewResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigNewResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigNewResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringConfigUpdateResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigUpdateResponseEnvelope]
type magicNetworkMonitoringConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicNetworkMonitoringConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigUpdateResponseEnvelopeErrors]
type magicNetworkMonitoringConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    magicNetworkMonitoringConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigUpdateResponseEnvelopeMessages]
type magicNetworkMonitoringConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigUpdateResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigUpdateResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigUpdateResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringConfigDeleteResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigDeleteResponseEnvelope]
type magicNetworkMonitoringConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicNetworkMonitoringConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigDeleteResponseEnvelopeErrors]
type magicNetworkMonitoringConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    magicNetworkMonitoringConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigDeleteResponseEnvelopeMessages]
type magicNetworkMonitoringConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigDeleteResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigDeleteResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigDeleteResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringConfigEditResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigEditResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringConfigEditResponseEnvelope]
type magicNetworkMonitoringConfigEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicNetworkMonitoringConfigEditResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigEditResponseEnvelopeErrors]
type magicNetworkMonitoringConfigEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicNetworkMonitoringConfigEditResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigEditResponseEnvelopeMessages]
type magicNetworkMonitoringConfigEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigEditResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigEditResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigEditResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringConfigGetResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigGetResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringConfigGetResponseEnvelope]
type magicNetworkMonitoringConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicNetworkMonitoringConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigGetResponseEnvelopeErrors]
type magicNetworkMonitoringConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    magicNetworkMonitoringConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigGetResponseEnvelopeMessages]
type magicNetworkMonitoringConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigGetResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigGetResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigGetResponseEnvelopeSuccess = true
)
