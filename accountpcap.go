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
func (r *AccountPcapService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *PcapsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create new PCAP request for account.
func (r *AccountPcapService) MagicPcapCollectionNewPcapRequest(ctx context.Context, accountIdentifier string, body AccountPcapMagicPcapCollectionNewPcapRequestParams, opts ...option.RequestOption) (res *PcapsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all packet capture requests for an account.
func (r *AccountPcapService) MagicPcapCollectionListPacketCaptureRequests(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *PcapsCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PcapsCollectionResponse struct {
	Errors     []PcapsCollectionResponseError    `json:"errors"`
	Messages   []PcapsCollectionResponseMessage  `json:"messages"`
	Result     []PcapsCollectionResponseResult   `json:"result"`
	ResultInfo PcapsCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PcapsCollectionResponseSuccess `json:"success"`
	JSON    pcapsCollectionResponseJSON    `json:"-"`
}

// pcapsCollectionResponseJSON contains the JSON metadata for the struct
// [PcapsCollectionResponse]
type pcapsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsCollectionResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    pcapsCollectionResponseErrorJSON `json:"-"`
}

// pcapsCollectionResponseErrorJSON contains the JSON metadata for the struct
// [PcapsCollectionResponseError]
type pcapsCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsCollectionResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    pcapsCollectionResponseMessageJSON `json:"-"`
}

// pcapsCollectionResponseMessageJSON contains the JSON metadata for the struct
// [PcapsCollectionResponseMessage]
type pcapsCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PcapsCollectionResponseResultPcapsResponseSimple] or
// [PcapsCollectionResponseResultPcapsResponseFull].
type PcapsCollectionResponseResult interface {
	implementsPcapsCollectionResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapsCollectionResponseResult)(nil)).Elem(), "")
}

type PcapsCollectionResponseResultPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapsCollectionResponseResultPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapsCollectionResponseResultPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapsCollectionResponseResultPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapsCollectionResponseResultPcapsResponseSimpleType `json:"type"`
	JSON pcapsCollectionResponseResultPcapsResponseSimpleJSON `json:"-"`
}

// pcapsCollectionResponseResultPcapsResponseSimpleJSON contains the JSON metadata
// for the struct [PcapsCollectionResponseResultPcapsResponseSimple]
type pcapsCollectionResponseResultPcapsResponseSimpleJSON struct {
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

func (r *PcapsCollectionResponseResultPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapsCollectionResponseResultPcapsResponseSimple) implementsPcapsCollectionResponseResult() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapsCollectionResponseResultPcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                      `json:"source_port"`
	JSON       pcapsCollectionResponseResultPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapsCollectionResponseResultPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PcapsCollectionResponseResultPcapsResponseSimpleFilterV1]
type pcapsCollectionResponseResultPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapsCollectionResponseResultPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapsCollectionResponseResultPcapsResponseSimpleStatus string

