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

// Union satisfied by [PcapNewResponse2Uqz5PlfPcapsResponseSimple] or
// [PcapNewResponse2Uqz5PlfPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse2Uqz5PlfPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2Uqz5PlfPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2Uqz5PlfPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse2Uqz5PlfPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse2Uqz5PlfPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse2Uqz5PlfPcapsResponseSimple]
type pcapNewResponse2Uqz5PlfPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse2Uqz5PlfPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2Uqz5PlfPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1]
type pcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2Uqz5PlfPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusUnknown           PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusSuccess           PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "success"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusPending           PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "pending"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusRunning           PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "running"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusConversionPending PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusConversionRunning PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusComplete          PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "complete"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatusFailed            PcapNewResponse2Uqz5PlfPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2Uqz5PlfPcapsResponseSimpleSystem string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleSystemMagicTransit PcapNewResponse2Uqz5PlfPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2Uqz5PlfPcapsResponseSimpleType string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleTypeSimple PcapNewResponse2Uqz5PlfPcapsResponseSimpleType = "simple"
	PcapNewResponse2Uqz5PlfPcapsResponseSimpleTypeFull   PcapNewResponse2Uqz5PlfPcapsResponseSimpleType = "full"
)

type PcapNewResponse2Uqz5PlfPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2Uqz5PlfPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2Uqz5PlfPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2Uqz5PlfPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse2Uqz5PlfPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse2Uqz5PlfPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse2Uqz5PlfPcapsResponseFull]
type pcapNewResponse2Uqz5PlfPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse2Uqz5PlfPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2Uqz5PlfPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1]
type pcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2Uqz5PlfPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2Uqz5PlfPcapsResponseFullStatus string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusUnknown           PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "unknown"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusSuccess           PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "success"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusPending           PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "pending"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusRunning           PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "running"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusConversionPending PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusConversionRunning PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusComplete          PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "complete"
	PcapNewResponse2Uqz5PlfPcapsResponseFullStatusFailed            PcapNewResponse2Uqz5PlfPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2Uqz5PlfPcapsResponseFullSystem string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseFullSystemMagicTransit PcapNewResponse2Uqz5PlfPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2Uqz5PlfPcapsResponseFullType string

const (
	PcapNewResponse2Uqz5PlfPcapsResponseFullTypeSimple PcapNewResponse2Uqz5PlfPcapsResponseFullType = "simple"
	PcapNewResponse2Uqz5PlfPcapsResponseFullTypeFull   PcapNewResponse2Uqz5PlfPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse2Uqz5PlfPcapsResponseSimple] or
// [PcapListResponse2Uqz5PlfPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse2Uqz5PlfPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2Uqz5PlfPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2Uqz5PlfPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse2Uqz5PlfPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse2Uqz5PlfPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse2Uqz5PlfPcapsResponseSimple]
type pcapListResponse2Uqz5PlfPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse2Uqz5PlfPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2Uqz5PlfPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1]
type pcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2Uqz5PlfPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus string

