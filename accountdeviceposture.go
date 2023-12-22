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

// AccountDevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDevicePostureService]
// method instead.
type AccountDevicePostureService struct {
	Options      []option.RequestOption
	Integrations *AccountDevicePostureIntegrationService
}

// NewAccountDevicePostureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDevicePostureService(opts ...option.RequestOption) (r *AccountDevicePostureService) {
	r = &AccountDevicePostureService{}
	r.Options = opts
	r.Integrations = NewAccountDevicePostureIntegrationService(opts...)
	return
}

// Fetch a single Device Posture Rule.
func (r *AccountDevicePostureService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SingleResponseVjy89cLt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Device Posture Rule.
func (r *AccountDevicePostureService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePostureUpdateParams, opts ...option.RequestOption) (res *SingleResponseVjy89cLt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Device Posture Rule.
func (r *AccountDevicePostureService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *IDResponseCOqJb73Q, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Device Posture Rule.
func (r *AccountDevicePostureService) DevicePostureRulesNewDevicePostureRule(ctx context.Context, identifier interface{}, body AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams, opts ...option.RequestOption) (res *SingleResponseVjy89cLt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Device Posture Rules for an account.
func (r *AccountDevicePostureService) DevicePostureRulesListDevicePostureRules(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionUDcVKm03, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IDResponseCOqJb73Q struct {
	Errors   []IDResponseCOqJb73QError   `json:"errors"`
	Messages []IDResponseCOqJb73QMessage `json:"messages"`
	Result   IDResponseCOqJb73QResult    `json:"result"`
	// Whether the API call was successful
	Success IDResponseCOqJb73QSuccess `json:"success"`
	JSON    idResponseCOqJb73QJSON    `json:"-"`
}

// idResponseCOqJb73QJSON contains the JSON metadata for the struct
// [IDResponseCOqJb73Q]
type idResponseCOqJb73QJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCOqJb73Q) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseCOqJb73QError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    idResponseCOqJb73QErrorJSON `json:"-"`
}

// idResponseCOqJb73QErrorJSON contains the JSON metadata for the struct
// [IDResponseCOqJb73QError]
type idResponseCOqJb73QErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCOqJb73QError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseCOqJb73QMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    idResponseCOqJb73QMessageJSON `json:"-"`
}

// idResponseCOqJb73QMessageJSON contains the JSON metadata for the struct
// [IDResponseCOqJb73QMessage]
type idResponseCOqJb73QMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCOqJb73QMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseCOqJb73QResult struct {
	// API uuid tag.
	ID   string                       `json:"id"`
	JSON idResponseCOqJb73QResultJSON `json:"-"`
}

// idResponseCOqJb73QResultJSON contains the JSON metadata for the struct
// [IDResponseCOqJb73QResult]
type idResponseCOqJb73QResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCOqJb73QResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IDResponseCOqJb73QSuccess bool

const (
	IDResponseCOqJb73QSuccessTrue IDResponseCOqJb73QSuccess = true
)

type ResponseCollectionUDcVKm03 struct {
	Errors     []ResponseCollectionUDcVKm03Error    `json:"errors"`
	Messages   []ResponseCollectionUDcVKm03Message  `json:"messages"`
	Result     []ResponseCollectionUDcVKm03Result   `json:"result"`
	ResultInfo ResponseCollectionUDcVKm03ResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionUDcVKm03Success `json:"success"`
	JSON    responseCollectionUDcVKm03JSON    `json:"-"`
}

// responseCollectionUDcVKm03JSON contains the JSON metadata for the struct
// [ResponseCollectionUDcVKm03]
type responseCollectionUDcVKm03JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUDcVKm03Error struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionUDcVKm03ErrorJSON `json:"-"`
}

// responseCollectionUDcVKm03ErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionUDcVKm03Error]
type responseCollectionUDcVKm03ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUDcVKm03Message struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionUDcVKm03MessageJSON `json:"-"`
}