const (
	PcapsCollectionResponseResultPcapsResponseSimpleStatusUnknown           PcapsCollectionResponseResultPcapsResponseSimpleStatus = "unknown"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusSuccess           PcapsCollectionResponseResultPcapsResponseSimpleStatus = "success"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusPending           PcapsCollectionResponseResultPcapsResponseSimpleStatus = "pending"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusRunning           PcapsCollectionResponseResultPcapsResponseSimpleStatus = "running"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusConversionPending PcapsCollectionResponseResultPcapsResponseSimpleStatus = "conversion_pending"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusConversionRunning PcapsCollectionResponseResultPcapsResponseSimpleStatus = "conversion_running"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusComplete          PcapsCollectionResponseResultPcapsResponseSimpleStatus = "complete"
	PcapsCollectionResponseResultPcapsResponseSimpleStatusFailed            PcapsCollectionResponseResultPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapsCollectionResponseResultPcapsResponseSimpleSystem string

const (
	PcapsCollectionResponseResultPcapsResponseSimpleSystemMagicTransit PcapsCollectionResponseResultPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapsCollectionResponseResultPcapsResponseSimpleType string

const (
	PcapsCollectionResponseResultPcapsResponseSimpleTypeSimple PcapsCollectionResponseResultPcapsResponseSimpleType = "simple"
	PcapsCollectionResponseResultPcapsResponseSimpleTypeFull   PcapsCollectionResponseResultPcapsResponseSimpleType = "full"
)

type PcapsCollectionResponseResultPcapsResponseFull struct {
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
	FilterV1 PcapsCollectionResponseResultPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapsCollectionResponseResultPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapsCollectionResponseResultPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapsCollectionResponseResultPcapsResponseFullType `json:"type"`
	JSON pcapsCollectionResponseResultPcapsResponseFullJSON `json:"-"`
}

// pcapsCollectionResponseResultPcapsResponseFullJSON contains the JSON metadata
// for the struct [PcapsCollectionResponseResultPcapsResponseFull]
type pcapsCollectionResponseResultPcapsResponseFullJSON struct {
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

func (r *PcapsCollectionResponseResultPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapsCollectionResponseResultPcapsResponseFull) implementsPcapsCollectionResponseResult() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapsCollectionResponseResultPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                    `json:"source_port"`
	JSON       pcapsCollectionResponseResultPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapsCollectionResponseResultPcapsResponseFullFilterV1JSON contains the JSON
// metadata for the struct [PcapsCollectionResponseResultPcapsResponseFullFilterV1]
type pcapsCollectionResponseResultPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapsCollectionResponseResultPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapsCollectionResponseResultPcapsResponseFullStatus string

const (
	PcapsCollectionResponseResultPcapsResponseFullStatusUnknown           PcapsCollectionResponseResultPcapsResponseFullStatus = "unknown"
	PcapsCollectionResponseResultPcapsResponseFullStatusSuccess           PcapsCollectionResponseResultPcapsResponseFullStatus = "success"
	PcapsCollectionResponseResultPcapsResponseFullStatusPending           PcapsCollectionResponseResultPcapsResponseFullStatus = "pending"
	PcapsCollectionResponseResultPcapsResponseFullStatusRunning           PcapsCollectionResponseResultPcapsResponseFullStatus = "running"
	PcapsCollectionResponseResultPcapsResponseFullStatusConversionPending PcapsCollectionResponseResultPcapsResponseFullStatus = "conversion_pending"
	PcapsCollectionResponseResultPcapsResponseFullStatusConversionRunning PcapsCollectionResponseResultPcapsResponseFullStatus = "conversion_running"
	PcapsCollectionResponseResultPcapsResponseFullStatusComplete          PcapsCollectionResponseResultPcapsResponseFullStatus = "complete"
	PcapsCollectionResponseResultPcapsResponseFullStatusFailed            PcapsCollectionResponseResultPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapsCollectionResponseResultPcapsResponseFullSystem string

const (
	PcapsCollectionResponseResultPcapsResponseFullSystemMagicTransit PcapsCollectionResponseResultPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapsCollectionResponseResultPcapsResponseFullType string

const (
	PcapsCollectionResponseResultPcapsResponseFullTypeSimple PcapsCollectionResponseResultPcapsResponseFullType = "simple"
	PcapsCollectionResponseResultPcapsResponseFullTypeFull   PcapsCollectionResponseResultPcapsResponseFullType = "full"
)

type PcapsCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       pcapsCollectionResponseResultInfoJSON `json:"-"`
}

// pcapsCollectionResponseResultInfoJSON contains the JSON metadata for the struct
// [PcapsCollectionResponseResultInfo]
type pcapsCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapsCollectionResponseSuccess bool

const (
	PcapsCollectionResponseSuccessTrue PcapsCollectionResponseSuccess = true
)

type PcapsSingleResponse struct {
	Errors   []PcapsSingleResponseError   `json:"errors"`
	Messages []PcapsSingleResponseMessage `json:"messages"`
	Result   PcapsSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PcapsSingleResponseSuccess `json:"success"`
	JSON    pcapsSingleResponseJSON    `json:"-"`
}

// pcapsSingleResponseJSON contains the JSON metadata for the struct
// [PcapsSingleResponse]
type pcapsSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsSingleResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    pcapsSingleResponseErrorJSON `json:"-"`
}

// pcapsSingleResponseErrorJSON contains the JSON metadata for the struct
// [PcapsSingleResponseError]
type pcapsSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsSingleResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    pcapsSingleResponseMessageJSON `json:"-"`
}

// pcapsSingleResponseMessageJSON contains the JSON metadata for the struct
// [PcapsSingleResponseMessage]
type pcapsSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PcapsSingleResponseResultPcapsResponseSimple] or
// [PcapsSingleResponseResultPcapsResponseFull].
type PcapsSingleResponseResult interface {
	implementsPcapsSingleResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapsSingleResponseResult)(nil)).Elem(), "")
}

type PcapsSingleResponseResultPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapsSingleResponseResultPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapsSingleResponseResultPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapsSingleResponseResultPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapsSingleResponseResultPcapsResponseSimpleType `json:"type"`
	JSON pcapsSingleResponseResultPcapsResponseSimpleJSON `json:"-"`
}

// pcapsSingleResponseResultPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapsSingleResponseResultPcapsResponseSimple]
type pcapsSingleResponseResultPcapsResponseSimpleJSON struct {
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

func (r *PcapsSingleResponseResultPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapsSingleResponseResultPcapsResponseSimple) implementsPcapsSingleResponseResult() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapsSingleResponseResultPcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                  `json:"source_port"`
	JSON       pcapsSingleResponseResultPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapsSingleResponseResultPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapsSingleResponseResultPcapsResponseSimpleFilterV1]
type pcapsSingleResponseResultPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapsSingleResponseResultPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapsSingleResponseResultPcapsResponseSimpleStatus string

