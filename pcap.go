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

// Union satisfied by [PCAPNewResponseBeWHv1YePCAPsResponseSimple] or
// [PCAPNewResponseBeWHv1YePCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseBeWHv1YePCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseBeWHv1YePCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseBeWHv1YePCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseBeWHv1YePCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseBeWHv1YePCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseBeWHv1YePCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseBeWHv1YePCAPsResponseSimple]
type pcapNewResponseBeWHv1YePCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseBeWHv1YePCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseBeWHv1YePCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseBeWHv1YePCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseBeWHv1YePCAPsResponseSimpleFilterV1]
type pcapNewResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseBeWHv1YePCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusUnknown           PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusSuccess           PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "success"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusPending           PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusRunning           PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "running"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusConversionPending PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusConversionRunning PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusComplete          PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatusFailed            PCAPNewResponseBeWHv1YePCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseBeWHv1YePCAPsResponseSimpleSystem string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleSystemMagicTransit PCAPNewResponseBeWHv1YePCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseBeWHv1YePCAPsResponseSimpleType string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleTypeSimple PCAPNewResponseBeWHv1YePCAPsResponseSimpleType = "simple"
	PCAPNewResponseBeWHv1YePCAPsResponseSimpleTypeFull   PCAPNewResponseBeWHv1YePCAPsResponseSimpleType = "full"
)

type PCAPNewResponseBeWHv1YePCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseBeWHv1YePCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseBeWHv1YePCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseBeWHv1YePCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseBeWHv1YePCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseBeWHv1YePCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseBeWHv1YePCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseBeWHv1YePCAPsResponseFull]
type pcapNewResponseBeWHv1YePCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseBeWHv1YePCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseBeWHv1YePCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseBeWHv1YePCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseBeWHv1YePCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseBeWHv1YePCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseBeWHv1YePCAPsResponseFullFilterV1]
type pcapNewResponseBeWHv1YePCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseBeWHv1YePCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseBeWHv1YePCAPsResponseFullStatus string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusUnknown           PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "unknown"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusSuccess           PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "success"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusPending           PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "pending"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusRunning           PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "running"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusConversionPending PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusConversionRunning PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusComplete          PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "complete"
	PCAPNewResponseBeWHv1YePCAPsResponseFullStatusFailed            PCAPNewResponseBeWHv1YePCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseBeWHv1YePCAPsResponseFullSystem string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseFullSystemMagicTransit PCAPNewResponseBeWHv1YePCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseBeWHv1YePCAPsResponseFullType string

const (
	PCAPNewResponseBeWHv1YePCAPsResponseFullTypeSimple PCAPNewResponseBeWHv1YePCAPsResponseFullType = "simple"
	PCAPNewResponseBeWHv1YePCAPsResponseFullTypeFull   PCAPNewResponseBeWHv1YePCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseBeWHv1YePCAPsResponseSimple] or
// [PCAPListResponseBeWHv1YePCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseBeWHv1YePCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseBeWHv1YePCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseBeWHv1YePCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseBeWHv1YePCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseBeWHv1YePCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseBeWHv1YePCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseBeWHv1YePCAPsResponseSimple]
type pcapListResponseBeWHv1YePCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseBeWHv1YePCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseBeWHv1YePCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseBeWHv1YePCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseBeWHv1YePCAPsResponseSimpleFilterV1]
type pcapListResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseBeWHv1YePCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus string

const (
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusUnknown           PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusSuccess           PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "success"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusPending           PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "pending"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusRunning           PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "running"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusConversionPending PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusConversionRunning PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusComplete          PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "complete"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleStatusFailed            PCAPListResponseBeWHv1YePCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseBeWHv1YePCAPsResponseSimpleSystem string

const (
	PCAPListResponseBeWHv1YePCAPsResponseSimpleSystemMagicTransit PCAPListResponseBeWHv1YePCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseBeWHv1YePCAPsResponseSimpleType string

const (
	PCAPListResponseBeWHv1YePCAPsResponseSimpleTypeSimple PCAPListResponseBeWHv1YePCAPsResponseSimpleType = "simple"
	PCAPListResponseBeWHv1YePCAPsResponseSimpleTypeFull   PCAPListResponseBeWHv1YePCAPsResponseSimpleType = "full"
)

type PCAPListResponseBeWHv1YePCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseBeWHv1YePCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseBeWHv1YePCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseBeWHv1YePCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseBeWHv1YePCAPsResponseFullType `json:"type"`
	JSON pcapListResponseBeWHv1YePCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseBeWHv1YePCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseBeWHv1YePCAPsResponseFull]
type pcapListResponseBeWHv1YePCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseBeWHv1YePCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseBeWHv1YePCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseBeWHv1YePCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseBeWHv1YePCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseBeWHv1YePCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseBeWHv1YePCAPsResponseFullFilterV1]
type pcapListResponseBeWHv1YePCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseBeWHv1YePCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseBeWHv1YePCAPsResponseFullStatus string

