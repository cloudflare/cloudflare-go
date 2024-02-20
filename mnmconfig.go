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

// Create a new network monitoring configuration.
func (r *MnmConfigService) New(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update fields in an existing network monitoring configuration.
func (r *MnmConfigService) Update(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists default sampling and router IPs for account.
func (r *MnmConfigService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Update an existing network monitoring configuration, requires the entire
// configuration to be updated at once.
func (r *MnmConfigService) Replace(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmConfigReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmConfigNewResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                   `json:"name,required"`
	RouterIPs []string                 `json:"router_ips,required"`
	JSON      mnmConfigNewResponseJSON `json:"-"`
}

// mnmConfigNewResponseJSON contains the JSON metadata for the struct
// [MnmConfigNewResponse]
type mnmConfigNewResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigUpdateResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                      `json:"name,required"`
	RouterIPs []string                    `json:"router_ips,required"`
	JSON      mnmConfigUpdateResponseJSON `json:"-"`
}

// mnmConfigUpdateResponseJSON contains the JSON metadata for the struct
// [MnmConfigUpdateResponse]
type mnmConfigUpdateResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigListResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                    `json:"name,required"`
	RouterIPs []string                  `json:"router_ips,required"`
	JSON      mnmConfigListResponseJSON `json:"-"`
}

// mnmConfigListResponseJSON contains the JSON metadata for the struct
// [MnmConfigListResponse]
type mnmConfigListResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type MnmConfigReplaceResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                       `json:"name,required"`
	RouterIPs []string                     `json:"router_ips,required"`
	JSON      mnmConfigReplaceResponseJSON `json:"-"`
}

// mnmConfigReplaceResponseJSON contains the JSON metadata for the struct
// [MnmConfigReplaceResponse]
type mnmConfigReplaceResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MnmConfigReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigNewResponseEnvelope struct {
	Errors   []MnmConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigNewResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmConfigNewResponseEnvelope]
type mnmConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    mnmConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MnmConfigNewResponseEnvelopeErrors]
type mnmConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    mnmConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmConfigNewResponseEnvelopeMessages]
type mnmConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigNewResponseEnvelopeSuccess bool

const (
	MnmConfigNewResponseEnvelopeSuccessTrue MnmConfigNewResponseEnvelopeSuccess = true
)

type MnmConfigUpdateResponseEnvelope struct {
	Errors   []MnmConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigUpdateResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmConfigUpdateResponseEnvelope]
type mnmConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmConfigUpdateResponseEnvelopeErrors]
type mnmConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmConfigUpdateResponseEnvelopeMessages]
type mnmConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigUpdateResponseEnvelopeSuccess bool

const (
	MnmConfigUpdateResponseEnvelopeSuccessTrue MnmConfigUpdateResponseEnvelopeSuccess = true
)

type MnmConfigListResponseEnvelope struct {
	Errors   []MnmConfigListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigListResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigListResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigListResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmConfigListResponseEnvelope]
type mnmConfigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmConfigListResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmConfigListResponseEnvelopeErrors]
type mnmConfigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmConfigListResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmConfigListResponseEnvelopeMessages]
type mnmConfigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigListResponseEnvelopeSuccess bool

const (
	MnmConfigListResponseEnvelopeSuccessTrue MnmConfigListResponseEnvelopeSuccess = true
)

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

type MnmConfigReplaceResponseEnvelope struct {
	Errors   []MnmConfigReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmConfigReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmConfigReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MnmConfigReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigReplaceResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmConfigReplaceResponseEnvelope]
type mnmConfigReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigReplaceResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    mnmConfigReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmConfigReplaceResponseEnvelopeErrors]
type mnmConfigReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmConfigReplaceResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    mnmConfigReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmConfigReplaceResponseEnvelopeMessages]
type mnmConfigReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmConfigReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmConfigReplaceResponseEnvelopeSuccess bool

const (
	MnmConfigReplaceResponseEnvelopeSuccessTrue MnmConfigReplaceResponseEnvelopeSuccess = true
)
