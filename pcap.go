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

// PCAPService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPCAPService] method instead.
type PCAPService struct {
	Options    []option.RequestOption
	Ownerships *PCAPOwnershipService
	Downloads  *PCAPDownloadService
}

// NewPCAPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPCAPService(opts ...option.RequestOption) (r *PCAPService) {
	r = &PCAPService{}
	r.Options = opts
	r.Ownerships = NewPCAPOwnershipService(opts...)
	r.Downloads = NewPCAPDownloadService(opts...)
	return
}

// Create new PCAP request for account.
func (r *PCAPService) New(ctx context.Context, accountIdentifier string, body PCAPNewParams, opts ...option.RequestOption) (res *PCAPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all packet capture requests for an account.
func (r *PCAPService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PCAPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a PCAP request by id.
func (r *PCAPService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *PCAPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PCAPNewResponse5pGRyj1RPCAPsResponseSimple] or
// [PCAPNewResponse5pGRyj1RPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponse5pGRyj1RPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponse5pGRyj1RPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse5pGRyj1RPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse5pGRyj1RPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponse5pGRyj1RpcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse5pGRyj1RpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponse5pGRyj1RPCAPsResponseSimple]
type pcapNewResponse5pGRyj1RpcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponse5pGRyj1RPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse5pGRyj1RPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse5pGRyj1RPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponse5pGRyj1RPCAPsResponseSimpleFilterV1]
type pcapNewResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse5pGRyj1RPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusUnknown           PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusSuccess           PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "success"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusPending           PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusRunning           PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "running"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusConversionPending PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusConversionRunning PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusComplete          PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatusFailed            PCAPNewResponse5pGRyj1RPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse5pGRyj1RPCAPsResponseSimpleSystem string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleSystemMagicTransit PCAPNewResponse5pGRyj1RPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse5pGRyj1RPCAPsResponseSimpleType string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleTypeSimple PCAPNewResponse5pGRyj1RPCAPsResponseSimpleType = "simple"
	PCAPNewResponse5pGRyj1RPCAPsResponseSimpleTypeFull   PCAPNewResponse5pGRyj1RPCAPsResponseSimpleType = "full"
)

type PCAPNewResponse5pGRyj1RPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponse5pGRyj1RPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse5pGRyj1RPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse5pGRyj1RPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponse5pGRyj1RpcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponse5pGRyj1RpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponse5pGRyj1RPCAPsResponseFull]
type pcapNewResponse5pGRyj1RpcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponse5pGRyj1RPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse5pGRyj1RPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse5pGRyj1RPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse5pGRyj1RpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse5pGRyj1RpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponse5pGRyj1RPCAPsResponseFullFilterV1]
type pcapNewResponse5pGRyj1RpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse5pGRyj1RPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusUnknown           PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "unknown"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusSuccess           PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "success"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusPending           PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "pending"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusRunning           PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "running"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusConversionPending PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusConversionRunning PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusComplete          PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "complete"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullStatusFailed            PCAPNewResponse5pGRyj1RPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse5pGRyj1RPCAPsResponseFullSystem string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseFullSystemMagicTransit PCAPNewResponse5pGRyj1RPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse5pGRyj1RPCAPsResponseFullType string

const (
	PCAPNewResponse5pGRyj1RPCAPsResponseFullTypeSimple PCAPNewResponse5pGRyj1RPCAPsResponseFullType = "simple"
	PCAPNewResponse5pGRyj1RPCAPsResponseFullTypeFull   PCAPNewResponse5pGRyj1RPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponse5pGRyj1RPCAPsResponseSimple] or