const (
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusUnknown           PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "unknown"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusSuccess           PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "success"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusPending           PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "pending"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusRunning           PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "running"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusConversionPending PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusConversionRunning PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusComplete          PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "complete"
	PCAPListResponseBeWHv1YePCAPsResponseFullStatusFailed            PCAPListResponseBeWHv1YePCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseBeWHv1YePCAPsResponseFullSystem string

const (
	PCAPListResponseBeWHv1YePCAPsResponseFullSystemMagicTransit PCAPListResponseBeWHv1YePCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseBeWHv1YePCAPsResponseFullType string

const (
	PCAPListResponseBeWHv1YePCAPsResponseFullTypeSimple PCAPListResponseBeWHv1YePCAPsResponseFullType = "simple"
	PCAPListResponseBeWHv1YePCAPsResponseFullTypeFull   PCAPListResponseBeWHv1YePCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseBeWHv1YePCAPsResponseSimple] or
// [PCAPGetResponseBeWHv1YePCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseBeWHv1YePCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseBeWHv1YePCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseBeWHv1YePCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseBeWHv1YePCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseBeWHv1YePCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseBeWHv1YePCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseBeWHv1YePCAPsResponseSimple]
type pcapGetResponseBeWHv1YePCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseBeWHv1YePCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseBeWHv1YePCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseBeWHv1YePCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseBeWHv1YePCAPsResponseSimpleFilterV1]
type pcapGetResponseBeWHv1YePCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseBeWHv1YePCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusUnknown           PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusSuccess           PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "success"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusPending           PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusRunning           PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "running"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusConversionPending PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusConversionRunning PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusComplete          PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatusFailed            PCAPGetResponseBeWHv1YePCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseBeWHv1YePCAPsResponseSimpleSystem string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleSystemMagicTransit PCAPGetResponseBeWHv1YePCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseBeWHv1YePCAPsResponseSimpleType string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleTypeSimple PCAPGetResponseBeWHv1YePCAPsResponseSimpleType = "simple"
	PCAPGetResponseBeWHv1YePCAPsResponseSimpleTypeFull   PCAPGetResponseBeWHv1YePCAPsResponseSimpleType = "full"
)

type PCAPGetResponseBeWHv1YePCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseBeWHv1YePCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseBeWHv1YePCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseBeWHv1YePCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseBeWHv1YePCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseBeWHv1YePCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseBeWHv1YePCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseBeWHv1YePCAPsResponseFull]
type pcapGetResponseBeWHv1YePCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseBeWHv1YePCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseBeWHv1YePCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseBeWHv1YePCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseBeWHv1YePCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseBeWHv1YePCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseBeWHv1YePCAPsResponseFullFilterV1]
type pcapGetResponseBeWHv1YePCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseBeWHv1YePCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseBeWHv1YePCAPsResponseFullStatus string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusUnknown           PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "unknown"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusSuccess           PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "success"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusPending           PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "pending"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusRunning           PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "running"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusConversionPending PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusConversionRunning PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusComplete          PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "complete"
	PCAPGetResponseBeWHv1YePCAPsResponseFullStatusFailed            PCAPGetResponseBeWHv1YePCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseBeWHv1YePCAPsResponseFullSystem string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseFullSystemMagicTransit PCAPGetResponseBeWHv1YePCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseBeWHv1YePCAPsResponseFullType string

const (
	PCAPGetResponseBeWHv1YePCAPsResponseFullTypeSimple PCAPGetResponseBeWHv1YePCAPsResponseFullType = "simple"
	PCAPGetResponseBeWHv1YePCAPsResponseFullTypeFull   PCAPGetResponseBeWHv1YePCAPsResponseFullType = "full"
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
