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

// Union satisfied by [PCAPNewResponse5aubKvqtPCAPsResponseSimple] or
// [PCAPNewResponse5aubKvqtPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponse5aubKvqtPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponse5aubKvqtPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse5aubKvqtPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse5aubKvqtPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponse5aubKvqtPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponse5aubKvqtPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponse5aubKvqtPCAPsResponseSimple]
type pcapNewResponse5aubKvqtPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponse5aubKvqtPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse5aubKvqtPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse5aubKvqtPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponse5aubKvqtPCAPsResponseSimpleFilterV1]
type pcapNewResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse5aubKvqtPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusUnknown           PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusSuccess           PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "success"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusPending           PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusRunning           PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "running"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusConversionPending PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusConversionRunning PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusComplete          PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatusFailed            PCAPNewResponse5aubKvqtPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse5aubKvqtPCAPsResponseSimpleSystem string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleSystemMagicTransit PCAPNewResponse5aubKvqtPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse5aubKvqtPCAPsResponseSimpleType string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleTypeSimple PCAPNewResponse5aubKvqtPCAPsResponseSimpleType = "simple"
	PCAPNewResponse5aubKvqtPCAPsResponseSimpleTypeFull   PCAPNewResponse5aubKvqtPCAPsResponseSimpleType = "full"
)

type PCAPNewResponse5aubKvqtPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponse5aubKvqtPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponse5aubKvqtPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponse5aubKvqtPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponse5aubKvqtPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponse5aubKvqtPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponse5aubKvqtPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponse5aubKvqtPCAPsResponseFull]
type pcapNewResponse5aubKvqtPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponse5aubKvqtPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponse5aubKvqtPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponse5aubKvqtPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponse5aubKvqtPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponse5aubKvqtPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponse5aubKvqtPCAPsResponseFullFilterV1]
type pcapNewResponse5aubKvqtPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponse5aubKvqtPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponse5aubKvqtPCAPsResponseFullStatus string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusUnknown           PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "unknown"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusSuccess           PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "success"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusPending           PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "pending"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusRunning           PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "running"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusConversionPending PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusConversionRunning PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusComplete          PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "complete"
	PCAPNewResponse5aubKvqtPCAPsResponseFullStatusFailed            PCAPNewResponse5aubKvqtPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponse5aubKvqtPCAPsResponseFullSystem string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseFullSystemMagicTransit PCAPNewResponse5aubKvqtPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponse5aubKvqtPCAPsResponseFullType string

const (
	PCAPNewResponse5aubKvqtPCAPsResponseFullTypeSimple PCAPNewResponse5aubKvqtPCAPsResponseFullType = "simple"
	PCAPNewResponse5aubKvqtPCAPsResponseFullTypeFull   PCAPNewResponse5aubKvqtPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponse5aubKvqtPCAPsResponseSimple] or
// [PCAPListResponse5aubKvqtPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponse5aubKvqtPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponse5aubKvqtPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse5aubKvqtPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse5aubKvqtPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponse5aubKvqtPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponse5aubKvqtPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponse5aubKvqtPCAPsResponseSimple]
type pcapListResponse5aubKvqtPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponse5aubKvqtPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse5aubKvqtPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse5aubKvqtPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponse5aubKvqtPCAPsResponseSimpleFilterV1]
type pcapListResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse5aubKvqtPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus string

const (
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusUnknown           PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusSuccess           PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "success"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusPending           PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "pending"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusRunning           PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "running"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusConversionPending PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusConversionRunning PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusComplete          PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "complete"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleStatusFailed            PCAPListResponse5aubKvqtPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse5aubKvqtPCAPsResponseSimpleSystem string

const (
	PCAPListResponse5aubKvqtPCAPsResponseSimpleSystemMagicTransit PCAPListResponse5aubKvqtPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse5aubKvqtPCAPsResponseSimpleType string

const (
	PCAPListResponse5aubKvqtPCAPsResponseSimpleTypeSimple PCAPListResponse5aubKvqtPCAPsResponseSimpleType = "simple"
	PCAPListResponse5aubKvqtPCAPsResponseSimpleTypeFull   PCAPListResponse5aubKvqtPCAPsResponseSimpleType = "full"
)

type PCAPListResponse5aubKvqtPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponse5aubKvqtPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponse5aubKvqtPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponse5aubKvqtPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponse5aubKvqtPCAPsResponseFullType `json:"type"`
	JSON pcapListResponse5aubKvqtPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponse5aubKvqtPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponse5aubKvqtPCAPsResponseFull]
type pcapListResponse5aubKvqtPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponse5aubKvqtPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponse5aubKvqtPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponse5aubKvqtPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponse5aubKvqtPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponse5aubKvqtPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponse5aubKvqtPCAPsResponseFullFilterV1]
type pcapListResponse5aubKvqtPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponse5aubKvqtPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponse5aubKvqtPCAPsResponseFullStatus string

