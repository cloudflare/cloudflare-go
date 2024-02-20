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

// Union satisfied by [PcapNewResponseM2NEk72FPcapsResponseSimple] or
// [PcapNewResponseM2NEk72FPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseM2NEk72FPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseM2NEk72FPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseM2NEk72FPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseM2NEk72FPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseM2NEk72FPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseM2NEk72FPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseM2NEk72FPcapsResponseSimple]
type pcapNewResponseM2NEk72FPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseM2NEk72FPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseM2NEk72FPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1]
type pcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseM2NEk72FPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseM2NEk72FPcapsResponseSimpleStatus string

const (
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusUnknown           PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusSuccess           PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "success"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusPending           PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "pending"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusRunning           PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "running"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusConversionPending PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusConversionRunning PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusComplete          PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "complete"
	PcapNewResponseM2NEk72FPcapsResponseSimpleStatusFailed            PcapNewResponseM2NEk72FPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseM2NEk72FPcapsResponseSimpleSystem string

const (
	PcapNewResponseM2NEk72FPcapsResponseSimpleSystemMagicTransit PcapNewResponseM2NEk72FPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseM2NEk72FPcapsResponseSimpleType string

const (
	PcapNewResponseM2NEk72FPcapsResponseSimpleTypeSimple PcapNewResponseM2NEk72FPcapsResponseSimpleType = "simple"
	PcapNewResponseM2NEk72FPcapsResponseSimpleTypeFull   PcapNewResponseM2NEk72FPcapsResponseSimpleType = "full"
)

type PcapNewResponseM2NEk72FPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseM2NEk72FPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseM2NEk72FPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseM2NEk72FPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseM2NEk72FPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseM2NEk72FPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseM2NEk72FPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseM2NEk72FPcapsResponseFull]
type pcapNewResponseM2NEk72FPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseM2NEk72FPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseM2NEk72FPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseM2NEk72FPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseM2NEk72FPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseM2NEk72FPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseM2NEk72FPcapsResponseFullFilterV1]
type pcapNewResponseM2NEk72FPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseM2NEk72FPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseM2NEk72FPcapsResponseFullStatus string

