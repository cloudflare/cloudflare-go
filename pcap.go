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

// Union satisfied by [PcapNewResponseZoW8vz5fPcapsResponseSimple] or
// [PcapNewResponseZoW8vz5fPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseZoW8vz5fPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseZoW8vz5fPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseZoW8vz5fPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseZoW8vz5fPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseZoW8vz5fPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseZoW8vz5fPcapsResponseSimple]
type pcapNewResponseZoW8vz5fPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseZoW8vz5fPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseZoW8vz5fPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1]
type pcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseZoW8vz5fPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus string

const (
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusUnknown           PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusSuccess           PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "success"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusPending           PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "pending"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusRunning           PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "running"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusConversionPending PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusConversionRunning PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusComplete          PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "complete"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleStatusFailed            PcapNewResponseZoW8vz5fPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseZoW8vz5fPcapsResponseSimpleSystem string

const (
	PcapNewResponseZoW8vz5fPcapsResponseSimpleSystemMagicTransit PcapNewResponseZoW8vz5fPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseZoW8vz5fPcapsResponseSimpleType string

const (
	PcapNewResponseZoW8vz5fPcapsResponseSimpleTypeSimple PcapNewResponseZoW8vz5fPcapsResponseSimpleType = "simple"
	PcapNewResponseZoW8vz5fPcapsResponseSimpleTypeFull   PcapNewResponseZoW8vz5fPcapsResponseSimpleType = "full"
)

type PcapNewResponseZoW8vz5fPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseZoW8vz5fPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseZoW8vz5fPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseZoW8vz5fPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseZoW8vz5fPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseZoW8vz5fPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseZoW8vz5fPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseZoW8vz5fPcapsResponseFull]
type pcapNewResponseZoW8vz5fPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseZoW8vz5fPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseZoW8vz5fPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseZoW8vz5fPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseZoW8vz5fPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseZoW8vz5fPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseZoW8vz5fPcapsResponseFullFilterV1]
type pcapNewResponseZoW8vz5fPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseZoW8vz5fPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseZoW8vz5fPcapsResponseFullStatus string