// [PCAPListResponse5pGRyj1RPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponse5pGRyj1RPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponse5pGRyj1RPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse5pGRyj1RPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse5pGRyj1RPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponse5pGRyj1RpcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponse5pGRyj1RpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponse5pGRyj1RPCAPsResponseSimple]
type pcapListResponse5pGRyj1RpcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponse5pGRyj1RPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse5pGRyj1RPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse5pGRyj1RPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponse5pGRyj1RPCAPsResponseSimpleFilterV1]
type pcapListResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse5pGRyj1RPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusUnknown           PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusSuccess           PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "success"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusPending           PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "pending"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusRunning           PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "running"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusConversionPending PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusConversionRunning PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusComplete          PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "complete"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatusFailed            PCAPListResponse5pGRyj1RPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse5pGRyj1RPCAPsResponseSimpleSystem string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleSystemMagicTransit PCAPListResponse5pGRyj1RPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse5pGRyj1RPCAPsResponseSimpleType string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleTypeSimple PCAPListResponse5pGRyj1RPCAPsResponseSimpleType = "simple"
	PCAPListResponse5pGRyj1RPCAPsResponseSimpleTypeFull   PCAPListResponse5pGRyj1RPCAPsResponseSimpleType = "full"
)

type PCAPListResponse5pGRyj1RPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponse5pGRyj1RPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse5pGRyj1RPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse5pGRyj1RPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse5pGRyj1RPCAPsResponseFullType `json:"type"`
	JSON pcapListResponse5pGRyj1RpcaPsResponseFullJSON `json:"-"`
}

// pcapListResponse5pGRyj1RpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponse5pGRyj1RPCAPsResponseFull]
type pcapListResponse5pGRyj1RpcaPsResponseFullJSON struct {
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

func (r *PCAPListResponse5pGRyj1RPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse5pGRyj1RPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse5pGRyj1RPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse5pGRyj1RpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse5pGRyj1RpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponse5pGRyj1RPCAPsResponseFullFilterV1]
type pcapListResponse5pGRyj1RpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse5pGRyj1RPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse5pGRyj1RPCAPsResponseFullStatus string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusUnknown           PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "unknown"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusSuccess           PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "success"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusPending           PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "pending"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusRunning           PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "running"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusConversionPending PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusConversionRunning PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusComplete          PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "complete"
	PCAPListResponse5pGRyj1RPCAPsResponseFullStatusFailed            PCAPListResponse5pGRyj1RPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse5pGRyj1RPCAPsResponseFullSystem string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseFullSystemMagicTransit PCAPListResponse5pGRyj1RPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse5pGRyj1RPCAPsResponseFullType string

const (
	PCAPListResponse5pGRyj1RPCAPsResponseFullTypeSimple PCAPListResponse5pGRyj1RPCAPsResponseFullType = "simple"
	PCAPListResponse5pGRyj1RPCAPsResponseFullTypeFull   PCAPListResponse5pGRyj1RPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponse5pGRyj1RPCAPsResponseSimple] or
// [PCAPGetResponse5pGRyj1RPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponse5pGRyj1RPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponse5pGRyj1RPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse5pGRyj1RPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse5pGRyj1RPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponse5pGRyj1RpcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse5pGRyj1RpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponse5pGRyj1RPCAPsResponseSimple]
type pcapGetResponse5pGRyj1RpcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponse5pGRyj1RPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse5pGRyj1RPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse5pGRyj1RPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponse5pGRyj1RPCAPsResponseSimpleFilterV1]
type pcapGetResponse5pGRyj1RpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse5pGRyj1RPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusUnknown           PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusSuccess           PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "success"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusPending           PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusRunning           PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "running"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusConversionPending PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusConversionRunning PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusComplete          PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatusFailed            PCAPGetResponse5pGRyj1RPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse5pGRyj1RPCAPsResponseSimpleSystem string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleSystemMagicTransit PCAPGetResponse5pGRyj1RPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse5pGRyj1RPCAPsResponseSimpleType string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleTypeSimple PCAPGetResponse5pGRyj1RPCAPsResponseSimpleType = "simple"
	PCAPGetResponse5pGRyj1RPCAPsResponseSimpleTypeFull   PCAPGetResponse5pGRyj1RPCAPsResponseSimpleType = "full"
)

