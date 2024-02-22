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

// Union satisfied by [PCAPNewResponseAN6TEptAPCAPsResponseSimple] or
// [PCAPNewResponseAN6TEptAPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseAN6TEptAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseAN6TEptAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseAN6TEptAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseAN6TEptAPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseAn6TEptApcaPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseAn6TEptApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseAN6TEptAPCAPsResponseSimple]
type pcapNewResponseAn6TEptApcaPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseAN6TEptAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseAN6TEptAPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseAN6TEptAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseAn6TEptApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseAn6TEptApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseAN6TEptAPCAPsResponseSimpleFilterV1]
type pcapNewResponseAn6TEptApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseAN6TEptAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusUnknown           PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusSuccess           PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusPending           PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusRunning           PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusConversionPending PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusComplete          PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatusFailed            PCAPNewResponseAN6TEptAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseAN6TEptAPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseAN6TEptAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseAN6TEptAPCAPsResponseSimpleType string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleTypeSimple PCAPNewResponseAN6TEptAPCAPsResponseSimpleType = "simple"
	PCAPNewResponseAN6TEptAPCAPsResponseSimpleTypeFull   PCAPNewResponseAN6TEptAPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseAN6TEptAPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseAN6TEptAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseAN6TEptAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseAN6TEptAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseAN6TEptAPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseAn6TEptApcaPsResponseFullJSON `json:"-"`
}

// pcapNewResponseAn6TEptApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseAN6TEptAPCAPsResponseFull]
type pcapNewResponseAn6TEptApcaPsResponseFullJSON struct {
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

func (r *PCAPNewResponseAN6TEptAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseAN6TEptAPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseAN6TEptAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseAn6TEptApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseAn6TEptApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseAN6TEptAPCAPsResponseFullFilterV1]
type pcapNewResponseAn6TEptApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseAN6TEptAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseAN6TEptAPCAPsResponseFullStatus string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusUnknown           PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusSuccess           PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "success"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusPending           PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "pending"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusRunning           PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "running"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusConversionPending PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusConversionRunning PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusComplete          PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "complete"
	PCAPNewResponseAN6TEptAPCAPsResponseFullStatusFailed            PCAPNewResponseAN6TEptAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseAN6TEptAPCAPsResponseFullSystem string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseFullSystemMagicTransit PCAPNewResponseAN6TEptAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseAN6TEptAPCAPsResponseFullType string

const (
	PCAPNewResponseAN6TEptAPCAPsResponseFullTypeSimple PCAPNewResponseAN6TEptAPCAPsResponseFullType = "simple"
	PCAPNewResponseAN6TEptAPCAPsResponseFullTypeFull   PCAPNewResponseAN6TEptAPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseAN6TEptAPCAPsResponseSimple] or
// [PCAPListResponseAN6TEptAPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseAN6TEptAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseAN6TEptAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseAN6TEptAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseAN6TEptAPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseAn6TEptApcaPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseAn6TEptApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseAN6TEptAPCAPsResponseSimple]
type pcapListResponseAn6TEptApcaPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseAN6TEptAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseAN6TEptAPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseAN6TEptAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseAn6TEptApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseAn6TEptApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseAN6TEptAPCAPsResponseSimpleFilterV1]
type pcapListResponseAn6TEptApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseAN6TEptAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus string