const (
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusUnknown           PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "unknown"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusSuccess           PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "success"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusPending           PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "pending"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusRunning           PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "running"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusConversionPending PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusConversionRunning PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusComplete          PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "complete"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleStatusFailed            PcapListResponse2Uqz5PlfPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2Uqz5PlfPcapsResponseSimpleSystem string

const (
	PcapListResponse2Uqz5PlfPcapsResponseSimpleSystemMagicTransit PcapListResponse2Uqz5PlfPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2Uqz5PlfPcapsResponseSimpleType string

const (
	PcapListResponse2Uqz5PlfPcapsResponseSimpleTypeSimple PcapListResponse2Uqz5PlfPcapsResponseSimpleType = "simple"
	PcapListResponse2Uqz5PlfPcapsResponseSimpleTypeFull   PcapListResponse2Uqz5PlfPcapsResponseSimpleType = "full"
)

type PcapListResponse2Uqz5PlfPcapsResponseFull struct {
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
	FilterV1 PcapListResponse2Uqz5PlfPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2Uqz5PlfPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2Uqz5PlfPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2Uqz5PlfPcapsResponseFullType `json:"type"`
	JSON pcapListResponse2Uqz5PlfPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse2Uqz5PlfPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse2Uqz5PlfPcapsResponseFull]
type pcapListResponse2Uqz5PlfPcapsResponseFullJSON struct {
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

func (r *PcapListResponse2Uqz5PlfPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2Uqz5PlfPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2Uqz5PlfPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse2Uqz5PlfPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse2Uqz5PlfPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse2Uqz5PlfPcapsResponseFullFilterV1]
type pcapListResponse2Uqz5PlfPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2Uqz5PlfPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2Uqz5PlfPcapsResponseFullStatus string

const (
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusUnknown           PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "unknown"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusSuccess           PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "success"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusPending           PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "pending"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusRunning           PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "running"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusConversionPending PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusConversionRunning PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_running"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusComplete          PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "complete"
	PcapListResponse2Uqz5PlfPcapsResponseFullStatusFailed            PcapListResponse2Uqz5PlfPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2Uqz5PlfPcapsResponseFullSystem string

const (
	PcapListResponse2Uqz5PlfPcapsResponseFullSystemMagicTransit PcapListResponse2Uqz5PlfPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2Uqz5PlfPcapsResponseFullType string

const (
	PcapListResponse2Uqz5PlfPcapsResponseFullTypeSimple PcapListResponse2Uqz5PlfPcapsResponseFullType = "simple"
	PcapListResponse2Uqz5PlfPcapsResponseFullTypeFull   PcapListResponse2Uqz5PlfPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse2Uqz5PlfPcapsResponseSimple] or
// [PcapGetResponse2Uqz5PlfPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse2Uqz5PlfPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2Uqz5PlfPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2Uqz5PlfPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse2Uqz5PlfPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse2Uqz5PlfPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse2Uqz5PlfPcapsResponseSimple]
type pcapGetResponse2Uqz5PlfPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse2Uqz5PlfPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2Uqz5PlfPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1]
type pcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2Uqz5PlfPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusUnknown           PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusSuccess           PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "success"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusPending           PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "pending"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusRunning           PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "running"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusConversionPending PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusConversionRunning PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusComplete          PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "complete"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatusFailed            PcapGetResponse2Uqz5PlfPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2Uqz5PlfPcapsResponseSimpleSystem string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleSystemMagicTransit PcapGetResponse2Uqz5PlfPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2Uqz5PlfPcapsResponseSimpleType string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleTypeSimple PcapGetResponse2Uqz5PlfPcapsResponseSimpleType = "simple"
	PcapGetResponse2Uqz5PlfPcapsResponseSimpleTypeFull   PcapGetResponse2Uqz5PlfPcapsResponseSimpleType = "full"
)

type PcapGetResponse2Uqz5PlfPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2Uqz5PlfPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2Uqz5PlfPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2Uqz5PlfPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse2Uqz5PlfPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse2Uqz5PlfPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse2Uqz5PlfPcapsResponseFull]
type pcapGetResponse2Uqz5PlfPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse2Uqz5PlfPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2Uqz5PlfPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1]
type pcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2Uqz5PlfPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2Uqz5PlfPcapsResponseFullStatus string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusUnknown           PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "unknown"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusSuccess           PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "success"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusPending           PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "pending"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusRunning           PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "running"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusConversionPending PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusConversionRunning PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusComplete          PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "complete"
	PcapGetResponse2Uqz5PlfPcapsResponseFullStatusFailed            PcapGetResponse2Uqz5PlfPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2Uqz5PlfPcapsResponseFullSystem string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseFullSystemMagicTransit PcapGetResponse2Uqz5PlfPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2Uqz5PlfPcapsResponseFullType string

const (
	PcapGetResponse2Uqz5PlfPcapsResponseFullTypeSimple PcapGetResponse2Uqz5PlfPcapsResponseFullType = "simple"
	PcapGetResponse2Uqz5PlfPcapsResponseFullTypeFull   PcapGetResponse2Uqz5PlfPcapsResponseFullType = "full"
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
