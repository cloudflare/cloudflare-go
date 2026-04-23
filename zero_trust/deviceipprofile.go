// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// DeviceIPProfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceIPProfileService] method instead.
type DeviceIPProfileService struct {
	Options []option.RequestOption
}

// NewDeviceIPProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceIPProfileService(opts ...option.RequestOption) (r *DeviceIPProfileService) {
	r = &DeviceIPProfileService{}
	r.Options = opts
	return
}

// Creates a WARP Device IP profile. Currently, only IPv4 Device subnets can be
// associated.
func (r *DeviceIPProfileService) New(ctx context.Context, params DeviceIPProfileNewParams, opts ...option.RequestOption) (res *IPProfile, err error) {
	var env DeviceIPProfileNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/ip-profiles", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a WARP Device IP profile. Currently, only IPv4 Device subnets can be
// associated.
func (r *DeviceIPProfileService) Update(ctx context.Context, profileID string, params DeviceIPProfileUpdateParams, opts ...option.RequestOption) (res *IPProfile, err error) {
	var env DeviceIPProfileUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/ip-profiles/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists WARP Device IP profiles.
func (r *DeviceIPProfileService) List(ctx context.Context, params DeviceIPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[IPProfile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/ip-profiles", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists WARP Device IP profiles.
func (r *DeviceIPProfileService) ListAutoPaging(ctx context.Context, params DeviceIPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IPProfile] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a WARP Device IP profile.
func (r *DeviceIPProfileService) Delete(ctx context.Context, profileID string, body DeviceIPProfileDeleteParams, opts ...option.RequestOption) (res *DeviceIPProfileDeleteResponse, err error) {
	var env DeviceIPProfileDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/ip-profiles/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a single WARP Device IP profile.
func (r *DeviceIPProfileService) Get(ctx context.Context, profileID string, query DeviceIPProfileGetParams, opts ...option.RequestOption) (res *IPProfile, err error) {
	var env DeviceIPProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/devices/ip-profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type IPProfile struct {
	// The ID of the Device IP profile.
	ID string `json:"id" api:"required"`
	// The RFC3339Nano timestamp when the Device IP profile was created.
	CreatedAt string `json:"created_at" api:"required"`
	// An optional description of the Device IP profile.
	Description string `json:"description" api:"required,nullable"`
	// Whether the Device IP profile is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// The wirefilter expression to match registrations. Available values:
	// "identity.name", "identity.email", "identity.groups.id", "identity.groups.name",
	// "identity.groups.email", "identity.saml_attributes".
	Match string `json:"match" api:"required"`
	// A user-friendly name for the Device IP profile.
	Name string `json:"name" api:"required"`
	// The precedence of the Device IP profile. Lower values indicate higher
	// precedence. Device IP profile will be evaluated in ascending order of this
	// field.
	Precedence int64 `json:"precedence" api:"required"`
	// The ID of the Subnet.
	SubnetID string `json:"subnet_id" api:"required"`
	// The RFC3339Nano timestamp when the Device IP profile was last updated.
	UpdatedAt string        `json:"updated_at" api:"required"`
	JSON      ipProfileJSON `json:"-"`
}

// ipProfileJSON contains the JSON metadata for the struct [IPProfile]
type ipProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Precedence  apijson.Field
	SubnetID    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipProfileJSON) RawJSON() string {
	return r.raw
}

type DeviceIPProfileDeleteResponse struct {
	// ID of the deleted Device IP profile.
	ID   string                            `json:"id"`
	JSON deviceIPProfileDeleteResponseJSON `json:"-"`
}

// deviceIPProfileDeleteResponseJSON contains the JSON metadata for the struct
// [DeviceIPProfileDeleteResponse]
type deviceIPProfileDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceIPProfileNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The wirefilter expression to match registrations. Available values:
	// "identity.name", "identity.email", "identity.groups.id", "identity.groups.name",
	// "identity.groups.email", "identity.saml_attributes".
	Match param.Field[string] `json:"match" api:"required"`
	// A user-friendly name for the Device IP profile.
	Name param.Field[string] `json:"name" api:"required"`
	// The precedence of the Device IP profile. Lower values indicate higher
	// precedence. Device IP profile will be evaluated in ascending order of this
	// field.
	Precedence param.Field[int64] `json:"precedence" api:"required"`
	// The ID of the Subnet.
	SubnetID param.Field[string] `json:"subnet_id" api:"required"`
	// An optional description of the Device IP profile.
	Description param.Field[string] `json:"description"`
	// Whether the Device IP profile will be applied to matching devices.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DeviceIPProfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceIPProfileNewResponseEnvelope struct {
	Errors   []DeviceIPProfileNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceIPProfileNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   IPProfile                                    `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success" api:"required"`
	JSON    deviceIPProfileNewResponseEnvelopeJSON `json:"-"`
}

// deviceIPProfileNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceIPProfileNewResponseEnvelope]
type deviceIPProfileNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	JSON    deviceIPProfileNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceIPProfileNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceIPProfileNewResponseEnvelopeErrors]
type deviceIPProfileNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	JSON    deviceIPProfileNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceIPProfileNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceIPProfileNewResponseEnvelopeMessages]
type deviceIPProfileNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceIPProfileUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// An optional description of the Device IP profile.
	Description param.Field[string] `json:"description"`
	// Whether the Device IP profile is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The wirefilter expression to match registrations. Available values:
	// "identity.name", "identity.email", "identity.groups.id", "identity.groups.name",
	// "identity.groups.email", "identity.saml_attributes".
	Match param.Field[string] `json:"match"`
	// A user-friendly name for the Device IP profile.
	Name param.Field[string] `json:"name"`
	// The precedence of the Device IP profile. Lower values indicate higher
	// precedence. Device IP profile will be evaluated in ascending order of this
	// field.
	Precedence param.Field[int64] `json:"precedence"`
	// The ID of the Subnet.
	SubnetID param.Field[string] `json:"subnet_id"`
}

func (r DeviceIPProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceIPProfileUpdateResponseEnvelope struct {
	Errors   []DeviceIPProfileUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceIPProfileUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   IPProfile                                       `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success" api:"required"`
	JSON    deviceIPProfileUpdateResponseEnvelopeJSON `json:"-"`
}

// deviceIPProfileUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceIPProfileUpdateResponseEnvelope]
type deviceIPProfileUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    deviceIPProfileUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceIPProfileUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceIPProfileUpdateResponseEnvelopeErrors]
type deviceIPProfileUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    deviceIPProfileUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceIPProfileUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceIPProfileUpdateResponseEnvelopeMessages]
type deviceIPProfileUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceIPProfileListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The number of IP profiles to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [DeviceIPProfileListParams]'s query parameters as
// `url.Values`.
func (r DeviceIPProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceIPProfileDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceIPProfileDeleteResponseEnvelope struct {
	Errors   []DeviceIPProfileDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceIPProfileDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DeviceIPProfileDeleteResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success" api:"required"`
	JSON    deviceIPProfileDeleteResponseEnvelopeJSON `json:"-"`
}

// deviceIPProfileDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceIPProfileDeleteResponseEnvelope]
type deviceIPProfileDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	JSON    deviceIPProfileDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceIPProfileDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceIPProfileDeleteResponseEnvelopeErrors]
type deviceIPProfileDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	JSON    deviceIPProfileDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceIPProfileDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceIPProfileDeleteResponseEnvelopeMessages]
type deviceIPProfileDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceIPProfileGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DeviceIPProfileGetResponseEnvelope struct {
	Errors   []DeviceIPProfileGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DeviceIPProfileGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   IPProfile                                    `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success" api:"required"`
	JSON    deviceIPProfileGetResponseEnvelopeJSON `json:"-"`
}

// deviceIPProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceIPProfileGetResponseEnvelope]
type deviceIPProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	JSON    deviceIPProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceIPProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceIPProfileGetResponseEnvelopeErrors]
type deviceIPProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceIPProfileGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	JSON    deviceIPProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceIPProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceIPProfileGetResponseEnvelopeMessages]
type deviceIPProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceIPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceIPProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
