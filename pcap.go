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

// Union satisfied by [PCAPNewResponseWuVsD7oCPCAPsResponseSimple] or
// [PCAPNewResponseWuVsD7oCPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseWuVsD7oCPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseWuVsD7oCPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseWuVsD7oCPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseWuVsD7oCPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseWuVsD7oCpcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseWuVsD7oCpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseWuVsD7oCPCAPsResponseSimple]
type pcapNewResponseWuVsD7oCpcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseWuVsD7oCPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseWuVsD7oCPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseWuVsD7oCPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseWuVsD7oCPCAPsResponseSimpleFilterV1]
type pcapNewResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseWuVsD7oCPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusUnknown           PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusSuccess           PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusPending           PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusRunning           PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusConversionPending PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusComplete          PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatusFailed            PCAPNewResponseWuVsD7oCPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseWuVsD7oCPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseWuVsD7oCPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseWuVsD7oCPCAPsResponseSimpleType string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleTypeSimple PCAPNewResponseWuVsD7oCPCAPsResponseSimpleType = "simple"
	PCAPNewResponseWuVsD7oCPCAPsResponseSimpleTypeFull   PCAPNewResponseWuVsD7oCPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseWuVsD7oCPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseWuVsD7oCPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseWuVsD7oCPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseWuVsD7oCPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseWuVsD7oCpcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponseWuVsD7oCpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseWuVsD7oCPCAPsResponseFull]
type pcapNewResponseWuVsD7oCpcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponseWuVsD7oCPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseWuVsD7oCPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseWuVsD7oCPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseWuVsD7oCpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseWuVsD7oCpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseWuVsD7oCPCAPsResponseFullFilterV1]
type pcapNewResponseWuVsD7oCpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseWuVsD7oCPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusUnknown           PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusSuccess           PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "success"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusPending           PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "pending"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusRunning           PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "running"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusConversionPending PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusConversionRunning PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusComplete          PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "complete"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullStatusFailed            PCAPNewResponseWuVsD7oCPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseWuVsD7oCPCAPsResponseFullSystem string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseFullSystemMagicTransit PCAPNewResponseWuVsD7oCPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseWuVsD7oCPCAPsResponseFullType string

const (
	PCAPNewResponseWuVsD7oCPCAPsResponseFullTypeSimple PCAPNewResponseWuVsD7oCPCAPsResponseFullType = "simple"
	PCAPNewResponseWuVsD7oCPCAPsResponseFullTypeFull   PCAPNewResponseWuVsD7oCPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseWuVsD7oCPCAPsResponseSimple] or
// [PCAPListResponseWuVsD7oCPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseWuVsD7oCPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseWuVsD7oCPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseWuVsD7oCPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseWuVsD7oCPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseWuVsD7oCpcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseWuVsD7oCpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseWuVsD7oCPCAPsResponseSimple]
type pcapListResponseWuVsD7oCpcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseWuVsD7oCPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseWuVsD7oCPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseWuVsD7oCPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseWuVsD7oCPCAPsResponseSimpleFilterV1]
type pcapListResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseWuVsD7oCPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusUnknown           PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusSuccess           PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "success"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusPending           PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusRunning           PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "running"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusConversionPending PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusConversionRunning PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusComplete          PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatusFailed            PCAPListResponseWuVsD7oCPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseWuVsD7oCPCAPsResponseSimpleSystem string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleSystemMagicTransit PCAPListResponseWuVsD7oCPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseWuVsD7oCPCAPsResponseSimpleType string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleTypeSimple PCAPListResponseWuVsD7oCPCAPsResponseSimpleType = "simple"
	PCAPListResponseWuVsD7oCPCAPsResponseSimpleTypeFull   PCAPListResponseWuVsD7oCPCAPsResponseSimpleType = "full"
)

type PCAPListResponseWuVsD7oCPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseWuVsD7oCPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseWuVsD7oCPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseWuVsD7oCPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseWuVsD7oCPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseWuVsD7oCpcaPsResponseFullJSON `json:"-"`
}

