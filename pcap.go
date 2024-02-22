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

// Union satisfied by [PCAPNewResponseZDm55LLxPCAPsResponseSimple] or
// [PCAPNewResponseZDm55LLxPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseZDm55LLxPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseZDm55LLxPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseZDm55LLxPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseZDm55LLxPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseZDm55LLxPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseZDm55LLxPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseZDm55LLxPCAPsResponseSimple]
type pcapNewResponseZDm55LLxPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseZDm55LLxPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseZDm55LLxPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseZDm55LLxPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseZDm55LLxPCAPsResponseSimpleFilterV1]
type pcapNewResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseZDm55LLxPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusUnknown           PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusSuccess           PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusPending           PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusRunning           PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusConversionPending PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusComplete          PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatusFailed            PCAPNewResponseZDm55LLxPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseZDm55LLxPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseZDm55LLxPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseZDm55LLxPCAPsResponseSimpleType string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleTypeSimple PCAPNewResponseZDm55LLxPCAPsResponseSimpleType = "simple"
	PCAPNewResponseZDm55LLxPCAPsResponseSimpleTypeFull   PCAPNewResponseZDm55LLxPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseZDm55LLxPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseZDm55LLxPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseZDm55LLxPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseZDm55LLxPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseZDm55LLxPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseZDm55LLxPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseZDm55LLxPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseZDm55LLxPCAPsResponseFull]
type pcapNewResponseZDm55LLxPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseZDm55LLxPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseZDm55LLxPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseZDm55LLxPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseZDm55LLxPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseZDm55LLxPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseZDm55LLxPCAPsResponseFullFilterV1]
type pcapNewResponseZDm55LLxPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseZDm55LLxPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseZDm55LLxPCAPsResponseFullStatus string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusUnknown           PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusSuccess           PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "success"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusPending           PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "pending"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusRunning           PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "running"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusConversionPending PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusConversionRunning PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusComplete          PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "complete"
	PCAPNewResponseZDm55LLxPCAPsResponseFullStatusFailed            PCAPNewResponseZDm55LLxPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseZDm55LLxPCAPsResponseFullSystem string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseFullSystemMagicTransit PCAPNewResponseZDm55LLxPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseZDm55LLxPCAPsResponseFullType string

const (
	PCAPNewResponseZDm55LLxPCAPsResponseFullTypeSimple PCAPNewResponseZDm55LLxPCAPsResponseFullType = "simple"
	PCAPNewResponseZDm55LLxPCAPsResponseFullTypeFull   PCAPNewResponseZDm55LLxPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseZDm55LLxPCAPsResponseSimple] or
// [PCAPListResponseZDm55LLxPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseZDm55LLxPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseZDm55LLxPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseZDm55LLxPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseZDm55LLxPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseZDm55LLxPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseZDm55LLxPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseZDm55LLxPCAPsResponseSimple]
type pcapListResponseZDm55LLxPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseZDm55LLxPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseZDm55LLxPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseZDm55LLxPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseZDm55LLxPCAPsResponseSimpleFilterV1]
type pcapListResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseZDm55LLxPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus string

const (
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusUnknown           PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusSuccess           PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "success"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusPending           PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusRunning           PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "running"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusConversionPending PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusConversionRunning PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusComplete          PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleStatusFailed            PCAPListResponseZDm55LLxPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseZDm55LLxPCAPsResponseSimpleSystem string

const (
	PCAPListResponseZDm55LLxPCAPsResponseSimpleSystemMagicTransit PCAPListResponseZDm55LLxPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseZDm55LLxPCAPsResponseSimpleType string

const (
	PCAPListResponseZDm55LLxPCAPsResponseSimpleTypeSimple PCAPListResponseZDm55LLxPCAPsResponseSimpleType = "simple"
	PCAPListResponseZDm55LLxPCAPsResponseSimpleTypeFull   PCAPListResponseZDm55LLxPCAPsResponseSimpleType = "full"
)

type PCAPListResponseZDm55LLxPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseZDm55LLxPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseZDm55LLxPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseZDm55LLxPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseZDm55LLxPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseZDm55LLxPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseZDm55LLxPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseZDm55LLxPCAPsResponseFull]
type pcapListResponseZDm55LLxPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseZDm55LLxPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseZDm55LLxPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseZDm55LLxPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseZDm55LLxPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseZDm55LLxPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseZDm55LLxPCAPsResponseFullFilterV1]
type pcapListResponseZDm55LLxPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseZDm55LLxPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseZDm55LLxPCAPsResponseFullStatus string

