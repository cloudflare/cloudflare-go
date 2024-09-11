// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// PCAPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPCAPService] method instead.
type PCAPService struct {
	Options   []option.RequestOption
	Ownership *PCAPOwnershipService
	Download  *PCAPDownloadService
}

// NewPCAPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPCAPService(opts ...option.RequestOption) (r *PCAPService) {
	r = &PCAPService{}
	r.Options = opts
	r.Ownership = NewPCAPOwnershipService(opts...)
	r.Download = NewPCAPDownloadService(opts...)
	return
}

// Create new PCAP request for account.
func (r *PCAPService) New(ctx context.Context, params PCAPNewParams, opts ...option.RequestOption) (res *PCAPNewResponse, err error) {
	var env PCAPNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps", params.AccountID)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
	var env PCAPGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pcapID == "" {
		err = errors.New("missing required pcap_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/%s", query.AccountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PCAP struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPFilter `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPType `json:"type"`
	JSON pcapJSON `json:"-"`
}

// pcapJSON contains the JSON metadata for the struct [PCAP]
type pcapJSON struct {
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

func (r *PCAP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapJSON) RawJSON() string {
	return r.raw
}

func (r PCAP) implementsMagicTransitPCAPNewResponse() {}

func (r PCAP) implementsMagicTransitPCAPListResponse() {}

func (r PCAP) implementsMagicTransitPCAPGetResponse() {}

// The status of the packet capture request.
type PCAPStatus string

const (
	PCAPStatusUnknown           PCAPStatus = "unknown"
	PCAPStatusSuccess           PCAPStatus = "success"
	PCAPStatusPending           PCAPStatus = "pending"
	PCAPStatusRunning           PCAPStatus = "running"
	PCAPStatusConversionPending PCAPStatus = "conversion_pending"
	PCAPStatusConversionRunning PCAPStatus = "conversion_running"
	PCAPStatusComplete          PCAPStatus = "complete"
	PCAPStatusFailed            PCAPStatus = "failed"
)

func (r PCAPStatus) IsKnown() bool {
	switch r {
	case PCAPStatusUnknown, PCAPStatusSuccess, PCAPStatusPending, PCAPStatusRunning, PCAPStatusConversionPending, PCAPStatusConversionRunning, PCAPStatusComplete, PCAPStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPSystem string

const (
	PCAPSystemMagicTransit PCAPSystem = "magic-transit"
)

func (r PCAPSystem) IsKnown() bool {
	switch r {
	case PCAPSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPType string

const (
	PCAPTypeSimple PCAPType = "simple"
	PCAPTypeFull   PCAPType = "full"
)

func (r PCAPType) IsKnown() bool {
	switch r {
	case PCAPTypeSimple, PCAPTypeFull:
		return true
	}
	return false
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPFilter struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64        `json:"source_port"`
	JSON       pcapFilterJSON `json:"-"`
}

// pcapFilterJSON contains the JSON metadata for the struct [PCAPFilter]
type pcapFilterJSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapFilterJSON) RawJSON() string {
	return r.raw
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPFilterParam struct {
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

func (r PCAPFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPNewResponse struct {
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPFilter `json:"filter_v1"`
	// The ID for the packet capture.
	ID string `json:"id"`
	// The status of the packet capture request.
	Status PCAPNewResponseStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseType `json:"type"`
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
	ErrorMessage string              `json:"error_message"`
	JSON         pcapNewResponseJSON `json:"-"`
	union        PCAPNewResponseUnion
}

// pcapNewResponseJSON contains the JSON metadata for the struct [PCAPNewResponse]
type pcapNewResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r pcapNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *PCAPNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = PCAPNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PCAPNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [magic_transit.PCAP],
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsResponseFull].
func (r PCAPNewResponse) AsUnion() PCAPNewResponseUnion {
	return r.union
}

// Union satisfied by [magic_transit.PCAP] or
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsResponseFull].
type PCAPNewResponseUnion interface {
	implementsMagicTransitPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPNewResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
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
	FilterV1 PCAPFilter `json:"filter_v1"`
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

func (r PCAPNewResponseMagicVisibilityPCAPsResponseFull) implementsMagicTransitPCAPNewResponse() {}

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

// The status of the packet capture request.
type PCAPNewResponseStatus string

const (
	PCAPNewResponseStatusUnknown           PCAPNewResponseStatus = "unknown"
	PCAPNewResponseStatusSuccess           PCAPNewResponseStatus = "success"
	PCAPNewResponseStatusPending           PCAPNewResponseStatus = "pending"
	PCAPNewResponseStatusRunning           PCAPNewResponseStatus = "running"
	PCAPNewResponseStatusConversionPending PCAPNewResponseStatus = "conversion_pending"
	PCAPNewResponseStatusConversionRunning PCAPNewResponseStatus = "conversion_running"
	PCAPNewResponseStatusComplete          PCAPNewResponseStatus = "complete"
	PCAPNewResponseStatusFailed            PCAPNewResponseStatus = "failed"
)

func (r PCAPNewResponseStatus) IsKnown() bool {
	switch r {
	case PCAPNewResponseStatusUnknown, PCAPNewResponseStatusSuccess, PCAPNewResponseStatusPending, PCAPNewResponseStatusRunning, PCAPNewResponseStatusConversionPending, PCAPNewResponseStatusConversionRunning, PCAPNewResponseStatusComplete, PCAPNewResponseStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewResponseSystem string

const (
	PCAPNewResponseSystemMagicTransit PCAPNewResponseSystem = "magic-transit"
)

func (r PCAPNewResponseSystem) IsKnown() bool {
	switch r {
	case PCAPNewResponseSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseType string

const (
	PCAPNewResponseTypeSimple PCAPNewResponseType = "simple"
	PCAPNewResponseTypeFull   PCAPNewResponseType = "full"
)

func (r PCAPNewResponseType) IsKnown() bool {
	switch r {
	case PCAPNewResponseTypeSimple, PCAPNewResponseTypeFull:
		return true
	}
	return false
}

type PCAPListResponse struct {
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPFilter `json:"filter_v1"`
	// The ID for the packet capture.
	ID string `json:"id"`
	// The status of the packet capture request.
	Status PCAPListResponseStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseType `json:"type"`
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
	ErrorMessage string               `json:"error_message"`
	JSON         pcapListResponseJSON `json:"-"`
	union        PCAPListResponseUnion
}

// pcapListResponseJSON contains the JSON metadata for the struct
// [PCAPListResponse]
type pcapListResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r pcapListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *PCAPListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = PCAPListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PCAPListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [magic_transit.PCAP],
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsResponseFull].
func (r PCAPListResponse) AsUnion() PCAPListResponseUnion {
	return r.union
}

// Union satisfied by [magic_transit.PCAP] or
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsResponseFull].
type PCAPListResponseUnion interface {
	implementsMagicTransitPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPListResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
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
	FilterV1 PCAPFilter `json:"filter_v1"`
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

func (r PCAPListResponseMagicVisibilityPCAPsResponseFull) implementsMagicTransitPCAPListResponse() {}

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

// The status of the packet capture request.
type PCAPListResponseStatus string

const (
	PCAPListResponseStatusUnknown           PCAPListResponseStatus = "unknown"
	PCAPListResponseStatusSuccess           PCAPListResponseStatus = "success"
	PCAPListResponseStatusPending           PCAPListResponseStatus = "pending"
	PCAPListResponseStatusRunning           PCAPListResponseStatus = "running"
	PCAPListResponseStatusConversionPending PCAPListResponseStatus = "conversion_pending"
	PCAPListResponseStatusConversionRunning PCAPListResponseStatus = "conversion_running"
	PCAPListResponseStatusComplete          PCAPListResponseStatus = "complete"
	PCAPListResponseStatusFailed            PCAPListResponseStatus = "failed"
)

func (r PCAPListResponseStatus) IsKnown() bool {
	switch r {
	case PCAPListResponseStatusUnknown, PCAPListResponseStatusSuccess, PCAPListResponseStatusPending, PCAPListResponseStatusRunning, PCAPListResponseStatusConversionPending, PCAPListResponseStatusConversionRunning, PCAPListResponseStatusComplete, PCAPListResponseStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPListResponseSystem string

const (
	PCAPListResponseSystemMagicTransit PCAPListResponseSystem = "magic-transit"
)

func (r PCAPListResponseSystem) IsKnown() bool {
	switch r {
	case PCAPListResponseSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseType string

const (
	PCAPListResponseTypeSimple PCAPListResponseType = "simple"
	PCAPListResponseTypeFull   PCAPListResponseType = "full"
)

func (r PCAPListResponseType) IsKnown() bool {
	switch r {
	case PCAPListResponseTypeSimple, PCAPListResponseTypeFull:
		return true
	}
	return false
}

type PCAPGetResponse struct {
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPFilter `json:"filter_v1"`
	// The ID for the packet capture.
	ID string `json:"id"`
	// The status of the packet capture request.
	Status PCAPGetResponseStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseType `json:"type"`
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
	ErrorMessage string              `json:"error_message"`
	JSON         pcapGetResponseJSON `json:"-"`
	union        PCAPGetResponseUnion
}

// pcapGetResponseJSON contains the JSON metadata for the struct [PCAPGetResponse]
type pcapGetResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r pcapGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *PCAPGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = PCAPGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PCAPGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [magic_transit.PCAP],
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsResponseFull].
func (r PCAPGetResponse) AsUnion() PCAPGetResponseUnion {
	return r.union
}

// Union satisfied by [magic_transit.PCAP] or
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsResponseFull].
type PCAPGetResponseUnion interface {
	implementsMagicTransitPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPGetResponseMagicVisibilityPCAPsResponseFull{}),
		},
	)
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
	FilterV1 PCAPFilter `json:"filter_v1"`
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

func (r PCAPGetResponseMagicVisibilityPCAPsResponseFull) implementsMagicTransitPCAPGetResponse() {}

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

// The status of the packet capture request.
type PCAPGetResponseStatus string

const (
	PCAPGetResponseStatusUnknown           PCAPGetResponseStatus = "unknown"
	PCAPGetResponseStatusSuccess           PCAPGetResponseStatus = "success"
	PCAPGetResponseStatusPending           PCAPGetResponseStatus = "pending"
	PCAPGetResponseStatusRunning           PCAPGetResponseStatus = "running"
	PCAPGetResponseStatusConversionPending PCAPGetResponseStatus = "conversion_pending"
	PCAPGetResponseStatusConversionRunning PCAPGetResponseStatus = "conversion_running"
	PCAPGetResponseStatusComplete          PCAPGetResponseStatus = "complete"
	PCAPGetResponseStatusFailed            PCAPGetResponseStatus = "failed"
)

func (r PCAPGetResponseStatus) IsKnown() bool {
	switch r {
	case PCAPGetResponseStatusUnknown, PCAPGetResponseStatusSuccess, PCAPGetResponseStatusPending, PCAPGetResponseStatusRunning, PCAPGetResponseStatusConversionPending, PCAPGetResponseStatusConversionRunning, PCAPGetResponseStatusComplete, PCAPGetResponseStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPGetResponseSystem string

const (
	PCAPGetResponseSystemMagicTransit PCAPGetResponseSystem = "magic-transit"
)

func (r PCAPGetResponseSystem) IsKnown() bool {
	switch r {
	case PCAPGetResponseSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseType string

const (
	PCAPGetResponseTypeSimple PCAPGetResponseType = "simple"
	PCAPGetResponseTypeFull   PCAPGetResponseType = "full"
)

func (r PCAPGetResponseType) IsKnown() bool {
	switch r {
	case PCAPGetResponseTypeSimple, PCAPGetResponseTypeFull:
		return true
	}
	return false
}

type PCAPNewParams struct {
	// Identifier
	AccountID param.Field[string]    `path:"account_id,required"`
	Body      PCAPNewParamsBodyUnion `json:"body,required"`
}

func (r PCAPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PCAPNewParamsBody struct {
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPFilterParam] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsBodySystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsBodyType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf"`
}

func (r PCAPNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBody) implementsMagicTransitPCAPNewParamsBodyUnion() {}

// Satisfied by [magic_transit.PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimple],
// [magic_transit.PCAPNewParamsBodyMagicVisibilityPCAPsRequestFull],
// [PCAPNewParamsBody].
type PCAPNewParamsBodyUnion interface {
	implementsMagicTransitPCAPNewParamsBodyUnion()
}

type PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPFilterParam] `json:"filter_v1"`
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimple) implementsMagicTransitPCAPNewParamsBodyUnion() {
}

