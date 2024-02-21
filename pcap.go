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

// Union satisfied by [PcapNewResponseFfMrAnLxPcapsResponseSimple] or
// [PcapNewResponseFfMrAnLxPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseFfMrAnLxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseFfMrAnLxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseFfMrAnLxPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseFfMrAnLxPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseFfMrAnLxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseFfMrAnLxPcapsResponseSimple]
type pcapNewResponseFfMrAnLxPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseFfMrAnLxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseFfMrAnLxPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1]
type pcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseFfMrAnLxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus string

const (
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusUnknown           PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusSuccess           PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "success"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusPending           PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "pending"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusRunning           PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "running"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusConversionPending PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusConversionRunning PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusComplete          PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "complete"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleStatusFailed            PcapNewResponseFfMrAnLxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseFfMrAnLxPcapsResponseSimpleSystem string

const (
	PcapNewResponseFfMrAnLxPcapsResponseSimpleSystemMagicTransit PcapNewResponseFfMrAnLxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseFfMrAnLxPcapsResponseSimpleType string

const (
	PcapNewResponseFfMrAnLxPcapsResponseSimpleTypeSimple PcapNewResponseFfMrAnLxPcapsResponseSimpleType = "simple"
	PcapNewResponseFfMrAnLxPcapsResponseSimpleTypeFull   PcapNewResponseFfMrAnLxPcapsResponseSimpleType = "full"
)

type PcapNewResponseFfMrAnLxPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseFfMrAnLxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseFfMrAnLxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseFfMrAnLxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseFfMrAnLxPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseFfMrAnLxPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseFfMrAnLxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseFfMrAnLxPcapsResponseFull]
type pcapNewResponseFfMrAnLxPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseFfMrAnLxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseFfMrAnLxPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseFfMrAnLxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseFfMrAnLxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseFfMrAnLxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseFfMrAnLxPcapsResponseFullFilterV1]
type pcapNewResponseFfMrAnLxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseFfMrAnLxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseFfMrAnLxPcapsResponseFullStatus string

