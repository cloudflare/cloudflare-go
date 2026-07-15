// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// DEXCommandService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandService] method instead.
type DEXCommandService struct {
	Options   []option.RequestOption
	Devices   *DEXCommandDeviceService
	Downloads *DEXCommandDownloadService
	Quota     *DEXCommandQuotaService
}

// NewDEXCommandService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXCommandService(opts ...option.RequestOption) (r *DEXCommandService) {
	r = &DEXCommandService{}
	r.Options = opts
	r.Devices = NewDEXCommandDeviceService(opts...)
	r.Downloads = NewDEXCommandDownloadService(opts...)
	r.Quota = NewDEXCommandQuotaService(opts...)
	return
}

// Initiate commands for up to 10 devices per account.
func (r *DEXCommandService) New(ctx context.Context, params DEXCommandNewParams, opts ...option.RequestOption) (res *DEXCommandNewResponse, err error) {
	var env DEXCommandNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a paginated list of commands issued to devices under the specified
// account, optionally filtered by time range, device, or other parameters
func (r *DEXCommandService) List(ctx context.Context, params DEXCommandListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[DEXCommandListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieves a paginated list of commands issued to devices under the specified
// account, optionally filtered by time range, device, or other parameters
func (r *DEXCommandService) ListAutoPaging(ctx context.Context, params DEXCommandListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[DEXCommandListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type DEXCommandNewResponse struct {
	// List of created commands
	Commands []DEXCommandNewResponseCommand `json:"commands"`
	JSON     dexCommandNewResponseJSON      `json:"-"`
}

// dexCommandNewResponseJSON contains the JSON metadata for the struct
// [DEXCommandNewResponse]
type dexCommandNewResponseJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseCommand struct {
	// Unique identifier for the command
	ID string `json:"id"`
	// Command arguments
	Args map[string]string `json:"args"`
	// Identifier for the device associated with the command
	DeviceID string `json:"device_id"`
	// Unique identifier for the device registration
	RegistrationID string `json:"registration_id"`
	// Current status of the command
	Status DEXCommandNewResponseCommandsStatus `json:"status"`
	// Type of the command (e.g., "pcap", "speed-test", or "warp-diag")
	Type string                           `json:"type"`
	JSON dexCommandNewResponseCommandJSON `json:"-"`
}

// dexCommandNewResponseCommandJSON contains the JSON metadata for the struct
// [DEXCommandNewResponseCommand]
type dexCommandNewResponseCommandJSON struct {
	ID             apijson.Field
	Args           apijson.Field
	DeviceID       apijson.Field
	RegistrationID apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DEXCommandNewResponseCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseCommandJSON) RawJSON() string {
	return r.raw
}

// Current status of the command
type DEXCommandNewResponseCommandsStatus string

const (
	DEXCommandNewResponseCommandsStatusPendingExec   DEXCommandNewResponseCommandsStatus = "PENDING_EXEC"
	DEXCommandNewResponseCommandsStatusPendingUpload DEXCommandNewResponseCommandsStatus = "PENDING_UPLOAD"
	DEXCommandNewResponseCommandsStatusSuccess       DEXCommandNewResponseCommandsStatus = "SUCCESS"
	DEXCommandNewResponseCommandsStatusFailed        DEXCommandNewResponseCommandsStatus = "FAILED"
)

func (r DEXCommandNewResponseCommandsStatus) IsKnown() bool {
	switch r {
	case DEXCommandNewResponseCommandsStatusPendingExec, DEXCommandNewResponseCommandsStatusPendingUpload, DEXCommandNewResponseCommandsStatusSuccess, DEXCommandNewResponseCommandsStatusFailed:
		return true
	}
	return false
}

type DEXCommandListResponse struct {
	Commands []DEXCommandListResponseCommand `json:"commands"`
	JSON     dexCommandListResponseJSON      `json:"-"`
}

// dexCommandListResponseJSON contains the JSON metadata for the struct
// [DEXCommandListResponse]
type dexCommandListResponseJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandListResponseCommand struct {
	ID            string    `json:"id"`
	CompletedDate time.Time `json:"completed_date" api:"nullable" format:"date-time"`
	CreatedDate   time.Time `json:"created_date" format:"date-time"`
	DeviceID      string    `json:"device_id"`
	Filename      string    `json:"filename" api:"nullable"`
	// Unique identifier for the device registration
	RegistrationID string                            `json:"registration_id"`
	Status         string                            `json:"status"`
	Type           string                            `json:"type"`
	UserEmail      string                            `json:"user_email"`
	JSON           dexCommandListResponseCommandJSON `json:"-"`
}

// dexCommandListResponseCommandJSON contains the JSON metadata for the struct
// [DEXCommandListResponseCommand]
type dexCommandListResponseCommandJSON struct {
	ID             apijson.Field
	CompletedDate  apijson.Field
	CreatedDate    apijson.Field
	DeviceID       apijson.Field
	Filename       apijson.Field
	RegistrationID apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	UserEmail      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DEXCommandListResponseCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandListResponseCommandJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// List of device-level commands to execute
	Commands param.Field[[]DEXCommandNewParamsCommand] `json:"commands" api:"required"`
}

func (r DEXCommandNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DEXCommandNewParamsCommand struct {
	// Unique identifier for the physical device
	DeviceID param.Field[string] `json:"device_id" api:"required"`
	// Type of command to execute on the device
	Type param.Field[DEXCommandNewParamsCommandsType] `json:"type" api:"required"`
	// Email tied to the device
	UserEmail param.Field[string] `json:"user_email" api:"required"`
	// Command arguments. Allowed fields depend on `type`.
	Args param.Field[DEXCommandNewParamsCommandsArgsUnion] `json:"args"`
	// Unique identifier for the device registration. Required for multi-user devices
	// to target the correct user session.
	RegistrationID param.Field[string] `json:"registration_id"`
}

func (r DEXCommandNewParamsCommand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of command to execute on the device
type DEXCommandNewParamsCommandsType string

const (
	DEXCommandNewParamsCommandsTypePCAP      DEXCommandNewParamsCommandsType = "pcap"
	DEXCommandNewParamsCommandsTypeSpeedTest DEXCommandNewParamsCommandsType = "speed-test"
	DEXCommandNewParamsCommandsTypeWARPDiag  DEXCommandNewParamsCommandsType = "warp-diag"
)

func (r DEXCommandNewParamsCommandsType) IsKnown() bool {
	switch r {
	case DEXCommandNewParamsCommandsTypePCAP, DEXCommandNewParamsCommandsTypeSpeedTest, DEXCommandNewParamsCommandsTypeWARPDiag:
		return true
	}
	return false
}

// Command arguments. Allowed fields depend on `type`.
type DEXCommandNewParamsCommandsArgs struct {
	Interfaces param.Field[interface{}] `json:"interfaces"`
	// Maximum file size (in MB) for the capture file. If the capture artifact exceeds
	// the specified max file size, it will NOT be uploaded.
	MaxFileSizeMB param.Field[float64] `json:"max-file-size-mb"`
	// Maximum number of bytes to save for each packet
	PacketSizeBytes param.Field[float64] `json:"packet-size-bytes"`
	// Test an IP address from all included or excluded ranges. Essentially the same as
	// running 'route get <ip>' and collecting the results. This option may increase
	// the time taken to collect the warp-diag.
	TestAllRoutes param.Field[bool] `json:"test-all-routes"`
	// Limit on capture duration (in minutes)
	TimeLimitMin param.Field[float64] `json:"time-limit-min"`
}

func (r DEXCommandNewParamsCommandsArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DEXCommandNewParamsCommandsArgs) implementsDEXCommandNewParamsCommandsArgsUnion() {}

// Command arguments. Allowed fields depend on `type`.
//
// Satisfied by [zero_trust.DEXCommandNewParamsCommandsArgsWARPDiagArgs],
// [zero_trust.DEXCommandNewParamsCommandsArgsPCAPArgs],
// [zero_trust.DEXCommandNewParamsCommandsArgsSpeedTestArgs],
// [DEXCommandNewParamsCommandsArgs].
type DEXCommandNewParamsCommandsArgsUnion interface {
	implementsDEXCommandNewParamsCommandsArgsUnion()
}

type DEXCommandNewParamsCommandsArgsWARPDiagArgs struct {
	// Test an IP address from all included or excluded ranges. Essentially the same as
	// running 'route get <ip>' and collecting the results. This option may increase
	// the time taken to collect the warp-diag.
	TestAllRoutes param.Field[bool] `json:"test-all-routes"`
}

func (r DEXCommandNewParamsCommandsArgsWARPDiagArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DEXCommandNewParamsCommandsArgsWARPDiagArgs) implementsDEXCommandNewParamsCommandsArgsUnion() {
}

type DEXCommandNewParamsCommandsArgsPCAPArgs struct {
	// Maximum file size (in MB) for the capture file. If the capture artifact exceeds
	// the specified max file size, it will NOT be uploaded.
	MaxFileSizeMB param.Field[float64] `json:"max-file-size-mb"`
	// Maximum number of bytes to save for each packet
	PacketSizeBytes param.Field[float64] `json:"packet-size-bytes"`
	// Limit on capture duration (in minutes)
	TimeLimitMin param.Field[float64] `json:"time-limit-min"`
}

func (r DEXCommandNewParamsCommandsArgsPCAPArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DEXCommandNewParamsCommandsArgsPCAPArgs) implementsDEXCommandNewParamsCommandsArgsUnion() {}

type DEXCommandNewParamsCommandsArgsSpeedTestArgs struct {
	// List of interfaces to run the speed test on
	Interfaces param.Field[[]DEXCommandNewParamsCommandsArgsSpeedTestArgsInterface] `json:"interfaces"`
}

func (r DEXCommandNewParamsCommandsArgsSpeedTestArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DEXCommandNewParamsCommandsArgsSpeedTestArgs) implementsDEXCommandNewParamsCommandsArgsUnion() {
}

type DEXCommandNewParamsCommandsArgsSpeedTestArgsInterface string

const (
	DEXCommandNewParamsCommandsArgsSpeedTestArgsInterfaceDefault DEXCommandNewParamsCommandsArgsSpeedTestArgsInterface = "default"
	DEXCommandNewParamsCommandsArgsSpeedTestArgsInterfaceTunnel  DEXCommandNewParamsCommandsArgsSpeedTestArgsInterface = "tunnel"
)

func (r DEXCommandNewParamsCommandsArgsSpeedTestArgsInterface) IsKnown() bool {
	switch r {
	case DEXCommandNewParamsCommandsArgsSpeedTestArgsInterfaceDefault, DEXCommandNewParamsCommandsArgsSpeedTestArgsInterfaceTunnel:
		return true
	}
	return false
}

type DEXCommandNewResponseEnvelope struct {
	Errors   []DEXCommandNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DEXCommandNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success    DEXCommandNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	Result     DEXCommandNewResponse                   `json:"result"`
	ResultInfo DEXCommandNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexCommandNewResponseEnvelopeJSON       `json:"-"`
}

// dexCommandNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXCommandNewResponseEnvelope]
type dexCommandNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DEXCommandNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexCommandNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexCommandNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeErrors]
type dexCommandNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dexCommandNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeErrorsSource]
type dexCommandNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DEXCommandNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexCommandNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexCommandNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeMessages]
type dexCommandNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dexCommandNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DEXCommandNewResponseEnvelopeMessagesSource]
type dexCommandNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXCommandNewResponseEnvelopeSuccess bool

const (
	DEXCommandNewResponseEnvelopeSuccessTrue DEXCommandNewResponseEnvelopeSuccess = true
)

func (r DEXCommandNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXCommandNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXCommandNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// The number of total pages in the entire result set.
	TotalPages float64                                     `json:"total_pages"`
	JSON       dexCommandNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeResultInfo]
type dexCommandNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DEXCommandListParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page" api:"required"`
	// Number of results per page.
	PerPage param.Field[float64] `query:"per_page" api:"required"`
	// Optionally filter executed commands by command type.
	CommandType param.Field[DEXCommandListParamsCommandType] `query:"command_type"`
	// Unique identifier for a device.
	DeviceID param.Field[string] `query:"device_id"`
	// Start time for the query in ISO (RFC3339 - ISO 8601) format.
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Optionally filter executed commands by status.
	Status param.Field[DEXCommandListParamsStatus] `query:"status"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format.
	To param.Field[time.Time] `query:"to" format:"date-time"`
	// Email tied to the device.
	UserEmail param.Field[string] `query:"user_email"`
}

// URLQuery serializes [DEXCommandListParams]'s query parameters as `url.Values`.
func (r DEXCommandListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Optionally filter executed commands by command type.
type DEXCommandListParamsCommandType string

const (
	DEXCommandListParamsCommandTypePCAP      DEXCommandListParamsCommandType = "pcap"
	DEXCommandListParamsCommandTypeSpeedTest DEXCommandListParamsCommandType = "speed-test"
	DEXCommandListParamsCommandTypeWARPDiag  DEXCommandListParamsCommandType = "warp-diag"
)

func (r DEXCommandListParamsCommandType) IsKnown() bool {
	switch r {
	case DEXCommandListParamsCommandTypePCAP, DEXCommandListParamsCommandTypeSpeedTest, DEXCommandListParamsCommandTypeWARPDiag:
		return true
	}
	return false
}

// Optionally filter executed commands by status.
type DEXCommandListParamsStatus string

const (
	DEXCommandListParamsStatusPendingExec   DEXCommandListParamsStatus = "PENDING_EXEC"
	DEXCommandListParamsStatusPendingUpload DEXCommandListParamsStatus = "PENDING_UPLOAD"
	DEXCommandListParamsStatusSuccess       DEXCommandListParamsStatus = "SUCCESS"
	DEXCommandListParamsStatusFailed        DEXCommandListParamsStatus = "FAILED"
)

func (r DEXCommandListParamsStatus) IsKnown() bool {
	switch r {
	case DEXCommandListParamsStatusPendingExec, DEXCommandListParamsStatusPendingUpload, DEXCommandListParamsStatusSuccess, DEXCommandListParamsStatusFailed:
		return true
	}
	return false
}
