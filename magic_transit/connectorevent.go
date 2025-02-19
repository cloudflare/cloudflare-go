// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
)

// ConnectorEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorEventService] method instead.
type ConnectorEventService struct {
	Options []option.RequestOption
}

// NewConnectorEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorEventService(opts ...option.RequestOption) (r *ConnectorEventService) {
	r = &ConnectorEventService{}
	r.Options = opts
	return
}

// List Events
func (r *ConnectorEventService) List(ctx context.Context, connectorID string, params ConnectorEventListParams, opts ...option.RequestOption) (res *ConnectorEventListResponse, err error) {
	var env ConnectorEventListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/magic/connectors/%s/telemetry/events", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Event
func (r *ConnectorEventService) Get(ctx context.Context, connectorID string, eventT float64, eventN float64, query ConnectorEventGetParams, opts ...option.RequestOption) (res *ConnectorEventGetResponse, err error) {
	var env ConnectorEventGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/magic/connectors/%s/telemetry/events/%v.%v", query.AccountID, connectorID, eventT, eventN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectorEventListResponse struct {
	Count  float64                          `json:"count,required"`
	Items  []ConnectorEventListResponseItem `json:"items,required"`
	Cursor string                           `json:"cursor"`
	JSON   connectorEventListResponseJSON   `json:"-"`
}

// connectorEventListResponseJSON contains the JSON metadata for the struct
// [ConnectorEventListResponse]
type connectorEventListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseItem struct {
	// Time the Event was collected (seconds since the Unix epoch)
	A float64 `json:"a,required"`
	// Kind
	K string `json:"k,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                            `json:"t,required"`
	JSON connectorEventListResponseItemJSON `json:"-"`
}

// connectorEventListResponseItemJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseItem]
type connectorEventListResponseItemJSON struct {
	A           apijson.Field
	K           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorEventGetResponse struct {
	E ConnectorEventGetResponseE `json:"e,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                       `json:"t,required"`
	JSON connectorEventGetResponseJSON `json:"-"`
}

// connectorEventGetResponseJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponse]
type connectorEventGetResponseJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseE struct {
	// Initialized process
	K ConnectorEventGetResponseEK `json:"k,required"`
	// Location of upgrade bundle
	URL   string                         `json:"url"`
	JSON  connectorEventGetResponseEJSON `json:"-"`
	union ConnectorEventGetResponseEUnion
}

// connectorEventGetResponseEJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseE]
type connectorEventGetResponseEJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectorEventGetResponseEJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectorEventGetResponseE) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectorEventGetResponseE{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectorEventGetResponseEUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEObject],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK].
func (r ConnectorEventGetResponseE) AsUnion() ConnectorEventGetResponseEUnion {
	return r.union
}

// Union satisfied by [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEObject],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK],
// [magic_transit.ConnectorEventGetResponseEK] or
// [magic_transit.ConnectorEventGetResponseEK].
type ConnectorEventGetResponseEUnion interface {
	implementsConnectorEventGetResponseE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectorEventGetResponseEUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorEventGetResponseEK{}),
		},
	)
}

type ConnectorEventGetResponseEK struct {
	// Initialized process
	K    ConnectorEventGetResponseEKK    `json:"k,required"`
	JSON connectorEventGetResponseEkJSON `json:"-"`
}

// connectorEventGetResponseEkJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEK]
type connectorEventGetResponseEkJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEK) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEkJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEK) implementsConnectorEventGetResponseE() {}

// Initialized process
type ConnectorEventGetResponseEKK string

const (
	ConnectorEventGetResponseEKKInit ConnectorEventGetResponseEKK = "Init"
)

func (r ConnectorEventGetResponseEKK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEKKInit:
		return true
	}
	return false
}

type ConnectorEventGetResponseEObject struct {
	// Started upgrade
	K ConnectorEventGetResponseEObjectK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                               `json:"url,required"`
	JSON connectorEventGetResponseEObjectJSON `json:"-"`
}

// connectorEventGetResponseEObjectJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEObject]
type connectorEventGetResponseEObjectJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEObjectJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEObject) implementsConnectorEventGetResponseE() {}

// Started upgrade
type ConnectorEventGetResponseEObjectK string

const (
	ConnectorEventGetResponseEObjectKStartUpgrade ConnectorEventGetResponseEObjectK = "StartUpgrade"
)

func (r ConnectorEventGetResponseEObjectK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEObjectKStartUpgrade:
		return true
	}
	return false
}

type ConnectorEventListParams struct {
	AccountID param.Field[float64] `path:"account_id,required"`
	From      param.Field[float64] `query:"from,required"`
	To        param.Field[float64] `query:"to,required"`
	Cursor    param.Field[string]  `query:"cursor"`
	Limit     param.Field[float64] `query:"limit"`
}

// URLQuery serializes [ConnectorEventListParams]'s query parameters as
// `url.Values`.
func (r ConnectorEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectorEventListResponseEnvelope struct {
	Result   ConnectorEventListResponse                   `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	Errors   []ConnectorEventListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventListResponseEnvelopeJSON       `json:"-"`
}

// connectorEventListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseEnvelope]
type connectorEventListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    connectorEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventListResponseEnvelopeErrors]
type connectorEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    connectorEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ConnectorEventListResponseEnvelopeMessages]
type connectorEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetParams struct {
	AccountID param.Field[float64] `path:"account_id,required"`
}

type ConnectorEventGetResponseEnvelope struct {
	// Recorded Event
	Result   ConnectorEventGetResponse                   `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	Errors   []ConnectorEventGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventGetResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventGetResponseEnvelopeJSON       `json:"-"`
}

// connectorEventGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEnvelope]
type connectorEventGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeErrors struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    connectorEventGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeErrors]
type connectorEventGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeMessages struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    connectorEventGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeMessages]
type connectorEventGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
