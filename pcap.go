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

// Union satisfied by [PcapNewResponseKQpxrDmsPcapsResponseSimple] or
// [PcapNewResponseKQpxrDmsPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseKQpxrDmsPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseKQpxrDmsPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseKQpxrDmsPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseKQpxrDmsPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseKQpxrDmsPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseKQpxrDmsPcapsResponseSimple]
type pcapNewResponseKQpxrDmsPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseKQpxrDmsPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseKQpxrDmsPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1]
type pcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseKQpxrDmsPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus string

const (
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusUnknown           PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusSuccess           PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "success"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusPending           PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "pending"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusRunning           PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "running"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusConversionPending PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusConversionRunning PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusComplete          PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "complete"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleStatusFailed            PcapNewResponseKQpxrDmsPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseKQpxrDmsPcapsResponseSimpleSystem string

const (
	PcapNewResponseKQpxrDmsPcapsResponseSimpleSystemMagicTransit PcapNewResponseKQpxrDmsPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseKQpxrDmsPcapsResponseSimpleType string

const (
	PcapNewResponseKQpxrDmsPcapsResponseSimpleTypeSimple PcapNewResponseKQpxrDmsPcapsResponseSimpleType = "simple"
	PcapNewResponseKQpxrDmsPcapsResponseSimpleTypeFull   PcapNewResponseKQpxrDmsPcapsResponseSimpleType = "full"
)

type PcapNewResponseKQpxrDmsPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseKQpxrDmsPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseKQpxrDmsPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseKQpxrDmsPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseKQpxrDmsPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseKQpxrDmsPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseKQpxrDmsPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseKQpxrDmsPcapsResponseFull]
type pcapNewResponseKQpxrDmsPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseKQpxrDmsPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseKQpxrDmsPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseKQpxrDmsPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseKQpxrDmsPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseKQpxrDmsPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseKQpxrDmsPcapsResponseFullFilterV1]
type pcapNewResponseKQpxrDmsPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseKQpxrDmsPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseKQpxrDmsPcapsResponseFullStatus string

