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

// Fetch a single Device Posture Integration.
func (r *AccountDevicePostureIntegrationService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SchemasSingleResponseAgeTAcZa, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Device Posture Integration.
func (r *AccountDevicePostureIntegrationService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePostureIntegrationUpdateParams, opts ...option.RequestOption) (res *SchemasSingleResponseAgeTAcZa, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a Device Posture Integration.
func (r *AccountDevicePostureIntegrationService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SchemasIDResponseO5wVo4qV, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Device Posture Integration.
func (r *AccountDevicePostureIntegrationService) DevicePostureIntegrationsNewDevicePostureIntegration(ctx context.Context, identifier interface{}, body AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams, opts ...option.RequestOption) (res *SchemasSingleResponseAgeTAcZa, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Device Posture Integrations for an account.
func (r *AccountDevicePostureIntegrationService) DevicePostureIntegrationsListDevicePostureIntegrations(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/integration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasIDResponseO5wVo4qV struct {
	Errors   []SchemasIDResponseO5wVo4qVError   `json:"errors"`
	Messages []SchemasIDResponseO5wVo4qVMessage `json:"messages"`
	Result   interface{}                        `json:"result,nullable"`
	// Whether the API call was successful
	Success SchemasIDResponseO5wVo4qVSuccess `json:"success"`
	JSON    schemasIDResponseO5wVo4qVJSON    `json:"-"`
}

// schemasIDResponseO5wVo4qVJSON contains the JSON metadata for the struct
// [SchemasIDResponseO5wVo4qV]
type schemasIDResponseO5wVo4qVJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseO5wVo4qV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseO5wVo4qVError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemasIDResponseO5wVo4qVErrorJSON `json:"-"`
}

// schemasIDResponseO5wVo4qVErrorJSON contains the JSON metadata for the struct
// [SchemasIDResponseO5wVo4qVError]
type schemasIDResponseO5wVo4qVErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseO5wVo4qVError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseO5wVo4qVMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasIDResponseO5wVo4qVMessageJSON `json:"-"`
}

// schemasIDResponseO5wVo4qVMessageJSON contains the JSON metadata for the struct
// [SchemasIDResponseO5wVo4qVMessage]
type schemasIDResponseO5wVo4qVMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseO5wVo4qVMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasIDResponseO5wVo4qVSuccess bool

const (
	SchemasIDResponseO5wVo4qVSuccessTrue SchemasIDResponseO5wVo4qVSuccess = true
)

type SchemasSingleResponseAgeTAcZa struct {
	Errors   []SchemasSingleResponseAgeTAcZaError   `json:"errors"`
	Messages []SchemasSingleResponseAgeTAcZaMessage `json:"messages"`
	Result   SchemasSingleResponseAgeTAcZaResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasSingleResponseAgeTAcZaSuccess `json:"success"`
	JSON    schemasSingleResponseAgeTAcZaJSON    `json:"-"`
}

// schemasSingleResponseAgeTAcZaJSON contains the JSON metadata for the struct
// [SchemasSingleResponseAgeTAcZa]
type schemasSingleResponseAgeTAcZaJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseAgeTAcZa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseAgeTAcZaError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasSingleResponseAgeTAcZaErrorJSON `json:"-"`
}

// schemasSingleResponseAgeTAcZaErrorJSON contains the JSON metadata for the struct
// [SchemasSingleResponseAgeTAcZaError]
type schemasSingleResponseAgeTAcZaErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseAgeTAcZaError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseAgeTAcZaMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasSingleResponseAgeTAcZaMessageJSON `json:"-"`
}

// schemasSingleResponseAgeTAcZaMessageJSON contains the JSON metadata for the
// struct [SchemasSingleResponseAgeTAcZaMessage]
type schemasSingleResponseAgeTAcZaMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseAgeTAcZaMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseAgeTAcZaResult struct {
	// API uuid tag.
	ID string `json:"id"`
	// The configuration object containing third party integration information.
	Config SchemasSingleResponseAgeTAcZaResultConfig `json:"config"`
	// The interval between each posture check with the third party API. Use "m" for
	// minutes (e.g. "5m") and "h" for hours (e.g. "12h").
	Interval string `json:"interval"`
	// The name of the Device Posture Integration.
	Name string `json:"name"`
	// The type of Device Posture Integration.
	Type SchemasSingleResponseAgeTAcZaResultType `json:"type"`
	JSON schemasSingleResponseAgeTAcZaResultJSON `json:"-"`
}

// schemasSingleResponseAgeTAcZaResultJSON contains the JSON metadata for the
// struct [SchemasSingleResponseAgeTAcZaResult]
type schemasSingleResponseAgeTAcZaResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseAgeTAcZaResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing third party integration information.
type SchemasSingleResponseAgeTAcZaResultConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                                        `json:"client_id,required"`
	JSON     schemasSingleResponseAgeTAcZaResultConfigJSON `json:"-"`
}

// schemasSingleResponseAgeTAcZaResultConfigJSON contains the JSON metadata for the
// struct [SchemasSingleResponseAgeTAcZaResultConfig]
type schemasSingleResponseAgeTAcZaResultConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseAgeTAcZaResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of Device Posture Integration.
type SchemasSingleResponseAgeTAcZaResultType string

