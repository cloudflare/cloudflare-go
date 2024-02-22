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

// Union satisfied by [PCAPNewResponseUz8acDqrPCAPsResponseSimple] or
// [PCAPNewResponseUz8acDqrPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseUz8acDqrPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseUz8acDqrPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseUz8acDqrPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseUz8acDqrPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseUz8acDqrPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseUz8acDqrPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseUz8acDqrPCAPsResponseSimple]
type pcapNewResponseUz8acDqrPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseUz8acDqrPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseUz8acDqrPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseUz8acDqrPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseUz8acDqrPCAPsResponseSimpleFilterV1]
type pcapNewResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseUz8acDqrPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusUnknown           PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusSuccess           PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusPending           PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusRunning           PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusConversionPending PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusComplete          PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatusFailed            PCAPNewResponseUz8acDqrPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseUz8acDqrPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseUz8acDqrPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseUz8acDqrPCAPsResponseSimpleType string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleTypeSimple PCAPNewResponseUz8acDqrPCAPsResponseSimpleType = "simple"
	PCAPNewResponseUz8acDqrPCAPsResponseSimpleTypeFull   PCAPNewResponseUz8acDqrPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseUz8acDqrPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseUz8acDqrPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseUz8acDqrPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseUz8acDqrPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseUz8acDqrPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseUz8acDqrPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseUz8acDqrPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseUz8acDqrPCAPsResponseFull]
type pcapNewResponseUz8acDqrPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseUz8acDqrPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseUz8acDqrPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseUz8acDqrPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseUz8acDqrPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseUz8acDqrPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseUz8acDqrPCAPsResponseFullFilterV1]
type pcapNewResponseUz8acDqrPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseUz8acDqrPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseUz8acDqrPCAPsResponseFullStatus string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusUnknown           PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusSuccess           PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "success"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusPending           PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "pending"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusRunning           PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "running"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusConversionPending PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusConversionRunning PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusComplete          PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "complete"
	PCAPNewResponseUz8acDqrPCAPsResponseFullStatusFailed            PCAPNewResponseUz8acDqrPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseUz8acDqrPCAPsResponseFullSystem string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseFullSystemMagicTransit PCAPNewResponseUz8acDqrPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseUz8acDqrPCAPsResponseFullType string

const (
	PCAPNewResponseUz8acDqrPCAPsResponseFullTypeSimple PCAPNewResponseUz8acDqrPCAPsResponseFullType = "simple"
	PCAPNewResponseUz8acDqrPCAPsResponseFullTypeFull   PCAPNewResponseUz8acDqrPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseUz8acDqrPCAPsResponseSimple] or
// [PCAPListResponseUz8acDqrPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseUz8acDqrPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseUz8acDqrPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseUz8acDqrPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseUz8acDqrPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseUz8acDqrPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseUz8acDqrPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseUz8acDqrPCAPsResponseSimple]
type pcapListResponseUz8acDqrPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseUz8acDqrPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseUz8acDqrPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseUz8acDqrPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseUz8acDqrPCAPsResponseSimpleFilterV1]
type pcapListResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseUz8acDqrPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus string

const (
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusUnknown           PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusSuccess           PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "success"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusPending           PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusRunning           PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "running"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusConversionPending PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusConversionRunning PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusComplete          PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleStatusFailed            PCAPListResponseUz8acDqrPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseUz8acDqrPCAPsResponseSimpleSystem string

const (
	PCAPListResponseUz8acDqrPCAPsResponseSimpleSystemMagicTransit PCAPListResponseUz8acDqrPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseUz8acDqrPCAPsResponseSimpleType string

const (
	PCAPListResponseUz8acDqrPCAPsResponseSimpleTypeSimple PCAPListResponseUz8acDqrPCAPsResponseSimpleType = "simple"
	PCAPListResponseUz8acDqrPCAPsResponseSimpleTypeFull   PCAPListResponseUz8acDqrPCAPsResponseSimpleType = "full"
)

type PCAPListResponseUz8acDqrPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseUz8acDqrPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseUz8acDqrPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseUz8acDqrPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseUz8acDqrPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseUz8acDqrPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseUz8acDqrPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseUz8acDqrPCAPsResponseFull]
type pcapListResponseUz8acDqrPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseUz8acDqrPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseUz8acDqrPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseUz8acDqrPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseUz8acDqrPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseUz8acDqrPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseUz8acDqrPCAPsResponseFullFilterV1]
type pcapListResponseUz8acDqrPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseUz8acDqrPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseUz8acDqrPCAPsResponseFullStatus string