// responseCollectionUDcVKm03MessageJSON contains the JSON metadata for the struct
// [ResponseCollectionUDcVKm03Message]
type responseCollectionUDcVKm03MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUDcVKm03Result struct {
	// API uuid tag.
	ID string `json:"id"`
	// The description of the Device Posture Rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input ResponseCollectionUDcVKm03ResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []ResponseCollectionUDcVKm03ResultMatch `json:"match"`
	// The name of the Device Posture Rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of Device Posture Rule.
	Type ResponseCollectionUDcVKm03ResultType `json:"type"`
	JSON responseCollectionUDcVKm03ResultJSON `json:"-"`
}

// responseCollectionUDcVKm03ResultJSON contains the JSON metadata for the struct
// [ResponseCollectionUDcVKm03Result]
type responseCollectionUDcVKm03ResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
type ResponseCollectionUDcVKm03ResultInput struct {
	// API uuid tag.
	ID string `json:"id"`
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                      `json:"requireAll"`
	JSON       responseCollectionUDcVKm03ResultInputJSON `json:"-"`
}

// responseCollectionUDcVKm03ResultInputJSON contains the JSON metadata for the
// struct [ResponseCollectionUDcVKm03ResultInput]
type responseCollectionUDcVKm03ResultInputJSON struct {
	ID          apijson.Field
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03ResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUDcVKm03ResultMatch struct {
	Platform ResponseCollectionUDcVKm03ResultMatchPlatform `json:"platform"`
	JSON     responseCollectionUDcVKm03ResultMatchJSON     `json:"-"`
}

// responseCollectionUDcVKm03ResultMatchJSON contains the JSON metadata for the
// struct [ResponseCollectionUDcVKm03ResultMatch]
type responseCollectionUDcVKm03ResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03ResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionUDcVKm03ResultMatchPlatform string

const (
	ResponseCollectionUDcVKm03ResultMatchPlatformWindows ResponseCollectionUDcVKm03ResultMatchPlatform = "windows"
	ResponseCollectionUDcVKm03ResultMatchPlatformMac     ResponseCollectionUDcVKm03ResultMatchPlatform = "mac"
	ResponseCollectionUDcVKm03ResultMatchPlatformLinux   ResponseCollectionUDcVKm03ResultMatchPlatform = "linux"
	ResponseCollectionUDcVKm03ResultMatchPlatformAndroid ResponseCollectionUDcVKm03ResultMatchPlatform = "android"
	ResponseCollectionUDcVKm03ResultMatchPlatformIos     ResponseCollectionUDcVKm03ResultMatchPlatform = "ios"
)

// The type of Device Posture Rule.
type ResponseCollectionUDcVKm03ResultType string

const (
	ResponseCollectionUDcVKm03ResultTypeFile           ResponseCollectionUDcVKm03ResultType = "file"
	ResponseCollectionUDcVKm03ResultTypeApplication    ResponseCollectionUDcVKm03ResultType = "application"
	ResponseCollectionUDcVKm03ResultTypeSerialNumber   ResponseCollectionUDcVKm03ResultType = "serial_number"
	ResponseCollectionUDcVKm03ResultTypeTanium         ResponseCollectionUDcVKm03ResultType = "tanium"
	ResponseCollectionUDcVKm03ResultTypeGateway        ResponseCollectionUDcVKm03ResultType = "gateway"
	ResponseCollectionUDcVKm03ResultTypeWarp           ResponseCollectionUDcVKm03ResultType = "warp"
	ResponseCollectionUDcVKm03ResultTypeDiskEncryption ResponseCollectionUDcVKm03ResultType = "disk_encryption"
)

type ResponseCollectionUDcVKm03ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionUDcVKm03ResultInfoJSON `json:"-"`
}

// responseCollectionUDcVKm03ResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionUDcVKm03ResultInfo]
type responseCollectionUDcVKm03ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionUDcVKm03ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionUDcVKm03Success bool

