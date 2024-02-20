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

// Union satisfied by [PcapNewResponseEqO6AITgPcapsResponseSimple] or
// [PcapNewResponseEqO6AITgPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseEqO6AITgPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseEqO6AITgPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseEqO6AITgPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseEqO6AITgPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseEqO6AITgPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseEqO6AITgPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseEqO6AITgPcapsResponseSimple]
type pcapNewResponseEqO6AITgPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseEqO6AITgPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseEqO6AITgPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1]
type pcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseEqO6AITgPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseEqO6AITgPcapsResponseSimpleStatus string

const (
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusUnknown           PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusSuccess           PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "success"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusPending           PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "pending"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusRunning           PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "running"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusConversionPending PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusConversionRunning PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusComplete          PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "complete"
	PcapNewResponseEqO6AITgPcapsResponseSimpleStatusFailed            PcapNewResponseEqO6AITgPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseEqO6AITgPcapsResponseSimpleSystem string

const (
	PcapNewResponseEqO6AITgPcapsResponseSimpleSystemMagicTransit PcapNewResponseEqO6AITgPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseEqO6AITgPcapsResponseSimpleType string

const (
	PcapNewResponseEqO6AITgPcapsResponseSimpleTypeSimple PcapNewResponseEqO6AITgPcapsResponseSimpleType = "simple"
	PcapNewResponseEqO6AITgPcapsResponseSimpleTypeFull   PcapNewResponseEqO6AITgPcapsResponseSimpleType = "full"
)

type PcapNewResponseEqO6AITgPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseEqO6AITgPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseEqO6AITgPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseEqO6AITgPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseEqO6AITgPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseEqO6AITgPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseEqO6AITgPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseEqO6AITgPcapsResponseFull]
type pcapNewResponseEqO6AITgPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseEqO6AITgPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseEqO6AITgPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseEqO6AITgPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseEqO6AITgPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseEqO6AITgPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseEqO6AITgPcapsResponseFullFilterV1]
type pcapNewResponseEqO6AITgPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseEqO6AITgPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseEqO6AITgPcapsResponseFullStatus string

