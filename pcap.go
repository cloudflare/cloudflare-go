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

// Union satisfied by [PCAPNewResponseX60QcBlGPCAPsResponseSimple] or
// [PCAPNewResponseX60QcBlGPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseX60QcBlGPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseX60QcBlGPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseX60QcBlGPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseX60QcBlGPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseX60QcBlGpcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseX60QcBlGpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseX60QcBlGPCAPsResponseSimple]
type pcapNewResponseX60QcBlGpcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseX60QcBlGPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseX60QcBlGPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseX60QcBlGPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseX60QcBlGPCAPsResponseSimpleFilterV1]
type pcapNewResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseX60QcBlGPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusUnknown           PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusSuccess           PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusPending           PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusRunning           PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusConversionPending PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusComplete          PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatusFailed            PCAPNewResponseX60QcBlGPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseX60QcBlGPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseX60QcBlGPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseX60QcBlGPCAPsResponseSimpleType string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleTypeSimple PCAPNewResponseX60QcBlGPCAPsResponseSimpleType = "simple"
	PCAPNewResponseX60QcBlGPCAPsResponseSimpleTypeFull   PCAPNewResponseX60QcBlGPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseX60QcBlGPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseX60QcBlGPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseX60QcBlGPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseX60QcBlGPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseX60QcBlGPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseX60QcBlGpcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponseX60QcBlGpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseX60QcBlGPCAPsResponseFull]
type pcapNewResponseX60QcBlGpcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponseX60QcBlGPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseX60QcBlGPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseX60QcBlGPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseX60QcBlGpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseX60QcBlGpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseX60QcBlGPCAPsResponseFullFilterV1]
type pcapNewResponseX60QcBlGpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseX60QcBlGPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseX60QcBlGPCAPsResponseFullStatus string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusUnknown           PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusSuccess           PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "success"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusPending           PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "pending"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusRunning           PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "running"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusConversionPending PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusConversionRunning PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusComplete          PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "complete"
	PCAPNewResponseX60QcBlGPCAPsResponseFullStatusFailed            PCAPNewResponseX60QcBlGPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseX60QcBlGPCAPsResponseFullSystem string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseFullSystemMagicTransit PCAPNewResponseX60QcBlGPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseX60QcBlGPCAPsResponseFullType string

const (
	PCAPNewResponseX60QcBlGPCAPsResponseFullTypeSimple PCAPNewResponseX60QcBlGPCAPsResponseFullType = "simple"
	PCAPNewResponseX60QcBlGPCAPsResponseFullTypeFull   PCAPNewResponseX60QcBlGPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseX60QcBlGPCAPsResponseSimple] or
// [PCAPListResponseX60QcBlGPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseX60QcBlGPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseX60QcBlGPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseX60QcBlGPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseX60QcBlGPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseX60QcBlGpcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseX60QcBlGpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseX60QcBlGPCAPsResponseSimple]
type pcapListResponseX60QcBlGpcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseX60QcBlGPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseX60QcBlGPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseX60QcBlGPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseX60QcBlGPCAPsResponseSimpleFilterV1]
type pcapListResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseX60QcBlGPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus string

