// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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

// SiteACLService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteACLService] method instead.
type SiteACLService struct {
	Options []option.RequestOption
}

// NewSiteACLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteACLService(opts ...option.RequestOption) (r *SiteACLService) {
	r = &SiteACLService{}
	r.Options = opts
	return
}

// Creates a new Site ACL.
func (r *SiteACLService) New(ctx context.Context, siteID string, params SiteACLNewParams, opts ...option.RequestOption) (res *SiteACLNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Site ACL.
func (r *SiteACLService) Update(ctx context.Context, siteID string, aclIdentifier string, params SiteACLUpdateParams, opts ...option.RequestOption) (res *SiteACLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", params.AccountID, siteID, aclIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Site ACLs associated with an account.
func (r *SiteACLService) List(ctx context.Context, siteID string, query SiteACLListParams, opts ...option.RequestOption) (res *SiteACLListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific Site ACL.
func (r *SiteACLService) Delete(ctx context.Context, siteID string, aclIdentifier string, params SiteACLDeleteParams, opts ...option.RequestOption) (res *SiteACLDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", params.AccountID, siteID, aclIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site ACL.
func (r *SiteACLService) Get(ctx context.Context, siteID string, aclIdentifier string, query SiteACLGetParams, opts ...option.RequestOption) (res *SiteACLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", query.AccountID, siteID, aclIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SiteACLNewResponse struct {
	ACLs []SiteACLNewResponseACL `json:"acls"`
	JSON siteACLNewResponseJSON  `json:"-"`
}

// siteACLNewResponseJSON contains the JSON metadata for the struct
// [SiteACLNewResponse]
type siteACLNewResponseJSON struct {
	ACLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseJSON) RawJSON() string {
	return r.raw
}

// Bidirectional ACL policy for network traffic within a site.
type SiteACLNewResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally bool                       `json:"forward_locally"`
	LAN1           SiteACLNewResponseACLsLAN1 `json:"lan_1"`
	LAN2           SiteACLNewResponseACLsLAN2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                      `json:"name"`
	Protocols []shared.UnnamedSchemaRef82 `json:"protocols"`
	JSON      siteACLNewResponseACLJSON   `json:"-"`
}

// siteACLNewResponseACLJSON contains the JSON metadata for the struct
// [SiteACLNewResponseACL]
type siteACLNewResponseACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SiteACLNewResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLNewResponseACLsLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLNewResponseACLsLAN1SubnetsUnion `json:"subnets"`
	JSON    siteACLNewResponseACLsLan1JSON           `json:"-"`
}

// siteACLNewResponseACLsLan1JSON contains the JSON metadata for the struct
// [SiteACLNewResponseACLsLAN1]
type siteACLNewResponseACLsLan1JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseACLsLAN1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLsLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLNewResponseACLsLAN1SubnetsUnion interface {
	ImplementsMagicTransitSiteACLNewResponseACLsLan1SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLNewResponseACLsLAN1SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLNewResponseACLsLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLNewResponseACLsLAN2SubnetsUnion `json:"subnets"`
	JSON    siteACLNewResponseACLsLan2JSON           `json:"-"`
}

// siteACLNewResponseACLsLan2JSON contains the JSON metadata for the struct
// [SiteACLNewResponseACLsLAN2]
type siteACLNewResponseACLsLan2JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseACLsLAN2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLsLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLNewResponseACLsLAN2SubnetsUnion interface {
	ImplementsMagicTransitSiteACLNewResponseACLsLan2SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLNewResponseACLsLAN2SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLUpdateResponse struct {
	// Bidirectional ACL policy for network traffic within a site.
	ACL  SiteACLUpdateResponseACL  `json:"acl"`
	JSON siteACLUpdateResponseJSON `json:"-"`
}

// siteACLUpdateResponseJSON contains the JSON metadata for the struct
// [SiteACLUpdateResponse]
type siteACLUpdateResponseJSON struct {
	ACL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Bidirectional ACL policy for network traffic within a site.
type SiteACLUpdateResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally bool                         `json:"forward_locally"`
	LAN1           SiteACLUpdateResponseACLLAN1 `json:"lan_1"`
	LAN2           SiteACLUpdateResponseACLLAN2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                       `json:"name"`
	Protocols []shared.UnnamedSchemaRef82  `json:"protocols"`
	JSON      siteACLUpdateResponseACLJSON `json:"-"`
}

// siteACLUpdateResponseACLJSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACL]
type siteACLUpdateResponseACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLUpdateResponseACLLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLUpdateResponseACLLAN1SubnetsUnion `json:"subnets"`
	JSON    siteACLUpdateResponseAcllan1JSON           `json:"-"`
}

// siteACLUpdateResponseAcllan1JSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACLLAN1]
type siteACLUpdateResponseAcllan1JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACLLAN1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseAcllan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLUpdateResponseACLLAN1SubnetsUnion interface {
	ImplementsMagicTransitSiteACLUpdateResponseAcllan1SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLUpdateResponseACLLAN1SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLUpdateResponseACLLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLUpdateResponseACLLAN2SubnetsUnion `json:"subnets"`
	JSON    siteACLUpdateResponseAcllan2JSON           `json:"-"`
}

// siteACLUpdateResponseAcllan2JSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACLLAN2]
type siteACLUpdateResponseAcllan2JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACLLAN2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseAcllan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLUpdateResponseACLLAN2SubnetsUnion interface {
	ImplementsMagicTransitSiteACLUpdateResponseAcllan2SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLUpdateResponseACLLAN2SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLListResponse struct {
	ACLs []SiteACLListResponseACL `json:"acls"`
	JSON siteACLListResponseJSON  `json:"-"`
}

// siteACLListResponseJSON contains the JSON metadata for the struct
// [SiteACLListResponse]
type siteACLListResponseJSON struct {
	ACLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseJSON) RawJSON() string {
	return r.raw
}

// Bidirectional ACL policy for network traffic within a site.
type SiteACLListResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally bool                        `json:"forward_locally"`
	LAN1           SiteACLListResponseACLsLAN1 `json:"lan_1"`
	LAN2           SiteACLListResponseACLsLAN2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                      `json:"name"`
	Protocols []shared.UnnamedSchemaRef82 `json:"protocols"`
	JSON      siteACLListResponseACLJSON  `json:"-"`
}

// siteACLListResponseACLJSON contains the JSON metadata for the struct
// [SiteACLListResponseACL]
type siteACLListResponseACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SiteACLListResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLListResponseACLsLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLListResponseACLsLAN1SubnetsUnion `json:"subnets"`
	JSON    siteACLListResponseACLsLan1JSON           `json:"-"`
}

// siteACLListResponseACLsLan1JSON contains the JSON metadata for the struct
// [SiteACLListResponseACLsLAN1]
type siteACLListResponseACLsLan1JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseACLsLAN1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLsLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLListResponseACLsLAN1SubnetsUnion interface {
	ImplementsMagicTransitSiteACLListResponseACLsLan1SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLListResponseACLsLAN1SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLListResponseACLsLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLListResponseACLsLAN2SubnetsUnion `json:"subnets"`
	JSON    siteACLListResponseACLsLan2JSON           `json:"-"`
}

// siteACLListResponseACLsLan2JSON contains the JSON metadata for the struct
// [SiteACLListResponseACLsLAN2]
type siteACLListResponseACLsLan2JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseACLsLAN2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLsLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLListResponseACLsLAN2SubnetsUnion interface {
	ImplementsMagicTransitSiteACLListResponseACLsLan2SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLListResponseACLsLAN2SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLDeleteResponse struct {
	Deleted bool `json:"deleted"`
	// Bidirectional ACL policy for network traffic within a site.
	DeletedACL SiteACLDeleteResponseDeletedACL `json:"deleted_acl"`
	JSON       siteACLDeleteResponseJSON       `json:"-"`
}

// siteACLDeleteResponseJSON contains the JSON metadata for the struct
// [SiteACLDeleteResponse]
type siteACLDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedACL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Bidirectional ACL policy for network traffic within a site.
type SiteACLDeleteResponseDeletedACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally bool                                `json:"forward_locally"`
	LAN1           SiteACLDeleteResponseDeletedACLLAN1 `json:"lan_1"`
	LAN2           SiteACLDeleteResponseDeletedACLLAN2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                              `json:"name"`
	Protocols []shared.UnnamedSchemaRef82         `json:"protocols"`
	JSON      siteACLDeleteResponseDeletedACLJSON `json:"-"`
}

// siteACLDeleteResponseDeletedACLJSON contains the JSON metadata for the struct
// [SiteACLDeleteResponseDeletedACL]
type siteACLDeleteResponseDeletedACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLDeleteResponseDeletedACLLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLDeleteResponseDeletedACLLAN1SubnetsUnion `json:"subnets"`
	JSON    siteACLDeleteResponseDeletedAcllan1JSON           `json:"-"`
}

