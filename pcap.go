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

// Union satisfied by [PcapNewResponseTy58MTbiPcapsResponseSimple] or
// [PcapNewResponseTy58MTbiPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseTy58MTbiPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseTy58MTbiPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseTy58MTbiPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseTy58MTbiPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseTy58MTbiPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseTy58MTbiPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseTy58MTbiPcapsResponseSimple]
type pcapNewResponseTy58MTbiPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseTy58MTbiPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseTy58MTbiPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1]
type pcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseTy58MTbiPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseTy58MTbiPcapsResponseSimpleStatus string

const (
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusUnknown           PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusSuccess           PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "success"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusPending           PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "pending"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusRunning           PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "running"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusConversionPending PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusConversionRunning PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusComplete          PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "complete"
	PcapNewResponseTy58MTbiPcapsResponseSimpleStatusFailed            PcapNewResponseTy58MTbiPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseTy58MTbiPcapsResponseSimpleSystem string

const (
	PcapNewResponseTy58MTbiPcapsResponseSimpleSystemMagicTransit PcapNewResponseTy58MTbiPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseTy58MTbiPcapsResponseSimpleType string

const (
	PcapNewResponseTy58MTbiPcapsResponseSimpleTypeSimple PcapNewResponseTy58MTbiPcapsResponseSimpleType = "simple"
	PcapNewResponseTy58MTbiPcapsResponseSimpleTypeFull   PcapNewResponseTy58MTbiPcapsResponseSimpleType = "full"
)

type PcapNewResponseTy58MTbiPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseTy58MTbiPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseTy58MTbiPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseTy58MTbiPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseTy58MTbiPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseTy58MTbiPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseTy58MTbiPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseTy58MTbiPcapsResponseFull]
type pcapNewResponseTy58MTbiPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseTy58MTbiPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseTy58MTbiPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseTy58MTbiPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseTy58MTbiPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseTy58MTbiPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseTy58MTbiPcapsResponseFullFilterV1]
type pcapNewResponseTy58MTbiPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseTy58MTbiPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseTy58MTbiPcapsResponseFullStatus string

