// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// LockdownService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLockdownService] method instead.
type LockdownService struct {
	Options []option.RequestOption
}

// NewLockdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockdownService(opts ...option.RequestOption) (r *LockdownService) {
	r = &LockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *LockdownService) New(ctx context.Context, zoneIdentifier string, body LockdownNewParams, opts ...option.RequestOption) (res *LockdownNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *LockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body LockdownUpdateParams, opts ...option.RequestOption) (res *LockdownUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) List(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[LockdownListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
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

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[LockdownListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *LockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, body LockdownDeleteParams, opts ...option.RequestOption) (res *LockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *LockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LockdownGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LockdownNewResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations LockdownNewResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                `json:"urls,required"`
	JSON lockdownNewResponseJSON `json:"-"`
}

// lockdownNewResponseJSON contains the JSON metadata for the struct
// [LockdownNewResponse]
type lockdownNewResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LockdownNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type LockdownNewResponseConfigurations struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownNewResponseConfigurationsTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                `json:"value"`
	JSON  lockdownNewResponseConfigurationsJSON `json:"-"`
	union LockdownNewResponseConfigurationsUnion
}

// lockdownNewResponseConfigurationsJSON contains the JSON metadata for the struct
// [LockdownNewResponseConfigurations]
type lockdownNewResponseConfigurationsJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lockdownNewResponseConfigurationsJSON) RawJSON() string {
	return r.raw
}

func (r *LockdownNewResponseConfigurations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r LockdownNewResponseConfigurations) AsUnion() LockdownNewResponseConfigurationsUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration] or
// [firewall.LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration].
type LockdownNewResponseConfigurationsUnion interface {
	implementsFirewallLockdownNewResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LockdownNewResponseConfigurationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration{}),
		},
	)
}

type LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                              `json:"value"`
	JSON  lockdownNewResponseConfigurationsFirewallSchemasIPConfigurationJSON `json:"-"`
}

// lockdownNewResponseConfigurationsFirewallSchemasIPConfigurationJSON contains the
// JSON metadata for the struct
// [LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration]
type lockdownNewResponseConfigurationsFirewallSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseConfigurationsFirewallSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownNewResponseConfigurationsFirewallSchemasIPConfiguration) implementsFirewallLockdownNewResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTarget string

const (
	LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTargetIP LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTarget = "ip"
)

func (r LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownNewResponseConfigurationsFirewallSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                `json:"value"`
	JSON  lockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationJSON `json:"-"`
}

// lockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationJSON contains
// the JSON metadata for the struct
// [LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration]
type lockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownNewResponseConfigurationsFirewallSchemasCIDRConfiguration) implementsFirewallLockdownNewResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTarget string

const (
	LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTarget = "ip_range"
)

func (r LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownNewResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownNewResponseConfigurationsTarget string

const (
	LockdownNewResponseConfigurationsTargetIP      LockdownNewResponseConfigurationsTarget = "ip"
	LockdownNewResponseConfigurationsTargetIPRange LockdownNewResponseConfigurationsTarget = "ip_range"
)

func (r LockdownNewResponseConfigurationsTarget) IsKnown() bool {
	switch r {
	case LockdownNewResponseConfigurationsTargetIP, LockdownNewResponseConfigurationsTargetIPRange:
		return true
	}
	return false
}

type LockdownUpdateResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations LockdownUpdateResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                   `json:"urls,required"`
	JSON lockdownUpdateResponseJSON `json:"-"`
}

// lockdownUpdateResponseJSON contains the JSON metadata for the struct
// [LockdownUpdateResponse]
type lockdownUpdateResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LockdownUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type LockdownUpdateResponseConfigurations struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownUpdateResponseConfigurationsTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                   `json:"value"`
	JSON  lockdownUpdateResponseConfigurationsJSON `json:"-"`
	union LockdownUpdateResponseConfigurationsUnion
}

// lockdownUpdateResponseConfigurationsJSON contains the JSON metadata for the
// struct [LockdownUpdateResponseConfigurations]
type lockdownUpdateResponseConfigurationsJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lockdownUpdateResponseConfigurationsJSON) RawJSON() string {
	return r.raw
}

func (r *LockdownUpdateResponseConfigurations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r LockdownUpdateResponseConfigurations) AsUnion() LockdownUpdateResponseConfigurationsUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration] or
// [firewall.LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration].
type LockdownUpdateResponseConfigurationsUnion interface {
	implementsFirewallLockdownUpdateResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LockdownUpdateResponseConfigurationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration{}),
		},
	)
}

type LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                 `json:"value"`
	JSON  lockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationJSON `json:"-"`
}

// lockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationJSON contains
// the JSON metadata for the struct
// [LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration]
type lockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownUpdateResponseConfigurationsFirewallSchemasIPConfiguration) implementsFirewallLockdownUpdateResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTarget string

const (
	LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTargetIP LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTarget = "ip"
)

func (r LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseConfigurationsFirewallSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                   `json:"value"`
	JSON  lockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationJSON `json:"-"`
}

// lockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationJSON
// contains the JSON metadata for the struct
// [LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration]
type lockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfiguration) implementsFirewallLockdownUpdateResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTarget string

const (
	LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTarget = "ip_range"
)

func (r LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownUpdateResponseConfigurationsTarget string

const (
	LockdownUpdateResponseConfigurationsTargetIP      LockdownUpdateResponseConfigurationsTarget = "ip"
	LockdownUpdateResponseConfigurationsTargetIPRange LockdownUpdateResponseConfigurationsTarget = "ip_range"
)

func (r LockdownUpdateResponseConfigurationsTarget) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseConfigurationsTargetIP, LockdownUpdateResponseConfigurationsTargetIPRange:
		return true
	}
	return false
}

type LockdownListResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations LockdownListResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                 `json:"urls,required"`
	JSON lockdownListResponseJSON `json:"-"`
}

// lockdownListResponseJSON contains the JSON metadata for the struct
// [LockdownListResponse]
type lockdownListResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LockdownListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownListResponseJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type LockdownListResponseConfigurations struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownListResponseConfigurationsTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                 `json:"value"`
	JSON  lockdownListResponseConfigurationsJSON `json:"-"`
	union LockdownListResponseConfigurationsUnion
}

// lockdownListResponseConfigurationsJSON contains the JSON metadata for the struct
// [LockdownListResponseConfigurations]
type lockdownListResponseConfigurationsJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lockdownListResponseConfigurationsJSON) RawJSON() string {
	return r.raw
}

func (r *LockdownListResponseConfigurations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r LockdownListResponseConfigurations) AsUnion() LockdownListResponseConfigurationsUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.LockdownListResponseConfigurationsFirewallSchemasIPConfiguration] or
// [firewall.LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration].
type LockdownListResponseConfigurationsUnion interface {
	implementsFirewallLockdownListResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LockdownListResponseConfigurationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownListResponseConfigurationsFirewallSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration{}),
		},
	)
}

type LockdownListResponseConfigurationsFirewallSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                               `json:"value"`
	JSON  lockdownListResponseConfigurationsFirewallSchemasIPConfigurationJSON `json:"-"`
}

// lockdownListResponseConfigurationsFirewallSchemasIPConfigurationJSON contains
// the JSON metadata for the struct
// [LockdownListResponseConfigurationsFirewallSchemasIPConfiguration]
type lockdownListResponseConfigurationsFirewallSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownListResponseConfigurationsFirewallSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownListResponseConfigurationsFirewallSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownListResponseConfigurationsFirewallSchemasIPConfiguration) implementsFirewallLockdownListResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTarget string

const (
	LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTargetIP LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTarget = "ip"
)

func (r LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownListResponseConfigurationsFirewallSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                 `json:"value"`
	JSON  lockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationJSON `json:"-"`
}

// lockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationJSON contains
// the JSON metadata for the struct
// [LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration]
type lockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownListResponseConfigurationsFirewallSchemasCIDRConfiguration) implementsFirewallLockdownListResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTarget string

const (
	LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTarget = "ip_range"
)

func (r LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownListResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownListResponseConfigurationsTarget string

const (
	LockdownListResponseConfigurationsTargetIP      LockdownListResponseConfigurationsTarget = "ip"
	LockdownListResponseConfigurationsTargetIPRange LockdownListResponseConfigurationsTarget = "ip_range"
)

func (r LockdownListResponseConfigurationsTarget) IsKnown() bool {
	switch r {
	case LockdownListResponseConfigurationsTargetIP, LockdownListResponseConfigurationsTargetIPRange:
		return true
	}
	return false
}

type LockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                     `json:"id"`
	JSON lockdownDeleteResponseJSON `json:"-"`
}

// lockdownDeleteResponseJSON contains the JSON metadata for the struct
// [LockdownDeleteResponse]
type lockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LockdownGetResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations LockdownGetResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                `json:"urls,required"`
	JSON lockdownGetResponseJSON `json:"-"`
}

