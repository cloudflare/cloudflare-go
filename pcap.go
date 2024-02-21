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

// Union satisfied by [PcapNewResponse7Ks0f1TyPcapsResponseSimple] or
// [PcapNewResponse7Ks0f1TyPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse7Ks0f1TyPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse7Ks0f1TyPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse7Ks0f1TyPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse7Ks0f1TyPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse7Ks0f1TyPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse7Ks0f1TyPcapsResponseSimple]
type pcapNewResponse7Ks0f1TyPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse7Ks0f1TyPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse7Ks0f1TyPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1]
type pcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse7Ks0f1TyPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusUnknown           PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusSuccess           PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "success"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusPending           PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "pending"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusRunning           PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "running"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusConversionPending PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusConversionRunning PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusComplete          PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "complete"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatusFailed            PcapNewResponse7Ks0f1TyPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse7Ks0f1TyPcapsResponseSimpleSystem string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleSystemMagicTransit PcapNewResponse7Ks0f1TyPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse7Ks0f1TyPcapsResponseSimpleType string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleTypeSimple PcapNewResponse7Ks0f1TyPcapsResponseSimpleType = "simple"
	PcapNewResponse7Ks0f1TyPcapsResponseSimpleTypeFull   PcapNewResponse7Ks0f1TyPcapsResponseSimpleType = "full"
)

type PcapNewResponse7Ks0f1TyPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse7Ks0f1TyPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse7Ks0f1TyPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse7Ks0f1TyPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse7Ks0f1TyPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse7Ks0f1TyPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse7Ks0f1TyPcapsResponseFull]
type pcapNewResponse7Ks0f1TyPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse7Ks0f1TyPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse7Ks0f1TyPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1]
type pcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse7Ks0f1TyPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse7Ks0f1TyPcapsResponseFullStatus string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusUnknown           PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "unknown"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusSuccess           PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "success"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusPending           PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "pending"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusRunning           PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "running"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusConversionPending PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusConversionRunning PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusComplete          PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "complete"
	PcapNewResponse7Ks0f1TyPcapsResponseFullStatusFailed            PcapNewResponse7Ks0f1TyPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse7Ks0f1TyPcapsResponseFullSystem string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseFullSystemMagicTransit PcapNewResponse7Ks0f1TyPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse7Ks0f1TyPcapsResponseFullType string

const (
	PcapNewResponse7Ks0f1TyPcapsResponseFullTypeSimple PcapNewResponse7Ks0f1TyPcapsResponseFullType = "simple"
	PcapNewResponse7Ks0f1TyPcapsResponseFullTypeFull   PcapNewResponse7Ks0f1TyPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse7Ks0f1TyPcapsResponseSimple] or
// [PcapListResponse7Ks0f1TyPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse7Ks0f1TyPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse7Ks0f1TyPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse7Ks0f1TyPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse7Ks0f1TyPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse7Ks0f1TyPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse7Ks0f1TyPcapsResponseSimple]
type pcapListResponse7Ks0f1TyPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse7Ks0f1TyPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse7Ks0f1TyPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1]
type pcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse7Ks0f1TyPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus string