const (
	ResponseCollectionUDcVKm03SuccessTrue ResponseCollectionUDcVKm03Success = true
)

type SingleResponseVjy89cLt struct {
	Errors   []SingleResponseVjy89cLtError   `json:"errors"`
	Messages []SingleResponseVjy89cLtMessage `json:"messages"`
	Result   SingleResponseVjy89cLtResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseVjy89cLtSuccess `json:"success"`
	JSON    singleResponseVjy89cLtJSON    `json:"-"`
}

// singleResponseVjy89cLtJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLt]
type singleResponseVjy89cLtJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseVjy89cLtError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseVjy89cLtErrorJSON `json:"-"`
}

// singleResponseVjy89cLtErrorJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLtError]
type singleResponseVjy89cLtErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLtError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseVjy89cLtMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseVjy89cLtMessageJSON `json:"-"`
}

// singleResponseVjy89cLtMessageJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLtMessage]
type singleResponseVjy89cLtMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLtMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseVjy89cLtResult struct {
	// API uuid tag.
	ID string `json:"id"`
	// The description of the Device Posture Rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input SingleResponseVjy89cLtResultInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []SingleResponseVjy89cLtResultMatch `json:"match"`
	// The name of the Device Posture Rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of Device Posture Rule.
	Type SingleResponseVjy89cLtResultType `json:"type"`
	JSON singleResponseVjy89cLtResultJSON `json:"-"`
}

// singleResponseVjy89cLtResultJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLtResult]
type singleResponseVjy89cLtResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLtResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The value to be checked against.
type SingleResponseVjy89cLtResultInput struct {
	// API uuid tag.
	ID string `json:"id"`
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                  `json:"requireAll"`
	JSON       singleResponseVjy89cLtResultInputJSON `json:"-"`
}

