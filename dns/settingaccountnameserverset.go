// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// SettingAccountNameserverSetService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAccountNameserverSetService] method instead.
type SettingAccountNameserverSetService struct {
	Options []option.RequestOption
}

// NewSettingAccountNameserverSetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingAccountNameserverSetService(opts ...option.RequestOption) (r *SettingAccountNameserverSetService) {
	r = &SettingAccountNameserverSetService{}
	r.Options = opts
	return
}

// Creates an immutable Custom Nameserver Set. To change a set, create a new one,
// move any zone assignments, and delete the old set.
func (r *SettingAccountNameserverSetService) New(ctx context.Context, params SettingAccountNameserverSetNewParams, opts ...option.RequestOption) (res *SettingAccountNameserverSetNewResponse, err error) {
	var env SettingAccountNameserverSetNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/nameserver_sets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists an account's Custom Nameserver Sets.
func (r *SettingAccountNameserverSetService) List(ctx context.Context, params SettingAccountNameserverSetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingAccountNameserverSetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/nameserver_sets", params.AccountID)
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

// Lists an account's Custom Nameserver Sets.
func (r *SettingAccountNameserverSetService) ListAutoPaging(ctx context.Context, params SettingAccountNameserverSetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingAccountNameserverSetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an unassigned Custom Nameserver Set.
func (r *SettingAccountNameserverSetService) Delete(ctx context.Context, nameserverSetID string, body SettingAccountNameserverSetDeleteParams, opts ...option.RequestOption) (res *SettingAccountNameserverSetDeleteResponse, err error) {
	var env SettingAccountNameserverSetDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if nameserverSetID == "" {
		err = errors.New("missing required nameserver_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/nameserver_sets/%s", body.AccountID, nameserverSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a Custom Nameserver Set owned by an account.
func (r *SettingAccountNameserverSetService) Get(ctx context.Context, nameserverSetID string, query SettingAccountNameserverSetGetParams, opts ...option.RequestOption) (res *SettingAccountNameserverSetGetResponse, err error) {
	var env SettingAccountNameserverSetGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if nameserverSetID == "" {
		err = errors.New("missing required nameserver_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/nameserver_sets/%s", query.AccountID, nameserverSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SettingAccountNameserverSetNewResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetNewResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet int64 `json:"ip_set" api:"required"`
	// This field can have the runtime type of
	// [[]SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserver],
	// [[]SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserver].
	Nameservers interface{}                                `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetNewResponseJSON `json:"-"`
	union       SettingAccountNameserverSetNewResponseUnion
}

// settingAccountNameserverSetNewResponseJSON contains the JSON metadata for the
// struct [SettingAccountNameserverSetNewResponse]
type settingAccountNameserverSetNewResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r settingAccountNameserverSetNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingAccountNameserverSetNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SettingAccountNameserverSetNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingAccountNameserverSetNewResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse],
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse].
func (r SettingAccountNameserverSetNewResponse) AsUnion() SettingAccountNameserverSetNewResponseUnion {
	return r.union
}

// Union satisfied by
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse]
// or
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse].
type SettingAccountNameserverSetNewResponseUnion interface {
	implementsSettingAccountNameserverSetNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingAccountNameserverSetNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse{}),
		},
	)
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                      `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseJSON         `json:"-"`
}

// settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse]
type settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponse) implementsSettingAccountNameserverSetNewResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvanced bool

const (
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvanced = false
)

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse:
		return true
	}
	return false
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                       `json:"name" api:"required"`
	JSON settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserver]
type settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseDNSSettingsNameserverSetStandardResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                      `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseJSON         `json:"-"`
}

// settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse]
type settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponse) implementsSettingAccountNameserverSetNewResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvanced bool

const (
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvanced = true
)

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV4Groups []SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group `json:"ipv4_groups" api:"required"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV6Groups []SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group `json:"ipv6_groups" api:"required"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                       `json:"name" api:"required"`
	JSON settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserver]
type settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV4Groups  apijson.Field
	IPV6        apijson.Field
	IPV6Groups  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group string

const (
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "a"
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "b"
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "c"
)

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA, SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB, SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC:
		return true
	}
	return false
}

type SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group string

const (
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "a"
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "b"
	SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "c"
)

func (r SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA, SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB, SettingAccountNameserverSetNewResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC:
		return true
	}
	return false
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetNewResponseAdvanced bool

const (
	SettingAccountNameserverSetNewResponseAdvancedFalse SettingAccountNameserverSetNewResponseAdvanced = false
	SettingAccountNameserverSetNewResponseAdvancedTrue  SettingAccountNameserverSetNewResponseAdvanced = true
)

func (r SettingAccountNameserverSetNewResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseAdvancedFalse, SettingAccountNameserverSetNewResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetListResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetListResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet int64 `json:"ip_set" api:"required"`
	// This field can have the runtime type of
	// [[]SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserver],
	// [[]SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserver].
	Nameservers interface{}                                 `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetListResponseJSON `json:"-"`
	union       SettingAccountNameserverSetListResponseUnion
}

// settingAccountNameserverSetListResponseJSON contains the JSON metadata for the
// struct [SettingAccountNameserverSetListResponse]
type settingAccountNameserverSetListResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r settingAccountNameserverSetListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingAccountNameserverSetListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SettingAccountNameserverSetListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingAccountNameserverSetListResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse],
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse].
func (r SettingAccountNameserverSetListResponse) AsUnion() SettingAccountNameserverSetListResponseUnion {
	return r.union
}

// Union satisfied by
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse]
// or
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse].
type SettingAccountNameserverSetListResponseUnion interface {
	implementsSettingAccountNameserverSetListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingAccountNameserverSetListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse{}),
		},
	)
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                       `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseJSON         `json:"-"`
}

// settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse]
type settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponse) implementsSettingAccountNameserverSetListResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvanced bool

const (
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvanced = false
)

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse:
		return true
	}
	return false
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                        `json:"name" api:"required"`
	JSON settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserver]
type settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetListResponseDNSSettingsNameserverSetStandardResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                       `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseJSON         `json:"-"`
}

// settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse]
type settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponse) implementsSettingAccountNameserverSetListResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvanced bool

const (
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvanced = true
)

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV4Groups []SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group `json:"ipv4_groups" api:"required"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV6Groups []SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group `json:"ipv6_groups" api:"required"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                        `json:"name" api:"required"`
	JSON settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserver]
type settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV4Groups  apijson.Field
	IPV6        apijson.Field
	IPV6Groups  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group string

const (
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "a"
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "b"
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "c"
)

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA, SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB, SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC:
		return true
	}
	return false
}

type SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group string

const (
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "a"
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "b"
	SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "c"
)

func (r SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA, SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB, SettingAccountNameserverSetListResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC:
		return true
	}
	return false
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetListResponseAdvanced bool

const (
	SettingAccountNameserverSetListResponseAdvancedFalse SettingAccountNameserverSetListResponseAdvanced = false
	SettingAccountNameserverSetListResponseAdvancedTrue  SettingAccountNameserverSetListResponseAdvanced = true
)

func (r SettingAccountNameserverSetListResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetListResponseAdvancedFalse, SettingAccountNameserverSetListResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetDeleteResponse struct {
	// Identifier for a nameserver set.
	ID   string                                        `json:"id" api:"required"`
	JSON settingAccountNameserverSetDeleteResponseJSON `json:"-"`
}

// settingAccountNameserverSetDeleteResponseJSON contains the JSON metadata for the
// struct [SettingAccountNameserverSetDeleteResponse]
type settingAccountNameserverSetDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetGetResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet int64 `json:"ip_set" api:"required"`
	// This field can have the runtime type of
	// [[]SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserver],
	// [[]SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserver].
	Nameservers interface{}                                `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetGetResponseJSON `json:"-"`
	union       SettingAccountNameserverSetGetResponseUnion
}

// settingAccountNameserverSetGetResponseJSON contains the JSON metadata for the
// struct [SettingAccountNameserverSetGetResponse]
type settingAccountNameserverSetGetResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r settingAccountNameserverSetGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingAccountNameserverSetGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SettingAccountNameserverSetGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingAccountNameserverSetGetResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse],
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse].
func (r SettingAccountNameserverSetGetResponse) AsUnion() SettingAccountNameserverSetGetResponseUnion {
	return r.union
}

// Union satisfied by
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse]
// or
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse].
type SettingAccountNameserverSetGetResponseUnion interface {
	implementsSettingAccountNameserverSetGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingAccountNameserverSetGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse{}),
		},
	)
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                      `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseJSON         `json:"-"`
}

// settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse]
type settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponse) implementsSettingAccountNameserverSetGetResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvanced bool

const (
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvanced = false
)

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseAdvancedFalse:
		return true
	}
	return false
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                       `json:"name" api:"required"`
	JSON settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserver]
type settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseDNSSettingsNameserverSetStandardResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse struct {
	// Identifier for a nameserver set.
	ID string `json:"id" api:"required"`
	// Whether the nameserver set uses Advanced anycast groups.
	Advanced SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvanced `json:"advanced" api:"required"`
	// When the nameserver set was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet       int64                                                                                      `json:"ip_set" api:"required"`
	Nameservers []SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserver `json:"nameservers" api:"required"`
	JSON        settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseJSON         `json:"-"`
}

// settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse]
type settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseJSON struct {
	ID          apijson.Field
	Advanced    apijson.Field
	CreatedOn   apijson.Field
	IPSet       apijson.Field
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseJSON) RawJSON() string {
	return r.raw
}

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponse) implementsSettingAccountNameserverSetGetResponse() {
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvanced bool

const (
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvanced = true
)

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserver struct {
	// IPv4 addresses assigned to the nameserver.
	IPV4 []string `json:"ipv4" api:"required" format:"ipv4"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV4Groups []SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group `json:"ipv4_groups" api:"required"`
	// IPv6 addresses assigned to the nameserver.
	IPV6 []string `json:"ipv6" api:"required" format:"ipv6"`
	// Advanced anycast group for each address in the corresponding address array.
	// Entries have the same order as, and correspond one-to-one with, the addresses.
	IPV6Groups []SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group `json:"ipv6_groups" api:"required"`
	// A unique lowercase Punycode nameserver name within the set.
	Name string                                                                                       `json:"name" api:"required"`
	JSON settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON `json:"-"`
}

// settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON
// contains the JSON metadata for the struct
// [SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserver]
type settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON struct {
	IPV4        apijson.Field
	IPV4Groups  apijson.Field
	IPV6        apijson.Field
	IPV6Groups  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserverJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group string

const (
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "a"
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "b"
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group = "c"
)

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupA, SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupB, SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV4GroupC:
		return true
	}
	return false
}

type SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group string

const (
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "a"
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "b"
	SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group = "c"
)

func (r SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6Group) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupA, SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupB, SettingAccountNameserverSetGetResponseDNSSettingsNameserverSetAdvancedResponseNameserversIPV6GroupC:
		return true
	}
	return false
}

// Whether the nameserver set uses Advanced anycast groups.
type SettingAccountNameserverSetGetResponseAdvanced bool

const (
	SettingAccountNameserverSetGetResponseAdvancedFalse SettingAccountNameserverSetGetResponseAdvanced = false
	SettingAccountNameserverSetGetResponseAdvancedTrue  SettingAccountNameserverSetGetResponseAdvanced = true
)

func (r SettingAccountNameserverSetGetResponseAdvanced) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseAdvancedFalse, SettingAccountNameserverSetGetResponseAdvancedTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Lists each nameserver and the number of addresses to allocate to it. Requires a
	// unique name for each entry in the set.
	Nameservers param.Field[[]SettingAccountNameserverSetNewParamsNameserver] `json:"nameservers" api:"required"`
	// Whether to allocate the nameservers from distinct Advanced anycast groups.
	Advanced param.Field[bool] `json:"advanced"`
	// Selects the account-specific IP set that supplies the nameserver addresses. The
	// account's entitlement determines the maximum value. Nameserver sets with the
	// same `ip_set` and `advanced` value may reuse addresses; otherwise, they use
	// disjoint address groups.
	IPSet param.Field[int64] `json:"ip_set"`
}

func (r SettingAccountNameserverSetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAccountNameserverSetNewParamsNameserver struct {
	// A unique lowercase Punycode nameserver name within the set.
	Name param.Field[string] `json:"name" api:"required"`
	// Number of IPv4 addresses and number of IPv6 addresses to allocate to this
	// nameserver. All nameservers in an Advanced nameserver set must use the same
	// value.
	IPCount param.Field[int64] `json:"ip_count"`
}

func (r SettingAccountNameserverSetNewParamsNameserver) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAccountNameserverSetNewResponseEnvelope struct {
	Errors   []SettingAccountNameserverSetNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAccountNameserverSetNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingAccountNameserverSetNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success SettingAccountNameserverSetNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    settingAccountNameserverSetNewResponseEnvelopeJSON    `json:"-"`
}

// settingAccountNameserverSetNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAccountNameserverSetNewResponseEnvelope]
type settingAccountNameserverSetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           SettingAccountNameserverSetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountNameserverSetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountNameserverSetNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingAccountNameserverSetNewResponseEnvelopeErrors]
type settingAccountNameserverSetNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    settingAccountNameserverSetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountNameserverSetNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [SettingAccountNameserverSetNewResponseEnvelopeErrorsSource]
type settingAccountNameserverSetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           SettingAccountNameserverSetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountNameserverSetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountNameserverSetNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingAccountNameserverSetNewResponseEnvelopeMessages]
type settingAccountNameserverSetNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    settingAccountNameserverSetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountNameserverSetNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingAccountNameserverSetNewResponseEnvelopeMessagesSource]
type settingAccountNameserverSetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountNameserverSetNewResponseEnvelopeSuccess bool

const (
	SettingAccountNameserverSetNewResponseEnvelopeSuccessTrue SettingAccountNameserverSetNewResponseEnvelopeSuccess = true
)

func (r SettingAccountNameserverSetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SettingAccountNameserverSetListParams]'s query parameters
// as `url.Values`.
func (r SettingAccountNameserverSetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SettingAccountNameserverSetDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingAccountNameserverSetDeleteResponseEnvelope struct {
	Errors   []SettingAccountNameserverSetDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAccountNameserverSetDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingAccountNameserverSetDeleteResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success SettingAccountNameserverSetDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    settingAccountNameserverSetDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingAccountNameserverSetDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAccountNameserverSetDeleteResponseEnvelope]
type settingAccountNameserverSetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetDeleteResponseEnvelopeErrors struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           SettingAccountNameserverSetDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountNameserverSetDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountNameserverSetDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingAccountNameserverSetDeleteResponseEnvelopeErrors]
type settingAccountNameserverSetDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    settingAccountNameserverSetDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountNameserverSetDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingAccountNameserverSetDeleteResponseEnvelopeErrorsSource]
type settingAccountNameserverSetDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetDeleteResponseEnvelopeMessages struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           SettingAccountNameserverSetDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountNameserverSetDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountNameserverSetDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAccountNameserverSetDeleteResponseEnvelopeMessages]
type settingAccountNameserverSetDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    settingAccountNameserverSetDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountNameserverSetDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingAccountNameserverSetDeleteResponseEnvelopeMessagesSource]
type settingAccountNameserverSetDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountNameserverSetDeleteResponseEnvelopeSuccess bool

const (
	SettingAccountNameserverSetDeleteResponseEnvelopeSuccessTrue SettingAccountNameserverSetDeleteResponseEnvelopeSuccess = true
)

func (r SettingAccountNameserverSetDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAccountNameserverSetGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingAccountNameserverSetGetResponseEnvelope struct {
	Errors   []SettingAccountNameserverSetGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingAccountNameserverSetGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   SettingAccountNameserverSetGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success SettingAccountNameserverSetGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    settingAccountNameserverSetGetResponseEnvelopeJSON    `json:"-"`
}

// settingAccountNameserverSetGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAccountNameserverSetGetResponseEnvelope]
type settingAccountNameserverSetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           SettingAccountNameserverSetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountNameserverSetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountNameserverSetGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingAccountNameserverSetGetResponseEnvelopeErrors]
type settingAccountNameserverSetGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    settingAccountNameserverSetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountNameserverSetGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [SettingAccountNameserverSetGetResponseEnvelopeErrorsSource]
type settingAccountNameserverSetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           SettingAccountNameserverSetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountNameserverSetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountNameserverSetGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingAccountNameserverSetGetResponseEnvelopeMessages]
type settingAccountNameserverSetGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountNameserverSetGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    settingAccountNameserverSetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountNameserverSetGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingAccountNameserverSetGetResponseEnvelopeMessagesSource]
type settingAccountNameserverSetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountNameserverSetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountNameserverSetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountNameserverSetGetResponseEnvelopeSuccess bool

const (
	SettingAccountNameserverSetGetResponseEnvelopeSuccessTrue SettingAccountNameserverSetGetResponseEnvelopeSuccess = true
)

func (r SettingAccountNameserverSetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountNameserverSetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