const (
	PcapNewResponseEqO6AITgPcapsResponseFullStatusUnknown           PcapNewResponseEqO6AITgPcapsResponseFullStatus = "unknown"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusSuccess           PcapNewResponseEqO6AITgPcapsResponseFullStatus = "success"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusPending           PcapNewResponseEqO6AITgPcapsResponseFullStatus = "pending"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusRunning           PcapNewResponseEqO6AITgPcapsResponseFullStatus = "running"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusConversionPending PcapNewResponseEqO6AITgPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusConversionRunning PcapNewResponseEqO6AITgPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusComplete          PcapNewResponseEqO6AITgPcapsResponseFullStatus = "complete"
	PcapNewResponseEqO6AITgPcapsResponseFullStatusFailed            PcapNewResponseEqO6AITgPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseEqO6AITgPcapsResponseFullSystem string

const (
	PcapNewResponseEqO6AITgPcapsResponseFullSystemMagicTransit PcapNewResponseEqO6AITgPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseEqO6AITgPcapsResponseFullType string

const (
	PcapNewResponseEqO6AITgPcapsResponseFullTypeSimple PcapNewResponseEqO6AITgPcapsResponseFullType = "simple"
	PcapNewResponseEqO6AITgPcapsResponseFullTypeFull   PcapNewResponseEqO6AITgPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseEqO6AITgPcapsResponseSimple] or
// [PcapListResponseEqO6AITgPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseEqO6AITgPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseEqO6AITgPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseEqO6AITgPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseEqO6AITgPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseEqO6AITgPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseEqO6AITgPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseEqO6AITgPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseEqO6AITgPcapsResponseSimple]
type pcapListResponseEqO6AITgPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseEqO6AITgPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseEqO6AITgPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseEqO6AITgPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseEqO6AITgPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseEqO6AITgPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseEqO6AITgPcapsResponseSimpleFilterV1]
type pcapListResponseEqO6AITgPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseEqO6AITgPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseEqO6AITgPcapsResponseSimpleStatus string

const (
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusUnknown           PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "unknown"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusSuccess           PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "success"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusPending           PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "pending"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusRunning           PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "running"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusConversionPending PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusConversionRunning PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusComplete          PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "complete"
	PcapListResponseEqO6AITgPcapsResponseSimpleStatusFailed            PcapListResponseEqO6AITgPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseEqO6AITgPcapsResponseSimpleSystem string

const (
	PcapListResponseEqO6AITgPcapsResponseSimpleSystemMagicTransit PcapListResponseEqO6AITgPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseEqO6AITgPcapsResponseSimpleType string

const (
	PcapListResponseEqO6AITgPcapsResponseSimpleTypeSimple PcapListResponseEqO6AITgPcapsResponseSimpleType = "simple"
	PcapListResponseEqO6AITgPcapsResponseSimpleTypeFull   PcapListResponseEqO6AITgPcapsResponseSimpleType = "full"
)

type PcapListResponseEqO6AITgPcapsResponseFull struct {
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
	FilterV1 PcapListResponseEqO6AITgPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseEqO6AITgPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseEqO6AITgPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseEqO6AITgPcapsResponseFullType `json:"type"`
	JSON pcapListResponseEqO6AITgPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseEqO6AITgPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseEqO6AITgPcapsResponseFull]
type pcapListResponseEqO6AITgPcapsResponseFullJSON struct {
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

func (r *PcapListResponseEqO6AITgPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseEqO6AITgPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseEqO6AITgPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseEqO6AITgPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseEqO6AITgPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseEqO6AITgPcapsResponseFullFilterV1]
type pcapListResponseEqO6AITgPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseEqO6AITgPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseEqO6AITgPcapsResponseFullStatus string

const (
	PcapListResponseEqO6AITgPcapsResponseFullStatusUnknown           PcapListResponseEqO6AITgPcapsResponseFullStatus = "unknown"
	PcapListResponseEqO6AITgPcapsResponseFullStatusSuccess           PcapListResponseEqO6AITgPcapsResponseFullStatus = "success"
	PcapListResponseEqO6AITgPcapsResponseFullStatusPending           PcapListResponseEqO6AITgPcapsResponseFullStatus = "pending"
	PcapListResponseEqO6AITgPcapsResponseFullStatusRunning           PcapListResponseEqO6AITgPcapsResponseFullStatus = "running"
	PcapListResponseEqO6AITgPcapsResponseFullStatusConversionPending PcapListResponseEqO6AITgPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseEqO6AITgPcapsResponseFullStatusConversionRunning PcapListResponseEqO6AITgPcapsResponseFullStatus = "conversion_running"
	PcapListResponseEqO6AITgPcapsResponseFullStatusComplete          PcapListResponseEqO6AITgPcapsResponseFullStatus = "complete"
	PcapListResponseEqO6AITgPcapsResponseFullStatusFailed            PcapListResponseEqO6AITgPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseEqO6AITgPcapsResponseFullSystem string

const (
	PcapListResponseEqO6AITgPcapsResponseFullSystemMagicTransit PcapListResponseEqO6AITgPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseEqO6AITgPcapsResponseFullType string

const (
	PcapListResponseEqO6AITgPcapsResponseFullTypeSimple PcapListResponseEqO6AITgPcapsResponseFullType = "simple"
	PcapListResponseEqO6AITgPcapsResponseFullTypeFull   PcapListResponseEqO6AITgPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseEqO6AITgPcapsResponseSimple] or
// [PcapGetResponseEqO6AITgPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseEqO6AITgPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseEqO6AITgPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseEqO6AITgPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseEqO6AITgPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseEqO6AITgPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseEqO6AITgPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseEqO6AITgPcapsResponseSimple]
type pcapGetResponseEqO6AITgPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseEqO6AITgPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseEqO6AITgPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1]
type pcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseEqO6AITgPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseEqO6AITgPcapsResponseSimpleStatus string

const (
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusUnknown           PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusSuccess           PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "success"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusPending           PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "pending"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusRunning           PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "running"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusConversionPending PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusConversionRunning PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusComplete          PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "complete"
	PcapGetResponseEqO6AITgPcapsResponseSimpleStatusFailed            PcapGetResponseEqO6AITgPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseEqO6AITgPcapsResponseSimpleSystem string

const (
	PcapGetResponseEqO6AITgPcapsResponseSimpleSystemMagicTransit PcapGetResponseEqO6AITgPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseEqO6AITgPcapsResponseSimpleType string

const (
	PcapGetResponseEqO6AITgPcapsResponseSimpleTypeSimple PcapGetResponseEqO6AITgPcapsResponseSimpleType = "simple"
	PcapGetResponseEqO6AITgPcapsResponseSimpleTypeFull   PcapGetResponseEqO6AITgPcapsResponseSimpleType = "full"
)

type PcapGetResponseEqO6AITgPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseEqO6AITgPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseEqO6AITgPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseEqO6AITgPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseEqO6AITgPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseEqO6AITgPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseEqO6AITgPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseEqO6AITgPcapsResponseFull]
type pcapGetResponseEqO6AITgPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseEqO6AITgPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseEqO6AITgPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseEqO6AITgPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseEqO6AITgPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseEqO6AITgPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseEqO6AITgPcapsResponseFullFilterV1]
type pcapGetResponseEqO6AITgPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseEqO6AITgPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseEqO6AITgPcapsResponseFullStatus string

const (
	PcapGetResponseEqO6AITgPcapsResponseFullStatusUnknown           PcapGetResponseEqO6AITgPcapsResponseFullStatus = "unknown"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusSuccess           PcapGetResponseEqO6AITgPcapsResponseFullStatus = "success"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusPending           PcapGetResponseEqO6AITgPcapsResponseFullStatus = "pending"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusRunning           PcapGetResponseEqO6AITgPcapsResponseFullStatus = "running"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusConversionPending PcapGetResponseEqO6AITgPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusConversionRunning PcapGetResponseEqO6AITgPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusComplete          PcapGetResponseEqO6AITgPcapsResponseFullStatus = "complete"
	PcapGetResponseEqO6AITgPcapsResponseFullStatusFailed            PcapGetResponseEqO6AITgPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseEqO6AITgPcapsResponseFullSystem string

const (
	PcapGetResponseEqO6AITgPcapsResponseFullSystemMagicTransit PcapGetResponseEqO6AITgPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseEqO6AITgPcapsResponseFullType string

const (
	PcapGetResponseEqO6AITgPcapsResponseFullTypeSimple PcapGetResponseEqO6AITgPcapsResponseFullType = "simple"
	PcapGetResponseEqO6AITgPcapsResponseFullTypeFull   PcapGetResponseEqO6AITgPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsEqO6AITgPcapsRequestSimple],
// [PcapNewParamsEqO6AITgPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsEqO6AITgPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsEqO6AITgPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsEqO6AITgPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsEqO6AITgPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsEqO6AITgPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsEqO6AITgPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsEqO6AITgPcapsRequestSimpleSystem string

const (
	PcapNewParamsEqO6AITgPcapsRequestSimpleSystemMagicTransit PcapNewParamsEqO6AITgPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsEqO6AITgPcapsRequestSimpleType string

const (
	PcapNewParamsEqO6AITgPcapsRequestSimpleTypeSimple PcapNewParamsEqO6AITgPcapsRequestSimpleType = "simple"
	PcapNewParamsEqO6AITgPcapsRequestSimpleTypeFull   PcapNewParamsEqO6AITgPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsEqO6AITgPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsEqO6AITgPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsEqO6AITgPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsEqO6AITgPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsEqO6AITgPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsEqO6AITgPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsEqO6AITgPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsEqO6AITgPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsEqO6AITgPcapsRequestFullSystem string

const (
	PcapNewParamsEqO6AITgPcapsRequestFullSystemMagicTransit PcapNewParamsEqO6AITgPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsEqO6AITgPcapsRequestFullType string

const (
	PcapNewParamsEqO6AITgPcapsRequestFullTypeSimple PcapNewParamsEqO6AITgPcapsRequestFullType = "simple"
	PcapNewParamsEqO6AITgPcapsRequestFullTypeFull   PcapNewParamsEqO6AITgPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsEqO6AITgPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsEqO6AITgPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