// singleResponseVjy89cLtResultInputJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLtResultInput]
type singleResponseVjy89cLtResultInputJSON struct {
	ID          apijson.Field
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLtResultInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseVjy89cLtResultMatch struct {
	Platform SingleResponseVjy89cLtResultMatchPlatform `json:"platform"`
	JSON     singleResponseVjy89cLtResultMatchJSON     `json:"-"`
}

// singleResponseVjy89cLtResultMatchJSON contains the JSON metadata for the struct
// [SingleResponseVjy89cLtResultMatch]
type singleResponseVjy89cLtResultMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseVjy89cLtResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseVjy89cLtResultMatchPlatform string

const (
	SingleResponseVjy89cLtResultMatchPlatformWindows SingleResponseVjy89cLtResultMatchPlatform = "windows"
	SingleResponseVjy89cLtResultMatchPlatformMac     SingleResponseVjy89cLtResultMatchPlatform = "mac"
	SingleResponseVjy89cLtResultMatchPlatformLinux   SingleResponseVjy89cLtResultMatchPlatform = "linux"
	SingleResponseVjy89cLtResultMatchPlatformAndroid SingleResponseVjy89cLtResultMatchPlatform = "android"
	SingleResponseVjy89cLtResultMatchPlatformIos     SingleResponseVjy89cLtResultMatchPlatform = "ios"
)

// The type of Device Posture Rule.
type SingleResponseVjy89cLtResultType string

const (
	SingleResponseVjy89cLtResultTypeFile           SingleResponseVjy89cLtResultType = "file"
	SingleResponseVjy89cLtResultTypeApplication    SingleResponseVjy89cLtResultType = "application"
	SingleResponseVjy89cLtResultTypeSerialNumber   SingleResponseVjy89cLtResultType = "serial_number"
	SingleResponseVjy89cLtResultTypeTanium         SingleResponseVjy89cLtResultType = "tanium"
	SingleResponseVjy89cLtResultTypeGateway        SingleResponseVjy89cLtResultType = "gateway"
	SingleResponseVjy89cLtResultTypeWarp           SingleResponseVjy89cLtResultType = "warp"
	SingleResponseVjy89cLtResultTypeDiskEncryption SingleResponseVjy89cLtResultType = "disk_encryption"
)

// Whether the API call was successful
type SingleResponseVjy89cLtSuccess bool

const (
	SingleResponseVjy89cLtSuccessTrue SingleResponseVjy89cLtSuccess = true
)

type AccountDevicePostureUpdateParams struct {
	// The name of the Device Posture Rule.
	Name param.Field[string] `json:"name,required"`
	// The type of Device Posture Rule.
	Type param.Field[AccountDevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the Device Posture Rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[AccountDevicePostureUpdateParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]AccountDevicePostureUpdateParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of Device Posture Rule.
type AccountDevicePostureUpdateParamsType string

const (
	AccountDevicePostureUpdateParamsTypeFile           AccountDevicePostureUpdateParamsType = "file"
	AccountDevicePostureUpdateParamsTypeApplication    AccountDevicePostureUpdateParamsType = "application"
	AccountDevicePostureUpdateParamsTypeSerialNumber   AccountDevicePostureUpdateParamsType = "serial_number"
	AccountDevicePostureUpdateParamsTypeTanium         AccountDevicePostureUpdateParamsType = "tanium"
	AccountDevicePostureUpdateParamsTypeGateway        AccountDevicePostureUpdateParamsType = "gateway"
	AccountDevicePostureUpdateParamsTypeWarp           AccountDevicePostureUpdateParamsType = "warp"
	AccountDevicePostureUpdateParamsTypeDiskEncryption AccountDevicePostureUpdateParamsType = "disk_encryption"
)

// The value to be checked against.
type AccountDevicePostureUpdateParamsInput struct {
	// API uuid tag.
	ID param.Field[string] `json:"id"`
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r AccountDevicePostureUpdateParamsInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureUpdateParamsMatch struct {
	Platform param.Field[AccountDevicePostureUpdateParamsMatchPlatform] `json:"platform"`
}

func (r AccountDevicePostureUpdateParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureUpdateParamsMatchPlatform string

const (
	AccountDevicePostureUpdateParamsMatchPlatformWindows AccountDevicePostureUpdateParamsMatchPlatform = "windows"
	AccountDevicePostureUpdateParamsMatchPlatformMac     AccountDevicePostureUpdateParamsMatchPlatform = "mac"
	AccountDevicePostureUpdateParamsMatchPlatformLinux   AccountDevicePostureUpdateParamsMatchPlatform = "linux"
	AccountDevicePostureUpdateParamsMatchPlatformAndroid AccountDevicePostureUpdateParamsMatchPlatform = "android"
	AccountDevicePostureUpdateParamsMatchPlatformIos     AccountDevicePostureUpdateParamsMatchPlatform = "ios"
)

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams struct {
	// The name of the Device Posture Rule.
	Name param.Field[string] `json:"name,required"`
	// The type of Device Posture Rule.
	Type param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType] `json:"type,required"`
	// The description of the Device Posture Rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of Device Posture Rule.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeFile           AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "file"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeApplication    AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "application"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeSerialNumber   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "serial_number"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeTanium         AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "tanium"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeGateway        AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "gateway"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeWarp           AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "warp"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsTypeDiskEncryption AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsType = "disk_encryption"
)

// The value to be checked against.
type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput struct {
	// API uuid tag.
	ID param.Field[string] `json:"id"`
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch struct {
	Platform param.Field[AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform] `json:"platform"`
}

func (r AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform string

const (
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformWindows AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "windows"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformMac     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "mac"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformLinux   AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "linux"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformAndroid AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "android"
	AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatformIos     AccountDevicePostureDevicePostureRulesNewDevicePostureRuleParamsMatchPlatform = "ios"
)
