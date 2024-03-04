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

// ZeroTrustDevicePostureIntegrationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDevicePostureIntegrationService] method instead.
type ZeroTrustDevicePostureIntegrationService struct {
	Options []option.RequestOption
}

// NewZeroTrustDevicePostureIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDevicePostureIntegrationService(opts ...option.RequestOption) (r *ZeroTrustDevicePostureIntegrationService) {
	r = &ZeroTrustDevicePostureIntegrationService{}
	r.Options = opts
	return
}

// Create a new device posture integration.
func (r *ZeroTrustDevicePostureIntegrationService) New(ctx context.Context, params ZeroTrustDevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureIntegrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureIntegrationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of device posture integrations for an account.
func (r *ZeroTrustDevicePostureIntegrationService) List(ctx context.Context, query ZeroTrustDevicePostureIntegrationListParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePostureIntegrationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureIntegrationListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured device posture integration.
func (r *ZeroTrustDevicePostureIntegrationService) Delete(ctx context.Context, integrationID string, body ZeroTrustDevicePostureIntegrationDeleteParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureIntegrationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", body.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device posture integration.
func (r *ZeroTrustDevicePostureIntegrationService) Edit(ctx context.Context, integrationID string, params ZeroTrustDevicePostureIntegrationEditParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureIntegrationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureIntegrationEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", params.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device posture integration.
func (r *ZeroTrustDevicePostureIntegrationService) Get(ctx context.Context, integrationID string, query ZeroTrustDevicePostureIntegrationGetParams, opts ...option.RequestOption) (res *ZeroTrustDevicePostureIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePostureIntegrationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", query.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePostureIntegrationNewResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config ZeroTrustDevicePostureIntegrationNewResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type ZeroTrustDevicePostureIntegrationNewResponseType `json:"type"`
	JSON zeroTrustDevicePostureIntegrationNewResponseJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationNewResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureIntegrationNewResponse]
type zeroTrustDevicePostureIntegrationNewResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type ZeroTrustDevicePostureIntegrationNewResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                 `json:"client_id,required"`
	JSON     zeroTrustDevicePostureIntegrationNewResponseConfigJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationNewResponseConfigJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationNewResponseConfig]
type zeroTrustDevicePostureIntegrationNewResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationNewResponseType string

const (
	ZeroTrustDevicePostureIntegrationNewResponseTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationNewResponseType = "workspace_one"
	ZeroTrustDevicePostureIntegrationNewResponseTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationNewResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationNewResponseTypeUptycs         ZeroTrustDevicePostureIntegrationNewResponseType = "uptycs"
	ZeroTrustDevicePostureIntegrationNewResponseTypeIntune         ZeroTrustDevicePostureIntegrationNewResponseType = "intune"
	ZeroTrustDevicePostureIntegrationNewResponseTypeKolide         ZeroTrustDevicePostureIntegrationNewResponseType = "kolide"
	ZeroTrustDevicePostureIntegrationNewResponseTypeTanium         ZeroTrustDevicePostureIntegrationNewResponseType = "tanium"
	ZeroTrustDevicePostureIntegrationNewResponseTypeSentineloneS2s ZeroTrustDevicePostureIntegrationNewResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureIntegrationListResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config ZeroTrustDevicePostureIntegrationListResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type ZeroTrustDevicePostureIntegrationListResponseType `json:"type"`
	JSON zeroTrustDevicePostureIntegrationListResponseJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureIntegrationListResponse]
type zeroTrustDevicePostureIntegrationListResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type ZeroTrustDevicePostureIntegrationListResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                  `json:"client_id,required"`
	JSON     zeroTrustDevicePostureIntegrationListResponseConfigJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseConfigJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationListResponseConfig]
type zeroTrustDevicePostureIntegrationListResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationListResponseType string

const (
	ZeroTrustDevicePostureIntegrationListResponseTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationListResponseType = "workspace_one"
	ZeroTrustDevicePostureIntegrationListResponseTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationListResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationListResponseTypeUptycs         ZeroTrustDevicePostureIntegrationListResponseType = "uptycs"
	ZeroTrustDevicePostureIntegrationListResponseTypeIntune         ZeroTrustDevicePostureIntegrationListResponseType = "intune"
	ZeroTrustDevicePostureIntegrationListResponseTypeKolide         ZeroTrustDevicePostureIntegrationListResponseType = "kolide"
	ZeroTrustDevicePostureIntegrationListResponseTypeTanium         ZeroTrustDevicePostureIntegrationListResponseType = "tanium"
	ZeroTrustDevicePostureIntegrationListResponseTypeSentineloneS2s ZeroTrustDevicePostureIntegrationListResponseType = "sentinelone_s2s"
)

// Union satisfied by [ZeroTrustDevicePostureIntegrationDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustDevicePostureIntegrationDeleteResponse interface {
	ImplementsZeroTrustDevicePostureIntegrationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustDevicePostureIntegrationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustDevicePostureIntegrationEditResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config ZeroTrustDevicePostureIntegrationEditResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type ZeroTrustDevicePostureIntegrationEditResponseType `json:"type"`
	JSON zeroTrustDevicePostureIntegrationEditResponseJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationEditResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureIntegrationEditResponse]
type zeroTrustDevicePostureIntegrationEditResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type ZeroTrustDevicePostureIntegrationEditResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                  `json:"client_id,required"`
	JSON     zeroTrustDevicePostureIntegrationEditResponseConfigJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationEditResponseConfigJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationEditResponseConfig]
type zeroTrustDevicePostureIntegrationEditResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationEditResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationEditResponseType string

const (
	ZeroTrustDevicePostureIntegrationEditResponseTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationEditResponseType = "workspace_one"
	ZeroTrustDevicePostureIntegrationEditResponseTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationEditResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationEditResponseTypeUptycs         ZeroTrustDevicePostureIntegrationEditResponseType = "uptycs"
	ZeroTrustDevicePostureIntegrationEditResponseTypeIntune         ZeroTrustDevicePostureIntegrationEditResponseType = "intune"
	ZeroTrustDevicePostureIntegrationEditResponseTypeKolide         ZeroTrustDevicePostureIntegrationEditResponseType = "kolide"
	ZeroTrustDevicePostureIntegrationEditResponseTypeTanium         ZeroTrustDevicePostureIntegrationEditResponseType = "tanium"
	ZeroTrustDevicePostureIntegrationEditResponseTypeSentineloneS2s ZeroTrustDevicePostureIntegrationEditResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureIntegrationGetResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config ZeroTrustDevicePostureIntegrationGetResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type ZeroTrustDevicePostureIntegrationGetResponseType `json:"type"`
	JSON zeroTrustDevicePostureIntegrationGetResponseJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationGetResponseJSON contains the JSON metadata for
// the struct [ZeroTrustDevicePostureIntegrationGetResponse]
type zeroTrustDevicePostureIntegrationGetResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type ZeroTrustDevicePostureIntegrationGetResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                 `json:"client_id,required"`
	JSON     zeroTrustDevicePostureIntegrationGetResponseConfigJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationGetResponseConfigJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationGetResponseConfig]
type zeroTrustDevicePostureIntegrationGetResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationGetResponseType string

