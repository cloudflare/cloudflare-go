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

// Union satisfied by [PcapNewResponseBvwDlDzxPcapsResponseSimple] or
// [PcapNewResponseBvwDlDzxPcapsResponseFull].
type PcapNewResponse interface {
	implementsPcapNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapNewResponse)(nil)).Elem(), "")
}

type PcapNewResponseBvwDlDzxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseBvwDlDzxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseBvwDlDzxPcapsResponseSimpleType `json:"type"`
	JSON pcapNewResponseBvwDlDzxPcapsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseBvwDlDzxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapNewResponseBvwDlDzxPcapsResponseSimple]
type pcapNewResponseBvwDlDzxPcapsResponseSimpleJSON struct {
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

func (r *PcapNewResponseBvwDlDzxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseBvwDlDzxPcapsResponseSimple) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1]
type pcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseBvwDlDzxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus string

const (
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusUnknown           PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "unknown"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusSuccess           PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "success"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusPending           PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "pending"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusRunning           PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "running"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusConversionPending PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_pending"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusConversionRunning PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_running"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusComplete          PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "complete"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleStatusFailed            PcapNewResponseBvwDlDzxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseBvwDlDzxPcapsResponseSimpleSystem string

const (
	PcapNewResponseBvwDlDzxPcapsResponseSimpleSystemMagicTransit PcapNewResponseBvwDlDzxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseBvwDlDzxPcapsResponseSimpleType string

const (
	PcapNewResponseBvwDlDzxPcapsResponseSimpleTypeSimple PcapNewResponseBvwDlDzxPcapsResponseSimpleType = "simple"
	PcapNewResponseBvwDlDzxPcapsResponseSimpleTypeFull   PcapNewResponseBvwDlDzxPcapsResponseSimpleType = "full"
)

type PcapNewResponseBvwDlDzxPcapsResponseFull struct {
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
	FilterV1 PcapNewResponseBvwDlDzxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapNewResponseBvwDlDzxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapNewResponseBvwDlDzxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapNewResponseBvwDlDzxPcapsResponseFullType `json:"type"`
	JSON pcapNewResponseBvwDlDzxPcapsResponseFullJSON `json:"-"`
}

// pcapNewResponseBvwDlDzxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapNewResponseBvwDlDzxPcapsResponseFull]
type pcapNewResponseBvwDlDzxPcapsResponseFullJSON struct {
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

func (r *PcapNewResponseBvwDlDzxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapNewResponseBvwDlDzxPcapsResponseFull) implementsPcapNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapNewResponseBvwDlDzxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseBvwDlDzxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseBvwDlDzxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapNewResponseBvwDlDzxPcapsResponseFullFilterV1]
type pcapNewResponseBvwDlDzxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapNewResponseBvwDlDzxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapNewResponseBvwDlDzxPcapsResponseFullStatus string

const (
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusUnknown           PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "unknown"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusSuccess           PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "success"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusPending           PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "pending"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusRunning           PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "running"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusConversionPending PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "conversion_pending"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusConversionRunning PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "conversion_running"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusComplete          PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "complete"
	PcapNewResponseBvwDlDzxPcapsResponseFullStatusFailed            PcapNewResponseBvwDlDzxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapNewResponseBvwDlDzxPcapsResponseFullSystem string

const (
	PcapNewResponseBvwDlDzxPcapsResponseFullSystemMagicTransit PcapNewResponseBvwDlDzxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapNewResponseBvwDlDzxPcapsResponseFullType string

const (
	PcapNewResponseBvwDlDzxPcapsResponseFullTypeSimple PcapNewResponseBvwDlDzxPcapsResponseFullType = "simple"
	PcapNewResponseBvwDlDzxPcapsResponseFullTypeFull   PcapNewResponseBvwDlDzxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapListResponseBvwDlDzxPcapsResponseSimple] or
// [PcapListResponseBvwDlDzxPcapsResponseFull].
type PcapListResponse interface {
	implementsPcapListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapListResponse)(nil)).Elem(), "")
}

type PcapListResponseBvwDlDzxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseBvwDlDzxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseBvwDlDzxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseBvwDlDzxPcapsResponseSimpleType `json:"type"`
	JSON pcapListResponseBvwDlDzxPcapsResponseSimpleJSON `json:"-"`
}

// pcapListResponseBvwDlDzxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapListResponseBvwDlDzxPcapsResponseSimple]
type pcapListResponseBvwDlDzxPcapsResponseSimpleJSON struct {
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

func (r *PcapListResponseBvwDlDzxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseBvwDlDzxPcapsResponseSimple) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1]
type pcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseBvwDlDzxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseBvwDlDzxPcapsResponseSimpleStatus string

const (
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusUnknown           PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "unknown"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusSuccess           PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "success"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusPending           PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "pending"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusRunning           PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "running"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusConversionPending PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_pending"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusConversionRunning PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_running"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusComplete          PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "complete"
	PcapListResponseBvwDlDzxPcapsResponseSimpleStatusFailed            PcapListResponseBvwDlDzxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseBvwDlDzxPcapsResponseSimpleSystem string

const (
	PcapListResponseBvwDlDzxPcapsResponseSimpleSystemMagicTransit PcapListResponseBvwDlDzxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseBvwDlDzxPcapsResponseSimpleType string

const (
	PcapListResponseBvwDlDzxPcapsResponseSimpleTypeSimple PcapListResponseBvwDlDzxPcapsResponseSimpleType = "simple"
	PcapListResponseBvwDlDzxPcapsResponseSimpleTypeFull   PcapListResponseBvwDlDzxPcapsResponseSimpleType = "full"
)

type PcapListResponseBvwDlDzxPcapsResponseFull struct {
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
	FilterV1 PcapListResponseBvwDlDzxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapListResponseBvwDlDzxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapListResponseBvwDlDzxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapListResponseBvwDlDzxPcapsResponseFullType `json:"type"`
	JSON pcapListResponseBvwDlDzxPcapsResponseFullJSON `json:"-"`
}

// pcapListResponseBvwDlDzxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapListResponseBvwDlDzxPcapsResponseFull]
type pcapListResponseBvwDlDzxPcapsResponseFullJSON struct {
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

func (r *PcapListResponseBvwDlDzxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapListResponseBvwDlDzxPcapsResponseFull) implementsPcapListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapListResponseBvwDlDzxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseBvwDlDzxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseBvwDlDzxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapListResponseBvwDlDzxPcapsResponseFullFilterV1]
type pcapListResponseBvwDlDzxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapListResponseBvwDlDzxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapListResponseBvwDlDzxPcapsResponseFullStatus string

const (
	PcapListResponseBvwDlDzxPcapsResponseFullStatusUnknown           PcapListResponseBvwDlDzxPcapsResponseFullStatus = "unknown"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusSuccess           PcapListResponseBvwDlDzxPcapsResponseFullStatus = "success"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusPending           PcapListResponseBvwDlDzxPcapsResponseFullStatus = "pending"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusRunning           PcapListResponseBvwDlDzxPcapsResponseFullStatus = "running"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusConversionPending PcapListResponseBvwDlDzxPcapsResponseFullStatus = "conversion_pending"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusConversionRunning PcapListResponseBvwDlDzxPcapsResponseFullStatus = "conversion_running"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusComplete          PcapListResponseBvwDlDzxPcapsResponseFullStatus = "complete"
	PcapListResponseBvwDlDzxPcapsResponseFullStatusFailed            PcapListResponseBvwDlDzxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapListResponseBvwDlDzxPcapsResponseFullSystem string

const (
	PcapListResponseBvwDlDzxPcapsResponseFullSystemMagicTransit PcapListResponseBvwDlDzxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapListResponseBvwDlDzxPcapsResponseFullType string

const (
	PcapListResponseBvwDlDzxPcapsResponseFullTypeSimple PcapListResponseBvwDlDzxPcapsResponseFullType = "simple"
	PcapListResponseBvwDlDzxPcapsResponseFullTypeFull   PcapListResponseBvwDlDzxPcapsResponseFullType = "full"
)

// Union satisfied by [PcapGetResponseBvwDlDzxPcapsResponseSimple] or
// [PcapGetResponseBvwDlDzxPcapsResponseFull].
type PcapGetResponse interface {
	implementsPcapGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PcapGetResponse)(nil)).Elem(), "")
}

type PcapGetResponseBvwDlDzxPcapsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseBvwDlDzxPcapsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseBvwDlDzxPcapsResponseSimpleType `json:"type"`
	JSON pcapGetResponseBvwDlDzxPcapsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseBvwDlDzxPcapsResponseSimpleJSON contains the JSON metadata for
// the struct [PcapGetResponseBvwDlDzxPcapsResponseSimple]
type pcapGetResponseBvwDlDzxPcapsResponseSimpleJSON struct {
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

func (r *PcapGetResponseBvwDlDzxPcapsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseBvwDlDzxPcapsResponseSimple) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1]
type pcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseBvwDlDzxPcapsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus string

const (
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusUnknown           PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "unknown"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusSuccess           PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "success"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusPending           PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "pending"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusRunning           PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "running"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusConversionPending PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_pending"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusConversionRunning PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "conversion_running"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusComplete          PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "complete"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleStatusFailed            PcapGetResponseBvwDlDzxPcapsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseBvwDlDzxPcapsResponseSimpleSystem string

const (
	PcapGetResponseBvwDlDzxPcapsResponseSimpleSystemMagicTransit PcapGetResponseBvwDlDzxPcapsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseBvwDlDzxPcapsResponseSimpleType string

const (
	PcapGetResponseBvwDlDzxPcapsResponseSimpleTypeSimple PcapGetResponseBvwDlDzxPcapsResponseSimpleType = "simple"
	PcapGetResponseBvwDlDzxPcapsResponseSimpleTypeFull   PcapGetResponseBvwDlDzxPcapsResponseSimpleType = "full"
)

type PcapGetResponseBvwDlDzxPcapsResponseFull struct {
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
	FilterV1 PcapGetResponseBvwDlDzxPcapsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PcapGetResponseBvwDlDzxPcapsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PcapGetResponseBvwDlDzxPcapsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PcapGetResponseBvwDlDzxPcapsResponseFullType `json:"type"`
	JSON pcapGetResponseBvwDlDzxPcapsResponseFullJSON `json:"-"`
}

// pcapGetResponseBvwDlDzxPcapsResponseFullJSON contains the JSON metadata for the
// struct [PcapGetResponseBvwDlDzxPcapsResponseFull]
type pcapGetResponseBvwDlDzxPcapsResponseFullJSON struct {
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

func (r *PcapGetResponseBvwDlDzxPcapsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PcapGetResponseBvwDlDzxPcapsResponseFull) implementsPcapGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PcapGetResponseBvwDlDzxPcapsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseBvwDlDzxPcapsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseBvwDlDzxPcapsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PcapGetResponseBvwDlDzxPcapsResponseFullFilterV1]
type pcapGetResponseBvwDlDzxPcapsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PcapGetResponseBvwDlDzxPcapsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PcapGetResponseBvwDlDzxPcapsResponseFullStatus string

const (
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusUnknown           PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "unknown"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusSuccess           PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "success"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusPending           PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "pending"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusRunning           PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "running"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusConversionPending PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "conversion_pending"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusConversionRunning PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "conversion_running"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusComplete          PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "complete"
	PcapGetResponseBvwDlDzxPcapsResponseFullStatusFailed            PcapGetResponseBvwDlDzxPcapsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PcapGetResponseBvwDlDzxPcapsResponseFullSystem string

const (
	PcapGetResponseBvwDlDzxPcapsResponseFullSystemMagicTransit PcapGetResponseBvwDlDzxPcapsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PcapGetResponseBvwDlDzxPcapsResponseFullType string

const (
	PcapGetResponseBvwDlDzxPcapsResponseFullTypeSimple PcapGetResponseBvwDlDzxPcapsResponseFullType = "simple"
	PcapGetResponseBvwDlDzxPcapsResponseFullTypeFull   PcapGetResponseBvwDlDzxPcapsResponseFullType = "full"
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
