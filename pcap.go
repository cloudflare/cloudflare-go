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

// Union satisfied by [PCAPNewResponseSochEXmAPCAPsResponseSimple] or
// [PCAPNewResponseSochEXmAPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseSochEXmAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseSochEXmAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseSochEXmAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseSochEXmAPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseSochEXmApcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseSochEXmApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseSochEXmAPCAPsResponseSimple]
type pcapNewResponseSochEXmApcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseSochEXmAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseSochEXmAPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseSochEXmAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseSochEXmApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseSochEXmApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseSochEXmAPCAPsResponseSimpleFilterV1]
type pcapNewResponseSochEXmApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseSochEXmAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusUnknown           PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusSuccess           PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusPending           PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusRunning           PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusConversionPending PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusComplete          PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleStatusFailed            PCAPNewResponseSochEXmAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseSochEXmAPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseSochEXmAPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseSochEXmAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseSochEXmAPCAPsResponseSimpleType string

const (
	PCAPNewResponseSochEXmAPCAPsResponseSimpleTypeSimple PCAPNewResponseSochEXmAPCAPsResponseSimpleType = "simple"
	PCAPNewResponseSochEXmAPCAPsResponseSimpleTypeFull   PCAPNewResponseSochEXmAPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseSochEXmAPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseSochEXmAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseSochEXmAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseSochEXmAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseSochEXmAPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseSochEXmApcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponseSochEXmApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseSochEXmAPCAPsResponseFull]
type pcapNewResponseSochEXmApcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponseSochEXmAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseSochEXmAPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseSochEXmAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseSochEXmApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseSochEXmApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseSochEXmAPCAPsResponseFullFilterV1]
type pcapNewResponseSochEXmApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseSochEXmAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseSochEXmAPCAPsResponseFullStatus string

