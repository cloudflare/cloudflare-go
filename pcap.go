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

// Union satisfied by [PcapNewResponseF0NvmgHxPcapsResponseSimple] or
// [PcapNewResponseF0NvmgHxPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseF0NvmgHxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseF0NvmgHxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseF0NvmgHxPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseF0NvmgHxPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseF0NvmgHxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseF0NvmgHxPcapsResponseSimple]
type pcapNewResponseF0NvmgHxPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseF0NvmgHxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseF0NvmgHxPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1]
type pcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseF0NvmgHxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus string

const (
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusUnknown           PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusSuccess           PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "success"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusPending           PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "pending"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusRunning           PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "running"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusConversionPending PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusConversionRunning PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusComplete          PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "complete"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleStatusFailed            PcapNewResponseF0NvmgHxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseF0NvmgHxPcapsResponseSimpleSystem string

const (
	PcapNewResponseF0NvmgHxPcapsResponseSimpleSystemMagicTransit PcapNewResponseF0NvmgHxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseF0NvmgHxPcapsResponseSimpleType string

const (
	PcapNewResponseF0NvmgHxPcapsResponseSimpleTypeSimple PcapNewResponseF0NvmgHxPcapsResponseSimpleType = "simple"
	PcapNewResponseF0NvmgHxPcapsResponseSimpleTypeFull   PcapNewResponseF0NvmgHxPcapsResponseSimpleType = "full"
)

type PcapNewResponseF0NvmgHxPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseF0NvmgHxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseF0NvmgHxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseF0NvmgHxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseF0NvmgHxPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseF0NvmgHxPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseF0NvmgHxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseF0NvmgHxPcapsResponseFull]
type pcapNewResponseF0NvmgHxPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseF0NvmgHxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseF0NvmgHxPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseF0NvmgHxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseF0NvmgHxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseF0NvmgHxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseF0NvmgHxPcapsResponseFullFilterV1]
type pcapNewResponseF0NvmgHxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseF0NvmgHxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseF0NvmgHxPcapsResponseFullStatus string

