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

// Union satisfied by [PcapNewResponseXlQyLgQwPcapsResponseSimple] or
// [PcapNewResponseXlQyLgQwPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseXlQyLgQwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseXlQyLgQwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseXlQyLgQwPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseXlQyLgQwPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseXlQyLgQwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseXlQyLgQwPcapsResponseSimple]
type pcapNewResponseXlQyLgQwPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseXlQyLgQwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseXlQyLgQwPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1]
type pcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseXlQyLgQwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus string

const (
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusUnknown           PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusSuccess           PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "success"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusPending           PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "pending"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusRunning           PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "running"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusConversionPending PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusConversionRunning PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusComplete          PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "complete"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleStatusFailed            PcapNewResponseXlQyLgQwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseXlQyLgQwPcapsResponseSimpleSystem string

const (
	PcapNewResponseXlQyLgQwPcapsResponseSimpleSystemMagicTransit PcapNewResponseXlQyLgQwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseXlQyLgQwPcapsResponseSimpleType string

const (
	PcapNewResponseXlQyLgQwPcapsResponseSimpleTypeSimple PcapNewResponseXlQyLgQwPcapsResponseSimpleType = "simple"
	PcapNewResponseXlQyLgQwPcapsResponseSimpleTypeFull   PcapNewResponseXlQyLgQwPcapsResponseSimpleType = "full"
)

type PcapNewResponseXlQyLgQwPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseXlQyLgQwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseXlQyLgQwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseXlQyLgQwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseXlQyLgQwPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseXlQyLgQwPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseXlQyLgQwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseXlQyLgQwPcapsResponseFull]
type pcapNewResponseXlQyLgQwPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseXlQyLgQwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseXlQyLgQwPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseXlQyLgQwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseXlQyLgQwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseXlQyLgQwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseXlQyLgQwPcapsResponseFullFilterV1]
type pcapNewResponseXlQyLgQwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseXlQyLgQwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseXlQyLgQwPcapsResponseFullStatus string

