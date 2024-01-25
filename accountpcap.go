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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPcapService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountPcapService] method
// instead.
type AccountPcapService struct {
	Options    []option.RequestOption
	Ownerships *AccountPcapOwnershipService
	Downloads  *AccountPcapDownloadService
}

// NewAccountPcapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountPcapService(opts ...option.RequestOption) (r *AccountPcapService) {
	r = &AccountPcapService{}
	r.Options = opts
	r.Ownerships = NewAccountPcapOwnershipService(opts...)
	r.Downloads = NewAccountPcapDownloadService(opts...)
	return
}

// Get information for a PCAP request by id.
func (r *AccountPcapService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountPcapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create new PCAP request for account.
func (r *AccountPcapService) MagicPcapCollectionNewPcapRequest(ctx context.Context, accountIdentifier string, body AccountPcapMagicPcapCollectionNewPcapRequestParams, opts ...option.RequestOption) (res *AccountPcapMagicPcapCollectionNewPcapRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all packet capture requests for an account.
func (r *AccountPcapService) MagicPcapCollectionListPacketCaptureRequests(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPcapGetResponse struct {
	Errors   []AccountPcapGetResponseError   `json:"errors"`
	Messages []AccountPcapGetResponseMessage `json:"messages"`
	Result   AccountPcapGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPcapGetResponseSuccess `json:"success"`
	JSON    accountPcapGetResponseJSON    `json:"-"`
}

// accountPcapGetResponseJSON contains the JSON metadata for the struct
// [AccountPcapGetResponse]
type accountPcapGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    accountPcapGetResponseErrorJSON `json:"-"`
}

// accountPcapGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountPcapGetResponseError]
type accountPcapGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountPcapGetResponseMessageJSON `json:"-"`
}

// accountPcapGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountPcapGetResponseMessage]
type accountPcapGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimple] or
// [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFull].
type AccountPcapGetResponseResult interface {
	implementsAccountPcapGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountPcapGetResponseResult)(nil)).Elem(), "")
}

type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleType `json:"type"`
	JSON accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleJSON `json:"-"`
}

// accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleJSON contains the JSON
// metadata for the struct
// [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimple]
type accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleJSON struct {
	ID          apijson.Field
	FilterV1    apijson.Field
	Status      apijson.Field
	Submitted   apijson.Field
	System      apijson.Field
	TimeLimit   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimple) implementsAccountPcapGetResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                             `json:"source_port"`
	JSON       accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON contains the
// JSON metadata for the struct
// [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1]
type accountPcapGetResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusUnknown           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "unknown"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusSuccess           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "success"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusPending           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "pending"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusRunning           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "running"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionPending AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_pending"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionRunning AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_running"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusComplete          AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "complete"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatusFailed            AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleSystem string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleSystemMagicTransit AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleType string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleTypeSimple AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleType = "simple"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleTypeFull   AccountPcapGetResponseResultJF8Dd6b2PcapsResponseSimpleType = "full"
)

type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFull struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullType `json:"type"`
	JSON accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullJSON `json:"-"`
}

// accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullJSON contains the JSON
// metadata for the struct [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFull]
type accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFull) implementsAccountPcapGetResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                           `json:"source_port"`
	JSON       accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON `json:"-"`
}

// accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON contains the
// JSON metadata for the struct
// [AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullFilterV1]
type accountPcapGetResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusUnknown           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "unknown"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusSuccess           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "success"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusPending           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "pending"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusRunning           AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "running"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusConversionPending AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_pending"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusConversionRunning AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_running"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusComplete          AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "complete"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatusFailed            AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullSystem string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullSystemMagicTransit AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullType string

const (
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullTypeSimple AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullType = "simple"
	AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullTypeFull   AccountPcapGetResponseResultJF8Dd6b2PcapsResponseFullType = "full"
)

// Whether the API call was successful
type AccountPcapGetResponseSuccess bool

const (
	AccountPcapGetResponseSuccessTrue AccountPcapGetResponseSuccess = true
)

