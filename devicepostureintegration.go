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

// Create a new device posture integration.
func (r *DevicePostureIntegrationService) New(ctx context.Context, accountID interface{}, body DevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *DevicePostureIntegrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]DevicePostureIntegrationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured device posture integration.
func (r *DevicePostureIntegrationService) Delete(ctx context.Context, accountID interface{}, integrationID string, opts ...option.RequestOption) (res *DevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device posture integration.
func (r *DevicePostureIntegrationService) Edit(ctx context.Context, accountID interface{}, integrationID string, body DevicePostureIntegrationEditParams, opts ...option.RequestOption) (res *DevicePostureIntegrationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device posture integration.
func (r *DevicePostureIntegrationService) Get(ctx context.Context, accountID interface{}, integrationID string, opts ...option.RequestOption) (res *DevicePostureIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureIntegrationNewResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationNewResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationNewResponseType `json:"type"`
	JSON devicePostureIntegrationNewResponseJSON `json:"-"`
}

// devicePostureIntegrationNewResponseJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationNewResponse]
type devicePostureIntegrationNewResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationNewResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                        `json:"client_id,required"`
	JSON     devicePostureIntegrationNewResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationNewResponseConfigJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationNewResponseConfig]
type devicePostureIntegrationNewResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationNewResponseType string

const (
	DevicePostureIntegrationNewResponseTypeWorkspaceOne   DevicePostureIntegrationNewResponseType = "workspace_one"
	DevicePostureIntegrationNewResponseTypeCrowdstrikeS2s DevicePostureIntegrationNewResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationNewResponseTypeUptycs         DevicePostureIntegrationNewResponseType = "uptycs"
	DevicePostureIntegrationNewResponseTypeIntune         DevicePostureIntegrationNewResponseType = "intune"
	DevicePostureIntegrationNewResponseTypeKolide         DevicePostureIntegrationNewResponseType = "kolide"
	DevicePostureIntegrationNewResponseTypeTanium         DevicePostureIntegrationNewResponseType = "tanium"
	DevicePostureIntegrationNewResponseTypeSentineloneS2s DevicePostureIntegrationNewResponseType = "sentinelone_s2s"
)

type DevicePostureIntegrationListResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationListResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationListResponseType `json:"type"`
	JSON devicePostureIntegrationListResponseJSON `json:"-"`
}

// devicePostureIntegrationListResponseJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationListResponse]
type devicePostureIntegrationListResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationListResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                         `json:"client_id,required"`
	JSON     devicePostureIntegrationListResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationListResponseConfigJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationListResponseConfig]
type devicePostureIntegrationListResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationListResponseType string

const (
	DevicePostureIntegrationListResponseTypeWorkspaceOne   DevicePostureIntegrationListResponseType = "workspace_one"
	DevicePostureIntegrationListResponseTypeCrowdstrikeS2s DevicePostureIntegrationListResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationListResponseTypeUptycs         DevicePostureIntegrationListResponseType = "uptycs"
	DevicePostureIntegrationListResponseTypeIntune         DevicePostureIntegrationListResponseType = "intune"
	DevicePostureIntegrationListResponseTypeKolide         DevicePostureIntegrationListResponseType = "kolide"
	DevicePostureIntegrationListResponseTypeTanium         DevicePostureIntegrationListResponseType = "tanium"
	DevicePostureIntegrationListResponseTypeSentineloneS2s DevicePostureIntegrationListResponseType = "sentinelone_s2s"
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

type DevicePostureIntegrationEditResponse struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationEditResponseConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationEditResponseType `json:"type"`
	JSON devicePostureIntegrationEditResponseJSON `json:"-"`
}

// devicePostureIntegrationEditResponseJSON contains the JSON metadata for the
// struct [DevicePostureIntegrationEditResponse]
type devicePostureIntegrationEditResponseJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationEditResponseConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                         `json:"client_id,required"`
	JSON     devicePostureIntegrationEditResponseConfigJSON `json:"-"`
}

// devicePostureIntegrationEditResponseConfigJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationEditResponseConfig]
type devicePostureIntegrationEditResponseConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type DevicePostureIntegrationEditResponseType string

const (
	DevicePostureIntegrationEditResponseTypeWorkspaceOne   DevicePostureIntegrationEditResponseType = "workspace_one"
	DevicePostureIntegrationEditResponseTypeCrowdstrikeS2s DevicePostureIntegrationEditResponseType = "crowdstrike_s2s"
	DevicePostureIntegrationEditResponseTypeUptycs         DevicePostureIntegrationEditResponseType = "uptycs"
	DevicePostureIntegrationEditResponseTypeIntune         DevicePostureIntegrationEditResponseType = "intune"
	DevicePostureIntegrationEditResponseTypeKolide         DevicePostureIntegrationEditResponseType = "kolide"
	DevicePostureIntegrationEditResponseTypeTanium         DevicePostureIntegrationEditResponseType = "tanium"
	DevicePostureIntegrationEditResponseTypeSentineloneS2s DevicePostureIntegrationEditResponseType = "sentinelone_s2s"
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

type DevicePostureIntegrationNewParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationNewParamsConfig] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationNewParamsType] `json:"type,required"`
}