// siteACLDeleteResponseDeletedAcllan1JSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseDeletedACLLAN1]
type siteACLDeleteResponseDeletedAcllan1JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACLLAN1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedAcllan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLDeleteResponseDeletedACLLAN1SubnetsUnion interface {
	ImplementsMagicTransitSiteACLDeleteResponseDeletedAcllan1SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLDeleteResponseDeletedACLLAN1SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLDeleteResponseDeletedACLLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLDeleteResponseDeletedACLLAN2SubnetsUnion `json:"subnets"`
	JSON    siteACLDeleteResponseDeletedAcllan2JSON           `json:"-"`
}

// siteACLDeleteResponseDeletedAcllan2JSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseDeletedACLLAN2]
type siteACLDeleteResponseDeletedAcllan2JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACLLAN2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedAcllan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLDeleteResponseDeletedACLLAN2SubnetsUnion interface {
	ImplementsMagicTransitSiteACLDeleteResponseDeletedAcllan2SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLDeleteResponseDeletedACLLAN2SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLGetResponse struct {
	// Bidirectional ACL policy for network traffic within a site.
	ACL  SiteACLGetResponseACL  `json:"acl"`
	JSON siteACLGetResponseJSON `json:"-"`
}

// siteACLGetResponseJSON contains the JSON metadata for the struct
// [SiteACLGetResponse]
type siteACLGetResponseJSON struct {
	ACL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseJSON) RawJSON() string {
	return r.raw
}

// Bidirectional ACL policy for network traffic within a site.
type SiteACLGetResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally bool                      `json:"forward_locally"`
	LAN1           SiteACLGetResponseACLLAN1 `json:"lan_1"`
	LAN2           SiteACLGetResponseACLLAN2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                      `json:"name"`
	Protocols []shared.UnnamedSchemaRef82 `json:"protocols"`
	JSON      siteACLGetResponseACLJSON   `json:"-"`
}