const (
	ZeroTrustDevicePostureIntegrationGetResponseTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationGetResponseType = "workspace_one"
	ZeroTrustDevicePostureIntegrationGetResponseTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationGetResponseType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationGetResponseTypeUptycs         ZeroTrustDevicePostureIntegrationGetResponseType = "uptycs"
	ZeroTrustDevicePostureIntegrationGetResponseTypeIntune         ZeroTrustDevicePostureIntegrationGetResponseType = "intune"
	ZeroTrustDevicePostureIntegrationGetResponseTypeKolide         ZeroTrustDevicePostureIntegrationGetResponseType = "kolide"
	ZeroTrustDevicePostureIntegrationGetResponseTypeTanium         ZeroTrustDevicePostureIntegrationGetResponseType = "tanium"
	ZeroTrustDevicePostureIntegrationGetResponseTypeSentineloneS2s ZeroTrustDevicePostureIntegrationGetResponseType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureIntegrationNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object containing third-party integration information.
	Config param.Field[ZeroTrustDevicePostureIntegrationNewParamsConfig] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[ZeroTrustDevicePostureIntegrationNewParamsType] `json:"type,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest],
// [ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type ZeroTrustDevicePostureIntegrationNewParamsConfig interface {
	implementsZeroTrustDevicePostureIntegrationNewParamsConfig()
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest struct {
	// The Tanium API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Tanium client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationNewParamsType string

const (
	ZeroTrustDevicePostureIntegrationNewParamsTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationNewParamsType = "workspace_one"
	ZeroTrustDevicePostureIntegrationNewParamsTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationNewParamsType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationNewParamsTypeUptycs         ZeroTrustDevicePostureIntegrationNewParamsType = "uptycs"
	ZeroTrustDevicePostureIntegrationNewParamsTypeIntune         ZeroTrustDevicePostureIntegrationNewParamsType = "intune"
	ZeroTrustDevicePostureIntegrationNewParamsTypeKolide         ZeroTrustDevicePostureIntegrationNewParamsType = "kolide"
	ZeroTrustDevicePostureIntegrationNewParamsTypeTanium         ZeroTrustDevicePostureIntegrationNewParamsType = "tanium"
	ZeroTrustDevicePostureIntegrationNewParamsTypeSentineloneS2s ZeroTrustDevicePostureIntegrationNewParamsType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureIntegrationNewResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureIntegrationNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureIntegrationNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationNewResponseEnvelope]
type zeroTrustDevicePostureIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationNewResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePostureIntegrationNewResponseEnvelopeErrors]
type zeroTrustDevicePostureIntegrationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationNewResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationNewResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationNewResponseEnvelopeMessages]
type zeroTrustDevicePostureIntegrationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureIntegrationNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureIntegrationNewResponseEnvelopeSuccessTrue ZeroTrustDevicePostureIntegrationNewResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureIntegrationListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureIntegrationListResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureIntegrationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureIntegrationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePostureIntegrationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePostureIntegrationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePostureIntegrationListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationListResponseEnvelope]
type zeroTrustDevicePostureIntegrationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationListResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationListResponseEnvelopeErrors]
type zeroTrustDevicePostureIntegrationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationListResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationListResponseEnvelopeMessages]
type zeroTrustDevicePostureIntegrationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureIntegrationListResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureIntegrationListResponseEnvelopeSuccessTrue ZeroTrustDevicePostureIntegrationListResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       zeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfo]
type zeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureIntegrationDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureIntegrationDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePostureIntegrationDeleteResponseEnvelope]
type zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrors]
type zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessages]
type zeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue ZeroTrustDevicePostureIntegrationDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureIntegrationEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object containing third-party integration information.
	Config param.Field[ZeroTrustDevicePostureIntegrationEditParamsConfig] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[ZeroTrustDevicePostureIntegrationEditParamsType] `json:"type"`
}

func (r ZeroTrustDevicePostureIntegrationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest],
// [ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type ZeroTrustDevicePostureIntegrationEditParamsConfig interface {
	implementsZeroTrustDevicePostureIntegrationEditParamsConfig()
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest struct {
	// The Tanium API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Tanium client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

type ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
}

// The type of device posture integration.
type ZeroTrustDevicePostureIntegrationEditParamsType string

const (
	ZeroTrustDevicePostureIntegrationEditParamsTypeWorkspaceOne   ZeroTrustDevicePostureIntegrationEditParamsType = "workspace_one"
	ZeroTrustDevicePostureIntegrationEditParamsTypeCrowdstrikeS2s ZeroTrustDevicePostureIntegrationEditParamsType = "crowdstrike_s2s"
	ZeroTrustDevicePostureIntegrationEditParamsTypeUptycs         ZeroTrustDevicePostureIntegrationEditParamsType = "uptycs"
	ZeroTrustDevicePostureIntegrationEditParamsTypeIntune         ZeroTrustDevicePostureIntegrationEditParamsType = "intune"
	ZeroTrustDevicePostureIntegrationEditParamsTypeKolide         ZeroTrustDevicePostureIntegrationEditParamsType = "kolide"
	ZeroTrustDevicePostureIntegrationEditParamsTypeTanium         ZeroTrustDevicePostureIntegrationEditParamsType = "tanium"
	ZeroTrustDevicePostureIntegrationEditParamsTypeSentineloneS2s ZeroTrustDevicePostureIntegrationEditParamsType = "sentinelone_s2s"
)

type ZeroTrustDevicePostureIntegrationEditResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureIntegrationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureIntegrationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureIntegrationEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureIntegrationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureIntegrationEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureIntegrationEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationEditResponseEnvelope]
type zeroTrustDevicePostureIntegrationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationEditResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationEditResponseEnvelopeErrors]
type zeroTrustDevicePostureIntegrationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationEditResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationEditResponseEnvelopeMessages]
type zeroTrustDevicePostureIntegrationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureIntegrationEditResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureIntegrationEditResponseEnvelopeSuccessTrue ZeroTrustDevicePostureIntegrationEditResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePostureIntegrationGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePostureIntegrationGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePostureIntegrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePostureIntegrationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDevicePostureIntegrationGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDevicePostureIntegrationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDevicePostureIntegrationGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDevicePostureIntegrationGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePostureIntegrationGetResponseEnvelope]
type zeroTrustDevicePostureIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationGetResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustDevicePostureIntegrationGetResponseEnvelopeErrors]
type zeroTrustDevicePostureIntegrationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePostureIntegrationGetResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustDevicePostureIntegrationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePostureIntegrationGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePostureIntegrationGetResponseEnvelopeMessages]
type zeroTrustDevicePostureIntegrationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePostureIntegrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePostureIntegrationGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePostureIntegrationGetResponseEnvelopeSuccessTrue ZeroTrustDevicePostureIntegrationGetResponseEnvelopeSuccess = true
)
