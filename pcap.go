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

// Create new PCAP request for account.
func (r *PcapService) New(ctx context.Context, accountIdentifier string, body PcapNewParams, opts ...option.RequestOption) (res *PcapNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all packet capture requests for an account.
func (r *PcapService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PcapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Union satisfied by [PcapNewResponseAGiBwgNdPcapsResponseSimple] or
// [PcapNewResponseAGiBwgNdPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseAGiBwgNdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseAGiBwgNdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseAGiBwgNdPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseAGiBwgNdPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseAGiBwgNdPcapsResponseSimple]
type pcapNewResponseAGiBwgNdPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseAGiBwgNdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseAGiBwgNdPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1]
type pcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseAGiBwgNdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus string

const (
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusUnknown           PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusSuccess           PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "success"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusPending           PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "pending"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusRunning           PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "running"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusConversionPending PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusConversionRunning PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusComplete          PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "complete"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleStatusFailed            PcapNewResponseAGiBwgNdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseAGiBwgNdPcapsResponseSimpleSystem string

const (
	PcapNewResponseAGiBwgNdPcapsResponseSimpleSystemMagicTransit PcapNewResponseAGiBwgNdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseAGiBwgNdPcapsResponseSimpleType string

const (
	PcapNewResponseAGiBwgNdPcapsResponseSimpleTypeSimple PcapNewResponseAGiBwgNdPcapsResponseSimpleType = "simple"
	PcapNewResponseAGiBwgNdPcapsResponseSimpleTypeFull   PcapNewResponseAGiBwgNdPcapsResponseSimpleType = "full"
)

type PcapNewResponseAGiBwgNdPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseAGiBwgNdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseAGiBwgNdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseAGiBwgNdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseAGiBwgNdPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseAGiBwgNdPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseAGiBwgNdPcapsResponseFull]
type pcapNewResponseAGiBwgNdPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseAGiBwgNdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseAGiBwgNdPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseAGiBwgNdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseAGiBwgNdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseAGiBwgNdPcapsResponseFullFilterV1]
type pcapNewResponseAGiBwgNdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseAGiBwgNdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseAGiBwgNdPcapsResponseFullStatus string

const (
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusUnknown           PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "unknown"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusSuccess           PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "success"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusPending           PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "pending"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusRunning           PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "running"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusConversionPending PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusConversionRunning PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusComplete          PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "complete"
	PcapNewResponseAGiBwgNdPcapsResponseFullStatusFailed            PcapNewResponseAGiBwgNdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseAGiBwgNdPcapsResponseFullSystem string

const (
	PcapNewResponseAGiBwgNdPcapsResponseFullSystemMagicTransit PcapNewResponseAGiBwgNdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseAGiBwgNdPcapsResponseFullType string

const (
	PcapNewResponseAGiBwgNdPcapsResponseFullTypeSimple PcapNewResponseAGiBwgNdPcapsResponseFullType = "simple"
	PcapNewResponseAGiBwgNdPcapsResponseFullTypeFull   PcapNewResponseAGiBwgNdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseAGiBwgNdPcapsResponseSimple] or
// [PcapListResponseAGiBwgNdPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseAGiBwgNdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseAGiBwgNdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseAGiBwgNdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseAGiBwgNdPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseAGiBwgNdPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseAGiBwgNdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseAGiBwgNdPcapsResponseSimple]
type pcapListResponseAGiBwgNdPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseAGiBwgNdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseAGiBwgNdPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                 `json:"source_port"`
	JSON       pcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1]
type pcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseAGiBwgNdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseAGiBwgNdPcapsResponseSimpleStatus string

const (
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusUnknown           PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "unknown"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusSuccess           PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "success"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusPending           PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "pending"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusRunning           PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "running"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusConversionPending PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusConversionRunning PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusComplete          PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "complete"
	PcapListResponseAGiBwgNdPcapsResponseSimpleStatusFailed            PcapListResponseAGiBwgNdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseAGiBwgNdPcapsResponseSimpleSystem string

const (
	PcapListResponseAGiBwgNdPcapsResponseSimpleSystemMagicTransit PcapListResponseAGiBwgNdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseAGiBwgNdPcapsResponseSimpleType string

const (
	PcapListResponseAGiBwgNdPcapsResponseSimpleTypeSimple PcapListResponseAGiBwgNdPcapsResponseSimpleType = "simple"
	PcapListResponseAGiBwgNdPcapsResponseSimpleTypeFull   PcapListResponseAGiBwgNdPcapsResponseSimpleType = "full"
)

type PcapListResponseAGiBwgNdPcapsResponseFull struct {
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
	FilterV1 PcapListResponseAGiBwgNdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseAGiBwgNdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseAGiBwgNdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseAGiBwgNdPcapsResponseFullType `json:"type"`
	JSON pcapListResponseAGiBwgNdPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseAGiBwgNdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseAGiBwgNdPcapsResponseFull]
type pcapListResponseAGiBwgNdPcapsResponseFullJSON struct {
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

func (r *PcapListResponseAGiBwgNdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseAGiBwgNdPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseAGiBwgNdPcapsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                               `json:"source_port"`
	JSON       pcapListResponseAGiBwgNdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseAGiBwgNdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseAGiBwgNdPcapsResponseFullFilterV1]
type pcapListResponseAGiBwgNdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseAGiBwgNdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseAGiBwgNdPcapsResponseFullStatus string

const (
	PcapListResponseAGiBwgNdPcapsResponseFullStatusUnknown           PcapListResponseAGiBwgNdPcapsResponseFullStatus = "unknown"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusSuccess           PcapListResponseAGiBwgNdPcapsResponseFullStatus = "success"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusPending           PcapListResponseAGiBwgNdPcapsResponseFullStatus = "pending"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusRunning           PcapListResponseAGiBwgNdPcapsResponseFullStatus = "running"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusConversionPending PcapListResponseAGiBwgNdPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusConversionRunning PcapListResponseAGiBwgNdPcapsResponseFullStatus = "conversion_running"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusComplete          PcapListResponseAGiBwgNdPcapsResponseFullStatus = "complete"
	PcapListResponseAGiBwgNdPcapsResponseFullStatusFailed            PcapListResponseAGiBwgNdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseAGiBwgNdPcapsResponseFullSystem string

const (
	PcapListResponseAGiBwgNdPcapsResponseFullSystemMagicTransit PcapListResponseAGiBwgNdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseAGiBwgNdPcapsResponseFullType string

const (
	PcapListResponseAGiBwgNdPcapsResponseFullTypeSimple PcapListResponseAGiBwgNdPcapsResponseFullType = "simple"
	PcapListResponseAGiBwgNdPcapsResponseFullTypeFull   PcapListResponseAGiBwgNdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseAGiBwgNdPcapsResponseSimple] or
// [PcapGetResponseAGiBwgNdPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseAGiBwgNdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseAGiBwgNdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseAGiBwgNdPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseAGiBwgNdPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseAGiBwgNdPcapsResponseSimple]
type pcapGetResponseAGiBwgNdPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseAGiBwgNdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseAGiBwgNdPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1]
type pcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseAGiBwgNdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus string

const (
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusUnknown           PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusSuccess           PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "success"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusPending           PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "pending"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusRunning           PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "running"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusConversionPending PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusConversionRunning PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusComplete          PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "complete"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleStatusFailed            PcapGetResponseAGiBwgNdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseAGiBwgNdPcapsResponseSimpleSystem string

const (
	PcapGetResponseAGiBwgNdPcapsResponseSimpleSystemMagicTransit PcapGetResponseAGiBwgNdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseAGiBwgNdPcapsResponseSimpleType string

const (
	PcapGetResponseAGiBwgNdPcapsResponseSimpleTypeSimple PcapGetResponseAGiBwgNdPcapsResponseSimpleType = "simple"
	PcapGetResponseAGiBwgNdPcapsResponseSimpleTypeFull   PcapGetResponseAGiBwgNdPcapsResponseSimpleType = "full"
)

type PcapGetResponseAGiBwgNdPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseAGiBwgNdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseAGiBwgNdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseAGiBwgNdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseAGiBwgNdPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseAGiBwgNdPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseAGiBwgNdPcapsResponseFull]
type pcapGetResponseAGiBwgNdPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseAGiBwgNdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseAGiBwgNdPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseAGiBwgNdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseAGiBwgNdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseAGiBwgNdPcapsResponseFullFilterV1]
type pcapGetResponseAGiBwgNdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseAGiBwgNdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseAGiBwgNdPcapsResponseFullStatus string

const (
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusUnknown           PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "unknown"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusSuccess           PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "success"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusPending           PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "pending"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusRunning           PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "running"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusConversionPending PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusConversionRunning PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusComplete          PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "complete"
	PcapGetResponseAGiBwgNdPcapsResponseFullStatusFailed            PcapGetResponseAGiBwgNdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseAGiBwgNdPcapsResponseFullSystem string

const (
	PcapGetResponseAGiBwgNdPcapsResponseFullSystemMagicTransit PcapGetResponseAGiBwgNdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseAGiBwgNdPcapsResponseFullType string

const (
	PcapGetResponseAGiBwgNdPcapsResponseFullTypeSimple PcapGetResponseAGiBwgNdPcapsResponseFullType = "simple"
	PcapGetResponseAGiBwgNdPcapsResponseFullTypeFull   PcapGetResponseAGiBwgNdPcapsResponseFullType = "full"
)

type PcapNewParams struct {
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string]                `json:"destination_conf"`
	FilterV1        param.Field[PcapNewParamsFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The system used to collect packet captures.
type PcapNewParamsSystem string

const (
	PcapNewParamsSystemMagicTransit PcapNewParamsSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsType string

const (
	PcapNewParamsTypeSimple PcapNewParamsType = "simple"
	PcapNewParamsTypeFull   PcapNewParamsType = "full"
)

type PcapNewParamsFilterV1 struct {
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

func (r PcapNewParamsFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewResponseEnvelope struct {
	Errors   []PcapNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapNewResponseEnvelopeJSON    `json:"-"`
}

// pcapNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PcapNewResponseEnvelope]
type pcapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pcapNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PcapNewResponseEnvelopeErrors]
type pcapNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pcapNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PcapNewResponseEnvelopeMessages]
type pcapNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapNewResponseEnvelopeSuccess bool

const (
	PcapNewResponseEnvelopeSuccessTrue PcapNewResponseEnvelopeSuccess = true
)

type PcapListResponseEnvelope struct {
	Errors   []PcapListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PcapListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PcapListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PcapListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapListResponseEnvelopeJSON       `json:"-"`
}

// pcapListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PcapListResponseEnvelope]
type pcapListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    pcapListResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PcapListResponseEnvelopeErrors]
type pcapListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    pcapListResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PcapListResponseEnvelopeMessages]
type pcapListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapListResponseEnvelopeSuccess bool

const (
	PcapListResponseEnvelopeSuccessTrue PcapListResponseEnvelopeSuccess = true
)

type PcapListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       pcapListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [PcapListResponseEnvelopeResultInfo]
type pcapListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