const (
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusUnknown           PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusSuccess           PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "success"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusPending           PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "pending"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusRunning           PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "running"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusConversionPending PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusConversionRunning PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusComplete          PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "complete"
	PCAPNewResponseSochEXmAPCAPsResponseFullStatusFailed            PCAPNewResponseSochEXmAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseSochEXmAPCAPsResponseFullSystem string

const (
	PCAPNewResponseSochEXmAPCAPsResponseFullSystemMagicTransit PCAPNewResponseSochEXmAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseSochEXmAPCAPsResponseFullType string

const (
	PCAPNewResponseSochEXmAPCAPsResponseFullTypeSimple PCAPNewResponseSochEXmAPCAPsResponseFullType = "simple"
	PCAPNewResponseSochEXmAPCAPsResponseFullTypeFull   PCAPNewResponseSochEXmAPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseSochEXmAPCAPsResponseSimple] or
// [PCAPListResponseSochEXmAPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseSochEXmAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseSochEXmAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseSochEXmAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseSochEXmAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseSochEXmAPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseSochEXmApcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseSochEXmApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseSochEXmAPCAPsResponseSimple]
type pcapListResponseSochEXmApcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseSochEXmAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseSochEXmAPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseSochEXmAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseSochEXmApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseSochEXmApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseSochEXmAPCAPsResponseSimpleFilterV1]
type pcapListResponseSochEXmApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseSochEXmAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseSochEXmAPCAPsResponseSimpleStatus string

const (
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusUnknown           PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusSuccess           PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "success"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusPending           PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusRunning           PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "running"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusConversionPending PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusConversionRunning PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusComplete          PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseSochEXmAPCAPsResponseSimpleStatusFailed            PCAPListResponseSochEXmAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseSochEXmAPCAPsResponseSimpleSystem string

const (
	PCAPListResponseSochEXmAPCAPsResponseSimpleSystemMagicTransit PCAPListResponseSochEXmAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseSochEXmAPCAPsResponseSimpleType string

const (
	PCAPListResponseSochEXmAPCAPsResponseSimpleTypeSimple PCAPListResponseSochEXmAPCAPsResponseSimpleType = "simple"
	PCAPListResponseSochEXmAPCAPsResponseSimpleTypeFull   PCAPListResponseSochEXmAPCAPsResponseSimpleType = "full"
)

type PCAPListResponseSochEXmAPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseSochEXmAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseSochEXmAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseSochEXmAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseSochEXmAPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseSochEXmApcaPsResponseFullJSON `json:"-"`
}

// pcapListResponseSochEXmApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseSochEXmAPCAPsResponseFull]
type pcapListResponseSochEXmApcaPsResponseFullJSON struct {
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

func (r *PCAPListResponseSochEXmAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseSochEXmAPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseSochEXmAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseSochEXmApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseSochEXmApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseSochEXmAPCAPsResponseFullFilterV1]
type pcapListResponseSochEXmApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseSochEXmAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseSochEXmAPCAPsResponseFullStatus string

const (
	PCAPListResponseSochEXmAPCAPsResponseFullStatusUnknown           PCAPListResponseSochEXmAPCAPsResponseFullStatus = "unknown"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusSuccess           PCAPListResponseSochEXmAPCAPsResponseFullStatus = "success"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusPending           PCAPListResponseSochEXmAPCAPsResponseFullStatus = "pending"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusRunning           PCAPListResponseSochEXmAPCAPsResponseFullStatus = "running"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusConversionPending PCAPListResponseSochEXmAPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusConversionRunning PCAPListResponseSochEXmAPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusComplete          PCAPListResponseSochEXmAPCAPsResponseFullStatus = "complete"
	PCAPListResponseSochEXmAPCAPsResponseFullStatusFailed            PCAPListResponseSochEXmAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseSochEXmAPCAPsResponseFullSystem string

const (
	PCAPListResponseSochEXmAPCAPsResponseFullSystemMagicTransit PCAPListResponseSochEXmAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseSochEXmAPCAPsResponseFullType string

const (
	PCAPListResponseSochEXmAPCAPsResponseFullTypeSimple PCAPListResponseSochEXmAPCAPsResponseFullType = "simple"
	PCAPListResponseSochEXmAPCAPsResponseFullTypeFull   PCAPListResponseSochEXmAPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseSochEXmAPCAPsResponseSimple] or
// [PCAPGetResponseSochEXmAPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseSochEXmAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseSochEXmAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseSochEXmAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseSochEXmAPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseSochEXmApcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseSochEXmApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseSochEXmAPCAPsResponseSimple]
type pcapGetResponseSochEXmApcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseSochEXmAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseSochEXmAPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseSochEXmAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseSochEXmApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseSochEXmApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseSochEXmAPCAPsResponseSimpleFilterV1]
type pcapGetResponseSochEXmApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseSochEXmAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusUnknown           PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusSuccess           PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusPending           PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusRunning           PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusConversionPending PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusComplete          PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleStatusFailed            PCAPGetResponseSochEXmAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseSochEXmAPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseSochEXmAPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseSochEXmAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseSochEXmAPCAPsResponseSimpleType string

const (
	PCAPGetResponseSochEXmAPCAPsResponseSimpleTypeSimple PCAPGetResponseSochEXmAPCAPsResponseSimpleType = "simple"
	PCAPGetResponseSochEXmAPCAPsResponseSimpleTypeFull   PCAPGetResponseSochEXmAPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseSochEXmAPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseSochEXmAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseSochEXmAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseSochEXmAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseSochEXmAPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseSochEXmApcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponseSochEXmApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseSochEXmAPCAPsResponseFull]
type pcapGetResponseSochEXmApcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponseSochEXmAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseSochEXmAPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseSochEXmAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseSochEXmApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseSochEXmApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseSochEXmAPCAPsResponseFullFilterV1]
type pcapGetResponseSochEXmApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseSochEXmAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseSochEXmAPCAPsResponseFullStatus string

const (
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusUnknown           PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusSuccess           PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "success"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusPending           PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "pending"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusRunning           PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "running"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusConversionPending PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusConversionRunning PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusComplete          PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "complete"
	PCAPGetResponseSochEXmAPCAPsResponseFullStatusFailed            PCAPGetResponseSochEXmAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseSochEXmAPCAPsResponseFullSystem string

const (
	PCAPGetResponseSochEXmAPCAPsResponseFullSystemMagicTransit PCAPGetResponseSochEXmAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseSochEXmAPCAPsResponseFullType string

const (
	PCAPGetResponseSochEXmAPCAPsResponseFullTypeSimple PCAPGetResponseSochEXmAPCAPsResponseFullType = "simple"
	PCAPGetResponseSochEXmAPCAPsResponseFullTypeFull   PCAPGetResponseSochEXmAPCAPsResponseFullType = "full"
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
