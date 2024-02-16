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

// Union satisfied by [PcapGetResponse2Ix5VjOtPcapsResponseSimple] or
// [PcapGetResponse2Ix5VjOtPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse2Ix5VjOtPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2Ix5VjOtPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2Ix5VjOtPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse2Ix5VjOtPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse2Ix5VjOtPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse2Ix5VjOtPcapsResponseSimple]
type pcapGetResponse2Ix5VjOtPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse2Ix5VjOtPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2Ix5VjOtPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1]
type pcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2Ix5VjOtPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusUnknown           PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusSuccess           PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "success"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusPending           PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "pending"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusRunning           PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "running"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusConversionPending PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusConversionRunning PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusComplete          PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "complete"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatusFailed            PcapGetResponse2Ix5VjOtPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2Ix5VjOtPcapsResponseSimpleSystem string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleSystemMagicTransit PcapGetResponse2Ix5VjOtPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2Ix5VjOtPcapsResponseSimpleType string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleTypeSimple PcapGetResponse2Ix5VjOtPcapsResponseSimpleType = "simple"
	PcapGetResponse2Ix5VjOtPcapsResponseSimpleTypeFull   PcapGetResponse2Ix5VjOtPcapsResponseSimpleType = "full"
)

type PcapGetResponse2Ix5VjOtPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2Ix5VjOtPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2Ix5VjOtPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2Ix5VjOtPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse2Ix5VjOtPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse2Ix5VjOtPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse2Ix5VjOtPcapsResponseFull]
type pcapGetResponse2Ix5VjOtPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse2Ix5VjOtPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2Ix5VjOtPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1]
type pcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2Ix5VjOtPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2Ix5VjOtPcapsResponseFullStatus string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusUnknown           PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "unknown"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusSuccess           PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "success"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusPending           PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "pending"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusRunning           PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "running"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusConversionPending PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusConversionRunning PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusComplete          PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "complete"
	PcapGetResponse2Ix5VjOtPcapsResponseFullStatusFailed            PcapGetResponse2Ix5VjOtPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2Ix5VjOtPcapsResponseFullSystem string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseFullSystemMagicTransit PcapGetResponse2Ix5VjOtPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2Ix5VjOtPcapsResponseFullType string

const (
	PcapGetResponse2Ix5VjOtPcapsResponseFullTypeSimple PcapGetResponse2Ix5VjOtPcapsResponseFullType = "simple"
	PcapGetResponse2Ix5VjOtPcapsResponseFullTypeFull   PcapGetResponse2Ix5VjOtPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponse2Ix5VjOtPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponse2Ix5VjOtPcapsResponseFullType = "full"
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
// [PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParams2Ix5VjOtPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