const (
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusUnknown           PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "unknown"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusSuccess           PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "success"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusPending           PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "pending"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusRunning           PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "running"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusConversionPending PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusConversionRunning PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusComplete          PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "complete"
	PcapNewResponseFfMrAnLxPcapsResponseFullStatusFailed            PcapNewResponseFfMrAnLxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseFfMrAnLxPcapsResponseFullSystem string

const (
	PcapNewResponseFfMrAnLxPcapsResponseFullSystemMagicTransit PcapNewResponseFfMrAnLxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseFfMrAnLxPcapsResponseFullType string

const (
	PcapNewResponseFfMrAnLxPcapsResponseFullTypeSimple PcapNewResponseFfMrAnLxPcapsResponseFullType = "simple"
	PcapNewResponseFfMrAnLxPcapsResponseFullTypeFull   PcapNewResponseFfMrAnLxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseFfMrAnLxPcapsResponseSimple] or
// [PcapListResponseFfMrAnLxPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseFfMrAnLxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseFfMrAnLxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseFfMrAnLxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseFfMrAnLxPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseFfMrAnLxPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseFfMrAnLxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseFfMrAnLxPcapsResponseSimple]
type pcapListResponseFfMrAnLxPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseFfMrAnLxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseFfMrAnLxPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1]
type pcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseFfMrAnLxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseFfMrAnLxPcapsResponseSimpleStatus string

const (
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusUnknown           PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "unknown"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusSuccess           PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "success"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusPending           PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "pending"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusRunning           PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "running"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusConversionPending PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusConversionRunning PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusComplete          PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "complete"
	PcapListResponseFfMrAnLxPcapsResponseSimpleStatusFailed            PcapListResponseFfMrAnLxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseFfMrAnLxPcapsResponseSimpleSystem string

const (
	PcapListResponseFfMrAnLxPcapsResponseSimpleSystemMagicTransit PcapListResponseFfMrAnLxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseFfMrAnLxPcapsResponseSimpleType string

const (
	PcapListResponseFfMrAnLxPcapsResponseSimpleTypeSimple PcapListResponseFfMrAnLxPcapsResponseSimpleType = "simple"
	PcapListResponseFfMrAnLxPcapsResponseSimpleTypeFull   PcapListResponseFfMrAnLxPcapsResponseSimpleType = "full"
)

type PcapListResponseFfMrAnLxPcapsResponseFull struct {
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
	FilterV1 PcapListResponseFfMrAnLxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseFfMrAnLxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseFfMrAnLxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseFfMrAnLxPcapsResponseFullType `json:"type"`
	JSON pcapListResponseFfMrAnLxPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseFfMrAnLxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseFfMrAnLxPcapsResponseFull]
type pcapListResponseFfMrAnLxPcapsResponseFullJSON struct {
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

func (r *PcapListResponseFfMrAnLxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseFfMrAnLxPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseFfMrAnLxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseFfMrAnLxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseFfMrAnLxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseFfMrAnLxPcapsResponseFullFilterV1]
type pcapListResponseFfMrAnLxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseFfMrAnLxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseFfMrAnLxPcapsResponseFullStatus string

const (
	PcapListResponseFfMrAnLxPcapsResponseFullStatusUnknown           PcapListResponseFfMrAnLxPcapsResponseFullStatus = "unknown"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusSuccess           PcapListResponseFfMrAnLxPcapsResponseFullStatus = "success"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusPending           PcapListResponseFfMrAnLxPcapsResponseFullStatus = "pending"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusRunning           PcapListResponseFfMrAnLxPcapsResponseFullStatus = "running"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusConversionPending PcapListResponseFfMrAnLxPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusConversionRunning PcapListResponseFfMrAnLxPcapsResponseFullStatus = "conversion_running"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusComplete          PcapListResponseFfMrAnLxPcapsResponseFullStatus = "complete"
	PcapListResponseFfMrAnLxPcapsResponseFullStatusFailed            PcapListResponseFfMrAnLxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseFfMrAnLxPcapsResponseFullSystem string

const (
	PcapListResponseFfMrAnLxPcapsResponseFullSystemMagicTransit PcapListResponseFfMrAnLxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseFfMrAnLxPcapsResponseFullType string

const (
	PcapListResponseFfMrAnLxPcapsResponseFullTypeSimple PcapListResponseFfMrAnLxPcapsResponseFullType = "simple"
	PcapListResponseFfMrAnLxPcapsResponseFullTypeFull   PcapListResponseFfMrAnLxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseFfMrAnLxPcapsResponseSimple] or
// [PcapGetResponseFfMrAnLxPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseFfMrAnLxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseFfMrAnLxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseFfMrAnLxPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseFfMrAnLxPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseFfMrAnLxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseFfMrAnLxPcapsResponseSimple]
type pcapGetResponseFfMrAnLxPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseFfMrAnLxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseFfMrAnLxPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1]
type pcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseFfMrAnLxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus string

const (
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusUnknown           PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusSuccess           PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "success"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusPending           PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "pending"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusRunning           PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "running"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusConversionPending PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusConversionRunning PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusComplete          PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "complete"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleStatusFailed            PcapGetResponseFfMrAnLxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseFfMrAnLxPcapsResponseSimpleSystem string

const (
	PcapGetResponseFfMrAnLxPcapsResponseSimpleSystemMagicTransit PcapGetResponseFfMrAnLxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseFfMrAnLxPcapsResponseSimpleType string

const (
	PcapGetResponseFfMrAnLxPcapsResponseSimpleTypeSimple PcapGetResponseFfMrAnLxPcapsResponseSimpleType = "simple"
	PcapGetResponseFfMrAnLxPcapsResponseSimpleTypeFull   PcapGetResponseFfMrAnLxPcapsResponseSimpleType = "full"
)

type PcapGetResponseFfMrAnLxPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseFfMrAnLxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseFfMrAnLxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseFfMrAnLxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseFfMrAnLxPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseFfMrAnLxPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseFfMrAnLxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseFfMrAnLxPcapsResponseFull]
type pcapGetResponseFfMrAnLxPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseFfMrAnLxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseFfMrAnLxPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseFfMrAnLxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseFfMrAnLxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseFfMrAnLxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseFfMrAnLxPcapsResponseFullFilterV1]
type pcapGetResponseFfMrAnLxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseFfMrAnLxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseFfMrAnLxPcapsResponseFullStatus string

const (
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusUnknown           PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "unknown"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusSuccess           PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "success"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusPending           PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "pending"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusRunning           PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "running"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusConversionPending PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusConversionRunning PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusComplete          PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "complete"
	PcapGetResponseFfMrAnLxPcapsResponseFullStatusFailed            PcapGetResponseFfMrAnLxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseFfMrAnLxPcapsResponseFullSystem string

const (
	PcapGetResponseFfMrAnLxPcapsResponseFullSystemMagicTransit PcapGetResponseFfMrAnLxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseFfMrAnLxPcapsResponseFullType string

const (
	PcapGetResponseFfMrAnLxPcapsResponseFullTypeSimple PcapGetResponseFfMrAnLxPcapsResponseFullType = "simple"
	PcapGetResponseFfMrAnLxPcapsResponseFullTypeFull   PcapGetResponseFfMrAnLxPcapsResponseFullType = "full"
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
