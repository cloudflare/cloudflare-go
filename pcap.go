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

// Union satisfied by [PcapNewResponsePgpenC68PcapsResponseSimple] or
// [PcapNewResponsePgpenC68PcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponsePgpenC68PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponsePgpenC68PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponsePgpenC68PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponsePgpenC68PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponsePgpenC68PcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponsePgpenC68PcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponsePgpenC68PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponsePgpenC68PcapsResponseSimple]
type pcapNewResponsePgpenC68PcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponsePgpenC68PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponsePgpenC68PcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponsePgpenC68PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponsePgpenC68PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponsePgpenC68PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponsePgpenC68PcapsResponseSimpleFilterV1]
type pcapNewResponsePgpenC68PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponsePgpenC68PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponsePgpenC68PcapsResponseSimpleStatus string

const (
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusUnknown           PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "unknown"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusSuccess           PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "success"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusPending           PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "pending"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusRunning           PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "running"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusConversionPending PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusConversionRunning PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusComplete          PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "complete"
	PcapNewResponsePgpenC68PcapsResponseSimpleStatusFailed            PcapNewResponsePgpenC68PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponsePgpenC68PcapsResponseSimpleSystem string

const (
	PcapNewResponsePgpenC68PcapsResponseSimpleSystemMagicTransit PcapNewResponsePgpenC68PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponsePgpenC68PcapsResponseSimpleType string

const (
	PcapNewResponsePgpenC68PcapsResponseSimpleTypeSimple PcapNewResponsePgpenC68PcapsResponseSimpleType = "simple"
	PcapNewResponsePgpenC68PcapsResponseSimpleTypeFull   PcapNewResponsePgpenC68PcapsResponseSimpleType = "full"
)

type PcapNewResponsePgpenC68PcapsResponseFull struct {
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
	FilterV1 PcapNewResponsePgpenC68PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponsePgpenC68PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponsePgpenC68PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponsePgpenC68PcapsResponseFullType `json:"type"`
	JSON pcapNewResponsePgpenC68PcapsResponseFullJSON `json:"-"`
}

// pcapNewResponsePgpenC68PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponsePgpenC68PcapsResponseFull]
type pcapNewResponsePgpenC68PcapsResponseFullJSON struct {
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

func (r *PcapNewResponsePgpenC68PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponsePgpenC68PcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponsePgpenC68PcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponsePgpenC68PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponsePgpenC68PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponsePgpenC68PcapsResponseFullFilterV1]
type pcapNewResponsePgpenC68PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponsePgpenC68PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponsePgpenC68PcapsResponseFullStatus string

const (
	PcapNewResponsePgpenC68PcapsResponseFullStatusUnknown           PcapNewResponsePgpenC68PcapsResponseFullStatus = "unknown"
	PcapNewResponsePgpenC68PcapsResponseFullStatusSuccess           PcapNewResponsePgpenC68PcapsResponseFullStatus = "success"
	PcapNewResponsePgpenC68PcapsResponseFullStatusPending           PcapNewResponsePgpenC68PcapsResponseFullStatus = "pending"
	PcapNewResponsePgpenC68PcapsResponseFullStatusRunning           PcapNewResponsePgpenC68PcapsResponseFullStatus = "running"
	PcapNewResponsePgpenC68PcapsResponseFullStatusConversionPending PcapNewResponsePgpenC68PcapsResponseFullStatus = "conversion_pending"
	PcapNewResponsePgpenC68PcapsResponseFullStatusConversionRunning PcapNewResponsePgpenC68PcapsResponseFullStatus = "conversion_running"
	PcapNewResponsePgpenC68PcapsResponseFullStatusComplete          PcapNewResponsePgpenC68PcapsResponseFullStatus = "complete"
	PcapNewResponsePgpenC68PcapsResponseFullStatusFailed            PcapNewResponsePgpenC68PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponsePgpenC68PcapsResponseFullSystem string

const (
	PcapNewResponsePgpenC68PcapsResponseFullSystemMagicTransit PcapNewResponsePgpenC68PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponsePgpenC68PcapsResponseFullType string

const (
	PcapNewResponsePgpenC68PcapsResponseFullTypeSimple PcapNewResponsePgpenC68PcapsResponseFullType = "simple"
	PcapNewResponsePgpenC68PcapsResponseFullTypeFull   PcapNewResponsePgpenC68PcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponsePgpenC68PcapsResponseSimple] or
// [PcapListResponsePgpenC68PcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponsePgpenC68PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponsePgpenC68PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponsePgpenC68PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponsePgpenC68PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponsePgpenC68PcapsResponseSimpleType `json:"type"`
	JSON pcapListResponsePgpenC68PcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponsePgpenC68PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponsePgpenC68PcapsResponseSimple]
type pcapListResponsePgpenC68PcapsResponseSimpleJSON struct {
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

func (r *PcapListResponsePgpenC68PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponsePgpenC68PcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponsePgpenC68PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponsePgpenC68PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponsePgpenC68PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponsePgpenC68PcapsResponseSimpleFilterV1]
type pcapListResponsePgpenC68PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponsePgpenC68PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponsePgpenC68PcapsResponseSimpleStatus string

const (
	PcapListResponsePgpenC68PcapsResponseSimpleStatusUnknown           PcapListResponsePgpenC68PcapsResponseSimpleStatus = "unknown"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusSuccess           PcapListResponsePgpenC68PcapsResponseSimpleStatus = "success"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusPending           PcapListResponsePgpenC68PcapsResponseSimpleStatus = "pending"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusRunning           PcapListResponsePgpenC68PcapsResponseSimpleStatus = "running"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusConversionPending PcapListResponsePgpenC68PcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusConversionRunning PcapListResponsePgpenC68PcapsResponseSimpleStatus = "conversion_running"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusComplete          PcapListResponsePgpenC68PcapsResponseSimpleStatus = "complete"
	PcapListResponsePgpenC68PcapsResponseSimpleStatusFailed            PcapListResponsePgpenC68PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponsePgpenC68PcapsResponseSimpleSystem string

const (
	PcapListResponsePgpenC68PcapsResponseSimpleSystemMagicTransit PcapListResponsePgpenC68PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponsePgpenC68PcapsResponseSimpleType string

const (
	PcapListResponsePgpenC68PcapsResponseSimpleTypeSimple PcapListResponsePgpenC68PcapsResponseSimpleType = "simple"
	PcapListResponsePgpenC68PcapsResponseSimpleTypeFull   PcapListResponsePgpenC68PcapsResponseSimpleType = "full"
)

type PcapListResponsePgpenC68PcapsResponseFull struct {
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
	FilterV1 PcapListResponsePgpenC68PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponsePgpenC68PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponsePgpenC68PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponsePgpenC68PcapsResponseFullType `json:"type"`
	JSON pcapListResponsePgpenC68PcapsResponseFullJSON `json:"-"`
}

// pcapListResponsePgpenC68PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponsePgpenC68PcapsResponseFull]
type pcapListResponsePgpenC68PcapsResponseFullJSON struct {
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

func (r *PcapListResponsePgpenC68PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponsePgpenC68PcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponsePgpenC68PcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponsePgpenC68PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponsePgpenC68PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponsePgpenC68PcapsResponseFullFilterV1]
type pcapListResponsePgpenC68PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponsePgpenC68PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponsePgpenC68PcapsResponseFullStatus string

const (
	PcapListResponsePgpenC68PcapsResponseFullStatusUnknown           PcapListResponsePgpenC68PcapsResponseFullStatus = "unknown"
	PcapListResponsePgpenC68PcapsResponseFullStatusSuccess           PcapListResponsePgpenC68PcapsResponseFullStatus = "success"
	PcapListResponsePgpenC68PcapsResponseFullStatusPending           PcapListResponsePgpenC68PcapsResponseFullStatus = "pending"
	PcapListResponsePgpenC68PcapsResponseFullStatusRunning           PcapListResponsePgpenC68PcapsResponseFullStatus = "running"
	PcapListResponsePgpenC68PcapsResponseFullStatusConversionPending PcapListResponsePgpenC68PcapsResponseFullStatus = "conversion_pending"
	PcapListResponsePgpenC68PcapsResponseFullStatusConversionRunning PcapListResponsePgpenC68PcapsResponseFullStatus = "conversion_running"
	PcapListResponsePgpenC68PcapsResponseFullStatusComplete          PcapListResponsePgpenC68PcapsResponseFullStatus = "complete"
	PcapListResponsePgpenC68PcapsResponseFullStatusFailed            PcapListResponsePgpenC68PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponsePgpenC68PcapsResponseFullSystem string

const (
	PcapListResponsePgpenC68PcapsResponseFullSystemMagicTransit PcapListResponsePgpenC68PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponsePgpenC68PcapsResponseFullType string

const (
	PcapListResponsePgpenC68PcapsResponseFullTypeSimple PcapListResponsePgpenC68PcapsResponseFullType = "simple"
	PcapListResponsePgpenC68PcapsResponseFullTypeFull   PcapListResponsePgpenC68PcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponsePgpenC68PcapsResponseSimple] or
// [PcapGetResponsePgpenC68PcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponsePgpenC68PcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponsePgpenC68PcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponsePgpenC68PcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponsePgpenC68PcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponsePgpenC68PcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponsePgpenC68PcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponsePgpenC68PcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponsePgpenC68PcapsResponseSimple]
type pcapGetResponsePgpenC68PcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponsePgpenC68PcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponsePgpenC68PcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponsePgpenC68PcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponsePgpenC68PcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponsePgpenC68PcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponsePgpenC68PcapsResponseSimpleFilterV1]
type pcapGetResponsePgpenC68PcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponsePgpenC68PcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponsePgpenC68PcapsResponseSimpleStatus string

const (
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusUnknown           PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "unknown"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusSuccess           PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "success"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusPending           PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "pending"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusRunning           PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "running"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusConversionPending PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusConversionRunning PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusComplete          PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "complete"
	PcapGetResponsePgpenC68PcapsResponseSimpleStatusFailed            PcapGetResponsePgpenC68PcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponsePgpenC68PcapsResponseSimpleSystem string

const (
	PcapGetResponsePgpenC68PcapsResponseSimpleSystemMagicTransit PcapGetResponsePgpenC68PcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponsePgpenC68PcapsResponseSimpleType string

const (
	PcapGetResponsePgpenC68PcapsResponseSimpleTypeSimple PcapGetResponsePgpenC68PcapsResponseSimpleType = "simple"
	PcapGetResponsePgpenC68PcapsResponseSimpleTypeFull   PcapGetResponsePgpenC68PcapsResponseSimpleType = "full"
)

type PcapGetResponsePgpenC68PcapsResponseFull struct {
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
	FilterV1 PcapGetResponsePgpenC68PcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponsePgpenC68PcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponsePgpenC68PcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponsePgpenC68PcapsResponseFullType `json:"type"`
	JSON pcapGetResponsePgpenC68PcapsResponseFullJSON `json:"-"`
}

// pcapGetResponsePgpenC68PcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponsePgpenC68PcapsResponseFull]
type pcapGetResponsePgpenC68PcapsResponseFullJSON struct {
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

func (r *PcapGetResponsePgpenC68PcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponsePgpenC68PcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponsePgpenC68PcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponsePgpenC68PcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponsePgpenC68PcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponsePgpenC68PcapsResponseFullFilterV1]
type pcapGetResponsePgpenC68PcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponsePgpenC68PcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponsePgpenC68PcapsResponseFullStatus string

const (
	PcapGetResponsePgpenC68PcapsResponseFullStatusUnknown           PcapGetResponsePgpenC68PcapsResponseFullStatus = "unknown"
	PcapGetResponsePgpenC68PcapsResponseFullStatusSuccess           PcapGetResponsePgpenC68PcapsResponseFullStatus = "success"
	PcapGetResponsePgpenC68PcapsResponseFullStatusPending           PcapGetResponsePgpenC68PcapsResponseFullStatus = "pending"
	PcapGetResponsePgpenC68PcapsResponseFullStatusRunning           PcapGetResponsePgpenC68PcapsResponseFullStatus = "running"
	PcapGetResponsePgpenC68PcapsResponseFullStatusConversionPending PcapGetResponsePgpenC68PcapsResponseFullStatus = "conversion_pending"
	PcapGetResponsePgpenC68PcapsResponseFullStatusConversionRunning PcapGetResponsePgpenC68PcapsResponseFullStatus = "conversion_running"
	PcapGetResponsePgpenC68PcapsResponseFullStatusComplete          PcapGetResponsePgpenC68PcapsResponseFullStatus = "complete"
	PcapGetResponsePgpenC68PcapsResponseFullStatusFailed            PcapGetResponsePgpenC68PcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponsePgpenC68PcapsResponseFullSystem string

const (
	PcapGetResponsePgpenC68PcapsResponseFullSystemMagicTransit PcapGetResponsePgpenC68PcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponsePgpenC68PcapsResponseFullType string

const (
	PcapGetResponsePgpenC68PcapsResponseFullTypeSimple PcapGetResponsePgpenC68PcapsResponseFullType = "simple"
	PcapGetResponsePgpenC68PcapsResponseFullTypeFull   PcapGetResponsePgpenC68PcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsPgpenC68PcapsRequestSimple],
// [PcapNewParamsPgpenC68PcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsPgpenC68PcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsPgpenC68PcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsPgpenC68PcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsPgpenC68PcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsPgpenC68PcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsPgpenC68PcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsPgpenC68PcapsRequestSimpleSystem string

const (
	PcapNewParamsPgpenC68PcapsRequestSimpleSystemMagicTransit PcapNewParamsPgpenC68PcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsPgpenC68PcapsRequestSimpleType string

const (
	PcapNewParamsPgpenC68PcapsRequestSimpleTypeSimple PcapNewParamsPgpenC68PcapsRequestSimpleType = "simple"
	PcapNewParamsPgpenC68PcapsRequestSimpleTypeFull   PcapNewParamsPgpenC68PcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsPgpenC68PcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsPgpenC68PcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsPgpenC68PcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsPgpenC68PcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsPgpenC68PcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsPgpenC68PcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsPgpenC68PcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsPgpenC68PcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsPgpenC68PcapsRequestFullSystem string

const (
	PcapNewParamsPgpenC68PcapsRequestFullSystemMagicTransit PcapNewParamsPgpenC68PcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsPgpenC68PcapsRequestFullType string

const (
	PcapNewParamsPgpenC68PcapsRequestFullTypeSimple PcapNewParamsPgpenC68PcapsRequestFullType = "simple"
	PcapNewParamsPgpenC68PcapsRequestFullTypeFull   PcapNewParamsPgpenC68PcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsPgpenC68PcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsPgpenC68PcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