const (
	SchemasSingleResponseAgeTAcZaResultTypeWorkspaceOne   SchemasSingleResponseAgeTAcZaResultType = "workspace_one"
	SchemasSingleResponseAgeTAcZaResultTypeCrowdstrikeS2s SchemasSingleResponseAgeTAcZaResultType = "crowdstrike_s2s"
	SchemasSingleResponseAgeTAcZaResultTypeUptycs         SchemasSingleResponseAgeTAcZaResultType = "uptycs"
	SchemasSingleResponseAgeTAcZaResultTypeIntune         SchemasSingleResponseAgeTAcZaResultType = "intune"
)

// Whether the API call was successful
type SchemasSingleResponseAgeTAcZaSuccess bool

const (
	SchemasSingleResponseAgeTAcZaSuccessTrue SchemasSingleResponseAgeTAcZaSuccess = true
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponse struct {
	Errors     []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseError    `json:"errors"`
	Messages   []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseMessage  `json:"messages"`
	Result     []AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResult   `json:"result"`
	ResultInfo AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
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
	// API uuid tag.
	ID string `json:"id"`
	// The configuration object containing third party integration information.
	Config AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultConfig `json:"config"`
	// The interval between each posture check with the third party API. Use "m" for
	// minutes (e.g. "5m") and "h" for hours (e.g. "12h").
	Interval string `json:"interval"`
	// The name of the Device Posture Integration.
	Name string `json:"name"`
	// The type of Device Posture Integration.
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

// The configuration object containing third party integration information.
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

// The type of Device Posture Integration.
type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType string

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeWorkspaceOne   AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "workspace_one"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeCrowdstrikeS2s AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeUptycs         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "uptycs"
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultTypeIntune         AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseResultType = "intune"
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

// Whether the API call was successful
type AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccess bool

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccessTrue AccountDevicePostureIntegrationDevicePostureIntegrationsListDevicePostureIntegrationsResponseSuccess = true
)

type AccountDevicePostureIntegrationUpdateParams struct {
	// The configuration object containing third party integration information.
	Config param.Field[AccountDevicePostureIntegrationUpdateParamsConfig] `json:"config"`
	// The interval between each posture check with the third party API. Use "m" for
	// minutes (e.g. "5m") and "h" for hours (e.g. "12h").
	Interval param.Field[string] `json:"interval"`
	// The name of the Device Posture Integration.
	Name param.Field[string] `json:"name"`
	// The type of Device Posture Integration.
	Type param.Field[AccountDevicePostureIntegrationUpdateParamsType] `json:"type"`
}

func (r AccountDevicePostureIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third party integration information.
//
// Satisfied by
// [AccountDevicePostureIntegrationUpdateParamsConfigWorkspaceOneConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigCrowdstrikeConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigUptycsConfigRequest],
// [AccountDevicePostureIntegrationUpdateParamsConfigIntuneConfigRequest].
type AccountDevicePostureIntegrationUpdateParamsConfig interface {
	implementsAccountDevicePostureIntegrationUpdateParamsConfig()
}

type AccountDevicePostureIntegrationUpdateParamsConfigWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigWorkspaceOneConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigCrowdstrikeConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigUptycsConfigRequest struct {
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigUptycsConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

type AccountDevicePostureIntegrationUpdateParamsConfigIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationUpdateParamsConfigIntuneConfigRequest) implementsAccountDevicePostureIntegrationUpdateParamsConfig() {
}

// The type of Device Posture Integration.
type AccountDevicePostureIntegrationUpdateParamsType string

const (
	AccountDevicePostureIntegrationUpdateParamsTypeWorkspaceOne   AccountDevicePostureIntegrationUpdateParamsType = "workspace_one"
	AccountDevicePostureIntegrationUpdateParamsTypeCrowdstrikeS2s AccountDevicePostureIntegrationUpdateParamsType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationUpdateParamsTypeUptycs         AccountDevicePostureIntegrationUpdateParamsType = "uptycs"
	AccountDevicePostureIntegrationUpdateParamsTypeIntune         AccountDevicePostureIntegrationUpdateParamsType = "intune"
)

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams struct {
	// The configuration object containing third party integration information.
	Config param.Field[AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig] `json:"config,required"`
	// The interval between each posture check with the third party API. Use "m" for
	// minutes (e.g. "5m") and "h" for hours (e.g. "12h").
	Interval param.Field[string] `json:"interval,required"`
	// The name of the Device Posture Integration.
	Name param.Field[string] `json:"name,required"`
	// The type of Device Posture Integration.
	Type param.Field[AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType] `json:"type,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing third party integration information.
//
// Satisfied by
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigWorkspaceOneConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigCrowdstrikeConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigUptycsConfigRequest],
// [AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigIntuneConfigRequest].
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig interface {
	implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig()
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigWorkspaceOneConfigRequest struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigWorkspaceOneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigWorkspaceOneConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigCrowdstrikeConfigRequest struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigCrowdstrikeConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigCrowdstrikeConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigUptycsConfigRequest struct {
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigUptycsConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigUptycsConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigIntuneConfigRequest struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigIntuneConfigRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfigIntuneConfigRequest) implementsAccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsConfig() {
}

// The type of Device Posture Integration.
type AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType string

const (
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeWorkspaceOne   AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "workspace_one"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeCrowdstrikeS2s AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "crowdstrike_s2s"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeUptycs         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "uptycs"
	AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsTypeIntune         AccountDevicePostureIntegrationDevicePostureIntegrationsNewDevicePostureIntegrationParamsType = "intune"
)
