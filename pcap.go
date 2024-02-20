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

// Union satisfied by [PcapNewResponseI4mKBGKnPcapsResponseSimple] or
// [PcapNewResponseI4mKBGKnPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseI4mKBGKnPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseI4mKBGKnPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseI4mKBGKnPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseI4mKBGKnPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseI4mKbgKnPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseI4mKbgKnPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseI4mKBGKnPcapsResponseSimple]
type pcapNewResponseI4mKbgKnPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseI4mKBGKnPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseI4mKBGKnPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseI4mKBGKnPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseI4mKBGKnPcapsResponseSimpleFilterV1]
type pcapNewResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseI4mKBGKnPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus string

const (
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusUnknown           PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusSuccess           PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "success"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusPending           PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "pending"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusRunning           PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "running"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusConversionPending PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusConversionRunning PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusComplete          PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "complete"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleStatusFailed            PcapNewResponseI4mKBGKnPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseI4mKBGKnPcapsResponseSimpleSystem string

const (
	PcapNewResponseI4mKBGKnPcapsResponseSimpleSystemMagicTransit PcapNewResponseI4mKBGKnPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseI4mKBGKnPcapsResponseSimpleType string

const (
	PcapNewResponseI4mKBGKnPcapsResponseSimpleTypeSimple PcapNewResponseI4mKBGKnPcapsResponseSimpleType = "simple"
	PcapNewResponseI4mKBGKnPcapsResponseSimpleTypeFull   PcapNewResponseI4mKBGKnPcapsResponseSimpleType = "full"
)

type PcapNewResponseI4mKBGKnPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseI4mKBGKnPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseI4mKBGKnPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseI4mKBGKnPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseI4mKBGKnPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseI4mKbgKnPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseI4mKbgKnPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseI4mKBGKnPcapsResponseFull]
type pcapNewResponseI4mKbgKnPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseI4mKBGKnPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseI4mKBGKnPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseI4mKBGKnPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseI4mKbgKnPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseI4mKbgKnPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseI4mKBGKnPcapsResponseFullFilterV1]
type pcapNewResponseI4mKbgKnPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseI4mKBGKnPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseI4mKBGKnPcapsResponseFullStatus string