const (
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusUnknown           PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "unknown"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusSuccess           PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "success"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusPending           PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "pending"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusRunning           PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "running"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusConversionPending PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusConversionRunning PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusComplete          PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "complete"
	PCAPListResponse5aubKvqtPCAPsResponseFullStatusFailed            PCAPListResponse5aubKvqtPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponse5aubKvqtPCAPsResponseFullSystem string

const (
	PCAPListResponse5aubKvqtPCAPsResponseFullSystemMagicTransit PCAPListResponse5aubKvqtPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponse5aubKvqtPCAPsResponseFullType string

const (
	PCAPListResponse5aubKvqtPCAPsResponseFullTypeSimple PCAPListResponse5aubKvqtPCAPsResponseFullType = "simple"
	PCAPListResponse5aubKvqtPCAPsResponseFullTypeFull   PCAPListResponse5aubKvqtPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponse5aubKvqtPCAPsResponseSimple] or
// [PCAPGetResponse5aubKvqtPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponse5aubKvqtPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponse5aubKvqtPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse5aubKvqtPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse5aubKvqtPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponse5aubKvqtPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponse5aubKvqtPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponse5aubKvqtPCAPsResponseSimple]
type pcapGetResponse5aubKvqtPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponse5aubKvqtPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse5aubKvqtPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse5aubKvqtPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponse5aubKvqtPCAPsResponseSimpleFilterV1]
type pcapGetResponse5aubKvqtPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse5aubKvqtPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusUnknown           PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusSuccess           PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "success"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusPending           PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusRunning           PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "running"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusConversionPending PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusConversionRunning PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusComplete          PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatusFailed            PCAPGetResponse5aubKvqtPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse5aubKvqtPCAPsResponseSimpleSystem string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleSystemMagicTransit PCAPGetResponse5aubKvqtPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse5aubKvqtPCAPsResponseSimpleType string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleTypeSimple PCAPGetResponse5aubKvqtPCAPsResponseSimpleType = "simple"
	PCAPGetResponse5aubKvqtPCAPsResponseSimpleTypeFull   PCAPGetResponse5aubKvqtPCAPsResponseSimpleType = "full"
)

type PCAPGetResponse5aubKvqtPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponse5aubKvqtPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponse5aubKvqtPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponse5aubKvqtPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponse5aubKvqtPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponse5aubKvqtPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponse5aubKvqtPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponse5aubKvqtPCAPsResponseFull]
type pcapGetResponse5aubKvqtPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponse5aubKvqtPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponse5aubKvqtPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponse5aubKvqtPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponse5aubKvqtPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponse5aubKvqtPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponse5aubKvqtPCAPsResponseFullFilterV1]
type pcapGetResponse5aubKvqtPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponse5aubKvqtPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponse5aubKvqtPCAPsResponseFullStatus string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusUnknown           PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "unknown"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusSuccess           PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "success"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusPending           PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "pending"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusRunning           PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "running"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusConversionPending PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusConversionRunning PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusComplete          PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "complete"
	PCAPGetResponse5aubKvqtPCAPsResponseFullStatusFailed            PCAPGetResponse5aubKvqtPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponse5aubKvqtPCAPsResponseFullSystem string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseFullSystemMagicTransit PCAPGetResponse5aubKvqtPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponse5aubKvqtPCAPsResponseFullType string

const (
	PCAPGetResponse5aubKvqtPCAPsResponseFullTypeSimple PCAPGetResponse5aubKvqtPCAPsResponseFullType = "simple"
	PCAPGetResponse5aubKvqtPCAPsResponseFullTypeFull   PCAPGetResponse5aubKvqtPCAPsResponseFullType = "full"
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
