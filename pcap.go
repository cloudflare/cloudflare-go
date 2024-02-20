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

// Union satisfied by [PcapNewResponse0F6lSEpTPcapsResponseSimple] or
// [PcapNewResponse0F6lSEpTPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse0F6lSEpTPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse0F6lSEpTPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse0F6lSEpTPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse0F6lSEpTPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse0F6lSEpTPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse0F6lSEpTPcapsResponseSimple]
type pcapNewResponse0F6lSEpTPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse0F6lSEpTPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse0F6lSEpTPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1]
type pcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse0F6lSEpTPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus string

const (
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusUnknown           PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusSuccess           PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "success"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusPending           PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "pending"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusRunning           PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "running"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusConversionPending PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusConversionRunning PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusComplete          PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "complete"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleStatusFailed            PcapNewResponse0F6lSEpTPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse0F6lSEpTPcapsResponseSimpleSystem string

const (
	PcapNewResponse0F6lSEpTPcapsResponseSimpleSystemMagicTransit PcapNewResponse0F6lSEpTPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse0F6lSEpTPcapsResponseSimpleType string

const (
	PcapNewResponse0F6lSEpTPcapsResponseSimpleTypeSimple PcapNewResponse0F6lSEpTPcapsResponseSimpleType = "simple"
	PcapNewResponse0F6lSEpTPcapsResponseSimpleTypeFull   PcapNewResponse0F6lSEpTPcapsResponseSimpleType = "full"
)

type PcapNewResponse0F6lSEpTPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse0F6lSEpTPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse0F6lSEpTPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse0F6lSEpTPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse0F6lSEpTPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse0F6lSEpTPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse0F6lSEpTPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse0F6lSEpTPcapsResponseFull]
type pcapNewResponse0F6lSEpTPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse0F6lSEpTPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse0F6lSEpTPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse0F6lSEpTPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse0F6lSEpTPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse0F6lSEpTPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse0F6lSEpTPcapsResponseFullFilterV1]
type pcapNewResponse0F6lSEpTPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse0F6lSEpTPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse0F6lSEpTPcapsResponseFullStatus string