const (
	PcapsSingleResponseResultPcapsResponseSimpleStatusUnknown           PcapsSingleResponseResultPcapsResponseSimpleStatus = "unknown"
	PcapsSingleResponseResultPcapsResponseSimpleStatusSuccess           PcapsSingleResponseResultPcapsResponseSimpleStatus = "success"
	PcapsSingleResponseResultPcapsResponseSimpleStatusPending           PcapsSingleResponseResultPcapsResponseSimpleStatus = "pending"
	PcapsSingleResponseResultPcapsResponseSimpleStatusRunning           PcapsSingleResponseResultPcapsResponseSimpleStatus = "running"
	PcapsSingleResponseResultPcapsResponseSimpleStatusConversionPending PcapsSingleResponseResultPcapsResponseSimpleStatus = "conversion_pending"
	PcapsSingleResponseResultPcapsResponseSimpleStatusConversionRunning PcapsSingleResponseResultPcapsResponseSimpleStatus = "conversion_running"
	PcapsSingleResponseResultPcapsResponseSimpleStatusComplete          PcapsSingleResponseResultPcapsResponseSimpleStatus = "complete"
	PcapsSingleResponseResultPcapsResponseSimpleStatusFailed            PcapsSingleResponseResultPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapsSingleResponseResultPcapsResponseSimpleSystem string

const (
	PcapsSingleResponseResultPcapsResponseSimpleSystemMagicTransit PcapsSingleResponseResultPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapsSingleResponseResultPcapsResponseSimpleType string

const (
	PcapsSingleResponseResultPcapsResponseSimpleTypeSimple PcapsSingleResponseResultPcapsResponseSimpleType = "simple"
	PcapsSingleResponseResultPcapsResponseSimpleTypeFull   PcapsSingleResponseResultPcapsResponseSimpleType = "full"
)

type PcapsSingleResponseResultPcapsResponseFull struct {
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
	FilterV1 PcapsSingleResponseResultPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapsSingleResponseResultPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapsSingleResponseResultPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapsSingleResponseResultPcapsResponseFullType `json:"type"`
	JSON pcapsSingleResponseResultPcapsResponseFullJSON `json:"-"`
}

// pcapsSingleResponseResultPcapsResponseFullJSON contains the JSON metadata for
// the struct [PcapsSingleResponseResultPcapsResponseFull]
type pcapsSingleResponseResultPcapsResponseFullJSON struct {
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

func (r *PcapsSingleResponseResultPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapsSingleResponseResultPcapsResponseFull) implementsPcapsSingleResponseResult() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapsSingleResponseResultPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                `json:"source_port"`
	JSON       pcapsSingleResponseResultPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapsSingleResponseResultPcapsResponseFullFilterV1JSON contains the JSON
// metadata for the struct [PcapsSingleResponseResultPcapsResponseFullFilterV1]
type pcapsSingleResponseResultPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapsSingleResponseResultPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapsSingleResponseResultPcapsResponseFullStatus string

const (
	PcapsSingleResponseResultPcapsResponseFullStatusUnknown           PcapsSingleResponseResultPcapsResponseFullStatus = "unknown"
	PcapsSingleResponseResultPcapsResponseFullStatusSuccess           PcapsSingleResponseResultPcapsResponseFullStatus = "success"
	PcapsSingleResponseResultPcapsResponseFullStatusPending           PcapsSingleResponseResultPcapsResponseFullStatus = "pending"
	PcapsSingleResponseResultPcapsResponseFullStatusRunning           PcapsSingleResponseResultPcapsResponseFullStatus = "running"
	PcapsSingleResponseResultPcapsResponseFullStatusConversionPending PcapsSingleResponseResultPcapsResponseFullStatus = "conversion_pending"
	PcapsSingleResponseResultPcapsResponseFullStatusConversionRunning PcapsSingleResponseResultPcapsResponseFullStatus = "conversion_running"
	PcapsSingleResponseResultPcapsResponseFullStatusComplete          PcapsSingleResponseResultPcapsResponseFullStatus = "complete"
	PcapsSingleResponseResultPcapsResponseFullStatusFailed            PcapsSingleResponseResultPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapsSingleResponseResultPcapsResponseFullSystem string

const (
	PcapsSingleResponseResultPcapsResponseFullSystemMagicTransit PcapsSingleResponseResultPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapsSingleResponseResultPcapsResponseFullType string

const (
	PcapsSingleResponseResultPcapsResponseFullTypeSimple PcapsSingleResponseResultPcapsResponseFullType = "simple"
	PcapsSingleResponseResultPcapsResponseFullTypeFull   PcapsSingleResponseResultPcapsResponseFullType = "full"
)

// Whether the API call was successful
type PcapsSingleResponseSuccess bool

const (
	PcapsSingleResponseSuccessTrue PcapsSingleResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimple],
// [AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFull].
type AccountPcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams()
}

type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimple) ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleFilterV1 struct {
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

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFull) ImplementsAccountPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullSystem string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullSystemMagicTransit AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullType string

const (
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullTypeSimple AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullType = "simple"
	AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullTypeFull   AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullFilterV1 struct {
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

func (r AccountPcapMagicPcapCollectionNewPcapRequestParamsPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
