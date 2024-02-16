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

// Union satisfied by [PcapGetResponseNiUlIPbzPcapsResponseSimple] or
// [PcapGetResponseNiUlIPbzPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseNiUlIPbzPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseNiUlIPbzPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseNiUlIPbzPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseNiUlIPbzPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseNiUlIPbzPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseNiUlIPbzPcapsResponseSimple]
type pcapGetResponseNiUlIPbzPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseNiUlIPbzPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseNiUlIPbzPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1]
type pcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseNiUlIPbzPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus string

const (
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusUnknown           PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusSuccess           PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "success"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusPending           PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "pending"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusRunning           PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "running"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusConversionPending PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusConversionRunning PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusComplete          PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "complete"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleStatusFailed            PcapGetResponseNiUlIPbzPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseNiUlIPbzPcapsResponseSimpleSystem string

const (
	PcapGetResponseNiUlIPbzPcapsResponseSimpleSystemMagicTransit PcapGetResponseNiUlIPbzPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseNiUlIPbzPcapsResponseSimpleType string

const (
	PcapGetResponseNiUlIPbzPcapsResponseSimpleTypeSimple PcapGetResponseNiUlIPbzPcapsResponseSimpleType = "simple"
	PcapGetResponseNiUlIPbzPcapsResponseSimpleTypeFull   PcapGetResponseNiUlIPbzPcapsResponseSimpleType = "full"
)

type PcapGetResponseNiUlIPbzPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseNiUlIPbzPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseNiUlIPbzPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseNiUlIPbzPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseNiUlIPbzPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseNiUlIPbzPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseNiUlIPbzPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseNiUlIPbzPcapsResponseFull]
type pcapGetResponseNiUlIPbzPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseNiUlIPbzPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseNiUlIPbzPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseNiUlIPbzPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseNiUlIPbzPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseNiUlIPbzPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseNiUlIPbzPcapsResponseFullFilterV1]
type pcapGetResponseNiUlIPbzPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseNiUlIPbzPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseNiUlIPbzPcapsResponseFullStatus string

const (
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusUnknown           PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "unknown"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusSuccess           PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "success"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusPending           PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "pending"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusRunning           PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "running"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusConversionPending PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusConversionRunning PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusComplete          PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "complete"
	PcapGetResponseNiUlIPbzPcapsResponseFullStatusFailed            PcapGetResponseNiUlIPbzPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseNiUlIPbzPcapsResponseFullSystem string

const (
	PcapGetResponseNiUlIPbzPcapsResponseFullSystemMagicTransit PcapGetResponseNiUlIPbzPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseNiUlIPbzPcapsResponseFullType string

const (
	PcapGetResponseNiUlIPbzPcapsResponseFullTypeSimple PcapGetResponseNiUlIPbzPcapsResponseFullType = "simple"
	PcapGetResponseNiUlIPbzPcapsResponseFullTypeFull   PcapGetResponseNiUlIPbzPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseNiUlIPbzPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseNiUlIPbzPcapsResponseFullType = "full"
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
// [PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsNiUlIPbzPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
