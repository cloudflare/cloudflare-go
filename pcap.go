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

// Union satisfied by [PcapNewResponseOSw5wAfZPcapsResponseSimple] or
// [PcapNewResponseOSw5wAfZPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseOSw5wAfZPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseOSw5wAfZPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseOSw5wAfZPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseOSw5wAfZPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseOSw5wAfZPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseOSw5wAfZPcapsResponseSimple]
type pcapNewResponseOSw5wAfZPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseOSw5wAfZPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseOSw5wAfZPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1]
type pcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseOSw5wAfZPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus string

const (
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusUnknown           PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusSuccess           PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "success"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusPending           PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "pending"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusRunning           PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "running"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusConversionPending PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusConversionRunning PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusComplete          PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "complete"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleStatusFailed            PcapNewResponseOSw5wAfZPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseOSw5wAfZPcapsResponseSimpleSystem string

const (
	PcapNewResponseOSw5wAfZPcapsResponseSimpleSystemMagicTransit PcapNewResponseOSw5wAfZPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseOSw5wAfZPcapsResponseSimpleType string

const (
	PcapNewResponseOSw5wAfZPcapsResponseSimpleTypeSimple PcapNewResponseOSw5wAfZPcapsResponseSimpleType = "simple"
	PcapNewResponseOSw5wAfZPcapsResponseSimpleTypeFull   PcapNewResponseOSw5wAfZPcapsResponseSimpleType = "full"
)

type PcapNewResponseOSw5wAfZPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseOSw5wAfZPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseOSw5wAfZPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseOSw5wAfZPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseOSw5wAfZPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseOSw5wAfZPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseOSw5wAfZPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseOSw5wAfZPcapsResponseFull]
type pcapNewResponseOSw5wAfZPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseOSw5wAfZPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseOSw5wAfZPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseOSw5wAfZPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseOSw5wAfZPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseOSw5wAfZPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseOSw5wAfZPcapsResponseFullFilterV1]
type pcapNewResponseOSw5wAfZPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseOSw5wAfZPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseOSw5wAfZPcapsResponseFullStatus string

