// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// AccessUserService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserService] method instead.
type AccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *AccessUserActiveSessionService
	LastSeenIdentity *AccessUserLastSeenIdentityService
	FailedLogins     *AccessUserFailedLoginService
}

// NewAccessUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessUserService(opts ...option.RequestOption) (r *AccessUserService) {
	r = &AccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewAccessUserFailedLoginService(opts...)
	return
}

// Creates a new user.
func (r *AccessUserService) New(ctx context.Context, params AccessUserNewParams, opts ...option.RequestOption) (res *AccessUserNewResponse, err error) {
	var env AccessUserNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific user's name for an account. Requires the user's current email
// as confirmation (email cannot be changed).
func (r *AccessUserService) Update(ctx context.Context, userID string, params AccessUserUpdateParams, opts ...option.RequestOption) (res *AccessUserUpdateResponse, err error) {
	var env AccessUserUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s", params.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of users for an account.
func (r *AccessUserService) List(ctx context.Context, params AccessUserListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessUserListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users", params.AccountID)
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

// Gets a list of users for an account.
func (r *AccessUserService) ListAutoPaging(ctx context.Context, params AccessUserListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessUserListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a specific user for an account. This will also revoke any active seats
// and tokens for the user.
func (r *AccessUserService) Delete(ctx context.Context, userID string, body AccessUserDeleteParams, opts ...option.RequestOption) (res *AccessUserDeleteResponse, err error) {
	var env AccessUserDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s", body.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a specific user for an account.
func (r *AccessUserService) Get(ctx context.Context, userID string, query AccessUserGetParams, opts ...option.RequestOption) (res *AccessUserGetResponse, err error) {
	var env AccessUserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s", query.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUser struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	ID string `json:"id"`
	// Determines the status of the SCIM User resource.
	Active bool `json:"active"`
	// The name of the SCIM User resource.
	DisplayName string            `json:"displayName"`
	Emails      []AccessUserEmail `json:"emails"`
	// The IdP-generated Id of the SCIM resource.
	ExternalID string `json:"externalId"`
	// The metadata of the SCIM resource.
	Meta AccessUserMeta `json:"meta"`
	// The list of URIs which indicate the attributes contained within a SCIM resource.
	Schemas []string       `json:"schemas"`
	JSON    accessUserJSON `json:"-"`
}

// accessUserJSON contains the JSON metadata for the struct [AccessUser]
type accessUserJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Schemas     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserJSON) RawJSON() string {
	return r.raw
}

type AccessUserEmail struct {
	// Indicates if the email address is the primary email belonging to the SCIM User
	// resource.
	Primary bool `json:"primary"`
	// Indicates the type of the email address.
	Type string `json:"type"`
	// The email address of the SCIM User resource.
	Value string              `json:"value" format:"email"`
	JSON  accessUserEmailJSON `json:"-"`
}

// accessUserEmailJSON contains the JSON metadata for the struct [AccessUserEmail]
type accessUserEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserEmailJSON) RawJSON() string {
	return r.raw
}

// The metadata of the SCIM resource.
type AccessUserMeta struct {
	// The timestamp of when the SCIM resource was created.
	Created time.Time `json:"created" format:"date-time"`
	// The timestamp of when the SCIM resource was last modified.
	LastModified time.Time          `json:"lastModified" format:"date-time"`
	JSON         accessUserMetaJSON `json:"-"`
}

// accessUserMetaJSON contains the JSON metadata for the struct [AccessUserMeta]
type accessUserMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewResponse struct {
	// UUID.
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID string `json:"seat_uid"`
	// The unique API identifier for the user.
	UID       string                    `json:"uid"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      accessUserNewResponseJSON `json:"-"`
}

// accessUserNewResponseJSON contains the JSON metadata for the struct
// [AccessUserNewResponse]
type accessUserNewResponseJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUID             apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserUpdateResponse struct {
	// UUID.
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID string `json:"seat_uid"`
	// The unique API identifier for the user.
	UID       string                       `json:"uid"`
	UpdatedAt time.Time                    `json:"updated_at" format:"date-time"`
	JSON      accessUserUpdateResponseJSON `json:"-"`
}

// accessUserUpdateResponseJSON contains the JSON metadata for the struct
// [AccessUserUpdateResponse]
type accessUserUpdateResponseJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUID             apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUserUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserListResponse struct {
	// UUID.
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID string `json:"seat_uid"`
	// The unique API identifier for the user.
	UID       string                     `json:"uid"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessUserListResponseJSON `json:"-"`
}

// accessUserListResponseJSON contains the JSON metadata for the struct
// [AccessUserListResponse]
type accessUserListResponseJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUID             apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserDeleteResponse = interface{}

type AccessUserGetResponse struct {
	// UUID.
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID string `json:"seat_uid"`
	// The unique API identifier for the user.
	UID       string                    `json:"uid"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      accessUserGetResponseJSON `json:"-"`
}

// accessUserGetResponseJSON contains the JSON metadata for the struct
// [AccessUserGetResponse]
type accessUserGetResponseJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUID             apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The name of the user.
	Name param.Field[string] `json:"name"`
}

func (r AccessUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserNewResponseEnvelope struct {
	Errors   []AccessUserNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessUserNewResponse                `json:"result"`
	JSON    accessUserNewResponseEnvelopeJSON    `json:"-"`
}

// accessUserNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessUserNewResponseEnvelope]
type accessUserNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccessUserNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessUserNewResponseEnvelopeErrors]
type accessUserNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accessUserNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AccessUserNewResponseEnvelopeErrorsSource]
type accessUserNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccessUserNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessUserNewResponseEnvelopeMessages]
type accessUserNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accessUserNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessUserNewResponseEnvelopeMessagesSource]
type accessUserNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserNewResponseEnvelopeSuccess bool

const (
	AccessUserNewResponseEnvelopeSuccessTrue AccessUserNewResponseEnvelopeSuccess = true
)

func (r AccessUserNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
	// The name of the user.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserUpdateResponseEnvelope struct {
	Errors   []AccessUserUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessUserUpdateResponse                `json:"result"`
	JSON    accessUserUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessUserUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessUserUpdateResponseEnvelope]
type accessUserUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccessUserUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessUserUpdateResponseEnvelopeErrors]
type accessUserUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accessUserUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessUserUpdateResponseEnvelopeErrorsSource]
type accessUserUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccessUserUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessUserUpdateResponseEnvelopeMessages]
type accessUserUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accessUserUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessUserUpdateResponseEnvelopeMessagesSource]
type accessUserUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserUpdateResponseEnvelopeSuccess bool

const (
	AccessUserUpdateResponseEnvelopeSuccessTrue AccessUserUpdateResponseEnvelopeSuccess = true
)

func (r AccessUserUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The email of the user.
	Email param.Field[string] `query:"email"`
	// The name of the user.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Search for users by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessUserListParams]'s query parameters as `url.Values`.
func (r AccessUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessUserDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessUserDeleteResponseEnvelope struct {
	Errors   []AccessUserDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessUserDeleteResponse                `json:"result,nullable"`
	JSON    accessUserDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessUserDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessUserDeleteResponseEnvelope]
type accessUserDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserDeleteResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccessUserDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessUserDeleteResponseEnvelopeErrors]
type accessUserDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accessUserDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessUserDeleteResponseEnvelopeErrorsSource]
type accessUserDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserDeleteResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccessUserDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessUserDeleteResponseEnvelopeMessages]
type accessUserDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accessUserDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessUserDeleteResponseEnvelopeMessagesSource]
type accessUserDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserDeleteResponseEnvelopeSuccess bool

const (
	AccessUserDeleteResponseEnvelopeSuccessTrue AccessUserDeleteResponseEnvelopeSuccess = true
)

func (r AccessUserDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessUserGetResponseEnvelope struct {
	Errors   []AccessUserGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessUserGetResponse                `json:"result"`
	JSON    accessUserGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessUserGetResponseEnvelope]
type accessUserGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccessUserGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessUserGetResponseEnvelopeErrors]
type accessUserGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accessUserGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AccessUserGetResponseEnvelopeErrorsSource]
type accessUserGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccessUserGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessUserGetResponseEnvelopeMessages]
type accessUserGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accessUserGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AccessUserGetResponseEnvelopeMessagesSource]
type accessUserGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserGetResponseEnvelopeSuccess bool

const (
	AccessUserGetResponseEnvelopeSuccessTrue AccessUserGetResponseEnvelopeSuccess = true
)

func (r AccessUserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
