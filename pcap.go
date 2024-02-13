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

// Union satisfied by [PcapGetResponse7dRf5S6WPcapsResponseSimple] or
// [PcapGetResponse7dRf5S6WPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse7dRf5S6WPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7dRf5S6WPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7dRf5S6WPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse7dRf5S6WPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse7dRf5S6WPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse7dRf5S6WPcapsResponseSimple]
type pcapGetResponse7dRf5S6WPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse7dRf5S6WPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7dRf5S6WPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1]
type pcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7dRf5S6WPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus string

const (
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusUnknown           PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusSuccess           PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "success"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusPending           PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "pending"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusRunning           PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "running"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusConversionPending PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusConversionRunning PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusComplete          PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "complete"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleStatusFailed            PcapGetResponse7dRf5S6WPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7dRf5S6WPcapsResponseSimpleSystem string

const (
	PcapGetResponse7dRf5S6WPcapsResponseSimpleSystemMagicTransit PcapGetResponse7dRf5S6WPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7dRf5S6WPcapsResponseSimpleType string

const (
	PcapGetResponse7dRf5S6WPcapsResponseSimpleTypeSimple PcapGetResponse7dRf5S6WPcapsResponseSimpleType = "simple"
	PcapGetResponse7dRf5S6WPcapsResponseSimpleTypeFull   PcapGetResponse7dRf5S6WPcapsResponseSimpleType = "full"
)

type PcapGetResponse7dRf5S6WPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse7dRf5S6WPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7dRf5S6WPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7dRf5S6WPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7dRf5S6WPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse7dRf5S6WPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse7dRf5S6WPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse7dRf5S6WPcapsResponseFull]
type pcapGetResponse7dRf5S6WPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse7dRf5S6WPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7dRf5S6WPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7dRf5S6WPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse7dRf5S6WPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse7dRf5S6WPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse7dRf5S6WPcapsResponseFullFilterV1]
type pcapGetResponse7dRf5S6WPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7dRf5S6WPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7dRf5S6WPcapsResponseFullStatus string

const (
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusUnknown           PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "unknown"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusSuccess           PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "success"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusPending           PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "pending"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusRunning           PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "running"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusConversionPending PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusConversionRunning PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusComplete          PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "complete"
	PcapGetResponse7dRf5S6WPcapsResponseFullStatusFailed            PcapGetResponse7dRf5S6WPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7dRf5S6WPcapsResponseFullSystem string

const (
	PcapGetResponse7dRf5S6WPcapsResponseFullSystemMagicTransit PcapGetResponse7dRf5S6WPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7dRf5S6WPcapsResponseFullType string

const (
	PcapGetResponse7dRf5S6WPcapsResponseFullTypeSimple PcapGetResponse7dRf5S6WPcapsResponseFullType = "simple"
	PcapGetResponse7dRf5S6WPcapsResponseFullTypeFull   PcapGetResponse7dRf5S6WPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimple] or
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFull].
type PcapMagicPcapCollectionNewPcapRequestResponse interface {
	implementsPcapMagicPcapCollectionNewPcapRequestResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionNewPcapRequestResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimple]
type pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimple) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFull]
type pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFull) implementsPcapMagicPcapCollectionNewPcapRequestResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusPending           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusRunning           PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusComplete          PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatusFailed            PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullTypeSimple PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullTypeFull   PcapMagicPcapCollectionNewPcapRequestResponse7dRf5S6WPcapsResponseFullType = "full"
)

// Union satisfied by
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimple]
// or
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFull].
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse interface {
	implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapMagicPcapCollectionListPacketCaptureRequestsResponse)(nil)).Elem(), "")
}

type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimple]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimple) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseSimpleType = "full"
)

type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFull struct {
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
	FilterV1 PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullType `json:"type"`
	JSON pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullJSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullJSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFull]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullJSON struct {
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

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFull) implementsPcapMagicPcapCollectionListPacketCaptureRequestsResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1 struct {
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
	JSON       pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1JSON
// contains the JSON metadata for the struct
// [PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1]
type pcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusUnknown           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "unknown"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusSuccess           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "success"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusPending           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusRunning           PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusConversionPending PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "conversion_pending"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusConversionRunning PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "conversion_running"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusComplete          PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "complete"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatusFailed            PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullSystem string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullSystemMagicTransit PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullType string

const (
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullTypeSimple PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullType = "simple"
	PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullTypeFull   PcapMagicPcapCollectionListPacketCaptureRequestsResponse7dRf5S6WPcapsResponseFullType = "full"
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
// [PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimple],
// [PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFull].
type PcapMagicPcapCollectionNewPcapRequestParams interface {
	ImplementsPcapMagicPcapCollectionNewPcapRequestParams()
}

type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimple) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleType string

const (
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleTypeSimple PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleTypeFull   PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFull) ImplementsPcapMagicPcapCollectionNewPcapRequestParams() {

}

// The system used to collect packet captures.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullSystem string

const (
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullSystemMagicTransit PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullType string

const (
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullTypeSimple PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullType = "simple"
	PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullTypeFull   PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullFilterV1 struct {
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

func (r PcapMagicPcapCollectionNewPcapRequestParams7dRf5S6WPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