// pcapListResponseWuVsD7oCpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseWuVsD7oCPCAPsResponseFull]
type pcapListResponseWuVsD7oCpcaPsResponseFullJSON struct {
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

func (r *PCAPListResponseWuVsD7oCPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseWuVsD7oCPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseWuVsD7oCPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseWuVsD7oCpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseWuVsD7oCpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseWuVsD7oCPCAPsResponseFullFilterV1]
type pcapListResponseWuVsD7oCpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseWuVsD7oCPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseWuVsD7oCPCAPsResponseFullStatus string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusUnknown           PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "unknown"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusSuccess           PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "success"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusPending           PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "pending"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusRunning           PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "running"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusConversionPending PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusConversionRunning PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusComplete          PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "complete"
	PCAPListResponseWuVsD7oCPCAPsResponseFullStatusFailed            PCAPListResponseWuVsD7oCPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseWuVsD7oCPCAPsResponseFullSystem string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseFullSystemMagicTransit PCAPListResponseWuVsD7oCPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseWuVsD7oCPCAPsResponseFullType string

const (
	PCAPListResponseWuVsD7oCPCAPsResponseFullTypeSimple PCAPListResponseWuVsD7oCPCAPsResponseFullType = "simple"
	PCAPListResponseWuVsD7oCPCAPsResponseFullTypeFull   PCAPListResponseWuVsD7oCPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseWuVsD7oCPCAPsResponseSimple] or
// [PCAPGetResponseWuVsD7oCPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseWuVsD7oCPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseWuVsD7oCPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseWuVsD7oCPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseWuVsD7oCPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseWuVsD7oCpcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseWuVsD7oCpcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseWuVsD7oCPCAPsResponseSimple]
type pcapGetResponseWuVsD7oCpcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseWuVsD7oCPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseWuVsD7oCPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseWuVsD7oCPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseWuVsD7oCPCAPsResponseSimpleFilterV1]
type pcapGetResponseWuVsD7oCpcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseWuVsD7oCPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusUnknown           PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusSuccess           PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusPending           PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusRunning           PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusConversionPending PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusComplete          PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatusFailed            PCAPGetResponseWuVsD7oCPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseWuVsD7oCPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseWuVsD7oCPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseWuVsD7oCPCAPsResponseSimpleType string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleTypeSimple PCAPGetResponseWuVsD7oCPCAPsResponseSimpleType = "simple"
	PCAPGetResponseWuVsD7oCPCAPsResponseSimpleTypeFull   PCAPGetResponseWuVsD7oCPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseWuVsD7oCPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseWuVsD7oCPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseWuVsD7oCPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseWuVsD7oCPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseWuVsD7oCpcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponseWuVsD7oCpcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseWuVsD7oCPCAPsResponseFull]
type pcapGetResponseWuVsD7oCpcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponseWuVsD7oCPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseWuVsD7oCPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseWuVsD7oCPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseWuVsD7oCpcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseWuVsD7oCpcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseWuVsD7oCPCAPsResponseFullFilterV1]
type pcapGetResponseWuVsD7oCpcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseWuVsD7oCPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusUnknown           PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusSuccess           PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "success"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusPending           PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "pending"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusRunning           PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "running"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusConversionPending PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusConversionRunning PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusComplete          PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "complete"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullStatusFailed            PCAPGetResponseWuVsD7oCPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseWuVsD7oCPCAPsResponseFullSystem string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseFullSystemMagicTransit PCAPGetResponseWuVsD7oCPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseWuVsD7oCPCAPsResponseFullType string

const (
	PCAPGetResponseWuVsD7oCPCAPsResponseFullTypeSimple PCAPGetResponseWuVsD7oCPCAPsResponseFullType = "simple"
	PCAPGetResponseWuVsD7oCPCAPsResponseFullTypeFull   PCAPGetResponseWuVsD7oCPCAPsResponseFullType = "full"
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
