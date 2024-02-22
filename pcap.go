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

// Union satisfied by [PCAPNewResponseO7MYbCYsPCAPsResponseSimple] or
// [PCAPNewResponseO7MYbCYsPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseO7MYbCYsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseO7MYbCYsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseO7MYbCYsPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseO7MYbCYsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseO7MYbCYsPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseO7MYbCYsPCAPsResponseSimple]
type pcapNewResponseO7MYbCYsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseO7MYbCYsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseO7MYbCYsPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1]
type pcapNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseO7MYbCYsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusUnknown           PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusSuccess           PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusPending           PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusRunning           PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusConversionPending PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusComplete          PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatusFailed            PCAPNewResponseO7MYbCYsPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseO7MYbCYsPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseO7MYbCYsPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseO7MYbCYsPCAPsResponseSimpleType string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleTypeSimple PCAPNewResponseO7MYbCYsPCAPsResponseSimpleType = "simple"
	PCAPNewResponseO7MYbCYsPCAPsResponseSimpleTypeFull   PCAPNewResponseO7MYbCYsPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseO7MYbCYsPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseO7MYbCYsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseO7MYbCYsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseO7MYbCYsPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseO7MYbCYsPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseO7MYbCYsPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseO7MYbCYsPCAPsResponseFull]
type pcapNewResponseO7MYbCYsPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseO7MYbCYsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseO7MYbCYsPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseO7MYbCYsPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseO7MYbCYsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseO7MYbCYsPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseO7MYbCYsPCAPsResponseFullFilterV1]
type pcapNewResponseO7MYbCYsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseO7MYbCYsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusUnknown           PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusSuccess           PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "success"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusPending           PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "pending"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusRunning           PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "running"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusConversionPending PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusConversionRunning PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusComplete          PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "complete"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullStatusFailed            PCAPNewResponseO7MYbCYsPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseO7MYbCYsPCAPsResponseFullSystem string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseFullSystemMagicTransit PCAPNewResponseO7MYbCYsPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseO7MYbCYsPCAPsResponseFullType string

const (
	PCAPNewResponseO7MYbCYsPCAPsResponseFullTypeSimple PCAPNewResponseO7MYbCYsPCAPsResponseFullType = "simple"
	PCAPNewResponseO7MYbCYsPCAPsResponseFullTypeFull   PCAPNewResponseO7MYbCYsPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseO7MYbCYsPCAPsResponseSimple] or
// [PCAPListResponseO7MYbCYsPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseO7MYbCYsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseO7MYbCYsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseO7MYbCYsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseO7MYbCYsPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseO7MYbCYsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseO7MYbCYsPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseO7MYbCYsPCAPsResponseSimple]
type pcapListResponseO7MYbCYsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseO7MYbCYsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseO7MYbCYsPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseO7MYbCYsPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseO7MYbCYsPCAPsResponseSimpleFilterV1]
type pcapListResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseO7MYbCYsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusUnknown           PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusSuccess           PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "success"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusPending           PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusRunning           PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "running"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusConversionPending PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusConversionRunning PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusComplete          PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatusFailed            PCAPListResponseO7MYbCYsPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseO7MYbCYsPCAPsResponseSimpleSystem string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleSystemMagicTransit PCAPListResponseO7MYbCYsPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseO7MYbCYsPCAPsResponseSimpleType string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleTypeSimple PCAPListResponseO7MYbCYsPCAPsResponseSimpleType = "simple"
	PCAPListResponseO7MYbCYsPCAPsResponseSimpleTypeFull   PCAPListResponseO7MYbCYsPCAPsResponseSimpleType = "full"
)

type PCAPListResponseO7MYbCYsPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseO7MYbCYsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseO7MYbCYsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseO7MYbCYsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseO7MYbCYsPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseO7MYbCYsPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseO7MYbCYsPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseO7MYbCYsPCAPsResponseFull]
type pcapListResponseO7MYbCYsPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseO7MYbCYsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseO7MYbCYsPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseO7MYbCYsPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseO7MYbCYsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseO7MYbCYsPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseO7MYbCYsPCAPsResponseFullFilterV1]
type pcapListResponseO7MYbCYsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseO7MYbCYsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseO7MYbCYsPCAPsResponseFullStatus string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusUnknown           PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "unknown"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusSuccess           PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "success"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusPending           PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "pending"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusRunning           PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "running"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusConversionPending PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusConversionRunning PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusComplete          PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "complete"
	PCAPListResponseO7MYbCYsPCAPsResponseFullStatusFailed            PCAPListResponseO7MYbCYsPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseO7MYbCYsPCAPsResponseFullSystem string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseFullSystemMagicTransit PCAPListResponseO7MYbCYsPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseO7MYbCYsPCAPsResponseFullType string

const (
	PCAPListResponseO7MYbCYsPCAPsResponseFullTypeSimple PCAPListResponseO7MYbCYsPCAPsResponseFullType = "simple"
	PCAPListResponseO7MYbCYsPCAPsResponseFullTypeFull   PCAPListResponseO7MYbCYsPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseO7MYbCYsPCAPsResponseSimple] or
// [PCAPGetResponseO7MYbCYsPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseO7MYbCYsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseO7MYbCYsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseO7MYbCYsPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseO7MYbCYsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseO7MYbCYsPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseO7MYbCYsPCAPsResponseSimple]
type pcapGetResponseO7MYbCYsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseO7MYbCYsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseO7MYbCYsPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1]
type pcapGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseO7MYbCYsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusUnknown           PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusSuccess           PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusPending           PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusRunning           PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusConversionPending PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusComplete          PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatusFailed            PCAPGetResponseO7MYbCYsPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseO7MYbCYsPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseO7MYbCYsPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseO7MYbCYsPCAPsResponseSimpleType string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleTypeSimple PCAPGetResponseO7MYbCYsPCAPsResponseSimpleType = "simple"
	PCAPGetResponseO7MYbCYsPCAPsResponseSimpleTypeFull   PCAPGetResponseO7MYbCYsPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseO7MYbCYsPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseO7MYbCYsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseO7MYbCYsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseO7MYbCYsPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseO7MYbCYsPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseO7MYbCYsPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseO7MYbCYsPCAPsResponseFull]
type pcapGetResponseO7MYbCYsPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseO7MYbCYsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseO7MYbCYsPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseO7MYbCYsPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseO7MYbCYsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseO7MYbCYsPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseO7MYbCYsPCAPsResponseFullFilterV1]
type pcapGetResponseO7MYbCYsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseO7MYbCYsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusUnknown           PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusSuccess           PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "success"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusPending           PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "pending"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusRunning           PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "running"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusConversionPending PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusConversionRunning PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusComplete          PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "complete"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullStatusFailed            PCAPGetResponseO7MYbCYsPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseO7MYbCYsPCAPsResponseFullSystem string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseFullSystemMagicTransit PCAPGetResponseO7MYbCYsPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseO7MYbCYsPCAPsResponseFullType string

const (
	PCAPGetResponseO7MYbCYsPCAPsResponseFullTypeSimple PCAPGetResponseO7MYbCYsPCAPsResponseFullType = "simple"
	PCAPGetResponseO7MYbCYsPCAPsResponseFullTypeFull   PCAPGetResponseO7MYbCYsPCAPsResponseFullType = "full"
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
