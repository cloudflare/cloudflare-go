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

// Union satisfied by [PCAPNewResponseUmFOsyUlPCAPsResponseSimple] or
// [PCAPNewResponseUmFOsyUlPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseUmFOsyUlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseUmFOsyUlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseUmFOsyUlPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseUmFOsyUlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseUmFOsyUlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseUmFOsyUlPCAPsResponseSimple]
type pcapNewResponseUmFOsyUlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseUmFOsyUlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseUmFOsyUlPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1]
type pcapNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseUmFOsyUlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusUnknown           PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusSuccess           PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusPending           PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusRunning           PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusConversionPending PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusComplete          PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatusFailed            PCAPNewResponseUmFOsyUlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseUmFOsyUlPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseUmFOsyUlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseUmFOsyUlPCAPsResponseSimpleType string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleTypeSimple PCAPNewResponseUmFOsyUlPCAPsResponseSimpleType = "simple"
	PCAPNewResponseUmFOsyUlPCAPsResponseSimpleTypeFull   PCAPNewResponseUmFOsyUlPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseUmFOsyUlPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseUmFOsyUlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseUmFOsyUlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseUmFOsyUlPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseUmFOsyUlPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseUmFOsyUlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseUmFOsyUlPCAPsResponseFull]
type pcapNewResponseUmFOsyUlPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseUmFOsyUlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseUmFOsyUlPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseUmFOsyUlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseUmFOsyUlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseUmFOsyUlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseUmFOsyUlPCAPsResponseFullFilterV1]
type pcapNewResponseUmFOsyUlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseUmFOsyUlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusUnknown           PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusSuccess           PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "success"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusPending           PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "pending"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusRunning           PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "running"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusConversionPending PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusConversionRunning PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusComplete          PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "complete"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullStatusFailed            PCAPNewResponseUmFOsyUlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseUmFOsyUlPCAPsResponseFullSystem string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseFullSystemMagicTransit PCAPNewResponseUmFOsyUlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseUmFOsyUlPCAPsResponseFullType string

const (
	PCAPNewResponseUmFOsyUlPCAPsResponseFullTypeSimple PCAPNewResponseUmFOsyUlPCAPsResponseFullType = "simple"
	PCAPNewResponseUmFOsyUlPCAPsResponseFullTypeFull   PCAPNewResponseUmFOsyUlPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseUmFOsyUlPCAPsResponseSimple] or
// [PCAPListResponseUmFOsyUlPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseUmFOsyUlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseUmFOsyUlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseUmFOsyUlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseUmFOsyUlPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseUmFOsyUlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseUmFOsyUlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseUmFOsyUlPCAPsResponseSimple]
type pcapListResponseUmFOsyUlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseUmFOsyUlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseUmFOsyUlPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseUmFOsyUlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseUmFOsyUlPCAPsResponseSimpleFilterV1]
type pcapListResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseUmFOsyUlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusUnknown           PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusSuccess           PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "success"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusPending           PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusRunning           PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "running"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusConversionPending PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusConversionRunning PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusComplete          PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatusFailed            PCAPListResponseUmFOsyUlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseUmFOsyUlPCAPsResponseSimpleSystem string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleSystemMagicTransit PCAPListResponseUmFOsyUlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseUmFOsyUlPCAPsResponseSimpleType string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleTypeSimple PCAPListResponseUmFOsyUlPCAPsResponseSimpleType = "simple"
	PCAPListResponseUmFOsyUlPCAPsResponseSimpleTypeFull   PCAPListResponseUmFOsyUlPCAPsResponseSimpleType = "full"
)

type PCAPListResponseUmFOsyUlPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseUmFOsyUlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseUmFOsyUlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseUmFOsyUlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseUmFOsyUlPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseUmFOsyUlPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseUmFOsyUlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseUmFOsyUlPCAPsResponseFull]
type pcapListResponseUmFOsyUlPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseUmFOsyUlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseUmFOsyUlPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseUmFOsyUlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseUmFOsyUlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseUmFOsyUlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseUmFOsyUlPCAPsResponseFullFilterV1]
type pcapListResponseUmFOsyUlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseUmFOsyUlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseUmFOsyUlPCAPsResponseFullStatus string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusUnknown           PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "unknown"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusSuccess           PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "success"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusPending           PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "pending"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusRunning           PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "running"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusConversionPending PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusConversionRunning PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusComplete          PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "complete"
	PCAPListResponseUmFOsyUlPCAPsResponseFullStatusFailed            PCAPListResponseUmFOsyUlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseUmFOsyUlPCAPsResponseFullSystem string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseFullSystemMagicTransit PCAPListResponseUmFOsyUlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseUmFOsyUlPCAPsResponseFullType string

const (
	PCAPListResponseUmFOsyUlPCAPsResponseFullTypeSimple PCAPListResponseUmFOsyUlPCAPsResponseFullType = "simple"
	PCAPListResponseUmFOsyUlPCAPsResponseFullTypeFull   PCAPListResponseUmFOsyUlPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseUmFOsyUlPCAPsResponseSimple] or
// [PCAPGetResponseUmFOsyUlPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseUmFOsyUlPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseUmFOsyUlPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseUmFOsyUlPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseUmFOsyUlPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseUmFOsyUlPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseUmFOsyUlPCAPsResponseSimple]
type pcapGetResponseUmFOsyUlPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseUmFOsyUlPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseUmFOsyUlPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1]
type pcapGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseUmFOsyUlPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusUnknown           PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusSuccess           PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusPending           PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusRunning           PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusConversionPending PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusComplete          PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatusFailed            PCAPGetResponseUmFOsyUlPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseUmFOsyUlPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseUmFOsyUlPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseUmFOsyUlPCAPsResponseSimpleType string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleTypeSimple PCAPGetResponseUmFOsyUlPCAPsResponseSimpleType = "simple"
	PCAPGetResponseUmFOsyUlPCAPsResponseSimpleTypeFull   PCAPGetResponseUmFOsyUlPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseUmFOsyUlPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseUmFOsyUlPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseUmFOsyUlPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseUmFOsyUlPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseUmFOsyUlPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseUmFOsyUlPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseUmFOsyUlPCAPsResponseFull]
type pcapGetResponseUmFOsyUlPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseUmFOsyUlPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseUmFOsyUlPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseUmFOsyUlPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseUmFOsyUlPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseUmFOsyUlPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseUmFOsyUlPCAPsResponseFullFilterV1]
type pcapGetResponseUmFOsyUlPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseUmFOsyUlPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusUnknown           PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusSuccess           PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "success"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusPending           PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "pending"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusRunning           PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "running"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusConversionPending PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusConversionRunning PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusComplete          PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "complete"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullStatusFailed            PCAPGetResponseUmFOsyUlPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseUmFOsyUlPCAPsResponseFullSystem string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseFullSystemMagicTransit PCAPGetResponseUmFOsyUlPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseUmFOsyUlPCAPsResponseFullType string

const (
	PCAPGetResponseUmFOsyUlPCAPsResponseFullTypeSimple PCAPGetResponseUmFOsyUlPCAPsResponseFullType = "simple"
	PCAPGetResponseUmFOsyUlPCAPsResponseFullTypeFull   PCAPGetResponseUmFOsyUlPCAPsResponseFullType = "full"
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
