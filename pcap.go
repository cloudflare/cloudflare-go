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

// Union satisfied by [PCAPNewResponseHj4JiX2vPCAPsResponseSimple] or
// [PCAPNewResponseHj4JiX2vPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseHj4JiX2vPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseHj4JiX2vPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseHj4JiX2vPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseHj4JiX2vPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseHj4JiX2vPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseHj4JiX2vPCAPsResponseSimple]
type pcapNewResponseHj4JiX2vPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseHj4JiX2vPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseHj4JiX2vPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1]
type pcapNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseHj4JiX2vPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusUnknown           PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusSuccess           PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusPending           PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusRunning           PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusConversionPending PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusComplete          PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatusFailed            PCAPNewResponseHj4JiX2vPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseHj4JiX2vPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseHj4JiX2vPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseHj4JiX2vPCAPsResponseSimpleType string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleTypeSimple PCAPNewResponseHj4JiX2vPCAPsResponseSimpleType = "simple"
	PCAPNewResponseHj4JiX2vPCAPsResponseSimpleTypeFull   PCAPNewResponseHj4JiX2vPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseHj4JiX2vPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseHj4JiX2vPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseHj4JiX2vPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseHj4JiX2vPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseHj4JiX2vPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseHj4JiX2vPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseHj4JiX2vPCAPsResponseFull]
type pcapNewResponseHj4JiX2vPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseHj4JiX2vPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseHj4JiX2vPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseHj4JiX2vPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseHj4JiX2vPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseHj4JiX2vPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseHj4JiX2vPCAPsResponseFullFilterV1]
type pcapNewResponseHj4JiX2vPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseHj4JiX2vPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusUnknown           PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusSuccess           PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "success"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusPending           PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "pending"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusRunning           PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "running"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusConversionPending PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusConversionRunning PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusComplete          PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "complete"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullStatusFailed            PCAPNewResponseHj4JiX2vPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseHj4JiX2vPCAPsResponseFullSystem string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseFullSystemMagicTransit PCAPNewResponseHj4JiX2vPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseHj4JiX2vPCAPsResponseFullType string

const (
	PCAPNewResponseHj4JiX2vPCAPsResponseFullTypeSimple PCAPNewResponseHj4JiX2vPCAPsResponseFullType = "simple"
	PCAPNewResponseHj4JiX2vPCAPsResponseFullTypeFull   PCAPNewResponseHj4JiX2vPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseHj4JiX2vPCAPsResponseSimple] or
// [PCAPListResponseHj4JiX2vPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseHj4JiX2vPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseHj4JiX2vPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseHj4JiX2vPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseHj4JiX2vPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseHj4JiX2vPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseHj4JiX2vPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseHj4JiX2vPCAPsResponseSimple]
type pcapListResponseHj4JiX2vPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseHj4JiX2vPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseHj4JiX2vPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseHj4JiX2vPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseHj4JiX2vPCAPsResponseSimpleFilterV1]
type pcapListResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseHj4JiX2vPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusUnknown           PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusSuccess           PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "success"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusPending           PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusRunning           PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "running"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusConversionPending PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusConversionRunning PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusComplete          PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatusFailed            PCAPListResponseHj4JiX2vPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseHj4JiX2vPCAPsResponseSimpleSystem string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleSystemMagicTransit PCAPListResponseHj4JiX2vPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseHj4JiX2vPCAPsResponseSimpleType string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleTypeSimple PCAPListResponseHj4JiX2vPCAPsResponseSimpleType = "simple"
	PCAPListResponseHj4JiX2vPCAPsResponseSimpleTypeFull   PCAPListResponseHj4JiX2vPCAPsResponseSimpleType = "full"
)

type PCAPListResponseHj4JiX2vPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseHj4JiX2vPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseHj4JiX2vPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseHj4JiX2vPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseHj4JiX2vPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseHj4JiX2vPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseHj4JiX2vPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseHj4JiX2vPCAPsResponseFull]
type pcapListResponseHj4JiX2vPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseHj4JiX2vPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseHj4JiX2vPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseHj4JiX2vPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseHj4JiX2vPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseHj4JiX2vPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseHj4JiX2vPCAPsResponseFullFilterV1]
type pcapListResponseHj4JiX2vPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseHj4JiX2vPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseHj4JiX2vPCAPsResponseFullStatus string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusUnknown           PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "unknown"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusSuccess           PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "success"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusPending           PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "pending"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusRunning           PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "running"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusConversionPending PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusConversionRunning PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusComplete          PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "complete"
	PCAPListResponseHj4JiX2vPCAPsResponseFullStatusFailed            PCAPListResponseHj4JiX2vPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseHj4JiX2vPCAPsResponseFullSystem string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseFullSystemMagicTransit PCAPListResponseHj4JiX2vPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseHj4JiX2vPCAPsResponseFullType string

const (
	PCAPListResponseHj4JiX2vPCAPsResponseFullTypeSimple PCAPListResponseHj4JiX2vPCAPsResponseFullType = "simple"
	PCAPListResponseHj4JiX2vPCAPsResponseFullTypeFull   PCAPListResponseHj4JiX2vPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseHj4JiX2vPCAPsResponseSimple] or
// [PCAPGetResponseHj4JiX2vPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseHj4JiX2vPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseHj4JiX2vPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseHj4JiX2vPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseHj4JiX2vPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseHj4JiX2vPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseHj4JiX2vPCAPsResponseSimple]
type pcapGetResponseHj4JiX2vPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseHj4JiX2vPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseHj4JiX2vPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1]
type pcapGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseHj4JiX2vPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusUnknown           PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusSuccess           PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusPending           PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusRunning           PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusConversionPending PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusComplete          PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatusFailed            PCAPGetResponseHj4JiX2vPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseHj4JiX2vPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseHj4JiX2vPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseHj4JiX2vPCAPsResponseSimpleType string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleTypeSimple PCAPGetResponseHj4JiX2vPCAPsResponseSimpleType = "simple"
	PCAPGetResponseHj4JiX2vPCAPsResponseSimpleTypeFull   PCAPGetResponseHj4JiX2vPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseHj4JiX2vPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseHj4JiX2vPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseHj4JiX2vPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseHj4JiX2vPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseHj4JiX2vPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseHj4JiX2vPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseHj4JiX2vPCAPsResponseFull]
type pcapGetResponseHj4JiX2vPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseHj4JiX2vPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseHj4JiX2vPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseHj4JiX2vPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseHj4JiX2vPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseHj4JiX2vPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseHj4JiX2vPCAPsResponseFullFilterV1]
type pcapGetResponseHj4JiX2vPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseHj4JiX2vPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusUnknown           PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusSuccess           PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "success"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusPending           PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "pending"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusRunning           PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "running"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusConversionPending PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusConversionRunning PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusComplete          PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "complete"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullStatusFailed            PCAPGetResponseHj4JiX2vPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseHj4JiX2vPCAPsResponseFullSystem string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseFullSystemMagicTransit PCAPGetResponseHj4JiX2vPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseHj4JiX2vPCAPsResponseFullType string

const (
	PCAPGetResponseHj4JiX2vPCAPsResponseFullTypeSimple PCAPGetResponseHj4JiX2vPCAPsResponseFullType = "simple"
	PCAPGetResponseHj4JiX2vPCAPsResponseFullTypeFull   PCAPGetResponseHj4JiX2vPCAPsResponseFullType = "full"
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