// lockdownGetResponseJSON contains the JSON metadata for the struct
// [LockdownGetResponse]
type lockdownGetResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LockdownGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type LockdownGetResponseConfigurations struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownGetResponseConfigurationsTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                `json:"value"`
	JSON  lockdownGetResponseConfigurationsJSON `json:"-"`
	union LockdownGetResponseConfigurationsUnion
}

// lockdownGetResponseConfigurationsJSON contains the JSON metadata for the struct
// [LockdownGetResponseConfigurations]
type lockdownGetResponseConfigurationsJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lockdownGetResponseConfigurationsJSON) RawJSON() string {
	return r.raw
}

func (r *LockdownGetResponseConfigurations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r LockdownGetResponseConfigurations) AsUnion() LockdownGetResponseConfigurationsUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration] or
// [firewall.LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration].
type LockdownGetResponseConfigurationsUnion interface {
	implementsFirewallLockdownGetResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LockdownGetResponseConfigurationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration{}),
		},
	)
}

type LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                              `json:"value"`
	JSON  lockdownGetResponseConfigurationsFirewallSchemasIPConfigurationJSON `json:"-"`
}

// lockdownGetResponseConfigurationsFirewallSchemasIPConfigurationJSON contains the
// JSON metadata for the struct
// [LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration]
type lockdownGetResponseConfigurationsFirewallSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseConfigurationsFirewallSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownGetResponseConfigurationsFirewallSchemasIPConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTarget string

const (
	LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTargetIP LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTarget = "ip"
)

func (r LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownGetResponseConfigurationsFirewallSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                `json:"value"`
	JSON  lockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationJSON `json:"-"`
}

// lockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationJSON contains
// the JSON metadata for the struct
// [LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration]
type lockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LockdownGetResponseConfigurationsFirewallSchemasCIDRConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTarget string

const (
	LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTarget = "ip_range"
)

func (r LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LockdownGetResponseConfigurationsFirewallSchemasCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LockdownGetResponseConfigurationsTarget string

const (
	LockdownGetResponseConfigurationsTargetIP      LockdownGetResponseConfigurationsTarget = "ip"
	LockdownGetResponseConfigurationsTargetIPRange LockdownGetResponseConfigurationsTarget = "ip_range"
)

func (r LockdownGetResponseConfigurationsTarget) IsKnown() bool {
	switch r {
	case LockdownGetResponseConfigurationsTargetIP, LockdownGetResponseConfigurationsTargetIPRange:
		return true
	}
	return false
}

type LockdownNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LockdownNewResponse                                       `json:"result,required"`
	// Whether the API call was successful
	Success LockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownNewResponseEnvelopeJSON    `json:"-"`
}

// lockdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownNewResponseEnvelope]
type lockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownNewResponseEnvelopeSuccess bool

const (
	LockdownNewResponseEnvelopeSuccessTrue LockdownNewResponseEnvelopeSuccess = true
)

func (r LockdownNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LockdownUpdateResponse                                    `json:"result,required"`
	// Whether the API call was successful
	Success LockdownUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownUpdateResponseEnvelopeJSON    `json:"-"`
}

// lockdownUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownUpdateResponseEnvelope]
type lockdownUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownUpdateResponseEnvelopeSuccess bool

const (
	LockdownUpdateResponseEnvelopeSuccessTrue LockdownUpdateResponseEnvelopeSuccess = true
)

func (r LockdownUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownListParams struct {
	// The timestamp of when the rule was created.
	CreatedOn param.Field[time.Time] `query:"created_on" format:"date-time"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// The timestamp of when the rule was last modified.
	ModifiedOn param.Field[time.Time] `query:"modified_on" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	URISearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [LockdownListParams]'s query parameters as `url.Values`.
func (r LockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LockdownDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownDeleteResponseEnvelope struct {
	Result LockdownDeleteResponse             `json:"result"`
	JSON   lockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// lockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownDeleteResponseEnvelope]
type lockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LockdownGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LockdownGetResponse                                       `json:"result,required"`
	// Whether the API call was successful
	Success LockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownGetResponseEnvelopeJSON    `json:"-"`
}

// lockdownGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownGetResponseEnvelope]
type lockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownGetResponseEnvelopeSuccess bool

const (
	LockdownGetResponseEnvelopeSuccessTrue LockdownGetResponseEnvelopeSuccess = true
)

func (r LockdownGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
