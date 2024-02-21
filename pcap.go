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

// Union satisfied by [PcapNewResponse7WjqRy2XPcapsResponseSimple] or
// [PcapNewResponse7WjqRy2XPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse7WjqRy2XPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse7WjqRy2XPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse7WjqRy2XPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse7WjqRy2XPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse7WjqRy2XPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse7WjqRy2XPcapsResponseSimple]
type pcapNewResponse7WjqRy2XPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse7WjqRy2XPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse7WjqRy2XPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1]
type pcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse7WjqRy2XPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus string

const (
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusUnknown           PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusSuccess           PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "success"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusPending           PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "pending"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusRunning           PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "running"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusConversionPending PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusConversionRunning PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusComplete          PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "complete"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleStatusFailed            PcapNewResponse7WjqRy2XPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse7WjqRy2XPcapsResponseSimpleSystem string

const (
	PcapNewResponse7WjqRy2XPcapsResponseSimpleSystemMagicTransit PcapNewResponse7WjqRy2XPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse7WjqRy2XPcapsResponseSimpleType string

const (
	PcapNewResponse7WjqRy2XPcapsResponseSimpleTypeSimple PcapNewResponse7WjqRy2XPcapsResponseSimpleType = "simple"
	PcapNewResponse7WjqRy2XPcapsResponseSimpleTypeFull   PcapNewResponse7WjqRy2XPcapsResponseSimpleType = "full"
)

type PcapNewResponse7WjqRy2XPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse7WjqRy2XPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse7WjqRy2XPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse7WjqRy2XPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse7WjqRy2XPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse7WjqRy2XPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse7WjqRy2XPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse7WjqRy2XPcapsResponseFull]
type pcapNewResponse7WjqRy2XPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse7WjqRy2XPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse7WjqRy2XPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse7WjqRy2XPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse7WjqRy2XPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse7WjqRy2XPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse7WjqRy2XPcapsResponseFullFilterV1]
type pcapNewResponse7WjqRy2XPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse7WjqRy2XPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse7WjqRy2XPcapsResponseFullStatus string

