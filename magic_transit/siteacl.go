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
func (r *SiteACLService) Delete(ctx context.Context, siteID string, aclIdentifier string, body SiteACLDeleteParams, opts ...option.RequestOption) (res *SiteACLDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteACLDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", body.AccountID, siteID, aclIdentifier)
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

// Bidirectional ACL policy for local network traffic within a site.
type SiteACLNewResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string                     `json:"description"`
	Lan1        SiteACLNewResponseACLsLan1 `json:"lan_1"`
	Lan2        SiteACLNewResponseACLsLan2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                           `json:"name"`
	Protocols []SiteACLNewResponseACLsProtocol `json:"protocols"`
	JSON      siteACLNewResponseACLJSON        `json:"-"`
}

// siteACLNewResponseACLJSON contains the JSON metadata for the struct
// [SiteACLNewResponseACL]
type siteACLNewResponseACLJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Lan1        apijson.Field
	Lan2        apijson.Field
	Name        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLNewResponseACLsLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLNewResponseACLsLan1Subnet `json:"subnets"`
	JSON    siteACLNewResponseACLsLan1JSON     `json:"-"`
}

// siteACLNewResponseACLsLan1JSON contains the JSON metadata for the struct
// [SiteACLNewResponseACLsLan1]
type siteACLNewResponseACLsLan1JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseACLsLan1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLsLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLNewResponseACLsLan1Subnet interface {
	ImplementsMagicTransitSiteACLNewResponseACLsLan1Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLNewResponseACLsLan1Subnet)(nil)).Elem(),
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

type SiteACLNewResponseACLsLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLNewResponseACLsLan2Subnet `json:"subnets"`
	JSON    siteACLNewResponseACLsLan2JSON     `json:"-"`
}

// siteACLNewResponseACLsLan2JSON contains the JSON metadata for the struct
// [SiteACLNewResponseACLsLan2]
type siteACLNewResponseACLsLan2JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseACLsLan2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseACLsLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLNewResponseACLsLan2Subnet interface {
	ImplementsMagicTransitSiteACLNewResponseACLsLan2Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLNewResponseACLsLan2Subnet)(nil)).Elem(),
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

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLNewResponseACLsProtocol string

const (
	SiteACLNewResponseACLsProtocolTcp  SiteACLNewResponseACLsProtocol = "tcp"
	SiteACLNewResponseACLsProtocolUdp  SiteACLNewResponseACLsProtocol = "udp"
	SiteACLNewResponseACLsProtocolIcmp SiteACLNewResponseACLsProtocol = "icmp"
)

func (r SiteACLNewResponseACLsProtocol) IsKnown() bool {
	switch r {
	case SiteACLNewResponseACLsProtocolTcp, SiteACLNewResponseACLsProtocolUdp, SiteACLNewResponseACLsProtocolIcmp:
		return true
	}
	return false
}

type SiteACLUpdateResponse struct {
	// Bidirectional ACL policy for local network traffic within a site.
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

// Bidirectional ACL policy for local network traffic within a site.
type SiteACLUpdateResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string                       `json:"description"`
	Lan1        SiteACLUpdateResponseACLLan1 `json:"lan_1"`
	Lan2        SiteACLUpdateResponseACLLan2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                             `json:"name"`
	Protocols []SiteACLUpdateResponseACLProtocol `json:"protocols"`
	JSON      siteACLUpdateResponseACLJSON       `json:"-"`
}

// siteACLUpdateResponseACLJSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACL]
type siteACLUpdateResponseACLJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Lan1        apijson.Field
	Lan2        apijson.Field
	Name        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLUpdateResponseACLLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLUpdateResponseACLLan1Subnet `json:"subnets"`
	JSON    siteACLUpdateResponseACLLan1JSON     `json:"-"`
}

