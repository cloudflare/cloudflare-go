// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// HyperdriveConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHyperdriveConfigService] method
// instead.
type HyperdriveConfigService struct {
	Options []option.RequestOption
}

// NewHyperdriveConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHyperdriveConfigService(opts ...option.RequestOption) (r *HyperdriveConfigService) {
	r = &HyperdriveConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration.
func (r *HyperdriveConfigService) New(ctx context.Context, accountID string, body HyperdriveConfigNewParams, opts ...option.RequestOption) (res *HyperdriveConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Update(ctx context.Context, accountID string, hyperdriveID string, body HyperdriveConfigUpdateParams, opts ...option.RequestOption) (res *HyperdriveConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Hyperdrives
func (r *HyperdriveConfigService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]HyperdriveConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the specified Hyperdrive.
func (r *HyperdriveConfigService) Delete(ctx context.Context, accountID string, hyperdriveID string, opts ...option.RequestOption) (res *HyperdriveConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches and returns the specified Hyperdrive configuration. Updates to the
// origin and caching settings are applied with an all-or-nothing approach.
func (r *HyperdriveConfigService) Edit(ctx context.Context, accountID string, hyperdriveID string, body HyperdriveConfigEditParams, opts ...option.RequestOption) (res *HyperdriveConfigEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Hyperdrive configuration.
func (r *HyperdriveConfigService) Get(ctx context.Context, accountID string, hyperdriveID string, opts ...option.RequestOption) (res *HyperdriveConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HyperdriveConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", accountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HyperdriveConfigNewResponse struct {
	// Identifier
	ID   string                          `json:"id"`
	JSON hyperdriveConfigNewResponseJSON `json:"-"`
}

// hyperdriveConfigNewResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigNewResponse]
type hyperdriveConfigNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponse struct {
	// Identifier
	ID   string                             `json:"id"`
	JSON hyperdriveConfigUpdateResponseJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigUpdateResponse]
type hyperdriveConfigUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponse struct {
	// Identifier
	ID   string                           `json:"id"`
	JSON hyperdriveConfigListResponseJSON `json:"-"`
}

// hyperdriveConfigListResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigListResponse]
type hyperdriveConfigListResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [HyperdriveConfigDeleteResponseUnknown] or
// [shared.UnionString].
type HyperdriveConfigDeleteResponse interface {
	ImplementsHyperdriveConfigDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HyperdriveConfigDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type HyperdriveConfigEditResponse struct {
	// Identifier
	ID   string                           `json:"id"`
	JSON hyperdriveConfigEditResponseJSON `json:"-"`
}

// hyperdriveConfigEditResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigEditResponse]
type hyperdriveConfigEditResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponse struct {
	// Identifier
	ID   string                          `json:"id"`
	JSON hyperdriveConfigGetResponseJSON `json:"-"`
}

// hyperdriveConfigGetResponseJSON contains the JSON metadata for the struct
// [HyperdriveConfigGetResponse]
type hyperdriveConfigGetResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewParams struct {
	Origin param.Field[HyperdriveConfigNewParamsOrigin] `json:"origin,required"`
}

func (r HyperdriveConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigNewParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r HyperdriveConfigNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigNewResponseEnvelope struct {
	Errors   []HyperdriveConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   HyperdriveConfigNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigNewResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigNewResponseEnvelope]
type hyperdriveConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    hyperdriveConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HyperdriveConfigNewResponseEnvelopeErrors]
type hyperdriveConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    hyperdriveConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigNewResponseEnvelopeMessages]
type hyperdriveConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigNewResponseEnvelopeSuccess bool

const (
	HyperdriveConfigNewResponseEnvelopeSuccessTrue HyperdriveConfigNewResponseEnvelopeSuccess = true
)

type HyperdriveConfigUpdateParams struct {
	Origin param.Field[HyperdriveConfigUpdateParamsOrigin] `json:"origin,required"`
}

func (r HyperdriveConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigUpdateParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r HyperdriveConfigUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigUpdateResponseEnvelope struct {
	Errors   []HyperdriveConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   HyperdriveConfigUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigUpdateResponseEnvelope]
type hyperdriveConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigUpdateResponseEnvelopeErrors]
type hyperdriveConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hyperdriveConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HyperdriveConfigUpdateResponseEnvelopeMessages]
type hyperdriveConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigUpdateResponseEnvelopeSuccess bool

const (
	HyperdriveConfigUpdateResponseEnvelopeSuccessTrue HyperdriveConfigUpdateResponseEnvelopeSuccess = true
)

type HyperdriveConfigListResponseEnvelope struct {
	Errors   []HyperdriveConfigListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigListResponseEnvelopeMessages `json:"messages,required"`
	Result   []HyperdriveConfigListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success HyperdriveConfigListResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigListResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigListResponseEnvelope]
type hyperdriveConfigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    hyperdriveConfigListResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigListResponseEnvelopeErrors]
type hyperdriveConfigListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigListResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigListResponseEnvelopeMessages]
type hyperdriveConfigListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigListResponseEnvelopeSuccess bool

const (
	HyperdriveConfigListResponseEnvelopeSuccessTrue HyperdriveConfigListResponseEnvelopeSuccess = true
)

type HyperdriveConfigDeleteResponseEnvelope struct {
	Errors   []HyperdriveConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HyperdriveConfigDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigDeleteResponseEnvelope]
type hyperdriveConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigDeleteResponseEnvelopeErrors]
type hyperdriveConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    hyperdriveConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HyperdriveConfigDeleteResponseEnvelopeMessages]
type hyperdriveConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigDeleteResponseEnvelopeSuccess bool

const (
	HyperdriveConfigDeleteResponseEnvelopeSuccessTrue HyperdriveConfigDeleteResponseEnvelopeSuccess = true
)

type HyperdriveConfigEditParams struct {
	Origin param.Field[HyperdriveConfigEditParamsOrigin] `json:"origin"`
}

func (r HyperdriveConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigEditParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r HyperdriveConfigEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HyperdriveConfigEditResponseEnvelope struct {
	Errors   []HyperdriveConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigEditResponseEnvelopeMessages `json:"messages,required"`
	Result   HyperdriveConfigEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigEditResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigEditResponseEnvelope]
type hyperdriveConfigEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    hyperdriveConfigEditResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HyperdriveConfigEditResponseEnvelopeErrors]
type hyperdriveConfigEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    hyperdriveConfigEditResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigEditResponseEnvelopeMessages]
type hyperdriveConfigEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigEditResponseEnvelopeSuccess bool

const (
	HyperdriveConfigEditResponseEnvelopeSuccessTrue HyperdriveConfigEditResponseEnvelopeSuccess = true
)

type HyperdriveConfigGetResponseEnvelope struct {
	Errors   []HyperdriveConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HyperdriveConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HyperdriveConfigGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success HyperdriveConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hyperdriveConfigGetResponseEnvelopeJSON    `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseEnvelope]
type hyperdriveConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    hyperdriveConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HyperdriveConfigGetResponseEnvelopeErrors]
type hyperdriveConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HyperdriveConfigGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    hyperdriveConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hyperdriveConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HyperdriveConfigGetResponseEnvelopeMessages]
type hyperdriveConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HyperdriveConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HyperdriveConfigGetResponseEnvelopeSuccess bool

const (
	HyperdriveConfigGetResponseEnvelopeSuccessTrue HyperdriveConfigGetResponseEnvelopeSuccess = true
)
