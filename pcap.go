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

// Union satisfied by [PcapNewResponse5cvYhiwuPcapsResponseSimple] or
// [PcapNewResponse5cvYhiwuPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse5cvYhiwuPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse5cvYhiwuPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse5cvYhiwuPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse5cvYhiwuPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse5cvYhiwuPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse5cvYhiwuPcapsResponseSimple]
type pcapNewResponse5cvYhiwuPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse5cvYhiwuPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse5cvYhiwuPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1]
type pcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse5cvYhiwuPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus string

const (
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusUnknown           PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusSuccess           PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "success"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusPending           PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "pending"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusRunning           PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "running"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusConversionPending PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusConversionRunning PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusComplete          PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "complete"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleStatusFailed            PcapNewResponse5cvYhiwuPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse5cvYhiwuPcapsResponseSimpleSystem string

const (
	PcapNewResponse5cvYhiwuPcapsResponseSimpleSystemMagicTransit PcapNewResponse5cvYhiwuPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse5cvYhiwuPcapsResponseSimpleType string

const (
	PcapNewResponse5cvYhiwuPcapsResponseSimpleTypeSimple PcapNewResponse5cvYhiwuPcapsResponseSimpleType = "simple"
	PcapNewResponse5cvYhiwuPcapsResponseSimpleTypeFull   PcapNewResponse5cvYhiwuPcapsResponseSimpleType = "full"
)

type PcapNewResponse5cvYhiwuPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse5cvYhiwuPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse5cvYhiwuPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse5cvYhiwuPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse5cvYhiwuPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse5cvYhiwuPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse5cvYhiwuPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse5cvYhiwuPcapsResponseFull]
type pcapNewResponse5cvYhiwuPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse5cvYhiwuPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse5cvYhiwuPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse5cvYhiwuPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse5cvYhiwuPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse5cvYhiwuPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse5cvYhiwuPcapsResponseFullFilterV1]
type pcapNewResponse5cvYhiwuPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse5cvYhiwuPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse5cvYhiwuPcapsResponseFullStatus string