// siteACLUpdateResponseACLLan1JSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACLLan1]
type siteACLUpdateResponseACLLan1JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACLLan1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseACLLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLUpdateResponseACLLan1Subnet interface {
	ImplementsMagicTransitSiteACLUpdateResponseACLLan1Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLUpdateResponseACLLan1Subnet)(nil)).Elem(),
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

type SiteACLUpdateResponseACLLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLUpdateResponseACLLan2Subnet `json:"subnets"`
	JSON    siteACLUpdateResponseACLLan2JSON     `json:"-"`
}

// siteACLUpdateResponseACLLan2JSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseACLLan2]
type siteACLUpdateResponseACLLan2JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseACLLan2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseACLLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLUpdateResponseACLLan2Subnet interface {
	ImplementsMagicTransitSiteACLUpdateResponseACLLan2Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLUpdateResponseACLLan2Subnet)(nil)).Elem(),
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

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLUpdateResponseACLProtocol string

const (
	SiteACLUpdateResponseACLProtocolTcp  SiteACLUpdateResponseACLProtocol = "tcp"
	SiteACLUpdateResponseACLProtocolUdp  SiteACLUpdateResponseACLProtocol = "udp"
	SiteACLUpdateResponseACLProtocolIcmp SiteACLUpdateResponseACLProtocol = "icmp"
)

func (r SiteACLUpdateResponseACLProtocol) IsKnown() bool {
	switch r {
	case SiteACLUpdateResponseACLProtocolTcp, SiteACLUpdateResponseACLProtocolUdp, SiteACLUpdateResponseACLProtocolIcmp:
		return true
	}
	return false
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

// Bidirectional ACL policy for local network traffic within a site.
type SiteACLListResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string                      `json:"description"`
	Lan1        SiteACLListResponseACLsLan1 `json:"lan_1"`
	Lan2        SiteACLListResponseACLsLan2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                            `json:"name"`
	Protocols []SiteACLListResponseACLsProtocol `json:"protocols"`
	JSON      siteACLListResponseACLJSON        `json:"-"`
}

// siteACLListResponseACLJSON contains the JSON metadata for the struct
// [SiteACLListResponseACL]
type siteACLListResponseACLJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Lan1        apijson.Field
	Lan2        apijson.Field
	Name        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLListResponseACLsLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLListResponseACLsLan1Subnet `json:"subnets"`
	JSON    siteACLListResponseACLsLan1JSON     `json:"-"`
}

// siteACLListResponseACLsLan1JSON contains the JSON metadata for the struct
// [SiteACLListResponseACLsLan1]
type siteACLListResponseACLsLan1JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseACLsLan1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLsLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLListResponseACLsLan1Subnet interface {
	ImplementsMagicTransitSiteACLListResponseACLsLan1Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLListResponseACLsLan1Subnet)(nil)).Elem(),
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

type SiteACLListResponseACLsLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLListResponseACLsLan2Subnet `json:"subnets"`
	JSON    siteACLListResponseACLsLan2JSON     `json:"-"`
}

// siteACLListResponseACLsLan2JSON contains the JSON metadata for the struct
// [SiteACLListResponseACLsLan2]
type siteACLListResponseACLsLan2JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseACLsLan2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseACLsLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLListResponseACLsLan2Subnet interface {
	ImplementsMagicTransitSiteACLListResponseACLsLan2Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLListResponseACLsLan2Subnet)(nil)).Elem(),
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

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLListResponseACLsProtocol string

const (
	SiteACLListResponseACLsProtocolTcp  SiteACLListResponseACLsProtocol = "tcp"
	SiteACLListResponseACLsProtocolUdp  SiteACLListResponseACLsProtocol = "udp"
	SiteACLListResponseACLsProtocolIcmp SiteACLListResponseACLsProtocol = "icmp"
)

func (r SiteACLListResponseACLsProtocol) IsKnown() bool {
	switch r {
	case SiteACLListResponseACLsProtocolTcp, SiteACLListResponseACLsProtocolUdp, SiteACLListResponseACLsProtocolIcmp:
		return true
	}
	return false
}

