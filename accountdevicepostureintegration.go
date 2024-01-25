// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDevicePostureIntegrationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDevicePostureIntegrationService] method instead.
type AccountDevicePostureIntegrationService struct {
	Options []option.RequestOption
}

// NewAccountDevicePostureIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePostureIntegrationService(opts ...option.RequestOption) (r *AccountDevicePostureIntegrationService) {
	r = &AccountDevicePostureIntegrationService{}
	r.Options = opts
	return
}

// Fetches details for a single device posture integration.
func (r *AccountDevicePostureIntegrationService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device posture integration.
func (r *AccountDevicePostureIntegrationService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePostureIntegrationUpdateParams, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a configured device posture integration.
func (r *AccountDevicePostureIntegrationService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new device posture integration.
func (r *AccountDevicePostureIntegrationService) DevicePostureIntegrationsNewDevicePostureIntegration(ctx context.Context, identifier interface{}, body AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the list of device posture integrations for an account.
func (r *AccountDevicePostureIntegrationService) DevicePostureIntegrationsListDevicePostureIntegrations(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDevicePostureIntegrationGetResponse struct {
	Errors   []AccountDevicePostureIntegrationGetResponseError   `json:"errors"`
	Messages []AccountDevicePostureIntegrationGetResponseMessage `json:"messages"`
	Result   AccountDevicePostureIntegrationGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationGetResponseSuccess `json:"success"`
	JSON    accountDevicePostureIntegrationGetResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationGetResponseJSON contains the JSON metadata for
// the struct [AccountDevicePostureIntegrationGetResponse]
type accountDevicePostureIntegrationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationGetResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountDevicePostureIntegrationGetResponseErrorJSON `json:"-"`
}

// accountDevicePostureIntegrationGetResponseErrorJSON contains the JSON metadata
// for the struct [AccountDevicePostureIntegrationGetResponseError]
type accountDevicePostureIntegrationGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationGetResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountDevicePostureIntegrationGetResponseMessageJSON `json:"-"`
}

// accountDevicePostureIntegrationGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountDevicePostureIntegrationGetResponseMessage]
type accountDevicePostureIntegrationGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationGetResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config AccountDevicePostureIntegrationGetResponseResultConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type AccountDevicePostureIntegrationGetResponseResultType `json:"type"`
	JSON accountDevicePostureIntegrationGetResponseResultJSON `json:"-"`
}

// accountDevicePostureIntegrationGetResponseResultJSON contains the JSON metadata
// for the struct [AccountDevicePostureIntegrationGetResponseResult]
type accountDevicePostureIntegrationGetResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type AccountDevicePostureIntegrationGetResponseResultConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                     `json:"client_id,required"`
	JSON     accountDevicePostureIntegrationGetResponseResultConfigJSON `json:"-"`
}

// accountDevicePostureIntegrationGetResponseResultConfigJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationGetResponseResultConfig]
type accountDevicePostureIntegrationGetResponseResultConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationGetResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type AccountDevicePostureIntegrationGetResponseResultType string

const (
	AccountDevicePostureIntegrationGetResponseResultTypeWorkspaceOne   AccountDevicePostureIntegrationGetResponseResultType = "workspace_one"
	AccountDevicePostureIntegrationGetResponseResultTypeCrowdstrikeS2s AccountDevicePostureIntegrationGetResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationGetResponseResultTypeUptycs         AccountDevicePostureIntegrationGetResponseResultType = "uptycs"
	AccountDevicePostureIntegrationGetResponseResultTypeIntune         AccountDevicePostureIntegrationGetResponseResultType = "intune"
	AccountDevicePostureIntegrationGetResponseResultTypeKolide         AccountDevicePostureIntegrationGetResponseResultType = "kolide"
	AccountDevicePostureIntegrationGetResponseResultTypeTanium         AccountDevicePostureIntegrationGetResponseResultType = "tanium"
	AccountDevicePostureIntegrationGetResponseResultTypeSentineloneS2s AccountDevicePostureIntegrationGetResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureIntegrationGetResponseSuccess bool

const (
	AccountDevicePostureIntegrationGetResponseSuccessTrue AccountDevicePostureIntegrationGetResponseSuccess = true
)

type AccountDevicePostureIntegrationUpdateResponse struct {
	Errors   []AccountDevicePostureIntegrationUpdateResponseError   `json:"errors"`
	Messages []AccountDevicePostureIntegrationUpdateResponseMessage `json:"messages"`
	Result   AccountDevicePostureIntegrationUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationUpdateResponseSuccess `json:"success"`
	JSON    accountDevicePostureIntegrationUpdateResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationUpdateResponseJSON contains the JSON metadata for
// the struct [AccountDevicePostureIntegrationUpdateResponse]
type accountDevicePostureIntegrationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationUpdateResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountDevicePostureIntegrationUpdateResponseErrorJSON `json:"-"`
}

// accountDevicePostureIntegrationUpdateResponseErrorJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationUpdateResponseError]
type accountDevicePostureIntegrationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationUpdateResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountDevicePostureIntegrationUpdateResponseMessageJSON `json:"-"`
}

// accountDevicePostureIntegrationUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationUpdateResponseMessage]
type accountDevicePostureIntegrationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationUpdateResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config AccountDevicePostureIntegrationUpdateResponseResultConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type AccountDevicePostureIntegrationUpdateResponseResultType `json:"type"`
	JSON accountDevicePostureIntegrationUpdateResponseResultJSON `json:"-"`
}

// accountDevicePostureIntegrationUpdateResponseResultJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationUpdateResponseResult]
type accountDevicePostureIntegrationUpdateResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type AccountDevicePostureIntegrationUpdateResponseResultConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                        `json:"client_id,required"`
	JSON     accountDevicePostureIntegrationUpdateResponseResultConfigJSON `json:"-"`
}