// siteACLGetResponseACLJSON contains the JSON metadata for the struct
// [SiteACLGetResponseACL]
type siteACLGetResponseACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SiteACLGetResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLGetResponseACLLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLGetResponseACLLAN1SubnetsUnion `json:"subnets"`
	JSON    siteACLGetResponseAcllan1JSON           `json:"-"`
}

// siteACLGetResponseAcllan1JSON contains the JSON metadata for the struct
// [SiteACLGetResponseACLLAN1]
type siteACLGetResponseAcllan1JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseACLLAN1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseAcllan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLGetResponseACLLAN1SubnetsUnion interface {
	ImplementsMagicTransitSiteACLGetResponseAcllan1SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLGetResponseACLLAN1SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLGetResponseACLLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLGetResponseACLLAN2SubnetsUnion `json:"subnets"`
	JSON    siteACLGetResponseAcllan2JSON           `json:"-"`
}

// siteACLGetResponseAcllan2JSON contains the JSON metadata for the struct
// [SiteACLGetResponseACLLAN2]
type siteACLGetResponseAcllan2JSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseACLLAN2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseAcllan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLGetResponseACLLAN2SubnetsUnion interface {
	ImplementsMagicTransitSiteACLGetResponseAcllan2SubnetsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLGetResponseACLLAN2SubnetsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SiteACLNewParams struct {
	// Identifier
	AccountID param.Field[string]              `path:"account_id,required"`
	ACL       param.Field[SiteACLNewParamsACL] `json:"acl"`
}

func (r SiteACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLNewParamsACL struct {
	LAN1 param.Field[SiteACLNewParamsACLLAN1] `json:"lan_1,required"`
	LAN2 param.Field[SiteACLNewParamsACLLAN2] `json:"lan_2,required"`
	// The name of the ACL.
	Name param.Field[string] `json:"name,required"`
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally param.Field[bool]                        `json:"forward_locally"`
	Protocols      param.Field[[]shared.UnnamedSchemaRef82] `json:"protocols"`
}

func (r SiteACLNewParamsACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLNewParamsACLLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLNewParamsACLLAN1SubnetUnion] `json:"subnets"`
}

func (r SiteACLNewParamsACLLAN1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLNewParamsACLLAN1SubnetUnion interface {
	ImplementsMagicTransitSiteACLNewParamsAcllan1SubnetUnion()
}

type SiteACLNewParamsACLLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLNewParamsACLLAN2SubnetUnion] `json:"subnets"`
}

