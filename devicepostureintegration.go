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

// DevicePostureIntegrationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDevicePostureIntegrationService] method instead.
type DevicePostureIntegrationService struct {
	Options []option.RequestOption
}

// NewDevicePostureIntegrationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePostureIntegrationService(opts ...option.RequestOption) (r *DevicePostureIntegrationService) {
	r = &DevicePostureIntegrationService{}
	r.Options = opts
	return
}

// Updates a configured device posture integration.
func (r *DevicePostureIntegrationService) Update(ctx context.Context, identifier interface{}, uuid string, body DevicePostureIntegrationUpdateParams, opts ...option.RequestOption) (res *DevicePostureIntegrationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured device posture integration.
func (r *DevicePostureIntegrationService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new device posture integration.
func (r *DevicePostureIntegrationService) DevicePostureIntegrationsNewDevicePostureIntegration(ctx context.Context, identifier interface{}, body DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams, opts ...option.RequestOption) (res *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) DevicePostureIntegrationsListDevicePostureIntegrations(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device posture integration.
func (r *DevicePostureIntegrationService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DevicePostureIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureIntegrationUpdateResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationUpdateResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationUpdateResponseType `json:"type"`
	JSON devicePostureIntegrationUpdateResponseJSON `json:"-"`
}

// devicePostureIntegrationUpdateResponseJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationUpdateResponse]
type devicePostureIntegrationUpdateResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationUpdateResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                           `json:"client_id,required"`
	JSON     devicePostureIntegrationUpdateResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationUpdateResponseConfigJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationUpdateResponseConfig]
type devicePostureIntegrationUpdateResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationUpdateResponseType string

const (
	DevicePostureIntegrationUpdateResponseTypeWorkspaceOne   DevicePostureIntegrationUpdateResponseType = "workspace_one"
	DevicePostureIntegrationUpdateResponseTypeCrowdstrikeS2s DevicePostureIntegrationUpdateResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationUpdateResponseTypeUptycs         DevicePostureIntegrationUpdateResponseType = "uptycs"
	DevicePostureIntegrationUpdateResponseTypeIntune         DevicePostureIntegrationUpdateResponseType = "intune"
	DevicePostureIntegrationUpdateResponseTypeKolide         DevicePostureIntegrationUpdateResponseType = "kolide"
	DevicePostureIntegrationUpdateResponseTypeTanium         DevicePostureIntegrationUpdateResponseType = "tanium"
	DevicePostureIntegrationUpdateResponseTypeSentineloneS2s DevicePostureIntegrationUpdateResponseType = "sentinelone_s2s"
)

// Union satisfied by [DevicePostureIntegrationDeleteResponseUnknown] or
// [shared.UnionString].
type DevicePostureIntegrationDeleteResponse interface {
	ImplementsDevicePostureIntegrationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DevicePostureIntegrationDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType `json:"type"`
	JSON devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse]
type devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                                                         `json:"client_id,required"`
	JSON     devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfigJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfig]
type devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType string

const (
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeWorkspaceOne   DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "workspace_one"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeCrowdstrikeS2s DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeUptycs         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "uptycs"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeIntune         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "intune"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeKolide         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "kolide"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeTanium         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "tanium"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseTypeSentineloneS2s DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseType = "sentinelone_s2s"
)

type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType `json:"type"`
	JSON devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                                                           `json:"client_id,required"`
	JSON     devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfigJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfig]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType string

const (
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeWorkspaceOne   DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "workspace_one"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeCrowdstrikeS2s DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeUptycs         DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "uptycs"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeIntune         DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "intune"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeKolide         DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "kolide"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeTanium         DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "tanium"
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseTypeSentineloneS2s DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseType = "sentinelone_s2s"
)

type DevicePostureIntegrationGetResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationGetResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationGetResponseType `json:"type"`
	JSON devicePostureIntegrationGetResponseJSON `json:"-"`
}

// devicePostureIntegrationGetResponseJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationGetResponse]
type devicePostureIntegrationGetResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationGetResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                        `json:"client_id,required"`
	JSON     devicePostureIntegrationGetResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationGetResponseConfigJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationGetResponseConfig]
type devicePostureIntegrationGetResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationGetResponseType string

