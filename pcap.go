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

// Union satisfied by [PcapNewResponse2PfkzezdPcapsResponseSimple] or
// [PcapNewResponse2PfkzezdPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse2PfkzezdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2PfkzezdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2PfkzezdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2PfkzezdPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse2PfkzezdPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse2PfkzezdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse2PfkzezdPcapsResponseSimple]
type pcapNewResponse2PfkzezdPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse2PfkzezdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2PfkzezdPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1]
type pcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2PfkzezdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2PfkzezdPcapsResponseSimpleStatus string

const (
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusUnknown           PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusSuccess           PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "success"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusPending           PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "pending"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusRunning           PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "running"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusConversionPending PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusConversionRunning PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusComplete          PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "complete"
	PcapNewResponse2PfkzezdPcapsResponseSimpleStatusFailed            PcapNewResponse2PfkzezdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2PfkzezdPcapsResponseSimpleSystem string

const (
	PcapNewResponse2PfkzezdPcapsResponseSimpleSystemMagicTransit PcapNewResponse2PfkzezdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2PfkzezdPcapsResponseSimpleType string

const (
	PcapNewResponse2PfkzezdPcapsResponseSimpleTypeSimple PcapNewResponse2PfkzezdPcapsResponseSimpleType = "simple"
	PcapNewResponse2PfkzezdPcapsResponseSimpleTypeFull   PcapNewResponse2PfkzezdPcapsResponseSimpleType = "full"
)

type PcapNewResponse2PfkzezdPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse2PfkzezdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2PfkzezdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2PfkzezdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2PfkzezdPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse2PfkzezdPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse2PfkzezdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse2PfkzezdPcapsResponseFull]
type pcapNewResponse2PfkzezdPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse2PfkzezdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2PfkzezdPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2PfkzezdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse2PfkzezdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse2PfkzezdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse2PfkzezdPcapsResponseFullFilterV1]
type pcapNewResponse2PfkzezdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2PfkzezdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2PfkzezdPcapsResponseFullStatus string

