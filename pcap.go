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

// Union satisfied by [PCAPNewResponseOkFeXgZfPCAPsResponseSimple] or
// [PCAPNewResponseOkFeXgZfPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseOkFeXgZfPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseOkFeXgZfPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseOkFeXgZfPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseOkFeXgZfPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseOkFeXgZfPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseOkFeXgZfPCAPsResponseSimple]
type pcapNewResponseOkFeXgZfPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseOkFeXgZfPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseOkFeXgZfPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1]
type pcapNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseOkFeXgZfPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusUnknown           PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusSuccess           PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusPending           PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusRunning           PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusConversionPending PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusComplete          PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatusFailed            PCAPNewResponseOkFeXgZfPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseOkFeXgZfPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseOkFeXgZfPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseOkFeXgZfPCAPsResponseSimpleType string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleTypeSimple PCAPNewResponseOkFeXgZfPCAPsResponseSimpleType = "simple"
	PCAPNewResponseOkFeXgZfPCAPsResponseSimpleTypeFull   PCAPNewResponseOkFeXgZfPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseOkFeXgZfPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseOkFeXgZfPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseOkFeXgZfPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseOkFeXgZfPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseOkFeXgZfPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseOkFeXgZfPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseOkFeXgZfPCAPsResponseFull]
type pcapNewResponseOkFeXgZfPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseOkFeXgZfPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseOkFeXgZfPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseOkFeXgZfPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseOkFeXgZfPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseOkFeXgZfPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseOkFeXgZfPCAPsResponseFullFilterV1]
type pcapNewResponseOkFeXgZfPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseOkFeXgZfPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusUnknown           PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusSuccess           PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "success"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusPending           PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "pending"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusRunning           PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "running"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusConversionPending PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusConversionRunning PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusComplete          PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "complete"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullStatusFailed            PCAPNewResponseOkFeXgZfPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseOkFeXgZfPCAPsResponseFullSystem string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseFullSystemMagicTransit PCAPNewResponseOkFeXgZfPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseOkFeXgZfPCAPsResponseFullType string

const (
	PCAPNewResponseOkFeXgZfPCAPsResponseFullTypeSimple PCAPNewResponseOkFeXgZfPCAPsResponseFullType = "simple"
	PCAPNewResponseOkFeXgZfPCAPsResponseFullTypeFull   PCAPNewResponseOkFeXgZfPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseOkFeXgZfPCAPsResponseSimple] or
// [PCAPListResponseOkFeXgZfPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseOkFeXgZfPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseOkFeXgZfPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseOkFeXgZfPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseOkFeXgZfPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseOkFeXgZfPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseOkFeXgZfPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseOkFeXgZfPCAPsResponseSimple]
type pcapListResponseOkFeXgZfPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseOkFeXgZfPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseOkFeXgZfPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseOkFeXgZfPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseOkFeXgZfPCAPsResponseSimpleFilterV1]
type pcapListResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseOkFeXgZfPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusUnknown           PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusSuccess           PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "success"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusPending           PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusRunning           PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "running"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusConversionPending PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusConversionRunning PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusComplete          PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatusFailed            PCAPListResponseOkFeXgZfPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseOkFeXgZfPCAPsResponseSimpleSystem string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleSystemMagicTransit PCAPListResponseOkFeXgZfPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseOkFeXgZfPCAPsResponseSimpleType string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleTypeSimple PCAPListResponseOkFeXgZfPCAPsResponseSimpleType = "simple"
	PCAPListResponseOkFeXgZfPCAPsResponseSimpleTypeFull   PCAPListResponseOkFeXgZfPCAPsResponseSimpleType = "full"
)

type PCAPListResponseOkFeXgZfPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseOkFeXgZfPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseOkFeXgZfPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseOkFeXgZfPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseOkFeXgZfPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseOkFeXgZfPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseOkFeXgZfPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseOkFeXgZfPCAPsResponseFull]
type pcapListResponseOkFeXgZfPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseOkFeXgZfPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseOkFeXgZfPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseOkFeXgZfPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseOkFeXgZfPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseOkFeXgZfPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseOkFeXgZfPCAPsResponseFullFilterV1]
type pcapListResponseOkFeXgZfPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseOkFeXgZfPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseOkFeXgZfPCAPsResponseFullStatus string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusUnknown           PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "unknown"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusSuccess           PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "success"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusPending           PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "pending"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusRunning           PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "running"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusConversionPending PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusConversionRunning PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusComplete          PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "complete"
	PCAPListResponseOkFeXgZfPCAPsResponseFullStatusFailed            PCAPListResponseOkFeXgZfPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseOkFeXgZfPCAPsResponseFullSystem string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseFullSystemMagicTransit PCAPListResponseOkFeXgZfPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseOkFeXgZfPCAPsResponseFullType string

const (
	PCAPListResponseOkFeXgZfPCAPsResponseFullTypeSimple PCAPListResponseOkFeXgZfPCAPsResponseFullType = "simple"
	PCAPListResponseOkFeXgZfPCAPsResponseFullTypeFull   PCAPListResponseOkFeXgZfPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseOkFeXgZfPCAPsResponseSimple] or
// [PCAPGetResponseOkFeXgZfPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseOkFeXgZfPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseOkFeXgZfPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseOkFeXgZfPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseOkFeXgZfPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseOkFeXgZfPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseOkFeXgZfPCAPsResponseSimple]
type pcapGetResponseOkFeXgZfPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseOkFeXgZfPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseOkFeXgZfPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1]
type pcapGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseOkFeXgZfPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusUnknown           PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusSuccess           PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusPending           PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusRunning           PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusConversionPending PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusComplete          PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatusFailed            PCAPGetResponseOkFeXgZfPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseOkFeXgZfPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseOkFeXgZfPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseOkFeXgZfPCAPsResponseSimpleType string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleTypeSimple PCAPGetResponseOkFeXgZfPCAPsResponseSimpleType = "simple"
	PCAPGetResponseOkFeXgZfPCAPsResponseSimpleTypeFull   PCAPGetResponseOkFeXgZfPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseOkFeXgZfPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseOkFeXgZfPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseOkFeXgZfPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseOkFeXgZfPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseOkFeXgZfPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseOkFeXgZfPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseOkFeXgZfPCAPsResponseFull]
type pcapGetResponseOkFeXgZfPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseOkFeXgZfPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseOkFeXgZfPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseOkFeXgZfPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseOkFeXgZfPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseOkFeXgZfPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseOkFeXgZfPCAPsResponseFullFilterV1]
type pcapGetResponseOkFeXgZfPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseOkFeXgZfPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusUnknown           PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusSuccess           PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "success"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusPending           PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "pending"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusRunning           PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "running"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusConversionPending PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusConversionRunning PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusComplete          PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "complete"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullStatusFailed            PCAPGetResponseOkFeXgZfPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseOkFeXgZfPCAPsResponseFullSystem string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseFullSystemMagicTransit PCAPGetResponseOkFeXgZfPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseOkFeXgZfPCAPsResponseFullType string

const (
	PCAPGetResponseOkFeXgZfPCAPsResponseFullTypeSimple PCAPGetResponseOkFeXgZfPCAPsResponseFullType = "simple"
	PCAPGetResponseOkFeXgZfPCAPsResponseFullTypeFull   PCAPGetResponseOkFeXgZfPCAPsResponseFullType = "full"
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