const (
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusUnknown           PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "unknown"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusSuccess           PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "success"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusPending           PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "pending"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusRunning           PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "running"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusConversionPending PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusConversionRunning PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusComplete          PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "complete"
	PcapNewResponseI4mKBGKnPcapsResponseFullStatusFailed            PcapNewResponseI4mKBGKnPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseI4mKBGKnPcapsResponseFullSystem string

const (
	PcapNewResponseI4mKBGKnPcapsResponseFullSystemMagicTransit PcapNewResponseI4mKBGKnPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseI4mKBGKnPcapsResponseFullType string

const (
	PcapNewResponseI4mKBGKnPcapsResponseFullTypeSimple PcapNewResponseI4mKBGKnPcapsResponseFullType = "simple"
	PcapNewResponseI4mKBGKnPcapsResponseFullTypeFull   PcapNewResponseI4mKBGKnPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseI4mKBGKnPcapsResponseSimple] or
// [PcapListResponseI4mKBGKnPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseI4mKBGKnPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseI4mKBGKnPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseI4mKBGKnPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseI4mKBGKnPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseI4mKBGKnPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseI4mKbgKnPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseI4mKbgKnPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseI4mKBGKnPcapsResponseSimple]
type pcapListResponseI4mKbgKnPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseI4mKBGKnPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseI4mKBGKnPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseI4mKBGKnPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseI4mKBGKnPcapsResponseSimpleFilterV1]
type pcapListResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseI4mKBGKnPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseI4mKBGKnPcapsResponseSimpleStatus string

const (
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusUnknown           PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "unknown"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusSuccess           PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "success"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusPending           PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "pending"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusRunning           PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "running"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusConversionPending PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusConversionRunning PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusComplete          PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "complete"
	PcapListResponseI4mKBGKnPcapsResponseSimpleStatusFailed            PcapListResponseI4mKBGKnPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseI4mKBGKnPcapsResponseSimpleSystem string

const (
	PcapListResponseI4mKBGKnPcapsResponseSimpleSystemMagicTransit PcapListResponseI4mKBGKnPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseI4mKBGKnPcapsResponseSimpleType string

const (
	PcapListResponseI4mKBGKnPcapsResponseSimpleTypeSimple PcapListResponseI4mKBGKnPcapsResponseSimpleType = "simple"
	PcapListResponseI4mKBGKnPcapsResponseSimpleTypeFull   PcapListResponseI4mKBGKnPcapsResponseSimpleType = "full"
)

type PcapListResponseI4mKBGKnPcapsResponseFull struct {
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
	FilterV1 PcapListResponseI4mKBGKnPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseI4mKBGKnPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseI4mKBGKnPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseI4mKBGKnPcapsResponseFullType `json:"type"`
	JSON pcapListResponseI4mKbgKnPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseI4mKbgKnPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseI4mKBGKnPcapsResponseFull]
type pcapListResponseI4mKbgKnPcapsResponseFullJSON struct {
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

func (r *PcapListResponseI4mKBGKnPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseI4mKBGKnPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseI4mKBGKnPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseI4mKbgKnPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseI4mKbgKnPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseI4mKBGKnPcapsResponseFullFilterV1]
type pcapListResponseI4mKbgKnPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseI4mKBGKnPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseI4mKBGKnPcapsResponseFullStatus string

const (
	PcapListResponseI4mKBGKnPcapsResponseFullStatusUnknown           PcapListResponseI4mKBGKnPcapsResponseFullStatus = "unknown"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusSuccess           PcapListResponseI4mKBGKnPcapsResponseFullStatus = "success"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusPending           PcapListResponseI4mKBGKnPcapsResponseFullStatus = "pending"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusRunning           PcapListResponseI4mKBGKnPcapsResponseFullStatus = "running"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusConversionPending PcapListResponseI4mKBGKnPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusConversionRunning PcapListResponseI4mKBGKnPcapsResponseFullStatus = "conversion_running"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusComplete          PcapListResponseI4mKBGKnPcapsResponseFullStatus = "complete"
	PcapListResponseI4mKBGKnPcapsResponseFullStatusFailed            PcapListResponseI4mKBGKnPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseI4mKBGKnPcapsResponseFullSystem string

const (
	PcapListResponseI4mKBGKnPcapsResponseFullSystemMagicTransit PcapListResponseI4mKBGKnPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseI4mKBGKnPcapsResponseFullType string

const (
	PcapListResponseI4mKBGKnPcapsResponseFullTypeSimple PcapListResponseI4mKBGKnPcapsResponseFullType = "simple"
	PcapListResponseI4mKBGKnPcapsResponseFullTypeFull   PcapListResponseI4mKBGKnPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseI4mKBGKnPcapsResponseSimple] or
// [PcapGetResponseI4mKBGKnPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseI4mKBGKnPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseI4mKBGKnPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseI4mKBGKnPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseI4mKBGKnPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseI4mKbgKnPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseI4mKbgKnPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseI4mKBGKnPcapsResponseSimple]
type pcapGetResponseI4mKbgKnPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseI4mKBGKnPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseI4mKBGKnPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseI4mKBGKnPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseI4mKBGKnPcapsResponseSimpleFilterV1]
type pcapGetResponseI4mKbgKnPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseI4mKBGKnPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus string

const (
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusUnknown           PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusSuccess           PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "success"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusPending           PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "pending"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusRunning           PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "running"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusConversionPending PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusConversionRunning PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusComplete          PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "complete"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleStatusFailed            PcapGetResponseI4mKBGKnPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseI4mKBGKnPcapsResponseSimpleSystem string

const (
	PcapGetResponseI4mKBGKnPcapsResponseSimpleSystemMagicTransit PcapGetResponseI4mKBGKnPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseI4mKBGKnPcapsResponseSimpleType string

const (
	PcapGetResponseI4mKBGKnPcapsResponseSimpleTypeSimple PcapGetResponseI4mKBGKnPcapsResponseSimpleType = "simple"
	PcapGetResponseI4mKBGKnPcapsResponseSimpleTypeFull   PcapGetResponseI4mKBGKnPcapsResponseSimpleType = "full"
)

type PcapGetResponseI4mKBGKnPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseI4mKBGKnPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseI4mKBGKnPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseI4mKBGKnPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseI4mKBGKnPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseI4mKbgKnPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseI4mKbgKnPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseI4mKBGKnPcapsResponseFull]
type pcapGetResponseI4mKbgKnPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseI4mKBGKnPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseI4mKBGKnPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseI4mKBGKnPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseI4mKbgKnPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseI4mKbgKnPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseI4mKBGKnPcapsResponseFullFilterV1]
type pcapGetResponseI4mKbgKnPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseI4mKBGKnPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseI4mKBGKnPcapsResponseFullStatus string

const (
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusUnknown           PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "unknown"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusSuccess           PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "success"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusPending           PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "pending"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusRunning           PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "running"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusConversionPending PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusConversionRunning PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusComplete          PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "complete"
	PcapGetResponseI4mKBGKnPcapsResponseFullStatusFailed            PcapGetResponseI4mKBGKnPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseI4mKBGKnPcapsResponseFullSystem string

const (
	PcapGetResponseI4mKBGKnPcapsResponseFullSystemMagicTransit PcapGetResponseI4mKBGKnPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseI4mKBGKnPcapsResponseFullType string

const (
	PcapGetResponseI4mKBGKnPcapsResponseFullTypeSimple PcapGetResponseI4mKBGKnPcapsResponseFullType = "simple"
	PcapGetResponseI4mKBGKnPcapsResponseFullTypeFull   PcapGetResponseI4mKBGKnPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsI4mKBGKnPcapsRequestSimple],
// [PcapNewParamsI4mKBGKnPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsI4mKBGKnPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsI4mKbgKnPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsI4mKbgKnPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsI4mKbgKnPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsI4mKBGKnPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsI4mKBGKnPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsI4mKbgKnPcapsRequestSimpleSystem string

const (
	PcapNewParamsI4mKbgKnPcapsRequestSimpleSystemMagicTransit PcapNewParamsI4mKbgKnPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsI4mKbgKnPcapsRequestSimpleType string

const (
	PcapNewParamsI4mKbgKnPcapsRequestSimpleTypeSimple PcapNewParamsI4mKbgKnPcapsRequestSimpleType = "simple"
	PcapNewParamsI4mKbgKnPcapsRequestSimpleTypeFull   PcapNewParamsI4mKbgKnPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsI4mKbgKnPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsI4mKbgKnPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsI4mKBGKnPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsI4mKbgKnPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsI4mKbgKnPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsI4mKbgKnPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsI4mKBGKnPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsI4mKBGKnPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsI4mKbgKnPcapsRequestFullSystem string

const (
	PcapNewParamsI4mKbgKnPcapsRequestFullSystemMagicTransit PcapNewParamsI4mKbgKnPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsI4mKbgKnPcapsRequestFullType string

const (
	PcapNewParamsI4mKbgKnPcapsRequestFullTypeSimple PcapNewParamsI4mKbgKnPcapsRequestFullType = "simple"
	PcapNewParamsI4mKbgKnPcapsRequestFullTypeFull   PcapNewParamsI4mKbgKnPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsI4mKbgKnPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsI4mKbgKnPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