const (
	PcapNewResponseTy58MTbiPcapsResponseFullStatusUnknown           PcapNewResponseTy58MTbiPcapsResponseFullStatus = "unknown"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusSuccess           PcapNewResponseTy58MTbiPcapsResponseFullStatus = "success"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusPending           PcapNewResponseTy58MTbiPcapsResponseFullStatus = "pending"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusRunning           PcapNewResponseTy58MTbiPcapsResponseFullStatus = "running"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusConversionPending PcapNewResponseTy58MTbiPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusConversionRunning PcapNewResponseTy58MTbiPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusComplete          PcapNewResponseTy58MTbiPcapsResponseFullStatus = "complete"
	PcapNewResponseTy58MTbiPcapsResponseFullStatusFailed            PcapNewResponseTy58MTbiPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseTy58MTbiPcapsResponseFullSystem string

const (
	PcapNewResponseTy58MTbiPcapsResponseFullSystemMagicTransit PcapNewResponseTy58MTbiPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseTy58MTbiPcapsResponseFullType string

const (
	PcapNewResponseTy58MTbiPcapsResponseFullTypeSimple PcapNewResponseTy58MTbiPcapsResponseFullType = "simple"
	PcapNewResponseTy58MTbiPcapsResponseFullTypeFull   PcapNewResponseTy58MTbiPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseTy58MTbiPcapsResponseSimple] or
// [PcapListResponseTy58MTbiPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseTy58MTbiPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseTy58MTbiPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseTy58MTbiPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseTy58MTbiPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseTy58MTbiPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseTy58MTbiPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseTy58MTbiPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseTy58MTbiPcapsResponseSimple]
type pcapListResponseTy58MTbiPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseTy58MTbiPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseTy58MTbiPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseTy58MTbiPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseTy58MTbiPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseTy58MTbiPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseTy58MTbiPcapsResponseSimpleFilterV1]
type pcapListResponseTy58MTbiPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseTy58MTbiPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseTy58MTbiPcapsResponseSimpleStatus string

const (
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusUnknown           PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "unknown"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusSuccess           PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "success"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusPending           PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "pending"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusRunning           PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "running"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusConversionPending PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusConversionRunning PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusComplete          PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "complete"
	PcapListResponseTy58MTbiPcapsResponseSimpleStatusFailed            PcapListResponseTy58MTbiPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseTy58MTbiPcapsResponseSimpleSystem string

const (
	PcapListResponseTy58MTbiPcapsResponseSimpleSystemMagicTransit PcapListResponseTy58MTbiPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseTy58MTbiPcapsResponseSimpleType string

const (
	PcapListResponseTy58MTbiPcapsResponseSimpleTypeSimple PcapListResponseTy58MTbiPcapsResponseSimpleType = "simple"
	PcapListResponseTy58MTbiPcapsResponseSimpleTypeFull   PcapListResponseTy58MTbiPcapsResponseSimpleType = "full"
)

type PcapListResponseTy58MTbiPcapsResponseFull struct {
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
	FilterV1 PcapListResponseTy58MTbiPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseTy58MTbiPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseTy58MTbiPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseTy58MTbiPcapsResponseFullType `json:"type"`
	JSON pcapListResponseTy58MTbiPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseTy58MTbiPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseTy58MTbiPcapsResponseFull]
type pcapListResponseTy58MTbiPcapsResponseFullJSON struct {
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

func (r *PcapListResponseTy58MTbiPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseTy58MTbiPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseTy58MTbiPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseTy58MTbiPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseTy58MTbiPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseTy58MTbiPcapsResponseFullFilterV1]
type pcapListResponseTy58MTbiPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseTy58MTbiPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseTy58MTbiPcapsResponseFullStatus string

const (
	PcapListResponseTy58MTbiPcapsResponseFullStatusUnknown           PcapListResponseTy58MTbiPcapsResponseFullStatus = "unknown"
	PcapListResponseTy58MTbiPcapsResponseFullStatusSuccess           PcapListResponseTy58MTbiPcapsResponseFullStatus = "success"
	PcapListResponseTy58MTbiPcapsResponseFullStatusPending           PcapListResponseTy58MTbiPcapsResponseFullStatus = "pending"
	PcapListResponseTy58MTbiPcapsResponseFullStatusRunning           PcapListResponseTy58MTbiPcapsResponseFullStatus = "running"
	PcapListResponseTy58MTbiPcapsResponseFullStatusConversionPending PcapListResponseTy58MTbiPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseTy58MTbiPcapsResponseFullStatusConversionRunning PcapListResponseTy58MTbiPcapsResponseFullStatus = "conversion_running"
	PcapListResponseTy58MTbiPcapsResponseFullStatusComplete          PcapListResponseTy58MTbiPcapsResponseFullStatus = "complete"
	PcapListResponseTy58MTbiPcapsResponseFullStatusFailed            PcapListResponseTy58MTbiPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseTy58MTbiPcapsResponseFullSystem string

const (
	PcapListResponseTy58MTbiPcapsResponseFullSystemMagicTransit PcapListResponseTy58MTbiPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseTy58MTbiPcapsResponseFullType string

const (
	PcapListResponseTy58MTbiPcapsResponseFullTypeSimple PcapListResponseTy58MTbiPcapsResponseFullType = "simple"
	PcapListResponseTy58MTbiPcapsResponseFullTypeFull   PcapListResponseTy58MTbiPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseTy58MTbiPcapsResponseSimple] or
// [PcapGetResponseTy58MTbiPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseTy58MTbiPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseTy58MTbiPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseTy58MTbiPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseTy58MTbiPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseTy58MTbiPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseTy58MTbiPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseTy58MTbiPcapsResponseSimple]
type pcapGetResponseTy58MTbiPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseTy58MTbiPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseTy58MTbiPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1]
type pcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseTy58MTbiPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseTy58MTbiPcapsResponseSimpleStatus string

const (
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusUnknown           PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusSuccess           PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "success"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusPending           PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "pending"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusRunning           PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "running"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusConversionPending PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusConversionRunning PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusComplete          PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "complete"
	PcapGetResponseTy58MTbiPcapsResponseSimpleStatusFailed            PcapGetResponseTy58MTbiPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseTy58MTbiPcapsResponseSimpleSystem string

const (
	PcapGetResponseTy58MTbiPcapsResponseSimpleSystemMagicTransit PcapGetResponseTy58MTbiPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseTy58MTbiPcapsResponseSimpleType string

const (
	PcapGetResponseTy58MTbiPcapsResponseSimpleTypeSimple PcapGetResponseTy58MTbiPcapsResponseSimpleType = "simple"
	PcapGetResponseTy58MTbiPcapsResponseSimpleTypeFull   PcapGetResponseTy58MTbiPcapsResponseSimpleType = "full"
)

type PcapGetResponseTy58MTbiPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseTy58MTbiPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseTy58MTbiPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseTy58MTbiPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseTy58MTbiPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseTy58MTbiPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseTy58MTbiPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseTy58MTbiPcapsResponseFull]
type pcapGetResponseTy58MTbiPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseTy58MTbiPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseTy58MTbiPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseTy58MTbiPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseTy58MTbiPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseTy58MTbiPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseTy58MTbiPcapsResponseFullFilterV1]
type pcapGetResponseTy58MTbiPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseTy58MTbiPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseTy58MTbiPcapsResponseFullStatus string

const (
	PcapGetResponseTy58MTbiPcapsResponseFullStatusUnknown           PcapGetResponseTy58MTbiPcapsResponseFullStatus = "unknown"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusSuccess           PcapGetResponseTy58MTbiPcapsResponseFullStatus = "success"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusPending           PcapGetResponseTy58MTbiPcapsResponseFullStatus = "pending"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusRunning           PcapGetResponseTy58MTbiPcapsResponseFullStatus = "running"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusConversionPending PcapGetResponseTy58MTbiPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusConversionRunning PcapGetResponseTy58MTbiPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusComplete          PcapGetResponseTy58MTbiPcapsResponseFullStatus = "complete"
	PcapGetResponseTy58MTbiPcapsResponseFullStatusFailed            PcapGetResponseTy58MTbiPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseTy58MTbiPcapsResponseFullSystem string

const (
	PcapGetResponseTy58MTbiPcapsResponseFullSystemMagicTransit PcapGetResponseTy58MTbiPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseTy58MTbiPcapsResponseFullType string

const (
	PcapGetResponseTy58MTbiPcapsResponseFullTypeSimple PcapGetResponseTy58MTbiPcapsResponseFullType = "simple"
	PcapGetResponseTy58MTbiPcapsResponseFullTypeFull   PcapGetResponseTy58MTbiPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsTy58MTbiPcapsRequestSimple],
// [PcapNewParamsTy58MTbiPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsTy58MTbiPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsTy58MTbiPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsTy58MTbiPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsTy58MTbiPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsTy58MTbiPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsTy58MTbiPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsTy58MTbiPcapsRequestSimpleSystem string

const (
	PcapNewParamsTy58MTbiPcapsRequestSimpleSystemMagicTransit PcapNewParamsTy58MTbiPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsTy58MTbiPcapsRequestSimpleType string

const (
	PcapNewParamsTy58MTbiPcapsRequestSimpleTypeSimple PcapNewParamsTy58MTbiPcapsRequestSimpleType = "simple"
	PcapNewParamsTy58MTbiPcapsRequestSimpleTypeFull   PcapNewParamsTy58MTbiPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsTy58MTbiPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsTy58MTbiPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsTy58MTbiPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsTy58MTbiPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsTy58MTbiPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsTy58MTbiPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsTy58MTbiPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsTy58MTbiPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsTy58MTbiPcapsRequestFullSystem string

const (
	PcapNewParamsTy58MTbiPcapsRequestFullSystemMagicTransit PcapNewParamsTy58MTbiPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsTy58MTbiPcapsRequestFullType string

const (
	PcapNewParamsTy58MTbiPcapsRequestFullTypeSimple PcapNewParamsTy58MTbiPcapsRequestFullType = "simple"
	PcapNewParamsTy58MTbiPcapsRequestFullTypeFull   PcapNewParamsTy58MTbiPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsTy58MTbiPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsTy58MTbiPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