const (
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusUnknown           PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusSuccess           PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "success"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusPending           PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusRunning           PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "running"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusConversionPending PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusConversionRunning PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusComplete          PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleStatusFailed            PCAPListResponseX60QcBlGPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseX60QcBlGPCAPsResponseSimpleSystem string

const (
	PCAPListResponseX60QcBlGPCAPsResponseSimpleSystemMagicTransit PCAPListResponseX60QcBlGPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseX60QcBlGPCAPsResponseSimpleType string

const (
	PCAPListResponseX60QcBlGPCAPsResponseSimpleTypeSimple PCAPListResponseX60QcBlGPCAPsResponseSimpleType = "simple"
	PCAPListResponseX60QcBlGPCAPsResponseSimpleTypeFull   PCAPListResponseX60QcBlGPCAPsResponseSimpleType = "full"
)

type PCAPListResponseX60QcBlGPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseX60QcBlGPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseX60QcBlGPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseX60QcBlGPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseX60QcBlGPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseX60QcBlGpcaPsResponseFullJSON `json:"-"`
}

// pcapListResponseX60QcBlGpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseX60QcBlGPCAPsResponseFull]
type pcapListResponseX60QcBlGpcaPsResponseFullJSON struct {
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

func (r *PCAPListResponseX60QcBlGPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseX60QcBlGPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseX60QcBlGPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseX60QcBlGpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseX60QcBlGpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseX60QcBlGPCAPsResponseFullFilterV1]
type pcapListResponseX60QcBlGpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseX60QcBlGPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseX60QcBlGPCAPsResponseFullStatus string

const (
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusUnknown           PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "unknown"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusSuccess           PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "success"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusPending           PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "pending"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusRunning           PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "running"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusConversionPending PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusConversionRunning PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusComplete          PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "complete"
	PCAPListResponseX60QcBlGPCAPsResponseFullStatusFailed            PCAPListResponseX60QcBlGPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseX60QcBlGPCAPsResponseFullSystem string

const (
	PCAPListResponseX60QcBlGPCAPsResponseFullSystemMagicTransit PCAPListResponseX60QcBlGPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseX60QcBlGPCAPsResponseFullType string

const (
	PCAPListResponseX60QcBlGPCAPsResponseFullTypeSimple PCAPListResponseX60QcBlGPCAPsResponseFullType = "simple"
	PCAPListResponseX60QcBlGPCAPsResponseFullTypeFull   PCAPListResponseX60QcBlGPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseX60QcBlGPCAPsResponseSimple] or
// [PCAPGetResponseX60QcBlGPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseX60QcBlGPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseX60QcBlGPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseX60QcBlGPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseX60QcBlGPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseX60QcBlGpcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseX60QcBlGpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseX60QcBlGPCAPsResponseSimple]
type pcapGetResponseX60QcBlGpcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseX60QcBlGPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseX60QcBlGPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseX60QcBlGPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseX60QcBlGPCAPsResponseSimpleFilterV1]
type pcapGetResponseX60QcBlGpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseX60QcBlGPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusUnknown           PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusSuccess           PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusPending           PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusRunning           PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusConversionPending PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusComplete          PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatusFailed            PCAPGetResponseX60QcBlGPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseX60QcBlGPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseX60QcBlGPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseX60QcBlGPCAPsResponseSimpleType string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleTypeSimple PCAPGetResponseX60QcBlGPCAPsResponseSimpleType = "simple"
	PCAPGetResponseX60QcBlGPCAPsResponseSimpleTypeFull   PCAPGetResponseX60QcBlGPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseX60QcBlGPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseX60QcBlGPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseX60QcBlGPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseX60QcBlGPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseX60QcBlGPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseX60QcBlGpcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponseX60QcBlGpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseX60QcBlGPCAPsResponseFull]
type pcapGetResponseX60QcBlGpcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponseX60QcBlGPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseX60QcBlGPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseX60QcBlGPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseX60QcBlGpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseX60QcBlGpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseX60QcBlGPCAPsResponseFullFilterV1]
type pcapGetResponseX60QcBlGpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseX60QcBlGPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseX60QcBlGPCAPsResponseFullStatus string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusUnknown           PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusSuccess           PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "success"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusPending           PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "pending"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusRunning           PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "running"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusConversionPending PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusConversionRunning PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusComplete          PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "complete"
	PCAPGetResponseX60QcBlGPCAPsResponseFullStatusFailed            PCAPGetResponseX60QcBlGPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseX60QcBlGPCAPsResponseFullSystem string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseFullSystemMagicTransit PCAPGetResponseX60QcBlGPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseX60QcBlGPCAPsResponseFullType string

const (
	PCAPGetResponseX60QcBlGPCAPsResponseFullTypeSimple PCAPGetResponseX60QcBlGPCAPsResponseFullType = "simple"
	PCAPGetResponseX60QcBlGPCAPsResponseFullTypeFull   PCAPGetResponseX60QcBlGPCAPsResponseFullType = "full"
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
