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

// Union satisfied by [PcapNewResponseUholmugaPcapsResponseSimple] or
// [PcapNewResponseUholmugaPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseUholmugaPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseUholmugaPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseUholmugaPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseUholmugaPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseUholmugaPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseUholmugaPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseUholmugaPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseUholmugaPcapsResponseSimple]
type pcapNewResponseUholmugaPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseUholmugaPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseUholmugaPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseUholmugaPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseUholmugaPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseUholmugaPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseUholmugaPcapsResponseSimpleFilterV1]
type pcapNewResponseUholmugaPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseUholmugaPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseUholmugaPcapsResponseSimpleStatus string

const (
	PcapNewResponseUholmugaPcapsResponseSimpleStatusUnknown           PcapNewResponseUholmugaPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusSuccess           PcapNewResponseUholmugaPcapsResponseSimpleStatus = "success"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusPending           PcapNewResponseUholmugaPcapsResponseSimpleStatus = "pending"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusRunning           PcapNewResponseUholmugaPcapsResponseSimpleStatus = "running"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusConversionPending PcapNewResponseUholmugaPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusConversionRunning PcapNewResponseUholmugaPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusComplete          PcapNewResponseUholmugaPcapsResponseSimpleStatus = "complete"
	PcapNewResponseUholmugaPcapsResponseSimpleStatusFailed            PcapNewResponseUholmugaPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseUholmugaPcapsResponseSimpleSystem string

const (
	PcapNewResponseUholmugaPcapsResponseSimpleSystemMagicTransit PcapNewResponseUholmugaPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseUholmugaPcapsResponseSimpleType string

const (
	PcapNewResponseUholmugaPcapsResponseSimpleTypeSimple PcapNewResponseUholmugaPcapsResponseSimpleType = "simple"
	PcapNewResponseUholmugaPcapsResponseSimpleTypeFull   PcapNewResponseUholmugaPcapsResponseSimpleType = "full"
)

type PcapNewResponseUholmugaPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseUholmugaPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseUholmugaPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseUholmugaPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseUholmugaPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseUholmugaPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseUholmugaPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseUholmugaPcapsResponseFull]
type pcapNewResponseUholmugaPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseUholmugaPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseUholmugaPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseUholmugaPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseUholmugaPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseUholmugaPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseUholmugaPcapsResponseFullFilterV1]
type pcapNewResponseUholmugaPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseUholmugaPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseUholmugaPcapsResponseFullStatus string

const (
	PcapNewResponseUholmugaPcapsResponseFullStatusUnknown           PcapNewResponseUholmugaPcapsResponseFullStatus = "unknown"
	PcapNewResponseUholmugaPcapsResponseFullStatusSuccess           PcapNewResponseUholmugaPcapsResponseFullStatus = "success"
	PcapNewResponseUholmugaPcapsResponseFullStatusPending           PcapNewResponseUholmugaPcapsResponseFullStatus = "pending"
	PcapNewResponseUholmugaPcapsResponseFullStatusRunning           PcapNewResponseUholmugaPcapsResponseFullStatus = "running"
	PcapNewResponseUholmugaPcapsResponseFullStatusConversionPending PcapNewResponseUholmugaPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseUholmugaPcapsResponseFullStatusConversionRunning PcapNewResponseUholmugaPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseUholmugaPcapsResponseFullStatusComplete          PcapNewResponseUholmugaPcapsResponseFullStatus = "complete"
	PcapNewResponseUholmugaPcapsResponseFullStatusFailed            PcapNewResponseUholmugaPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseUholmugaPcapsResponseFullSystem string

const (
	PcapNewResponseUholmugaPcapsResponseFullSystemMagicTransit PcapNewResponseUholmugaPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseUholmugaPcapsResponseFullType string

const (
	PcapNewResponseUholmugaPcapsResponseFullTypeSimple PcapNewResponseUholmugaPcapsResponseFullType = "simple"
	PcapNewResponseUholmugaPcapsResponseFullTypeFull   PcapNewResponseUholmugaPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseUholmugaPcapsResponseSimple] or
// [PcapListResponseUholmugaPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseUholmugaPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseUholmugaPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseUholmugaPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseUholmugaPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseUholmugaPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseUholmugaPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseUholmugaPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseUholmugaPcapsResponseSimple]
type pcapListResponseUholmugaPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseUholmugaPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseUholmugaPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseUholmugaPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseUholmugaPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseUholmugaPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseUholmugaPcapsResponseSimpleFilterV1]
type pcapListResponseUholmugaPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseUholmugaPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseUholmugaPcapsResponseSimpleStatus string

const (
	PcapListResponseUholmugaPcapsResponseSimpleStatusUnknown           PcapListResponseUholmugaPcapsResponseSimpleStatus = "unknown"
	PcapListResponseUholmugaPcapsResponseSimpleStatusSuccess           PcapListResponseUholmugaPcapsResponseSimpleStatus = "success"
	PcapListResponseUholmugaPcapsResponseSimpleStatusPending           PcapListResponseUholmugaPcapsResponseSimpleStatus = "pending"
	PcapListResponseUholmugaPcapsResponseSimpleStatusRunning           PcapListResponseUholmugaPcapsResponseSimpleStatus = "running"
	PcapListResponseUholmugaPcapsResponseSimpleStatusConversionPending PcapListResponseUholmugaPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseUholmugaPcapsResponseSimpleStatusConversionRunning PcapListResponseUholmugaPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseUholmugaPcapsResponseSimpleStatusComplete          PcapListResponseUholmugaPcapsResponseSimpleStatus = "complete"
	PcapListResponseUholmugaPcapsResponseSimpleStatusFailed            PcapListResponseUholmugaPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseUholmugaPcapsResponseSimpleSystem string

const (
	PcapListResponseUholmugaPcapsResponseSimpleSystemMagicTransit PcapListResponseUholmugaPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseUholmugaPcapsResponseSimpleType string

const (
	PcapListResponseUholmugaPcapsResponseSimpleTypeSimple PcapListResponseUholmugaPcapsResponseSimpleType = "simple"
	PcapListResponseUholmugaPcapsResponseSimpleTypeFull   PcapListResponseUholmugaPcapsResponseSimpleType = "full"
)

type PcapListResponseUholmugaPcapsResponseFull struct {
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
	FilterV1 PcapListResponseUholmugaPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseUholmugaPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseUholmugaPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseUholmugaPcapsResponseFullType `json:"type"`
	JSON pcapListResponseUholmugaPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseUholmugaPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseUholmugaPcapsResponseFull]
type pcapListResponseUholmugaPcapsResponseFullJSON struct {
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

func (r *PcapListResponseUholmugaPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseUholmugaPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseUholmugaPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseUholmugaPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseUholmugaPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseUholmugaPcapsResponseFullFilterV1]
type pcapListResponseUholmugaPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseUholmugaPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseUholmugaPcapsResponseFullStatus string

const (
	PcapListResponseUholmugaPcapsResponseFullStatusUnknown           PcapListResponseUholmugaPcapsResponseFullStatus = "unknown"
	PcapListResponseUholmugaPcapsResponseFullStatusSuccess           PcapListResponseUholmugaPcapsResponseFullStatus = "success"
	PcapListResponseUholmugaPcapsResponseFullStatusPending           PcapListResponseUholmugaPcapsResponseFullStatus = "pending"
	PcapListResponseUholmugaPcapsResponseFullStatusRunning           PcapListResponseUholmugaPcapsResponseFullStatus = "running"
	PcapListResponseUholmugaPcapsResponseFullStatusConversionPending PcapListResponseUholmugaPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseUholmugaPcapsResponseFullStatusConversionRunning PcapListResponseUholmugaPcapsResponseFullStatus = "conversion_running"
	PcapListResponseUholmugaPcapsResponseFullStatusComplete          PcapListResponseUholmugaPcapsResponseFullStatus = "complete"
	PcapListResponseUholmugaPcapsResponseFullStatusFailed            PcapListResponseUholmugaPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseUholmugaPcapsResponseFullSystem string

const (
	PcapListResponseUholmugaPcapsResponseFullSystemMagicTransit PcapListResponseUholmugaPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseUholmugaPcapsResponseFullType string

const (
	PcapListResponseUholmugaPcapsResponseFullTypeSimple PcapListResponseUholmugaPcapsResponseFullType = "simple"
	PcapListResponseUholmugaPcapsResponseFullTypeFull   PcapListResponseUholmugaPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseUholmugaPcapsResponseSimple] or
// [PcapGetResponseUholmugaPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseUholmugaPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseUholmugaPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseUholmugaPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseUholmugaPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseUholmugaPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseUholmugaPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseUholmugaPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseUholmugaPcapsResponseSimple]
type pcapGetResponseUholmugaPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseUholmugaPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseUholmugaPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseUholmugaPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseUholmugaPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseUholmugaPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseUholmugaPcapsResponseSimpleFilterV1]
type pcapGetResponseUholmugaPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseUholmugaPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseUholmugaPcapsResponseSimpleStatus string

const (
	PcapGetResponseUholmugaPcapsResponseSimpleStatusUnknown           PcapGetResponseUholmugaPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusSuccess           PcapGetResponseUholmugaPcapsResponseSimpleStatus = "success"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusPending           PcapGetResponseUholmugaPcapsResponseSimpleStatus = "pending"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusRunning           PcapGetResponseUholmugaPcapsResponseSimpleStatus = "running"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusConversionPending PcapGetResponseUholmugaPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusConversionRunning PcapGetResponseUholmugaPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusComplete          PcapGetResponseUholmugaPcapsResponseSimpleStatus = "complete"
	PcapGetResponseUholmugaPcapsResponseSimpleStatusFailed            PcapGetResponseUholmugaPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseUholmugaPcapsResponseSimpleSystem string

const (
	PcapGetResponseUholmugaPcapsResponseSimpleSystemMagicTransit PcapGetResponseUholmugaPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseUholmugaPcapsResponseSimpleType string

const (
	PcapGetResponseUholmugaPcapsResponseSimpleTypeSimple PcapGetResponseUholmugaPcapsResponseSimpleType = "simple"
	PcapGetResponseUholmugaPcapsResponseSimpleTypeFull   PcapGetResponseUholmugaPcapsResponseSimpleType = "full"
)

type PcapGetResponseUholmugaPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseUholmugaPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseUholmugaPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseUholmugaPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseUholmugaPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseUholmugaPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseUholmugaPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseUholmugaPcapsResponseFull]
type pcapGetResponseUholmugaPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseUholmugaPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseUholmugaPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseUholmugaPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseUholmugaPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseUholmugaPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseUholmugaPcapsResponseFullFilterV1]
type pcapGetResponseUholmugaPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseUholmugaPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseUholmugaPcapsResponseFullStatus string

const (
	PcapGetResponseUholmugaPcapsResponseFullStatusUnknown           PcapGetResponseUholmugaPcapsResponseFullStatus = "unknown"
	PcapGetResponseUholmugaPcapsResponseFullStatusSuccess           PcapGetResponseUholmugaPcapsResponseFullStatus = "success"
	PcapGetResponseUholmugaPcapsResponseFullStatusPending           PcapGetResponseUholmugaPcapsResponseFullStatus = "pending"
	PcapGetResponseUholmugaPcapsResponseFullStatusRunning           PcapGetResponseUholmugaPcapsResponseFullStatus = "running"
	PcapGetResponseUholmugaPcapsResponseFullStatusConversionPending PcapGetResponseUholmugaPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseUholmugaPcapsResponseFullStatusConversionRunning PcapGetResponseUholmugaPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseUholmugaPcapsResponseFullStatusComplete          PcapGetResponseUholmugaPcapsResponseFullStatus = "complete"
	PcapGetResponseUholmugaPcapsResponseFullStatusFailed            PcapGetResponseUholmugaPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseUholmugaPcapsResponseFullSystem string

const (
	PcapGetResponseUholmugaPcapsResponseFullSystemMagicTransit PcapGetResponseUholmugaPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseUholmugaPcapsResponseFullType string

const (
	PcapGetResponseUholmugaPcapsResponseFullTypeSimple PcapGetResponseUholmugaPcapsResponseFullType = "simple"
	PcapGetResponseUholmugaPcapsResponseFullTypeFull   PcapGetResponseUholmugaPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsUholmugaPcapsRequestSimple],
// [PcapNewParamsUholmugaPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsUholmugaPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsUholmugaPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsUholmugaPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsUholmugaPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsUholmugaPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsUholmugaPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsUholmugaPcapsRequestSimpleSystem string

const (
	PcapNewParamsUholmugaPcapsRequestSimpleSystemMagicTransit PcapNewParamsUholmugaPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsUholmugaPcapsRequestSimpleType string

const (
	PcapNewParamsUholmugaPcapsRequestSimpleTypeSimple PcapNewParamsUholmugaPcapsRequestSimpleType = "simple"
	PcapNewParamsUholmugaPcapsRequestSimpleTypeFull   PcapNewParamsUholmugaPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsUholmugaPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsUholmugaPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsUholmugaPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsUholmugaPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsUholmugaPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsUholmugaPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsUholmugaPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsUholmugaPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsUholmugaPcapsRequestFullSystem string

const (
	PcapNewParamsUholmugaPcapsRequestFullSystemMagicTransit PcapNewParamsUholmugaPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsUholmugaPcapsRequestFullType string

const (
	PcapNewParamsUholmugaPcapsRequestFullTypeSimple PcapNewParamsUholmugaPcapsRequestFullType = "simple"
	PcapNewParamsUholmugaPcapsRequestFullTypeFull   PcapNewParamsUholmugaPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsUholmugaPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsUholmugaPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
