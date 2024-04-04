// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pcaps

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// PCAPService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPCAPService] method instead.
type PCAPService struct {
	Options   []option.RequestOption
	Ownership *OwnershipService
	Download  *DownloadService
}

// NewPCAPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPCAPService(opts ...option.RequestOption) (r *PCAPService) {
	r = &PCAPService{}
	r.Options = opts
	r.Ownership = NewOwnershipService(opts...)
	r.Download = NewDownloadService(opts...)
	return
}

// Create new PCAP request for account.
func (r *PCAPService) New(ctx context.Context, params PCAPNewParams, opts ...option.RequestOption) (res *PCAPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps", params.getAccountID())
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all packet capture requests for an account.
func (r *PCAPService) List(ctx context.Context, query PCAPListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PCAPListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all packet capture requests for an account.
func (r *PCAPService) ListAutoPaging(ctx context.Context, query PCAPListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PCAPListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Get information for a PCAP request by id.
func (r *PCAPService) Get(ctx context.Context, pcapID string, query PCAPGetParams, opts ...option.RequestOption) (res *PCAPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PCAPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/%s", query.AccountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [pcaps.PCAPNewResponseMagicVisibilityPCAPsResponseSimple] or
// [pcaps.PCAPNewResponseMagicVisibilityPCAPsResponseFull].
type PCAPNewResponse interface {
	implementsPCAPsPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPNewResponseMagicVisibilityPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPNewResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
}

type PCAPNewResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON metadata
// for the struct [PCAPNewResponseMagicVisibilityPCAPsResponseSimple]
type pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPNewResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                       `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusPending, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusRunning, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusComplete, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeSimple, PCAPNewResponseMagicVisibilityPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPNewResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPNewResponseMagicVisibilityPCAPsResponseFull]
type pcapNewResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFull) implementsPCAPsPCAPNewResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                     `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusUnknown, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusSuccess, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusPending, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusRunning, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionPending, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusComplete, PCAPNewResponseMagicVisibilityPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPNewResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPNewResponseMagicVisibilityPCAPsResponseFullType = "full"
)

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeSimple, PCAPNewResponseMagicVisibilityPCAPsResponseFullTypeFull:
		return true
	}
	return false
}

// Union satisfied by [pcaps.PCAPListResponseMagicVisibilityPCAPsResponseSimple] or
// [pcaps.PCAPListResponseMagicVisibilityPCAPsResponseFull].
type PCAPListResponse interface {
	implementsPCAPsPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPListResponseMagicVisibilityPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPListResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
}

type PCAPListResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON
// metadata for the struct [PCAPListResponseMagicVisibilityPCAPsResponseSimple]
type pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPListResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                        `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusPending, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusRunning, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusComplete, PCAPListResponseMagicVisibilityPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPListResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPListResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeSimple, PCAPListResponseMagicVisibilityPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPListResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPListResponseMagicVisibilityPCAPsResponseFull]
type pcapListResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPListResponseMagicVisibilityPCAPsResponseFull) implementsPCAPsPCAPListResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                      `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPListResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseFullStatusUnknown, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusSuccess, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusPending, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusRunning, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionPending, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusComplete, PCAPListResponseMagicVisibilityPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPListResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPListResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPListResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPListResponseMagicVisibilityPCAPsResponseFullType = "full"
)

func (r PCAPListResponseMagicVisibilityPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsResponseFullTypeSimple, PCAPListResponseMagicVisibilityPCAPsResponseFullTypeFull:
		return true
	}
	return false
}

// Union satisfied by [pcaps.PCAPGetResponseMagicVisibilityPCAPsResponseSimple] or
// [pcaps.PCAPGetResponseMagicVisibilityPCAPsResponseFull].
type PCAPGetResponse interface {
	implementsPCAPsPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPGetResponseMagicVisibilityPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPGetResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
}

type PCAPGetResponseMagicVisibilityPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON contains the JSON metadata
// for the struct [PCAPGetResponseMagicVisibilityPCAPsResponseSimple]
type pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPGetResponseMagicVisibilityPCAPsResponseSimple) implementsPCAPsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                       `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1]
type pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusPending           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusRunning           PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusComplete          PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusFailed            PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusUnknown, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusSuccess, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusPending, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusRunning, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionPending, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusConversionRunning, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusComplete, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeSimple PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeFull   PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType = "full"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeSimple, PCAPGetResponseMagicVisibilityPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPGetResponseMagicVisibilityPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseFullJSON contains the JSON metadata
// for the struct [PCAPGetResponseMagicVisibilityPCAPsResponseFull]
type pcapGetResponseMagicVisibilityPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFull) implementsPCAPsPCAPGetResponse() {}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                     `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON contains the JSON
// metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1]
type pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusPending           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusRunning           PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusComplete          PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusFailed            PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus = "failed"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusUnknown, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusSuccess, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusPending, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusRunning, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionPending, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusConversionRunning, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusComplete, PCAPGetResponseMagicVisibilityPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsResponseFullType string