const (
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusUnknown           PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "unknown"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusSuccess           PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "success"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusPending           PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "pending"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusRunning           PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "running"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusConversionPending PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusConversionRunning PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusComplete          PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "complete"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleStatusFailed            PcapListResponse7Ks0f1TyPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse7Ks0f1TyPcapsResponseSimpleSystem string

const (
	PcapListResponse7Ks0f1TyPcapsResponseSimpleSystemMagicTransit PcapListResponse7Ks0f1TyPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse7Ks0f1TyPcapsResponseSimpleType string

const (
	PcapListResponse7Ks0f1TyPcapsResponseSimpleTypeSimple PcapListResponse7Ks0f1TyPcapsResponseSimpleType = "simple"
	PcapListResponse7Ks0f1TyPcapsResponseSimpleTypeFull   PcapListResponse7Ks0f1TyPcapsResponseSimpleType = "full"
)

type PcapListResponse7Ks0f1TyPcapsResponseFull struct {
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
	FilterV1 PcapListResponse7Ks0f1TyPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse7Ks0f1TyPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse7Ks0f1TyPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse7Ks0f1TyPcapsResponseFullType `json:"type"`
	JSON pcapListResponse7Ks0f1TyPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse7Ks0f1TyPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse7Ks0f1TyPcapsResponseFull]
type pcapListResponse7Ks0f1TyPcapsResponseFullJSON struct {
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

func (r *PcapListResponse7Ks0f1TyPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse7Ks0f1TyPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse7Ks0f1TyPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse7Ks0f1TyPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse7Ks0f1TyPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse7Ks0f1TyPcapsResponseFullFilterV1]
type pcapListResponse7Ks0f1TyPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse7Ks0f1TyPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse7Ks0f1TyPcapsResponseFullStatus string

const (
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusUnknown           PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "unknown"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusSuccess           PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "success"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusPending           PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "pending"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusRunning           PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "running"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusConversionPending PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusConversionRunning PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_running"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusComplete          PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "complete"
	PcapListResponse7Ks0f1TyPcapsResponseFullStatusFailed            PcapListResponse7Ks0f1TyPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse7Ks0f1TyPcapsResponseFullSystem string

const (
	PcapListResponse7Ks0f1TyPcapsResponseFullSystemMagicTransit PcapListResponse7Ks0f1TyPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse7Ks0f1TyPcapsResponseFullType string

const (
	PcapListResponse7Ks0f1TyPcapsResponseFullTypeSimple PcapListResponse7Ks0f1TyPcapsResponseFullType = "simple"
	PcapListResponse7Ks0f1TyPcapsResponseFullTypeFull   PcapListResponse7Ks0f1TyPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse7Ks0f1TyPcapsResponseSimple] or
// [PcapGetResponse7Ks0f1TyPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse7Ks0f1TyPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7Ks0f1TyPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7Ks0f1TyPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse7Ks0f1TyPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse7Ks0f1TyPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse7Ks0f1TyPcapsResponseSimple]
type pcapGetResponse7Ks0f1TyPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse7Ks0f1TyPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7Ks0f1TyPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1]
type pcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7Ks0f1TyPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusUnknown           PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusSuccess           PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "success"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusPending           PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "pending"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusRunning           PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "running"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusConversionPending PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusConversionRunning PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusComplete          PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "complete"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatusFailed            PcapGetResponse7Ks0f1TyPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7Ks0f1TyPcapsResponseSimpleSystem string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleSystemMagicTransit PcapGetResponse7Ks0f1TyPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7Ks0f1TyPcapsResponseSimpleType string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleTypeSimple PcapGetResponse7Ks0f1TyPcapsResponseSimpleType = "simple"
	PcapGetResponse7Ks0f1TyPcapsResponseSimpleTypeFull   PcapGetResponse7Ks0f1TyPcapsResponseSimpleType = "full"
)

type PcapGetResponse7Ks0f1TyPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse7Ks0f1TyPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse7Ks0f1TyPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse7Ks0f1TyPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse7Ks0f1TyPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse7Ks0f1TyPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse7Ks0f1TyPcapsResponseFull]
type pcapGetResponse7Ks0f1TyPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse7Ks0f1TyPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse7Ks0f1TyPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1]
type pcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse7Ks0f1TyPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse7Ks0f1TyPcapsResponseFullStatus string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusUnknown           PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "unknown"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusSuccess           PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "success"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusPending           PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "pending"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusRunning           PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "running"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusConversionPending PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusConversionRunning PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusComplete          PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "complete"
	PcapGetResponse7Ks0f1TyPcapsResponseFullStatusFailed            PcapGetResponse7Ks0f1TyPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse7Ks0f1TyPcapsResponseFullSystem string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseFullSystemMagicTransit PcapGetResponse7Ks0f1TyPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse7Ks0f1TyPcapsResponseFullType string

const (
	PcapGetResponse7Ks0f1TyPcapsResponseFullTypeSimple PcapGetResponse7Ks0f1TyPcapsResponseFullType = "simple"
	PcapGetResponse7Ks0f1TyPcapsResponseFullTypeFull   PcapGetResponse7Ks0f1TyPcapsResponseFullType = "full"
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
