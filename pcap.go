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

// Union satisfied by [PcapGetResponseFZsXwDu1PcapsResponseSimple] or
// [PcapGetResponseFZsXwDu1PcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseFZsXwDu1PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseFZsXwDu1PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseFZsXwDu1PcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseFZsXwDu1PcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseFZsXwDu1PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseFZsXwDu1PcapsResponseSimple]
type pcapGetResponseFZsXwDu1PcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseFZsXwDu1PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseFZsXwDu1PcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1]
type pcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseFZsXwDu1PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus string

const (
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusUnknown           PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "unknown"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusSuccess           PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "success"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusPending           PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "pending"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusRunning           PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "running"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusConversionPending PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusConversionRunning PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusComplete          PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "complete"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleStatusFailed            PcapGetResponseFZsXwDu1PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseFZsXwDu1PcapsResponseSimpleSystem string

const (
	PcapGetResponseFZsXwDu1PcapsResponseSimpleSystemMagicTransit PcapGetResponseFZsXwDu1PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseFZsXwDu1PcapsResponseSimpleType string

const (
	PcapGetResponseFZsXwDu1PcapsResponseSimpleTypeSimple PcapGetResponseFZsXwDu1PcapsResponseSimpleType = "simple"
	PcapGetResponseFZsXwDu1PcapsResponseSimpleTypeFull   PcapGetResponseFZsXwDu1PcapsResponseSimpleType = "full"
)

type PcapGetResponseFZsXwDu1PcapsResponseFull struct {
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
	FilterV1 PcapGetResponseFZsXwDu1PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseFZsXwDu1PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseFZsXwDu1PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseFZsXwDu1PcapsResponseFullType `json:"type"`
	JSON pcapGetResponseFZsXwDu1PcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseFZsXwDu1PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseFZsXwDu1PcapsResponseFull]
type pcapGetResponseFZsXwDu1PcapsResponseFullJSON struct {
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

func (r *PcapGetResponseFZsXwDu1PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseFZsXwDu1PcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseFZsXwDu1PcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseFZsXwDu1PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseFZsXwDu1PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseFZsXwDu1PcapsResponseFullFilterV1]
type pcapGetResponseFZsXwDu1PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseFZsXwDu1PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseFZsXwDu1PcapsResponseFullStatus string

const (
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusUnknown           PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "unknown"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusSuccess           PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "success"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusPending           PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "pending"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusRunning           PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "running"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusConversionPending PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusConversionRunning PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "conversion_running"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusComplete          PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "complete"
	PcapGetResponseFZsXwDu1PcapsResponseFullStatusFailed            PcapGetResponseFZsXwDu1PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseFZsXwDu1PcapsResponseFullSystem string

const (
	PcapGetResponseFZsXwDu1PcapsResponseFullSystemMagicTransit PcapGetResponseFZsXwDu1PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseFZsXwDu1PcapsResponseFullType string

const (
	PcapGetResponseFZsXwDu1PcapsResponseFullTypeSimple PcapGetResponseFZsXwDu1PcapsResponseFullType = "simple"
	PcapGetResponseFZsXwDu1PcapsResponseFullTypeFull   PcapGetResponseFZsXwDu1PcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseFZsXwDu1PcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseFZsXwDu1PcapsResponseFullType = "full"
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
// [PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsFZsXwDu1PcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