const (
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusUnknown           PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "unknown"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusSuccess           PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "success"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusPending           PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "pending"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusRunning           PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "running"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusConversionPending PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusConversionRunning PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusComplete          PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "complete"
	PcapNewResponse5cvYhiwuPcapsResponseFullStatusFailed            PcapNewResponse5cvYhiwuPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse5cvYhiwuPcapsResponseFullSystem string

const (
	PcapNewResponse5cvYhiwuPcapsResponseFullSystemMagicTransit PcapNewResponse5cvYhiwuPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse5cvYhiwuPcapsResponseFullType string

const (
	PcapNewResponse5cvYhiwuPcapsResponseFullTypeSimple PcapNewResponse5cvYhiwuPcapsResponseFullType = "simple"
	PcapNewResponse5cvYhiwuPcapsResponseFullTypeFull   PcapNewResponse5cvYhiwuPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse5cvYhiwuPcapsResponseSimple] or
// [PcapListResponse5cvYhiwuPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse5cvYhiwuPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse5cvYhiwuPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse5cvYhiwuPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse5cvYhiwuPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse5cvYhiwuPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse5cvYhiwuPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse5cvYhiwuPcapsResponseSimple]
type pcapListResponse5cvYhiwuPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse5cvYhiwuPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse5cvYhiwuPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1]
type pcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse5cvYhiwuPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse5cvYhiwuPcapsResponseSimpleStatus string

const (
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusUnknown           PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "unknown"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusSuccess           PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "success"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusPending           PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "pending"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusRunning           PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "running"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusConversionPending PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusConversionRunning PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusComplete          PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "complete"
	PcapListResponse5cvYhiwuPcapsResponseSimpleStatusFailed            PcapListResponse5cvYhiwuPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse5cvYhiwuPcapsResponseSimpleSystem string

const (
	PcapListResponse5cvYhiwuPcapsResponseSimpleSystemMagicTransit PcapListResponse5cvYhiwuPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse5cvYhiwuPcapsResponseSimpleType string

const (
	PcapListResponse5cvYhiwuPcapsResponseSimpleTypeSimple PcapListResponse5cvYhiwuPcapsResponseSimpleType = "simple"
	PcapListResponse5cvYhiwuPcapsResponseSimpleTypeFull   PcapListResponse5cvYhiwuPcapsResponseSimpleType = "full"
)

type PcapListResponse5cvYhiwuPcapsResponseFull struct {
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
	FilterV1 PcapListResponse5cvYhiwuPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse5cvYhiwuPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse5cvYhiwuPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse5cvYhiwuPcapsResponseFullType `json:"type"`
	JSON pcapListResponse5cvYhiwuPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse5cvYhiwuPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse5cvYhiwuPcapsResponseFull]
type pcapListResponse5cvYhiwuPcapsResponseFullJSON struct {
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

func (r *PcapListResponse5cvYhiwuPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse5cvYhiwuPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse5cvYhiwuPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse5cvYhiwuPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse5cvYhiwuPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse5cvYhiwuPcapsResponseFullFilterV1]
type pcapListResponse5cvYhiwuPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse5cvYhiwuPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse5cvYhiwuPcapsResponseFullStatus string

const (
	PcapListResponse5cvYhiwuPcapsResponseFullStatusUnknown           PcapListResponse5cvYhiwuPcapsResponseFullStatus = "unknown"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusSuccess           PcapListResponse5cvYhiwuPcapsResponseFullStatus = "success"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusPending           PcapListResponse5cvYhiwuPcapsResponseFullStatus = "pending"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusRunning           PcapListResponse5cvYhiwuPcapsResponseFullStatus = "running"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusConversionPending PcapListResponse5cvYhiwuPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusConversionRunning PcapListResponse5cvYhiwuPcapsResponseFullStatus = "conversion_running"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusComplete          PcapListResponse5cvYhiwuPcapsResponseFullStatus = "complete"
	PcapListResponse5cvYhiwuPcapsResponseFullStatusFailed            PcapListResponse5cvYhiwuPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse5cvYhiwuPcapsResponseFullSystem string

const (
	PcapListResponse5cvYhiwuPcapsResponseFullSystemMagicTransit PcapListResponse5cvYhiwuPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse5cvYhiwuPcapsResponseFullType string

const (
	PcapListResponse5cvYhiwuPcapsResponseFullTypeSimple PcapListResponse5cvYhiwuPcapsResponseFullType = "simple"
	PcapListResponse5cvYhiwuPcapsResponseFullTypeFull   PcapListResponse5cvYhiwuPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse5cvYhiwuPcapsResponseSimple] or
// [PcapGetResponse5cvYhiwuPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse5cvYhiwuPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse5cvYhiwuPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse5cvYhiwuPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse5cvYhiwuPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse5cvYhiwuPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse5cvYhiwuPcapsResponseSimple]
type pcapGetResponse5cvYhiwuPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse5cvYhiwuPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse5cvYhiwuPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1]
type pcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse5cvYhiwuPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus string

const (
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusUnknown           PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusSuccess           PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "success"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusPending           PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "pending"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusRunning           PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "running"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusConversionPending PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusConversionRunning PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusComplete          PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "complete"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleStatusFailed            PcapGetResponse5cvYhiwuPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse5cvYhiwuPcapsResponseSimpleSystem string

const (
	PcapGetResponse5cvYhiwuPcapsResponseSimpleSystemMagicTransit PcapGetResponse5cvYhiwuPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse5cvYhiwuPcapsResponseSimpleType string

const (
	PcapGetResponse5cvYhiwuPcapsResponseSimpleTypeSimple PcapGetResponse5cvYhiwuPcapsResponseSimpleType = "simple"
	PcapGetResponse5cvYhiwuPcapsResponseSimpleTypeFull   PcapGetResponse5cvYhiwuPcapsResponseSimpleType = "full"
)

type PcapGetResponse5cvYhiwuPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse5cvYhiwuPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse5cvYhiwuPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse5cvYhiwuPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse5cvYhiwuPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse5cvYhiwuPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse5cvYhiwuPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse5cvYhiwuPcapsResponseFull]
type pcapGetResponse5cvYhiwuPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse5cvYhiwuPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse5cvYhiwuPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse5cvYhiwuPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse5cvYhiwuPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse5cvYhiwuPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse5cvYhiwuPcapsResponseFullFilterV1]
type pcapGetResponse5cvYhiwuPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse5cvYhiwuPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse5cvYhiwuPcapsResponseFullStatus string

const (
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusUnknown           PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "unknown"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusSuccess           PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "success"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusPending           PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "pending"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusRunning           PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "running"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusConversionPending PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusConversionRunning PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusComplete          PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "complete"
	PcapGetResponse5cvYhiwuPcapsResponseFullStatusFailed            PcapGetResponse5cvYhiwuPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse5cvYhiwuPcapsResponseFullSystem string

const (
	PcapGetResponse5cvYhiwuPcapsResponseFullSystemMagicTransit PcapGetResponse5cvYhiwuPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse5cvYhiwuPcapsResponseFullType string

const (
	PcapGetResponse5cvYhiwuPcapsResponseFullTypeSimple PcapGetResponse5cvYhiwuPcapsResponseFullType = "simple"
	PcapGetResponse5cvYhiwuPcapsResponseFullTypeFull   PcapGetResponse5cvYhiwuPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParams5cvYhiwuPcapsRequestSimple],
// [PcapNewParams5cvYhiwuPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParams5cvYhiwuPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams5cvYhiwuPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams5cvYhiwuPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams5cvYhiwuPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParams5cvYhiwuPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams5cvYhiwuPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams5cvYhiwuPcapsRequestSimpleSystem string

const (
	PcapNewParams5cvYhiwuPcapsRequestSimpleSystemMagicTransit PcapNewParams5cvYhiwuPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams5cvYhiwuPcapsRequestSimpleType string

const (
	PcapNewParams5cvYhiwuPcapsRequestSimpleTypeSimple PcapNewParams5cvYhiwuPcapsRequestSimpleType = "simple"
	PcapNewParams5cvYhiwuPcapsRequestSimpleTypeFull   PcapNewParams5cvYhiwuPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams5cvYhiwuPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParams5cvYhiwuPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParams5cvYhiwuPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams5cvYhiwuPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams5cvYhiwuPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams5cvYhiwuPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParams5cvYhiwuPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams5cvYhiwuPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams5cvYhiwuPcapsRequestFullSystem string

const (
	PcapNewParams5cvYhiwuPcapsRequestFullSystemMagicTransit PcapNewParams5cvYhiwuPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams5cvYhiwuPcapsRequestFullType string

const (
	PcapNewParams5cvYhiwuPcapsRequestFullTypeSimple PcapNewParams5cvYhiwuPcapsRequestFullType = "simple"
	PcapNewParams5cvYhiwuPcapsRequestFullTypeFull   PcapNewParams5cvYhiwuPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams5cvYhiwuPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParams5cvYhiwuPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