const (
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusUnknown           PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusSuccess           PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "success"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusPending           PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusRunning           PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "running"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusConversionPending PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusConversionRunning PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusComplete          PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleStatusFailed            PCAPListResponseAN6TEptAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseAN6TEptAPCAPsResponseSimpleSystem string

const (
	PCAPListResponseAN6TEptAPCAPsResponseSimpleSystemMagicTransit PCAPListResponseAN6TEptAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseAN6TEptAPCAPsResponseSimpleType string

const (
	PCAPListResponseAN6TEptAPCAPsResponseSimpleTypeSimple PCAPListResponseAN6TEptAPCAPsResponseSimpleType = "simple"
	PCAPListResponseAN6TEptAPCAPsResponseSimpleTypeFull   PCAPListResponseAN6TEptAPCAPsResponseSimpleType = "full"
)

type PCAPListResponseAN6TEptAPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseAN6TEptAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseAN6TEptAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseAN6TEptAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseAN6TEptAPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseAn6TEptApcaPsResponseFullJSON `json:"-"`
}

// pcapListResponseAn6TEptApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseAN6TEptAPCAPsResponseFull]
type pcapListResponseAn6TEptApcaPsResponseFullJSON struct {
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

func (r *PCAPListResponseAN6TEptAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseAN6TEptAPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseAN6TEptAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseAn6TEptApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseAn6TEptApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseAN6TEptAPCAPsResponseFullFilterV1]
type pcapListResponseAn6TEptApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseAN6TEptAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseAN6TEptAPCAPsResponseFullStatus string

const (
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusUnknown           PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "unknown"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusSuccess           PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "success"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusPending           PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "pending"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusRunning           PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "running"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusConversionPending PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusConversionRunning PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusComplete          PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "complete"
	PCAPListResponseAN6TEptAPCAPsResponseFullStatusFailed            PCAPListResponseAN6TEptAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseAN6TEptAPCAPsResponseFullSystem string

const (
	PCAPListResponseAN6TEptAPCAPsResponseFullSystemMagicTransit PCAPListResponseAN6TEptAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseAN6TEptAPCAPsResponseFullType string

const (
	PCAPListResponseAN6TEptAPCAPsResponseFullTypeSimple PCAPListResponseAN6TEptAPCAPsResponseFullType = "simple"
	PCAPListResponseAN6TEptAPCAPsResponseFullTypeFull   PCAPListResponseAN6TEptAPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseAN6TEptAPCAPsResponseSimple] or
// [PCAPGetResponseAN6TEptAPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseAN6TEptAPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseAN6TEptAPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseAN6TEptAPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseAN6TEptAPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseAn6TEptApcaPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseAn6TEptApcaPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseAN6TEptAPCAPsResponseSimple]
type pcapGetResponseAn6TEptApcaPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseAN6TEptAPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseAN6TEptAPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseAN6TEptAPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseAn6TEptApcaPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseAn6TEptApcaPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseAN6TEptAPCAPsResponseSimpleFilterV1]
type pcapGetResponseAn6TEptApcaPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseAN6TEptAPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusUnknown           PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusSuccess           PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusPending           PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusRunning           PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusConversionPending PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusComplete          PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatusFailed            PCAPGetResponseAN6TEptAPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseAN6TEptAPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseAN6TEptAPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseAN6TEptAPCAPsResponseSimpleType string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleTypeSimple PCAPGetResponseAN6TEptAPCAPsResponseSimpleType = "simple"
	PCAPGetResponseAN6TEptAPCAPsResponseSimpleTypeFull   PCAPGetResponseAN6TEptAPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseAN6TEptAPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseAN6TEptAPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseAN6TEptAPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseAN6TEptAPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseAN6TEptAPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseAn6TEptApcaPsResponseFullJSON `json:"-"`
}

// pcapGetResponseAn6TEptApcaPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseAN6TEptAPCAPsResponseFull]
type pcapGetResponseAn6TEptApcaPsResponseFullJSON struct {
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

func (r *PCAPGetResponseAN6TEptAPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseAN6TEptAPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseAN6TEptAPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseAn6TEptApcaPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseAn6TEptApcaPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseAN6TEptAPCAPsResponseFullFilterV1]
type pcapGetResponseAn6TEptApcaPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseAN6TEptAPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseAN6TEptAPCAPsResponseFullStatus string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusUnknown           PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusSuccess           PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "success"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusPending           PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "pending"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusRunning           PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "running"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusConversionPending PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusConversionRunning PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusComplete          PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "complete"
	PCAPGetResponseAN6TEptAPCAPsResponseFullStatusFailed            PCAPGetResponseAN6TEptAPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseAN6TEptAPCAPsResponseFullSystem string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseFullSystemMagicTransit PCAPGetResponseAN6TEptAPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseAN6TEptAPCAPsResponseFullType string

const (
	PCAPGetResponseAN6TEptAPCAPsResponseFullTypeSimple PCAPGetResponseAN6TEptAPCAPsResponseFullType = "simple"
	PCAPGetResponseAN6TEptAPCAPsResponseFullTypeFull   PCAPGetResponseAN6TEptAPCAPsResponseFullType = "full"
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
