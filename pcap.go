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

// Union satisfied by [PCAPNewResponseRbkcI8PlPCAPsResponseSimple] or
// [PCAPNewResponseRbkcI8PlPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseRbkcI8PlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseRbkcI8PlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseRbkcI8PlPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseRbkcI8PlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseRbkcI8PlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseRbkcI8PlPCAPsResponseSimple]
type pcapNewResponseRbkcI8PlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseRbkcI8PlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseRbkcI8PlPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1]
type pcapNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseRbkcI8PlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusUnknown           PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusSuccess           PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusPending           PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusRunning           PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusConversionPending PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusComplete          PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatusFailed            PCAPNewResponseRbkcI8PlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseRbkcI8PlPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseRbkcI8PlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseRbkcI8PlPCAPsResponseSimpleType string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleTypeSimple PCAPNewResponseRbkcI8PlPCAPsResponseSimpleType = "simple"
	PCAPNewResponseRbkcI8PlPCAPsResponseSimpleTypeFull   PCAPNewResponseRbkcI8PlPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseRbkcI8PlPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseRbkcI8PlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseRbkcI8PlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseRbkcI8PlPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseRbkcI8PlPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseRbkcI8PlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseRbkcI8PlPCAPsResponseFull]
type pcapNewResponseRbkcI8PlPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseRbkcI8PlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseRbkcI8PlPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseRbkcI8PlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseRbkcI8PlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseRbkcI8PlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseRbkcI8PlPCAPsResponseFullFilterV1]
type pcapNewResponseRbkcI8PlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseRbkcI8PlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusUnknown           PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusSuccess           PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "success"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusPending           PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "pending"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusRunning           PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "running"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusConversionPending PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusConversionRunning PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusComplete          PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "complete"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullStatusFailed            PCAPNewResponseRbkcI8PlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseRbkcI8PlPCAPsResponseFullSystem string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseFullSystemMagicTransit PCAPNewResponseRbkcI8PlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseRbkcI8PlPCAPsResponseFullType string

const (
	PCAPNewResponseRbkcI8PlPCAPsResponseFullTypeSimple PCAPNewResponseRbkcI8PlPCAPsResponseFullType = "simple"
	PCAPNewResponseRbkcI8PlPCAPsResponseFullTypeFull   PCAPNewResponseRbkcI8PlPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseRbkcI8PlPCAPsResponseSimple] or
// [PCAPListResponseRbkcI8PlPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseRbkcI8PlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseRbkcI8PlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseRbkcI8PlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseRbkcI8PlPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseRbkcI8PlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseRbkcI8PlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseRbkcI8PlPCAPsResponseSimple]
type pcapListResponseRbkcI8PlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseRbkcI8PlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseRbkcI8PlPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseRbkcI8PlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseRbkcI8PlPCAPsResponseSimpleFilterV1]
type pcapListResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseRbkcI8PlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusUnknown           PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusSuccess           PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "success"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusPending           PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusRunning           PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "running"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusConversionPending PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusConversionRunning PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusComplete          PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatusFailed            PCAPListResponseRbkcI8PlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseRbkcI8PlPCAPsResponseSimpleSystem string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleSystemMagicTransit PCAPListResponseRbkcI8PlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseRbkcI8PlPCAPsResponseSimpleType string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleTypeSimple PCAPListResponseRbkcI8PlPCAPsResponseSimpleType = "simple"
	PCAPListResponseRbkcI8PlPCAPsResponseSimpleTypeFull   PCAPListResponseRbkcI8PlPCAPsResponseSimpleType = "full"
)

type PCAPListResponseRbkcI8PlPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseRbkcI8PlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseRbkcI8PlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseRbkcI8PlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseRbkcI8PlPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseRbkcI8PlPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseRbkcI8PlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseRbkcI8PlPCAPsResponseFull]
type pcapListResponseRbkcI8PlPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseRbkcI8PlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseRbkcI8PlPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseRbkcI8PlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseRbkcI8PlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseRbkcI8PlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseRbkcI8PlPCAPsResponseFullFilterV1]
type pcapListResponseRbkcI8PlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseRbkcI8PlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseRbkcI8PlPCAPsResponseFullStatus string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusUnknown           PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "unknown"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusSuccess           PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "success"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusPending           PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "pending"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusRunning           PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "running"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusConversionPending PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusConversionRunning PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusComplete          PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "complete"
	PCAPListResponseRbkcI8PlPCAPsResponseFullStatusFailed            PCAPListResponseRbkcI8PlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseRbkcI8PlPCAPsResponseFullSystem string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseFullSystemMagicTransit PCAPListResponseRbkcI8PlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseRbkcI8PlPCAPsResponseFullType string

const (
	PCAPListResponseRbkcI8PlPCAPsResponseFullTypeSimple PCAPListResponseRbkcI8PlPCAPsResponseFullType = "simple"
	PCAPListResponseRbkcI8PlPCAPsResponseFullTypeFull   PCAPListResponseRbkcI8PlPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseRbkcI8PlPCAPsResponseSimple] or
// [PCAPGetResponseRbkcI8PlPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseRbkcI8PlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseRbkcI8PlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseRbkcI8PlPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseRbkcI8PlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseRbkcI8PlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseRbkcI8PlPCAPsResponseSimple]
type pcapGetResponseRbkcI8PlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseRbkcI8PlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseRbkcI8PlPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1]
type pcapGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseRbkcI8PlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusUnknown           PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusSuccess           PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusPending           PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusRunning           PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusConversionPending PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusComplete          PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatusFailed            PCAPGetResponseRbkcI8PlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseRbkcI8PlPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseRbkcI8PlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseRbkcI8PlPCAPsResponseSimpleType string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleTypeSimple PCAPGetResponseRbkcI8PlPCAPsResponseSimpleType = "simple"
	PCAPGetResponseRbkcI8PlPCAPsResponseSimpleTypeFull   PCAPGetResponseRbkcI8PlPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseRbkcI8PlPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseRbkcI8PlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseRbkcI8PlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseRbkcI8PlPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseRbkcI8PlPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseRbkcI8PlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseRbkcI8PlPCAPsResponseFull]
type pcapGetResponseRbkcI8PlPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseRbkcI8PlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseRbkcI8PlPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseRbkcI8PlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseRbkcI8PlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseRbkcI8PlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseRbkcI8PlPCAPsResponseFullFilterV1]
type pcapGetResponseRbkcI8PlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseRbkcI8PlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusUnknown           PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusSuccess           PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "success"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusPending           PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "pending"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusRunning           PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "running"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusConversionPending PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusConversionRunning PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusComplete          PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "complete"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullStatusFailed            PCAPGetResponseRbkcI8PlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseRbkcI8PlPCAPsResponseFullSystem string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseFullSystemMagicTransit PCAPGetResponseRbkcI8PlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseRbkcI8PlPCAPsResponseFullType string

const (
	PCAPGetResponseRbkcI8PlPCAPsResponseFullTypeSimple PCAPGetResponseRbkcI8PlPCAPsResponseFullType = "simple"
	PCAPGetResponseRbkcI8PlPCAPsResponseFullTypeFull   PCAPGetResponseRbkcI8PlPCAPsResponseFullType = "full"
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
