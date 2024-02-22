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

// Union satisfied by [PCAPNewResponse4LpHxpI3PCAPsResponseSimple] or
// [PCAPNewResponse4LpHxpI3PCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponse4LpHxpI3PCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse4LpHxpI3PCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse4LpHxpI3PCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponse4LpHxpI3PCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse4LpHxpI3PCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponse4LpHxpI3PCAPsResponseSimple]
type pcapNewResponse4LpHxpI3PCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponse4LpHxpI3PCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse4LpHxpI3PCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1]
type pcapNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse4LpHxpI3PCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusUnknown           PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusSuccess           PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "success"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusPending           PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "pending"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusRunning           PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "running"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusConversionPending PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusConversionRunning PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusComplete          PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "complete"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatusFailed            PCAPNewResponse4LpHxpI3PCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse4LpHxpI3PCAPsResponseSimpleSystem string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleSystemMagicTransit PCAPNewResponse4LpHxpI3PCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse4LpHxpI3PCAPsResponseSimpleType string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleTypeSimple PCAPNewResponse4LpHxpI3PCAPsResponseSimpleType = "simple"
	PCAPNewResponse4LpHxpI3PCAPsResponseSimpleTypeFull   PCAPNewResponse4LpHxpI3PCAPsResponseSimpleType = "full"
)

type PCAPNewResponse4LpHxpI3PCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponse4LpHxpI3PCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse4LpHxpI3PCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse4LpHxpI3PCAPsResponseFullType `json:"type"`
	JSON pcapNewResponse4LpHxpI3PCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponse4LpHxpI3PCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponse4LpHxpI3PCAPsResponseFull]
type pcapNewResponse4LpHxpI3PCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponse4LpHxpI3PCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse4LpHxpI3PCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse4LpHxpI3PCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse4LpHxpI3PCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse4LpHxpI3PCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponse4LpHxpI3PCAPsResponseFullFilterV1]
type pcapNewResponse4LpHxpI3PCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse4LpHxpI3PCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusUnknown           PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "unknown"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusSuccess           PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "success"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusPending           PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "pending"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusRunning           PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "running"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusConversionPending PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusConversionRunning PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusComplete          PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "complete"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullStatusFailed            PCAPNewResponse4LpHxpI3PCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse4LpHxpI3PCAPsResponseFullSystem string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseFullSystemMagicTransit PCAPNewResponse4LpHxpI3PCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse4LpHxpI3PCAPsResponseFullType string

const (
	PCAPNewResponse4LpHxpI3PCAPsResponseFullTypeSimple PCAPNewResponse4LpHxpI3PCAPsResponseFullType = "simple"
	PCAPNewResponse4LpHxpI3PCAPsResponseFullTypeFull   PCAPNewResponse4LpHxpI3PCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponse4LpHxpI3PCAPsResponseSimple] or
// [PCAPListResponse4LpHxpI3PCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponse4LpHxpI3PCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponse4LpHxpI3PCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse4LpHxpI3PCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse4LpHxpI3PCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponse4LpHxpI3PCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponse4LpHxpI3PCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponse4LpHxpI3PCAPsResponseSimple]
type pcapListResponse4LpHxpI3PCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponse4LpHxpI3PCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse4LpHxpI3PCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse4LpHxpI3PCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponse4LpHxpI3PCAPsResponseSimpleFilterV1]
type pcapListResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse4LpHxpI3PCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusUnknown           PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "unknown"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusSuccess           PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "success"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusPending           PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "pending"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusRunning           PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "running"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusConversionPending PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusConversionRunning PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusComplete          PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "complete"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatusFailed            PCAPListResponse4LpHxpI3PCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse4LpHxpI3PCAPsResponseSimpleSystem string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleSystemMagicTransit PCAPListResponse4LpHxpI3PCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse4LpHxpI3PCAPsResponseSimpleType string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleTypeSimple PCAPListResponse4LpHxpI3PCAPsResponseSimpleType = "simple"
	PCAPListResponse4LpHxpI3PCAPsResponseSimpleTypeFull   PCAPListResponse4LpHxpI3PCAPsResponseSimpleType = "full"
)