type SiteACLDeleteResponse struct {
	Deleted bool `json:"deleted"`
	// Bidirectional ACL policy for local network traffic within a site.
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

// Bidirectional ACL policy for local network traffic within a site.
type SiteACLDeleteResponseDeletedACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string                              `json:"description"`
	Lan1        SiteACLDeleteResponseDeletedACLLan1 `json:"lan_1"`
	Lan2        SiteACLDeleteResponseDeletedACLLan2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                                    `json:"name"`
	Protocols []SiteACLDeleteResponseDeletedACLProtocol `json:"protocols"`
	JSON      siteACLDeleteResponseDeletedACLJSON       `json:"-"`
}

// siteACLDeleteResponseDeletedACLJSON contains the JSON metadata for the struct
// [SiteACLDeleteResponseDeletedACL]
type siteACLDeleteResponseDeletedACLJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Lan1        apijson.Field
	Lan2        apijson.Field
	Name        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLDeleteResponseDeletedACLLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLDeleteResponseDeletedACLLan1Subnet `json:"subnets"`
	JSON    siteACLDeleteResponseDeletedACLLan1JSON     `json:"-"`
}

// siteACLDeleteResponseDeletedACLLan1JSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseDeletedACLLan1]
type siteACLDeleteResponseDeletedACLLan1JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACLLan1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedACLLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLDeleteResponseDeletedACLLan1Subnet interface {
	ImplementsMagicTransitSiteACLDeleteResponseDeletedACLLan1Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLDeleteResponseDeletedACLLan1Subnet)(nil)).Elem(),
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

type SiteACLDeleteResponseDeletedACLLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLDeleteResponseDeletedACLLan2Subnet `json:"subnets"`
	JSON    siteACLDeleteResponseDeletedACLLan2JSON     `json:"-"`
}

// siteACLDeleteResponseDeletedACLLan2JSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseDeletedACLLan2]
type siteACLDeleteResponseDeletedACLLan2JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseDeletedACLLan2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseDeletedACLLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLDeleteResponseDeletedACLLan2Subnet interface {
	ImplementsMagicTransitSiteACLDeleteResponseDeletedACLLan2Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLDeleteResponseDeletedACLLan2Subnet)(nil)).Elem(),
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

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLDeleteResponseDeletedACLProtocol string

const (
	SiteACLDeleteResponseDeletedACLProtocolTcp  SiteACLDeleteResponseDeletedACLProtocol = "tcp"
	SiteACLDeleteResponseDeletedACLProtocolUdp  SiteACLDeleteResponseDeletedACLProtocol = "udp"
	SiteACLDeleteResponseDeletedACLProtocolIcmp SiteACLDeleteResponseDeletedACLProtocol = "icmp"
)

func (r SiteACLDeleteResponseDeletedACLProtocol) IsKnown() bool {
	switch r {
	case SiteACLDeleteResponseDeletedACLProtocolTcp, SiteACLDeleteResponseDeletedACLProtocolUdp, SiteACLDeleteResponseDeletedACLProtocolIcmp:
		return true
	}
	return false
}

type SiteACLGetResponse struct {
	// Bidirectional ACL policy for local network traffic within a site.
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

// Bidirectional ACL policy for local network traffic within a site.
type SiteACLGetResponseACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string                    `json:"description"`
	Lan1        SiteACLGetResponseACLLan1 `json:"lan_1"`
	Lan2        SiteACLGetResponseACLLan2 `json:"lan_2"`
	// The name of the ACL.
	Name      string                          `json:"name"`
	Protocols []SiteACLGetResponseACLProtocol `json:"protocols"`
	JSON      siteACLGetResponseACLJSON       `json:"-"`
}

// siteACLGetResponseACLJSON contains the JSON metadata for the struct
// [SiteACLGetResponseACL]
type siteACLGetResponseACLJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Lan1        apijson.Field
	Lan2        apijson.Field
	Name        apijson.Field
	Protocols   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseACLJSON) RawJSON() string {
	return r.raw
}

type SiteACLGetResponseACLLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLGetResponseACLLan1Subnet `json:"subnets"`
	JSON    siteACLGetResponseACLLan1JSON     `json:"-"`
}

// siteACLGetResponseACLLan1JSON contains the JSON metadata for the struct
// [SiteACLGetResponseACLLan1]
type siteACLGetResponseACLLan1JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseACLLan1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseACLLan1JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLGetResponseACLLan1Subnet interface {
	ImplementsMagicTransitSiteACLGetResponseACLLan1Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLGetResponseACLLan1Subnet)(nil)).Elem(),
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

type SiteACLGetResponseACLLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []SiteACLGetResponseACLLan2Subnet `json:"subnets"`
	JSON    siteACLGetResponseACLLan2JSON     `json:"-"`
}

// siteACLGetResponseACLLan2JSON contains the JSON metadata for the struct
// [SiteACLGetResponseACLLan2]
type siteACLGetResponseACLLan2JSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseACLLan2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseACLLan2JSON) RawJSON() string {
	return r.raw
}

// A valid IPv4 address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type SiteACLGetResponseACLLan2Subnet interface {
	ImplementsMagicTransitSiteACLGetResponseACLLan2Subnet()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SiteACLGetResponseACLLan2Subnet)(nil)).Elem(),
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

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLGetResponseACLProtocol string

const (
	SiteACLGetResponseACLProtocolTcp  SiteACLGetResponseACLProtocol = "tcp"
	SiteACLGetResponseACLProtocolUdp  SiteACLGetResponseACLProtocol = "udp"
	SiteACLGetResponseACLProtocolIcmp SiteACLGetResponseACLProtocol = "icmp"
)

func (r SiteACLGetResponseACLProtocol) IsKnown() bool {
	switch r {
	case SiteACLGetResponseACLProtocolTcp, SiteACLGetResponseACLProtocolUdp, SiteACLGetResponseACLProtocolIcmp:
		return true
	}
	return false
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
	Lan1 param.Field[SiteACLNewParamsACLLan1] `json:"lan_1,required"`
	Lan2 param.Field[SiteACLNewParamsACLLan2] `json:"lan_2,required"`
	// The name of the ACL.
	Name param.Field[string] `json:"name,required"`
	// Description for the ACL.
	Description param.Field[string]                        `json:"description"`
	Protocols   param.Field[[]SiteACLNewParamsACLProtocol] `json:"protocols"`
}

func (r SiteACLNewParamsACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLNewParamsACLLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLNewParamsACLLan1Subnet] `json:"subnets"`
}

func (r SiteACLNewParamsACLLan1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLNewParamsACLLan1Subnet interface {
	ImplementsMagicTransitSiteACLNewParamsACLLan1Subnet()
}

type SiteACLNewParamsACLLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLNewParamsACLLan2Subnet] `json:"subnets"`
}

func (r SiteACLNewParamsACLLan2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLNewParamsACLLan2Subnet interface {
	ImplementsMagicTransitSiteACLNewParamsACLLan2Subnet()
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLNewParamsACLProtocol string

const (
	SiteACLNewParamsACLProtocolTcp  SiteACLNewParamsACLProtocol = "tcp"
	SiteACLNewParamsACLProtocolUdp  SiteACLNewParamsACLProtocol = "udp"
	SiteACLNewParamsACLProtocolIcmp SiteACLNewParamsACLProtocol = "icmp"
)

func (r SiteACLNewParamsACLProtocol) IsKnown() bool {
	switch r {
	case SiteACLNewParamsACLProtocolTcp, SiteACLNewParamsACLProtocolUdp, SiteACLNewParamsACLProtocolIcmp:
		return true
	}
	return false
}

type SiteACLNewResponseEnvelope struct {
	Errors   []SiteACLNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteACLNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteACLNewResponse                   `json:"result,required"`
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

type SiteACLNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteACLNewResponseEnvelopeErrorsJSON `json:"-"`
}

// siteACLNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteACLNewResponseEnvelopeErrors]
type siteACLNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteACLNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteACLNewResponseEnvelopeMessagesJSON `json:"-"`
}

// siteACLNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteACLNewResponseEnvelopeMessages]
type siteACLNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Description param.Field[string]                     `json:"description"`
	Lan1        param.Field[SiteACLUpdateParamsACLLan1] `json:"lan_1"`
	Lan2        param.Field[SiteACLUpdateParamsACLLan2] `json:"lan_2"`
	// The name of the ACL.
	Name      param.Field[string]                           `json:"name"`
	Protocols param.Field[[]SiteACLUpdateParamsACLProtocol] `json:"protocols"`
}

