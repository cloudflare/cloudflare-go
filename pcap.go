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

// Union satisfied by [PCAPNewResponseXti6IAgpPCAPsResponseSimple] or
// [PCAPNewResponseXti6IAgpPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseXti6IAgpPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseXti6IAgpPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseXti6IAgpPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseXti6IAgpPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseXti6IAgpPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseXti6IAgpPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseXti6IAgpPCAPsResponseSimple]
type pcapNewResponseXti6IAgpPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseXti6IAgpPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseXti6IAgpPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseXti6IAgpPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseXti6IAgpPCAPsResponseSimpleFilterV1]
type pcapNewResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseXti6IAgpPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusUnknown           PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusSuccess           PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusPending           PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusRunning           PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusConversionPending PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusComplete          PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatusFailed            PCAPNewResponseXti6IAgpPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseXti6IAgpPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseXti6IAgpPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseXti6IAgpPCAPsResponseSimpleType string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleTypeSimple PCAPNewResponseXti6IAgpPCAPsResponseSimpleType = "simple"
	PCAPNewResponseXti6IAgpPCAPsResponseSimpleTypeFull   PCAPNewResponseXti6IAgpPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseXti6IAgpPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseXti6IAgpPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseXti6IAgpPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseXti6IAgpPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseXti6IAgpPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseXti6IAgpPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseXti6IAgpPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseXti6IAgpPCAPsResponseFull]
type pcapNewResponseXti6IAgpPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseXti6IAgpPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseXti6IAgpPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseXti6IAgpPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseXti6IAgpPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseXti6IAgpPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseXti6IAgpPCAPsResponseFullFilterV1]
type pcapNewResponseXti6IAgpPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseXti6IAgpPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseXti6IAgpPCAPsResponseFullStatus string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusUnknown           PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusSuccess           PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "success"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusPending           PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "pending"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusRunning           PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "running"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusConversionPending PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusConversionRunning PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusComplete          PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "complete"
	PCAPNewResponseXti6IAgpPCAPsResponseFullStatusFailed            PCAPNewResponseXti6IAgpPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseXti6IAgpPCAPsResponseFullSystem string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseFullSystemMagicTransit PCAPNewResponseXti6IAgpPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseXti6IAgpPCAPsResponseFullType string

const (
	PCAPNewResponseXti6IAgpPCAPsResponseFullTypeSimple PCAPNewResponseXti6IAgpPCAPsResponseFullType = "simple"
	PCAPNewResponseXti6IAgpPCAPsResponseFullTypeFull   PCAPNewResponseXti6IAgpPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseXti6IAgpPCAPsResponseSimple] or
// [PCAPListResponseXti6IAgpPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseXti6IAgpPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseXti6IAgpPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseXti6IAgpPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseXti6IAgpPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseXti6IAgpPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseXti6IAgpPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseXti6IAgpPCAPsResponseSimple]
type pcapListResponseXti6IAgpPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseXti6IAgpPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseXti6IAgpPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseXti6IAgpPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseXti6IAgpPCAPsResponseSimpleFilterV1]
type pcapListResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseXti6IAgpPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus string

const (
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusUnknown           PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusSuccess           PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "success"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusPending           PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusRunning           PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "running"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusConversionPending PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusConversionRunning PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusComplete          PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleStatusFailed            PCAPListResponseXti6IAgpPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseXti6IAgpPCAPsResponseSimpleSystem string

const (
	PCAPListResponseXti6IAgpPCAPsResponseSimpleSystemMagicTransit PCAPListResponseXti6IAgpPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseXti6IAgpPCAPsResponseSimpleType string

const (
	PCAPListResponseXti6IAgpPCAPsResponseSimpleTypeSimple PCAPListResponseXti6IAgpPCAPsResponseSimpleType = "simple"
	PCAPListResponseXti6IAgpPCAPsResponseSimpleTypeFull   PCAPListResponseXti6IAgpPCAPsResponseSimpleType = "full"
)

type PCAPListResponseXti6IAgpPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseXti6IAgpPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseXti6IAgpPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseXti6IAgpPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseXti6IAgpPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseXti6IAgpPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseXti6IAgpPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseXti6IAgpPCAPsResponseFull]
type pcapListResponseXti6IAgpPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseXti6IAgpPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseXti6IAgpPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseXti6IAgpPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseXti6IAgpPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseXti6IAgpPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseXti6IAgpPCAPsResponseFullFilterV1]
type pcapListResponseXti6IAgpPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseXti6IAgpPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseXti6IAgpPCAPsResponseFullStatus string

