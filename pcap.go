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
func (r *PCAPService) New(ctx context.Context, params PCAPNewParams, opts ...option.RequestOption) (res *PCAPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all packet capture requests for an account.
func (r *PCAPService) List(ctx context.Context, query PCAPListParams, opts ...option.RequestOption) (res *[]PCAPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a PCAP request by id.
func (r *PCAPService) Get(ctx context.Context, pcapID string, query PCAPGetParams, opts ...option.RequestOption) (res *PCAPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/%s", query.AccountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PCAPNewResponseMagicVisibilityPCAPsResponseSimple] or
// [PCAPNewResponseMagicVisibilityPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON metadata
// for the struct [PCAPNewResponseMagicVisibilityPCAPsResponseSimple]
type pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                       `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPNewResponseMagicVisibilityPCAPsResponseFull]
type pcapNewResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                     `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPNewResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPNewResponseMagicVisibilityPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseMagicVisibilityPCAPsResponseSimple] or
// [PCAPListResponseMagicVisibilityPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON
// metadata for the struct [PCAPListResponseMagicVisibilityPCAPsResponseSimple]
type pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                        `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPListResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPListResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

type PCAPListResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPListResponseMagicVisibilityPCAPsResponseFull]
type pcapListResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseMagicVisibilityPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                      `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPListResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPListResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPListResponseMagicVisibilityPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseMagicVisibilityPCAPsResponseSimple] or
// [PCAPGetResponseMagicVisibilityPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON metadata
// for the struct [PCAPGetResponseMagicVisibilityPCAPsResponseSimple]
type pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                       `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPGetResponseMagicVisibilityPCAPsResponseFull]
type pcapGetResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                     `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPGetResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPGetResponseMagicVisibilityPCAPsResponseFullType = "full"
)

type PCAPNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

type PCAPListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

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

type PCAPGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