const (
	PcapNewResponse2PfkzezdPcapsResponseFullStatusUnknown           PcapNewResponse2PfkzezdPcapsResponseFullStatus = "unknown"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusSuccess           PcapNewResponse2PfkzezdPcapsResponseFullStatus = "success"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusPending           PcapNewResponse2PfkzezdPcapsResponseFullStatus = "pending"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusRunning           PcapNewResponse2PfkzezdPcapsResponseFullStatus = "running"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusConversionPending PcapNewResponse2PfkzezdPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusConversionRunning PcapNewResponse2PfkzezdPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusComplete          PcapNewResponse2PfkzezdPcapsResponseFullStatus = "complete"
	PcapNewResponse2PfkzezdPcapsResponseFullStatusFailed            PcapNewResponse2PfkzezdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2PfkzezdPcapsResponseFullSystem string

const (
	PcapNewResponse2PfkzezdPcapsResponseFullSystemMagicTransit PcapNewResponse2PfkzezdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2PfkzezdPcapsResponseFullType string

const (
	PcapNewResponse2PfkzezdPcapsResponseFullTypeSimple PcapNewResponse2PfkzezdPcapsResponseFullType = "simple"
	PcapNewResponse2PfkzezdPcapsResponseFullTypeFull   PcapNewResponse2PfkzezdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse2PfkzezdPcapsResponseSimple] or
// [PcapListResponse2PfkzezdPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse2PfkzezdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse2PfkzezdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2PfkzezdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2PfkzezdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2PfkzezdPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse2PfkzezdPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse2PfkzezdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse2PfkzezdPcapsResponseSimple]
type pcapListResponse2PfkzezdPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse2PfkzezdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2PfkzezdPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2PfkzezdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse2PfkzezdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse2PfkzezdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse2PfkzezdPcapsResponseSimpleFilterV1]
type pcapListResponse2PfkzezdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2PfkzezdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2PfkzezdPcapsResponseSimpleStatus string

const (
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusUnknown           PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "unknown"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusSuccess           PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "success"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusPending           PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "pending"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusRunning           PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "running"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusConversionPending PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusConversionRunning PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusComplete          PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "complete"
	PcapListResponse2PfkzezdPcapsResponseSimpleStatusFailed            PcapListResponse2PfkzezdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2PfkzezdPcapsResponseSimpleSystem string

const (
	PcapListResponse2PfkzezdPcapsResponseSimpleSystemMagicTransit PcapListResponse2PfkzezdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2PfkzezdPcapsResponseSimpleType string

const (
	PcapListResponse2PfkzezdPcapsResponseSimpleTypeSimple PcapListResponse2PfkzezdPcapsResponseSimpleType = "simple"
	PcapListResponse2PfkzezdPcapsResponseSimpleTypeFull   PcapListResponse2PfkzezdPcapsResponseSimpleType = "full"
)

type PcapListResponse2PfkzezdPcapsResponseFull struct {
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
	FilterV1 PcapListResponse2PfkzezdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2PfkzezdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2PfkzezdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2PfkzezdPcapsResponseFullType `json:"type"`
	JSON pcapListResponse2PfkzezdPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse2PfkzezdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse2PfkzezdPcapsResponseFull]
type pcapListResponse2PfkzezdPcapsResponseFullJSON struct {
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

func (r *PcapListResponse2PfkzezdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2PfkzezdPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2PfkzezdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse2PfkzezdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse2PfkzezdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse2PfkzezdPcapsResponseFullFilterV1]
type pcapListResponse2PfkzezdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2PfkzezdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2PfkzezdPcapsResponseFullStatus string

const (
	PcapListResponse2PfkzezdPcapsResponseFullStatusUnknown           PcapListResponse2PfkzezdPcapsResponseFullStatus = "unknown"
	PcapListResponse2PfkzezdPcapsResponseFullStatusSuccess           PcapListResponse2PfkzezdPcapsResponseFullStatus = "success"
	PcapListResponse2PfkzezdPcapsResponseFullStatusPending           PcapListResponse2PfkzezdPcapsResponseFullStatus = "pending"
	PcapListResponse2PfkzezdPcapsResponseFullStatusRunning           PcapListResponse2PfkzezdPcapsResponseFullStatus = "running"
	PcapListResponse2PfkzezdPcapsResponseFullStatusConversionPending PcapListResponse2PfkzezdPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse2PfkzezdPcapsResponseFullStatusConversionRunning PcapListResponse2PfkzezdPcapsResponseFullStatus = "conversion_running"
	PcapListResponse2PfkzezdPcapsResponseFullStatusComplete          PcapListResponse2PfkzezdPcapsResponseFullStatus = "complete"
	PcapListResponse2PfkzezdPcapsResponseFullStatusFailed            PcapListResponse2PfkzezdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2PfkzezdPcapsResponseFullSystem string

const (
	PcapListResponse2PfkzezdPcapsResponseFullSystemMagicTransit PcapListResponse2PfkzezdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2PfkzezdPcapsResponseFullType string

const (
	PcapListResponse2PfkzezdPcapsResponseFullTypeSimple PcapListResponse2PfkzezdPcapsResponseFullType = "simple"
	PcapListResponse2PfkzezdPcapsResponseFullTypeFull   PcapListResponse2PfkzezdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse2PfkzezdPcapsResponseSimple] or
// [PcapGetResponse2PfkzezdPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse2PfkzezdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2PfkzezdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2PfkzezdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2PfkzezdPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse2PfkzezdPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse2PfkzezdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse2PfkzezdPcapsResponseSimple]
type pcapGetResponse2PfkzezdPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse2PfkzezdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2PfkzezdPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1]
type pcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2PfkzezdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2PfkzezdPcapsResponseSimpleStatus string

const (
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusUnknown           PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusSuccess           PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "success"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusPending           PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "pending"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusRunning           PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "running"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusConversionPending PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusConversionRunning PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusComplete          PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "complete"
	PcapGetResponse2PfkzezdPcapsResponseSimpleStatusFailed            PcapGetResponse2PfkzezdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2PfkzezdPcapsResponseSimpleSystem string

const (
	PcapGetResponse2PfkzezdPcapsResponseSimpleSystemMagicTransit PcapGetResponse2PfkzezdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2PfkzezdPcapsResponseSimpleType string

const (
	PcapGetResponse2PfkzezdPcapsResponseSimpleTypeSimple PcapGetResponse2PfkzezdPcapsResponseSimpleType = "simple"
	PcapGetResponse2PfkzezdPcapsResponseSimpleTypeFull   PcapGetResponse2PfkzezdPcapsResponseSimpleType = "full"
)

type PcapGetResponse2PfkzezdPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse2PfkzezdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2PfkzezdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2PfkzezdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2PfkzezdPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse2PfkzezdPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse2PfkzezdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse2PfkzezdPcapsResponseFull]
type pcapGetResponse2PfkzezdPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse2PfkzezdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2PfkzezdPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2PfkzezdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse2PfkzezdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse2PfkzezdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse2PfkzezdPcapsResponseFullFilterV1]
type pcapGetResponse2PfkzezdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2PfkzezdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2PfkzezdPcapsResponseFullStatus string

const (
	PcapGetResponse2PfkzezdPcapsResponseFullStatusUnknown           PcapGetResponse2PfkzezdPcapsResponseFullStatus = "unknown"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusSuccess           PcapGetResponse2PfkzezdPcapsResponseFullStatus = "success"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusPending           PcapGetResponse2PfkzezdPcapsResponseFullStatus = "pending"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusRunning           PcapGetResponse2PfkzezdPcapsResponseFullStatus = "running"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusConversionPending PcapGetResponse2PfkzezdPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusConversionRunning PcapGetResponse2PfkzezdPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusComplete          PcapGetResponse2PfkzezdPcapsResponseFullStatus = "complete"
	PcapGetResponse2PfkzezdPcapsResponseFullStatusFailed            PcapGetResponse2PfkzezdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2PfkzezdPcapsResponseFullSystem string

const (
	PcapGetResponse2PfkzezdPcapsResponseFullSystemMagicTransit PcapGetResponse2PfkzezdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2PfkzezdPcapsResponseFullType string

const (
	PcapGetResponse2PfkzezdPcapsResponseFullTypeSimple PcapGetResponse2PfkzezdPcapsResponseFullType = "simple"
	PcapGetResponse2PfkzezdPcapsResponseFullTypeFull   PcapGetResponse2PfkzezdPcapsResponseFullType = "full"
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
