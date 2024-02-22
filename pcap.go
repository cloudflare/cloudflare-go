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

// Union satisfied by [PCAPNewResponseVpTbFtFcPCAPsResponseSimple] or
// [PCAPNewResponseVpTbFtFcPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(), "")
}

type PCAPNewResponseVpTbFtFcPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseVpTbFtFcPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseVpTbFtFcPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseVpTbFtFcPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseVpTbFtFcPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPNewResponseVpTbFtFcPCAPsResponseSimple]
type pcapNewResponseVpTbFtFcPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseVpTbFtFcPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseVpTbFtFcPCAPsResponseSimple) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1]
type pcapNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseVpTbFtFcPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusUnknown           PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusSuccess           PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusPending           PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusRunning           PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusConversionPending PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusComplete          PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatusFailed            PCAPNewResponseVpTbFtFcPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseVpTbFtFcPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseVpTbFtFcPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseVpTbFtFcPCAPsResponseSimpleType string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleTypeSimple PCAPNewResponseVpTbFtFcPCAPsResponseSimpleType = "simple"
	PCAPNewResponseVpTbFtFcPCAPsResponseSimpleTypeFull   PCAPNewResponseVpTbFtFcPCAPsResponseSimpleType = "full"
)

type PCAPNewResponseVpTbFtFcPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseVpTbFtFcPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseVpTbFtFcPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseVpTbFtFcPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseVpTbFtFcPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseVpTbFtFcPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPNewResponseVpTbFtFcPCAPsResponseFull]
type pcapNewResponseVpTbFtFcPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseVpTbFtFcPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPNewResponseVpTbFtFcPCAPsResponseFull) implementsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseVpTbFtFcPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapNewResponseVpTbFtFcPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseVpTbFtFcPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPNewResponseVpTbFtFcPCAPsResponseFullFilterV1]
type pcapNewResponseVpTbFtFcPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseVpTbFtFcPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusUnknown           PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusSuccess           PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "success"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusPending           PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "pending"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusRunning           PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "running"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusConversionPending PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusConversionRunning PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusComplete          PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "complete"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullStatusFailed            PCAPNewResponseVpTbFtFcPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPNewResponseVpTbFtFcPCAPsResponseFullSystem string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseFullSystemMagicTransit PCAPNewResponseVpTbFtFcPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseVpTbFtFcPCAPsResponseFullType string

const (
	PCAPNewResponseVpTbFtFcPCAPsResponseFullTypeSimple PCAPNewResponseVpTbFtFcPCAPsResponseFullType = "simple"
	PCAPNewResponseVpTbFtFcPCAPsResponseFullTypeFull   PCAPNewResponseVpTbFtFcPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPListResponseVpTbFtFcPCAPsResponseSimple] or
// [PCAPListResponseVpTbFtFcPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPListResponse)(nil)).Elem(), "")
}

type PCAPListResponseVpTbFtFcPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseVpTbFtFcPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseVpTbFtFcPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseVpTbFtFcPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseVpTbFtFcPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseVpTbFtFcPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPListResponseVpTbFtFcPCAPsResponseSimple]
type pcapListResponseVpTbFtFcPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseVpTbFtFcPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseVpTbFtFcPCAPsResponseSimple) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseVpTbFtFcPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapListResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPListResponseVpTbFtFcPCAPsResponseSimpleFilterV1]
type pcapListResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseVpTbFtFcPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusUnknown           PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusSuccess           PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "success"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusPending           PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusRunning           PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "running"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusConversionPending PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusConversionRunning PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusComplete          PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatusFailed            PCAPListResponseVpTbFtFcPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseVpTbFtFcPCAPsResponseSimpleSystem string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleSystemMagicTransit PCAPListResponseVpTbFtFcPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseVpTbFtFcPCAPsResponseSimpleType string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleTypeSimple PCAPListResponseVpTbFtFcPCAPsResponseSimpleType = "simple"
	PCAPListResponseVpTbFtFcPCAPsResponseSimpleTypeFull   PCAPListResponseVpTbFtFcPCAPsResponseSimpleType = "full"
)

type PCAPListResponseVpTbFtFcPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseVpTbFtFcPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseVpTbFtFcPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseVpTbFtFcPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseVpTbFtFcPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseVpTbFtFcPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseVpTbFtFcPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPListResponseVpTbFtFcPCAPsResponseFull]
type pcapListResponseVpTbFtFcPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseVpTbFtFcPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPListResponseVpTbFtFcPCAPsResponseFull) implementsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseVpTbFtFcPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapListResponseVpTbFtFcPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseVpTbFtFcPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPListResponseVpTbFtFcPCAPsResponseFullFilterV1]
type pcapListResponseVpTbFtFcPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseVpTbFtFcPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPListResponseVpTbFtFcPCAPsResponseFullStatus string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusUnknown           PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "unknown"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusSuccess           PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "success"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusPending           PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "pending"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusRunning           PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "running"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusConversionPending PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusConversionRunning PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusComplete          PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "complete"
	PCAPListResponseVpTbFtFcPCAPsResponseFullStatusFailed            PCAPListResponseVpTbFtFcPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPListResponseVpTbFtFcPCAPsResponseFullSystem string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseFullSystemMagicTransit PCAPListResponseVpTbFtFcPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseVpTbFtFcPCAPsResponseFullType string

const (
	PCAPListResponseVpTbFtFcPCAPsResponseFullTypeSimple PCAPListResponseVpTbFtFcPCAPsResponseFullType = "simple"
	PCAPListResponseVpTbFtFcPCAPsResponseFullTypeFull   PCAPListResponseVpTbFtFcPCAPsResponseFullType = "full"
)

// Union satisfied by [PCAPGetResponseVpTbFtFcPCAPsResponseSimple] or
// [PCAPGetResponseVpTbFtFcPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(), "")
}

type PCAPGetResponseVpTbFtFcPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseVpTbFtFcPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseVpTbFtFcPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseVpTbFtFcPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseVpTbFtFcPCAPsResponseSimpleJSON contains the JSON metadata for
// the struct [PCAPGetResponseVpTbFtFcPCAPsResponseSimple]
type pcapGetResponseVpTbFtFcPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseVpTbFtFcPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseVpTbFtFcPCAPsResponseSimple) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1 struct {
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
	JSON       pcapGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct [PCAPGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1]
type pcapGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseVpTbFtFcPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusUnknown           PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusSuccess           PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusPending           PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusRunning           PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusConversionPending PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusComplete          PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatusFailed            PCAPGetResponseVpTbFtFcPCAPsResponseSimpleStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseVpTbFtFcPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseVpTbFtFcPCAPsResponseSimpleSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseVpTbFtFcPCAPsResponseSimpleType string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleTypeSimple PCAPGetResponseVpTbFtFcPCAPsResponseSimpleType = "simple"
	PCAPGetResponseVpTbFtFcPCAPsResponseSimpleTypeFull   PCAPGetResponseVpTbFtFcPCAPsResponseSimpleType = "full"
)

type PCAPGetResponseVpTbFtFcPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseVpTbFtFcPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseVpTbFtFcPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseVpTbFtFcPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseVpTbFtFcPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseVpTbFtFcPCAPsResponseFullJSON contains the JSON metadata for the
// struct [PCAPGetResponseVpTbFtFcPCAPsResponseFull]
type pcapGetResponseVpTbFtFcPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseVpTbFtFcPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PCAPGetResponseVpTbFtFcPCAPsResponseFull) implementsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseVpTbFtFcPCAPsResponseFullFilterV1 struct {
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
	JSON       pcapGetResponseVpTbFtFcPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseVpTbFtFcPCAPsResponseFullFilterV1JSON contains the JSON metadata
// for the struct [PCAPGetResponseVpTbFtFcPCAPsResponseFullFilterV1]
type pcapGetResponseVpTbFtFcPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseVpTbFtFcPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the packet capture request.
type PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusUnknown           PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusSuccess           PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "success"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusPending           PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "pending"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusRunning           PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "running"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusConversionPending PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusConversionRunning PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusComplete          PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "complete"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullStatusFailed            PCAPGetResponseVpTbFtFcPCAPsResponseFullStatus = "failed"
)

// The system used to collect packet captures.
type PCAPGetResponseVpTbFtFcPCAPsResponseFullSystem string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseFullSystemMagicTransit PCAPGetResponseVpTbFtFcPCAPsResponseFullSystem = "magic-transit"
)

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseVpTbFtFcPCAPsResponseFullType string

const (
	PCAPGetResponseVpTbFtFcPCAPsResponseFullTypeSimple PCAPGetResponseVpTbFtFcPCAPsResponseFullType = "simple"
	PCAPGetResponseVpTbFtFcPCAPsResponseFullTypeFull   PCAPGetResponseVpTbFtFcPCAPsResponseFullType = "full"
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