const (
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusUnknown           PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "unknown"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusSuccess           PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "success"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusPending           PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "pending"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusRunning           PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "running"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusConversionPending PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusConversionRunning PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusComplete          PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "complete"
	PcapNewResponseXlQyLgQwPcapsResponseFullStatusFailed            PcapNewResponseXlQyLgQwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseXlQyLgQwPcapsResponseFullSystem string

const (
	PcapNewResponseXlQyLgQwPcapsResponseFullSystemMagicTransit PcapNewResponseXlQyLgQwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseXlQyLgQwPcapsResponseFullType string

const (
	PcapNewResponseXlQyLgQwPcapsResponseFullTypeSimple PcapNewResponseXlQyLgQwPcapsResponseFullType = "simple"
	PcapNewResponseXlQyLgQwPcapsResponseFullTypeFull   PcapNewResponseXlQyLgQwPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseXlQyLgQwPcapsResponseSimple] or
// [PcapListResponseXlQyLgQwPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseXlQyLgQwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseXlQyLgQwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseXlQyLgQwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseXlQyLgQwPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseXlQyLgQwPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseXlQyLgQwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseXlQyLgQwPcapsResponseSimple]
type pcapListResponseXlQyLgQwPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseXlQyLgQwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseXlQyLgQwPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1]
type pcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseXlQyLgQwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseXlQyLgQwPcapsResponseSimpleStatus string

const (
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusUnknown           PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "unknown"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusSuccess           PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "success"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusPending           PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "pending"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusRunning           PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "running"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusConversionPending PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusConversionRunning PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusComplete          PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "complete"
	PcapListResponseXlQyLgQwPcapsResponseSimpleStatusFailed            PcapListResponseXlQyLgQwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseXlQyLgQwPcapsResponseSimpleSystem string

const (
	PcapListResponseXlQyLgQwPcapsResponseSimpleSystemMagicTransit PcapListResponseXlQyLgQwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseXlQyLgQwPcapsResponseSimpleType string

const (
	PcapListResponseXlQyLgQwPcapsResponseSimpleTypeSimple PcapListResponseXlQyLgQwPcapsResponseSimpleType = "simple"
	PcapListResponseXlQyLgQwPcapsResponseSimpleTypeFull   PcapListResponseXlQyLgQwPcapsResponseSimpleType = "full"
)

type PcapListResponseXlQyLgQwPcapsResponseFull struct {
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
	FilterV1 PcapListResponseXlQyLgQwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseXlQyLgQwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseXlQyLgQwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseXlQyLgQwPcapsResponseFullType `json:"type"`
	JSON pcapListResponseXlQyLgQwPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseXlQyLgQwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseXlQyLgQwPcapsResponseFull]
type pcapListResponseXlQyLgQwPcapsResponseFullJSON struct {
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

func (r *PcapListResponseXlQyLgQwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseXlQyLgQwPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseXlQyLgQwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseXlQyLgQwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseXlQyLgQwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseXlQyLgQwPcapsResponseFullFilterV1]
type pcapListResponseXlQyLgQwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseXlQyLgQwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseXlQyLgQwPcapsResponseFullStatus string

const (
	PcapListResponseXlQyLgQwPcapsResponseFullStatusUnknown           PcapListResponseXlQyLgQwPcapsResponseFullStatus = "unknown"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusSuccess           PcapListResponseXlQyLgQwPcapsResponseFullStatus = "success"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusPending           PcapListResponseXlQyLgQwPcapsResponseFullStatus = "pending"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusRunning           PcapListResponseXlQyLgQwPcapsResponseFullStatus = "running"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusConversionPending PcapListResponseXlQyLgQwPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusConversionRunning PcapListResponseXlQyLgQwPcapsResponseFullStatus = "conversion_running"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusComplete          PcapListResponseXlQyLgQwPcapsResponseFullStatus = "complete"
	PcapListResponseXlQyLgQwPcapsResponseFullStatusFailed            PcapListResponseXlQyLgQwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseXlQyLgQwPcapsResponseFullSystem string

const (
	PcapListResponseXlQyLgQwPcapsResponseFullSystemMagicTransit PcapListResponseXlQyLgQwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseXlQyLgQwPcapsResponseFullType string

const (
	PcapListResponseXlQyLgQwPcapsResponseFullTypeSimple PcapListResponseXlQyLgQwPcapsResponseFullType = "simple"
	PcapListResponseXlQyLgQwPcapsResponseFullTypeFull   PcapListResponseXlQyLgQwPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseXlQyLgQwPcapsResponseSimple] or
// [PcapGetResponseXlQyLgQwPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseXlQyLgQwPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseXlQyLgQwPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseXlQyLgQwPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseXlQyLgQwPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseXlQyLgQwPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseXlQyLgQwPcapsResponseSimple]
type pcapGetResponseXlQyLgQwPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseXlQyLgQwPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseXlQyLgQwPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1]
type pcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseXlQyLgQwPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus string

const (
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusUnknown           PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusSuccess           PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "success"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusPending           PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "pending"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusRunning           PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "running"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusConversionPending PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusConversionRunning PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusComplete          PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "complete"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleStatusFailed            PcapGetResponseXlQyLgQwPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseXlQyLgQwPcapsResponseSimpleSystem string

const (
	PcapGetResponseXlQyLgQwPcapsResponseSimpleSystemMagicTransit PcapGetResponseXlQyLgQwPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseXlQyLgQwPcapsResponseSimpleType string

const (
	PcapGetResponseXlQyLgQwPcapsResponseSimpleTypeSimple PcapGetResponseXlQyLgQwPcapsResponseSimpleType = "simple"
	PcapGetResponseXlQyLgQwPcapsResponseSimpleTypeFull   PcapGetResponseXlQyLgQwPcapsResponseSimpleType = "full"
)

type PcapGetResponseXlQyLgQwPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseXlQyLgQwPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseXlQyLgQwPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseXlQyLgQwPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseXlQyLgQwPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseXlQyLgQwPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseXlQyLgQwPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseXlQyLgQwPcapsResponseFull]
type pcapGetResponseXlQyLgQwPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseXlQyLgQwPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseXlQyLgQwPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseXlQyLgQwPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseXlQyLgQwPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseXlQyLgQwPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseXlQyLgQwPcapsResponseFullFilterV1]
type pcapGetResponseXlQyLgQwPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseXlQyLgQwPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseXlQyLgQwPcapsResponseFullStatus string

const (
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusUnknown           PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "unknown"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusSuccess           PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "success"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusPending           PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "pending"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusRunning           PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "running"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusConversionPending PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusConversionRunning PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusComplete          PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "complete"
	PcapGetResponseXlQyLgQwPcapsResponseFullStatusFailed            PcapGetResponseXlQyLgQwPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseXlQyLgQwPcapsResponseFullSystem string

const (
	PcapGetResponseXlQyLgQwPcapsResponseFullSystemMagicTransit PcapGetResponseXlQyLgQwPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseXlQyLgQwPcapsResponseFullType string

const (
	PcapGetResponseXlQyLgQwPcapsResponseFullTypeSimple PcapGetResponseXlQyLgQwPcapsResponseFullType = "simple"
	PcapGetResponseXlQyLgQwPcapsResponseFullTypeFull   PcapGetResponseXlQyLgQwPcapsResponseFullType = "full"
)

// This interface is a union satisfied by one of the following:
// [PcapNewParamsXlQyLgQwPcapsRequestSimple],
// [PcapNewParamsXlQyLgQwPcapsRequestFull].
type PcapNewParams interface {
	ImplementsPcapNewParams()
}

type PcapNewParamsXlQyLgQwPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsXlQyLgQwPcapsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsXlQyLgQwPcapsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsXlQyLgQwPcapsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PcapNewParamsXlQyLgQwPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsXlQyLgQwPcapsRequestSimple) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsXlQyLgQwPcapsRequestSimpleSystem string

const (
	PcapNewParamsXlQyLgQwPcapsRequestSimpleSystemMagicTransit PcapNewParamsXlQyLgQwPcapsRequestSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsXlQyLgQwPcapsRequestSimpleType string

const (
	PcapNewParamsXlQyLgQwPcapsRequestSimpleTypeSimple PcapNewParamsXlQyLgQwPcapsRequestSimpleType = "simple"
	PcapNewParamsXlQyLgQwPcapsRequestSimpleTypeFull   PcapNewParamsXlQyLgQwPcapsRequestSimpleType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsXlQyLgQwPcapsRequestSimpleFilterV1 struct {
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

func (r PcapNewParamsXlQyLgQwPcapsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapNewParamsXlQyLgQwPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PcapNewParamsXlQyLgQwPcapsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PcapNewParamsXlQyLgQwPcapsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PcapNewParamsXlQyLgQwPcapsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PcapNewParamsXlQyLgQwPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PcapNewParamsXlQyLgQwPcapsRequestFull) ImplementsPcapNewParams() {

}

// The system used to collect packet captures.
type PcapNewParamsXlQyLgQwPcapsRequestFullSystem string

const (
	PcapNewParamsXlQyLgQwPcapsRequestFullSystemMagicTransit PcapNewParamsXlQyLgQwPcapsRequestFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewParamsXlQyLgQwPcapsRequestFullType string

const (
	PcapNewParamsXlQyLgQwPcapsRequestFullTypeSimple PcapNewParamsXlQyLgQwPcapsRequestFullType = "simple"
	PcapNewParamsXlQyLgQwPcapsRequestFullTypeFull   PcapNewParamsXlQyLgQwPcapsRequestFullType = "full"
)

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewParamsXlQyLgQwPcapsRequestFullFilterV1 struct {
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

func (r PcapNewParamsXlQyLgQwPcapsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
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
