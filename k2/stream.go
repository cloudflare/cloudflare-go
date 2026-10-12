// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package k2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// StreamService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
	Options       []option.RequestOption
	Subscriptions *StreamSubscriptionService
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r *StreamService) {
	r = &StreamService{}
	r.Options = opts
	r.Subscriptions = NewStreamSubscriptionService(opts...)
	return
}

// Create a new K2 stream. HTTP is disabled when `http` is omitted. Enabled HTTP
// requires authentication and allows all origins unless `authentication` or `cors`
// say otherwise. At least one input must be enabled.
func (r *StreamService) New(ctx context.Context, params StreamNewParams, opts ...option.RequestOption) (res *StreamNewResponse, err error) {
	var env StreamNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update a K2 stream. Omitted `http` settings, such as `authentication` and
// `cors`, keep their current values. Disabling HTTP keeps them, so enabling HTTP
// again restores them. At least one input must remain enabled.
func (r *StreamService) Update(ctx context.Context, streamID string, params StreamUpdateParams, opts ...option.RequestOption) (res *StreamUpdateResponse, err error) {
	var env StreamUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams/%s", params.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List or filter K2 streams in an account.
func (r *StreamService) List(ctx context.Context, params StreamListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[StreamListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List or filter K2 streams in an account.
func (r *StreamService) ListAutoPaging(ctx context.Context, params StreamListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[StreamListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a K2 stream in an account. Deleting a stream that does not exist also
// succeeds.
func (r *StreamService) Delete(ctx context.Context, streamID string, body StreamDeleteParams, opts ...option.RequestOption) (res *StreamDeleteResponse, err error) {
	var env StreamDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams/%s", body.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get K2 stream details.
func (r *StreamService) Get(ctx context.Context, streamID string, query StreamGetParams, opts ...option.RequestOption) (res *StreamGetResponse, err error) {
	var env StreamGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams/%s", query.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type StreamNewResponse struct {
	// Specifies the public ID of the K2 stream.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Indicates the base HTTP endpoint for producing and consuming records.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP       StreamNewResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time             `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the K2 stream.
	Name string `json:"name" api:"required"`
	// Shows the configured record retention period from 1 hour (3600 seconds) to 30
	// days (2592000 seconds), inclusive.
	RetentionSeconds int64                          `json:"retention_seconds" api:"required"`
	WorkerBinding    StreamNewResponseWorkerBinding `json:"worker_binding" api:"required"`
	JSON             streamNewResponseJSON          `json:"-"`
}

// streamNewResponseJSON contains the JSON metadata for the struct
// [StreamNewResponse]
type streamNewResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Endpoint         apijson.Field
	HTTP             apijson.Field
	ModifiedAt       apijson.Field
	Name             apijson.Field
	RetentionSeconds apijson.Field
	WorkerBinding    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseJSON) RawJSON() string {
	return r.raw
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamNewResponseHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled bool `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication bool                      `json:"authentication"`
	CORS           StreamNewResponseHTTPCORS `json:"cors"`
	JSON           streamNewResponseHTTPJSON `json:"-"`
}

// streamNewResponseHTTPJSON contains the JSON metadata for the struct
// [StreamNewResponseHTTP]
type streamNewResponseHTTPJSON struct {
	Enabled        apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamNewResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseHTTPJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins []string                      `json:"origins"`
	JSON    streamNewResponseHttpcorsJSON `json:"-"`
}

// streamNewResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamNewResponseHTTPCORS]
type streamNewResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamNewResponseWorkerBindingEnabled `json:"enabled" api:"required"`
	JSON    streamNewResponseWorkerBindingJSON    `json:"-"`
	union   StreamNewResponseWorkerBindingUnion
}

// streamNewResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamNewResponseWorkerBinding]
type streamNewResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamNewResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

func (r *StreamNewResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	*r = StreamNewResponseWorkerBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamNewResponseWorkerBindingUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamNewResponseWorkerBindingEnabled],
// [StreamNewResponseWorkerBindingEnabled].
func (r StreamNewResponseWorkerBinding) AsUnion() StreamNewResponseWorkerBindingUnion {
	return r.union
}

// Union satisfied by [StreamNewResponseWorkerBindingEnabled] or
// [StreamNewResponseWorkerBindingEnabled].
type StreamNewResponseWorkerBindingUnion interface {
	implementsStreamNewResponseWorkerBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamNewResponseWorkerBindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamNewResponseWorkerBindingEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamNewResponseWorkerBindingEnabled{}),
		},
	)
}

type StreamNewResponseWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamNewResponseWorkerBindingEnabledEnabled `json:"enabled" api:"required"`
	JSON    streamNewResponseWorkerBindingEnabledJSON    `json:"-"`
}

// streamNewResponseWorkerBindingEnabledJSON contains the JSON metadata for the
// struct [StreamNewResponseWorkerBindingEnabled]
type streamNewResponseWorkerBindingEnabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseWorkerBindingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseWorkerBindingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseWorkerBindingEnabled) implementsStreamNewResponseWorkerBinding() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamNewResponseWorkerBindingEnabledEnabled bool

const (
	StreamNewResponseWorkerBindingEnabledEnabledFalse StreamNewResponseWorkerBindingEnabledEnabled = false
)

func (r StreamNewResponseWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamNewResponseWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamUpdateResponse struct {
	// Specifies the public ID of the K2 stream.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Indicates the base HTTP endpoint for producing and consuming records.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP       StreamUpdateResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time                `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the K2 stream.
	Name string `json:"name" api:"required"`
	// Shows the configured record retention period from 1 hour (3600 seconds) to 30
	// days (2592000 seconds), inclusive.
	RetentionSeconds int64                             `json:"retention_seconds" api:"required"`
	WorkerBinding    StreamUpdateResponseWorkerBinding `json:"worker_binding" api:"required"`
	JSON             streamUpdateResponseJSON          `json:"-"`
}

// streamUpdateResponseJSON contains the JSON metadata for the struct
// [StreamUpdateResponse]
type streamUpdateResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Endpoint         apijson.Field
	HTTP             apijson.Field
	ModifiedAt       apijson.Field
	Name             apijson.Field
	RetentionSeconds apijson.Field
	WorkerBinding    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamUpdateResponseHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled bool `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication bool                         `json:"authentication"`
	CORS           StreamUpdateResponseHTTPCORS `json:"cors"`
	JSON           streamUpdateResponseHTTPJSON `json:"-"`
}

// streamUpdateResponseHTTPJSON contains the JSON metadata for the struct
// [StreamUpdateResponseHTTP]
type streamUpdateResponseHTTPJSON struct {
	Enabled        apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamUpdateResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseHTTPJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins []string                         `json:"origins"`
	JSON    streamUpdateResponseHttpcorsJSON `json:"-"`
}

// streamUpdateResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamUpdateResponseHTTPCORS]
type streamUpdateResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamUpdateResponseWorkerBindingEnabled `json:"enabled" api:"required"`
	JSON    streamUpdateResponseWorkerBindingJSON    `json:"-"`
	union   StreamUpdateResponseWorkerBindingUnion
}

// streamUpdateResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamUpdateResponseWorkerBinding]
type streamUpdateResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamUpdateResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

func (r *StreamUpdateResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	*r = StreamUpdateResponseWorkerBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamUpdateResponseWorkerBindingUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StreamUpdateResponseWorkerBindingEnabled],
// [StreamUpdateResponseWorkerBindingEnabled].
func (r StreamUpdateResponseWorkerBinding) AsUnion() StreamUpdateResponseWorkerBindingUnion {
	return r.union
}

// Union satisfied by [StreamUpdateResponseWorkerBindingEnabled] or
// [StreamUpdateResponseWorkerBindingEnabled].
type StreamUpdateResponseWorkerBindingUnion interface {
	implementsStreamUpdateResponseWorkerBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamUpdateResponseWorkerBindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamUpdateResponseWorkerBindingEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamUpdateResponseWorkerBindingEnabled{}),
		},
	)
}

type StreamUpdateResponseWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamUpdateResponseWorkerBindingEnabledEnabled `json:"enabled" api:"required"`
	JSON    streamUpdateResponseWorkerBindingEnabledJSON    `json:"-"`
}

// streamUpdateResponseWorkerBindingEnabledJSON contains the JSON metadata for the
// struct [StreamUpdateResponseWorkerBindingEnabled]
type streamUpdateResponseWorkerBindingEnabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseWorkerBindingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseWorkerBindingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r StreamUpdateResponseWorkerBindingEnabled) implementsStreamUpdateResponseWorkerBinding() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamUpdateResponseWorkerBindingEnabledEnabled bool

const (
	StreamUpdateResponseWorkerBindingEnabledEnabledFalse StreamUpdateResponseWorkerBindingEnabledEnabled = false
)

func (r StreamUpdateResponseWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamUpdateResponseWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamListResponse struct {
	// Specifies the public ID of the K2 stream.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Indicates the base HTTP endpoint for producing and consuming records.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP       StreamListResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time              `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the K2 stream.
	Name string `json:"name" api:"required"`
	// Shows the configured record retention period from 1 hour (3600 seconds) to 30
	// days (2592000 seconds), inclusive.
	RetentionSeconds int64                           `json:"retention_seconds" api:"required"`
	WorkerBinding    StreamListResponseWorkerBinding `json:"worker_binding" api:"required"`
	JSON             streamListResponseJSON          `json:"-"`
}

// streamListResponseJSON contains the JSON metadata for the struct
// [StreamListResponse]
type streamListResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Endpoint         apijson.Field
	HTTP             apijson.Field
	ModifiedAt       apijson.Field
	Name             apijson.Field
	RetentionSeconds apijson.Field
	WorkerBinding    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseJSON) RawJSON() string {
	return r.raw
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamListResponseHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled bool `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication bool                       `json:"authentication"`
	CORS           StreamListResponseHTTPCORS `json:"cors"`
	JSON           streamListResponseHTTPJSON `json:"-"`
}

// streamListResponseHTTPJSON contains the JSON metadata for the struct
// [StreamListResponseHTTP]
type streamListResponseHTTPJSON struct {
	Enabled        apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamListResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseHTTPJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins []string                       `json:"origins"`
	JSON    streamListResponseHttpcorsJSON `json:"-"`
}

// streamListResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamListResponseHTTPCORS]
type streamListResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamListResponseWorkerBindingEnabled `json:"enabled" api:"required"`
	JSON    streamListResponseWorkerBindingJSON    `json:"-"`
	union   StreamListResponseWorkerBindingUnion
}

// streamListResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamListResponseWorkerBinding]
type streamListResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamListResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

func (r *StreamListResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	*r = StreamListResponseWorkerBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamListResponseWorkerBindingUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [StreamListResponseWorkerBindingEnabled],
// [StreamListResponseWorkerBindingEnabled].
func (r StreamListResponseWorkerBinding) AsUnion() StreamListResponseWorkerBindingUnion {
	return r.union
}

// Union satisfied by [StreamListResponseWorkerBindingEnabled] or
// [StreamListResponseWorkerBindingEnabled].
type StreamListResponseWorkerBindingUnion interface {
	implementsStreamListResponseWorkerBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamListResponseWorkerBindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamListResponseWorkerBindingEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamListResponseWorkerBindingEnabled{}),
		},
	)
}

type StreamListResponseWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamListResponseWorkerBindingEnabledEnabled `json:"enabled" api:"required"`
	JSON    streamListResponseWorkerBindingEnabledJSON    `json:"-"`
}

// streamListResponseWorkerBindingEnabledJSON contains the JSON metadata for the
// struct [StreamListResponseWorkerBindingEnabled]
type streamListResponseWorkerBindingEnabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseWorkerBindingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseWorkerBindingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseWorkerBindingEnabled) implementsStreamListResponseWorkerBinding() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamListResponseWorkerBindingEnabledEnabled bool

const (
	StreamListResponseWorkerBindingEnabledEnabledFalse StreamListResponseWorkerBindingEnabledEnabled = false
)

func (r StreamListResponseWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamListResponseWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamDeleteResponse = interface{}

type StreamGetResponse struct {
	// Specifies the public ID of the K2 stream.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Indicates the base HTTP endpoint for producing and consuming records.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP       StreamGetResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time             `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the K2 stream.
	Name string `json:"name" api:"required"`
	// Shows the configured record retention period from 1 hour (3600 seconds) to 30
	// days (2592000 seconds), inclusive.
	RetentionSeconds int64                          `json:"retention_seconds" api:"required"`
	WorkerBinding    StreamGetResponseWorkerBinding `json:"worker_binding" api:"required"`
	JSON             streamGetResponseJSON          `json:"-"`
}

// streamGetResponseJSON contains the JSON metadata for the struct
// [StreamGetResponse]
type streamGetResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Endpoint         apijson.Field
	HTTP             apijson.Field
	ModifiedAt       apijson.Field
	Name             apijson.Field
	RetentionSeconds apijson.Field
	WorkerBinding    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StreamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseJSON) RawJSON() string {
	return r.raw
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamGetResponseHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled bool `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication bool                      `json:"authentication"`
	CORS           StreamGetResponseHTTPCORS `json:"cors"`
	JSON           streamGetResponseHTTPJSON `json:"-"`
}

// streamGetResponseHTTPJSON contains the JSON metadata for the struct
// [StreamGetResponseHTTP]
type streamGetResponseHTTPJSON struct {
	Enabled        apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamGetResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseHTTPJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins []string                      `json:"origins"`
	JSON    streamGetResponseHttpcorsJSON `json:"-"`
}

// streamGetResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamGetResponseHTTPCORS]
type streamGetResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamGetResponseWorkerBindingEnabled `json:"enabled" api:"required"`
	JSON    streamGetResponseWorkerBindingJSON    `json:"-"`
	union   StreamGetResponseWorkerBindingUnion
}

// streamGetResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamGetResponseWorkerBinding]
type streamGetResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamGetResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

func (r *StreamGetResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	*r = StreamGetResponseWorkerBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamGetResponseWorkerBindingUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamGetResponseWorkerBindingEnabled],
// [StreamGetResponseWorkerBindingEnabled].
func (r StreamGetResponseWorkerBinding) AsUnion() StreamGetResponseWorkerBindingUnion {
	return r.union
}

// Union satisfied by [StreamGetResponseWorkerBindingEnabled] or
// [StreamGetResponseWorkerBindingEnabled].
type StreamGetResponseWorkerBindingUnion interface {
	implementsStreamGetResponseWorkerBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamGetResponseWorkerBindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamGetResponseWorkerBindingEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(StreamGetResponseWorkerBindingEnabled{}),
		},
	)
}

type StreamGetResponseWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled StreamGetResponseWorkerBindingEnabledEnabled `json:"enabled" api:"required"`
	JSON    streamGetResponseWorkerBindingEnabledJSON    `json:"-"`
}

// streamGetResponseWorkerBindingEnabledJSON contains the JSON metadata for the
// struct [StreamGetResponseWorkerBindingEnabled]
type streamGetResponseWorkerBindingEnabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseWorkerBindingEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseWorkerBindingEnabledJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseWorkerBindingEnabled) implementsStreamGetResponseWorkerBinding() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamGetResponseWorkerBindingEnabledEnabled bool

const (
	StreamGetResponseWorkerBindingEnabledEnabledFalse StreamGetResponseWorkerBindingEnabledEnabled = false
)

func (r StreamGetResponseWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamGetResponseWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamNewParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specifies the name of the K2 stream.
	Name param.Field[string] `json:"name" api:"required"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP param.Field[StreamNewParamsHTTP] `json:"http"`
	// Sets the record retention period from 1 hour (3600 seconds) to 30 days (2592000
	// seconds), inclusive.
	RetentionSeconds param.Field[int64]                             `json:"retention_seconds"`
	WorkerBinding    param.Field[StreamNewParamsWorkerBindingUnion] `json:"worker_binding"`
}

func (r StreamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamNewParamsHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication param.Field[bool]                    `json:"authentication"`
	CORS           param.Field[StreamNewParamsHTTPCORS] `json:"cors"`
}

func (r StreamNewParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins param.Field[[]string] `json:"origins"`
}

func (r StreamNewParamsHTTPCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled param.Field[StreamNewParamsWorkerBindingEnabled] `json:"enabled" api:"required"`
}

func (r StreamNewParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsWorkerBinding) implementsStreamNewParamsWorkerBindingUnion() {}

// Satisfied by [k2.StreamNewParamsWorkerBindingEnabled],
// [k2.StreamNewParamsWorkerBindingEnabled], [StreamNewParamsWorkerBinding].
type StreamNewParamsWorkerBindingUnion interface {
	implementsStreamNewParamsWorkerBindingUnion()
}

type StreamNewParamsWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled param.Field[StreamNewParamsWorkerBindingEnabledEnabled] `json:"enabled" api:"required"`
}

func (r StreamNewParamsWorkerBindingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsWorkerBindingEnabled) implementsStreamNewParamsWorkerBindingUnion() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamNewParamsWorkerBindingEnabledEnabled bool

const (
	StreamNewParamsWorkerBindingEnabledEnabledFalse StreamNewParamsWorkerBindingEnabledEnabled = false
)

func (r StreamNewParamsWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamNewParamsWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamNewResponseEnvelope struct {
	Errors   []StreamNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []StreamNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   StreamNewResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success" api:"required"`
	JSON    streamNewResponseEnvelopeJSON `json:"-"`
}

// streamNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamNewResponseEnvelope]
type streamNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseEnvelopeErrors struct {
	Message string                              `json:"message" api:"required"`
	Code    int64                               `json:"code"`
	JSON    streamNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamNewResponseEnvelopeErrors]
type streamNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseEnvelopeMessages struct {
	Message string                                `json:"message" api:"required"`
	Code    int64                                 `json:"code"`
	JSON    streamNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [StreamNewResponseEnvelopeMessages]
type streamNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
	// so enabling it again restores them.
	HTTP param.Field[StreamUpdateParamsHTTP] `json:"http"`
	// Sets the record retention period from 1 hour (3600 seconds) to 30 days (2592000
	// seconds), inclusive.
	RetentionSeconds param.Field[int64]                                `json:"retention_seconds"`
	WorkerBinding    param.Field[StreamUpdateParamsWorkerBindingUnion] `json:"worker_binding"`
}

func (r StreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the HTTP endpoint. Disabling HTTP keeps `authentication` and `cors`,
// so enabling it again restores them.
type StreamUpdateParamsHTTP struct {
	// Indicates whether the HTTP endpoint accepts records.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Indicates whether the HTTP endpoint requires an API token with K2 produce
	// permission. When false, the endpoint accepts unauthenticated records. Defaults
	// to true when HTTP is enabled without a stored value.
	Authentication param.Field[bool]                       `json:"authentication"`
	CORS           param.Field[StreamUpdateParamsHTTPCORS] `json:"cors"`
}

func (r StreamUpdateParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateParamsHTTPCORS struct {
	// Allows browser requests from these HTTP or HTTPS origins. Use a wildcard only as
	// the sole origin. An empty list blocks cross-origin browser requests. Defaults to
	// `['*']` when HTTP is enabled without stored origins.
	Origins param.Field[[]string] `json:"origins"`
}

func (r StreamUpdateParamsHTTPCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateParamsWorkerBinding struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled param.Field[StreamUpdateParamsWorkerBindingEnabled] `json:"enabled" api:"required"`
}

func (r StreamUpdateParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamUpdateParamsWorkerBinding) implementsStreamUpdateParamsWorkerBindingUnion() {}

// Satisfied by [k2.StreamUpdateParamsWorkerBindingEnabled],
// [k2.StreamUpdateParamsWorkerBindingEnabled], [StreamUpdateParamsWorkerBinding].
type StreamUpdateParamsWorkerBindingUnion interface {
	implementsStreamUpdateParamsWorkerBindingUnion()
}

type StreamUpdateParamsWorkerBindingEnabled struct {
	// Indicates whether Workers bindings can produce records to the stream.
	Enabled param.Field[StreamUpdateParamsWorkerBindingEnabledEnabled] `json:"enabled" api:"required"`
}

func (r StreamUpdateParamsWorkerBindingEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamUpdateParamsWorkerBindingEnabled) implementsStreamUpdateParamsWorkerBindingUnion() {}

// Indicates whether Workers bindings can produce records to the stream.
type StreamUpdateParamsWorkerBindingEnabledEnabled bool

const (
	StreamUpdateParamsWorkerBindingEnabledEnabledFalse StreamUpdateParamsWorkerBindingEnabledEnabled = false
)

func (r StreamUpdateParamsWorkerBindingEnabledEnabled) IsKnown() bool {
	switch r {
	case StreamUpdateParamsWorkerBindingEnabledEnabledFalse:
		return true
	}
	return false
}

type StreamUpdateResponseEnvelope struct {
	Errors   []StreamUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []StreamUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   StreamUpdateResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                             `json:"success" api:"required"`
	JSON    streamUpdateResponseEnvelopeJSON `json:"-"`
}

// streamUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamUpdateResponseEnvelope]
type streamUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseEnvelopeErrors struct {
	Message string                                 `json:"message" api:"required"`
	Code    int64                                  `json:"code"`
	JSON    streamUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamUpdateResponseEnvelopeErrors]
type streamUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseEnvelopeMessages struct {
	Message string                                   `json:"message" api:"required"`
	Code    int64                                    `json:"code"`
	JSON    streamUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamUpdateResponseEnvelopeMessages]
type streamUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamListParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters K2 streams by name using a case-insensitive substring.
	Name    param.Field[string] `query:"name"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [StreamListParams]'s query parameters as `url.Values`.
func (r StreamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StreamDeleteParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type StreamDeleteResponseEnvelope struct {
	Errors   []StreamDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []StreamDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   StreamDeleteResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                             `json:"success" api:"required"`
	JSON    streamDeleteResponseEnvelopeJSON `json:"-"`
}

// streamDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDeleteResponseEnvelope]
type streamDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamDeleteResponseEnvelopeErrors struct {
	Message string                                 `json:"message" api:"required"`
	Code    int64                                  `json:"code"`
	JSON    streamDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamDeleteResponseEnvelopeErrors]
type streamDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamDeleteResponseEnvelopeMessages struct {
	Message string                                   `json:"message" api:"required"`
	Code    int64                                    `json:"code"`
	JSON    streamDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamDeleteResponseEnvelopeMessages]
type streamDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type StreamGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type StreamGetResponseEnvelope struct {
	Errors   []StreamGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []StreamGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   StreamGetResponse                   `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success" api:"required"`
	JSON    streamGetResponseEnvelopeJSON `json:"-"`
}

// streamGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelope]
type streamGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeErrors struct {
	Message string                              `json:"message" api:"required"`
	Code    int64                               `json:"code"`
	JSON    streamGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelopeErrors]
type streamGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseEnvelopeMessages struct {
	Message string                                `json:"message" api:"required"`
	Code    int64                                 `json:"code"`
	JSON    streamGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelopeMessages]
type streamGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
