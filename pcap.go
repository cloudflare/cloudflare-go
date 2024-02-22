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

// Union satisfied by [PCAPNewResponse9Iuma6lRPCAPsResponseSimple] or
// [PCAPNewResponse9Iuma6lRPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponse9Iuma6lRPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponse9Iuma6lRPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse9Iuma6lRPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse9Iuma6lRPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponse9Iuma6lRpcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse9Iuma6lRpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponse9Iuma6lRPCAPsResponseSimple]
type pcapNewResponse9Iuma6lRpcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponse9Iuma6lRPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse9Iuma6lRPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse9Iuma6lRPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponse9Iuma6lRPCAPsResponseSimpleFilterV1]
type pcapNewResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse9Iuma6lRPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusUnknown           PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusSuccess           PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "success"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusPending           PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusRunning           PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "running"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusConversionPending PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusConversionRunning PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusComplete          PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatusFailed            PCAPNewResponse9Iuma6lRPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse9Iuma6lRPCAPsResponseSimpleSystem string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleSystemMagicTransit PCAPNewResponse9Iuma6lRPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse9Iuma6lRPCAPsResponseSimpleType string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleTypeSimple PCAPNewResponse9Iuma6lRPCAPsResponseSimpleType = "simple"
	PCAPNewResponse9Iuma6lRPCAPsResponseSimpleTypeFull   PCAPNewResponse9Iuma6lRPCAPsResponseSimpleType = "full"
)

type PCAPNewResponse9Iuma6lRPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponse9Iuma6lRPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse9Iuma6lRPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse9Iuma6lRPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponse9Iuma6lRpcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponse9Iuma6lRpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponse9Iuma6lRPCAPsResponseFull]
type pcapNewResponse9Iuma6lRpcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponse9Iuma6lRPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse9Iuma6lRPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse9Iuma6lRPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse9Iuma6lRpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse9Iuma6lRpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponse9Iuma6lRPCAPsResponseFullFilterV1]
type pcapNewResponse9Iuma6lRpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse9Iuma6lRPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusUnknown           PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "unknown"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusSuccess           PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "success"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusPending           PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "pending"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusRunning           PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "running"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusConversionPending PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusConversionRunning PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusComplete          PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "complete"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullStatusFailed            PCAPNewResponse9Iuma6lRPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse9Iuma6lRPCAPsResponseFullSystem string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseFullSystemMagicTransit PCAPNewResponse9Iuma6lRPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse9Iuma6lRPCAPsResponseFullType string

const (
	PCAPNewResponse9Iuma6lRPCAPsResponseFullTypeSimple PCAPNewResponse9Iuma6lRPCAPsResponseFullType = "simple"
	PCAPNewResponse9Iuma6lRPCAPsResponseFullTypeFull   PCAPNewResponse9Iuma6lRPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponse9Iuma6lRPCAPsResponseSimple] or
// [PCAPListResponse9Iuma6lRPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponse9Iuma6lRPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponse9Iuma6lRPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse9Iuma6lRPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse9Iuma6lRPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponse9Iuma6lRpcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponse9Iuma6lRpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponse9Iuma6lRPCAPsResponseSimple]
type pcapListResponse9Iuma6lRpcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponse9Iuma6lRPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse9Iuma6lRPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse9Iuma6lRPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponse9Iuma6lRPCAPsResponseSimpleFilterV1]
type pcapListResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse9Iuma6lRPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusUnknown           PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusSuccess           PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "success"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusPending           PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "pending"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusRunning           PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "running"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusConversionPending PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusConversionRunning PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusComplete          PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "complete"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatusFailed            PCAPListResponse9Iuma6lRPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse9Iuma6lRPCAPsResponseSimpleSystem string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleSystemMagicTransit PCAPListResponse9Iuma6lRPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse9Iuma6lRPCAPsResponseSimpleType string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleTypeSimple PCAPListResponse9Iuma6lRPCAPsResponseSimpleType = "simple"
	PCAPListResponse9Iuma6lRPCAPsResponseSimpleTypeFull   PCAPListResponse9Iuma6lRPCAPsResponseSimpleType = "full"
)

