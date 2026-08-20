// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// TunnelWARPConnectorConfigurationService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorConfigurationService] method instead.
type TunnelWARPConnectorConfigurationService struct {
	Options []option.RequestOption
}

// NewTunnelWARPConnectorConfigurationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorConfigurationService(opts ...option.RequestOption) (r *TunnelWARPConnectorConfigurationService) {
	r = &TunnelWARPConnectorConfigurationService{}
	r.Options = opts
	return
}

// Adds or updates the high-availability configuration for a WARP Connector tunnel.
func (r *TunnelWARPConnectorConfigurationService) Update(ctx context.Context, tunnelID string, params TunnelWARPConnectorConfigurationUpdateParams, opts ...option.RequestOption) (res *TunnelWARPConnectorConfigurationUpdateResponse, err error) {
	var env TunnelWARPConnectorConfigurationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/configurations", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets the high-availability configuration for a WARP Connector tunnel.
func (r *TunnelWARPConnectorConfigurationService) Get(ctx context.Context, tunnelID string, query TunnelWARPConnectorConfigurationGetParams, opts ...option.RequestOption) (res *TunnelWARPConnectorConfigurationGetResponse, err error) {
	var env TunnelWARPConnectorConfigurationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/configurations", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TunnelWARPConnectorConfigurationUpdateResponse struct {
	// Monotonically increasing configuration version, incremented on each PUT.
	ConfigurationVersion int64 `json:"configuration_version" api:"required"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// High-availability mode for the WARP Connector tunnel. `none` means HA is enabled
	// but no provider is configured yet (newly created tunnels default to this).
	// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
	// failover. `local` uses virtual IPs (VIPs) on the local interface.
	HaMode TunnelWARPConnectorConfigurationUpdateResponseHaMode `json:"ha_mode" api:"required"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" api:"required" format:"uuid"`
	// Provider-specific configuration. Present for `aws` and `local` modes.
	Config TunnelWARPConnectorConfigurationUpdateResponseConfig `json:"config" api:"nullable"`
	// Timestamp of the last update. Null if never updated.
	UpdatedAt time.Time                                          `json:"updated_at" api:"nullable" format:"date-time"`
	JSON      tunnelWARPConnectorConfigurationUpdateResponseJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseJSON contains the JSON metadata
// for the struct [TunnelWARPConnectorConfigurationUpdateResponse]
type tunnelWARPConnectorConfigurationUpdateResponseJSON struct {
	ConfigurationVersion apijson.Field
	CreatedAt            apijson.Field
	HaMode               apijson.Field
	TunnelID             apijson.Field
	Config               apijson.Field
	UpdatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// TunnelWARPConnectorConfigurationUpdateResponseHaMode high-availability mode for the WARP Connector tunnel. `none` means HA is enabled
// but no provider is configured yet (newly created tunnels default to this).
// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
// failover. `local` uses virtual IPs (VIPs) on the local interface.
type TunnelWARPConnectorConfigurationUpdateResponseHaMode string

const (
	TunnelWARPConnectorConfigurationUpdateResponseHaModeNone     TunnelWARPConnectorConfigurationUpdateResponseHaMode = "none"
	TunnelWARPConnectorConfigurationUpdateResponseHaModeDisabled TunnelWARPConnectorConfigurationUpdateResponseHaMode = "disabled"
	TunnelWARPConnectorConfigurationUpdateResponseHaModeAws      TunnelWARPConnectorConfigurationUpdateResponseHaMode = "aws"
	TunnelWARPConnectorConfigurationUpdateResponseHaModeLocal    TunnelWARPConnectorConfigurationUpdateResponseHaMode = "local"
)

func (r TunnelWARPConnectorConfigurationUpdateResponseHaMode) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConfigurationUpdateResponseHaModeNone, TunnelWARPConnectorConfigurationUpdateResponseHaModeDisabled, TunnelWARPConnectorConfigurationUpdateResponseHaModeAws, TunnelWARPConnectorConfigurationUpdateResponseHaModeLocal:
		return true
	}
	return false
}

// TunnelWARPConnectorConfigurationUpdateResponseConfig provider-specific configuration. Present for `aws` and `local` modes.
type TunnelWARPConnectorConfigurationUpdateResponseConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID string `json:"fnr_id"`
	// This field can have the runtime type of
	// [[]TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVip].
	Vips interface{} `json:"vips"`
	// This field can have the runtime type of
	// [[]TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPrevious].
	VipsPrevious interface{}                                              `json:"vips_previous"`
	JSON         tunnelWARPConnectorConfigurationUpdateResponseConfigJSON `json:"-"`
	union        TunnelWARPConnectorConfigurationUpdateResponseConfigUnion
}

