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

// Union satisfied by [PcapNewResponseOvZ3JqIqPcapsResponseSimple] or
// [PcapNewResponseOvZ3JqIqPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseOvZ3JqIqPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseOvZ3JqIqPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseOvZ3JqIqPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseOvZ3JqIqPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseOvZ3JqIqPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseOvZ3JqIqPcapsResponseSimple]
type pcapNewResponseOvZ3JqIqPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseOvZ3JqIqPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseOvZ3JqIqPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1]
type pcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseOvZ3JqIqPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusUnknown           PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusSuccess           PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "success"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusPending           PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "pending"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusRunning           PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "running"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusConversionPending PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusConversionRunning PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusComplete          PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "complete"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatusFailed            PcapNewResponseOvZ3JqIqPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseOvZ3JqIqPcapsResponseSimpleSystem string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleSystemMagicTransit PcapNewResponseOvZ3JqIqPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseOvZ3JqIqPcapsResponseSimpleType string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleTypeSimple PcapNewResponseOvZ3JqIqPcapsResponseSimpleType = "simple"
	PcapNewResponseOvZ3JqIqPcapsResponseSimpleTypeFull   PcapNewResponseOvZ3JqIqPcapsResponseSimpleType = "full"
)

type PcapNewResponseOvZ3JqIqPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseOvZ3JqIqPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseOvZ3JqIqPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseOvZ3JqIqPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseOvZ3JqIqPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseOvZ3JqIqPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseOvZ3JqIqPcapsResponseFull]
type pcapNewResponseOvZ3JqIqPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseOvZ3JqIqPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseOvZ3JqIqPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1]
type pcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseOvZ3JqIqPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseOvZ3JqIqPcapsResponseFullStatus string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusUnknown           PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "unknown"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusSuccess           PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "success"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusPending           PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "pending"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusRunning           PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "running"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusConversionPending PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusConversionRunning PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusComplete          PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "complete"
	PcapNewResponseOvZ3JqIqPcapsResponseFullStatusFailed            PcapNewResponseOvZ3JqIqPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseOvZ3JqIqPcapsResponseFullSystem string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseFullSystemMagicTransit PcapNewResponseOvZ3JqIqPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseOvZ3JqIqPcapsResponseFullType string

const (
	PcapNewResponseOvZ3JqIqPcapsResponseFullTypeSimple PcapNewResponseOvZ3JqIqPcapsResponseFullType = "simple"
	PcapNewResponseOvZ3JqIqPcapsResponseFullTypeFull   PcapNewResponseOvZ3JqIqPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseOvZ3JqIqPcapsResponseSimple] or
// [PcapListResponseOvZ3JqIqPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseOvZ3JqIqPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseOvZ3JqIqPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseOvZ3JqIqPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseOvZ3JqIqPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseOvZ3JqIqPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseOvZ3JqIqPcapsResponseSimple]
type pcapListResponseOvZ3JqIqPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseOvZ3JqIqPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseOvZ3JqIqPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1]
type pcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseOvZ3JqIqPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus string