type PCAPListResponse9Iuma6lRPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponse9Iuma6lRPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse9Iuma6lRPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse9Iuma6lRPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse9Iuma6lRPCAPsResponseFullType `json:"type"`
	JSON pcapListResponse9Iuma6lRpcaPsResponseFullJSON `json:"-"`
}

// pcapListResponse9Iuma6lRpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponse9Iuma6lRPCAPsResponseFull]
type pcapListResponse9Iuma6lRpcaPsResponseFullJSON struct {
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

func (r *PCAPListResponse9Iuma6lRPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse9Iuma6lRPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse9Iuma6lRPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse9Iuma6lRpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse9Iuma6lRpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponse9Iuma6lRPCAPsResponseFullFilterV1]
type pcapListResponse9Iuma6lRpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse9Iuma6lRPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse9Iuma6lRPCAPsResponseFullStatus string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusUnknown           PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "unknown"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusSuccess           PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "success"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusPending           PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "pending"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusRunning           PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "running"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusConversionPending PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusConversionRunning PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusComplete          PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "complete"
	PCAPListResponse9Iuma6lRPCAPsResponseFullStatusFailed            PCAPListResponse9Iuma6lRPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse9Iuma6lRPCAPsResponseFullSystem string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseFullSystemMagicTransit PCAPListResponse9Iuma6lRPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse9Iuma6lRPCAPsResponseFullType string

const (
	PCAPListResponse9Iuma6lRPCAPsResponseFullTypeSimple PCAPListResponse9Iuma6lRPCAPsResponseFullType = "simple"
	PCAPListResponse9Iuma6lRPCAPsResponseFullTypeFull   PCAPListResponse9Iuma6lRPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponse9Iuma6lRPCAPsResponseSimple] or
// [PCAPGetResponse9Iuma6lRPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponse9Iuma6lRPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponse9Iuma6lRPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse9Iuma6lRPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse9Iuma6lRPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponse9Iuma6lRpcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse9Iuma6lRpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponse9Iuma6lRPCAPsResponseSimple]
type pcapGetResponse9Iuma6lRpcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponse9Iuma6lRPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse9Iuma6lRPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse9Iuma6lRPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponse9Iuma6lRPCAPsResponseSimpleFilterV1]
type pcapGetResponse9Iuma6lRpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse9Iuma6lRPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusUnknown           PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusSuccess           PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "success"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusPending           PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusRunning           PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "running"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusConversionPending PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusConversionRunning PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusComplete          PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatusFailed            PCAPGetResponse9Iuma6lRPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse9Iuma6lRPCAPsResponseSimpleSystem string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleSystemMagicTransit PCAPGetResponse9Iuma6lRPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse9Iuma6lRPCAPsResponseSimpleType string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleTypeSimple PCAPGetResponse9Iuma6lRPCAPsResponseSimpleType = "simple"
	PCAPGetResponse9Iuma6lRPCAPsResponseSimpleTypeFull   PCAPGetResponse9Iuma6lRPCAPsResponseSimpleType = "full"
)

type PCAPGetResponse9Iuma6lRPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponse9Iuma6lRPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse9Iuma6lRPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse9Iuma6lRPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponse9Iuma6lRpcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponse9Iuma6lRpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponse9Iuma6lRPCAPsResponseFull]
type pcapGetResponse9Iuma6lRpcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponse9Iuma6lRPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse9Iuma6lRPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse9Iuma6lRPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse9Iuma6lRpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse9Iuma6lRpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponse9Iuma6lRPCAPsResponseFullFilterV1]
type pcapGetResponse9Iuma6lRpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse9Iuma6lRPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusUnknown           PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "unknown"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusSuccess           PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "success"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusPending           PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "pending"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusRunning           PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "running"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusConversionPending PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusConversionRunning PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusComplete          PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "complete"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullStatusFailed            PCAPGetResponse9Iuma6lRPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse9Iuma6lRPCAPsResponseFullSystem string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseFullSystemMagicTransit PCAPGetResponse9Iuma6lRPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse9Iuma6lRPCAPsResponseFullType string

const (
	PCAPGetResponse9Iuma6lRPCAPsResponseFullTypeSimple PCAPGetResponse9Iuma6lRPCAPsResponseFullType = "simple"
	PCAPGetResponse9Iuma6lRPCAPsResponseFullTypeFull   PCAPGetResponse9Iuma6lRPCAPsResponseFullType = "full"
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