const (
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusUnknown           PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "unknown"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusSuccess           PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "success"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusPending           PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "pending"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusRunning           PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "running"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusConversionPending PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusConversionRunning PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusComplete          PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "complete"
	PcapNewResponseOSw5wAfZPcapsResponseFullStatusFailed            PcapNewResponseOSw5wAfZPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseOSw5wAfZPcapsResponseFullSystem string

const (
	PcapNewResponseOSw5wAfZPcapsResponseFullSystemMagicTransit PcapNewResponseOSw5wAfZPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseOSw5wAfZPcapsResponseFullType string

const (
	PcapNewResponseOSw5wAfZPcapsResponseFullTypeSimple PcapNewResponseOSw5wAfZPcapsResponseFullType = "simple"
	PcapNewResponseOSw5wAfZPcapsResponseFullTypeFull   PcapNewResponseOSw5wAfZPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseOSw5wAfZPcapsResponseSimple] or
// [PcapListResponseOSw5wAfZPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseOSw5wAfZPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseOSw5wAfZPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseOSw5wAfZPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseOSw5wAfZPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseOSw5wAfZPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseOSw5wAfZPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseOSw5wAfZPcapsResponseSimple]
type pcapListResponseOSw5wAfZPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseOSw5wAfZPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseOSw5wAfZPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1]
type pcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseOSw5wAfZPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseOSw5wAfZPcapsResponseSimpleStatus string

const (
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusUnknown           PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "unknown"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusSuccess           PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "success"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusPending           PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "pending"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusRunning           PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "running"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusConversionPending PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusConversionRunning PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusComplete          PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "complete"
	PcapListResponseOSw5wAfZPcapsResponseSimpleStatusFailed            PcapListResponseOSw5wAfZPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseOSw5wAfZPcapsResponseSimpleSystem string

const (
	PcapListResponseOSw5wAfZPcapsResponseSimpleSystemMagicTransit PcapListResponseOSw5wAfZPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseOSw5wAfZPcapsResponseSimpleType string

const (
	PcapListResponseOSw5wAfZPcapsResponseSimpleTypeSimple PcapListResponseOSw5wAfZPcapsResponseSimpleType = "simple"
	PcapListResponseOSw5wAfZPcapsResponseSimpleTypeFull   PcapListResponseOSw5wAfZPcapsResponseSimpleType = "full"
)

type PcapListResponseOSw5wAfZPcapsResponseFull struct {
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
	FilterV1 PcapListResponseOSw5wAfZPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseOSw5wAfZPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseOSw5wAfZPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseOSw5wAfZPcapsResponseFullType `json:"type"`
	JSON pcapListResponseOSw5wAfZPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseOSw5wAfZPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseOSw5wAfZPcapsResponseFull]
type pcapListResponseOSw5wAfZPcapsResponseFullJSON struct {
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

func (r *PcapListResponseOSw5wAfZPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseOSw5wAfZPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseOSw5wAfZPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseOSw5wAfZPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseOSw5wAfZPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseOSw5wAfZPcapsResponseFullFilterV1]
type pcapListResponseOSw5wAfZPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseOSw5wAfZPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseOSw5wAfZPcapsResponseFullStatus string

const (
	PcapListResponseOSw5wAfZPcapsResponseFullStatusUnknown           PcapListResponseOSw5wAfZPcapsResponseFullStatus = "unknown"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusSuccess           PcapListResponseOSw5wAfZPcapsResponseFullStatus = "success"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusPending           PcapListResponseOSw5wAfZPcapsResponseFullStatus = "pending"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusRunning           PcapListResponseOSw5wAfZPcapsResponseFullStatus = "running"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusConversionPending PcapListResponseOSw5wAfZPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusConversionRunning PcapListResponseOSw5wAfZPcapsResponseFullStatus = "conversion_running"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusComplete          PcapListResponseOSw5wAfZPcapsResponseFullStatus = "complete"
	PcapListResponseOSw5wAfZPcapsResponseFullStatusFailed            PcapListResponseOSw5wAfZPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseOSw5wAfZPcapsResponseFullSystem string

const (
	PcapListResponseOSw5wAfZPcapsResponseFullSystemMagicTransit PcapListResponseOSw5wAfZPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseOSw5wAfZPcapsResponseFullType string

const (
	PcapListResponseOSw5wAfZPcapsResponseFullTypeSimple PcapListResponseOSw5wAfZPcapsResponseFullType = "simple"
	PcapListResponseOSw5wAfZPcapsResponseFullTypeFull   PcapListResponseOSw5wAfZPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseOSw5wAfZPcapsResponseSimple] or
// [PcapGetResponseOSw5wAfZPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseOSw5wAfZPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseOSw5wAfZPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseOSw5wAfZPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseOSw5wAfZPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseOSw5wAfZPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseOSw5wAfZPcapsResponseSimple]
type pcapGetResponseOSw5wAfZPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseOSw5wAfZPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseOSw5wAfZPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1]
type pcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseOSw5wAfZPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus string

const (
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusUnknown           PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusSuccess           PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "success"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusPending           PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "pending"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusRunning           PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "running"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusConversionPending PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusConversionRunning PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusComplete          PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "complete"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleStatusFailed            PcapGetResponseOSw5wAfZPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseOSw5wAfZPcapsResponseSimpleSystem string

const (
	PcapGetResponseOSw5wAfZPcapsResponseSimpleSystemMagicTransit PcapGetResponseOSw5wAfZPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseOSw5wAfZPcapsResponseSimpleType string

const (
	PcapGetResponseOSw5wAfZPcapsResponseSimpleTypeSimple PcapGetResponseOSw5wAfZPcapsResponseSimpleType = "simple"
	PcapGetResponseOSw5wAfZPcapsResponseSimpleTypeFull   PcapGetResponseOSw5wAfZPcapsResponseSimpleType = "full"
)

type PcapGetResponseOSw5wAfZPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseOSw5wAfZPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseOSw5wAfZPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseOSw5wAfZPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseOSw5wAfZPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseOSw5wAfZPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseOSw5wAfZPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseOSw5wAfZPcapsResponseFull]
type pcapGetResponseOSw5wAfZPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseOSw5wAfZPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseOSw5wAfZPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseOSw5wAfZPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseOSw5wAfZPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseOSw5wAfZPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseOSw5wAfZPcapsResponseFullFilterV1]
type pcapGetResponseOSw5wAfZPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseOSw5wAfZPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseOSw5wAfZPcapsResponseFullStatus string

const (
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusUnknown           PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "unknown"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusSuccess           PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "success"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusPending           PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "pending"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusRunning           PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "running"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusConversionPending PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusConversionRunning PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusComplete          PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "complete"
	PcapGetResponseOSw5wAfZPcapsResponseFullStatusFailed            PcapGetResponseOSw5wAfZPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseOSw5wAfZPcapsResponseFullSystem string

const (
	PcapGetResponseOSw5wAfZPcapsResponseFullSystemMagicTransit PcapGetResponseOSw5wAfZPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseOSw5wAfZPcapsResponseFullType string

const (
	PcapGetResponseOSw5wAfZPcapsResponseFullTypeSimple PcapGetResponseOSw5wAfZPcapsResponseFullType = "simple"
	PcapGetResponseOSw5wAfZPcapsResponseFullTypeFull   PcapGetResponseOSw5wAfZPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsOSw5wAfZPcapsRequestSimple],
// [PcapNewParamsOSw5wAfZPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsOSw5wAfZPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsOSw5wAfZPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsOSw5wAfZPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsOSw5wAfZPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsOSw5wAfZPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsOSw5wAfZPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsOSw5wAfZPcapsRequestSimpleSystem string

const (
	PcapNewParamsOSw5wAfZPcapsRequestSimpleSystemMagicTransit PcapNewParamsOSw5wAfZPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsOSw5wAfZPcapsRequestSimpleType string

const (
	PcapNewParamsOSw5wAfZPcapsRequestSimpleTypeSimple PcapNewParamsOSw5wAfZPcapsRequestSimpleType = "simple"
	PcapNewParamsOSw5wAfZPcapsRequestSimpleTypeFull   PcapNewParamsOSw5wAfZPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsOSw5wAfZPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsOSw5wAfZPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsOSw5wAfZPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsOSw5wAfZPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsOSw5wAfZPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsOSw5wAfZPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsOSw5wAfZPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsOSw5wAfZPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsOSw5wAfZPcapsRequestFullSystem string

const (
	PcapNewParamsOSw5wAfZPcapsRequestFullSystemMagicTransit PcapNewParamsOSw5wAfZPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsOSw5wAfZPcapsRequestFullType string

const (
	PcapNewParamsOSw5wAfZPcapsRequestFullTypeSimple PcapNewParamsOSw5wAfZPcapsRequestFullType = "simple"
	PcapNewParamsOSw5wAfZPcapsRequestFullTypeFull   PcapNewParamsOSw5wAfZPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsOSw5wAfZPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsOSw5wAfZPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
