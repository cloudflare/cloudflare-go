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

// Union satisfied by [PcapNewResponse2aE72SzcPcapsResponseSimple] or
// [PcapNewResponse2aE72SzcPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponse2aE72SzcPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2aE72SzcPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2aE72SzcPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2aE72SzcPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponse2aE72SzcPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse2aE72SzcPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponse2aE72SzcPcapsResponseSimple]
type pcapNewResponse2aE72SzcPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponse2aE72SzcPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2aE72SzcPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1]
type pcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2aE72SzcPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2aE72SzcPcapsResponseSimpleStatus string

const (
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusUnknown           PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "unknown"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusSuccess           PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "success"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusPending           PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "pending"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusRunning           PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "running"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusConversionPending PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusConversionRunning PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusComplete          PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "complete"
	PcapNewResponse2aE72SzcPcapsResponseSimpleStatusFailed            PcapNewResponse2aE72SzcPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2aE72SzcPcapsResponseSimpleSystem string

const (
	PcapNewResponse2aE72SzcPcapsResponseSimpleSystemMagicTransit PcapNewResponse2aE72SzcPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2aE72SzcPcapsResponseSimpleType string

const (
	PcapNewResponse2aE72SzcPcapsResponseSimpleTypeSimple PcapNewResponse2aE72SzcPcapsResponseSimpleType = "simple"
	PcapNewResponse2aE72SzcPcapsResponseSimpleTypeFull   PcapNewResponse2aE72SzcPcapsResponseSimpleType = "full"
)

type PcapNewResponse2aE72SzcPcapsResponseFull struct {
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
	FilterV1 PcapNewResponse2aE72SzcPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponse2aE72SzcPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponse2aE72SzcPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponse2aE72SzcPcapsResponseFullType `json:"type"`
	JSON pcapNewResponse2aE72SzcPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponse2aE72SzcPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponse2aE72SzcPcapsResponseFull]
type pcapNewResponse2aE72SzcPcapsResponseFullJSON struct {
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

func (r *PcapNewResponse2aE72SzcPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponse2aE72SzcPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponse2aE72SzcPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse2aE72SzcPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse2aE72SzcPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponse2aE72SzcPcapsResponseFullFilterV1]
type pcapNewResponse2aE72SzcPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponse2aE72SzcPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponse2aE72SzcPcapsResponseFullStatus string

const (
	PcapNewResponse2aE72SzcPcapsResponseFullStatusUnknown           PcapNewResponse2aE72SzcPcapsResponseFullStatus = "unknown"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusSuccess           PcapNewResponse2aE72SzcPcapsResponseFullStatus = "success"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusPending           PcapNewResponse2aE72SzcPcapsResponseFullStatus = "pending"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusRunning           PcapNewResponse2aE72SzcPcapsResponseFullStatus = "running"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusConversionPending PcapNewResponse2aE72SzcPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusConversionRunning PcapNewResponse2aE72SzcPcapsResponseFullStatus = "conversion_running"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusComplete          PcapNewResponse2aE72SzcPcapsResponseFullStatus = "complete"
	PcapNewResponse2aE72SzcPcapsResponseFullStatusFailed            PcapNewResponse2aE72SzcPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponse2aE72SzcPcapsResponseFullSystem string

const (
	PcapNewResponse2aE72SzcPcapsResponseFullSystemMagicTransit PcapNewResponse2aE72SzcPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponse2aE72SzcPcapsResponseFullType string

const (
	PcapNewResponse2aE72SzcPcapsResponseFullTypeSimple PcapNewResponse2aE72SzcPcapsResponseFullType = "simple"
	PcapNewResponse2aE72SzcPcapsResponseFullTypeFull   PcapNewResponse2aE72SzcPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponse2aE72SzcPcapsResponseSimple] or
// [PcapListResponse2aE72SzcPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponse2aE72SzcPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponse2aE72SzcPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2aE72SzcPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2aE72SzcPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2aE72SzcPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponse2aE72SzcPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponse2aE72SzcPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponse2aE72SzcPcapsResponseSimple]
type pcapListResponse2aE72SzcPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponse2aE72SzcPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2aE72SzcPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2aE72SzcPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse2aE72SzcPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse2aE72SzcPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponse2aE72SzcPcapsResponseSimpleFilterV1]
type pcapListResponse2aE72SzcPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2aE72SzcPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2aE72SzcPcapsResponseSimpleStatus string

const (
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusUnknown           PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "unknown"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusSuccess           PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "success"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusPending           PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "pending"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusRunning           PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "running"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusConversionPending PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusConversionRunning PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusComplete          PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "complete"
	PcapListResponse2aE72SzcPcapsResponseSimpleStatusFailed            PcapListResponse2aE72SzcPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2aE72SzcPcapsResponseSimpleSystem string

const (
	PcapListResponse2aE72SzcPcapsResponseSimpleSystemMagicTransit PcapListResponse2aE72SzcPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2aE72SzcPcapsResponseSimpleType string

const (
	PcapListResponse2aE72SzcPcapsResponseSimpleTypeSimple PcapListResponse2aE72SzcPcapsResponseSimpleType = "simple"
	PcapListResponse2aE72SzcPcapsResponseSimpleTypeFull   PcapListResponse2aE72SzcPcapsResponseSimpleType = "full"
)

type PcapListResponse2aE72SzcPcapsResponseFull struct {
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
	FilterV1 PcapListResponse2aE72SzcPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponse2aE72SzcPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponse2aE72SzcPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponse2aE72SzcPcapsResponseFullType `json:"type"`
	JSON pcapListResponse2aE72SzcPcapsResponseFullJSON `json:"-"`
}

// pcapListResponse2aE72SzcPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponse2aE72SzcPcapsResponseFull]
type pcapListResponse2aE72SzcPcapsResponseFullJSON struct {
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

func (r *PcapListResponse2aE72SzcPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponse2aE72SzcPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponse2aE72SzcPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse2aE72SzcPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse2aE72SzcPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponse2aE72SzcPcapsResponseFullFilterV1]
type pcapListResponse2aE72SzcPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponse2aE72SzcPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponse2aE72SzcPcapsResponseFullStatus string

const (
	PcapListResponse2aE72SzcPcapsResponseFullStatusUnknown           PcapListResponse2aE72SzcPcapsResponseFullStatus = "unknown"
	PcapListResponse2aE72SzcPcapsResponseFullStatusSuccess           PcapListResponse2aE72SzcPcapsResponseFullStatus = "success"
	PcapListResponse2aE72SzcPcapsResponseFullStatusPending           PcapListResponse2aE72SzcPcapsResponseFullStatus = "pending"
	PcapListResponse2aE72SzcPcapsResponseFullStatusRunning           PcapListResponse2aE72SzcPcapsResponseFullStatus = "running"
	PcapListResponse2aE72SzcPcapsResponseFullStatusConversionPending PcapListResponse2aE72SzcPcapsResponseFullStatus = "conversion_pending"
	PcapListResponse2aE72SzcPcapsResponseFullStatusConversionRunning PcapListResponse2aE72SzcPcapsResponseFullStatus = "conversion_running"
	PcapListResponse2aE72SzcPcapsResponseFullStatusComplete          PcapListResponse2aE72SzcPcapsResponseFullStatus = "complete"
	PcapListResponse2aE72SzcPcapsResponseFullStatusFailed            PcapListResponse2aE72SzcPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponse2aE72SzcPcapsResponseFullSystem string

const (
	PcapListResponse2aE72SzcPcapsResponseFullSystemMagicTransit PcapListResponse2aE72SzcPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponse2aE72SzcPcapsResponseFullType string

const (
	PcapListResponse2aE72SzcPcapsResponseFullTypeSimple PcapListResponse2aE72SzcPcapsResponseFullType = "simple"
	PcapListResponse2aE72SzcPcapsResponseFullTypeFull   PcapListResponse2aE72SzcPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponse2aE72SzcPcapsResponseSimple] or
// [PcapGetResponse2aE72SzcPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponse2aE72SzcPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2aE72SzcPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2aE72SzcPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2aE72SzcPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponse2aE72SzcPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse2aE72SzcPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponse2aE72SzcPcapsResponseSimple]
type pcapGetResponse2aE72SzcPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponse2aE72SzcPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2aE72SzcPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1]
type pcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2aE72SzcPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2aE72SzcPcapsResponseSimpleStatus string

const (
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusUnknown           PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "unknown"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusSuccess           PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "success"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusPending           PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "pending"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusRunning           PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "running"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusConversionPending PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusConversionRunning PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusComplete          PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "complete"
	PcapGetResponse2aE72SzcPcapsResponseSimpleStatusFailed            PcapGetResponse2aE72SzcPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2aE72SzcPcapsResponseSimpleSystem string

const (
	PcapGetResponse2aE72SzcPcapsResponseSimpleSystemMagicTransit PcapGetResponse2aE72SzcPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2aE72SzcPcapsResponseSimpleType string

const (
	PcapGetResponse2aE72SzcPcapsResponseSimpleTypeSimple PcapGetResponse2aE72SzcPcapsResponseSimpleType = "simple"
	PcapGetResponse2aE72SzcPcapsResponseSimpleTypeFull   PcapGetResponse2aE72SzcPcapsResponseSimpleType = "full"
)

type PcapGetResponse2aE72SzcPcapsResponseFull struct {
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
	FilterV1 PcapGetResponse2aE72SzcPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponse2aE72SzcPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponse2aE72SzcPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponse2aE72SzcPcapsResponseFullType `json:"type"`
	JSON pcapGetResponse2aE72SzcPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponse2aE72SzcPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponse2aE72SzcPcapsResponseFull]
type pcapGetResponse2aE72SzcPcapsResponseFullJSON struct {
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

func (r *PcapGetResponse2aE72SzcPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponse2aE72SzcPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponse2aE72SzcPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse2aE72SzcPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse2aE72SzcPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponse2aE72SzcPcapsResponseFullFilterV1]
type pcapGetResponse2aE72SzcPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponse2aE72SzcPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponse2aE72SzcPcapsResponseFullStatus string

const (
	PcapGetResponse2aE72SzcPcapsResponseFullStatusUnknown           PcapGetResponse2aE72SzcPcapsResponseFullStatus = "unknown"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusSuccess           PcapGetResponse2aE72SzcPcapsResponseFullStatus = "success"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusPending           PcapGetResponse2aE72SzcPcapsResponseFullStatus = "pending"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusRunning           PcapGetResponse2aE72SzcPcapsResponseFullStatus = "running"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusConversionPending PcapGetResponse2aE72SzcPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusConversionRunning PcapGetResponse2aE72SzcPcapsResponseFullStatus = "conversion_running"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusComplete          PcapGetResponse2aE72SzcPcapsResponseFullStatus = "complete"
	PcapGetResponse2aE72SzcPcapsResponseFullStatusFailed            PcapGetResponse2aE72SzcPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponse2aE72SzcPcapsResponseFullSystem string

const (
	PcapGetResponse2aE72SzcPcapsResponseFullSystemMagicTransit PcapGetResponse2aE72SzcPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponse2aE72SzcPcapsResponseFullType string

const (
	PcapGetResponse2aE72SzcPcapsResponseFullTypeSimple PcapGetResponse2aE72SzcPcapsResponseFullType = "simple"
	PcapGetResponse2aE72SzcPcapsResponseFullTypeFull   PcapGetResponse2aE72SzcPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParams2aE72SzcPcapsRequestSimple],
// [PcapNewParams2aE72SzcPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParams2aE72SzcPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams2aE72SzcPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams2aE72SzcPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams2aE72SzcPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParams2aE72SzcPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams2aE72SzcPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams2aE72SzcPcapsRequestSimpleSystem string

const (
	PcapNewParams2aE72SzcPcapsRequestSimpleSystemMagicTransit PcapNewParams2aE72SzcPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams2aE72SzcPcapsRequestSimpleType string

const (
	PcapNewParams2aE72SzcPcapsRequestSimpleTypeSimple PcapNewParams2aE72SzcPcapsRequestSimpleType = "simple"
	PcapNewParams2aE72SzcPcapsRequestSimpleTypeFull   PcapNewParams2aE72SzcPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams2aE72SzcPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParams2aE72SzcPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParams2aE72SzcPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParams2aE72SzcPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParams2aE72SzcPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParams2aE72SzcPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParams2aE72SzcPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParams2aE72SzcPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParams2aE72SzcPcapsRequestFullSystem string

const (
	PcapNewParams2aE72SzcPcapsRequestFullSystemMagicTransit PcapNewParams2aE72SzcPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParams2aE72SzcPcapsRequestFullType string

const (
	PcapNewParams2aE72SzcPcapsRequestFullTypeSimple PcapNewParams2aE72SzcPcapsRequestFullType = "simple"
	PcapNewParams2aE72SzcPcapsRequestFullTypeFull   PcapNewParams2aE72SzcPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParams2aE72SzcPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParams2aE72SzcPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