const (
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusUnknown           PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "unknown"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusSuccess           PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "success"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusPending           PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "pending"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusRunning           PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "running"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusConversionPending PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusConversionRunning PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusComplete          PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "complete"
	PCAPListResponseUz8acDqrPCAPsResponseFullStatusFailed            PCAPListResponseUz8acDqrPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseUz8acDqrPCAPsResponseFullSystem string

const (
	PCAPListResponseUz8acDqrPCAPsResponseFullSystemMagicTransit PCAPListResponseUz8acDqrPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseUz8acDqrPCAPsResponseFullType string

const (
	PCAPListResponseUz8acDqrPCAPsResponseFullTypeSimple PCAPListResponseUz8acDqrPCAPsResponseFullType = "simple"
	PCAPListResponseUz8acDqrPCAPsResponseFullTypeFull   PCAPListResponseUz8acDqrPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseUz8acDqrPCAPsResponseSimple] or
// [PCAPGetResponseUz8acDqrPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseUz8acDqrPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseUz8acDqrPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseUz8acDqrPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseUz8acDqrPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseUz8acDqrPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseUz8acDqrPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseUz8acDqrPCAPsResponseSimple]
type pcapGetResponseUz8acDqrPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseUz8acDqrPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseUz8acDqrPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseUz8acDqrPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseUz8acDqrPCAPsResponseSimpleFilterV1]
type pcapGetResponseUz8acDqrPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseUz8acDqrPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusUnknown           PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusSuccess           PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusPending           PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusRunning           PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusConversionPending PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusComplete          PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatusFailed            PCAPGetResponseUz8acDqrPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseUz8acDqrPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseUz8acDqrPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseUz8acDqrPCAPsResponseSimpleType string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleTypeSimple PCAPGetResponseUz8acDqrPCAPsResponseSimpleType = "simple"
	PCAPGetResponseUz8acDqrPCAPsResponseSimpleTypeFull   PCAPGetResponseUz8acDqrPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseUz8acDqrPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseUz8acDqrPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseUz8acDqrPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseUz8acDqrPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseUz8acDqrPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseUz8acDqrPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseUz8acDqrPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseUz8acDqrPCAPsResponseFull]
type pcapGetResponseUz8acDqrPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseUz8acDqrPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseUz8acDqrPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseUz8acDqrPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseUz8acDqrPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseUz8acDqrPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseUz8acDqrPCAPsResponseFullFilterV1]
type pcapGetResponseUz8acDqrPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseUz8acDqrPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseUz8acDqrPCAPsResponseFullStatus string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusUnknown           PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusSuccess           PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "success"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusPending           PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "pending"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusRunning           PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "running"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusConversionPending PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusConversionRunning PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusComplete          PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "complete"
	PCAPGetResponseUz8acDqrPCAPsResponseFullStatusFailed            PCAPGetResponseUz8acDqrPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseUz8acDqrPCAPsResponseFullSystem string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseFullSystemMagicTransit PCAPGetResponseUz8acDqrPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseUz8acDqrPCAPsResponseFullType string

const (
	PCAPGetResponseUz8acDqrPCAPsResponseFullTypeSimple PCAPGetResponseUz8acDqrPCAPsResponseFullType = "simple"
	PCAPGetResponseUz8acDqrPCAPsResponseFullTypeFull   PCAPGetResponseUz8acDqrPCAPsResponseFullType = "full"
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