const (
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusUnknown           PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "unknown"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusSuccess           PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "success"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusPending           PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "pending"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusRunning           PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "running"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusConversionPending PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusConversionRunning PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusComplete          PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "complete"
	PcapNewResponseF0NvmgHxPcapsResponseFullStatusFailed            PcapNewResponseF0NvmgHxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseF0NvmgHxPcapsResponseFullSystem string

const (
	PcapNewResponseF0NvmgHxPcapsResponseFullSystemMagicTransit PcapNewResponseF0NvmgHxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseF0NvmgHxPcapsResponseFullType string

const (
	PcapNewResponseF0NvmgHxPcapsResponseFullTypeSimple PcapNewResponseF0NvmgHxPcapsResponseFullType = "simple"
	PcapNewResponseF0NvmgHxPcapsResponseFullTypeFull   PcapNewResponseF0NvmgHxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseF0NvmgHxPcapsResponseSimple] or
// [PcapListResponseF0NvmgHxPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseF0NvmgHxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseF0NvmgHxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseF0NvmgHxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseF0NvmgHxPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseF0NvmgHxPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseF0NvmgHxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseF0NvmgHxPcapsResponseSimple]
type pcapListResponseF0NvmgHxPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseF0NvmgHxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseF0NvmgHxPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1]
type pcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseF0NvmgHxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseF0NvmgHxPcapsResponseSimpleStatus string

const (
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusUnknown           PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "unknown"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusSuccess           PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "success"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusPending           PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "pending"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusRunning           PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "running"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusConversionPending PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusConversionRunning PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusComplete          PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "complete"
	PcapListResponseF0NvmgHxPcapsResponseSimpleStatusFailed            PcapListResponseF0NvmgHxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseF0NvmgHxPcapsResponseSimpleSystem string

const (
	PcapListResponseF0NvmgHxPcapsResponseSimpleSystemMagicTransit PcapListResponseF0NvmgHxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseF0NvmgHxPcapsResponseSimpleType string

const (
	PcapListResponseF0NvmgHxPcapsResponseSimpleTypeSimple PcapListResponseF0NvmgHxPcapsResponseSimpleType = "simple"
	PcapListResponseF0NvmgHxPcapsResponseSimpleTypeFull   PcapListResponseF0NvmgHxPcapsResponseSimpleType = "full"
)

type PcapListResponseF0NvmgHxPcapsResponseFull struct {
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
	FilterV1 PcapListResponseF0NvmgHxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseF0NvmgHxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseF0NvmgHxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseF0NvmgHxPcapsResponseFullType `json:"type"`
	JSON pcapListResponseF0NvmgHxPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseF0NvmgHxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseF0NvmgHxPcapsResponseFull]
type pcapListResponseF0NvmgHxPcapsResponseFullJSON struct {
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

func (r *PcapListResponseF0NvmgHxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseF0NvmgHxPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseF0NvmgHxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseF0NvmgHxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseF0NvmgHxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseF0NvmgHxPcapsResponseFullFilterV1]
type pcapListResponseF0NvmgHxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseF0NvmgHxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseF0NvmgHxPcapsResponseFullStatus string

const (
	PcapListResponseF0NvmgHxPcapsResponseFullStatusUnknown           PcapListResponseF0NvmgHxPcapsResponseFullStatus = "unknown"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusSuccess           PcapListResponseF0NvmgHxPcapsResponseFullStatus = "success"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusPending           PcapListResponseF0NvmgHxPcapsResponseFullStatus = "pending"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusRunning           PcapListResponseF0NvmgHxPcapsResponseFullStatus = "running"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusConversionPending PcapListResponseF0NvmgHxPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusConversionRunning PcapListResponseF0NvmgHxPcapsResponseFullStatus = "conversion_running"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusComplete          PcapListResponseF0NvmgHxPcapsResponseFullStatus = "complete"
	PcapListResponseF0NvmgHxPcapsResponseFullStatusFailed            PcapListResponseF0NvmgHxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseF0NvmgHxPcapsResponseFullSystem string

const (
	PcapListResponseF0NvmgHxPcapsResponseFullSystemMagicTransit PcapListResponseF0NvmgHxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseF0NvmgHxPcapsResponseFullType string

const (
	PcapListResponseF0NvmgHxPcapsResponseFullTypeSimple PcapListResponseF0NvmgHxPcapsResponseFullType = "simple"
	PcapListResponseF0NvmgHxPcapsResponseFullTypeFull   PcapListResponseF0NvmgHxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseF0NvmgHxPcapsResponseSimple] or
// [PcapGetResponseF0NvmgHxPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseF0NvmgHxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseF0NvmgHxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseF0NvmgHxPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseF0NvmgHxPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseF0NvmgHxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseF0NvmgHxPcapsResponseSimple]
type pcapGetResponseF0NvmgHxPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseF0NvmgHxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseF0NvmgHxPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1]
type pcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseF0NvmgHxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus string

const (
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusUnknown           PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusSuccess           PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "success"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusPending           PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "pending"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusRunning           PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "running"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusConversionPending PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusConversionRunning PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusComplete          PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "complete"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleStatusFailed            PcapGetResponseF0NvmgHxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseF0NvmgHxPcapsResponseSimpleSystem string

const (
	PcapGetResponseF0NvmgHxPcapsResponseSimpleSystemMagicTransit PcapGetResponseF0NvmgHxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseF0NvmgHxPcapsResponseSimpleType string

const (
	PcapGetResponseF0NvmgHxPcapsResponseSimpleTypeSimple PcapGetResponseF0NvmgHxPcapsResponseSimpleType = "simple"
	PcapGetResponseF0NvmgHxPcapsResponseSimpleTypeFull   PcapGetResponseF0NvmgHxPcapsResponseSimpleType = "full"
)

type PcapGetResponseF0NvmgHxPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseF0NvmgHxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseF0NvmgHxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseF0NvmgHxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseF0NvmgHxPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseF0NvmgHxPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseF0NvmgHxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseF0NvmgHxPcapsResponseFull]
type pcapGetResponseF0NvmgHxPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseF0NvmgHxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseF0NvmgHxPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseF0NvmgHxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseF0NvmgHxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseF0NvmgHxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseF0NvmgHxPcapsResponseFullFilterV1]
type pcapGetResponseF0NvmgHxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseF0NvmgHxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseF0NvmgHxPcapsResponseFullStatus string

const (
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusUnknown           PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "unknown"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusSuccess           PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "success"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusPending           PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "pending"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusRunning           PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "running"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusConversionPending PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusConversionRunning PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusComplete          PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "complete"
	PcapGetResponseF0NvmgHxPcapsResponseFullStatusFailed            PcapGetResponseF0NvmgHxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseF0NvmgHxPcapsResponseFullSystem string

const (
	PcapGetResponseF0NvmgHxPcapsResponseFullSystemMagicTransit PcapGetResponseF0NvmgHxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseF0NvmgHxPcapsResponseFullType string

const (
	PcapGetResponseF0NvmgHxPcapsResponseFullTypeSimple PcapGetResponseF0NvmgHxPcapsResponseFullType = "simple"
	PcapGetResponseF0NvmgHxPcapsResponseFullTypeFull   PcapGetResponseF0NvmgHxPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsF0NvmgHxPcapsRequestSimple],
// [PcapNewParamsF0NvmgHxPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsF0NvmgHxPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsF0NvmgHxPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsF0NvmgHxPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsF0NvmgHxPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsF0NvmgHxPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsF0NvmgHxPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsF0NvmgHxPcapsRequestSimpleSystem string

const (
	PcapNewParamsF0NvmgHxPcapsRequestSimpleSystemMagicTransit PcapNewParamsF0NvmgHxPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsF0NvmgHxPcapsRequestSimpleType string

const (
	PcapNewParamsF0NvmgHxPcapsRequestSimpleTypeSimple PcapNewParamsF0NvmgHxPcapsRequestSimpleType = "simple"
	PcapNewParamsF0NvmgHxPcapsRequestSimpleTypeFull   PcapNewParamsF0NvmgHxPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsF0NvmgHxPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsF0NvmgHxPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsF0NvmgHxPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsF0NvmgHxPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsF0NvmgHxPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsF0NvmgHxPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsF0NvmgHxPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsF0NvmgHxPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsF0NvmgHxPcapsRequestFullSystem string

const (
	PcapNewParamsF0NvmgHxPcapsRequestFullSystemMagicTransit PcapNewParamsF0NvmgHxPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsF0NvmgHxPcapsRequestFullType string

const (
	PcapNewParamsF0NvmgHxPcapsRequestFullTypeSimple PcapNewParamsF0NvmgHxPcapsRequestFullType = "simple"
	PcapNewParamsF0NvmgHxPcapsRequestFullTypeFull   PcapNewParamsF0NvmgHxPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsF0NvmgHxPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsF0NvmgHxPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