const (
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusUnknown           PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "unknown"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusSuccess           PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "success"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusPending           PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "pending"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusRunning           PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "running"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusConversionPending PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusConversionRunning PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusComplete          PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "complete"
	PcapNewResponseZoW8vz5fPcapsResponseFullStatusFailed            PcapNewResponseZoW8vz5fPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseZoW8vz5fPcapsResponseFullSystem string

const (
	PcapNewResponseZoW8vz5fPcapsResponseFullSystemMagicTransit PcapNewResponseZoW8vz5fPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseZoW8vz5fPcapsResponseFullType string

const (
	PcapNewResponseZoW8vz5fPcapsResponseFullTypeSimple PcapNewResponseZoW8vz5fPcapsResponseFullType = "simple"
	PcapNewResponseZoW8vz5fPcapsResponseFullTypeFull   PcapNewResponseZoW8vz5fPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseZoW8vz5fPcapsResponseSimple] or
// [PcapListResponseZoW8vz5fPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseZoW8vz5fPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseZoW8vz5fPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseZoW8vz5fPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseZoW8vz5fPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseZoW8vz5fPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseZoW8vz5fPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseZoW8vz5fPcapsResponseSimple]
type pcapListResponseZoW8vz5fPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseZoW8vz5fPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseZoW8vz5fPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1]
type pcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseZoW8vz5fPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseZoW8vz5fPcapsResponseSimpleStatus string

const (
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusUnknown           PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "unknown"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusSuccess           PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "success"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusPending           PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "pending"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusRunning           PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "running"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusConversionPending PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusConversionRunning PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusComplete          PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "complete"
	PcapListResponseZoW8vz5fPcapsResponseSimpleStatusFailed            PcapListResponseZoW8vz5fPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseZoW8vz5fPcapsResponseSimpleSystem string

const (
	PcapListResponseZoW8vz5fPcapsResponseSimpleSystemMagicTransit PcapListResponseZoW8vz5fPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseZoW8vz5fPcapsResponseSimpleType string

const (
	PcapListResponseZoW8vz5fPcapsResponseSimpleTypeSimple PcapListResponseZoW8vz5fPcapsResponseSimpleType = "simple"
	PcapListResponseZoW8vz5fPcapsResponseSimpleTypeFull   PcapListResponseZoW8vz5fPcapsResponseSimpleType = "full"
)

type PcapListResponseZoW8vz5fPcapsResponseFull struct {
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
	FilterV1 PcapListResponseZoW8vz5fPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseZoW8vz5fPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseZoW8vz5fPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseZoW8vz5fPcapsResponseFullType `json:"type"`
	JSON pcapListResponseZoW8vz5fPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseZoW8vz5fPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseZoW8vz5fPcapsResponseFull]
type pcapListResponseZoW8vz5fPcapsResponseFullJSON struct {
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

func (r *PcapListResponseZoW8vz5fPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseZoW8vz5fPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseZoW8vz5fPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseZoW8vz5fPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseZoW8vz5fPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseZoW8vz5fPcapsResponseFullFilterV1]
type pcapListResponseZoW8vz5fPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseZoW8vz5fPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseZoW8vz5fPcapsResponseFullStatus string

const (
	PcapListResponseZoW8vz5fPcapsResponseFullStatusUnknown           PcapListResponseZoW8vz5fPcapsResponseFullStatus = "unknown"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusSuccess           PcapListResponseZoW8vz5fPcapsResponseFullStatus = "success"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusPending           PcapListResponseZoW8vz5fPcapsResponseFullStatus = "pending"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusRunning           PcapListResponseZoW8vz5fPcapsResponseFullStatus = "running"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusConversionPending PcapListResponseZoW8vz5fPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusConversionRunning PcapListResponseZoW8vz5fPcapsResponseFullStatus = "conversion_running"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusComplete          PcapListResponseZoW8vz5fPcapsResponseFullStatus = "complete"
	PcapListResponseZoW8vz5fPcapsResponseFullStatusFailed            PcapListResponseZoW8vz5fPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseZoW8vz5fPcapsResponseFullSystem string

const (
	PcapListResponseZoW8vz5fPcapsResponseFullSystemMagicTransit PcapListResponseZoW8vz5fPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseZoW8vz5fPcapsResponseFullType string

const (
	PcapListResponseZoW8vz5fPcapsResponseFullTypeSimple PcapListResponseZoW8vz5fPcapsResponseFullType = "simple"
	PcapListResponseZoW8vz5fPcapsResponseFullTypeFull   PcapListResponseZoW8vz5fPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseZoW8vz5fPcapsResponseSimple] or
// [PcapGetResponseZoW8vz5fPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseZoW8vz5fPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseZoW8vz5fPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseZoW8vz5fPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseZoW8vz5fPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseZoW8vz5fPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseZoW8vz5fPcapsResponseSimple]
type pcapGetResponseZoW8vz5fPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseZoW8vz5fPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseZoW8vz5fPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1]
type pcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseZoW8vz5fPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus string

const (
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusUnknown           PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusSuccess           PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "success"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusPending           PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "pending"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusRunning           PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "running"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusConversionPending PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusConversionRunning PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusComplete          PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "complete"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleStatusFailed            PcapGetResponseZoW8vz5fPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseZoW8vz5fPcapsResponseSimpleSystem string

const (
	PcapGetResponseZoW8vz5fPcapsResponseSimpleSystemMagicTransit PcapGetResponseZoW8vz5fPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseZoW8vz5fPcapsResponseSimpleType string

const (
	PcapGetResponseZoW8vz5fPcapsResponseSimpleTypeSimple PcapGetResponseZoW8vz5fPcapsResponseSimpleType = "simple"
	PcapGetResponseZoW8vz5fPcapsResponseSimpleTypeFull   PcapGetResponseZoW8vz5fPcapsResponseSimpleType = "full"
)

type PcapGetResponseZoW8vz5fPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseZoW8vz5fPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseZoW8vz5fPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseZoW8vz5fPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseZoW8vz5fPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseZoW8vz5fPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseZoW8vz5fPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseZoW8vz5fPcapsResponseFull]
type pcapGetResponseZoW8vz5fPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseZoW8vz5fPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseZoW8vz5fPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseZoW8vz5fPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseZoW8vz5fPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseZoW8vz5fPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseZoW8vz5fPcapsResponseFullFilterV1]
type pcapGetResponseZoW8vz5fPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseZoW8vz5fPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseZoW8vz5fPcapsResponseFullStatus string

const (
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusUnknown           PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "unknown"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusSuccess           PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "success"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusPending           PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "pending"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusRunning           PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "running"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusConversionPending PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusConversionRunning PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusComplete          PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "complete"
	PcapGetResponseZoW8vz5fPcapsResponseFullStatusFailed            PcapGetResponseZoW8vz5fPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseZoW8vz5fPcapsResponseFullSystem string

const (
	PcapGetResponseZoW8vz5fPcapsResponseFullSystemMagicTransit PcapGetResponseZoW8vz5fPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseZoW8vz5fPcapsResponseFullType string

const (
	PcapGetResponseZoW8vz5fPcapsResponseFullTypeSimple PcapGetResponseZoW8vz5fPcapsResponseFullType = "simple"
	PcapGetResponseZoW8vz5fPcapsResponseFullTypeFull   PcapGetResponseZoW8vz5fPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsZoW8vz5fPcapsRequestSimple],
// [PcapNewParamsZoW8vz5fPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsZoW8vz5fPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsZoW8vz5fPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsZoW8vz5fPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsZoW8vz5fPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsZoW8vz5fPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsZoW8vz5fPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsZoW8vz5fPcapsRequestSimpleSystem string

const (
	PcapNewParamsZoW8vz5fPcapsRequestSimpleSystemMagicTransit PcapNewParamsZoW8vz5fPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsZoW8vz5fPcapsRequestSimpleType string

const (
	PcapNewParamsZoW8vz5fPcapsRequestSimpleTypeSimple PcapNewParamsZoW8vz5fPcapsRequestSimpleType = "simple"
	PcapNewParamsZoW8vz5fPcapsRequestSimpleTypeFull   PcapNewParamsZoW8vz5fPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsZoW8vz5fPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsZoW8vz5fPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsZoW8vz5fPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsZoW8vz5fPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsZoW8vz5fPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsZoW8vz5fPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsZoW8vz5fPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsZoW8vz5fPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsZoW8vz5fPcapsRequestFullSystem string

const (
	PcapNewParamsZoW8vz5fPcapsRequestFullSystemMagicTransit PcapNewParamsZoW8vz5fPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsZoW8vz5fPcapsRequestFullType string

const (
	PcapNewParamsZoW8vz5fPcapsRequestFullTypeSimple PcapNewParamsZoW8vz5fPcapsRequestFullType = "simple"
	PcapNewParamsZoW8vz5fPcapsRequestFullTypeFull   PcapNewParamsZoW8vz5fPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsZoW8vz5fPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsZoW8vz5fPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
