// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Creates and returns a new Hyperdrive configuration.
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *ConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Hyperdrive configuration.
func (r *ConfigService) Update(ctx context.Context, hyperdriveID string, params ConfigUpdateParams, opts ...option.RequestOption) (res *ConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Hyperdrives
func (r *ConfigService) List(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) (res *shared.SinglePage[ConfigListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of Hyperdrives
func (r *ConfigService) ListAutoPaging(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[ConfigListResponse] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the specified Hyperdrive.
func (r *ConfigService) Delete(ctx context.Context, hyperdriveID string, body ConfigDeleteParams, opts ...option.RequestOption) (res *ConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", body.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches and returns the specified Hyperdrive configuration. Updates to the
// origin and caching settings are applied with an all-or-nothing approach.
func (r *ConfigService) Edit(ctx context.Context, hyperdriveID string, params ConfigEditParams, opts ...option.RequestOption) (res *ConfigEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", params.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Hyperdrive configuration.
func (r *ConfigService) Get(ctx context.Context, hyperdriveID string, query ConfigGetParams, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/hyperdrive/configs/%s", query.AccountID, hyperdriveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigNewResponse struct {
	// Identifier
	ID   string                `json:"id"`
	JSON configNewResponseJSON `json:"-"`
}

// configNewResponseJSON contains the JSON metadata for the struct
// [ConfigNewResponse]
type configNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigUpdateResponse struct {
	// Identifier
	ID   string                   `json:"id"`
	JSON configUpdateResponseJSON `json:"-"`
}

// configUpdateResponseJSON contains the JSON metadata for the struct
// [ConfigUpdateResponse]
type configUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigListResponse struct {
	// Identifier
	ID   string                 `json:"id"`
	JSON configListResponseJSON `json:"-"`
}

// configListResponseJSON contains the JSON metadata for the struct
// [ConfigListResponse]
type configListResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [hyperdrive.ConfigDeleteResponseUnknown] or
// [shared.UnionString].
type ConfigDeleteResponse interface {
	ImplementsHyperdriveConfigDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ConfigEditResponse struct {
	// Identifier
	ID   string                 `json:"id"`
	JSON configEditResponseJSON `json:"-"`
}

// configEditResponseJSON contains the JSON metadata for the struct
// [ConfigEditResponse]
type configEditResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponse struct {
	// Identifier
	ID   string                `json:"id"`
	JSON configGetResponseJSON `json:"-"`
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigNewParams struct {
	// Identifier
	AccountID param.Field[string]                `path:"account_id,required"`
	Origin    param.Field[ConfigNewParamsOrigin] `json:"origin,required"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r ConfigNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewResponseEnvelope struct {
	Errors   []ConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigNewResponse                   `json:"result,required,nullable"`
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
	// Identifier
	AccountID param.Field[string]                   `path:"account_id,required"`
	Origin    param.Field[ConfigUpdateParamsOrigin] `json:"origin,required"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r ConfigUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateResponseEnvelope struct {
	Errors   []ConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigUpdateResponse                   `json:"result,required,nullable"`
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

type ConfigListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []ConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigDeleteResponse                   `json:"result,required,nullable"`
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
	// Identifier
	AccountID param.Field[string]                 `path:"account_id,required"`
	Origin    param.Field[ConfigEditParamsOrigin] `json:"origin"`
}

func (r ConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigEditParamsOrigin struct {
	// The password required to access your origin database. This value is write-only
	// and never returned by the API.
	Password param.Field[string] `json:"password,required"`
}

func (r ConfigEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigEditResponseEnvelope struct {
	Errors   []ConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigEditResponse                   `json:"result,required,nullable"`
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []ConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ConfigGetResponse                   `json:"result,required,nullable"`
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