const (
	PcapNewResponseM2NEk72FPcapsResponseFullStatusUnknown           PcapNewResponseM2NEk72FPcapsResponseFullStatus = "unknown"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusSuccess           PcapNewResponseM2NEk72FPcapsResponseFullStatus = "success"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusPending           PcapNewResponseM2NEk72FPcapsResponseFullStatus = "pending"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusRunning           PcapNewResponseM2NEk72FPcapsResponseFullStatus = "running"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusConversionPending PcapNewResponseM2NEk72FPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusConversionRunning PcapNewResponseM2NEk72FPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusComplete          PcapNewResponseM2NEk72FPcapsResponseFullStatus = "complete"
	PcapNewResponseM2NEk72FPcapsResponseFullStatusFailed            PcapNewResponseM2NEk72FPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseM2NEk72FPcapsResponseFullSystem string

const (
	PcapNewResponseM2NEk72FPcapsResponseFullSystemMagicTransit PcapNewResponseM2NEk72FPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseM2NEk72FPcapsResponseFullType string

const (
	PcapNewResponseM2NEk72FPcapsResponseFullTypeSimple PcapNewResponseM2NEk72FPcapsResponseFullType = "simple"
	PcapNewResponseM2NEk72FPcapsResponseFullTypeFull   PcapNewResponseM2NEk72FPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseM2NEk72FPcapsResponseSimple] or
// [PcapListResponseM2NEk72FPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseM2NEk72FPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseM2NEk72FPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseM2NEk72FPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseM2NEk72FPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseM2NEk72FPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseM2NEk72FPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseM2NEk72FPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseM2NEk72FPcapsResponseSimple]
type pcapListResponseM2NEk72FPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseM2NEk72FPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseM2NEk72FPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseM2NEk72FPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseM2NEk72FPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseM2NEk72FPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseM2NEk72FPcapsResponseSimpleFilterV1]
type pcapListResponseM2NEk72FPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseM2NEk72FPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseM2NEk72FPcapsResponseSimpleStatus string

const (
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusUnknown           PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "unknown"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusSuccess           PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "success"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusPending           PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "pending"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusRunning           PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "running"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusConversionPending PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusConversionRunning PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusComplete          PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "complete"
	PcapListResponseM2NEk72FPcapsResponseSimpleStatusFailed            PcapListResponseM2NEk72FPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseM2NEk72FPcapsResponseSimpleSystem string

const (
	PcapListResponseM2NEk72FPcapsResponseSimpleSystemMagicTransit PcapListResponseM2NEk72FPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseM2NEk72FPcapsResponseSimpleType string

const (
	PcapListResponseM2NEk72FPcapsResponseSimpleTypeSimple PcapListResponseM2NEk72FPcapsResponseSimpleType = "simple"
	PcapListResponseM2NEk72FPcapsResponseSimpleTypeFull   PcapListResponseM2NEk72FPcapsResponseSimpleType = "full"
)

type PcapListResponseM2NEk72FPcapsResponseFull struct {
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
	FilterV1 PcapListResponseM2NEk72FPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseM2NEk72FPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseM2NEk72FPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseM2NEk72FPcapsResponseFullType `json:"type"`
	JSON pcapListResponseM2NEk72FPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseM2NEk72FPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseM2NEk72FPcapsResponseFull]
type pcapListResponseM2NEk72FPcapsResponseFullJSON struct {
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

func (r *PcapListResponseM2NEk72FPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseM2NEk72FPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseM2NEk72FPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseM2NEk72FPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseM2NEk72FPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseM2NEk72FPcapsResponseFullFilterV1]
type pcapListResponseM2NEk72FPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseM2NEk72FPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseM2NEk72FPcapsResponseFullStatus string

const (
	PcapListResponseM2NEk72FPcapsResponseFullStatusUnknown           PcapListResponseM2NEk72FPcapsResponseFullStatus = "unknown"
	PcapListResponseM2NEk72FPcapsResponseFullStatusSuccess           PcapListResponseM2NEk72FPcapsResponseFullStatus = "success"
	PcapListResponseM2NEk72FPcapsResponseFullStatusPending           PcapListResponseM2NEk72FPcapsResponseFullStatus = "pending"
	PcapListResponseM2NEk72FPcapsResponseFullStatusRunning           PcapListResponseM2NEk72FPcapsResponseFullStatus = "running"
	PcapListResponseM2NEk72FPcapsResponseFullStatusConversionPending PcapListResponseM2NEk72FPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseM2NEk72FPcapsResponseFullStatusConversionRunning PcapListResponseM2NEk72FPcapsResponseFullStatus = "conversion_running"
	PcapListResponseM2NEk72FPcapsResponseFullStatusComplete          PcapListResponseM2NEk72FPcapsResponseFullStatus = "complete"
	PcapListResponseM2NEk72FPcapsResponseFullStatusFailed            PcapListResponseM2NEk72FPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseM2NEk72FPcapsResponseFullSystem string

const (
	PcapListResponseM2NEk72FPcapsResponseFullSystemMagicTransit PcapListResponseM2NEk72FPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseM2NEk72FPcapsResponseFullType string

const (
	PcapListResponseM2NEk72FPcapsResponseFullTypeSimple PcapListResponseM2NEk72FPcapsResponseFullType = "simple"
	PcapListResponseM2NEk72FPcapsResponseFullTypeFull   PcapListResponseM2NEk72FPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseM2NEk72FPcapsResponseSimple] or
// [PcapGetResponseM2NEk72FPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseM2NEk72FPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseM2NEk72FPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseM2NEk72FPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseM2NEk72FPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseM2NEk72FPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseM2NEk72FPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseM2NEk72FPcapsResponseSimple]
type pcapGetResponseM2NEk72FPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseM2NEk72FPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseM2NEk72FPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1]
type pcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseM2NEk72FPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseM2NEk72FPcapsResponseSimpleStatus string

const (
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusUnknown           PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusSuccess           PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "success"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusPending           PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "pending"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusRunning           PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "running"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusConversionPending PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusConversionRunning PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusComplete          PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "complete"
	PcapGetResponseM2NEk72FPcapsResponseSimpleStatusFailed            PcapGetResponseM2NEk72FPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseM2NEk72FPcapsResponseSimpleSystem string

const (
	PcapGetResponseM2NEk72FPcapsResponseSimpleSystemMagicTransit PcapGetResponseM2NEk72FPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseM2NEk72FPcapsResponseSimpleType string

const (
	PcapGetResponseM2NEk72FPcapsResponseSimpleTypeSimple PcapGetResponseM2NEk72FPcapsResponseSimpleType = "simple"
	PcapGetResponseM2NEk72FPcapsResponseSimpleTypeFull   PcapGetResponseM2NEk72FPcapsResponseSimpleType = "full"
)

type PcapGetResponseM2NEk72FPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseM2NEk72FPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseM2NEk72FPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseM2NEk72FPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseM2NEk72FPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseM2NEk72FPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseM2NEk72FPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseM2NEk72FPcapsResponseFull]
type pcapGetResponseM2NEk72FPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseM2NEk72FPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseM2NEk72FPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseM2NEk72FPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseM2NEk72FPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseM2NEk72FPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseM2NEk72FPcapsResponseFullFilterV1]
type pcapGetResponseM2NEk72FPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseM2NEk72FPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseM2NEk72FPcapsResponseFullStatus string

const (
	PcapGetResponseM2NEk72FPcapsResponseFullStatusUnknown           PcapGetResponseM2NEk72FPcapsResponseFullStatus = "unknown"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusSuccess           PcapGetResponseM2NEk72FPcapsResponseFullStatus = "success"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusPending           PcapGetResponseM2NEk72FPcapsResponseFullStatus = "pending"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusRunning           PcapGetResponseM2NEk72FPcapsResponseFullStatus = "running"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusConversionPending PcapGetResponseM2NEk72FPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusConversionRunning PcapGetResponseM2NEk72FPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusComplete          PcapGetResponseM2NEk72FPcapsResponseFullStatus = "complete"
	PcapGetResponseM2NEk72FPcapsResponseFullStatusFailed            PcapGetResponseM2NEk72FPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseM2NEk72FPcapsResponseFullSystem string

const (
	PcapGetResponseM2NEk72FPcapsResponseFullSystemMagicTransit PcapGetResponseM2NEk72FPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseM2NEk72FPcapsResponseFullType string

const (
	PcapGetResponseM2NEk72FPcapsResponseFullTypeSimple PcapGetResponseM2NEk72FPcapsResponseFullType = "simple"
	PcapGetResponseM2NEk72FPcapsResponseFullTypeFull   PcapGetResponseM2NEk72FPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsM2NEk72FPcapsRequestSimple],
// [PcapNewParamsM2NEk72FPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsM2NEk72FPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsM2NEk72FPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsM2NEk72FPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsM2NEk72FPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsM2NEk72FPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsM2NEk72FPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsM2NEk72FPcapsRequestSimpleSystem string

const (
	PcapNewParamsM2NEk72FPcapsRequestSimpleSystemMagicTransit PcapNewParamsM2NEk72FPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsM2NEk72FPcapsRequestSimpleType string

const (
	PcapNewParamsM2NEk72FPcapsRequestSimpleTypeSimple PcapNewParamsM2NEk72FPcapsRequestSimpleType = "simple"
	PcapNewParamsM2NEk72FPcapsRequestSimpleTypeFull   PcapNewParamsM2NEk72FPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsM2NEk72FPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsM2NEk72FPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsM2NEk72FPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsM2NEk72FPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsM2NEk72FPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsM2NEk72FPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsM2NEk72FPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsM2NEk72FPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsM2NEk72FPcapsRequestFullSystem string

const (
	PcapNewParamsM2NEk72FPcapsRequestFullSystemMagicTransit PcapNewParamsM2NEk72FPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsM2NEk72FPcapsRequestFullType string

const (
	PcapNewParamsM2NEk72FPcapsRequestFullTypeSimple PcapNewParamsM2NEk72FPcapsRequestFullType = "simple"
	PcapNewParamsM2NEk72FPcapsRequestFullTypeFull   PcapNewParamsM2NEk72FPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsM2NEk72FPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsM2NEk72FPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