// accountDevicePostureIntegrationUpdateResponseResultConfigJSON contains the JSON
// metadata for the struct
// [AccountDevicePostureIntegrationUpdateResponseResultConfig]
type accountDevicePostureIntegrationUpdateResponseResultConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationUpdateResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type AccountDevicePostureIntegrationUpdateResponseResultType string

const (
	AccountDevicePostureIntegrationUpdateResponseResultTypeWorkspaceOne   AccountDevicePostureIntegrationUpdateResponseResultType = "workspace_one"
	AccountDevicePostureIntegrationUpdateResponseResultTypeCrowdstrikeS2s AccountDevicePostureIntegrationUpdateResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationUpdateResponseResultTypeUptycs         AccountDevicePostureIntegrationUpdateResponseResultType = "uptycs"
	AccountDevicePostureIntegrationUpdateResponseResultTypeIntune         AccountDevicePostureIntegrationUpdateResponseResultType = "intune"
	AccountDevicePostureIntegrationUpdateResponseResultTypeKolide         AccountDevicePostureIntegrationUpdateResponseResultType = "kolide"
	AccountDevicePostureIntegrationUpdateResponseResultTypeTanium         AccountDevicePostureIntegrationUpdateResponseResultType = "tanium"
	AccountDevicePostureIntegrationUpdateResponseResultTypeSentineloneS2s AccountDevicePostureIntegrationUpdateResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureIntegrationUpdateResponseSuccess bool

const (
	AccountDevicePostureIntegrationUpdateResponseSuccessTrue AccountDevicePostureIntegrationUpdateResponseSuccess = true
)

type AccountDevicePostureIntegrationDeleteResponse struct {
	Errors   []AccountDevicePostureIntegrationDeleteResponseError   `json:"errors"`
	Messages []AccountDevicePostureIntegrationDeleteResponseMessage `json:"messages"`
	Result   interface{}                                            `json:"result,nullable"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationDeleteResponseSuccess `json:"success"`
	JSON    accountDevicePostureIntegrationDeleteResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationDeleteResponseJSON contains the JSON metadata for
// the struct [AccountDevicePostureIntegrationDeleteResponse]
type accountDevicePostureIntegrationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDeleteResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountDevicePostureIntegrationDeleteResponseErrorJSON `json:"-"`
}

// accountDevicePostureIntegrationDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationDeleteResponseError]
type accountDevicePostureIntegrationDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDeleteResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountDevicePostureIntegrationDeleteResponseMessageJSON `json:"-"`
}

// accountDevicePostureIntegrationDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationDeleteResponseMessage]
type accountDevicePostureIntegrationDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePostureIntegrationDeleteResponseSuccess bool

const (
	AccountDevicePostureIntegrationDeleteResponseSuccessTrue AccountDevicePostureIntegrationDeleteResponseSuccess = true
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse struct {
	Errors   []AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseError   `json:"errors"`
	Messages []AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessage `json:"messages"`
	Result   AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseSuccess `json:"success"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse]
type accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseErrorJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseError]
type accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessageJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessage]
type accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType `json:"type"`
	JSON accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResult]
type accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                                                                      `json:"client_id,required"`
	JSON     accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfigJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfigJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfig]
type accountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType string

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeWorkspaceOne   AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "workspace_one"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeCrowdstrikeS2s AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeUptycs         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "uptycs"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeIntune         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "intune"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeKolide         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "kolide"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeTanium         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "tanium"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultTypeSentineloneS2s AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseResultType = "sentinelone_s2s"
)

// Whether the API call was successful.
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseSuccess bool

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseSuccessTrue AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationResponseSuccess = true
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse struct {
	Errors     []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseError    `json:"errors"`
	Messages   []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessage  `json:"messages"`
	Result     []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResult   `json:"result"`
	ResultInfo AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccess `json:"success"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseError struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseErrorJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseError]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessage struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessageJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessage]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResult struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType `json:"type"`
	JSON accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResult]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third-party integration information.
type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                                                                                        `json:"client_id,required"`
	JSON     accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfigJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfigJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfig]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device posture integration.
type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType string

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeWorkspaceOne   AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "workspace_one"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeCrowdstrikeS2s AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeUptycs         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "uptycs"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeIntune         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "intune"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeKolide         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "kolide"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeTanium         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "tanium"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeSentineloneS2s AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "sentinelone_s2s"
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                     `json:"total_count"`
	JSON       accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfoJSON `json:"-"`
}

// accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfo]
type accountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccess bool

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccessTrue AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccess = true
)

type AccountDevicePostureIntegrationUpdateParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[AccountDevicePostureIntegrationUpdateParamsConfig] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[AccountDevicePostureIntegrationUpdateParamsType] `json:"type"`
}

func (r AccountDevicePostureIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhCrowdstrikeConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhUptycsConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhIntuneConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhKolideConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhTaniumConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhSentineloneS2sConfigRequest].
type AccountDevicePostureIntegrationUpdateParamsConfig interface {
	implementsAccountDevicePostureIntegrationUpdateParamsConfig()
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhWorkspaceOneConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhCrowdstrikeConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhUptycsConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhIntuneConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhKolideConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhTaniumConfigRequest struct {
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

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhTaniumConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigZR7Sv6YhSentineloneS2sConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

// The type of device posture integration.
type AccountDevicePostureIntegrationUpdateParamsType string

const (
	AccountDevicePostureIntegrationUpdateParamsTypeWorkspaceOne   AccountDevicePostureIntegrationUpdateParamsType = "workspace_one"
	AccountDevicePostureIntegrationUpdateParamsTypeCrowdstrikeS2s AccountDevicePostureIntegrationUpdateParamsType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationUpdateParamsTypeUptycs         AccountDevicePostureIntegrationUpdateParamsType = "uptycs"
	AccountDevicePostureIntegrationUpdateParamsTypeIntune         AccountDevicePostureIntegrationUpdateParamsType = "intune"
	AccountDevicePostureIntegrationUpdateParamsTypeKolide         AccountDevicePostureIntegrationUpdateParamsType = "kolide"
	AccountDevicePostureIntegrationUpdateParamsTypeTanium         AccountDevicePostureIntegrationUpdateParamsType = "tanium"
	AccountDevicePostureIntegrationUpdateParamsTypeSentineloneS2s AccountDevicePostureIntegrationUpdateParamsType = "sentinelone_s2s"
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType] `json:"type,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhCrowdstrikeConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhUptycsConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhIntuneConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhKolideConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhTaniumConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhSentineloneS2sConfigRequest].
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig interface {
	implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig()
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhWorkspaceOneConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhCrowdstrikeConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhUptycsConfigRequest struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhUptycsConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhIntuneConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhKolideConfigRequest struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhKolideConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhKolideConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhTaniumConfigRequest struct {
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

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhTaniumConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhTaniumConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhSentineloneS2sConfigRequest struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhSentineloneS2sConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigZR7Sv6YhSentineloneS2sConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

// The type of device posture integration.
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType string

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeWorkspaceOne   AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "workspace_one"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeCrowdstrikeS2s AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeUptycs         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "uptycs"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeIntune         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "intune"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeKolide         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "kolide"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeTanium         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "tanium"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeSentineloneS2s AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "sentinelone_s2s"
)