const (
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusUnknown           PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "unknown"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusSuccess           PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "success"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusPending           PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "pending"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusRunning           PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "running"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusConversionPending PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusConversionRunning PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusComplete          PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "complete"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleStatusFailed            PcapListResponseOvZ3JqIqPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseOvZ3JqIqPcapsResponseSimpleSystem string

const (
	PcapListResponseOvZ3JqIqPcapsResponseSimpleSystemMagicTransit PcapListResponseOvZ3JqIqPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseOvZ3JqIqPcapsResponseSimpleType string

const (
	PcapListResponseOvZ3JqIqPcapsResponseSimpleTypeSimple PcapListResponseOvZ3JqIqPcapsResponseSimpleType = "simple"
	PcapListResponseOvZ3JqIqPcapsResponseSimpleTypeFull   PcapListResponseOvZ3JqIqPcapsResponseSimpleType = "full"
)

type PcapListResponseOvZ3JqIqPcapsResponseFull struct {
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
	FilterV1 PcapListResponseOvZ3JqIqPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseOvZ3JqIqPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseOvZ3JqIqPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseOvZ3JqIqPcapsResponseFullType `json:"type"`
	JSON pcapListResponseOvZ3JqIqPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseOvZ3JqIqPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseOvZ3JqIqPcapsResponseFull]
type pcapListResponseOvZ3JqIqPcapsResponseFullJSON struct {
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

func (r *PcapListResponseOvZ3JqIqPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseOvZ3JqIqPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseOvZ3JqIqPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseOvZ3JqIqPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseOvZ3JqIqPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseOvZ3JqIqPcapsResponseFullFilterV1]
type pcapListResponseOvZ3JqIqPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseOvZ3JqIqPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseOvZ3JqIqPcapsResponseFullStatus string

const (
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusUnknown           PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "unknown"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusSuccess           PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "success"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusPending           PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "pending"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusRunning           PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "running"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusConversionPending PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusConversionRunning PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_running"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusComplete          PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "complete"
	PcapListResponseOvZ3JqIqPcapsResponseFullStatusFailed            PcapListResponseOvZ3JqIqPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseOvZ3JqIqPcapsResponseFullSystem string

const (
	PcapListResponseOvZ3JqIqPcapsResponseFullSystemMagicTransit PcapListResponseOvZ3JqIqPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseOvZ3JqIqPcapsResponseFullType string

const (
	PcapListResponseOvZ3JqIqPcapsResponseFullTypeSimple PcapListResponseOvZ3JqIqPcapsResponseFullType = "simple"
	PcapListResponseOvZ3JqIqPcapsResponseFullTypeFull   PcapListResponseOvZ3JqIqPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseOvZ3JqIqPcapsResponseSimple] or
// [PcapGetResponseOvZ3JqIqPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseOvZ3JqIqPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseOvZ3JqIqPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseOvZ3JqIqPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseOvZ3JqIqPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseOvZ3JqIqPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseOvZ3JqIqPcapsResponseSimple]
type pcapGetResponseOvZ3JqIqPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseOvZ3JqIqPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseOvZ3JqIqPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1]
type pcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseOvZ3JqIqPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusUnknown           PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusSuccess           PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "success"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusPending           PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "pending"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusRunning           PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "running"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusConversionPending PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusConversionRunning PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusComplete          PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "complete"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatusFailed            PcapGetResponseOvZ3JqIqPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseOvZ3JqIqPcapsResponseSimpleSystem string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleSystemMagicTransit PcapGetResponseOvZ3JqIqPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseOvZ3JqIqPcapsResponseSimpleType string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleTypeSimple PcapGetResponseOvZ3JqIqPcapsResponseSimpleType = "simple"
	PcapGetResponseOvZ3JqIqPcapsResponseSimpleTypeFull   PcapGetResponseOvZ3JqIqPcapsResponseSimpleType = "full"
)

type PcapGetResponseOvZ3JqIqPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseOvZ3JqIqPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseOvZ3JqIqPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseOvZ3JqIqPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseOvZ3JqIqPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseOvZ3JqIqPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseOvZ3JqIqPcapsResponseFull]
type pcapGetResponseOvZ3JqIqPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseOvZ3JqIqPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseOvZ3JqIqPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1]
type pcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseOvZ3JqIqPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseOvZ3JqIqPcapsResponseFullStatus string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusUnknown           PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "unknown"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusSuccess           PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "success"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusPending           PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "pending"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusRunning           PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "running"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusConversionPending PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusConversionRunning PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusComplete          PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "complete"
	PcapGetResponseOvZ3JqIqPcapsResponseFullStatusFailed            PcapGetResponseOvZ3JqIqPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseOvZ3JqIqPcapsResponseFullSystem string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseFullSystemMagicTransit PcapGetResponseOvZ3JqIqPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseOvZ3JqIqPcapsResponseFullType string

const (
	PcapGetResponseOvZ3JqIqPcapsResponseFullTypeSimple PcapGetResponseOvZ3JqIqPcapsResponseFullType = "simple"
	PcapGetResponseOvZ3JqIqPcapsResponseFullTypeFull   PcapGetResponseOvZ3JqIqPcapsResponseFullType = "full"
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
