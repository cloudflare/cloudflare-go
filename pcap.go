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

// Union satisfied by [PcapNewResponseHlEvaQ18PcapsResponseSimple] or
// [PcapNewResponseHlEvaQ18PcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseHlEvaQ18PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseHlEvaQ18PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseHlEvaQ18PcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseHlEvaQ18PcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseHlEvaQ18PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseHlEvaQ18PcapsResponseSimple]
type pcapNewResponseHlEvaQ18PcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseHlEvaQ18PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseHlEvaQ18PcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1]
type pcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseHlEvaQ18PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus string

const (
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusUnknown           PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "unknown"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusSuccess           PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "success"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusPending           PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "pending"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusRunning           PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "running"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusConversionPending PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusConversionRunning PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusComplete          PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "complete"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleStatusFailed            PcapNewResponseHlEvaQ18PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseHlEvaQ18PcapsResponseSimpleSystem string

const (
	PcapNewResponseHlEvaQ18PcapsResponseSimpleSystemMagicTransit PcapNewResponseHlEvaQ18PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseHlEvaQ18PcapsResponseSimpleType string

const (
	PcapNewResponseHlEvaQ18PcapsResponseSimpleTypeSimple PcapNewResponseHlEvaQ18PcapsResponseSimpleType = "simple"
	PcapNewResponseHlEvaQ18PcapsResponseSimpleTypeFull   PcapNewResponseHlEvaQ18PcapsResponseSimpleType = "full"
)

type PcapNewResponseHlEvaQ18PcapsResponseFull struct {
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
	FilterV1 PcapNewResponseHlEvaQ18PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseHlEvaQ18PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseHlEvaQ18PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseHlEvaQ18PcapsResponseFullType `json:"type"`
	JSON pcapNewResponseHlEvaQ18PcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseHlEvaQ18PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseHlEvaQ18PcapsResponseFull]
type pcapNewResponseHlEvaQ18PcapsResponseFullJSON struct {
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

func (r *PcapNewResponseHlEvaQ18PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseHlEvaQ18PcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseHlEvaQ18PcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseHlEvaQ18PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseHlEvaQ18PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseHlEvaQ18PcapsResponseFullFilterV1]
type pcapNewResponseHlEvaQ18PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseHlEvaQ18PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseHlEvaQ18PcapsResponseFullStatus string

const (
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusUnknown           PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "unknown"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusSuccess           PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "success"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusPending           PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "pending"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusRunning           PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "running"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusConversionPending PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusConversionRunning PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "conversion_running"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusComplete          PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "complete"
	PcapNewResponseHlEvaQ18PcapsResponseFullStatusFailed            PcapNewResponseHlEvaQ18PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseHlEvaQ18PcapsResponseFullSystem string

const (
	PcapNewResponseHlEvaQ18PcapsResponseFullSystemMagicTransit PcapNewResponseHlEvaQ18PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseHlEvaQ18PcapsResponseFullType string

const (
	PcapNewResponseHlEvaQ18PcapsResponseFullTypeSimple PcapNewResponseHlEvaQ18PcapsResponseFullType = "simple"
	PcapNewResponseHlEvaQ18PcapsResponseFullTypeFull   PcapNewResponseHlEvaQ18PcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseHlEvaQ18PcapsResponseSimple] or
// [PcapListResponseHlEvaQ18PcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseHlEvaQ18PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseHlEvaQ18PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseHlEvaQ18PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseHlEvaQ18PcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseHlEvaQ18PcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseHlEvaQ18PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseHlEvaQ18PcapsResponseSimple]
type pcapListResponseHlEvaQ18PcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseHlEvaQ18PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseHlEvaQ18PcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1]
type pcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseHlEvaQ18PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseHlEvaQ18PcapsResponseSimpleStatus string

const (
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusUnknown           PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "unknown"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusSuccess           PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "success"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusPending           PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "pending"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusRunning           PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "running"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusConversionPending PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusConversionRunning PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusComplete          PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "complete"
	PcapListResponseHlEvaQ18PcapsResponseSimpleStatusFailed            PcapListResponseHlEvaQ18PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseHlEvaQ18PcapsResponseSimpleSystem string

const (
	PcapListResponseHlEvaQ18PcapsResponseSimpleSystemMagicTransit PcapListResponseHlEvaQ18PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseHlEvaQ18PcapsResponseSimpleType string

const (
	PcapListResponseHlEvaQ18PcapsResponseSimpleTypeSimple PcapListResponseHlEvaQ18PcapsResponseSimpleType = "simple"
	PcapListResponseHlEvaQ18PcapsResponseSimpleTypeFull   PcapListResponseHlEvaQ18PcapsResponseSimpleType = "full"
)

type PcapListResponseHlEvaQ18PcapsResponseFull struct {
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
	FilterV1 PcapListResponseHlEvaQ18PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseHlEvaQ18PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseHlEvaQ18PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseHlEvaQ18PcapsResponseFullType `json:"type"`
	JSON pcapListResponseHlEvaQ18PcapsResponseFullJSON `json:"-"`
}

// pcapListResponseHlEvaQ18PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseHlEvaQ18PcapsResponseFull]
type pcapListResponseHlEvaQ18PcapsResponseFullJSON struct {
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

func (r *PcapListResponseHlEvaQ18PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseHlEvaQ18PcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseHlEvaQ18PcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseHlEvaQ18PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseHlEvaQ18PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseHlEvaQ18PcapsResponseFullFilterV1]
type pcapListResponseHlEvaQ18PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseHlEvaQ18PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseHlEvaQ18PcapsResponseFullStatus string

const (
	PcapListResponseHlEvaQ18PcapsResponseFullStatusUnknown           PcapListResponseHlEvaQ18PcapsResponseFullStatus = "unknown"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusSuccess           PcapListResponseHlEvaQ18PcapsResponseFullStatus = "success"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusPending           PcapListResponseHlEvaQ18PcapsResponseFullStatus = "pending"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusRunning           PcapListResponseHlEvaQ18PcapsResponseFullStatus = "running"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusConversionPending PcapListResponseHlEvaQ18PcapsResponseFullStatus = "conversion_pending"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusConversionRunning PcapListResponseHlEvaQ18PcapsResponseFullStatus = "conversion_running"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusComplete          PcapListResponseHlEvaQ18PcapsResponseFullStatus = "complete"
	PcapListResponseHlEvaQ18PcapsResponseFullStatusFailed            PcapListResponseHlEvaQ18PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseHlEvaQ18PcapsResponseFullSystem string

const (
	PcapListResponseHlEvaQ18PcapsResponseFullSystemMagicTransit PcapListResponseHlEvaQ18PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseHlEvaQ18PcapsResponseFullType string

const (
	PcapListResponseHlEvaQ18PcapsResponseFullTypeSimple PcapListResponseHlEvaQ18PcapsResponseFullType = "simple"
	PcapListResponseHlEvaQ18PcapsResponseFullTypeFull   PcapListResponseHlEvaQ18PcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseHlEvaQ18PcapsResponseSimple] or
// [PcapGetResponseHlEvaQ18PcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseHlEvaQ18PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseHlEvaQ18PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseHlEvaQ18PcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseHlEvaQ18PcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseHlEvaQ18PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseHlEvaQ18PcapsResponseSimple]
type pcapGetResponseHlEvaQ18PcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseHlEvaQ18PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseHlEvaQ18PcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1]
type pcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseHlEvaQ18PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus string

const (
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusUnknown           PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "unknown"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusSuccess           PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "success"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusPending           PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "pending"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusRunning           PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "running"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusConversionPending PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusConversionRunning PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusComplete          PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "complete"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleStatusFailed            PcapGetResponseHlEvaQ18PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseHlEvaQ18PcapsResponseSimpleSystem string

const (
	PcapGetResponseHlEvaQ18PcapsResponseSimpleSystemMagicTransit PcapGetResponseHlEvaQ18PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseHlEvaQ18PcapsResponseSimpleType string

const (
	PcapGetResponseHlEvaQ18PcapsResponseSimpleTypeSimple PcapGetResponseHlEvaQ18PcapsResponseSimpleType = "simple"
	PcapGetResponseHlEvaQ18PcapsResponseSimpleTypeFull   PcapGetResponseHlEvaQ18PcapsResponseSimpleType = "full"
)

type PcapGetResponseHlEvaQ18PcapsResponseFull struct {
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
	FilterV1 PcapGetResponseHlEvaQ18PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseHlEvaQ18PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseHlEvaQ18PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseHlEvaQ18PcapsResponseFullType `json:"type"`
	JSON pcapGetResponseHlEvaQ18PcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseHlEvaQ18PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseHlEvaQ18PcapsResponseFull]
type pcapGetResponseHlEvaQ18PcapsResponseFullJSON struct {
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

func (r *PcapGetResponseHlEvaQ18PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseHlEvaQ18PcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseHlEvaQ18PcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseHlEvaQ18PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseHlEvaQ18PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseHlEvaQ18PcapsResponseFullFilterV1]
type pcapGetResponseHlEvaQ18PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseHlEvaQ18PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseHlEvaQ18PcapsResponseFullStatus string

const (
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusUnknown           PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "unknown"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusSuccess           PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "success"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusPending           PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "pending"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusRunning           PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "running"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusConversionPending PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusConversionRunning PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "conversion_running"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusComplete          PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "complete"
	PcapGetResponseHlEvaQ18PcapsResponseFullStatusFailed            PcapGetResponseHlEvaQ18PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseHlEvaQ18PcapsResponseFullSystem string

const (
	PcapGetResponseHlEvaQ18PcapsResponseFullSystemMagicTransit PcapGetResponseHlEvaQ18PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseHlEvaQ18PcapsResponseFullType string

const (
	PcapGetResponseHlEvaQ18PcapsResponseFullTypeSimple PcapGetResponseHlEvaQ18PcapsResponseFullType = "simple"
	PcapGetResponseHlEvaQ18PcapsResponseFullTypeFull   PcapGetResponseHlEvaQ18PcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsHlEvaQ18PcapsRequestSimple],
// [PcapNewParamsHlEvaQ18PcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsHlEvaQ18PcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsHlEvaQ18PcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsHlEvaQ18PcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsHlEvaQ18PcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsHlEvaQ18PcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsHlEvaQ18PcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsHlEvaQ18PcapsRequestSimpleSystem string

const (
	PcapNewParamsHlEvaQ18PcapsRequestSimpleSystemMagicTransit PcapNewParamsHlEvaQ18PcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsHlEvaQ18PcapsRequestSimpleType string

const (
	PcapNewParamsHlEvaQ18PcapsRequestSimpleTypeSimple PcapNewParamsHlEvaQ18PcapsRequestSimpleType = "simple"
	PcapNewParamsHlEvaQ18PcapsRequestSimpleTypeFull   PcapNewParamsHlEvaQ18PcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsHlEvaQ18PcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsHlEvaQ18PcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsHlEvaQ18PcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsHlEvaQ18PcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsHlEvaQ18PcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsHlEvaQ18PcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsHlEvaQ18PcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsHlEvaQ18PcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsHlEvaQ18PcapsRequestFullSystem string

const (
	PcapNewParamsHlEvaQ18PcapsRequestFullSystemMagicTransit PcapNewParamsHlEvaQ18PcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsHlEvaQ18PcapsRequestFullType string

const (
	PcapNewParamsHlEvaQ18PcapsRequestFullTypeSimple PcapNewParamsHlEvaQ18PcapsRequestFullType = "simple"
	PcapNewParamsHlEvaQ18PcapsRequestFullTypeFull   PcapNewParamsHlEvaQ18PcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsHlEvaQ18PcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsHlEvaQ18PcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