// tunnelWARPConnectorConfigurationUpdateResponseConfigJSON contains the JSON
// metadata for the struct [TunnelWARPConnectorConfigurationUpdateResponseConfig]
type tunnelWARPConnectorConfigurationUpdateResponseConfigJSON struct {
	FnrID        apijson.Field
	Vips         apijson.Field
	VipsPrevious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r tunnelWARPConnectorConfigurationUpdateResponseConfigJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelWARPConnectorConfigurationUpdateResponseConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelWARPConnectorConfigurationUpdateResponseConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig],
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig].
func (r TunnelWARPConnectorConfigurationUpdateResponseConfig) AsUnion() TunnelWARPConnectorConfigurationUpdateResponseConfigUnion {
	return r.union
}

// TunnelWARPConnectorConfigurationUpdateResponseConfigUnion provider-specific configuration. Present for `aws` and `local` modes.
//
// Union satisfied by
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig] or
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig].
type TunnelWARPConnectorConfigurationUpdateResponseConfigUnion interface {
	implementsTunnelWARPConnectorConfigurationUpdateResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelWARPConnectorConfigurationUpdateResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig{}),
		},
	)
}

type TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID string                                                                      `json:"fnr_id" api:"required"`
	JSON  tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfigJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfigJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig]
type tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfigJSON struct {
	FnrID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfigJSON) RawJSON() string {
	return r.raw
}

func (r TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshAwsConfig) implementsTunnelWARPConnectorConfigurationUpdateResponseConfig() {
}

type TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig struct {
	// VIPs to assign on the CloudflareWARP interface.
	Vips []TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVip `json:"vips" api:"required"`
	// VIPs to clean up on demotion or version drift.
	VipsPrevious []TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPrevious `json:"vips_previous"`
	JSON         tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigJSON           `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig]
type tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigJSON struct {
	Vips         apijson.Field
	VipsPrevious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigJSON) RawJSON() string {
	return r.raw
}

func (r TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfig) implementsTunnelWARPConnectorConfigurationUpdateResponseConfig() {
}

type TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVip struct {
	// Virtual IP address (IPv4 or IPv6).
	Address string                                                                           `json:"address" api:"required"`
	JSON    tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVip]
type tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPrevious struct {
	// Virtual IP address (IPv4 or IPv6).
	Address string                                                                                    `json:"address" api:"required"`
	JSON    tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPreviousJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPreviousJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPrevious]
type tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPreviousJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPrevious) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseConfigTunnelMeshLocalConfigVipsPreviousJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponse struct {
	// Monotonically increasing configuration version, incremented on each PUT.
	ConfigurationVersion int64 `json:"configuration_version" api:"required"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// High-availability mode for the WARP Connector tunnel. `none` means HA is enabled
	// but no provider is configured yet (newly created tunnels default to this).
	// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
	// failover. `local` uses virtual IPs (VIPs) on the local interface.
	HaMode TunnelWARPConnectorConfigurationGetResponseHaMode `json:"ha_mode" api:"required"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" api:"required" format:"uuid"`
	// Provider-specific configuration. Present for `aws` and `local` modes.
	Config TunnelWARPConnectorConfigurationGetResponseConfig `json:"config" api:"nullable"`
	// Timestamp of the last update. Null if never updated.
	UpdatedAt time.Time                                       `json:"updated_at" api:"nullable" format:"date-time"`
	JSON      tunnelWARPConnectorConfigurationGetResponseJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseJSON contains the JSON metadata for
// the struct [TunnelWARPConnectorConfigurationGetResponse]
type tunnelWARPConnectorConfigurationGetResponseJSON struct {
	ConfigurationVersion apijson.Field
	CreatedAt            apijson.Field
	HaMode               apijson.Field
	TunnelID             apijson.Field
	Config               apijson.Field
	UpdatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

// TunnelWARPConnectorConfigurationGetResponseHaMode high-availability mode for the WARP Connector tunnel. `none` means HA is enabled
// but no provider is configured yet (newly created tunnels default to this).
// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
// failover. `local` uses virtual IPs (VIPs) on the local interface.
type TunnelWARPConnectorConfigurationGetResponseHaMode string

const (
	TunnelWARPConnectorConfigurationGetResponseHaModeNone     TunnelWARPConnectorConfigurationGetResponseHaMode = "none"
	TunnelWARPConnectorConfigurationGetResponseHaModeDisabled TunnelWARPConnectorConfigurationGetResponseHaMode = "disabled"
	TunnelWARPConnectorConfigurationGetResponseHaModeAws      TunnelWARPConnectorConfigurationGetResponseHaMode = "aws"
	TunnelWARPConnectorConfigurationGetResponseHaModeLocal    TunnelWARPConnectorConfigurationGetResponseHaMode = "local"
)

func (r TunnelWARPConnectorConfigurationGetResponseHaMode) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConfigurationGetResponseHaModeNone, TunnelWARPConnectorConfigurationGetResponseHaModeDisabled, TunnelWARPConnectorConfigurationGetResponseHaModeAws, TunnelWARPConnectorConfigurationGetResponseHaModeLocal:
		return true
	}
	return false
}