func (r SiteACLNewParamsACLLAN2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLNewParamsACLLAN2SubnetUnion interface {
	ImplementsMagicTransitSiteACLNewParamsAcllan2SubnetUnion()
}

type SiteACLNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteACLNewResponse    `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLNewResponseEnvelopeJSON    `json:"-"`
}

// siteACLNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLNewResponseEnvelope]
type siteACLNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLNewResponseEnvelopeSuccess bool

const (
	SiteACLNewResponseEnvelopeSuccessTrue SiteACLNewResponseEnvelopeSuccess = true
)

func (r SiteACLNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                 `path:"account_id,required"`
	ACL       param.Field[SiteACLUpdateParamsACL] `json:"acl"`
}

func (r SiteACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLUpdateParamsACL struct {
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic WAN Connector. If not included in request, will
	// default to false.
	ForwardLocally param.Field[bool]                       `json:"forward_locally"`
	LAN1           param.Field[SiteACLUpdateParamsACLLAN1] `json:"lan_1"`
	LAN2           param.Field[SiteACLUpdateParamsACLLAN2] `json:"lan_2"`
	// The name of the ACL.
	Name      param.Field[string]                      `json:"name"`
	Protocols param.Field[[]shared.UnnamedSchemaRef82] `json:"protocols"`
}

func (r SiteACLUpdateParamsACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLUpdateParamsACLLAN1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLUpdateParamsACLLAN1SubnetUnion] `json:"subnets"`
}

func (r SiteACLUpdateParamsACLLAN1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLUpdateParamsACLLAN1SubnetUnion interface {
	ImplementsMagicTransitSiteACLUpdateParamsAcllan1SubnetUnion()
}

type SiteACLUpdateParamsACLLAN2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLUpdateParamsACLLAN2SubnetUnion] `json:"subnets"`
}

func (r SiteACLUpdateParamsACLLAN2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLUpdateParamsACLLAN2SubnetUnion interface {
	ImplementsMagicTransitSiteACLUpdateParamsAcllan2SubnetUnion()
}

type SiteACLUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteACLUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteACLUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseEnvelope]
type siteACLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLUpdateResponseEnvelopeSuccess bool

const (
	SiteACLUpdateResponseEnvelopeSuccessTrue SiteACLUpdateResponseEnvelopeSuccess = true
)

func (r SiteACLUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteACLListResponse   `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLListResponseEnvelopeJSON    `json:"-"`
}

// siteACLListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLListResponseEnvelope]
type siteACLListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLListResponseEnvelopeSuccess bool

const (
	SiteACLListResponseEnvelopeSuccessTrue SiteACLListResponseEnvelopeSuccess = true
)

func (r SiteACLListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SiteACLDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SiteACLDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteACLDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteACLDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLDeleteResponseEnvelope]
type siteACLDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLDeleteResponseEnvelopeSuccess bool

const (
	SiteACLDeleteResponseEnvelopeSuccessTrue SiteACLDeleteResponseEnvelopeSuccess = true
)

func (r SiteACLDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteACLGetResponse    `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLGetResponseEnvelopeJSON    `json:"-"`
}

// siteACLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLGetResponseEnvelope]
type siteACLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLGetResponseEnvelopeSuccess bool

const (
	SiteACLGetResponseEnvelopeSuccessTrue SiteACLGetResponseEnvelopeSuccess = true
)

func (r SiteACLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
