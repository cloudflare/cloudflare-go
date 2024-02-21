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

// Union satisfied by [PcapNewResponseGne8pwAwPcapsResponseSimple] or
// [PcapNewResponseGne8pwAwPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseGne8pwAwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseGne8pwAwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseGne8pwAwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseGne8pwAwPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseGne8pwAwPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseGne8pwAwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseGne8pwAwPcapsResponseSimple]
type pcapNewResponseGne8pwAwPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseGne8pwAwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseGne8pwAwPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1]
type pcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseGne8pwAwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseGne8pwAwPcapsResponseSimpleStatus string

const (
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusUnknown           PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusSuccess           PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "success"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusPending           PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "pending"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusRunning           PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "running"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusConversionPending PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusConversionRunning PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusComplete          PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "complete"
	PcapNewResponseGne8pwAwPcapsResponseSimpleStatusFailed            PcapNewResponseGne8pwAwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseGne8pwAwPcapsResponseSimpleSystem string

const (
	PcapNewResponseGne8pwAwPcapsResponseSimpleSystemMagicTransit PcapNewResponseGne8pwAwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseGne8pwAwPcapsResponseSimpleType string

const (
	PcapNewResponseGne8pwAwPcapsResponseSimpleTypeSimple PcapNewResponseGne8pwAwPcapsResponseSimpleType = "simple"
	PcapNewResponseGne8pwAwPcapsResponseSimpleTypeFull   PcapNewResponseGne8pwAwPcapsResponseSimpleType = "full"
)

type PcapNewResponseGne8pwAwPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseGne8pwAwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseGne8pwAwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseGne8pwAwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseGne8pwAwPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseGne8pwAwPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseGne8pwAwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseGne8pwAwPcapsResponseFull]
type pcapNewResponseGne8pwAwPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseGne8pwAwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseGne8pwAwPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseGne8pwAwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseGne8pwAwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseGne8pwAwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseGne8pwAwPcapsResponseFullFilterV1]
type pcapNewResponseGne8pwAwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseGne8pwAwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseGne8pwAwPcapsResponseFullStatus string