const (
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusUnknown           PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "unknown"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusSuccess           PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "success"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusPending           PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "pending"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusRunning           PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "running"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusConversionPending PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusConversionRunning PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusComplete          PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "complete"
	PCAPListResponseZDm55LLxPCAPsResponseFullStatusFailed            PCAPListResponseZDm55LLxPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseZDm55LLxPCAPsResponseFullSystem string

const (
	PCAPListResponseZDm55LLxPCAPsResponseFullSystemMagicTransit PCAPListResponseZDm55LLxPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseZDm55LLxPCAPsResponseFullType string

const (
	PCAPListResponseZDm55LLxPCAPsResponseFullTypeSimple PCAPListResponseZDm55LLxPCAPsResponseFullType = "simple"
	PCAPListResponseZDm55LLxPCAPsResponseFullTypeFull   PCAPListResponseZDm55LLxPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseZDm55LLxPCAPsResponseSimple] or
// [PCAPGetResponseZDm55LLxPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseZDm55LLxPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseZDm55LLxPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseZDm55LLxPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseZDm55LLxPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseZDm55LLxPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseZDm55LLxPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseZDm55LLxPCAPsResponseSimple]
type pcapGetResponseZDm55LLxPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseZDm55LLxPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseZDm55LLxPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseZDm55LLxPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseZDm55LLxPCAPsResponseSimpleFilterV1]
type pcapGetResponseZDm55LLxPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseZDm55LLxPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusUnknown           PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusSuccess           PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusPending           PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusRunning           PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusConversionPending PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusComplete          PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatusFailed            PCAPGetResponseZDm55LLxPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseZDm55LLxPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseZDm55LLxPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseZDm55LLxPCAPsResponseSimpleType string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleTypeSimple PCAPGetResponseZDm55LLxPCAPsResponseSimpleType = "simple"
	PCAPGetResponseZDm55LLxPCAPsResponseSimpleTypeFull   PCAPGetResponseZDm55LLxPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseZDm55LLxPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseZDm55LLxPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseZDm55LLxPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseZDm55LLxPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseZDm55LLxPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseZDm55LLxPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseZDm55LLxPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseZDm55LLxPCAPsResponseFull]
type pcapGetResponseZDm55LLxPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseZDm55LLxPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseZDm55LLxPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseZDm55LLxPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseZDm55LLxPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseZDm55LLxPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseZDm55LLxPCAPsResponseFullFilterV1]
type pcapGetResponseZDm55LLxPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseZDm55LLxPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseZDm55LLxPCAPsResponseFullStatus string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusUnknown           PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusSuccess           PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "success"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusPending           PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "pending"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusRunning           PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "running"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusConversionPending PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusConversionRunning PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusComplete          PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "complete"
	PCAPGetResponseZDm55LLxPCAPsResponseFullStatusFailed            PCAPGetResponseZDm55LLxPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseZDm55LLxPCAPsResponseFullSystem string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseFullSystemMagicTransit PCAPGetResponseZDm55LLxPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseZDm55LLxPCAPsResponseFullType string

const (
	PCAPGetResponseZDm55LLxPCAPsResponseFullTypeSimple PCAPGetResponseZDm55LLxPCAPsResponseFullType = "simple"
	PCAPGetResponseZDm55LLxPCAPsResponseFullTypeFull   PCAPGetResponseZDm55LLxPCAPsResponseFullType = "full"
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
