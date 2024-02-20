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

// Union satisfied by [PcapNewResponseEErxb5cBPcapsResponseSimple] or
// [PcapNewResponseEErxb5cBPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseEErxb5cBPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseEErxb5cBPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseEErxb5cBPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseEErxb5cBPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseEErxb5cBPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseEErxb5cBPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseEErxb5cBPcapsResponseSimple]
type pcapNewResponseEErxb5cBPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseEErxb5cBPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseEErxb5cBPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1]
type pcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseEErxb5cBPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseEErxb5cBPcapsResponseSimpleStatus string

const (
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusUnknown           PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusSuccess           PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "success"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusPending           PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "pending"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusRunning           PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "running"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusConversionPending PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusConversionRunning PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusComplete          PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "complete"
	PcapNewResponseEErxb5cBPcapsResponseSimpleStatusFailed            PcapNewResponseEErxb5cBPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseEErxb5cBPcapsResponseSimpleSystem string

const (
	PcapNewResponseEErxb5cBPcapsResponseSimpleSystemMagicTransit PcapNewResponseEErxb5cBPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseEErxb5cBPcapsResponseSimpleType string

const (
	PcapNewResponseEErxb5cBPcapsResponseSimpleTypeSimple PcapNewResponseEErxb5cBPcapsResponseSimpleType = "simple"
	PcapNewResponseEErxb5cBPcapsResponseSimpleTypeFull   PcapNewResponseEErxb5cBPcapsResponseSimpleType = "full"
)

type PcapNewResponseEErxb5cBPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseEErxb5cBPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseEErxb5cBPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseEErxb5cBPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseEErxb5cBPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseEErxb5cBPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseEErxb5cBPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseEErxb5cBPcapsResponseFull]
type pcapNewResponseEErxb5cBPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseEErxb5cBPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseEErxb5cBPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseEErxb5cBPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseEErxb5cBPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseEErxb5cBPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseEErxb5cBPcapsResponseFullFilterV1]
type pcapNewResponseEErxb5cBPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseEErxb5cBPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseEErxb5cBPcapsResponseFullStatus string