// TunnelWARPConnectorConfigurationGetResponseConfig provider-specific configuration. Present for `aws` and `local` modes.
type TunnelWARPConnectorConfigurationGetResponseConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID string `json:"fnr_id"`
	// This field can have the runtime type of
	// [[]TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVip].
	Vips interface{} `json:"vips"`
	// This field can have the runtime type of
	// [[]TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPrevious].
	VipsPrevious interface{}                                           `json:"vips_previous"`
	JSON         tunnelWARPConnectorConfigurationGetResponseConfigJSON `json:"-"`
	union        TunnelWARPConnectorConfigurationGetResponseConfigUnion
}

// tunnelWARPConnectorConfigurationGetResponseConfigJSON contains the JSON metadata
// for the struct [TunnelWARPConnectorConfigurationGetResponseConfig]
type tunnelWARPConnectorConfigurationGetResponseConfigJSON struct {
	FnrID        apijson.Field
	Vips         apijson.Field
	VipsPrevious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r tunnelWARPConnectorConfigurationGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelWARPConnectorConfigurationGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelWARPConnectorConfigurationGetResponseConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelWARPConnectorConfigurationGetResponseConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig],
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig].
func (r TunnelWARPConnectorConfigurationGetResponseConfig) AsUnion() TunnelWARPConnectorConfigurationGetResponseConfigUnion {
	return r.union
}

// TunnelWARPConnectorConfigurationGetResponseConfigUnion provider-specific configuration. Present for `aws` and `local` modes.
//
// Union satisfied by
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig] or
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig].
type TunnelWARPConnectorConfigurationGetResponseConfigUnion interface {
	implementsTunnelWARPConnectorConfigurationGetResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelWARPConnectorConfigurationGetResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig{}),
		},
	)
}

type TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID string                                                                   `json:"fnr_id" api:"required"`
	JSON  tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfigJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfigJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig]
type tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfigJSON struct {
	FnrID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfigJSON) RawJSON() string {
	return r.raw
}

func (r TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshAwsConfig) implementsTunnelWARPConnectorConfigurationGetResponseConfig() {
}

type TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig struct {
	// VIPs to assign on the CloudflareWARP interface.
	Vips []TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVip `json:"vips" api:"required"`
	// VIPs to clean up on demotion or version drift.
	VipsPrevious []TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPrevious `json:"vips_previous"`
	JSON         tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigJSON           `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig]
type tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigJSON struct {
	Vips         apijson.Field
	VipsPrevious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigJSON) RawJSON() string {
	return r.raw
}

func (r TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfig) implementsTunnelWARPConnectorConfigurationGetResponseConfig() {
}

type TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVip struct {
	// Virtual IP address (IPv4 or IPv6).
	Address string                                                                        `json:"address" api:"required"`
	JSON    tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVip]
type tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPrevious struct {
	// Virtual IP address (IPv4 or IPv6).
	Address string                                                                                 `json:"address" api:"required"`
	JSON    tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPreviousJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPreviousJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPrevious]
type tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPreviousJSON struct {
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPrevious) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseConfigTunnelMeshLocalConfigVipsPreviousJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// High-availability mode for the WARP Connector tunnel. `none` means HA is enabled
	// but no provider is configured yet (newly created tunnels default to this).
	// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
	// failover. `local` uses virtual IPs (VIPs) on the local interface.
	HaMode param.Field[TunnelWARPConnectorConfigurationUpdateParamsHaMode] `json:"ha_mode" api:"required"`
	// Provider-specific configuration. Required shape depends on ha_mode. For `aws`,
	// must contain `fnr_id`. For `local`, must contain `vips`. For `none` and
	// `disabled`, must be empty or omitted.
	Config param.Field[TunnelWARPConnectorConfigurationUpdateParamsConfigUnion] `json:"config"`
}

func (r TunnelWARPConnectorConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TunnelWARPConnectorConfigurationUpdateParamsHaMode high-availability mode for the WARP Connector tunnel. `none` means HA is enabled
// but no provider is configured yet (newly created tunnels default to this).
// `disabled` means HA is explicitly turned off. `aws` uses AWS ENI move for
// failover. `local` uses virtual IPs (VIPs) on the local interface.
type TunnelWARPConnectorConfigurationUpdateParamsHaMode string

const (
	TunnelWARPConnectorConfigurationUpdateParamsHaModeNone     TunnelWARPConnectorConfigurationUpdateParamsHaMode = "none"
	TunnelWARPConnectorConfigurationUpdateParamsHaModeDisabled TunnelWARPConnectorConfigurationUpdateParamsHaMode = "disabled"
	TunnelWARPConnectorConfigurationUpdateParamsHaModeAws      TunnelWARPConnectorConfigurationUpdateParamsHaMode = "aws"
	TunnelWARPConnectorConfigurationUpdateParamsHaModeLocal    TunnelWARPConnectorConfigurationUpdateParamsHaMode = "local"
)

func (r TunnelWARPConnectorConfigurationUpdateParamsHaMode) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConfigurationUpdateParamsHaModeNone, TunnelWARPConnectorConfigurationUpdateParamsHaModeDisabled, TunnelWARPConnectorConfigurationUpdateParamsHaModeAws, TunnelWARPConnectorConfigurationUpdateParamsHaModeLocal:
		return true
	}
	return false
}

// TunnelWARPConnectorConfigurationUpdateParamsConfig provider-specific configuration. Required shape depends on ha_mode. For `aws`,
// must contain `fnr_id`. For `local`, must contain `vips`. For `none` and
// `disabled`, must be empty or omitted.
type TunnelWARPConnectorConfigurationUpdateParamsConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID        param.Field[string]      `json:"fnr_id"`
	Vips         param.Field[interface{}] `json:"vips"`
	VipsPrevious param.Field[interface{}] `json:"vips_previous"`
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfig) implementsTunnelWARPConnectorConfigurationUpdateParamsConfigUnion() {
}

// TunnelWARPConnectorConfigurationUpdateParamsConfigUnion provider-specific configuration. Required shape depends on ha_mode. For `aws`,
// must contain `fnr_id`. For `local`, must contain `vips`. For `none` and
// `disabled`, must be empty or omitted.
//
// Satisfied by
// [zero_trust.TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshAwsConfig],
// [zero_trust.TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfig],
// [TunnelWARPConnectorConfigurationUpdateParamsConfig].
//
// Use [Raw()] to specify an arbitrary value for this param
type TunnelWARPConnectorConfigurationUpdateParamsConfigUnion interface {
	implementsTunnelWARPConnectorConfigurationUpdateParamsConfigUnion()
}

type TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshAwsConfig struct {
	// Floating Network Resource ID — the secondary ENI that is moved between nodes on
	// failover.
	FnrID param.Field[string] `json:"fnr_id" api:"required"`
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshAwsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshAwsConfig) implementsTunnelWARPConnectorConfigurationUpdateParamsConfigUnion() {
}

type TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfig struct {
	// VIPs to assign on the CloudflareWARP interface.
	Vips param.Field[[]TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVip] `json:"vips" api:"required"`
	// VIPs to clean up on demotion or version drift.
	VipsPrevious param.Field[[]TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVipsPrevious] `json:"vips_previous"`
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfig) implementsTunnelWARPConnectorConfigurationUpdateParamsConfigUnion() {
}

type TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVip struct {
	// Virtual IP address (IPv4 or IPv6).
	Address param.Field[string] `json:"address" api:"required"`
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVipsPrevious struct {
	// Virtual IP address (IPv4 or IPv6).
	Address param.Field[string] `json:"address" api:"required"`
}

func (r TunnelWARPConnectorConfigurationUpdateParamsConfigTunnelMeshLocalConfigVipsPrevious) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelWARPConnectorConfigurationUpdateResponseEnvelope struct {
	Errors   []TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TunnelWARPConnectorConfigurationUpdateResponse                `json:"result"`
	JSON    tunnelWARPConnectorConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [TunnelWARPConnectorConfigurationUpdateResponseEnvelope]
type tunnelWARPConnectorConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                              `json:"code" api:"required"`
	Message          string                                                             `json:"message" api:"required"`
	DocumentationURL string                                                             `json:"documentation_url"`
	Source           TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrors]
type tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                                 `json:"pointer"`
	JSON    tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSource]
type tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                                `json:"code" api:"required"`
	Message          string                                                               `json:"message" api:"required"`
	DocumentationURL string                                                               `json:"documentation_url"`
	Source           TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessages]
type tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                                   `json:"pointer"`
	JSON    tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSource]
type tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccessTrue TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelWARPConnectorConfigurationGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type TunnelWARPConnectorConfigurationGetResponseEnvelope struct {
	Errors   []TunnelWARPConnectorConfigurationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TunnelWARPConnectorConfigurationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TunnelWARPConnectorConfigurationGetResponse                `json:"result"`
	JSON    tunnelWARPConnectorConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [TunnelWARPConnectorConfigurationGetResponseEnvelope]
type tunnelWARPConnectorConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           TunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseEnvelopeErrors]
type tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSource]
type tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           TunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseEnvelopeMessages]
type tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [TunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSource]
type tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorConfigurationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccess indicates whether the API call was successful.
type TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccessTrue TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
