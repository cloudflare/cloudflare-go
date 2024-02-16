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

// PcapService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPcapService] method instead.
type PcapService struct {
	Options    []option.RequestOption
	Ownerships *PcapOwnershipService
	Downloads  *PcapDownloadService
}

// NewPcapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPcapService(opts ...option.RequestOption) (r *PcapService) {
	r = &PcapService{}
	r.Options = opts
	r.Ownerships = NewPcapOwnershipService(opts...)
	r.Downloads = NewPcapDownloadService(opts...)
	return
}

// Get information for a PCAP request by id.
func (r *PcapService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *PcapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create new PCAP request for account.
func (r *PcapService) MagicPcapCollectionNewPcapRequest(ctx context.Context, accountIdentifier string, body PcapMagicPcapCollectionNewPcapRequestParams, opts ...option.RequestOption) (res *PcapMagicPcapCollectionNewPcapRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapMagicPcapCollectionNewPcapRequestResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all packet capture requests for an account.
func (r *PcapService) MagicPcapCollectionListPacketCaptureRequests(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PcapMagicPcapCollectionListPacketCaptureRequestsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PcapGetResponseWq9B1j1jPcapsResponseSimple] or
// [PcapGetResponseWq9B1j1jPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseWq9B1j1jPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseWq9B1j1jPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseWq9B1j1jPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseWq9B1j1jPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseWq9B1j1jPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseWq9B1j1jPcapsResponseSimple]
type pcapGetResponseWq9B1j1jPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseWq9B1j1jPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseWq9B1j1jPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1]
type pcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseWq9B1j1jPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus string

const (
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusUnknown           PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusSuccess           PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "success"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusPending           PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "pending"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusRunning           PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "running"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusConversionPending PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusConversionRunning PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusComplete          PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "complete"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleStatusFailed            PcapGetResponseWq9B1j1jPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseWq9B1j1jPcapsResponseSimpleSystem string

const (
	PcapGetResponseWq9B1j1jPcapsResponseSimpleSystemMagicTransit PcapGetResponseWq9B1j1jPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseWq9B1j1jPcapsResponseSimpleType string

const (
	PcapGetResponseWq9B1j1jPcapsResponseSimpleTypeSimple PcapGetResponseWq9B1j1jPcapsResponseSimpleType = "simple"
	PcapGetResponseWq9B1j1jPcapsResponseSimpleTypeFull   PcapGetResponseWq9B1j1jPcapsResponseSimpleType = "full"
)

type PcapGetResponseWq9B1j1jPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseWq9B1j1jPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseWq9B1j1jPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseWq9B1j1jPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseWq9B1j1jPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseWq9B1j1jPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseWq9B1j1jPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseWq9B1j1jPcapsResponseFull]
type pcapGetResponseWq9B1j1jPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseWq9B1j1jPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseWq9B1j1jPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseWq9B1j1jPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                              `json:"source_port"`
	JSON       pcapGetResponseWq9B1j1jPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseWq9B1j1jPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseWq9B1j1jPcapsResponseFullFilterV1]
type pcapGetResponseWq9B1j1jPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseWq9B1j1jPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseWq9B1j1jPcapsResponseFullStatus string

const (
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusUnknown           PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "unknown"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusSuccess           PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "success"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusPending           PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "pending"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusRunning           PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "running"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusConversionPending PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusConversionRunning PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusComplete          PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "complete"
	PcapGetResponseWq9B1j1jPcapsResponseFullStatusFailed            PcapGetResponseWq9B1j1jPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseWq9B1j1jPcapsResponseFullSystem string

const (
	PcapGetResponseWq9B1j1jPcapsResponseFullSystemMagicTransit PcapGetResponseWq9B1j1jPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseWq9B1j1jPcapsResponseFullType string

const (
	PcapGetResponseWq9B1j1jPcapsResponseFullTypeSimple PcapGetResponseWq9B1j1jPcapsResponseFullType = "simple"
	PcapGetResponseWq9B1j1jPcapsResponseFullTypeFull   PcapGetResponseWq9B1j1jPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                              `json:"source_port"`
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                            `json:"source_port"`
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseWq9B1j1jPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                                                       `json:"source_port"`
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseWq9B1j1jPcapsResponseFullType = "full"
)

type PcapGetResponseEnvelope struct {
	Errors   []PcapGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapGetResponseEnvelopeJSON    `json:"-"`
}

// pcapGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PcapGetResponseEnvelope]
type pcapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pcapGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PcapGetResponseEnvelopeErrors]
type pcapGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pcapGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PcapGetResponseEnvelopeMessages]
type pcapGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapGetResponseEnvelopeSuccess bool

const (
	PcapGetResponseEnvelopeSuccessTrue PcapGetResponseEnvelopeSuccess = true
)

// This interface is a union satisfied by one of the following:
// [PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsWq9B1j1jPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestResponseEnvelope struct {
	Errors   []PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapMagicPcapCollectionNewPcapRequestResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeJSON    `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeJSON contains the JSON
// metadata for the struct [PcapMagicPcapCollectionNewPcapRequestResponseEnvelope]
type pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrors]
type pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessages]
type pcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeSuccess bool

const (
	PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeSuccessTrue PcapMagicPcapCollectionNewPcapRequestResponseEnvelopeSuccess = true
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelope struct {
	Errors   []PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessages `json:"messages,required"`
	Result   []PcapMagicPcapCollectionListPacketCaptureRequestsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeJSON       `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelope]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrors]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessages]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeSuccess bool

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeSuccessTrue PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeSuccess = true
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                        `json:"total_count"`
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfo]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
