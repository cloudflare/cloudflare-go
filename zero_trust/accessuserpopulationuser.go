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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AccessUserPopulationUserService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserPopulationUserService] method instead.
type AccessUserPopulationUserService struct {
	Options []option.RequestOption
}

// NewAccessUserPopulationUserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessUserPopulationUserService(opts ...option.RequestOption) (r *AccessUserPopulationUserService) {
	r = &AccessUserPopulationUserService{}
	r.Options = opts
	return
}

// Creates a SCIM user in a user population.
func (r *AccessUserPopulationUserService) New(ctx context.Context, userPopulationID string, params AccessUserPopulationUserNewParams, opts ...option.RequestOption) (res *AccessUserPopulationUserNewResponse, err error) {
	var env AccessUserPopulationUserNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users", params.AccountID, userPopulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a SCIM user in a user population.
func (r *AccessUserPopulationUserService) Update(ctx context.Context, userPopulationID string, userID string, params AccessUserPopulationUserUpdateParams, opts ...option.RequestOption) (res *AccessUserPopulationUserUpdateResponse, err error) {
	var env AccessUserPopulationUserUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users/%s", params.AccountID, userPopulationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists SCIM users in a user population.
func (r *AccessUserPopulationUserService) List(ctx context.Context, userPopulationID string, params AccessUserPopulationUserListParams, opts ...option.RequestOption) (res *AccessUserPopulationUserListResponse, err error) {
	var env AccessUserPopulationUserListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users", params.AccountID, userPopulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes a SCIM user from a user population.
func (r *AccessUserPopulationUserService) Delete(ctx context.Context, userPopulationID string, userID string, body AccessUserPopulationUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users/%s", body.AccountID, userPopulationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetches a SCIM user from a user population.
func (r *AccessUserPopulationUserService) Get(ctx context.Context, userPopulationID string, userID string, params AccessUserPopulationUserGetParams, opts ...option.RequestOption) (res *AccessUserPopulationUserGetResponse, err error) {
	var env AccessUserPopulationUserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users/%s", params.AccountID, userPopulationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces a SCIM user in a user population.
func (r *AccessUserPopulationUserService) Replace(ctx context.Context, userPopulationID string, userID string, params AccessUserPopulationUserReplaceParams, opts ...option.RequestOption) (res *AccessUserPopulationUserReplaceResponse, err error) {
	var env AccessUserPopulationUserReplaceResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s/users/%s", params.AccountID, userPopulationID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A SCIM User resource. Additional SCIM extension attributes are supported.
type AccessUserPopulationUserNewResponse struct {
	ID          string                                     `json:"id"`
	Active      bool                                       `json:"active"`
	DisplayName string                                     `json:"displayName"`
	Emails      []AccessUserPopulationUserNewResponseEmail `json:"emails"`
	ExternalID  string                                     `json:"externalId"`
	Meta        AccessUserPopulationUserNewResponseMeta    `json:"meta"`
	Name        AccessUserPopulationUserNewResponseName    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas     []string                                `json:"schemas"`
	UserName    string                                  `json:"userName"`
	ExtraFields map[string]interface{}                  `json:"-" api:"extrafields"`
	JSON        accessUserPopulationUserNewResponseJSON `json:"-"`
}

// accessUserPopulationUserNewResponseJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserNewResponse]
type accessUserPopulationUserNewResponseJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Schemas     apijson.Field
	UserName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseEmail struct {
	Primary bool                                         `json:"primary"`
	Type    string                                       `json:"type"`
	Value   string                                       `json:"value"`
	JSON    accessUserPopulationUserNewResponseEmailJSON `json:"-"`
}

// accessUserPopulationUserNewResponseEmailJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserNewResponseEmail]
type accessUserPopulationUserNewResponseEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEmailJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseMeta struct {
	Created      time.Time                                   `json:"created" format:"date-time"`
	LastModified time.Time                                   `json:"lastModified" format:"date-time"`
	Location     string                                      `json:"location" format:"uri"`
	ResourceType string                                      `json:"resourceType"`
	Version      string                                      `json:"version"`
	JSON         accessUserPopulationUserNewResponseMetaJSON `json:"-"`
}

// accessUserPopulationUserNewResponseMetaJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserNewResponseMeta]
type accessUserPopulationUserNewResponseMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	Location     apijson.Field
	ResourceType apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseName struct {
	FamilyName      string                                      `json:"familyName"`
	Formatted       string                                      `json:"formatted"`
	GivenName       string                                      `json:"givenName"`
	HonorificPrefix string                                      `json:"honorificPrefix"`
	HonorificSuffix string                                      `json:"honorificSuffix"`
	MiddleName      string                                      `json:"middleName"`
	JSON            accessUserPopulationUserNewResponseNameJSON `json:"-"`
}

// accessUserPopulationUserNewResponseNameJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserNewResponseName]
type accessUserPopulationUserNewResponseNameJSON struct {
	FamilyName      apijson.Field
	Formatted       apijson.Field
	GivenName       apijson.Field
	HonorificPrefix apijson.Field
	HonorificSuffix apijson.Field
	MiddleName      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseNameJSON) RawJSON() string {
	return r.raw
}

// A SCIM User resource. Additional SCIM extension attributes are supported.
type AccessUserPopulationUserUpdateResponse struct {
	ID          string                                        `json:"id"`
	Active      bool                                          `json:"active"`
	DisplayName string                                        `json:"displayName"`
	Emails      []AccessUserPopulationUserUpdateResponseEmail `json:"emails"`
	ExternalID  string                                        `json:"externalId"`
	Meta        AccessUserPopulationUserUpdateResponseMeta    `json:"meta"`
	Name        AccessUserPopulationUserUpdateResponseName    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas     []string                                   `json:"schemas"`
	UserName    string                                     `json:"userName"`
	ExtraFields map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON        accessUserPopulationUserUpdateResponseJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserUpdateResponse]
type accessUserPopulationUserUpdateResponseJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Schemas     apijson.Field
	UserName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseEmail struct {
	Primary bool                                            `json:"primary"`
	Type    string                                          `json:"type"`
	Value   string                                          `json:"value"`
	JSON    accessUserPopulationUserUpdateResponseEmailJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseEmailJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserUpdateResponseEmail]
type accessUserPopulationUserUpdateResponseEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEmailJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseMeta struct {
	Created      time.Time                                      `json:"created" format:"date-time"`
	LastModified time.Time                                      `json:"lastModified" format:"date-time"`
	Location     string                                         `json:"location" format:"uri"`
	ResourceType string                                         `json:"resourceType"`
	Version      string                                         `json:"version"`
	JSON         accessUserPopulationUserUpdateResponseMetaJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseMetaJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserUpdateResponseMeta]
type accessUserPopulationUserUpdateResponseMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	Location     apijson.Field
	ResourceType apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseName struct {
	FamilyName      string                                         `json:"familyName"`
	Formatted       string                                         `json:"formatted"`
	GivenName       string                                         `json:"givenName"`
	HonorificPrefix string                                         `json:"honorificPrefix"`
	HonorificSuffix string                                         `json:"honorificSuffix"`
	MiddleName      string                                         `json:"middleName"`
	JSON            accessUserPopulationUserUpdateResponseNameJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseNameJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserUpdateResponseName]
type accessUserPopulationUserUpdateResponseNameJSON struct {
	FamilyName      apijson.Field
	Formatted       apijson.Field
	GivenName       apijson.Field
	HonorificPrefix apijson.Field
	HonorificSuffix apijson.Field
	MiddleName      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseNameJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponse struct {
	ItemsPerPage int64                                          `json:"itemsPerPage" api:"required"`
	Resources    []AccessUserPopulationUserListResponseResource `json:"Resources" api:"required"`
	Schemas      []string                                       `json:"schemas" api:"required"`
	StartIndex   int64                                          `json:"startIndex" api:"required"`
	TotalResults int64                                          `json:"totalResults" api:"required"`
	JSON         accessUserPopulationUserListResponseJSON       `json:"-"`
}

// accessUserPopulationUserListResponseJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserListResponse]
type accessUserPopulationUserListResponseJSON struct {
	ItemsPerPage apijson.Field
	Resources    apijson.Field
	Schemas      apijson.Field
	StartIndex   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseJSON) RawJSON() string {
	return r.raw
}

// A SCIM User resource. Additional SCIM extension attributes are supported.
type AccessUserPopulationUserListResponseResource struct {
	ID          string                                               `json:"id"`
	Active      bool                                                 `json:"active"`
	DisplayName string                                               `json:"displayName"`
	Emails      []AccessUserPopulationUserListResponseResourcesEmail `json:"emails"`
	ExternalID  string                                               `json:"externalId"`
	Meta        AccessUserPopulationUserListResponseResourcesMeta    `json:"meta"`
	Name        AccessUserPopulationUserListResponseResourcesName    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas     []string                                         `json:"schemas"`
	UserName    string                                           `json:"userName"`
	ExtraFields map[string]interface{}                           `json:"-" api:"extrafields"`
	JSON        accessUserPopulationUserListResponseResourceJSON `json:"-"`
}

// accessUserPopulationUserListResponseResourceJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserListResponseResource]
type accessUserPopulationUserListResponseResourceJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Schemas     apijson.Field
	UserName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseResourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseResourcesEmail struct {
	Primary bool                                                   `json:"primary"`
	Type    string                                                 `json:"type"`
	Value   string                                                 `json:"value"`
	JSON    accessUserPopulationUserListResponseResourcesEmailJSON `json:"-"`
}

// accessUserPopulationUserListResponseResourcesEmailJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserListResponseResourcesEmail]
type accessUserPopulationUserListResponseResourcesEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseResourcesEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseResourcesEmailJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseResourcesMeta struct {
	Created      time.Time                                             `json:"created" format:"date-time"`
	LastModified time.Time                                             `json:"lastModified" format:"date-time"`
	Location     string                                                `json:"location" format:"uri"`
	ResourceType string                                                `json:"resourceType"`
	Version      string                                                `json:"version"`
	JSON         accessUserPopulationUserListResponseResourcesMetaJSON `json:"-"`
}

// accessUserPopulationUserListResponseResourcesMetaJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserListResponseResourcesMeta]
type accessUserPopulationUserListResponseResourcesMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	Location     apijson.Field
	ResourceType apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseResourcesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseResourcesMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseResourcesName struct {
	FamilyName      string                                                `json:"familyName"`
	Formatted       string                                                `json:"formatted"`
	GivenName       string                                                `json:"givenName"`
	HonorificPrefix string                                                `json:"honorificPrefix"`
	HonorificSuffix string                                                `json:"honorificSuffix"`
	MiddleName      string                                                `json:"middleName"`
	JSON            accessUserPopulationUserListResponseResourcesNameJSON `json:"-"`
}

// accessUserPopulationUserListResponseResourcesNameJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserListResponseResourcesName]
type accessUserPopulationUserListResponseResourcesNameJSON struct {
	FamilyName      apijson.Field
	Formatted       apijson.Field
	GivenName       apijson.Field
	HonorificPrefix apijson.Field
	HonorificSuffix apijson.Field
	MiddleName      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseResourcesName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseResourcesNameJSON) RawJSON() string {
	return r.raw
}

// A SCIM User resource. Additional SCIM extension attributes are supported.
type AccessUserPopulationUserGetResponse struct {
	ID          string                                     `json:"id"`
	Active      bool                                       `json:"active"`
	DisplayName string                                     `json:"displayName"`
	Emails      []AccessUserPopulationUserGetResponseEmail `json:"emails"`
	ExternalID  string                                     `json:"externalId"`
	Meta        AccessUserPopulationUserGetResponseMeta    `json:"meta"`
	Name        AccessUserPopulationUserGetResponseName    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas     []string                                `json:"schemas"`
	UserName    string                                  `json:"userName"`
	ExtraFields map[string]interface{}                  `json:"-" api:"extrafields"`
	JSON        accessUserPopulationUserGetResponseJSON `json:"-"`
}

// accessUserPopulationUserGetResponseJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserGetResponse]
type accessUserPopulationUserGetResponseJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Schemas     apijson.Field
	UserName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseEmail struct {
	Primary bool                                         `json:"primary"`
	Type    string                                       `json:"type"`
	Value   string                                       `json:"value"`
	JSON    accessUserPopulationUserGetResponseEmailJSON `json:"-"`
}

// accessUserPopulationUserGetResponseEmailJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserGetResponseEmail]
type accessUserPopulationUserGetResponseEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEmailJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseMeta struct {
	Created      time.Time                                   `json:"created" format:"date-time"`
	LastModified time.Time                                   `json:"lastModified" format:"date-time"`
	Location     string                                      `json:"location" format:"uri"`
	ResourceType string                                      `json:"resourceType"`
	Version      string                                      `json:"version"`
	JSON         accessUserPopulationUserGetResponseMetaJSON `json:"-"`
}

// accessUserPopulationUserGetResponseMetaJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserGetResponseMeta]
type accessUserPopulationUserGetResponseMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	Location     apijson.Field
	ResourceType apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseName struct {
	FamilyName      string                                      `json:"familyName"`
	Formatted       string                                      `json:"formatted"`
	GivenName       string                                      `json:"givenName"`
	HonorificPrefix string                                      `json:"honorificPrefix"`
	HonorificSuffix string                                      `json:"honorificSuffix"`
	MiddleName      string                                      `json:"middleName"`
	JSON            accessUserPopulationUserGetResponseNameJSON `json:"-"`
}

// accessUserPopulationUserGetResponseNameJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserGetResponseName]
type accessUserPopulationUserGetResponseNameJSON struct {
	FamilyName      apijson.Field
	Formatted       apijson.Field
	GivenName       apijson.Field
	HonorificPrefix apijson.Field
	HonorificSuffix apijson.Field
	MiddleName      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseNameJSON) RawJSON() string {
	return r.raw
}

// A SCIM User resource. Additional SCIM extension attributes are supported.
type AccessUserPopulationUserReplaceResponse struct {
	ID          string                                         `json:"id"`
	Active      bool                                           `json:"active"`
	DisplayName string                                         `json:"displayName"`
	Emails      []AccessUserPopulationUserReplaceResponseEmail `json:"emails"`
	ExternalID  string                                         `json:"externalId"`
	Meta        AccessUserPopulationUserReplaceResponseMeta    `json:"meta"`
	Name        AccessUserPopulationUserReplaceResponseName    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas     []string                                    `json:"schemas"`
	UserName    string                                      `json:"userName"`
	ExtraFields map[string]interface{}                      `json:"-" api:"extrafields"`
	JSON        accessUserPopulationUserReplaceResponseJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseJSON contains the JSON metadata for the
// struct [AccessUserPopulationUserReplaceResponse]
type accessUserPopulationUserReplaceResponseJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Schemas     apijson.Field
	UserName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseEmail struct {
	Primary bool                                             `json:"primary"`
	Type    string                                           `json:"type"`
	Value   string                                           `json:"value"`
	JSON    accessUserPopulationUserReplaceResponseEmailJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseEmailJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserReplaceResponseEmail]
type accessUserPopulationUserReplaceResponseEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEmailJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseMeta struct {
	Created      time.Time                                       `json:"created" format:"date-time"`
	LastModified time.Time                                       `json:"lastModified" format:"date-time"`
	Location     string                                          `json:"location" format:"uri"`
	ResourceType string                                          `json:"resourceType"`
	Version      string                                          `json:"version"`
	JSON         accessUserPopulationUserReplaceResponseMetaJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseMetaJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserReplaceResponseMeta]
type accessUserPopulationUserReplaceResponseMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	Location     apijson.Field
	ResourceType apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseName struct {
	FamilyName      string                                          `json:"familyName"`
	Formatted       string                                          `json:"formatted"`
	GivenName       string                                          `json:"givenName"`
	HonorificPrefix string                                          `json:"honorificPrefix"`
	HonorificSuffix string                                          `json:"honorificSuffix"`
	MiddleName      string                                          `json:"middleName"`
	JSON            accessUserPopulationUserReplaceResponseNameJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseNameJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserReplaceResponseName]
type accessUserPopulationUserReplaceResponseNameJSON struct {
	FamilyName      apijson.Field
	Formatted       apijson.Field
	GivenName       apijson.Field
	HonorificPrefix apijson.Field
	HonorificSuffix apijson.Field
	MiddleName      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseNameJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewParams struct {
	// Identifier.
	AccountID   param.Field[string]                                   `path:"account_id" api:"required"`
	Active      param.Field[bool]                                     `json:"active"`
	DisplayName param.Field[string]                                   `json:"displayName"`
	Emails      param.Field[[]AccessUserPopulationUserNewParamsEmail] `json:"emails"`
	ExternalID  param.Field[string]                                   `json:"externalId"`
	Meta        param.Field[AccessUserPopulationUserNewParamsMeta]    `json:"meta"`
	Name        param.Field[AccessUserPopulationUserNewParamsName]    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas  param.Field[[]string] `json:"schemas"`
	UserName param.Field[string]   `json:"userName"`
}

func (r AccessUserPopulationUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserNewParamsEmail struct {
	Primary param.Field[bool]   `json:"primary"`
	Type    param.Field[string] `json:"type"`
	Value   param.Field[string] `json:"value"`
}

func (r AccessUserPopulationUserNewParamsEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserNewParamsMeta struct {
	Created      param.Field[time.Time] `json:"created" format:"date-time"`
	LastModified param.Field[time.Time] `json:"lastModified" format:"date-time"`
	Location     param.Field[string]    `json:"location" format:"uri"`
	ResourceType param.Field[string]    `json:"resourceType"`
	Version      param.Field[string]    `json:"version"`
}

func (r AccessUserPopulationUserNewParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserNewParamsName struct {
	FamilyName      param.Field[string] `json:"familyName"`
	Formatted       param.Field[string] `json:"formatted"`
	GivenName       param.Field[string] `json:"givenName"`
	HonorificPrefix param.Field[string] `json:"honorificPrefix"`
	HonorificSuffix param.Field[string] `json:"honorificSuffix"`
	MiddleName      param.Field[string] `json:"middleName"`
}

func (r AccessUserPopulationUserNewParamsName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserNewResponseEnvelope struct {
	Errors   []AccessUserPopulationUserNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUserNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUserNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A SCIM User resource. Additional SCIM extension attributes are supported.
	Result AccessUserPopulationUserNewResponse             `json:"result"`
	JSON   accessUserPopulationUserNewResponseEnvelopeJSON `json:"-"`
}

// accessUserPopulationUserNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserNewResponseEnvelope]
type accessUserPopulationUserNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessUserPopulationUserNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUserNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUserNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserNewResponseEnvelopeErrors]
type accessUserPopulationUserNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessUserPopulationUserNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUserNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserNewResponseEnvelopeErrorsSource]
type accessUserPopulationUserNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessUserPopulationUserNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUserNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUserNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserNewResponseEnvelopeMessages]
type accessUserPopulationUserNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessUserPopulationUserNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUserNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserNewResponseEnvelopeMessagesSource]
type accessUserPopulationUserNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUserNewResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUserNewResponseEnvelopeSuccessTrue AccessUserPopulationUserNewResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUserNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationUserUpdateParams struct {
	// Identifier.
	AccountID  param.Field[string]                                          `path:"account_id" api:"required"`
	Operations param.Field[[]AccessUserPopulationUserUpdateParamsOperation] `json:"Operations" api:"required"`
	Schemas    param.Field[[]string]                                        `json:"schemas"`
}

func (r AccessUserPopulationUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserUpdateParamsOperation struct {
	Op    param.Field[AccessUserPopulationUserUpdateParamsOperationsOp]         `json:"op" api:"required"`
	Path  param.Field[string]                                                   `json:"path"`
	Value param.Field[AccessUserPopulationUserUpdateParamsOperationsValueUnion] `json:"value"`
}

func (r AccessUserPopulationUserUpdateParamsOperation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserUpdateParamsOperationsOp string

const (
	AccessUserPopulationUserUpdateParamsOperationsOpAdd     AccessUserPopulationUserUpdateParamsOperationsOp = "add"
	AccessUserPopulationUserUpdateParamsOperationsOpRemove  AccessUserPopulationUserUpdateParamsOperationsOp = "remove"
	AccessUserPopulationUserUpdateParamsOperationsOpReplace AccessUserPopulationUserUpdateParamsOperationsOp = "replace"
)

func (r AccessUserPopulationUserUpdateParamsOperationsOp) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserUpdateParamsOperationsOpAdd, AccessUserPopulationUserUpdateParamsOperationsOpRemove, AccessUserPopulationUserUpdateParamsOperationsOpReplace:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [zero_trust.AccessUserPopulationUserUpdateParamsOperationsValueArray],
// [zero_trust.AccessUserPopulationUserUpdateParamsOperationsValueMap].
type AccessUserPopulationUserUpdateParamsOperationsValueUnion interface {
	ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion()
}

type AccessUserPopulationUserUpdateParamsOperationsValueArray []interface{}

func (r AccessUserPopulationUserUpdateParamsOperationsValueArray) ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion() {
}

type AccessUserPopulationUserUpdateParamsOperationsValueMap map[string]interface{}

func (r AccessUserPopulationUserUpdateParamsOperationsValueMap) ImplementsAccessUserPopulationUserUpdateParamsOperationsValueUnion() {
}

type AccessUserPopulationUserUpdateResponseEnvelope struct {
	Errors   []AccessUserPopulationUserUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUserUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUserUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// A SCIM User resource. Additional SCIM extension attributes are supported.
	Result AccessUserPopulationUserUpdateResponse             `json:"result"`
	JSON   accessUserPopulationUserUpdateResponseEnvelopeJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserUpdateResponseEnvelope]
type accessUserPopulationUserUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessUserPopulationUserUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUserUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUserUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserUpdateResponseEnvelopeErrors]
type accessUserPopulationUserUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessUserPopulationUserUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserUpdateResponseEnvelopeErrorsSource]
type accessUserPopulationUserUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessUserPopulationUserUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUserUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUserUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserUpdateResponseEnvelopeMessages]
type accessUserPopulationUserUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessUserPopulationUserUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUserUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessUserPopulationUserUpdateResponseEnvelopeMessagesSource]
type accessUserPopulationUserUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUserUpdateResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUserUpdateResponseEnvelopeSuccessTrue AccessUserPopulationUserUpdateResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUserUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationUserListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Comma-separated SCIM attributes to return.
	Attributes param.Field[string] `query:"attributes"`
	// The maximum number of results to return.
	Count param.Field[int64] `query:"count"`
	// Comma-separated SCIM attributes to exclude from the response.
	ExcludedAttributes param.Field[string] `query:"excludedAttributes"`
	// A SCIM filter expression.
	Filter param.Field[string] `query:"filter"`
	// The 1-based index of the first result to return.
	StartIndex param.Field[int64] `query:"startIndex"`
}

// URLQuery serializes [AccessUserPopulationUserListParams]'s query parameters as
// `url.Values`.
func (r AccessUserPopulationUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessUserPopulationUserListResponseEnvelope struct {
	Errors   []AccessUserPopulationUserListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUserListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUserListResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPopulationUserListResponse                `json:"result"`
	JSON    accessUserPopulationUserListResponseEnvelopeJSON    `json:"-"`
}

// accessUserPopulationUserListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserListResponseEnvelope]
type accessUserPopulationUserListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessUserPopulationUserListResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUserListResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUserListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserListResponseEnvelopeErrors]
type accessUserPopulationUserListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessUserPopulationUserListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUserListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserListResponseEnvelopeErrorsSource]
type accessUserPopulationUserListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessUserPopulationUserListResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUserListResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUserListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserListResponseEnvelopeMessages]
type accessUserPopulationUserListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserListResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessUserPopulationUserListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUserListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserListResponseEnvelopeMessagesSource]
type accessUserPopulationUserListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUserListResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUserListResponseEnvelopeSuccessTrue AccessUserPopulationUserListResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUserListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationUserDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessUserPopulationUserGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Comma-separated SCIM attributes to return.
	Attributes param.Field[string] `query:"attributes"`
	// Comma-separated SCIM attributes to exclude from the response.
	ExcludedAttributes param.Field[string] `query:"excludedAttributes"`
}

// URLQuery serializes [AccessUserPopulationUserGetParams]'s query parameters as
// `url.Values`.
func (r AccessUserPopulationUserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessUserPopulationUserGetResponseEnvelope struct {
	Errors   []AccessUserPopulationUserGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUserGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUserGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A SCIM User resource. Additional SCIM extension attributes are supported.
	Result AccessUserPopulationUserGetResponse             `json:"result"`
	JSON   accessUserPopulationUserGetResponseEnvelopeJSON `json:"-"`
}

// accessUserPopulationUserGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserPopulationUserGetResponseEnvelope]
type accessUserPopulationUserGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessUserPopulationUserGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUserGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUserGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserGetResponseEnvelopeErrors]
type accessUserPopulationUserGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessUserPopulationUserGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUserGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserGetResponseEnvelopeErrorsSource]
type accessUserPopulationUserGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessUserPopulationUserGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUserGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUserGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserGetResponseEnvelopeMessages]
type accessUserPopulationUserGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessUserPopulationUserGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUserGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserGetResponseEnvelopeMessagesSource]
type accessUserPopulationUserGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUserGetResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUserGetResponseEnvelopeSuccessTrue AccessUserPopulationUserGetResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationUserReplaceParams struct {
	// Identifier.
	AccountID   param.Field[string]                                       `path:"account_id" api:"required"`
	Active      param.Field[bool]                                         `json:"active"`
	DisplayName param.Field[string]                                       `json:"displayName"`
	Emails      param.Field[[]AccessUserPopulationUserReplaceParamsEmail] `json:"emails"`
	ExternalID  param.Field[string]                                       `json:"externalId"`
	Meta        param.Field[AccessUserPopulationUserReplaceParamsMeta]    `json:"meta"`
	Name        param.Field[AccessUserPopulationUserReplaceParamsName]    `json:"name"`
	// The SCIM schemas that apply to this resource.
	Schemas  param.Field[[]string] `json:"schemas"`
	UserName param.Field[string]   `json:"userName"`
}