const (
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusUnknown           PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "unknown"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusSuccess           PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "success"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusPending           PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "pending"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusRunning           PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "running"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusConversionPending PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusConversionRunning PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusComplete          PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "complete"
	PcapNewResponseKQpxrDmsPcapsResponseFullStatusFailed            PcapNewResponseKQpxrDmsPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseKQpxrDmsPcapsResponseFullSystem string

const (
	PcapNewResponseKQpxrDmsPcapsResponseFullSystemMagicTransit PcapNewResponseKQpxrDmsPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseKQpxrDmsPcapsResponseFullType string

const (
	PcapNewResponseKQpxrDmsPcapsResponseFullTypeSimple PcapNewResponseKQpxrDmsPcapsResponseFullType = "simple"
	PcapNewResponseKQpxrDmsPcapsResponseFullTypeFull   PcapNewResponseKQpxrDmsPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseKQpxrDmsPcapsResponseSimple] or
// [PcapListResponseKQpxrDmsPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseKQpxrDmsPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseKQpxrDmsPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseKQpxrDmsPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseKQpxrDmsPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseKQpxrDmsPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseKQpxrDmsPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseKQpxrDmsPcapsResponseSimple]
type pcapListResponseKQpxrDmsPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseKQpxrDmsPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseKQpxrDmsPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1]
type pcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseKQpxrDmsPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseKQpxrDmsPcapsResponseSimpleStatus string

const (
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusUnknown           PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "unknown"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusSuccess           PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "success"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusPending           PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "pending"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusRunning           PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "running"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusConversionPending PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusConversionRunning PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusComplete          PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "complete"
	PcapListResponseKQpxrDmsPcapsResponseSimpleStatusFailed            PcapListResponseKQpxrDmsPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseKQpxrDmsPcapsResponseSimpleSystem string

const (
	PcapListResponseKQpxrDmsPcapsResponseSimpleSystemMagicTransit PcapListResponseKQpxrDmsPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseKQpxrDmsPcapsResponseSimpleType string

const (
	PcapListResponseKQpxrDmsPcapsResponseSimpleTypeSimple PcapListResponseKQpxrDmsPcapsResponseSimpleType = "simple"
	PcapListResponseKQpxrDmsPcapsResponseSimpleTypeFull   PcapListResponseKQpxrDmsPcapsResponseSimpleType = "full"
)

type PcapListResponseKQpxrDmsPcapsResponseFull struct {
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
	FilterV1 PcapListResponseKQpxrDmsPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseKQpxrDmsPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseKQpxrDmsPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseKQpxrDmsPcapsResponseFullType `json:"type"`
	JSON pcapListResponseKQpxrDmsPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseKQpxrDmsPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseKQpxrDmsPcapsResponseFull]
type pcapListResponseKQpxrDmsPcapsResponseFullJSON struct {
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

func (r *PcapListResponseKQpxrDmsPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseKQpxrDmsPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseKQpxrDmsPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseKQpxrDmsPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseKQpxrDmsPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseKQpxrDmsPcapsResponseFullFilterV1]
type pcapListResponseKQpxrDmsPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseKQpxrDmsPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseKQpxrDmsPcapsResponseFullStatus string

const (
	PcapListResponseKQpxrDmsPcapsResponseFullStatusUnknown           PcapListResponseKQpxrDmsPcapsResponseFullStatus = "unknown"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusSuccess           PcapListResponseKQpxrDmsPcapsResponseFullStatus = "success"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusPending           PcapListResponseKQpxrDmsPcapsResponseFullStatus = "pending"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusRunning           PcapListResponseKQpxrDmsPcapsResponseFullStatus = "running"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusConversionPending PcapListResponseKQpxrDmsPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusConversionRunning PcapListResponseKQpxrDmsPcapsResponseFullStatus = "conversion_running"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusComplete          PcapListResponseKQpxrDmsPcapsResponseFullStatus = "complete"
	PcapListResponseKQpxrDmsPcapsResponseFullStatusFailed            PcapListResponseKQpxrDmsPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseKQpxrDmsPcapsResponseFullSystem string

const (
	PcapListResponseKQpxrDmsPcapsResponseFullSystemMagicTransit PcapListResponseKQpxrDmsPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseKQpxrDmsPcapsResponseFullType string

const (
	PcapListResponseKQpxrDmsPcapsResponseFullTypeSimple PcapListResponseKQpxrDmsPcapsResponseFullType = "simple"
	PcapListResponseKQpxrDmsPcapsResponseFullTypeFull   PcapListResponseKQpxrDmsPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseKQpxrDmsPcapsResponseSimple] or
// [PcapGetResponseKQpxrDmsPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseKQpxrDmsPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseKQpxrDmsPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseKQpxrDmsPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseKQpxrDmsPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseKQpxrDmsPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseKQpxrDmsPcapsResponseSimple]
type pcapGetResponseKQpxrDmsPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseKQpxrDmsPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseKQpxrDmsPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1]
type pcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseKQpxrDmsPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus string

const (
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusUnknown           PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusSuccess           PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "success"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusPending           PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "pending"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusRunning           PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "running"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusConversionPending PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusConversionRunning PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusComplete          PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "complete"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleStatusFailed            PcapGetResponseKQpxrDmsPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseKQpxrDmsPcapsResponseSimpleSystem string

const (
	PcapGetResponseKQpxrDmsPcapsResponseSimpleSystemMagicTransit PcapGetResponseKQpxrDmsPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseKQpxrDmsPcapsResponseSimpleType string

const (
	PcapGetResponseKQpxrDmsPcapsResponseSimpleTypeSimple PcapGetResponseKQpxrDmsPcapsResponseSimpleType = "simple"
	PcapGetResponseKQpxrDmsPcapsResponseSimpleTypeFull   PcapGetResponseKQpxrDmsPcapsResponseSimpleType = "full"
)

type PcapGetResponseKQpxrDmsPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseKQpxrDmsPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseKQpxrDmsPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseKQpxrDmsPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseKQpxrDmsPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseKQpxrDmsPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseKQpxrDmsPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseKQpxrDmsPcapsResponseFull]
type pcapGetResponseKQpxrDmsPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseKQpxrDmsPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseKQpxrDmsPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseKQpxrDmsPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseKQpxrDmsPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseKQpxrDmsPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseKQpxrDmsPcapsResponseFullFilterV1]
type pcapGetResponseKQpxrDmsPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseKQpxrDmsPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseKQpxrDmsPcapsResponseFullStatus string

const (
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusUnknown           PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "unknown"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusSuccess           PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "success"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusPending           PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "pending"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusRunning           PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "running"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusConversionPending PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusConversionRunning PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusComplete          PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "complete"
	PcapGetResponseKQpxrDmsPcapsResponseFullStatusFailed            PcapGetResponseKQpxrDmsPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseKQpxrDmsPcapsResponseFullSystem string

const (
	PcapGetResponseKQpxrDmsPcapsResponseFullSystemMagicTransit PcapGetResponseKQpxrDmsPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseKQpxrDmsPcapsResponseFullType string

const (
	PcapGetResponseKQpxrDmsPcapsResponseFullTypeSimple PcapGetResponseKQpxrDmsPcapsResponseFullType = "simple"
	PcapGetResponseKQpxrDmsPcapsResponseFullTypeFull   PcapGetResponseKQpxrDmsPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsKQpxrDmsPcapsRequestSimple],
// [PcapNewParamsKQpxrDmsPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsKQpxrDmsPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsKQpxrDmsPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsKQpxrDmsPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsKQpxrDmsPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsKQpxrDmsPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsKQpxrDmsPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsKQpxrDmsPcapsRequestSimpleSystem string

const (
	PcapNewParamsKQpxrDmsPcapsRequestSimpleSystemMagicTransit PcapNewParamsKQpxrDmsPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsKQpxrDmsPcapsRequestSimpleType string

const (
	PcapNewParamsKQpxrDmsPcapsRequestSimpleTypeSimple PcapNewParamsKQpxrDmsPcapsRequestSimpleType = "simple"
	PcapNewParamsKQpxrDmsPcapsRequestSimpleTypeFull   PcapNewParamsKQpxrDmsPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsKQpxrDmsPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsKQpxrDmsPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsKQpxrDmsPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsKQpxrDmsPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsKQpxrDmsPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsKQpxrDmsPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsKQpxrDmsPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsKQpxrDmsPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsKQpxrDmsPcapsRequestFullSystem string

const (
	PcapNewParamsKQpxrDmsPcapsRequestFullSystemMagicTransit PcapNewParamsKQpxrDmsPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsKQpxrDmsPcapsRequestFullType string

const (
	PcapNewParamsKQpxrDmsPcapsRequestFullTypeSimple PcapNewParamsKQpxrDmsPcapsRequestFullType = "simple"
	PcapNewParamsKQpxrDmsPcapsRequestFullTypeFull   PcapNewParamsKQpxrDmsPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsKQpxrDmsPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsKQpxrDmsPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