type AccountPcapMagicPcapCollectionNewPcapRequestResponse struct {
	Errors   []AccountPcapMagicPcapCollectionNewPcapRequestResponseError   `json:"errors"`
	Messages []AccountPcapMagicPcapCollectionNewPcapRequestResponseMessage `json:"messages"`
	Result   AccountPcapMagicPcapCollectionNewPcapRequestResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPcapMagicPcapCollectionNewPcapRequestResponseSuccess `json:"success"`
	JSON    accountPcapMagicPcapCollectionNewPcapRequestResponseJSON    `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseJSON contains the JSON
// metadata for the struct [AccountPcapMagicPcapCollectionNewPcapRequestResponse]
type accountPcapMagicPcapCollectionNewPcapRequestResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapMagicPcapCollectionNewPcapRequestResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountPcapMagicPcapCollectionNewPcapRequestResponseErrorJSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseError]
type accountPcapMagicPcapCollectionNewPcapRequestResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapMagicPcapCollectionNewPcapRequestResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountPcapMagicPcapCollectionNewPcapRequestResponseMessageJSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseMessage]
type accountPcapMagicPcapCollectionNewPcapRequestResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimple]
// or
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFull].
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResult interface {
	implementsAccountPcapMagicPcapCollectionNewPcapRequestResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountPcapMagicPcapCollectionNewPcapRequestResponseResult)(nil)).Elem(), "")
}

type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleType `json:"type"`
	JSON accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleJSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimple]
type accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleJSON struct {
	ID          apijson.Field
	FilterV1    apijson.Field
	Status      apijson.Field
	Submitted   apijson.Field
	System      apijson.Field
	TimeLimit   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimple) implementsAccountPcapMagicPcapCollectionNewPcapRequestResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                                           `json:"source_port"`
	JSON       accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1]
type accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusUnknown           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "unknown"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusSuccess           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "success"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusPending           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "pending"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusRunning           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "running"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionPending AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_pending"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionRunning AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_running"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusComplete          AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "complete"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatusFailed            AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseSimpleType = "full"
)

type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFull struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullType `json:"type"`
	JSON accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullJSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFull]
type accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFull) implementsAccountPcapMagicPcapCollectionNewPcapRequestResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                                         `json:"source_port"`
	JSON       accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON `json:"-"`
}

// accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullFilterV1]
type accountPcapMagicPcapCollectionNewPcapRequestResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusUnknown           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "unknown"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusSuccess           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "success"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusPending           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "pending"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusRunning           AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "running"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusConversionPending AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_pending"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusConversionRunning AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_running"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusComplete          AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "complete"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatusFailed            AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestResponseResultJF8Dd6b2PcapsResponseFullType = "full"
)

// Whether the API call was successful
type AccountPcapMagicPcapCollectionNewPcapRequestResponseSuccess bool

const (
	AccountPcapMagicPcapCollectionNewPcapRequestResponseSuccessTrue AccountPcapMagicPcapCollectionNewPcapRequestResponseSuccess = true
)

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponse struct {
	Errors     []AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseError    `json:"errors"`
	Messages   []AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessage  `json:"messages"`
	Result     []AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult   `json:"result"`
	ResultInfo AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseSuccess `json:"success"`
	JSON    accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseJSON    `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseJSON contains the
// JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponse]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseErrorJSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseError]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessageJSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessage]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimple]
// or
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFull].
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult interface {
	implementsAccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult)(nil)).Elem(), "")
}

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleType `json:"type"`
	JSON accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleJSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimple]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleJSON struct {
	ID          apijson.Field
	FilterV1    apijson.Field
	Status      apijson.Field
	Submitted   apijson.Field
	System      apijson.Field
	TimeLimit   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimple) implementsAccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                                                      `json:"source_port"`
	JSON       accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusUnknown           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "unknown"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusSuccess           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "success"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusPending           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "pending"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusRunning           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "running"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionPending AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_pending"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusConversionRunning AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "conversion_running"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusComplete          AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "complete"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatusFailed            AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleSystem string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleSystemMagicTransit AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleType string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleTypeSimple AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleType = "simple"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleTypeFull   AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseSimpleType = "full"
)

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFull struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullType `json:"type"`
	JSON accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullJSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFull]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFull) implementsAccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResult() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                                                    `json:"source_port"`
	JSON       accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullFilterV1]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJf8Dd6b2PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusUnknown           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "unknown"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusSuccess           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "success"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusPending           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "pending"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusRunning           AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "running"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusConversionPending AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_pending"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusConversionRunning AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "conversion_running"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusComplete          AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "complete"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatusFailed            AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullSystem string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullSystemMagicTransit AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullType string

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullTypeSimple AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullType = "simple"
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullTypeFull   AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultJF8Dd6b2PcapsResponseFullType = "full"
)

type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfoJSON `json:"-"`
}

// accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfo]
type accountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseSuccess bool

const (
	AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseSuccessTrue AccountPcapMagicPcapCollectionListPacketCaptureRequestsResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestSimple],
// [AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestFull].
type AccountPcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams()
}

type AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestSimple) ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress param.Field[string] `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort param.Field[float64] `json:"destination_port"`
	// The protocol number of the packet.
	Protocol param.Field[float64] `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress param.Field[string] `json:"source_address"`
	// The source port of the packet.
	SourcePort param.Field[float64] `json:"source_port"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountPcapMagicPcapCollectionNewPcapRequestParamsJF8Dd6b2PcapsRequestFull) ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress param.Field[string] `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort param.Field[float64] `json:"destination_port"`
	// The protocol number of the packet.
	Protocol param.Field[float64] `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress param.Field[string] `json:"source_address"`
	// The source port of the packet.
	SourcePort param.Field[float64] `json:"source_port"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsJf8Dd6b2PcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