const (
	PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeSimple PCAPGetResponseMagicVisibilityPCAPsResponseFullType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeFull   PCAPGetResponseMagicVisibilityPCAPsResponseFullType = "full"
)

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeSimple, PCAPGetResponseMagicVisibilityPCAPsResponseFullTypeFull:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [PCAPNewParamsMagicVisibilityPCAPsRequestSimple],
// [PCAPNewParamsMagicVisibilityPCAPsRequestFull].
type PCAPNewParams interface {
	ImplementsPCAPNewParams()

	getAccountID() param.Field[string]
}

type PCAPNewParamsMagicVisibilityPCAPsRequestSimple struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PCAPNewParamsMagicVisibilityPCAPsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsMagicVisibilityPCAPsRequestSimple) getAccountID() param.Field[string] {
	return r.AccountID
}

func (PCAPNewParamsMagicVisibilityPCAPsRequestSimple) ImplementsPCAPNewParams() {

}

// The system used to collect packet captures.
type PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystem string

const (
	PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystemMagicTransit PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystem = "magic-transit"
)

func (r PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsMagicVisibilityPCAPsRequestSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsMagicVisibilityPCAPsRequestSimpleType string

const (
	PCAPNewParamsMagicVisibilityPCAPsRequestSimpleTypeSimple PCAPNewParamsMagicVisibilityPCAPsRequestSimpleType = "simple"
	PCAPNewParamsMagicVisibilityPCAPsRequestSimpleTypeFull   PCAPNewParamsMagicVisibilityPCAPsRequestSimpleType = "full"
)

func (r PCAPNewParamsMagicVisibilityPCAPsRequestSimpleType) IsKnown() bool {
	switch r {
	case PCAPNewParamsMagicVisibilityPCAPsRequestSimpleTypeSimple, PCAPNewParamsMagicVisibilityPCAPsRequestSimpleTypeFull:
		return true
	}
	return false
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewParamsMagicVisibilityPCAPsRequestSimpleFilterV1 struct {
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

func (r PCAPNewParamsMagicVisibilityPCAPsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPNewParamsMagicVisibilityPCAPsRequestFull struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPNewParamsMagicVisibilityPCAPsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PCAPNewParamsMagicVisibilityPCAPsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsMagicVisibilityPCAPsRequestFull) getAccountID() param.Field[string] {
	return r.AccountID
}

func (PCAPNewParamsMagicVisibilityPCAPsRequestFull) ImplementsPCAPNewParams() {

}

// The system used to collect packet captures.
type PCAPNewParamsMagicVisibilityPCAPsRequestFullSystem string

const (
	PCAPNewParamsMagicVisibilityPCAPsRequestFullSystemMagicTransit PCAPNewParamsMagicVisibilityPCAPsRequestFullSystem = "magic-transit"
)

func (r PCAPNewParamsMagicVisibilityPCAPsRequestFullSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsMagicVisibilityPCAPsRequestFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsMagicVisibilityPCAPsRequestFullType string

const (
	PCAPNewParamsMagicVisibilityPCAPsRequestFullTypeSimple PCAPNewParamsMagicVisibilityPCAPsRequestFullType = "simple"
	PCAPNewParamsMagicVisibilityPCAPsRequestFullTypeFull   PCAPNewParamsMagicVisibilityPCAPsRequestFullType = "full"
)

func (r PCAPNewParamsMagicVisibilityPCAPsRequestFullType) IsKnown() bool {
	switch r {
	case PCAPNewParamsMagicVisibilityPCAPsRequestFullTypeSimple, PCAPNewParamsMagicVisibilityPCAPsRequestFullTypeFull:
		return true
	}
	return false
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewParamsMagicVisibilityPCAPsRequestFullFilterV1 struct {
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

func (r PCAPNewParamsMagicVisibilityPCAPsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   PCAPNewResponse              `json:"result,required"`
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

func (r pcapNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PCAPNewResponseEnvelopeSuccess bool

const (
	PCAPNewResponseEnvelopeSuccessTrue PCAPNewResponseEnvelopeSuccess = true
)

func (r PCAPNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PCAPNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PCAPListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PCAPGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PCAPGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   PCAPGetResponse              `json:"result,required"`
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

func (r pcapGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PCAPGetResponseEnvelopeSuccess bool

const (
	PCAPGetResponseEnvelopeSuccessTrue PCAPGetResponseEnvelopeSuccess = true
)

func (r PCAPGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PCAPGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