const (
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusUnknown           PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "unknown"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusSuccess           PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "success"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusPending           PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "pending"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusRunning           PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "running"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusConversionPending PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusConversionRunning PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusComplete          PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "complete"
	PcapNewResponse7WjqRy2XPcapsResponseFullStatusFailed            PcapNewResponse7WjqRy2XPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse7WjqRy2XPcapsResponseFullSystem string

const (
	PcapNewResponse7WjqRy2XPcapsResponseFullSystemMagicTransit PcapNewResponse7WjqRy2XPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse7WjqRy2XPcapsResponseFullType string

const (
	PcapNewResponse7WjqRy2XPcapsResponseFullTypeSimple PcapNewResponse7WjqRy2XPcapsResponseFullType = "simple"
	PcapNewResponse7WjqRy2XPcapsResponseFullTypeFull   PcapNewResponse7WjqRy2XPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse7WjqRy2XPcapsResponseSimple] or
// [PcapListResponse7WjqRy2XPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse7WjqRy2XPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse7WjqRy2XPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse7WjqRy2XPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse7WjqRy2XPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse7WjqRy2XPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse7WjqRy2XPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse7WjqRy2XPcapsResponseSimple]
type pcapListResponse7WjqRy2XPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse7WjqRy2XPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse7WjqRy2XPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1]
type pcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse7WjqRy2XPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse7WjqRy2XPcapsResponseSimpleStatus string

const (
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusUnknown           PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "unknown"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusSuccess           PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "success"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusPending           PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "pending"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusRunning           PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "running"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusConversionPending PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusConversionRunning PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusComplete          PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "complete"
	PcapListResponse7WjqRy2XPcapsResponseSimpleStatusFailed            PcapListResponse7WjqRy2XPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse7WjqRy2XPcapsResponseSimpleSystem string

const (
	PcapListResponse7WjqRy2XPcapsResponseSimpleSystemMagicTransit PcapListResponse7WjqRy2XPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse7WjqRy2XPcapsResponseSimpleType string

const (
	PcapListResponse7WjqRy2XPcapsResponseSimpleTypeSimple PcapListResponse7WjqRy2XPcapsResponseSimpleType = "simple"
	PcapListResponse7WjqRy2XPcapsResponseSimpleTypeFull   PcapListResponse7WjqRy2XPcapsResponseSimpleType = "full"
)

type PcapListResponse7WjqRy2XPcapsResponseFull struct {
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
	FilterV1 PcapListResponse7WjqRy2XPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse7WjqRy2XPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse7WjqRy2XPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse7WjqRy2XPcapsResponseFullType `json:"type"`
	JSON pcapListResponse7WjqRy2XPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse7WjqRy2XPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse7WjqRy2XPcapsResponseFull]
type pcapListResponse7WjqRy2XPcapsResponseFullJSON struct {
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

func (r *PcapListResponse7WjqRy2XPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse7WjqRy2XPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse7WjqRy2XPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse7WjqRy2XPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse7WjqRy2XPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse7WjqRy2XPcapsResponseFullFilterV1]
type pcapListResponse7WjqRy2XPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse7WjqRy2XPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse7WjqRy2XPcapsResponseFullStatus string

const (
	PcapListResponse7WjqRy2XPcapsResponseFullStatusUnknown           PcapListResponse7WjqRy2XPcapsResponseFullStatus = "unknown"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusSuccess           PcapListResponse7WjqRy2XPcapsResponseFullStatus = "success"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusPending           PcapListResponse7WjqRy2XPcapsResponseFullStatus = "pending"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusRunning           PcapListResponse7WjqRy2XPcapsResponseFullStatus = "running"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusConversionPending PcapListResponse7WjqRy2XPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusConversionRunning PcapListResponse7WjqRy2XPcapsResponseFullStatus = "conversion_running"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusComplete          PcapListResponse7WjqRy2XPcapsResponseFullStatus = "complete"
	PcapListResponse7WjqRy2XPcapsResponseFullStatusFailed            PcapListResponse7WjqRy2XPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse7WjqRy2XPcapsResponseFullSystem string

const (
	PcapListResponse7WjqRy2XPcapsResponseFullSystemMagicTransit PcapListResponse7WjqRy2XPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse7WjqRy2XPcapsResponseFullType string

const (
	PcapListResponse7WjqRy2XPcapsResponseFullTypeSimple PcapListResponse7WjqRy2XPcapsResponseFullType = "simple"
	PcapListResponse7WjqRy2XPcapsResponseFullTypeFull   PcapListResponse7WjqRy2XPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse7WjqRy2XPcapsResponseSimple] or
// [PcapGetResponse7WjqRy2XPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse7WjqRy2XPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7WjqRy2XPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7WjqRy2XPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse7WjqRy2XPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse7WjqRy2XPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse7WjqRy2XPcapsResponseSimple]
type pcapGetResponse7WjqRy2XPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse7WjqRy2XPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7WjqRy2XPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1]
type pcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7WjqRy2XPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus string

const (
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusUnknown           PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusSuccess           PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "success"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusPending           PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "pending"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusRunning           PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "running"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusConversionPending PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusConversionRunning PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusComplete          PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "complete"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleStatusFailed            PcapGetResponse7WjqRy2XPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7WjqRy2XPcapsResponseSimpleSystem string

const (
	PcapGetResponse7WjqRy2XPcapsResponseSimpleSystemMagicTransit PcapGetResponse7WjqRy2XPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7WjqRy2XPcapsResponseSimpleType string

const (
	PcapGetResponse7WjqRy2XPcapsResponseSimpleTypeSimple PcapGetResponse7WjqRy2XPcapsResponseSimpleType = "simple"
	PcapGetResponse7WjqRy2XPcapsResponseSimpleTypeFull   PcapGetResponse7WjqRy2XPcapsResponseSimpleType = "full"
)

type PcapGetResponse7WjqRy2XPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse7WjqRy2XPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7WjqRy2XPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7WjqRy2XPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7WjqRy2XPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse7WjqRy2XPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse7WjqRy2XPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse7WjqRy2XPcapsResponseFull]
type pcapGetResponse7WjqRy2XPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse7WjqRy2XPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7WjqRy2XPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7WjqRy2XPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse7WjqRy2XPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse7WjqRy2XPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse7WjqRy2XPcapsResponseFullFilterV1]
type pcapGetResponse7WjqRy2XPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7WjqRy2XPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7WjqRy2XPcapsResponseFullStatus string

const (
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusUnknown           PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "unknown"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusSuccess           PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "success"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusPending           PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "pending"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusRunning           PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "running"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusConversionPending PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusConversionRunning PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusComplete          PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "complete"
	PcapGetResponse7WjqRy2XPcapsResponseFullStatusFailed            PcapGetResponse7WjqRy2XPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7WjqRy2XPcapsResponseFullSystem string

const (
	PcapGetResponse7WjqRy2XPcapsResponseFullSystemMagicTransit PcapGetResponse7WjqRy2XPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7WjqRy2XPcapsResponseFullType string

const (
	PcapGetResponse7WjqRy2XPcapsResponseFullTypeSimple PcapGetResponse7WjqRy2XPcapsResponseFullType = "simple"
	PcapGetResponse7WjqRy2XPcapsResponseFullTypeFull   PcapGetResponse7WjqRy2XPcapsResponseFullType = "full"
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