func (r DevicePostureIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest],
// [DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationNewParamsConfig interface {
	implementsDevicePostureIntegrationNewParamsConfig()
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest struct {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

type DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationNewParamsConfig() {
}

// The type of device posture integration.
type DevicePostureIntegrationNewParamsType string

const (
	DevicePostureIntegrationNewParamsTypeWorkspaceOne   DevicePostureIntegrationNewParamsType = "workspace_one"
	DevicePostureIntegrationNewParamsTypeCrowdstrikeS2s DevicePostureIntegrationNewParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationNewParamsTypeUptycs         DevicePostureIntegrationNewParamsType = "uptycs"
	DevicePostureIntegrationNewParamsTypeIntune         DevicePostureIntegrationNewParamsType = "intune"
	DevicePostureIntegrationNewParamsTypeKolide         DevicePostureIntegrationNewParamsType = "kolide"
	DevicePostureIntegrationNewParamsTypeTanium         DevicePostureIntegrationNewParamsType = "tanium"
	DevicePostureIntegrationNewParamsTypeSentineloneS2s DevicePostureIntegrationNewParamsType = "sentinelone_s2s"
)

type DevicePostureIntegrationNewResponseEnvelope struct {
	Errors   []DevicePostureIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationNewResponseEnvelope]
type devicePostureIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationNewResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    devicePostureIntegrationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DevicePostureIntegrationNewResponseEnvelopeErrors]
type devicePostureIntegrationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationNewResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    devicePostureIntegrationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationNewResponseEnvelopeMessages]
type devicePostureIntegrationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationNewResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationNewResponseEnvelopeSuccessTrue DevicePostureIntegrationNewResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationListResponseEnvelope struct {
	Errors   []DevicePostureIntegrationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePostureIntegrationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePostureIntegrationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePostureIntegrationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePostureIntegrationListResponseEnvelopeJSON       `json:"-"`
}

// devicePostureIntegrationListResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationListResponseEnvelope]
type devicePostureIntegrationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    devicePostureIntegrationListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationListResponseEnvelopeErrors]
type devicePostureIntegrationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePostureIntegrationListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationListResponseEnvelopeMessages]
type devicePostureIntegrationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationListResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationListResponseEnvelopeSuccessTrue DevicePostureIntegrationListResponseEnvelopeSuccess = true
)

type DevicePostureIntegrationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       devicePostureIntegrationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePostureIntegrationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationListResponseEnvelopeResultInfo]
type devicePostureIntegrationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type DevicePostureIntegrationEditParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[DevicePostureIntegrationEditParamsConfig] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[DevicePostureIntegrationEditParamsType] `json:"type"`
}

func (r DevicePostureIntegrationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest],
// [DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationEditParamsConfig interface {
	implementsDevicePostureIntegrationEditParamsConfig()
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest struct {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

type DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsDevicePostureIntegrationEditParamsConfig() {
}

// The type of device posture integration.
type DevicePostureIntegrationEditParamsType string

const (
	DevicePostureIntegrationEditParamsTypeWorkspaceOne   DevicePostureIntegrationEditParamsType = "workspace_one"
	DevicePostureIntegrationEditParamsTypeCrowdstrikeS2s DevicePostureIntegrationEditParamsType = "crowdstrike_s2s"
	DevicePostureIntegrationEditParamsTypeUptycs         DevicePostureIntegrationEditParamsType = "uptycs"
	DevicePostureIntegrationEditParamsTypeIntune         DevicePostureIntegrationEditParamsType = "intune"
	DevicePostureIntegrationEditParamsTypeKolide         DevicePostureIntegrationEditParamsType = "kolide"
	DevicePostureIntegrationEditParamsTypeTanium         DevicePostureIntegrationEditParamsType = "tanium"
	DevicePostureIntegrationEditParamsTypeSentineloneS2s DevicePostureIntegrationEditParamsType = "sentinelone_s2s"
)

type DevicePostureIntegrationEditResponseEnvelope struct {
	Errors   []DevicePostureIntegrationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrationEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureIntegrationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureIntegrationEditResponseEnvelopeJSON    `json:"-"`
}

// devicePostureIntegrationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePostureIntegrationEditResponseEnvelope]
type devicePostureIntegrationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    devicePostureIntegrationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePostureIntegrationEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationEditResponseEnvelopeErrors]
type devicePostureIntegrationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePostureIntegrationEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePostureIntegrationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePostureIntegrationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePostureIntegrationEditResponseEnvelopeMessages]
type devicePostureIntegrationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePostureIntegrationEditResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationEditResponseEnvelopeSuccessTrue DevicePostureIntegrationEditResponseEnvelopeSuccess = true
)

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
