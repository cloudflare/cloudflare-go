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
func (r *DevicePostureIntegrationService) New(ctx context.Context, identifier interface{}, body DevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *DevicePostureIntegrationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
