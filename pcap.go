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

// Union satisfied by [PCAPNewResponseMxJLrR6tPCAPsResponseSimple] or
// [PCAPNewResponseMxJLrR6tPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseMxJLrR6tPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMxJLrR6tPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMxJLrR6tPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseMxJLrR6tPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseMxJLrR6tPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseMxJLrR6tPCAPsResponseSimple]
type pcapNewResponseMxJLrR6tPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseMxJLrR6tPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseMxJLrR6tPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1]
type pcapNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMxJLrR6tPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusUnknown           PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusSuccess           PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusPending           PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusRunning           PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusConversionPending PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusComplete          PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatusFailed            PCAPNewResponseMxJLrR6tPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseMxJLrR6tPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseMxJLrR6tPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMxJLrR6tPCAPsResponseSimpleType string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleTypeSimple PCAPNewResponseMxJLrR6tPCAPsResponseSimpleType = "simple"
	PCAPNewResponseMxJLrR6tPCAPsResponseSimpleTypeFull   PCAPNewResponseMxJLrR6tPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseMxJLrR6tPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseMxJLrR6tPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMxJLrR6tPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMxJLrR6tPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseMxJLrR6tPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseMxJLrR6tPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseMxJLrR6tPCAPsResponseFull]
type pcapNewResponseMxJLrR6tPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseMxJLrR6tPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseMxJLrR6tPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMxJLrR6tPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseMxJLrR6tPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseMxJLrR6tPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseMxJLrR6tPCAPsResponseFullFilterV1]
type pcapNewResponseMxJLrR6tPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMxJLrR6tPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusUnknown           PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusSuccess           PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "success"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusPending           PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "pending"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusRunning           PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "running"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusConversionPending PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusConversionRunning PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusComplete          PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "complete"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullStatusFailed            PCAPNewResponseMxJLrR6tPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseMxJLrR6tPCAPsResponseFullSystem string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseFullSystemMagicTransit PCAPNewResponseMxJLrR6tPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMxJLrR6tPCAPsResponseFullType string

const (
	PCAPNewResponseMxJLrR6tPCAPsResponseFullTypeSimple PCAPNewResponseMxJLrR6tPCAPsResponseFullType = "simple"
	PCAPNewResponseMxJLrR6tPCAPsResponseFullTypeFull   PCAPNewResponseMxJLrR6tPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseMxJLrR6tPCAPsResponseSimple] or
// [PCAPListResponseMxJLrR6tPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseMxJLrR6tPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseMxJLrR6tPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMxJLrR6tPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMxJLrR6tPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseMxJLrR6tPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseMxJLrR6tPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseMxJLrR6tPCAPsResponseSimple]
type pcapListResponseMxJLrR6tPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseMxJLrR6tPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseMxJLrR6tPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMxJLrR6tPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseMxJLrR6tPCAPsResponseSimpleFilterV1]
type pcapListResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMxJLrR6tPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusUnknown           PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusSuccess           PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "success"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusPending           PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusRunning           PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "running"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusConversionPending PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusConversionRunning PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusComplete          PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatusFailed            PCAPListResponseMxJLrR6tPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseMxJLrR6tPCAPsResponseSimpleSystem string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleSystemMagicTransit PCAPListResponseMxJLrR6tPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMxJLrR6tPCAPsResponseSimpleType string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleTypeSimple PCAPListResponseMxJLrR6tPCAPsResponseSimpleType = "simple"
	PCAPListResponseMxJLrR6tPCAPsResponseSimpleTypeFull   PCAPListResponseMxJLrR6tPCAPsResponseSimpleType = "full"
)

type PCAPListResponseMxJLrR6tPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseMxJLrR6tPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMxJLrR6tPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMxJLrR6tPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMxJLrR6tPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseMxJLrR6tPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseMxJLrR6tPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseMxJLrR6tPCAPsResponseFull]
type pcapListResponseMxJLrR6tPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseMxJLrR6tPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseMxJLrR6tPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMxJLrR6tPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseMxJLrR6tPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseMxJLrR6tPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseMxJLrR6tPCAPsResponseFullFilterV1]
type pcapListResponseMxJLrR6tPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMxJLrR6tPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseMxJLrR6tPCAPsResponseFullStatus string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusUnknown           PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "unknown"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusSuccess           PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "success"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusPending           PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "pending"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusRunning           PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "running"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusConversionPending PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusConversionRunning PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusComplete          PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "complete"
	PCAPListResponseMxJLrR6tPCAPsResponseFullStatusFailed            PCAPListResponseMxJLrR6tPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseMxJLrR6tPCAPsResponseFullSystem string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseFullSystemMagicTransit PCAPListResponseMxJLrR6tPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMxJLrR6tPCAPsResponseFullType string

const (
	PCAPListResponseMxJLrR6tPCAPsResponseFullTypeSimple PCAPListResponseMxJLrR6tPCAPsResponseFullType = "simple"
	PCAPListResponseMxJLrR6tPCAPsResponseFullTypeFull   PCAPListResponseMxJLrR6tPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseMxJLrR6tPCAPsResponseSimple] or
// [PCAPGetResponseMxJLrR6tPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseMxJLrR6tPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMxJLrR6tPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMxJLrR6tPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseMxJLrR6tPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseMxJLrR6tPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseMxJLrR6tPCAPsResponseSimple]
type pcapGetResponseMxJLrR6tPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseMxJLrR6tPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseMxJLrR6tPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1]
type pcapGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMxJLrR6tPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusUnknown           PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusSuccess           PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusPending           PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusRunning           PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusConversionPending PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusComplete          PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatusFailed            PCAPGetResponseMxJLrR6tPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseMxJLrR6tPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseMxJLrR6tPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMxJLrR6tPCAPsResponseSimpleType string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleTypeSimple PCAPGetResponseMxJLrR6tPCAPsResponseSimpleType = "simple"
	PCAPGetResponseMxJLrR6tPCAPsResponseSimpleTypeFull   PCAPGetResponseMxJLrR6tPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseMxJLrR6tPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseMxJLrR6tPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMxJLrR6tPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMxJLrR6tPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseMxJLrR6tPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseMxJLrR6tPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseMxJLrR6tPCAPsResponseFull]
type pcapGetResponseMxJLrR6tPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseMxJLrR6tPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseMxJLrR6tPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMxJLrR6tPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseMxJLrR6tPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseMxJLrR6tPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseMxJLrR6tPCAPsResponseFullFilterV1]
type pcapGetResponseMxJLrR6tPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMxJLrR6tPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusUnknown           PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusSuccess           PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "success"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusPending           PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "pending"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusRunning           PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "running"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusConversionPending PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusConversionRunning PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusComplete          PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "complete"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullStatusFailed            PCAPGetResponseMxJLrR6tPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseMxJLrR6tPCAPsResponseFullSystem string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseFullSystemMagicTransit PCAPGetResponseMxJLrR6tPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMxJLrR6tPCAPsResponseFullType string

const (
	PCAPGetResponseMxJLrR6tPCAPsResponseFullTypeSimple PCAPGetResponseMxJLrR6tPCAPsResponseFullType = "simple"
	PCAPGetResponseMxJLrR6tPCAPsResponseFullTypeFull   PCAPGetResponseMxJLrR6tPCAPsResponseFullType = "full"
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