const (
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusUnknown           PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "unknown"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusSuccess           PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "success"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusPending           PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "pending"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusRunning           PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "running"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusConversionPending PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusConversionRunning PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusComplete          PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "complete"
	PCAPListResponseXti6IAgpPCAPsResponseFullStatusFailed            PCAPListResponseXti6IAgpPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseXti6IAgpPCAPsResponseFullSystem string

const (
	PCAPListResponseXti6IAgpPCAPsResponseFullSystemMagicTransit PCAPListResponseXti6IAgpPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseXti6IAgpPCAPsResponseFullType string

const (
	PCAPListResponseXti6IAgpPCAPsResponseFullTypeSimple PCAPListResponseXti6IAgpPCAPsResponseFullType = "simple"
	PCAPListResponseXti6IAgpPCAPsResponseFullTypeFull   PCAPListResponseXti6IAgpPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseXti6IAgpPCAPsResponseSimple] or
// [PCAPGetResponseXti6IAgpPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseXti6IAgpPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseXti6IAgpPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseXti6IAgpPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseXti6IAgpPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseXti6IAgpPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseXti6IAgpPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseXti6IAgpPCAPsResponseSimple]
type pcapGetResponseXti6IAgpPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseXti6IAgpPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseXti6IAgpPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseXti6IAgpPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseXti6IAgpPCAPsResponseSimpleFilterV1]
type pcapGetResponseXti6IAgpPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseXti6IAgpPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusUnknown           PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusSuccess           PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusPending           PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusRunning           PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusConversionPending PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusComplete          PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatusFailed            PCAPGetResponseXti6IAgpPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseXti6IAgpPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseXti6IAgpPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseXti6IAgpPCAPsResponseSimpleType string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleTypeSimple PCAPGetResponseXti6IAgpPCAPsResponseSimpleType = "simple"
	PCAPGetResponseXti6IAgpPCAPsResponseSimpleTypeFull   PCAPGetResponseXti6IAgpPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseXti6IAgpPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseXti6IAgpPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseXti6IAgpPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseXti6IAgpPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseXti6IAgpPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseXti6IAgpPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseXti6IAgpPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseXti6IAgpPCAPsResponseFull]
type pcapGetResponseXti6IAgpPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseXti6IAgpPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseXti6IAgpPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseXti6IAgpPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseXti6IAgpPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseXti6IAgpPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseXti6IAgpPCAPsResponseFullFilterV1]
type pcapGetResponseXti6IAgpPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseXti6IAgpPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseXti6IAgpPCAPsResponseFullStatus string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusUnknown           PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusSuccess           PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "success"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusPending           PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "pending"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusRunning           PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "running"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusConversionPending PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusConversionRunning PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusComplete          PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "complete"
	PCAPGetResponseXti6IAgpPCAPsResponseFullStatusFailed            PCAPGetResponseXti6IAgpPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseXti6IAgpPCAPsResponseFullSystem string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseFullSystemMagicTransit PCAPGetResponseXti6IAgpPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseXti6IAgpPCAPsResponseFullType string

const (
	PCAPGetResponseXti6IAgpPCAPsResponseFullTypeSimple PCAPGetResponseXti6IAgpPCAPsResponseFullType = "simple"
	PCAPGetResponseXti6IAgpPCAPsResponseFullTypeFull   PCAPGetResponseXti6IAgpPCAPsResponseFullType = "full"
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
