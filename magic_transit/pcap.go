// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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

type PCAPNewResponse struct {
	// This field can have the runtime type of
	// [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1],
	// [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1].
	FilterV1 interface{} `json:"filter_v1,required"`
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
	Type  PCAPNewResponseType `json:"type"`
	JSON  pcapNewResponseJSON `json:"-"`
	union PCAPNewResponseUnion
}

// pcapNewResponseJSON contains the JSON metadata for the struct [PCAPNewResponse]
type pcapNewResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
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
// Possible runtime types of the union are
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple],
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull].
func (r PCAPNewResponse) AsUnion() PCAPNewResponseUnion {
	return r.union
}

// Union satisfied by
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple] or
// [magic_transit.PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull].
type PCAPNewResponseUnion interface {
	implementsMagicTransitPCAPNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull{}),
		},
	)
}

type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON contains the JSON
// metadata for the struct [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple]
type pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimple) implementsMagicTransitPCAPNewResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                            `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1]
type pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete          PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed            PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleType string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull   PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "full"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull struct {
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
	FilterV1 PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullType `json:"type"`
	JSON pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullJSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullJSON contains the JSON
// metadata for the struct [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull]
type pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullJSON struct {
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

func (r *PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFull) implementsMagicTransitPCAPNewResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                          `json:"source_port"`
	JSON       pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1]
type pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapNewResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "unknown"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "success"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "pending"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning           PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "running"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_pending"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_running"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete          PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "complete"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed            PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "failed"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystem string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullType string

const (
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullType = "simple"
	PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull   PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullType = "full"
)

func (r PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple, PCAPNewResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull:
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
	// This field can have the runtime type of
	// [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1],
	// [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1].
	FilterV1 interface{} `json:"filter_v1,required"`
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
	Type  PCAPListResponseType `json:"type"`
	JSON  pcapListResponseJSON `json:"-"`
	union PCAPListResponseUnion
}

// pcapListResponseJSON contains the JSON metadata for the struct
// [PCAPListResponse]
type pcapListResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
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
// Possible runtime types of the union are
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple],
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull].
func (r PCAPListResponse) AsUnion() PCAPListResponseUnion {
	return r.union
}

// Union satisfied by
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple] or
// [magic_transit.PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull].
type PCAPListResponseUnion interface {
	implementsMagicTransitPCAPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull{}),
		},
	)
}

type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON contains the JSON
// metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple]
type pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimple) implementsMagicTransitPCAPListResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                             `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1]
type pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete          PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed            PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleType string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "simple"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull   PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "full"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull struct {
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
	FilterV1 PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullType `json:"type"`
	JSON pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullJSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullJSON contains the JSON
// metadata for the struct [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull]
type pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullJSON struct {
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

func (r *PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFull) implementsMagicTransitPCAPListResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                           `json:"source_port"`
	JSON       pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1]
type pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapListResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "unknown"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "success"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "pending"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning           PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "running"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_pending"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_running"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete          PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "complete"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed            PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "failed"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystem string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullType string

const (
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullType = "simple"
	PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull   PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullType = "full"
)

func (r PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple, PCAPListResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull:
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
	// This field can have the runtime type of
	// [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1],
	// [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1].
	FilterV1 interface{} `json:"filter_v1,required"`
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
	Type  PCAPGetResponseType `json:"type"`
	JSON  pcapGetResponseJSON `json:"-"`
	union PCAPGetResponseUnion
}

// pcapGetResponseJSON contains the JSON metadata for the struct [PCAPGetResponse]
type pcapGetResponseJSON struct {
	FilterV1        apijson.Field
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
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
// Possible runtime types of the union are
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple],
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull].
func (r PCAPGetResponse) AsUnion() PCAPGetResponseUnion {
	return r.union
}

// Union satisfied by
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple] or
// [magic_transit.PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull].
type PCAPGetResponseUnion interface {
	implementsMagicTransitPCAPGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PCAPGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull{}),
		},
	)
}

type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON contains the JSON
// metadata for the struct [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple]
type pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimple) implementsMagicTransitPCAPGetResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                            `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1]
type pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete          PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed            PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus = "failed"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatus) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusUnknown, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusSuccess, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusPending, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusRunning, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionPending, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusConversionRunning, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusComplete, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem = "magic-transit"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleType string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull   PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleType = "full"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleType) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeSimple, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseSimpleTypeFull:
		return true
	}
	return false
}

type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull struct {
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
	FilterV1 PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystem `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullType `json:"type"`
	JSON pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullJSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullJSON contains the JSON
// metadata for the struct [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull]
type pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullJSON struct {
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

func (r *PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullJSON) RawJSON() string {
	return r.raw
}

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFull) implementsMagicTransitPCAPGetResponse() {
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64                                                          `json:"source_port"`
	JSON       pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON `json:"-"`
}

// pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON contains the
// JSON metadata for the struct
// [PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1]
type pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pcapGetResponseMagicVisibilityPCAPsPCAPsResponseFullFilterV1JSON) RawJSON() string {
	return r.raw
}

// The status of the packet capture request.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "unknown"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "success"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "pending"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning           PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "running"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_pending"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "conversion_running"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete          PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "complete"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed            PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus = "failed"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatus) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusUnknown, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusSuccess, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusPending, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusRunning, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionPending, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusConversionRunning, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusComplete, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullStatusFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystem string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystem = "magic-transit"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystem) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullType string

const (
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullType = "simple"
	PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull   PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullType = "full"
)

func (r PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullType) IsKnown() bool {
	switch r {
	case PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullTypeSimple, PCAPGetResponseMagicVisibilityPCAPsPCAPsResponseFullTypeFull:
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
	FilterV1 param.Field[interface{}] `json:"filter_v1,required"`
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
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PCAPNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBody) implementsMagicTransitPCAPNewParamsBodyUnion() {}

// Satisfied by
// [magic_transit.PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimple],
// [magic_transit.PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFull],
// [PCAPNewParamsBody].
type PCAPNewParamsBodyUnion interface {
	implementsMagicTransitPCAPNewParamsBodyUnion()
}

type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleFilterV1] `json:"filter_v1"`
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimple) implementsMagicTransitPCAPNewParamsBodyUnion() {
}

// The system used to collect packet captures.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystem string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystemMagicTransit PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystem = "magic-transit"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleType string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleTypeSimple PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleType = "simple"
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleTypeFull   PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleType = "full"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleType) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleTypeSimple, PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleTypeFull:
		return true
	}
	return false
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleFilterV1 struct {
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

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestSimpleFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystem] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullFilterV1] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFull) implementsMagicTransitPCAPNewParamsBodyUnion() {
}

// The system used to collect packet captures.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystem string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystemMagicTransit PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystem = "magic-transit"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystem) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullSystemMagicTransit:
		return true
	}
	return false
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullType string

const (
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullTypeSimple PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullType = "simple"
	PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullTypeFull   PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullType = "full"
)

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullType) IsKnown() bool {
	switch r {
	case PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullTypeSimple, PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullTypeFull:
		return true
	}
	return false
}

// The packet capture filter. When this field is empty, all packets are captured.
type PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullFilterV1 struct {
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

func (r PCAPNewParamsBodyMagicVisibilityPCAPsPCAPsRequestFullFilterV1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