// The system used to collect packet captures.
type PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystem string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystemMagicTransit PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystem = "magic-transit"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleType string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleTypeSimple PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleType = "simple"
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleTypeFull   PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleType = "full"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleType) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleTypeSimple, PCAPNewParamsBodyMagicVisibilityPCAPsRequestSimpleTypeFull:
		return true
	}
	return false
}

type PCAPNewParamsBodyMagicVisibilityPCAPsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPFilterParam] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestFull) implementsMagicTransitPCAPNewParamsBodyUnion() {
}

// The system used to collect packet captures.
type PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystem string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystemMagicTransit PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystem = "magic-transit"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullType string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullTypeSimple PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullType = "simple"
	PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullTypeFull   PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullType = "full"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullType) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullTypeSimple, PCAPNewParamsBodyMagicVisibilityPCAPsRequestFullTypeFull:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewParamsBodySystem string

const (
	PCAPNewParamsBodySystemMagicTransit PCAPNewParamsBodySystem = "magic-transit"
)

func (r PCAPNewParamsBodySystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodySystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsBodyType string

const (
	PCAPNewParamsBodyTypeSimple PCAPNewParamsBodyType = "simple"
	PCAPNewParamsBodyTypeFull   PCAPNewParamsBodyType = "full"
)

func (r PCAPNewParamsBodyType) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyTypeSimple, PCAPNewParamsBodyTypeFull:
		return true
	}
	return false
}

type PCAPNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PCAPNewResponse       `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PCAPGetResponse       `json:"result,required"`
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