func (r AccessUserPopulationUserReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserReplaceParamsEmail struct {
	Primary param.Field[bool]   `json:"primary"`
	Type    param.Field[string] `json:"type"`
	Value   param.Field[string] `json:"value"`
}

func (r AccessUserPopulationUserReplaceParamsEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserReplaceParamsMeta struct {
	Created      param.Field[time.Time] `json:"created" format:"date-time"`
	LastModified param.Field[time.Time] `json:"lastModified" format:"date-time"`
	Location     param.Field[string]    `json:"location" format:"uri"`
	ResourceType param.Field[string]    `json:"resourceType"`
	Version      param.Field[string]    `json:"version"`
}

func (r AccessUserPopulationUserReplaceParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserReplaceParamsName struct {
	FamilyName      param.Field[string] `json:"familyName"`
	Formatted       param.Field[string] `json:"formatted"`
	GivenName       param.Field[string] `json:"givenName"`
	HonorificPrefix param.Field[string] `json:"honorificPrefix"`
	HonorificSuffix param.Field[string] `json:"honorificSuffix"`
	MiddleName      param.Field[string] `json:"middleName"`
}

func (r AccessUserPopulationUserReplaceParamsName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUserReplaceResponseEnvelope struct {
	Errors   []AccessUserPopulationUserReplaceResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUserReplaceResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUserReplaceResponseEnvelopeSuccess `json:"success" api:"required"`
	// A SCIM User resource. Additional SCIM extension attributes are supported.
	Result AccessUserPopulationUserReplaceResponse             `json:"result"`
	JSON   accessUserPopulationUserReplaceResponseEnvelopeJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessUserPopulationUserReplaceResponseEnvelope]
type accessUserPopulationUserReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessUserPopulationUserReplaceResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUserReplaceResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUserReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserPopulationUserReplaceResponseEnvelopeErrors]
type accessUserPopulationUserReplaceResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessUserPopulationUserReplaceResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccessUserPopulationUserReplaceResponseEnvelopeErrorsSource]
type accessUserPopulationUserReplaceResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           AccessUserPopulationUserReplaceResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUserReplaceResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUserReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUserReplaceResponseEnvelopeMessages]
type accessUserPopulationUserReplaceResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUserReplaceResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    accessUserPopulationUserReplaceResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUserReplaceResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessUserPopulationUserReplaceResponseEnvelopeMessagesSource]
type accessUserPopulationUserReplaceResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUserReplaceResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUserReplaceResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUserReplaceResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUserReplaceResponseEnvelopeSuccessTrue AccessUserPopulationUserReplaceResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUserReplaceResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUserReplaceResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
