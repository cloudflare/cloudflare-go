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

// Union satisfied by [PcapNewResponseHwGa9CrdPcapsResponseSimple] or
// [PcapNewResponseHwGa9CrdPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseHwGa9CrdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseHwGa9CrdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseHwGa9CrdPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseHwGa9CrdPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseHwGa9CrdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseHwGa9CrdPcapsResponseSimple]
type pcapNewResponseHwGa9CrdPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseHwGa9CrdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseHwGa9CrdPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1]
type pcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseHwGa9CrdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus string

const (
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusUnknown           PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusSuccess           PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "success"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusPending           PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "pending"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusRunning           PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "running"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusConversionPending PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusConversionRunning PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusComplete          PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "complete"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleStatusFailed            PcapNewResponseHwGa9CrdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseHwGa9CrdPcapsResponseSimpleSystem string

const (
	PcapNewResponseHwGa9CrdPcapsResponseSimpleSystemMagicTransit PcapNewResponseHwGa9CrdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseHwGa9CrdPcapsResponseSimpleType string

const (
	PcapNewResponseHwGa9CrdPcapsResponseSimpleTypeSimple PcapNewResponseHwGa9CrdPcapsResponseSimpleType = "simple"
	PcapNewResponseHwGa9CrdPcapsResponseSimpleTypeFull   PcapNewResponseHwGa9CrdPcapsResponseSimpleType = "full"
)

type PcapNewResponseHwGa9CrdPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseHwGa9CrdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseHwGa9CrdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseHwGa9CrdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseHwGa9CrdPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseHwGa9CrdPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseHwGa9CrdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseHwGa9CrdPcapsResponseFull]
type pcapNewResponseHwGa9CrdPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseHwGa9CrdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseHwGa9CrdPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseHwGa9CrdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseHwGa9CrdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseHwGa9CrdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseHwGa9CrdPcapsResponseFullFilterV1]
type pcapNewResponseHwGa9CrdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseHwGa9CrdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseHwGa9CrdPcapsResponseFullStatus string

