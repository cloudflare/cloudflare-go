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

// MNMConfigService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMNMConfigService] method instead.
type MNMConfigService struct {
	Options []option.RequestOption
	Fulls   *MNMConfigFullService
}

// NewMNMConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMNMConfigService(opts ...option.RequestOption) (r *MNMConfigService) {
	r = &MNMConfigService{}
	r.Options = opts
	r.Fulls = NewMNMConfigFullService(opts...)
	return
}

// Create a new network monitoring configuration.
func (r *MNMConfigService) New(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigNewResponseEnvelope
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
func (r *MNMConfigService) Update(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling and router IPs for account.
func (r *MNMConfigService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an existing network monitoring configuration.
func (r *MNMConfigService) Delete(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *MNMConfigService) Edit(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MNMConfigNewResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                   `json:"name,required"`
	RouterIPs []string                 `json:"router_ips,required"`
	JSON      mnmConfigNewResponseJSON `json:"-"`
}

// mnmConfigNewResponseJSON contains the JSON metadata for the struct
// [MNMConfigNewResponse]
type mnmConfigNewResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigUpdateResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                      `json:"name,required"`
	RouterIPs []string                    `json:"router_ips,required"`
	JSON      mnmConfigUpdateResponseJSON `json:"-"`
}

// mnmConfigUpdateResponseJSON contains the JSON metadata for the struct
// [MNMConfigUpdateResponse]
type mnmConfigUpdateResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigListResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                    `json:"name,required"`
	RouterIPs []string                  `json:"router_ips,required"`
	JSON      mnmConfigListResponseJSON `json:"-"`
}

// mnmConfigListResponseJSON contains the JSON metadata for the struct
// [MNMConfigListResponse]
type mnmConfigListResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigDeleteResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                      `json:"name,required"`
	RouterIPs []string                    `json:"router_ips,required"`
	JSON      mnmConfigDeleteResponseJSON `json:"-"`
}

// mnmConfigDeleteResponseJSON contains the JSON metadata for the struct
// [MNMConfigDeleteResponse]
type mnmConfigDeleteResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigEditResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                    `json:"name,required"`
	RouterIPs []string                  `json:"router_ips,required"`
	JSON      mnmConfigEditResponseJSON `json:"-"`
}

// mnmConfigEditResponseJSON contains the JSON metadata for the struct
// [MNMConfigEditResponse]
type mnmConfigEditResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigNewResponseEnvelope struct {
	Errors   []MNMConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigNewResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigNewResponseEnvelope]
type mnmConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    mnmConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MNMConfigNewResponseEnvelopeErrors]
type mnmConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    mnmConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigNewResponseEnvelopeMessages]
type mnmConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigNewResponseEnvelopeSuccess bool

const (
	MNMConfigNewResponseEnvelopeSuccessTrue MNMConfigNewResponseEnvelopeSuccess = true
)

type MNMConfigUpdateResponseEnvelope struct {
	Errors   []MNMConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigUpdateResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigUpdateResponseEnvelope]
type mnmConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMConfigUpdateResponseEnvelopeErrors]
type mnmConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigUpdateResponseEnvelopeMessages]
type mnmConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigUpdateResponseEnvelopeSuccess bool

const (
	MNMConfigUpdateResponseEnvelopeSuccessTrue MNMConfigUpdateResponseEnvelopeSuccess = true
)

type MNMConfigListResponseEnvelope struct {
	Errors   []MNMConfigListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigListResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigListResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigListResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigListResponseEnvelope]
type mnmConfigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmConfigListResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMConfigListResponseEnvelopeErrors]
type mnmConfigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigListResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigListResponseEnvelopeMessages]
type mnmConfigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigListResponseEnvelopeSuccess bool

const (
	MNMConfigListResponseEnvelopeSuccessTrue MNMConfigListResponseEnvelopeSuccess = true
)

type MNMConfigDeleteResponseEnvelope struct {
	Errors   []MNMConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigDeleteResponseEnvelope]
type mnmConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMConfigDeleteResponseEnvelopeErrors]
type mnmConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigDeleteResponseEnvelopeMessages]
type mnmConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigDeleteResponseEnvelopeSuccess bool

const (
	MNMConfigDeleteResponseEnvelopeSuccessTrue MNMConfigDeleteResponseEnvelopeSuccess = true
)

type MNMConfigEditResponseEnvelope struct {
	Errors   []MNMConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigEditResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigEditResponseEnvelope]
type mnmConfigEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigEditResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmConfigEditResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMConfigEditResponseEnvelopeErrors]
type mnmConfigEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigEditResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigEditResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigEditResponseEnvelopeMessages]
type mnmConfigEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigEditResponseEnvelopeSuccess bool

const (
	MNMConfigEditResponseEnvelopeSuccessTrue MNMConfigEditResponseEnvelopeSuccess = true
)