const (
	DevicePostureIntegrationGetResponseTypeWorkspaceOne   DevicePostureIntegrationGetResponseType = "workspace_one"
	DevicePostureIntegrationGetResponseTypeCrowdstrikeS2s DevicePostureIntegrationGetResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationGetResponseTypeUptycs         DevicePostureIntegrationGetResponseType = "uptycs"
	DevicePostureIntegrationGetResponseTypeIntune         DevicePostureIntegrationGetResponseType = "intune"
	DevicePostureIntegrationGetResponseTypeKolide         DevicePostureIntegrationGetResponseType = "kolide"
	DevicePostureIntegrationGetResponseTypeTanium         DevicePostureIntegrationGetResponseType = "tanium"
	DevicePostureIntegrationGetResponseTypeSentineloneS2s DevicePostureIntegrationGetResponseType = "sentinelone_s2s"
)

type DevicePostureIntegrationUpdateParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationUpdateParamsConfig] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationUpdateParamsType] `json:"type"`
}

func (r DevicePostureIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesUptycsConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesIntuneConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesKolideConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesTaniumConfigRequest],
// [DevicePostureIntegrationUpdateParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationUpdateParamsConfig interface {
	implementsDevicePostureIntegrationUpdateParamsConfig()
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesTaniumConfigRequest struct {
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

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

type DevicePostureIntegrationUpdateParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationUpdateParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationUpdateParamsConfig() {
}

// The type of device posture integration.
type DevicePostureIntegrationUpdateParamsType string

const (
	DevicePostureIntegrationUpdateParamsTypeWorkspaceOne   DevicePostureIntegrationUpdateParamsType = "workspace_one"
	DevicePostureIntegrationUpdateParamsTypeCrowdstrikeS2s DevicePostureIntegrationUpdateParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationUpdateParamsTypeUptycs         DevicePostureIntegrationUpdateParamsType = "uptycs"
	DevicePostureIntegrationUpdateParamsTypeIntune         DevicePostureIntegrationUpdateParamsType = "intune"
	DevicePostureIntegrationUpdateParamsTypeKolide         DevicePostureIntegrationUpdateParamsType = "kolide"
	DevicePostureIntegrationUpdateParamsTypeTanium         DevicePostureIntegrationUpdateParamsType = "tanium"
	DevicePostureIntegrationUpdateParamsTypeSentineloneS2s DevicePostureIntegrationUpdateParamsType = "sentinelone_s2s"
)

type DevicePostureIntegrationUpdateResponseEnvelope struct {
	Errors   []DevicePostureIntegrationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePostureIntegrationUpdateResponseEnvelope]
type devicePostureIntegrationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePostureIntegrationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationUpdateResponseEnvelopeErrors]
type devicePostureIntegrationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    devicePostureIntegrationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationUpdateResponseEnvelopeMessages]
type devicePostureIntegrationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationUpdateResponseEnvelopeSuccessTrue DevicePostureIntegrationUpdateResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationDeleteResponseEnvelope struct {
	Errors   []DevicePostureIntegrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePostureIntegrationDeleteResponseEnvelope]
type devicePostureIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePostureIntegrationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationDeleteResponseEnvelopeErrors]
type devicePostureIntegrationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    devicePostureIntegrationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationDeleteResponseEnvelopeMessages]
type devicePostureIntegrationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue DevicePostureIntegrationDeleteResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType] `json:"type,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesUptycsConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesIntuneConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesKolideConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesTaniumConfigRequest],
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig interface {
	implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig()
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesTaniumConfigRequest struct {
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

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

// The type of device posture integration.
type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType string

const (
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeWorkspaceOne   DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "workspace_one"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeCrowdstrikeS2s DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeUptycs         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "uptycs"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeIntune         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "intune"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeKolide         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "kolide"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeTanium         DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "tanium"
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeSentineloneS2s DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "sentinelone_s2s"
)

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelope struct {
	Errors   []DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelope]
type devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrors]
type devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessages]
type devicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeSuccessTrue DevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelope struct {
	Errors   []DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeJSON       `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelope]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrors]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessages]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeSuccessTrue DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                      `json:"total_count"`
	JSON       devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfo]
type devicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationGetResponseEnvelope struct {
	Errors   []DevicePostureIntegrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationGetResponseEnvelope]
type devicePostureIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    devicePostureIntegrationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePostureIntegrationGetResponseEnvelopeErrors]
type devicePostureIntegrationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    devicePostureIntegrationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationGetResponseEnvelopeMessages]
type devicePostureIntegrationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationGetResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationGetResponseEnvelopeSuccessTrue DevicePostureIntegrationGetResponseEnvelopeSuccess = true
)