const (
	PcapNewResponseGne8pwAwPcapsResponseFullStatusUnknown           PcapNewResponseGne8pwAwPcapsResponseFullStatus = "unknown"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusSuccess           PcapNewResponseGne8pwAwPcapsResponseFullStatus = "success"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusPending           PcapNewResponseGne8pwAwPcapsResponseFullStatus = "pending"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusRunning           PcapNewResponseGne8pwAwPcapsResponseFullStatus = "running"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusConversionPending PcapNewResponseGne8pwAwPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusConversionRunning PcapNewResponseGne8pwAwPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusComplete          PcapNewResponseGne8pwAwPcapsResponseFullStatus = "complete"
	PcapNewResponseGne8pwAwPcapsResponseFullStatusFailed            PcapNewResponseGne8pwAwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseGne8pwAwPcapsResponseFullSystem string

const (
	PcapNewResponseGne8pwAwPcapsResponseFullSystemMagicTransit PcapNewResponseGne8pwAwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseGne8pwAwPcapsResponseFullType string

const (
	PcapNewResponseGne8pwAwPcapsResponseFullTypeSimple PcapNewResponseGne8pwAwPcapsResponseFullType = "simple"
	PcapNewResponseGne8pwAwPcapsResponseFullTypeFull   PcapNewResponseGne8pwAwPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseGne8pwAwPcapsResponseSimple] or
// [PcapListResponseGne8pwAwPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseGne8pwAwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseGne8pwAwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseGne8pwAwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseGne8pwAwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseGne8pwAwPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseGne8pwAwPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseGne8pwAwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseGne8pwAwPcapsResponseSimple]
type pcapListResponseGne8pwAwPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseGne8pwAwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseGne8pwAwPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseGne8pwAwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseGne8pwAwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseGne8pwAwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseGne8pwAwPcapsResponseSimpleFilterV1]
type pcapListResponseGne8pwAwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseGne8pwAwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseGne8pwAwPcapsResponseSimpleStatus string

const (
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusUnknown           PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "unknown"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusSuccess           PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "success"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusPending           PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "pending"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusRunning           PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "running"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusConversionPending PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusConversionRunning PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusComplete          PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "complete"
	PcapListResponseGne8pwAwPcapsResponseSimpleStatusFailed            PcapListResponseGne8pwAwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseGne8pwAwPcapsResponseSimpleSystem string

const (
	PcapListResponseGne8pwAwPcapsResponseSimpleSystemMagicTransit PcapListResponseGne8pwAwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseGne8pwAwPcapsResponseSimpleType string

const (
	PcapListResponseGne8pwAwPcapsResponseSimpleTypeSimple PcapListResponseGne8pwAwPcapsResponseSimpleType = "simple"
	PcapListResponseGne8pwAwPcapsResponseSimpleTypeFull   PcapListResponseGne8pwAwPcapsResponseSimpleType = "full"
)

type PcapListResponseGne8pwAwPcapsResponseFull struct {
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
	FilterV1 PcapListResponseGne8pwAwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseGne8pwAwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseGne8pwAwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseGne8pwAwPcapsResponseFullType `json:"type"`
	JSON pcapListResponseGne8pwAwPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseGne8pwAwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseGne8pwAwPcapsResponseFull]
type pcapListResponseGne8pwAwPcapsResponseFullJSON struct {
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

func (r *PcapListResponseGne8pwAwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseGne8pwAwPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseGne8pwAwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseGne8pwAwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseGne8pwAwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseGne8pwAwPcapsResponseFullFilterV1]
type pcapListResponseGne8pwAwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseGne8pwAwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseGne8pwAwPcapsResponseFullStatus string

const (
	PcapListResponseGne8pwAwPcapsResponseFullStatusUnknown           PcapListResponseGne8pwAwPcapsResponseFullStatus = "unknown"
	PcapListResponseGne8pwAwPcapsResponseFullStatusSuccess           PcapListResponseGne8pwAwPcapsResponseFullStatus = "success"
	PcapListResponseGne8pwAwPcapsResponseFullStatusPending           PcapListResponseGne8pwAwPcapsResponseFullStatus = "pending"
	PcapListResponseGne8pwAwPcapsResponseFullStatusRunning           PcapListResponseGne8pwAwPcapsResponseFullStatus = "running"
	PcapListResponseGne8pwAwPcapsResponseFullStatusConversionPending PcapListResponseGne8pwAwPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseGne8pwAwPcapsResponseFullStatusConversionRunning PcapListResponseGne8pwAwPcapsResponseFullStatus = "conversion_running"
	PcapListResponseGne8pwAwPcapsResponseFullStatusComplete          PcapListResponseGne8pwAwPcapsResponseFullStatus = "complete"
	PcapListResponseGne8pwAwPcapsResponseFullStatusFailed            PcapListResponseGne8pwAwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseGne8pwAwPcapsResponseFullSystem string

const (
	PcapListResponseGne8pwAwPcapsResponseFullSystemMagicTransit PcapListResponseGne8pwAwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseGne8pwAwPcapsResponseFullType string

const (
	PcapListResponseGne8pwAwPcapsResponseFullTypeSimple PcapListResponseGne8pwAwPcapsResponseFullType = "simple"
	PcapListResponseGne8pwAwPcapsResponseFullTypeFull   PcapListResponseGne8pwAwPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseGne8pwAwPcapsResponseSimple] or
// [PcapGetResponseGne8pwAwPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseGne8pwAwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseGne8pwAwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseGne8pwAwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseGne8pwAwPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseGne8pwAwPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseGne8pwAwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseGne8pwAwPcapsResponseSimple]
type pcapGetResponseGne8pwAwPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseGne8pwAwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseGne8pwAwPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1]
type pcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseGne8pwAwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseGne8pwAwPcapsResponseSimpleStatus string

const (
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusUnknown           PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusSuccess           PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "success"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusPending           PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "pending"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusRunning           PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "running"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusConversionPending PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusConversionRunning PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusComplete          PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "complete"
	PcapGetResponseGne8pwAwPcapsResponseSimpleStatusFailed            PcapGetResponseGne8pwAwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseGne8pwAwPcapsResponseSimpleSystem string

const (
	PcapGetResponseGne8pwAwPcapsResponseSimpleSystemMagicTransit PcapGetResponseGne8pwAwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseGne8pwAwPcapsResponseSimpleType string

const (
	PcapGetResponseGne8pwAwPcapsResponseSimpleTypeSimple PcapGetResponseGne8pwAwPcapsResponseSimpleType = "simple"
	PcapGetResponseGne8pwAwPcapsResponseSimpleTypeFull   PcapGetResponseGne8pwAwPcapsResponseSimpleType = "full"
)

type PcapGetResponseGne8pwAwPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseGne8pwAwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseGne8pwAwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseGne8pwAwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseGne8pwAwPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseGne8pwAwPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseGne8pwAwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseGne8pwAwPcapsResponseFull]
type pcapGetResponseGne8pwAwPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseGne8pwAwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseGne8pwAwPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseGne8pwAwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseGne8pwAwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseGne8pwAwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseGne8pwAwPcapsResponseFullFilterV1]
type pcapGetResponseGne8pwAwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseGne8pwAwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseGne8pwAwPcapsResponseFullStatus string

const (
	PcapGetResponseGne8pwAwPcapsResponseFullStatusUnknown           PcapGetResponseGne8pwAwPcapsResponseFullStatus = "unknown"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusSuccess           PcapGetResponseGne8pwAwPcapsResponseFullStatus = "success"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusPending           PcapGetResponseGne8pwAwPcapsResponseFullStatus = "pending"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusRunning           PcapGetResponseGne8pwAwPcapsResponseFullStatus = "running"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusConversionPending PcapGetResponseGne8pwAwPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusConversionRunning PcapGetResponseGne8pwAwPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusComplete          PcapGetResponseGne8pwAwPcapsResponseFullStatus = "complete"
	PcapGetResponseGne8pwAwPcapsResponseFullStatusFailed            PcapGetResponseGne8pwAwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseGne8pwAwPcapsResponseFullSystem string

const (
	PcapGetResponseGne8pwAwPcapsResponseFullSystemMagicTransit PcapGetResponseGne8pwAwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseGne8pwAwPcapsResponseFullType string

const (
	PcapGetResponseGne8pwAwPcapsResponseFullTypeSimple PcapGetResponseGne8pwAwPcapsResponseFullType = "simple"
	PcapGetResponseGne8pwAwPcapsResponseFullTypeFull   PcapGetResponseGne8pwAwPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsGne8pwAwPcapsRequestSimple],
// [PcapNewParamsGne8pwAwPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsGne8pwAwPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsGne8pwAwPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsGne8pwAwPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsGne8pwAwPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsGne8pwAwPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsGne8pwAwPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsGne8pwAwPcapsRequestSimpleSystem string

const (
	PcapNewParamsGne8pwAwPcapsRequestSimpleSystemMagicTransit PcapNewParamsGne8pwAwPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsGne8pwAwPcapsRequestSimpleType string

const (
	PcapNewParamsGne8pwAwPcapsRequestSimpleTypeSimple PcapNewParamsGne8pwAwPcapsRequestSimpleType = "simple"
	PcapNewParamsGne8pwAwPcapsRequestSimpleTypeFull   PcapNewParamsGne8pwAwPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsGne8pwAwPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsGne8pwAwPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsGne8pwAwPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsGne8pwAwPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsGne8pwAwPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsGne8pwAwPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsGne8pwAwPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsGne8pwAwPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsGne8pwAwPcapsRequestFullSystem string

const (
	PcapNewParamsGne8pwAwPcapsRequestFullSystemMagicTransit PcapNewParamsGne8pwAwPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsGne8pwAwPcapsRequestFullType string

const (
	PcapNewParamsGne8pwAwPcapsRequestFullTypeSimple PcapNewParamsGne8pwAwPcapsRequestFullType = "simple"
	PcapNewParamsGne8pwAwPcapsRequestFullTypeFull   PcapNewParamsGne8pwAwPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsGne8pwAwPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsGne8pwAwPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