const (
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusUnknown           PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "unknown"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusSuccess           PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "success"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusPending           PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "pending"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusRunning           PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "running"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusConversionPending PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusConversionRunning PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusComplete          PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "complete"
	PcapNewResponseHwGa9CrdPcapsResponseFullStatusFailed            PcapNewResponseHwGa9CrdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseHwGa9CrdPcapsResponseFullSystem string

const (
	PcapNewResponseHwGa9CrdPcapsResponseFullSystemMagicTransit PcapNewResponseHwGa9CrdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseHwGa9CrdPcapsResponseFullType string

const (
	PcapNewResponseHwGa9CrdPcapsResponseFullTypeSimple PcapNewResponseHwGa9CrdPcapsResponseFullType = "simple"
	PcapNewResponseHwGa9CrdPcapsResponseFullTypeFull   PcapNewResponseHwGa9CrdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseHwGa9CrdPcapsResponseSimple] or
// [PcapListResponseHwGa9CrdPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseHwGa9CrdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseHwGa9CrdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseHwGa9CrdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseHwGa9CrdPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseHwGa9CrdPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseHwGa9CrdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseHwGa9CrdPcapsResponseSimple]
type pcapListResponseHwGa9CrdPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseHwGa9CrdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseHwGa9CrdPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1]
type pcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseHwGa9CrdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseHwGa9CrdPcapsResponseSimpleStatus string

const (
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusUnknown           PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "unknown"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusSuccess           PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "success"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusPending           PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "pending"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusRunning           PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "running"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusConversionPending PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusConversionRunning PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusComplete          PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "complete"
	PcapListResponseHwGa9CrdPcapsResponseSimpleStatusFailed            PcapListResponseHwGa9CrdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseHwGa9CrdPcapsResponseSimpleSystem string

const (
	PcapListResponseHwGa9CrdPcapsResponseSimpleSystemMagicTransit PcapListResponseHwGa9CrdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseHwGa9CrdPcapsResponseSimpleType string

const (
	PcapListResponseHwGa9CrdPcapsResponseSimpleTypeSimple PcapListResponseHwGa9CrdPcapsResponseSimpleType = "simple"
	PcapListResponseHwGa9CrdPcapsResponseSimpleTypeFull   PcapListResponseHwGa9CrdPcapsResponseSimpleType = "full"
)

type PcapListResponseHwGa9CrdPcapsResponseFull struct {
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
	FilterV1 PcapListResponseHwGa9CrdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseHwGa9CrdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseHwGa9CrdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseHwGa9CrdPcapsResponseFullType `json:"type"`
	JSON pcapListResponseHwGa9CrdPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseHwGa9CrdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseHwGa9CrdPcapsResponseFull]
type pcapListResponseHwGa9CrdPcapsResponseFullJSON struct {
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

func (r *PcapListResponseHwGa9CrdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseHwGa9CrdPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseHwGa9CrdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseHwGa9CrdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseHwGa9CrdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseHwGa9CrdPcapsResponseFullFilterV1]
type pcapListResponseHwGa9CrdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseHwGa9CrdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseHwGa9CrdPcapsResponseFullStatus string

const (
	PcapListResponseHwGa9CrdPcapsResponseFullStatusUnknown           PcapListResponseHwGa9CrdPcapsResponseFullStatus = "unknown"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusSuccess           PcapListResponseHwGa9CrdPcapsResponseFullStatus = "success"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusPending           PcapListResponseHwGa9CrdPcapsResponseFullStatus = "pending"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusRunning           PcapListResponseHwGa9CrdPcapsResponseFullStatus = "running"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusConversionPending PcapListResponseHwGa9CrdPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusConversionRunning PcapListResponseHwGa9CrdPcapsResponseFullStatus = "conversion_running"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusComplete          PcapListResponseHwGa9CrdPcapsResponseFullStatus = "complete"
	PcapListResponseHwGa9CrdPcapsResponseFullStatusFailed            PcapListResponseHwGa9CrdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseHwGa9CrdPcapsResponseFullSystem string

const (
	PcapListResponseHwGa9CrdPcapsResponseFullSystemMagicTransit PcapListResponseHwGa9CrdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseHwGa9CrdPcapsResponseFullType string

const (
	PcapListResponseHwGa9CrdPcapsResponseFullTypeSimple PcapListResponseHwGa9CrdPcapsResponseFullType = "simple"
	PcapListResponseHwGa9CrdPcapsResponseFullTypeFull   PcapListResponseHwGa9CrdPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseHwGa9CrdPcapsResponseSimple] or
// [PcapGetResponseHwGa9CrdPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseHwGa9CrdPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseHwGa9CrdPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseHwGa9CrdPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseHwGa9CrdPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseHwGa9CrdPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseHwGa9CrdPcapsResponseSimple]
type pcapGetResponseHwGa9CrdPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseHwGa9CrdPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseHwGa9CrdPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1]
type pcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseHwGa9CrdPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus string

const (
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusUnknown           PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusSuccess           PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "success"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusPending           PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "pending"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusRunning           PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "running"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusConversionPending PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusConversionRunning PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusComplete          PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "complete"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleStatusFailed            PcapGetResponseHwGa9CrdPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseHwGa9CrdPcapsResponseSimpleSystem string

const (
	PcapGetResponseHwGa9CrdPcapsResponseSimpleSystemMagicTransit PcapGetResponseHwGa9CrdPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseHwGa9CrdPcapsResponseSimpleType string

const (
	PcapGetResponseHwGa9CrdPcapsResponseSimpleTypeSimple PcapGetResponseHwGa9CrdPcapsResponseSimpleType = "simple"
	PcapGetResponseHwGa9CrdPcapsResponseSimpleTypeFull   PcapGetResponseHwGa9CrdPcapsResponseSimpleType = "full"
)

type PcapGetResponseHwGa9CrdPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseHwGa9CrdPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseHwGa9CrdPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseHwGa9CrdPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseHwGa9CrdPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseHwGa9CrdPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseHwGa9CrdPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseHwGa9CrdPcapsResponseFull]
type pcapGetResponseHwGa9CrdPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseHwGa9CrdPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseHwGa9CrdPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseHwGa9CrdPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseHwGa9CrdPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseHwGa9CrdPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseHwGa9CrdPcapsResponseFullFilterV1]
type pcapGetResponseHwGa9CrdPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseHwGa9CrdPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseHwGa9CrdPcapsResponseFullStatus string

const (
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusUnknown           PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "unknown"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusSuccess           PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "success"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusPending           PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "pending"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusRunning           PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "running"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusConversionPending PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusConversionRunning PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusComplete          PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "complete"
	PcapGetResponseHwGa9CrdPcapsResponseFullStatusFailed            PcapGetResponseHwGa9CrdPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseHwGa9CrdPcapsResponseFullSystem string

const (
	PcapGetResponseHwGa9CrdPcapsResponseFullSystemMagicTransit PcapGetResponseHwGa9CrdPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseHwGa9CrdPcapsResponseFullType string

const (
	PcapGetResponseHwGa9CrdPcapsResponseFullTypeSimple PcapGetResponseHwGa9CrdPcapsResponseFullType = "simple"
	PcapGetResponseHwGa9CrdPcapsResponseFullTypeFull   PcapGetResponseHwGa9CrdPcapsResponseFullType = "full"
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
