// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *DevicePostureIntegrationService) New(ctx context.Context, params DevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *DevicePostureIntegrations, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) List(ctx context.Context, query DevicePostureIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureIntegrations], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", query.AccountID)
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

// Fetches the list of device posture integrations for an account.
func (r *DevicePostureIntegrationService) ListAutoPaging(ctx context.Context, query DevicePostureIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureIntegrations] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured device posture integration.
func (r *DevicePostureIntegrationService) Delete(ctx context.Context, integrationID string, body DevicePostureIntegrationDeleteParams, opts ...option.RequestOption) (res *DevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", body.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device posture integration.
func (r *DevicePostureIntegrationService) Edit(ctx context.Context, integrationID string, params DevicePostureIntegrationEditParams, opts ...option.RequestOption) (res *DevicePostureIntegrations, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", params.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single device posture integration.
func (r *DevicePostureIntegrationService) Get(ctx context.Context, integrationID string, query DevicePostureIntegrationGetParams, opts ...option.RequestOption) (res *DevicePostureIntegrations, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureIntegrationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", query.AccountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePostureIntegrations struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationsConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DevicePostureIntegrationsType `json:"type"`
	JSON devicePostureIntegrationsJSON `json:"-"`
}

// devicePostureIntegrationsJSON contains the JSON metadata for the struct
// [DevicePostureIntegrations]
type devicePostureIntegrationsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationsJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationsConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                              `json:"client_id,required"`
	JSON     devicePostureIntegrationsConfigJSON `json:"-"`
}

// devicePostureIntegrationsConfigJSON contains the JSON metadata for the struct
// [DevicePostureIntegrationsConfig]
type devicePostureIntegrationsConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device posture integration.
type DevicePostureIntegrationsType string

const (
	DevicePostureIntegrationsTypeWorkspaceOne   DevicePostureIntegrationsType = "workspace_one"
	DevicePostureIntegrationsTypeCrowdstrikeS2s DevicePostureIntegrationsType = "crowdstrike_s2s"
	DevicePostureIntegrationsTypeUptycs         DevicePostureIntegrationsType = "uptycs"
	DevicePostureIntegrationsTypeIntune         DevicePostureIntegrationsType = "intune"
	DevicePostureIntegrationsTypeKolide         DevicePostureIntegrationsType = "kolide"
	DevicePostureIntegrationsTypeTanium         DevicePostureIntegrationsType = "tanium"
	DevicePostureIntegrationsTypeSentineloneS2s DevicePostureIntegrationsType = "sentinelone_s2s"
)

func (r DevicePostureIntegrationsType) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationsTypeWorkspaceOne, DevicePostureIntegrationsTypeCrowdstrikeS2s, DevicePostureIntegrationsTypeUptycs, DevicePostureIntegrationsTypeIntune, DevicePostureIntegrationsTypeKolide, DevicePostureIntegrationsTypeTanium, DevicePostureIntegrationsTypeSentineloneS2s:
		return true
	}
	return false
}

// Union satisfied by [zero_trust.DevicePostureIntegrationDeleteResponseUnknown] or
// [shared.UnionString].
type DevicePostureIntegrationDeleteResponse interface {
	ImplementsZeroTrustDevicePostureIntegrationDeleteResponse()
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

type DevicePostureIntegrationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest],
// [zero_trust.DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationNewParamsConfig interface {
	implementsZeroTrustDevicePostureIntegrationNewParamsConfig()
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesUptycsConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesIntuneConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesKolideConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesTaniumConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsZeroTrustDevicePostureIntegrationNewParamsConfig() {
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

func (r DevicePostureIntegrationNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationNewParamsTypeWorkspaceOne, DevicePostureIntegrationNewParamsTypeCrowdstrikeS2s, DevicePostureIntegrationNewParamsTypeUptycs, DevicePostureIntegrationNewParamsTypeIntune, DevicePostureIntegrationNewParamsTypeKolide, DevicePostureIntegrationNewParamsTypeTanium, DevicePostureIntegrationNewParamsTypeSentineloneS2s:
		return true
	}
	return false
}

type DevicePostureIntegrationNewResponseEnvelope struct {
	Errors   []DevicePostureIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrations                             `json:"result,required,nullable"`
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

func (r devicePostureIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationNewResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationNewResponseEnvelopeSuccessTrue DevicePostureIntegrationNewResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r devicePostureIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue DevicePostureIntegrationDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest],
// [zero_trust.DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest].
type DevicePostureIntegrationEditParamsConfig interface {
	implementsZeroTrustDevicePostureIntegrationEditParamsConfig()
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesWorkspaceOneConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesCrowdstrikeConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesUptycsConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesIntuneConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesKolideConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesTaniumConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsConfigTeamsDevicesSentineloneS2sConfigRequest) implementsZeroTrustDevicePostureIntegrationEditParamsConfig() {
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

func (r DevicePostureIntegrationEditParamsType) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationEditParamsTypeWorkspaceOne, DevicePostureIntegrationEditParamsTypeCrowdstrikeS2s, DevicePostureIntegrationEditParamsTypeUptycs, DevicePostureIntegrationEditParamsTypeIntune, DevicePostureIntegrationEditParamsTypeKolide, DevicePostureIntegrationEditParamsTypeTanium, DevicePostureIntegrationEditParamsTypeSentineloneS2s:
		return true
	}
	return false
}

type DevicePostureIntegrationEditResponseEnvelope struct {
	Errors   []DevicePostureIntegrationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrations                              `json:"result,required,nullable"`
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

func (r devicePostureIntegrationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationEditResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationEditResponseEnvelopeSuccessTrue DevicePostureIntegrationEditResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureIntegrationGetResponseEnvelope struct {
	Errors   []DevicePostureIntegrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePostureIntegrationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DevicePostureIntegrations                             `json:"result,required,nullable"`
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

func (r devicePostureIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r devicePostureIntegrationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureIntegrationGetResponseEnvelopeSuccess bool

const (
	DevicePostureIntegrationGetResponseEnvelopeSuccessTrue DevicePostureIntegrationGetResponseEnvelopeSuccess = true
)

func (r DevicePostureIntegrationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureIntegrationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
