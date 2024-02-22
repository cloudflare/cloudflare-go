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

// Union satisfied by [PCAPNewResponseDqIPKuAyPCAPsResponseSimple] or
// [PCAPNewResponseDqIPKuAyPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseDqIPKuAyPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseDqIPKuAyPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseDqIPKuAyPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseDqIPKuAyPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseDqIPKuAyPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseDqIPKuAyPCAPsResponseSimple]
type pcapNewResponseDqIPKuAyPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseDqIPKuAyPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseDqIPKuAyPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1]
type pcapNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseDqIPKuAyPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusUnknown           PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusSuccess           PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusPending           PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusRunning           PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusConversionPending PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusComplete          PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatusFailed            PCAPNewResponseDqIPKuAyPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseDqIPKuAyPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseDqIPKuAyPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseDqIPKuAyPCAPsResponseSimpleType string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleTypeSimple PCAPNewResponseDqIPKuAyPCAPsResponseSimpleType = "simple"
	PCAPNewResponseDqIPKuAyPCAPsResponseSimpleTypeFull   PCAPNewResponseDqIPKuAyPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseDqIPKuAyPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseDqIPKuAyPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseDqIPKuAyPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseDqIPKuAyPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseDqIPKuAyPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseDqIPKuAyPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseDqIPKuAyPCAPsResponseFull]
type pcapNewResponseDqIPKuAyPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseDqIPKuAyPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseDqIPKuAyPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseDqIPKuAyPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseDqIPKuAyPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseDqIPKuAyPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseDqIPKuAyPCAPsResponseFullFilterV1]
type pcapNewResponseDqIPKuAyPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseDqIPKuAyPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusUnknown           PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusSuccess           PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "success"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusPending           PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "pending"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusRunning           PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "running"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusConversionPending PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusConversionRunning PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusComplete          PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "complete"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullStatusFailed            PCAPNewResponseDqIPKuAyPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseDqIPKuAyPCAPsResponseFullSystem string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseFullSystemMagicTransit PCAPNewResponseDqIPKuAyPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseDqIPKuAyPCAPsResponseFullType string

const (
	PCAPNewResponseDqIPKuAyPCAPsResponseFullTypeSimple PCAPNewResponseDqIPKuAyPCAPsResponseFullType = "simple"
	PCAPNewResponseDqIPKuAyPCAPsResponseFullTypeFull   PCAPNewResponseDqIPKuAyPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseDqIPKuAyPCAPsResponseSimple] or
// [PCAPListResponseDqIPKuAyPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseDqIPKuAyPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseDqIPKuAyPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseDqIPKuAyPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseDqIPKuAyPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseDqIPKuAyPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseDqIPKuAyPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseDqIPKuAyPCAPsResponseSimple]
type pcapListResponseDqIPKuAyPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseDqIPKuAyPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseDqIPKuAyPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseDqIPKuAyPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseDqIPKuAyPCAPsResponseSimpleFilterV1]
type pcapListResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseDqIPKuAyPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusUnknown           PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusSuccess           PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "success"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusPending           PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusRunning           PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "running"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusConversionPending PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusConversionRunning PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusComplete          PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatusFailed            PCAPListResponseDqIPKuAyPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseDqIPKuAyPCAPsResponseSimpleSystem string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleSystemMagicTransit PCAPListResponseDqIPKuAyPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseDqIPKuAyPCAPsResponseSimpleType string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleTypeSimple PCAPListResponseDqIPKuAyPCAPsResponseSimpleType = "simple"
	PCAPListResponseDqIPKuAyPCAPsResponseSimpleTypeFull   PCAPListResponseDqIPKuAyPCAPsResponseSimpleType = "full"
)

type PCAPListResponseDqIPKuAyPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseDqIPKuAyPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseDqIPKuAyPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseDqIPKuAyPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseDqIPKuAyPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseDqIPKuAyPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseDqIPKuAyPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseDqIPKuAyPCAPsResponseFull]
type pcapListResponseDqIPKuAyPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseDqIPKuAyPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseDqIPKuAyPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseDqIPKuAyPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseDqIPKuAyPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseDqIPKuAyPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseDqIPKuAyPCAPsResponseFullFilterV1]
type pcapListResponseDqIPKuAyPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseDqIPKuAyPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseDqIPKuAyPCAPsResponseFullStatus string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusUnknown           PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "unknown"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusSuccess           PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "success"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusPending           PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "pending"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusRunning           PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "running"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusConversionPending PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusConversionRunning PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusComplete          PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "complete"
	PCAPListResponseDqIPKuAyPCAPsResponseFullStatusFailed            PCAPListResponseDqIPKuAyPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseDqIPKuAyPCAPsResponseFullSystem string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseFullSystemMagicTransit PCAPListResponseDqIPKuAyPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseDqIPKuAyPCAPsResponseFullType string

const (
	PCAPListResponseDqIPKuAyPCAPsResponseFullTypeSimple PCAPListResponseDqIPKuAyPCAPsResponseFullType = "simple"
	PCAPListResponseDqIPKuAyPCAPsResponseFullTypeFull   PCAPListResponseDqIPKuAyPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseDqIPKuAyPCAPsResponseSimple] or
// [PCAPGetResponseDqIPKuAyPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseDqIPKuAyPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseDqIPKuAyPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseDqIPKuAyPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseDqIPKuAyPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseDqIPKuAyPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseDqIPKuAyPCAPsResponseSimple]
type pcapGetResponseDqIPKuAyPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseDqIPKuAyPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseDqIPKuAyPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1]
type pcapGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseDqIPKuAyPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusUnknown           PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusSuccess           PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusPending           PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusRunning           PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusConversionPending PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusComplete          PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatusFailed            PCAPGetResponseDqIPKuAyPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseDqIPKuAyPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseDqIPKuAyPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseDqIPKuAyPCAPsResponseSimpleType string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleTypeSimple PCAPGetResponseDqIPKuAyPCAPsResponseSimpleType = "simple"
	PCAPGetResponseDqIPKuAyPCAPsResponseSimpleTypeFull   PCAPGetResponseDqIPKuAyPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseDqIPKuAyPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseDqIPKuAyPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseDqIPKuAyPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseDqIPKuAyPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseDqIPKuAyPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseDqIPKuAyPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseDqIPKuAyPCAPsResponseFull]
type pcapGetResponseDqIPKuAyPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseDqIPKuAyPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseDqIPKuAyPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseDqIPKuAyPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseDqIPKuAyPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseDqIPKuAyPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseDqIPKuAyPCAPsResponseFullFilterV1]
type pcapGetResponseDqIPKuAyPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseDqIPKuAyPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusUnknown           PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusSuccess           PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "success"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusPending           PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "pending"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusRunning           PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "running"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusConversionPending PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusConversionRunning PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusComplete          PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "complete"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullStatusFailed            PCAPGetResponseDqIPKuAyPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseDqIPKuAyPCAPsResponseFullSystem string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseFullSystemMagicTransit PCAPGetResponseDqIPKuAyPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseDqIPKuAyPCAPsResponseFullType string

const (
	PCAPGetResponseDqIPKuAyPCAPsResponseFullTypeSimple PCAPGetResponseDqIPKuAyPCAPsResponseFullType = "simple"
	PCAPGetResponseDqIPKuAyPCAPsResponseFullTypeFull   PCAPGetResponseDqIPKuAyPCAPsResponseFullType = "full"
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