const (
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusUnknown           PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "unknown"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusSuccess           PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "success"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusPending           PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "pending"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusRunning           PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "running"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusConversionPending PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusConversionRunning PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusComplete          PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "complete"
	PcapNewResponse0F6lSEpTPcapsResponseFullStatusFailed            PcapNewResponse0F6lSEpTPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse0F6lSEpTPcapsResponseFullSystem string

const (
	PcapNewResponse0F6lSEpTPcapsResponseFullSystemMagicTransit PcapNewResponse0F6lSEpTPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse0F6lSEpTPcapsResponseFullType string

const (
	PcapNewResponse0F6lSEpTPcapsResponseFullTypeSimple PcapNewResponse0F6lSEpTPcapsResponseFullType = "simple"
	PcapNewResponse0F6lSEpTPcapsResponseFullTypeFull   PcapNewResponse0F6lSEpTPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse0F6lSEpTPcapsResponseSimple] or
// [PcapListResponse0F6lSEpTPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse0F6lSEpTPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse0F6lSEpTPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse0F6lSEpTPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse0F6lSEpTPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse0F6lSEpTPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse0F6lSEpTPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse0F6lSEpTPcapsResponseSimple]
type pcapListResponse0F6lSEpTPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse0F6lSEpTPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse0F6lSEpTPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1]
type pcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse0F6lSEpTPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse0F6lSEpTPcapsResponseSimpleStatus string

const (
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusUnknown           PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "unknown"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusSuccess           PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "success"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusPending           PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "pending"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusRunning           PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "running"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusConversionPending PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusConversionRunning PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusComplete          PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "complete"
	PcapListResponse0F6lSEpTPcapsResponseSimpleStatusFailed            PcapListResponse0F6lSEpTPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse0F6lSEpTPcapsResponseSimpleSystem string

const (
	PcapListResponse0F6lSEpTPcapsResponseSimpleSystemMagicTransit PcapListResponse0F6lSEpTPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse0F6lSEpTPcapsResponseSimpleType string

const (
	PcapListResponse0F6lSEpTPcapsResponseSimpleTypeSimple PcapListResponse0F6lSEpTPcapsResponseSimpleType = "simple"
	PcapListResponse0F6lSEpTPcapsResponseSimpleTypeFull   PcapListResponse0F6lSEpTPcapsResponseSimpleType = "full"
)

type PcapListResponse0F6lSEpTPcapsResponseFull struct {
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
	FilterV1 PcapListResponse0F6lSEpTPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse0F6lSEpTPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse0F6lSEpTPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse0F6lSEpTPcapsResponseFullType `json:"type"`
	JSON pcapListResponse0F6lSEpTPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse0F6lSEpTPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse0F6lSEpTPcapsResponseFull]
type pcapListResponse0F6lSEpTPcapsResponseFullJSON struct {
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

func (r *PcapListResponse0F6lSEpTPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse0F6lSEpTPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse0F6lSEpTPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse0F6lSEpTPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse0F6lSEpTPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse0F6lSEpTPcapsResponseFullFilterV1]
type pcapListResponse0F6lSEpTPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse0F6lSEpTPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse0F6lSEpTPcapsResponseFullStatus string

const (
	PcapListResponse0F6lSEpTPcapsResponseFullStatusUnknown           PcapListResponse0F6lSEpTPcapsResponseFullStatus = "unknown"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusSuccess           PcapListResponse0F6lSEpTPcapsResponseFullStatus = "success"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusPending           PcapListResponse0F6lSEpTPcapsResponseFullStatus = "pending"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusRunning           PcapListResponse0F6lSEpTPcapsResponseFullStatus = "running"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusConversionPending PcapListResponse0F6lSEpTPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusConversionRunning PcapListResponse0F6lSEpTPcapsResponseFullStatus = "conversion_running"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusComplete          PcapListResponse0F6lSEpTPcapsResponseFullStatus = "complete"
	PcapListResponse0F6lSEpTPcapsResponseFullStatusFailed            PcapListResponse0F6lSEpTPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse0F6lSEpTPcapsResponseFullSystem string

const (
	PcapListResponse0F6lSEpTPcapsResponseFullSystemMagicTransit PcapListResponse0F6lSEpTPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse0F6lSEpTPcapsResponseFullType string

const (
	PcapListResponse0F6lSEpTPcapsResponseFullTypeSimple PcapListResponse0F6lSEpTPcapsResponseFullType = "simple"
	PcapListResponse0F6lSEpTPcapsResponseFullTypeFull   PcapListResponse0F6lSEpTPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse0F6lSEpTPcapsResponseSimple] or
// [PcapGetResponse0F6lSEpTPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse0F6lSEpTPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse0F6lSEpTPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse0F6lSEpTPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse0F6lSEpTPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse0F6lSEpTPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse0F6lSEpTPcapsResponseSimple]
type pcapGetResponse0F6lSEpTPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse0F6lSEpTPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse0F6lSEpTPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1]
type pcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse0F6lSEpTPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus string

const (
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusUnknown           PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusSuccess           PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "success"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusPending           PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "pending"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusRunning           PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "running"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusConversionPending PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusConversionRunning PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusComplete          PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "complete"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleStatusFailed            PcapGetResponse0F6lSEpTPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse0F6lSEpTPcapsResponseSimpleSystem string

const (
	PcapGetResponse0F6lSEpTPcapsResponseSimpleSystemMagicTransit PcapGetResponse0F6lSEpTPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse0F6lSEpTPcapsResponseSimpleType string

const (
	PcapGetResponse0F6lSEpTPcapsResponseSimpleTypeSimple PcapGetResponse0F6lSEpTPcapsResponseSimpleType = "simple"
	PcapGetResponse0F6lSEpTPcapsResponseSimpleTypeFull   PcapGetResponse0F6lSEpTPcapsResponseSimpleType = "full"
)

type PcapGetResponse0F6lSEpTPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse0F6lSEpTPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse0F6lSEpTPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse0F6lSEpTPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse0F6lSEpTPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse0F6lSEpTPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse0F6lSEpTPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse0F6lSEpTPcapsResponseFull]
type pcapGetResponse0F6lSEpTPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse0F6lSEpTPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse0F6lSEpTPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse0F6lSEpTPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse0F6lSEpTPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse0F6lSEpTPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse0F6lSEpTPcapsResponseFullFilterV1]
type pcapGetResponse0F6lSEpTPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse0F6lSEpTPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse0F6lSEpTPcapsResponseFullStatus string

const (
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusUnknown           PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "unknown"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusSuccess           PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "success"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusPending           PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "pending"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusRunning           PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "running"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusConversionPending PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusConversionRunning PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusComplete          PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "complete"
	PcapGetResponse0F6lSEpTPcapsResponseFullStatusFailed            PcapGetResponse0F6lSEpTPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse0F6lSEpTPcapsResponseFullSystem string

const (
	PcapGetResponse0F6lSEpTPcapsResponseFullSystemMagicTransit PcapGetResponse0F6lSEpTPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse0F6lSEpTPcapsResponseFullType string

const (
	PcapGetResponse0F6lSEpTPcapsResponseFullTypeSimple PcapGetResponse0F6lSEpTPcapsResponseFullType = "simple"
	PcapGetResponse0F6lSEpTPcapsResponseFullTypeFull   PcapGetResponse0F6lSEpTPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParams0F6lSEpTPcapsRequestSimple],
// [PcapNewParams0F6lSEpTPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParams0F6lSEpTPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams0F6lSEpTPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams0F6lSEpTPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams0F6lSEpTPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParams0F6lSEpTPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams0F6lSEpTPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams0F6lSEpTPcapsRequestSimpleSystem string

const (
	PcapNewParams0F6lSEpTPcapsRequestSimpleSystemMagicTransit PcapNewParams0F6lSEpTPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams0F6lSEpTPcapsRequestSimpleType string

const (
	PcapNewParams0F6lSEpTPcapsRequestSimpleTypeSimple PcapNewParams0F6lSEpTPcapsRequestSimpleType = "simple"
	PcapNewParams0F6lSEpTPcapsRequestSimpleTypeFull   PcapNewParams0F6lSEpTPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams0F6lSEpTPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParams0F6lSEpTPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParams0F6lSEpTPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams0F6lSEpTPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams0F6lSEpTPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams0F6lSEpTPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParams0F6lSEpTPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams0F6lSEpTPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams0F6lSEpTPcapsRequestFullSystem string

const (
	PcapNewParams0F6lSEpTPcapsRequestFullSystemMagicTransit PcapNewParams0F6lSEpTPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams0F6lSEpTPcapsRequestFullType string

const (
	PcapNewParams0F6lSEpTPcapsRequestFullTypeSimple PcapNewParams0F6lSEpTPcapsRequestFullType = "simple"
	PcapNewParams0F6lSEpTPcapsRequestFullTypeFull   PcapNewParams0F6lSEpTPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams0F6lSEpTPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParams0F6lSEpTPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