const (
	PcapNewResponseEErxb5cBPcapsResponseFullStatusUnknown           PcapNewResponseEErxb5cBPcapsResponseFullStatus = "unknown"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusSuccess           PcapNewResponseEErxb5cBPcapsResponseFullStatus = "success"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusPending           PcapNewResponseEErxb5cBPcapsResponseFullStatus = "pending"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusRunning           PcapNewResponseEErxb5cBPcapsResponseFullStatus = "running"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusConversionPending PcapNewResponseEErxb5cBPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusConversionRunning PcapNewResponseEErxb5cBPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusComplete          PcapNewResponseEErxb5cBPcapsResponseFullStatus = "complete"
	PcapNewResponseEErxb5cBPcapsResponseFullStatusFailed            PcapNewResponseEErxb5cBPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseEErxb5cBPcapsResponseFullSystem string

const (
	PcapNewResponseEErxb5cBPcapsResponseFullSystemMagicTransit PcapNewResponseEErxb5cBPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseEErxb5cBPcapsResponseFullType string

const (
	PcapNewResponseEErxb5cBPcapsResponseFullTypeSimple PcapNewResponseEErxb5cBPcapsResponseFullType = "simple"
	PcapNewResponseEErxb5cBPcapsResponseFullTypeFull   PcapNewResponseEErxb5cBPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseEErxb5cBPcapsResponseSimple] or
// [PcapListResponseEErxb5cBPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseEErxb5cBPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseEErxb5cBPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseEErxb5cBPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseEErxb5cBPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseEErxb5cBPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseEErxb5cBPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseEErxb5cBPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseEErxb5cBPcapsResponseSimple]
type pcapListResponseEErxb5cBPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseEErxb5cBPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseEErxb5cBPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseEErxb5cBPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseEErxb5cBPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseEErxb5cBPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseEErxb5cBPcapsResponseSimpleFilterV1]
type pcapListResponseEErxb5cBPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseEErxb5cBPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseEErxb5cBPcapsResponseSimpleStatus string

const (
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusUnknown           PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "unknown"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusSuccess           PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "success"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusPending           PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "pending"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusRunning           PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "running"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusConversionPending PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusConversionRunning PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusComplete          PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "complete"
	PcapListResponseEErxb5cBPcapsResponseSimpleStatusFailed            PcapListResponseEErxb5cBPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseEErxb5cBPcapsResponseSimpleSystem string

const (
	PcapListResponseEErxb5cBPcapsResponseSimpleSystemMagicTransit PcapListResponseEErxb5cBPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseEErxb5cBPcapsResponseSimpleType string

const (
	PcapListResponseEErxb5cBPcapsResponseSimpleTypeSimple PcapListResponseEErxb5cBPcapsResponseSimpleType = "simple"
	PcapListResponseEErxb5cBPcapsResponseSimpleTypeFull   PcapListResponseEErxb5cBPcapsResponseSimpleType = "full"
)

type PcapListResponseEErxb5cBPcapsResponseFull struct {
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
	FilterV1 PcapListResponseEErxb5cBPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseEErxb5cBPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseEErxb5cBPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseEErxb5cBPcapsResponseFullType `json:"type"`
	JSON pcapListResponseEErxb5cBPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseEErxb5cBPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseEErxb5cBPcapsResponseFull]
type pcapListResponseEErxb5cBPcapsResponseFullJSON struct {
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

func (r *PcapListResponseEErxb5cBPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseEErxb5cBPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseEErxb5cBPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseEErxb5cBPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseEErxb5cBPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseEErxb5cBPcapsResponseFullFilterV1]
type pcapListResponseEErxb5cBPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseEErxb5cBPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseEErxb5cBPcapsResponseFullStatus string

const (
	PcapListResponseEErxb5cBPcapsResponseFullStatusUnknown           PcapListResponseEErxb5cBPcapsResponseFullStatus = "unknown"
	PcapListResponseEErxb5cBPcapsResponseFullStatusSuccess           PcapListResponseEErxb5cBPcapsResponseFullStatus = "success"
	PcapListResponseEErxb5cBPcapsResponseFullStatusPending           PcapListResponseEErxb5cBPcapsResponseFullStatus = "pending"
	PcapListResponseEErxb5cBPcapsResponseFullStatusRunning           PcapListResponseEErxb5cBPcapsResponseFullStatus = "running"
	PcapListResponseEErxb5cBPcapsResponseFullStatusConversionPending PcapListResponseEErxb5cBPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseEErxb5cBPcapsResponseFullStatusConversionRunning PcapListResponseEErxb5cBPcapsResponseFullStatus = "conversion_running"
	PcapListResponseEErxb5cBPcapsResponseFullStatusComplete          PcapListResponseEErxb5cBPcapsResponseFullStatus = "complete"
	PcapListResponseEErxb5cBPcapsResponseFullStatusFailed            PcapListResponseEErxb5cBPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseEErxb5cBPcapsResponseFullSystem string

const (
	PcapListResponseEErxb5cBPcapsResponseFullSystemMagicTransit PcapListResponseEErxb5cBPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseEErxb5cBPcapsResponseFullType string

const (
	PcapListResponseEErxb5cBPcapsResponseFullTypeSimple PcapListResponseEErxb5cBPcapsResponseFullType = "simple"
	PcapListResponseEErxb5cBPcapsResponseFullTypeFull   PcapListResponseEErxb5cBPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseEErxb5cBPcapsResponseSimple] or
// [PcapGetResponseEErxb5cBPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseEErxb5cBPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseEErxb5cBPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseEErxb5cBPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseEErxb5cBPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseEErxb5cBPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseEErxb5cBPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseEErxb5cBPcapsResponseSimple]
type pcapGetResponseEErxb5cBPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseEErxb5cBPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseEErxb5cBPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1]
type pcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseEErxb5cBPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseEErxb5cBPcapsResponseSimpleStatus string

const (
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusUnknown           PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusSuccess           PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "success"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusPending           PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "pending"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusRunning           PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "running"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusConversionPending PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusConversionRunning PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusComplete          PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "complete"
	PcapGetResponseEErxb5cBPcapsResponseSimpleStatusFailed            PcapGetResponseEErxb5cBPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseEErxb5cBPcapsResponseSimpleSystem string

const (
	PcapGetResponseEErxb5cBPcapsResponseSimpleSystemMagicTransit PcapGetResponseEErxb5cBPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseEErxb5cBPcapsResponseSimpleType string

const (
	PcapGetResponseEErxb5cBPcapsResponseSimpleTypeSimple PcapGetResponseEErxb5cBPcapsResponseSimpleType = "simple"
	PcapGetResponseEErxb5cBPcapsResponseSimpleTypeFull   PcapGetResponseEErxb5cBPcapsResponseSimpleType = "full"
)

type PcapGetResponseEErxb5cBPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseEErxb5cBPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseEErxb5cBPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseEErxb5cBPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseEErxb5cBPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseEErxb5cBPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseEErxb5cBPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseEErxb5cBPcapsResponseFull]
type pcapGetResponseEErxb5cBPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseEErxb5cBPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseEErxb5cBPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseEErxb5cBPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseEErxb5cBPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseEErxb5cBPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseEErxb5cBPcapsResponseFullFilterV1]
type pcapGetResponseEErxb5cBPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseEErxb5cBPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseEErxb5cBPcapsResponseFullStatus string

const (
	PcapGetResponseEErxb5cBPcapsResponseFullStatusUnknown           PcapGetResponseEErxb5cBPcapsResponseFullStatus = "unknown"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusSuccess           PcapGetResponseEErxb5cBPcapsResponseFullStatus = "success"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusPending           PcapGetResponseEErxb5cBPcapsResponseFullStatus = "pending"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusRunning           PcapGetResponseEErxb5cBPcapsResponseFullStatus = "running"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusConversionPending PcapGetResponseEErxb5cBPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusConversionRunning PcapGetResponseEErxb5cBPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusComplete          PcapGetResponseEErxb5cBPcapsResponseFullStatus = "complete"
	PcapGetResponseEErxb5cBPcapsResponseFullStatusFailed            PcapGetResponseEErxb5cBPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseEErxb5cBPcapsResponseFullSystem string

const (
	PcapGetResponseEErxb5cBPcapsResponseFullSystemMagicTransit PcapGetResponseEErxb5cBPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseEErxb5cBPcapsResponseFullType string

const (
	PcapGetResponseEErxb5cBPcapsResponseFullTypeSimple PcapGetResponseEErxb5cBPcapsResponseFullType = "simple"
	PcapGetResponseEErxb5cBPcapsResponseFullTypeFull   PcapGetResponseEErxb5cBPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsEErxb5cBPcapsRequestSimple],
// [PcapNewParamsEErxb5cBPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsEErxb5cBPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsEErxb5cBPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsEErxb5cBPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsEErxb5cBPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsEErxb5cBPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsEErxb5cBPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsEErxb5cBPcapsRequestSimpleSystem string

const (
	PcapNewParamsEErxb5cBPcapsRequestSimpleSystemMagicTransit PcapNewParamsEErxb5cBPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsEErxb5cBPcapsRequestSimpleType string

const (
	PcapNewParamsEErxb5cBPcapsRequestSimpleTypeSimple PcapNewParamsEErxb5cBPcapsRequestSimpleType = "simple"
	PcapNewParamsEErxb5cBPcapsRequestSimpleTypeFull   PcapNewParamsEErxb5cBPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsEErxb5cBPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsEErxb5cBPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsEErxb5cBPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsEErxb5cBPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsEErxb5cBPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsEErxb5cBPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsEErxb5cBPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsEErxb5cBPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsEErxb5cBPcapsRequestFullSystem string

const (
	PcapNewParamsEErxb5cBPcapsRequestFullSystemMagicTransit PcapNewParamsEErxb5cBPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsEErxb5cBPcapsRequestFullType string

const (
	PcapNewParamsEErxb5cBPcapsRequestFullTypeSimple PcapNewParamsEErxb5cBPcapsRequestFullType = "simple"
	PcapNewParamsEErxb5cBPcapsRequestFullTypeFull   PcapNewParamsEErxb5cBPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsEErxb5cBPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsEErxb5cBPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
