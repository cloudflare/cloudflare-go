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

// Union satisfied by [PcapNewResponse3K3498u9PcapsResponseSimple] or
// [PcapNewResponse3K3498u9PcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse3K3498u9PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse3K3498u9PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse3K3498u9PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse3K3498u9PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse3K3498u9PcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse3K3498u9PcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse3K3498u9PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse3K3498u9PcapsResponseSimple]
type pcapNewResponse3K3498u9PcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse3K3498u9PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse3K3498u9PcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse3K3498u9PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse3K3498u9PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse3K3498u9PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse3K3498u9PcapsResponseSimpleFilterV1]
type pcapNewResponse3K3498u9PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse3K3498u9PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse3K3498u9PcapsResponseSimpleStatus string

const (
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusUnknown           PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "unknown"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusSuccess           PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "success"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusPending           PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "pending"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusRunning           PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "running"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusConversionPending PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusConversionRunning PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusComplete          PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "complete"
	PcapNewResponse3K3498u9PcapsResponseSimpleStatusFailed            PcapNewResponse3K3498u9PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse3K3498u9PcapsResponseSimpleSystem string

const (
	PcapNewResponse3K3498u9PcapsResponseSimpleSystemMagicTransit PcapNewResponse3K3498u9PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse3K3498u9PcapsResponseSimpleType string

const (
	PcapNewResponse3K3498u9PcapsResponseSimpleTypeSimple PcapNewResponse3K3498u9PcapsResponseSimpleType = "simple"
	PcapNewResponse3K3498u9PcapsResponseSimpleTypeFull   PcapNewResponse3K3498u9PcapsResponseSimpleType = "full"
)

type PcapNewResponse3K3498u9PcapsResponseFull struct {
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
	FilterV1 PcapNewResponse3K3498u9PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse3K3498u9PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse3K3498u9PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse3K3498u9PcapsResponseFullType `json:"type"`
	JSON pcapNewResponse3K3498u9PcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse3K3498u9PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse3K3498u9PcapsResponseFull]
type pcapNewResponse3K3498u9PcapsResponseFullJSON struct {
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

func (r *PcapNewResponse3K3498u9PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse3K3498u9PcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse3K3498u9PcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse3K3498u9PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse3K3498u9PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse3K3498u9PcapsResponseFullFilterV1]
type pcapNewResponse3K3498u9PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse3K3498u9PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse3K3498u9PcapsResponseFullStatus string

const (
	PcapNewResponse3K3498u9PcapsResponseFullStatusUnknown           PcapNewResponse3K3498u9PcapsResponseFullStatus = "unknown"
	PcapNewResponse3K3498u9PcapsResponseFullStatusSuccess           PcapNewResponse3K3498u9PcapsResponseFullStatus = "success"
	PcapNewResponse3K3498u9PcapsResponseFullStatusPending           PcapNewResponse3K3498u9PcapsResponseFullStatus = "pending"
	PcapNewResponse3K3498u9PcapsResponseFullStatusRunning           PcapNewResponse3K3498u9PcapsResponseFullStatus = "running"
	PcapNewResponse3K3498u9PcapsResponseFullStatusConversionPending PcapNewResponse3K3498u9PcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse3K3498u9PcapsResponseFullStatusConversionRunning PcapNewResponse3K3498u9PcapsResponseFullStatus = "conversion_running"
	PcapNewResponse3K3498u9PcapsResponseFullStatusComplete          PcapNewResponse3K3498u9PcapsResponseFullStatus = "complete"
	PcapNewResponse3K3498u9PcapsResponseFullStatusFailed            PcapNewResponse3K3498u9PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse3K3498u9PcapsResponseFullSystem string

const (
	PcapNewResponse3K3498u9PcapsResponseFullSystemMagicTransit PcapNewResponse3K3498u9PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse3K3498u9PcapsResponseFullType string

const (
	PcapNewResponse3K3498u9PcapsResponseFullTypeSimple PcapNewResponse3K3498u9PcapsResponseFullType = "simple"
	PcapNewResponse3K3498u9PcapsResponseFullTypeFull   PcapNewResponse3K3498u9PcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse3K3498u9PcapsResponseSimple] or
// [PcapListResponse3K3498u9PcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse3K3498u9PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse3K3498u9PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse3K3498u9PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse3K3498u9PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse3K3498u9PcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse3K3498u9PcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse3K3498u9PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse3K3498u9PcapsResponseSimple]
type pcapListResponse3K3498u9PcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse3K3498u9PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse3K3498u9PcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse3K3498u9PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse3K3498u9PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse3K3498u9PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse3K3498u9PcapsResponseSimpleFilterV1]
type pcapListResponse3K3498u9PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse3K3498u9PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse3K3498u9PcapsResponseSimpleStatus string

const (
	PcapListResponse3K3498u9PcapsResponseSimpleStatusUnknown           PcapListResponse3K3498u9PcapsResponseSimpleStatus = "unknown"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusSuccess           PcapListResponse3K3498u9PcapsResponseSimpleStatus = "success"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusPending           PcapListResponse3K3498u9PcapsResponseSimpleStatus = "pending"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusRunning           PcapListResponse3K3498u9PcapsResponseSimpleStatus = "running"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusConversionPending PcapListResponse3K3498u9PcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusConversionRunning PcapListResponse3K3498u9PcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusComplete          PcapListResponse3K3498u9PcapsResponseSimpleStatus = "complete"
	PcapListResponse3K3498u9PcapsResponseSimpleStatusFailed            PcapListResponse3K3498u9PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse3K3498u9PcapsResponseSimpleSystem string

const (
	PcapListResponse3K3498u9PcapsResponseSimpleSystemMagicTransit PcapListResponse3K3498u9PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse3K3498u9PcapsResponseSimpleType string

const (
	PcapListResponse3K3498u9PcapsResponseSimpleTypeSimple PcapListResponse3K3498u9PcapsResponseSimpleType = "simple"
	PcapListResponse3K3498u9PcapsResponseSimpleTypeFull   PcapListResponse3K3498u9PcapsResponseSimpleType = "full"
)

type PcapListResponse3K3498u9PcapsResponseFull struct {
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
	FilterV1 PcapListResponse3K3498u9PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse3K3498u9PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse3K3498u9PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse3K3498u9PcapsResponseFullType `json:"type"`
	JSON pcapListResponse3K3498u9PcapsResponseFullJSON `json:"-"`
}

// pcapListResponse3K3498u9PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse3K3498u9PcapsResponseFull]
type pcapListResponse3K3498u9PcapsResponseFullJSON struct {
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

func (r *PcapListResponse3K3498u9PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse3K3498u9PcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse3K3498u9PcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse3K3498u9PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse3K3498u9PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse3K3498u9PcapsResponseFullFilterV1]
type pcapListResponse3K3498u9PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse3K3498u9PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse3K3498u9PcapsResponseFullStatus string

const (
	PcapListResponse3K3498u9PcapsResponseFullStatusUnknown           PcapListResponse3K3498u9PcapsResponseFullStatus = "unknown"
	PcapListResponse3K3498u9PcapsResponseFullStatusSuccess           PcapListResponse3K3498u9PcapsResponseFullStatus = "success"
	PcapListResponse3K3498u9PcapsResponseFullStatusPending           PcapListResponse3K3498u9PcapsResponseFullStatus = "pending"
	PcapListResponse3K3498u9PcapsResponseFullStatusRunning           PcapListResponse3K3498u9PcapsResponseFullStatus = "running"
	PcapListResponse3K3498u9PcapsResponseFullStatusConversionPending PcapListResponse3K3498u9PcapsResponseFullStatus = "conversion_pending"
	PcapListResponse3K3498u9PcapsResponseFullStatusConversionRunning PcapListResponse3K3498u9PcapsResponseFullStatus = "conversion_running"
	PcapListResponse3K3498u9PcapsResponseFullStatusComplete          PcapListResponse3K3498u9PcapsResponseFullStatus = "complete"
	PcapListResponse3K3498u9PcapsResponseFullStatusFailed            PcapListResponse3K3498u9PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse3K3498u9PcapsResponseFullSystem string

const (
	PcapListResponse3K3498u9PcapsResponseFullSystemMagicTransit PcapListResponse3K3498u9PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse3K3498u9PcapsResponseFullType string

const (
	PcapListResponse3K3498u9PcapsResponseFullTypeSimple PcapListResponse3K3498u9PcapsResponseFullType = "simple"
	PcapListResponse3K3498u9PcapsResponseFullTypeFull   PcapListResponse3K3498u9PcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse3K3498u9PcapsResponseSimple] or
// [PcapGetResponse3K3498u9PcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse3K3498u9PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse3K3498u9PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse3K3498u9PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse3K3498u9PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse3K3498u9PcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse3K3498u9PcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse3K3498u9PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse3K3498u9PcapsResponseSimple]
type pcapGetResponse3K3498u9PcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse3K3498u9PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse3K3498u9PcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse3K3498u9PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse3K3498u9PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse3K3498u9PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse3K3498u9PcapsResponseSimpleFilterV1]
type pcapGetResponse3K3498u9PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse3K3498u9PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse3K3498u9PcapsResponseSimpleStatus string

const (
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusUnknown           PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "unknown"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusSuccess           PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "success"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusPending           PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "pending"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusRunning           PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "running"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusConversionPending PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusConversionRunning PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusComplete          PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "complete"
	PcapGetResponse3K3498u9PcapsResponseSimpleStatusFailed            PcapGetResponse3K3498u9PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse3K3498u9PcapsResponseSimpleSystem string

const (
	PcapGetResponse3K3498u9PcapsResponseSimpleSystemMagicTransit PcapGetResponse3K3498u9PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse3K3498u9PcapsResponseSimpleType string

const (
	PcapGetResponse3K3498u9PcapsResponseSimpleTypeSimple PcapGetResponse3K3498u9PcapsResponseSimpleType = "simple"
	PcapGetResponse3K3498u9PcapsResponseSimpleTypeFull   PcapGetResponse3K3498u9PcapsResponseSimpleType = "full"
)

type PcapGetResponse3K3498u9PcapsResponseFull struct {
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
	FilterV1 PcapGetResponse3K3498u9PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse3K3498u9PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse3K3498u9PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse3K3498u9PcapsResponseFullType `json:"type"`
	JSON pcapGetResponse3K3498u9PcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse3K3498u9PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse3K3498u9PcapsResponseFull]
type pcapGetResponse3K3498u9PcapsResponseFullJSON struct {
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

func (r *PcapGetResponse3K3498u9PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse3K3498u9PcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse3K3498u9PcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse3K3498u9PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse3K3498u9PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse3K3498u9PcapsResponseFullFilterV1]
type pcapGetResponse3K3498u9PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse3K3498u9PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse3K3498u9PcapsResponseFullStatus string

const (
	PcapGetResponse3K3498u9PcapsResponseFullStatusUnknown           PcapGetResponse3K3498u9PcapsResponseFullStatus = "unknown"
	PcapGetResponse3K3498u9PcapsResponseFullStatusSuccess           PcapGetResponse3K3498u9PcapsResponseFullStatus = "success"
	PcapGetResponse3K3498u9PcapsResponseFullStatusPending           PcapGetResponse3K3498u9PcapsResponseFullStatus = "pending"
	PcapGetResponse3K3498u9PcapsResponseFullStatusRunning           PcapGetResponse3K3498u9PcapsResponseFullStatus = "running"
	PcapGetResponse3K3498u9PcapsResponseFullStatusConversionPending PcapGetResponse3K3498u9PcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse3K3498u9PcapsResponseFullStatusConversionRunning PcapGetResponse3K3498u9PcapsResponseFullStatus = "conversion_running"
	PcapGetResponse3K3498u9PcapsResponseFullStatusComplete          PcapGetResponse3K3498u9PcapsResponseFullStatus = "complete"
	PcapGetResponse3K3498u9PcapsResponseFullStatusFailed            PcapGetResponse3K3498u9PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse3K3498u9PcapsResponseFullSystem string

const (
	PcapGetResponse3K3498u9PcapsResponseFullSystemMagicTransit PcapGetResponse3K3498u9PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse3K3498u9PcapsResponseFullType string

const (
	PcapGetResponse3K3498u9PcapsResponseFullTypeSimple PcapGetResponse3K3498u9PcapsResponseFullType = "simple"
	PcapGetResponse3K3498u9PcapsResponseFullTypeFull   PcapGetResponse3K3498u9PcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParams3K3498u9PcapsRequestSimple],
// [PcapNewParams3K3498u9PcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParams3K3498u9PcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams3K3498u9PcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams3K3498u9PcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams3K3498u9PcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParams3K3498u9PcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams3K3498u9PcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams3K3498u9PcapsRequestSimpleSystem string

const (
	PcapNewParams3K3498u9PcapsRequestSimpleSystemMagicTransit PcapNewParams3K3498u9PcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams3K3498u9PcapsRequestSimpleType string

const (
	PcapNewParams3K3498u9PcapsRequestSimpleTypeSimple PcapNewParams3K3498u9PcapsRequestSimpleType = "simple"
	PcapNewParams3K3498u9PcapsRequestSimpleTypeFull   PcapNewParams3K3498u9PcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams3K3498u9PcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParams3K3498u9PcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParams3K3498u9PcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams3K3498u9PcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams3K3498u9PcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams3K3498u9PcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParams3K3498u9PcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams3K3498u9PcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams3K3498u9PcapsRequestFullSystem string

const (
	PcapNewParams3K3498u9PcapsRequestFullSystemMagicTransit PcapNewParams3K3498u9PcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams3K3498u9PcapsRequestFullType string

const (
	PcapNewParams3K3498u9PcapsRequestFullTypeSimple PcapNewParams3K3498u9PcapsRequestFullType = "simple"
	PcapNewParams3K3498u9PcapsRequestFullTypeFull   PcapNewParams3K3498u9PcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams3K3498u9PcapsRequestFullFilterV1 struct {
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

func (r PcapNewParams3K3498u9PcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
