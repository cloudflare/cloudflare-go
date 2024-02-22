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

// Union satisfied by [PCAPNewResponseAGiBwgNdPCAPsResponseSimple] or
// [PCAPNewResponseAGiBwgNdPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseAGiBwgNdPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseAGiBwgNdPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseAGiBwgNdPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseAGiBwgNdPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseAGiBwgNdPCAPsResponseSimple]
type pcapNewResponseAGiBwgNdPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseAGiBwgNdPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseAGiBwgNdPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1]
type pcapNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseAGiBwgNdPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusUnknown           PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusSuccess           PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusPending           PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusRunning           PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusConversionPending PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusComplete          PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatusFailed            PCAPNewResponseAGiBwgNdPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseAGiBwgNdPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseAGiBwgNdPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseAGiBwgNdPCAPsResponseSimpleType string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleTypeSimple PCAPNewResponseAGiBwgNdPCAPsResponseSimpleType = "simple"
	PCAPNewResponseAGiBwgNdPCAPsResponseSimpleTypeFull   PCAPNewResponseAGiBwgNdPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseAGiBwgNdPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseAGiBwgNdPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseAGiBwgNdPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseAGiBwgNdPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseAGiBwgNdPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseAGiBwgNdPCAPsResponseFull]
type pcapNewResponseAGiBwgNdPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseAGiBwgNdPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseAGiBwgNdPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseAGiBwgNdPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseAGiBwgNdPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseAGiBwgNdPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseAGiBwgNdPCAPsResponseFullFilterV1]
type pcapNewResponseAGiBwgNdPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseAGiBwgNdPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusUnknown           PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusSuccess           PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "success"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusPending           PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "pending"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusRunning           PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "running"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusConversionPending PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusConversionRunning PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusComplete          PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "complete"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullStatusFailed            PCAPNewResponseAGiBwgNdPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseAGiBwgNdPCAPsResponseFullSystem string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseFullSystemMagicTransit PCAPNewResponseAGiBwgNdPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseAGiBwgNdPCAPsResponseFullType string

const (
	PCAPNewResponseAGiBwgNdPCAPsResponseFullTypeSimple PCAPNewResponseAGiBwgNdPCAPsResponseFullType = "simple"
	PCAPNewResponseAGiBwgNdPCAPsResponseFullTypeFull   PCAPNewResponseAGiBwgNdPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseAGiBwgNdPCAPsResponseSimple] or
// [PCAPListResponseAGiBwgNdPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseAGiBwgNdPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseAGiBwgNdPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseAGiBwgNdPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseAGiBwgNdPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseAGiBwgNdPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseAGiBwgNdPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseAGiBwgNdPCAPsResponseSimple]
type pcapListResponseAGiBwgNdPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseAGiBwgNdPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseAGiBwgNdPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseAGiBwgNdPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseAGiBwgNdPCAPsResponseSimpleFilterV1]
type pcapListResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseAGiBwgNdPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusUnknown           PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusSuccess           PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "success"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusPending           PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusRunning           PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "running"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusConversionPending PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusConversionRunning PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusComplete          PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatusFailed            PCAPListResponseAGiBwgNdPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseAGiBwgNdPCAPsResponseSimpleSystem string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleSystemMagicTransit PCAPListResponseAGiBwgNdPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseAGiBwgNdPCAPsResponseSimpleType string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleTypeSimple PCAPListResponseAGiBwgNdPCAPsResponseSimpleType = "simple"
	PCAPListResponseAGiBwgNdPCAPsResponseSimpleTypeFull   PCAPListResponseAGiBwgNdPCAPsResponseSimpleType = "full"
)

type PCAPListResponseAGiBwgNdPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseAGiBwgNdPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseAGiBwgNdPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseAGiBwgNdPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseAGiBwgNdPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseAGiBwgNdPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseAGiBwgNdPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseAGiBwgNdPCAPsResponseFull]
type pcapListResponseAGiBwgNdPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseAGiBwgNdPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseAGiBwgNdPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseAGiBwgNdPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseAGiBwgNdPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseAGiBwgNdPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseAGiBwgNdPCAPsResponseFullFilterV1]
type pcapListResponseAGiBwgNdPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseAGiBwgNdPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseAGiBwgNdPCAPsResponseFullStatus string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusUnknown           PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "unknown"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusSuccess           PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "success"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusPending           PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "pending"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusRunning           PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "running"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusConversionPending PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusConversionRunning PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusComplete          PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "complete"
	PCAPListResponseAGiBwgNdPCAPsResponseFullStatusFailed            PCAPListResponseAGiBwgNdPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseAGiBwgNdPCAPsResponseFullSystem string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseFullSystemMagicTransit PCAPListResponseAGiBwgNdPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseAGiBwgNdPCAPsResponseFullType string

const (
	PCAPListResponseAGiBwgNdPCAPsResponseFullTypeSimple PCAPListResponseAGiBwgNdPCAPsResponseFullType = "simple"
	PCAPListResponseAGiBwgNdPCAPsResponseFullTypeFull   PCAPListResponseAGiBwgNdPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseAGiBwgNdPCAPsResponseSimple] or
// [PCAPGetResponseAGiBwgNdPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseAGiBwgNdPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseAGiBwgNdPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseAGiBwgNdPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseAGiBwgNdPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseAGiBwgNdPCAPsResponseSimple]
type pcapGetResponseAGiBwgNdPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseAGiBwgNdPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseAGiBwgNdPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1]
type pcapGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseAGiBwgNdPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusUnknown           PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusSuccess           PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusPending           PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusRunning           PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusConversionPending PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusComplete          PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatusFailed            PCAPGetResponseAGiBwgNdPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseAGiBwgNdPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseAGiBwgNdPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseAGiBwgNdPCAPsResponseSimpleType string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleTypeSimple PCAPGetResponseAGiBwgNdPCAPsResponseSimpleType = "simple"
	PCAPGetResponseAGiBwgNdPCAPsResponseSimpleTypeFull   PCAPGetResponseAGiBwgNdPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseAGiBwgNdPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseAGiBwgNdPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseAGiBwgNdPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseAGiBwgNdPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseAGiBwgNdPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseAGiBwgNdPCAPsResponseFull]
type pcapGetResponseAGiBwgNdPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseAGiBwgNdPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseAGiBwgNdPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseAGiBwgNdPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseAGiBwgNdPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseAGiBwgNdPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseAGiBwgNdPCAPsResponseFullFilterV1]
type pcapGetResponseAGiBwgNdPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseAGiBwgNdPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusUnknown           PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusSuccess           PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "success"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusPending           PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "pending"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusRunning           PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "running"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusConversionPending PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusConversionRunning PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusComplete          PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "complete"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullStatusFailed            PCAPGetResponseAGiBwgNdPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseAGiBwgNdPCAPsResponseFullSystem string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseFullSystemMagicTransit PCAPGetResponseAGiBwgNdPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseAGiBwgNdPCAPsResponseFullType string

const (
	PCAPGetResponseAGiBwgNdPCAPsResponseFullTypeSimple PCAPGetResponseAGiBwgNdPCAPsResponseFullType = "simple"
	PCAPGetResponseAGiBwgNdPCAPsResponseFullTypeFull   PCAPGetResponseAGiBwgNdPCAPsResponseFullType = "full"
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
