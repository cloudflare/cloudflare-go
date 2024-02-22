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

// Union satisfied by [PCAPNewResponseT8F7ioUkPCAPsResponseSimple] or
// [PCAPNewResponseT8F7ioUkPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseT8F7ioUkPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseT8F7ioUkPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseT8F7ioUkPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseT8F7ioUkPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseT8F7ioUkPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseT8F7ioUkPCAPsResponseSimple]
type pcapNewResponseT8F7ioUkPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseT8F7ioUkPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseT8F7ioUkPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1]
type pcapNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseT8F7ioUkPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusUnknown           PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusSuccess           PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusPending           PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusRunning           PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusConversionPending PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusComplete          PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatusFailed            PCAPNewResponseT8F7ioUkPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseT8F7ioUkPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseT8F7ioUkPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseT8F7ioUkPCAPsResponseSimpleType string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleTypeSimple PCAPNewResponseT8F7ioUkPCAPsResponseSimpleType = "simple"
	PCAPNewResponseT8F7ioUkPCAPsResponseSimpleTypeFull   PCAPNewResponseT8F7ioUkPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseT8F7ioUkPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseT8F7ioUkPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseT8F7ioUkPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseT8F7ioUkPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseT8F7ioUkPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseT8F7ioUkPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseT8F7ioUkPCAPsResponseFull]
type pcapNewResponseT8F7ioUkPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseT8F7ioUkPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseT8F7ioUkPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseT8F7ioUkPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseT8F7ioUkPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseT8F7ioUkPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseT8F7ioUkPCAPsResponseFullFilterV1]
type pcapNewResponseT8F7ioUkPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseT8F7ioUkPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusUnknown           PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusSuccess           PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "success"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusPending           PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "pending"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusRunning           PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "running"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusConversionPending PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusConversionRunning PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusComplete          PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "complete"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullStatusFailed            PCAPNewResponseT8F7ioUkPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseT8F7ioUkPCAPsResponseFullSystem string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseFullSystemMagicTransit PCAPNewResponseT8F7ioUkPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseT8F7ioUkPCAPsResponseFullType string

const (
	PCAPNewResponseT8F7ioUkPCAPsResponseFullTypeSimple PCAPNewResponseT8F7ioUkPCAPsResponseFullType = "simple"
	PCAPNewResponseT8F7ioUkPCAPsResponseFullTypeFull   PCAPNewResponseT8F7ioUkPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseT8F7ioUkPCAPsResponseSimple] or
// [PCAPListResponseT8F7ioUkPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseT8F7ioUkPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseT8F7ioUkPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseT8F7ioUkPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseT8F7ioUkPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseT8F7ioUkPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseT8F7ioUkPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseT8F7ioUkPCAPsResponseSimple]
type pcapListResponseT8F7ioUkPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseT8F7ioUkPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseT8F7ioUkPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseT8F7ioUkPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseT8F7ioUkPCAPsResponseSimpleFilterV1]
type pcapListResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseT8F7ioUkPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusUnknown           PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusSuccess           PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "success"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusPending           PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusRunning           PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "running"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusConversionPending PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusConversionRunning PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusComplete          PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatusFailed            PCAPListResponseT8F7ioUkPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseT8F7ioUkPCAPsResponseSimpleSystem string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleSystemMagicTransit PCAPListResponseT8F7ioUkPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseT8F7ioUkPCAPsResponseSimpleType string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleTypeSimple PCAPListResponseT8F7ioUkPCAPsResponseSimpleType = "simple"
	PCAPListResponseT8F7ioUkPCAPsResponseSimpleTypeFull   PCAPListResponseT8F7ioUkPCAPsResponseSimpleType = "full"
)

type PCAPListResponseT8F7ioUkPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseT8F7ioUkPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseT8F7ioUkPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseT8F7ioUkPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseT8F7ioUkPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseT8F7ioUkPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseT8F7ioUkPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseT8F7ioUkPCAPsResponseFull]
type pcapListResponseT8F7ioUkPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseT8F7ioUkPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseT8F7ioUkPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseT8F7ioUkPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseT8F7ioUkPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseT8F7ioUkPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseT8F7ioUkPCAPsResponseFullFilterV1]
type pcapListResponseT8F7ioUkPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseT8F7ioUkPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseT8F7ioUkPCAPsResponseFullStatus string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusUnknown           PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "unknown"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusSuccess           PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "success"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusPending           PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "pending"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusRunning           PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "running"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusConversionPending PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusConversionRunning PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusComplete          PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "complete"
	PCAPListResponseT8F7ioUkPCAPsResponseFullStatusFailed            PCAPListResponseT8F7ioUkPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseT8F7ioUkPCAPsResponseFullSystem string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseFullSystemMagicTransit PCAPListResponseT8F7ioUkPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseT8F7ioUkPCAPsResponseFullType string

const (
	PCAPListResponseT8F7ioUkPCAPsResponseFullTypeSimple PCAPListResponseT8F7ioUkPCAPsResponseFullType = "simple"
	PCAPListResponseT8F7ioUkPCAPsResponseFullTypeFull   PCAPListResponseT8F7ioUkPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseT8F7ioUkPCAPsResponseSimple] or
// [PCAPGetResponseT8F7ioUkPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseT8F7ioUkPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseT8F7ioUkPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseT8F7ioUkPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseT8F7ioUkPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseT8F7ioUkPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseT8F7ioUkPCAPsResponseSimple]
type pcapGetResponseT8F7ioUkPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseT8F7ioUkPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseT8F7ioUkPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1]
type pcapGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseT8F7ioUkPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusUnknown           PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusSuccess           PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusPending           PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusRunning           PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusConversionPending PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusComplete          PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatusFailed            PCAPGetResponseT8F7ioUkPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseT8F7ioUkPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseT8F7ioUkPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseT8F7ioUkPCAPsResponseSimpleType string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleTypeSimple PCAPGetResponseT8F7ioUkPCAPsResponseSimpleType = "simple"
	PCAPGetResponseT8F7ioUkPCAPsResponseSimpleTypeFull   PCAPGetResponseT8F7ioUkPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseT8F7ioUkPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseT8F7ioUkPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseT8F7ioUkPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseT8F7ioUkPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseT8F7ioUkPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseT8F7ioUkPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseT8F7ioUkPCAPsResponseFull]
type pcapGetResponseT8F7ioUkPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseT8F7ioUkPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseT8F7ioUkPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseT8F7ioUkPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseT8F7ioUkPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseT8F7ioUkPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseT8F7ioUkPCAPsResponseFullFilterV1]
type pcapGetResponseT8F7ioUkPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseT8F7ioUkPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusUnknown           PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusSuccess           PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "success"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusPending           PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "pending"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusRunning           PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "running"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusConversionPending PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusConversionRunning PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusComplete          PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "complete"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullStatusFailed            PCAPGetResponseT8F7ioUkPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseT8F7ioUkPCAPsResponseFullSystem string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseFullSystemMagicTransit PCAPGetResponseT8F7ioUkPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseT8F7ioUkPCAPsResponseFullType string

const (
	PCAPGetResponseT8F7ioUkPCAPsResponseFullTypeSimple PCAPGetResponseT8F7ioUkPCAPsResponseFullType = "simple"
	PCAPGetResponseT8F7ioUkPCAPsResponseFullTypeFull   PCAPGetResponseT8F7ioUkPCAPsResponseFullType = "full"
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