type PCAPGetResponse5pGRyj1RPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponse5pGRyj1RPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse5pGRyj1RPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse5pGRyj1RPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponse5pGRyj1RpcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponse5pGRyj1RpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponse5pGRyj1RPCAPsResponseFull]
type pcapGetResponse5pGRyj1RpcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponse5pGRyj1RPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse5pGRyj1RPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse5pGRyj1RPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse5pGRyj1RpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse5pGRyj1RpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponse5pGRyj1RPCAPsResponseFullFilterV1]
type pcapGetResponse5pGRyj1RpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse5pGRyj1RPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusUnknown           PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "unknown"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusSuccess           PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "success"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusPending           PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "pending"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusRunning           PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "running"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusConversionPending PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusConversionRunning PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusComplete          PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "complete"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullStatusFailed            PCAPGetResponse5pGRyj1RPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse5pGRyj1RPCAPsResponseFullSystem string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseFullSystemMagicTransit PCAPGetResponse5pGRyj1RPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse5pGRyj1RPCAPsResponseFullType string

const (
	PCAPGetResponse5pGRyj1RPCAPsResponseFullTypeSimple PCAPGetResponse5pGRyj1RPCAPsResponseFullType = "simple"
	PCAPGetResponse5pGRyj1RPCAPsResponseFullTypeFull   PCAPGetResponse5pGRyj1RPCAPsResponseFullType = "full"
)

type PCAPNewParams struct {
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string]                `json:"destination_conf"`
	FilterV1        param.Field[PCAPNewParamsFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PCAPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The system used to collect packet captures.
type PCAPNewParamsSystem string

const (
	PCAPNewParamsSystemMagicTransit PCAPNewParamsSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsType string

const (
	PCAPNewParamsTypeSimple PCAPNewParamsType = "simple"
	PCAPNewParamsTypeFull   PCAPNewParamsType = "full"
)

type PCAPNewParamsFilterV1 struct {
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

func (r PCAPNewParamsFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPNewResponseEnvelope struct {
	Errors   []PCAPNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PCAPNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PCAPNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapNewResponseEnvelopeJSON    `json:"-"`
}

// pcapNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPNewResponseEnvelope]
type pcapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pcapNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PCAPNewResponseEnvelopeErrors]
type pcapNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pcapNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PCAPNewResponseEnvelopeMessages]
type pcapNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPNewResponseEnvelopeSuccess bool

const (
	PCAPNewResponseEnvelopeSuccessTrue PCAPNewResponseEnvelopeSuccess = true
)

type PCAPListResponseEnvelope struct {
	Errors   []PCAPListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PCAPListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PCAPListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PCAPListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapListResponseEnvelopeJSON       `json:"-"`
}

// pcapListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPListResponseEnvelope]
type pcapListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    pcapListResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PCAPListResponseEnvelopeErrors]
type pcapListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    pcapListResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PCAPListResponseEnvelopeMessages]
type pcapListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPListResponseEnvelopeSuccess bool

const (
	PCAPListResponseEnvelopeSuccessTrue PCAPListResponseEnvelopeSuccess = true
)

type PCAPListResponseEnvelopeResultInfo struct {
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
// [PCAPListResponseEnvelopeResultInfo]
type pcapListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPGetResponseEnvelope struct {
	Errors   []PCAPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PCAPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PCAPGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PCAPGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapGetResponseEnvelopeJSON    `json:"-"`
}

// pcapGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PCAPGetResponseEnvelope]
type pcapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pcapGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PCAPGetResponseEnvelopeErrors]
type pcapGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PCAPGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pcapGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PCAPGetResponseEnvelopeMessages]
type pcapGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PCAPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PCAPGetResponseEnvelopeSuccess bool

const (
	PCAPGetResponseEnvelopeSuccessTrue PCAPGetResponseEnvelopeSuccess = true
)