type PCAPListResponse4LpHxpI3PCAPsResponseFull struct {
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
	FilterV1 PCAPListResponse4LpHxpI3PCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse4LpHxpI3PCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse4LpHxpI3PCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse4LpHxpI3PCAPsResponseFullType `json:"type"`
	JSON pcapListResponse4LpHxpI3PCAPsResponseFullJSON `json:"-"`
}

// pcapListResponse4LpHxpI3PCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponse4LpHxpI3PCAPsResponseFull]
type pcapListResponse4LpHxpI3PCAPsResponseFullJSON struct {
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

func (r *PCAPListResponse4LpHxpI3PCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse4LpHxpI3PCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse4LpHxpI3PCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse4LpHxpI3PCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse4LpHxpI3PCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponse4LpHxpI3PCAPsResponseFullFilterV1]
type pcapListResponse4LpHxpI3PCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse4LpHxpI3PCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse4LpHxpI3PCAPsResponseFullStatus string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusUnknown           PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "unknown"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusSuccess           PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "success"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusPending           PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "pending"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusRunning           PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "running"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusConversionPending PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusConversionRunning PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_running"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusComplete          PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "complete"
	PCAPListResponse4LpHxpI3PCAPsResponseFullStatusFailed            PCAPListResponse4LpHxpI3PCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse4LpHxpI3PCAPsResponseFullSystem string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseFullSystemMagicTransit PCAPListResponse4LpHxpI3PCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse4LpHxpI3PCAPsResponseFullType string

const (
	PCAPListResponse4LpHxpI3PCAPsResponseFullTypeSimple PCAPListResponse4LpHxpI3PCAPsResponseFullType = "simple"
	PCAPListResponse4LpHxpI3PCAPsResponseFullTypeFull   PCAPListResponse4LpHxpI3PCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponse4LpHxpI3PCAPsResponseSimple] or
// [PCAPGetResponse4LpHxpI3PCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponse4LpHxpI3PCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse4LpHxpI3PCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse4LpHxpI3PCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponse4LpHxpI3PCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse4LpHxpI3PCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponse4LpHxpI3PCAPsResponseSimple]
type pcapGetResponse4LpHxpI3PCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponse4LpHxpI3PCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse4LpHxpI3PCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1]
type pcapGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse4LpHxpI3PCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusUnknown           PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusSuccess           PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "success"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusPending           PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "pending"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusRunning           PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "running"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusConversionPending PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusConversionRunning PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusComplete          PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "complete"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatusFailed            PCAPGetResponse4LpHxpI3PCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse4LpHxpI3PCAPsResponseSimpleSystem string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleSystemMagicTransit PCAPGetResponse4LpHxpI3PCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse4LpHxpI3PCAPsResponseSimpleType string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleTypeSimple PCAPGetResponse4LpHxpI3PCAPsResponseSimpleType = "simple"
	PCAPGetResponse4LpHxpI3PCAPsResponseSimpleTypeFull   PCAPGetResponse4LpHxpI3PCAPsResponseSimpleType = "full"
)

type PCAPGetResponse4LpHxpI3PCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponse4LpHxpI3PCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse4LpHxpI3PCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse4LpHxpI3PCAPsResponseFullType `json:"type"`
	JSON pcapGetResponse4LpHxpI3PCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponse4LpHxpI3PCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponse4LpHxpI3PCAPsResponseFull]
type pcapGetResponse4LpHxpI3PCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponse4LpHxpI3PCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse4LpHxpI3PCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse4LpHxpI3PCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse4LpHxpI3PCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse4LpHxpI3PCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponse4LpHxpI3PCAPsResponseFullFilterV1]
type pcapGetResponse4LpHxpI3PCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse4LpHxpI3PCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusUnknown           PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "unknown"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusSuccess           PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "success"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusPending           PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "pending"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusRunning           PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "running"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusConversionPending PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusConversionRunning PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusComplete          PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "complete"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullStatusFailed            PCAPGetResponse4LpHxpI3PCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse4LpHxpI3PCAPsResponseFullSystem string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseFullSystemMagicTransit PCAPGetResponse4LpHxpI3PCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse4LpHxpI3PCAPsResponseFullType string

const (
	PCAPGetResponse4LpHxpI3PCAPsResponseFullTypeSimple PCAPGetResponse4LpHxpI3PCAPsResponseFullType = "simple"
	PCAPGetResponse4LpHxpI3PCAPsResponseFullTypeFull   PCAPGetResponse4LpHxpI3PCAPsResponseFullType = "full"
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
