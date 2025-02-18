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

// ConnectorTelemetryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorTelemetryService] method instead.
type ConnectorTelemetryService struct {
	Options []option.RequestOption
}

// NewConnectorTelemetryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorTelemetryService(opts ...option.RequestOption) (r *ConnectorTelemetryService) {
	r = &ConnectorTelemetryService{}
	r.Options = opts
	return
}

// List Events
func (r *ConnectorTelemetryService) List(ctx context.Context, connectorID string, params ConnectorTelemetryListParams, opts ...option.RequestOption) (res *ConnectorTelemetryListResponse, err error) {
	var env ConnectorTelemetryListResponseEnvelope
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
func (r *ConnectorTelemetryService) Get(ctx context.Context, connectorID string, eventT float64, eventN float64, query ConnectorTelemetryGetParams, opts ...option.RequestOption) (res *ConnectorTelemetryGetResponse, err error) {
	var env ConnectorTelemetryGetResponseEnvelope
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

type ConnectorTelemetryListResponse struct {
	Count  float64                              `json:"count,required"`
	Items  []ConnectorTelemetryListResponseItem `json:"items,required"`
	Cursor string                               `json:"cursor"`
	JSON   connectorTelemetryListResponseJSON   `json:"-"`
}

// connectorTelemetryListResponseJSON contains the JSON metadata for the struct
// [ConnectorTelemetryListResponse]
type connectorTelemetryListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryListResponseItem struct {
	// Time the Event was collected (seconds since the Unix epoch)
	A float64 `json:"a,required"`
	// Kind
	K string `json:"k,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                                `json:"t,required"`
	JSON connectorTelemetryListResponseItemJSON `json:"-"`
}

// connectorTelemetryListResponseItemJSON contains the JSON metadata for the struct
// [ConnectorTelemetryListResponseItem]
type connectorTelemetryListResponseItemJSON struct {
	A           apijson.Field
	K           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorTelemetryGetResponse struct {
	E ConnectorTelemetryGetResponseE `json:"e,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                           `json:"t,required"`
	JSON connectorTelemetryGetResponseJSON `json:"-"`
}

// connectorTelemetryGetResponseJSON contains the JSON metadata for the struct
// [ConnectorTelemetryGetResponse]
type connectorTelemetryGetResponseJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryGetResponseE struct {
	// Initialized process
	K ConnectorTelemetryGetResponseEK `json:"k,required"`
	// Location of upgrade bundle
	URL   string                             `json:"url"`
	JSON  connectorTelemetryGetResponseEJSON `json:"-"`
	union ConnectorTelemetryGetResponseEUnion
}

// connectorTelemetryGetResponseEJSON contains the JSON metadata for the struct
// [ConnectorTelemetryGetResponseE]
type connectorTelemetryGetResponseEJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectorTelemetryGetResponseEJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectorTelemetryGetResponseE) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectorTelemetryGetResponseE{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectorTelemetryGetResponseEUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEObject],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK].
func (r ConnectorTelemetryGetResponseE) AsUnion() ConnectorTelemetryGetResponseEUnion {
	return r.union
}

// Union satisfied by [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEObject],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK],
// [magic_transit.ConnectorTelemetryGetResponseEK] or
// [magic_transit.ConnectorTelemetryGetResponseEK].
type ConnectorTelemetryGetResponseEUnion interface {
	implementsConnectorTelemetryGetResponseE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectorTelemetryGetResponseEUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectorTelemetryGetResponseEK{}),
		},
	)
}

type ConnectorTelemetryGetResponseEK struct {
	// Initialized process
	K    ConnectorTelemetryGetResponseEKK    `json:"k,required"`
	JSON connectorTelemetryGetResponseEkJSON `json:"-"`
}

// connectorTelemetryGetResponseEkJSON contains the JSON metadata for the struct
// [ConnectorTelemetryGetResponseEK]
type connectorTelemetryGetResponseEkJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponseEK) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseEkJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorTelemetryGetResponseEK) implementsConnectorTelemetryGetResponseE() {}

