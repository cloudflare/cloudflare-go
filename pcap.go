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

// Union satisfied by [PcapGetResponseY652Sr8cPcapsResponseSimple] or
// [PcapGetResponseY652Sr8cPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseY652Sr8cPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseY652Sr8cPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseY652Sr8cPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseY652Sr8cPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseY652Sr8cPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseY652Sr8cPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseY652Sr8cPcapsResponseSimple]
type pcapGetResponseY652Sr8cPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseY652Sr8cPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseY652Sr8cPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1]
type pcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseY652Sr8cPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseY652Sr8cPcapsResponseSimpleStatus string

const (
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusUnknown           PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusSuccess           PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "success"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusPending           PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "pending"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusRunning           PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "running"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusConversionPending PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusConversionRunning PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusComplete          PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "complete"
	PcapGetResponseY652Sr8cPcapsResponseSimpleStatusFailed            PcapGetResponseY652Sr8cPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseY652Sr8cPcapsResponseSimpleSystem string

const (
	PcapGetResponseY652Sr8cPcapsResponseSimpleSystemMagicTransit PcapGetResponseY652Sr8cPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseY652Sr8cPcapsResponseSimpleType string

const (
	PcapGetResponseY652Sr8cPcapsResponseSimpleTypeSimple PcapGetResponseY652Sr8cPcapsResponseSimpleType = "simple"
	PcapGetResponseY652Sr8cPcapsResponseSimpleTypeFull   PcapGetResponseY652Sr8cPcapsResponseSimpleType = "full"
)

type PcapGetResponseY652Sr8cPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseY652Sr8cPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseY652Sr8cPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseY652Sr8cPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseY652Sr8cPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseY652Sr8cPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseY652Sr8cPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseY652Sr8cPcapsResponseFull]
type pcapGetResponseY652Sr8cPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseY652Sr8cPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseY652Sr8cPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseY652Sr8cPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseY652Sr8cPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseY652Sr8cPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseY652Sr8cPcapsResponseFullFilterV1]
type pcapGetResponseY652Sr8cPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseY652Sr8cPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseY652Sr8cPcapsResponseFullStatus string

const (
	PcapGetResponseY652Sr8cPcapsResponseFullStatusUnknown           PcapGetResponseY652Sr8cPcapsResponseFullStatus = "unknown"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusSuccess           PcapGetResponseY652Sr8cPcapsResponseFullStatus = "success"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusPending           PcapGetResponseY652Sr8cPcapsResponseFullStatus = "pending"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusRunning           PcapGetResponseY652Sr8cPcapsResponseFullStatus = "running"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusConversionPending PcapGetResponseY652Sr8cPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusConversionRunning PcapGetResponseY652Sr8cPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusComplete          PcapGetResponseY652Sr8cPcapsResponseFullStatus = "complete"
	PcapGetResponseY652Sr8cPcapsResponseFullStatusFailed            PcapGetResponseY652Sr8cPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseY652Sr8cPcapsResponseFullSystem string

const (
	PcapGetResponseY652Sr8cPcapsResponseFullSystemMagicTransit PcapGetResponseY652Sr8cPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseY652Sr8cPcapsResponseFullType string

const (
	PcapGetResponseY652Sr8cPcapsResponseFullTypeSimple PcapGetResponseY652Sr8cPcapsResponseFullType = "simple"
	PcapGetResponseY652Sr8cPcapsResponseFullTypeFull   PcapGetResponseY652Sr8cPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponseY652Sr8cPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponseY652Sr8cPcapsResponseFullType = "full"
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
// [PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParamsY652Sr8cPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