func (r SiteACLUpdateParamsACL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLUpdateParamsACLLan1 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLUpdateParamsACLLan1Subnet] `json:"subnets"`
}

func (r SiteACLUpdateParamsACLLan1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLUpdateParamsACLLan1Subnet interface {
	ImplementsMagicTransitSiteACLUpdateParamsACLLan1Subnet()
}

type SiteACLUpdateParamsACLLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName param.Field[string] `json:"lan_name"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SiteACLUpdateParamsACLLan2Subnet] `json:"subnets"`
}

func (r SiteACLUpdateParamsACLLan2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid IPv4 address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type SiteACLUpdateParamsACLLan2Subnet interface {
	ImplementsMagicTransitSiteACLUpdateParamsACLLan2Subnet()
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type SiteACLUpdateParamsACLProtocol string

const (
	SiteACLUpdateParamsACLProtocolTcp  SiteACLUpdateParamsACLProtocol = "tcp"
	SiteACLUpdateParamsACLProtocolUdp  SiteACLUpdateParamsACLProtocol = "udp"
	SiteACLUpdateParamsACLProtocolIcmp SiteACLUpdateParamsACLProtocol = "icmp"
)

func (r SiteACLUpdateParamsACLProtocol) IsKnown() bool {
	switch r {
	case SiteACLUpdateParamsACLProtocolTcp, SiteACLUpdateParamsACLProtocolUdp, SiteACLUpdateParamsACLProtocolIcmp:
		return true
	}
	return false
}

type SiteACLUpdateResponseEnvelope struct {
	Errors   []SiteACLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteACLUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteACLUpdateResponse                   `json:"result,required"`
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

type SiteACLUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteACLUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// siteACLUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteACLUpdateResponseEnvelopeErrors]
type siteACLUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteACLUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteACLUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// siteACLUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteACLUpdateResponseEnvelopeMessages]
type siteACLUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []SiteACLListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteACLListResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteACLListResponse                   `json:"result,required"`
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

type SiteACLListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    siteACLListResponseEnvelopeErrorsJSON `json:"-"`
}

// siteACLListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteACLListResponseEnvelopeErrors]
type siteACLListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteACLListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteACLListResponseEnvelopeMessagesJSON `json:"-"`
}

// siteACLListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteACLListResponseEnvelopeMessages]
type siteACLListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLListResponseEnvelopeMessagesJSON) RawJSON() string {
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLDeleteResponseEnvelope struct {
	Errors   []SiteACLDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteACLDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteACLDeleteResponse                   `json:"result,required"`
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

type SiteACLDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteACLDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// siteACLDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseEnvelopeErrors]
type siteACLDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteACLDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteACLDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// siteACLDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteACLDeleteResponseEnvelopeMessages]
type siteACLDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []SiteACLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteACLGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteACLGetResponse                   `json:"result,required"`
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

type SiteACLGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteACLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// siteACLGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteACLGetResponseEnvelopeErrors]
type siteACLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteACLGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteACLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// siteACLGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteACLGetResponseEnvelopeMessages]
type siteACLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