// Initialized process
type ConnectorTelemetryGetResponseEKK string

const (
	ConnectorTelemetryGetResponseEKKInit ConnectorTelemetryGetResponseEKK = "Init"
)

func (r ConnectorTelemetryGetResponseEKK) IsKnown() bool {
	switch r {
	case ConnectorTelemetryGetResponseEKKInit:
		return true
	}
	return false
}

type ConnectorTelemetryGetResponseEObject struct {
	// Started upgrade
	K ConnectorTelemetryGetResponseEObjectK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                                   `json:"url,required"`
	JSON connectorTelemetryGetResponseEObjectJSON `json:"-"`
}

// connectorTelemetryGetResponseEObjectJSON contains the JSON metadata for the
// struct [ConnectorTelemetryGetResponseEObject]
type connectorTelemetryGetResponseEObjectJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponseEObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseEObjectJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorTelemetryGetResponseEObject) implementsConnectorTelemetryGetResponseE() {}

// Started upgrade
type ConnectorTelemetryGetResponseEObjectK string

const (
	ConnectorTelemetryGetResponseEObjectKStartUpgrade ConnectorTelemetryGetResponseEObjectK = "StartUpgrade"
)

func (r ConnectorTelemetryGetResponseEObjectK) IsKnown() bool {
	switch r {
	case ConnectorTelemetryGetResponseEObjectKStartUpgrade:
		return true
	}
	return false
}

type ConnectorTelemetryListParams struct {
	AccountID param.Field[float64] `path:"account_id,required"`
	From      param.Field[float64] `query:"from,required"`
	To        param.Field[float64] `query:"to,required"`
	Cursor    param.Field[string]  `query:"cursor"`
	Limit     param.Field[float64] `query:"limit"`
}

// URLQuery serializes [ConnectorTelemetryListParams]'s query parameters as
// `url.Values`.
func (r ConnectorTelemetryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectorTelemetryListResponseEnvelope struct {
	Result   ConnectorTelemetryListResponse                   `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	Errors   []ConnectorTelemetryListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorTelemetryListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorTelemetryListResponseEnvelopeJSON       `json:"-"`
}

// connectorTelemetryListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectorTelemetryListResponseEnvelope]
type connectorTelemetryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryListResponseEnvelopeErrors struct {
	Code    float64                                          `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    connectorTelemetryListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorTelemetryListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ConnectorTelemetryListResponseEnvelopeErrors]
type connectorTelemetryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryListResponseEnvelopeMessages struct {
	Code    float64                                            `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    connectorTelemetryListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorTelemetryListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ConnectorTelemetryListResponseEnvelopeMessages]
type connectorTelemetryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryGetParams struct {
	AccountID param.Field[float64] `path:"account_id,required"`
}

type ConnectorTelemetryGetResponseEnvelope struct {
	// Recorded Event
	Result   ConnectorTelemetryGetResponse                   `json:"result,required"`
	Success  bool                                            `json:"success,required"`
	Errors   []ConnectorTelemetryGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorTelemetryGetResponseEnvelopeMessages `json:"messages"`
	JSON     connectorTelemetryGetResponseEnvelopeJSON       `json:"-"`
}

// connectorTelemetryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectorTelemetryGetResponseEnvelope]
type connectorTelemetryGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryGetResponseEnvelopeErrors struct {
	Code    float64                                         `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    connectorTelemetryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorTelemetryGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ConnectorTelemetryGetResponseEnvelopeErrors]
type connectorTelemetryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorTelemetryGetResponseEnvelopeMessages struct {
	Code    float64                                           `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    connectorTelemetryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorTelemetryGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ConnectorTelemetryGetResponseEnvelopeMessages]
type connectorTelemetryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorTelemetryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorTelemetryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
