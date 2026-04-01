// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// OperationLabelService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationLabelService] method instead.
type OperationLabelService struct {
	Options []option.RequestOption
}

// NewOperationLabelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationLabelService(opts ...option.RequestOption) (r *OperationLabelService) {
	r = &OperationLabelService{}
	r.Options = opts
	return
}

// Attach label(s) on an operation in endpoint management
func (r *OperationLabelService) New(ctx context.Context, operationID string, params OperationLabelNewParams, opts ...option.RequestOption) (res *OperationLabelNewResponse, err error) {
	var env OperationLabelNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/labels", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replace label(s) on an operation in endpoint management
func (r *OperationLabelService) Update(ctx context.Context, operationID string, params OperationLabelUpdateParams, opts ...option.RequestOption) (res *OperationLabelUpdateResponse, err error) {
	var env OperationLabelUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/labels", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Remove label(s) on an operation in endpoint management
func (r *OperationLabelService) Delete(ctx context.Context, operationID string, body OperationLabelDeleteParams, opts ...option.RequestOption) (res *OperationLabelDeleteResponse, err error) {
	var env OperationLabelDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/labels", body.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Bulk attach label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkNew(ctx context.Context, params OperationLabelBulkNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[OperationLabelBulkNewResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/labels", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Bulk attach label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkNewAutoPaging(ctx context.Context, params OperationLabelBulkNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OperationLabelBulkNewResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkNew(ctx, params, opts...))
}

// Bulk remove label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkDelete(ctx context.Context, body OperationLabelBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[OperationLabelBulkDeleteResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/labels", body.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Bulk remove label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkDeleteAutoPaging(ctx context.Context, body OperationLabelBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OperationLabelBulkDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, body, opts...))
}

// Bulk replace label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkUpdate(ctx context.Context, params OperationLabelBulkUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[OperationLabelBulkUpdateResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/labels", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Bulk replace label(s) on operation(s) in endpoint management
func (r *OperationLabelService) BulkUpdateAutoPaging(ctx context.Context, params OperationLabelBulkUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OperationLabelBulkUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkUpdate(ctx, params, opts...))
}

type OperationLabelNewResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelNewResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                           `json:"operation_id" api:"required"`
	Labels      []OperationLabelNewResponseLabel `json:"labels"`
	JSON        operationLabelNewResponseJSON    `json:"-"`
}

// operationLabelNewResponseJSON contains the JSON metadata for the struct
// [OperationLabelNewResponse]
type operationLabelNewResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelNewResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelNewResponseMethod string

const (
	OperationLabelNewResponseMethodGet     OperationLabelNewResponseMethod = "GET"
	OperationLabelNewResponseMethodPost    OperationLabelNewResponseMethod = "POST"
	OperationLabelNewResponseMethodHead    OperationLabelNewResponseMethod = "HEAD"
	OperationLabelNewResponseMethodOptions OperationLabelNewResponseMethod = "OPTIONS"
	OperationLabelNewResponseMethodPut     OperationLabelNewResponseMethod = "PUT"
	OperationLabelNewResponseMethodDelete  OperationLabelNewResponseMethod = "DELETE"
	OperationLabelNewResponseMethodConnect OperationLabelNewResponseMethod = "CONNECT"
	OperationLabelNewResponseMethodPatch   OperationLabelNewResponseMethod = "PATCH"
	OperationLabelNewResponseMethodTrace   OperationLabelNewResponseMethod = "TRACE"
)

func (r OperationLabelNewResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelNewResponseMethodGet, OperationLabelNewResponseMethodPost, OperationLabelNewResponseMethodHead, OperationLabelNewResponseMethodOptions, OperationLabelNewResponseMethodPut, OperationLabelNewResponseMethodDelete, OperationLabelNewResponseMethodConnect, OperationLabelNewResponseMethodPatch, OperationLabelNewResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelNewResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelNewResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelNewResponseLabelJSON    `json:"-"`
}

// operationLabelNewResponseLabelJSON contains the JSON metadata for the struct
// [OperationLabelNewResponseLabel]
type operationLabelNewResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelNewResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelNewResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelNewResponseLabelsSource string

const (
	OperationLabelNewResponseLabelsSourceUser    OperationLabelNewResponseLabelsSource = "user"
	OperationLabelNewResponseLabelsSourceManaged OperationLabelNewResponseLabelsSource = "managed"
)

func (r OperationLabelNewResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelNewResponseLabelsSourceUser, OperationLabelNewResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelUpdateResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelUpdateResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                              `json:"operation_id" api:"required"`
	Labels      []OperationLabelUpdateResponseLabel `json:"labels"`
	JSON        operationLabelUpdateResponseJSON    `json:"-"`
}

// operationLabelUpdateResponseJSON contains the JSON metadata for the struct
// [OperationLabelUpdateResponse]
type operationLabelUpdateResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelUpdateResponseMethod string

const (
	OperationLabelUpdateResponseMethodGet     OperationLabelUpdateResponseMethod = "GET"
	OperationLabelUpdateResponseMethodPost    OperationLabelUpdateResponseMethod = "POST"
	OperationLabelUpdateResponseMethodHead    OperationLabelUpdateResponseMethod = "HEAD"
	OperationLabelUpdateResponseMethodOptions OperationLabelUpdateResponseMethod = "OPTIONS"
	OperationLabelUpdateResponseMethodPut     OperationLabelUpdateResponseMethod = "PUT"
	OperationLabelUpdateResponseMethodDelete  OperationLabelUpdateResponseMethod = "DELETE"
	OperationLabelUpdateResponseMethodConnect OperationLabelUpdateResponseMethod = "CONNECT"
	OperationLabelUpdateResponseMethodPatch   OperationLabelUpdateResponseMethod = "PATCH"
	OperationLabelUpdateResponseMethodTrace   OperationLabelUpdateResponseMethod = "TRACE"
)

func (r OperationLabelUpdateResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelUpdateResponseMethodGet, OperationLabelUpdateResponseMethodPost, OperationLabelUpdateResponseMethodHead, OperationLabelUpdateResponseMethodOptions, OperationLabelUpdateResponseMethodPut, OperationLabelUpdateResponseMethodDelete, OperationLabelUpdateResponseMethodConnect, OperationLabelUpdateResponseMethodPatch, OperationLabelUpdateResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelUpdateResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelUpdateResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelUpdateResponseLabelJSON    `json:"-"`
}

// operationLabelUpdateResponseLabelJSON contains the JSON metadata for the struct
// [OperationLabelUpdateResponseLabel]
type operationLabelUpdateResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelUpdateResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelUpdateResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelUpdateResponseLabelsSource string

const (
	OperationLabelUpdateResponseLabelsSourceUser    OperationLabelUpdateResponseLabelsSource = "user"
	OperationLabelUpdateResponseLabelsSourceManaged OperationLabelUpdateResponseLabelsSource = "managed"
)

func (r OperationLabelUpdateResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelUpdateResponseLabelsSourceUser, OperationLabelUpdateResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelDeleteResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelDeleteResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                              `json:"operation_id" api:"required"`
	Labels      []OperationLabelDeleteResponseLabel `json:"labels"`
	JSON        operationLabelDeleteResponseJSON    `json:"-"`
}

// operationLabelDeleteResponseJSON contains the JSON metadata for the struct
// [OperationLabelDeleteResponse]
type operationLabelDeleteResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelDeleteResponseMethod string

const (
	OperationLabelDeleteResponseMethodGet     OperationLabelDeleteResponseMethod = "GET"
	OperationLabelDeleteResponseMethodPost    OperationLabelDeleteResponseMethod = "POST"
	OperationLabelDeleteResponseMethodHead    OperationLabelDeleteResponseMethod = "HEAD"
	OperationLabelDeleteResponseMethodOptions OperationLabelDeleteResponseMethod = "OPTIONS"
	OperationLabelDeleteResponseMethodPut     OperationLabelDeleteResponseMethod = "PUT"
	OperationLabelDeleteResponseMethodDelete  OperationLabelDeleteResponseMethod = "DELETE"
	OperationLabelDeleteResponseMethodConnect OperationLabelDeleteResponseMethod = "CONNECT"
	OperationLabelDeleteResponseMethodPatch   OperationLabelDeleteResponseMethod = "PATCH"
	OperationLabelDeleteResponseMethodTrace   OperationLabelDeleteResponseMethod = "TRACE"
)

func (r OperationLabelDeleteResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelDeleteResponseMethodGet, OperationLabelDeleteResponseMethodPost, OperationLabelDeleteResponseMethodHead, OperationLabelDeleteResponseMethodOptions, OperationLabelDeleteResponseMethodPut, OperationLabelDeleteResponseMethodDelete, OperationLabelDeleteResponseMethodConnect, OperationLabelDeleteResponseMethodPatch, OperationLabelDeleteResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelDeleteResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelDeleteResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelDeleteResponseLabelJSON    `json:"-"`
}

// operationLabelDeleteResponseLabelJSON contains the JSON metadata for the struct
// [OperationLabelDeleteResponseLabel]
type operationLabelDeleteResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelDeleteResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelDeleteResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelDeleteResponseLabelsSource string

const (
	OperationLabelDeleteResponseLabelsSourceUser    OperationLabelDeleteResponseLabelsSource = "user"
	OperationLabelDeleteResponseLabelsSourceManaged OperationLabelDeleteResponseLabelsSource = "managed"
)

func (r OperationLabelDeleteResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelDeleteResponseLabelsSourceUser, OperationLabelDeleteResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelBulkNewResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelBulkNewResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                               `json:"operation_id" api:"required"`
	Labels      []OperationLabelBulkNewResponseLabel `json:"labels"`
	JSON        operationLabelBulkNewResponseJSON    `json:"-"`
}

// operationLabelBulkNewResponseJSON contains the JSON metadata for the struct
// [OperationLabelBulkNewResponse]
type operationLabelBulkNewResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelBulkNewResponseMethod string

const (
	OperationLabelBulkNewResponseMethodGet     OperationLabelBulkNewResponseMethod = "GET"
	OperationLabelBulkNewResponseMethodPost    OperationLabelBulkNewResponseMethod = "POST"
	OperationLabelBulkNewResponseMethodHead    OperationLabelBulkNewResponseMethod = "HEAD"
	OperationLabelBulkNewResponseMethodOptions OperationLabelBulkNewResponseMethod = "OPTIONS"
	OperationLabelBulkNewResponseMethodPut     OperationLabelBulkNewResponseMethod = "PUT"
	OperationLabelBulkNewResponseMethodDelete  OperationLabelBulkNewResponseMethod = "DELETE"
	OperationLabelBulkNewResponseMethodConnect OperationLabelBulkNewResponseMethod = "CONNECT"
	OperationLabelBulkNewResponseMethodPatch   OperationLabelBulkNewResponseMethod = "PATCH"
	OperationLabelBulkNewResponseMethodTrace   OperationLabelBulkNewResponseMethod = "TRACE"
)

func (r OperationLabelBulkNewResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelBulkNewResponseMethodGet, OperationLabelBulkNewResponseMethodPost, OperationLabelBulkNewResponseMethodHead, OperationLabelBulkNewResponseMethodOptions, OperationLabelBulkNewResponseMethodPut, OperationLabelBulkNewResponseMethodDelete, OperationLabelBulkNewResponseMethodConnect, OperationLabelBulkNewResponseMethodPatch, OperationLabelBulkNewResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelBulkNewResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelBulkNewResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelBulkNewResponseLabelJSON    `json:"-"`
}

// operationLabelBulkNewResponseLabelJSON contains the JSON metadata for the struct
// [OperationLabelBulkNewResponseLabel]
type operationLabelBulkNewResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkNewResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkNewResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelBulkNewResponseLabelsSource string

const (
	OperationLabelBulkNewResponseLabelsSourceUser    OperationLabelBulkNewResponseLabelsSource = "user"
	OperationLabelBulkNewResponseLabelsSourceManaged OperationLabelBulkNewResponseLabelsSource = "managed"
)

func (r OperationLabelBulkNewResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelBulkNewResponseLabelsSourceUser, OperationLabelBulkNewResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelBulkDeleteResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelBulkDeleteResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                                  `json:"operation_id" api:"required"`
	Labels      []OperationLabelBulkDeleteResponseLabel `json:"labels"`
	JSON        operationLabelBulkDeleteResponseJSON    `json:"-"`
}

// operationLabelBulkDeleteResponseJSON contains the JSON metadata for the struct
// [OperationLabelBulkDeleteResponse]
type operationLabelBulkDeleteResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelBulkDeleteResponseMethod string

const (
	OperationLabelBulkDeleteResponseMethodGet     OperationLabelBulkDeleteResponseMethod = "GET"
	OperationLabelBulkDeleteResponseMethodPost    OperationLabelBulkDeleteResponseMethod = "POST"
	OperationLabelBulkDeleteResponseMethodHead    OperationLabelBulkDeleteResponseMethod = "HEAD"
	OperationLabelBulkDeleteResponseMethodOptions OperationLabelBulkDeleteResponseMethod = "OPTIONS"
	OperationLabelBulkDeleteResponseMethodPut     OperationLabelBulkDeleteResponseMethod = "PUT"
	OperationLabelBulkDeleteResponseMethodDelete  OperationLabelBulkDeleteResponseMethod = "DELETE"
	OperationLabelBulkDeleteResponseMethodConnect OperationLabelBulkDeleteResponseMethod = "CONNECT"
	OperationLabelBulkDeleteResponseMethodPatch   OperationLabelBulkDeleteResponseMethod = "PATCH"
	OperationLabelBulkDeleteResponseMethodTrace   OperationLabelBulkDeleteResponseMethod = "TRACE"
)

func (r OperationLabelBulkDeleteResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelBulkDeleteResponseMethodGet, OperationLabelBulkDeleteResponseMethodPost, OperationLabelBulkDeleteResponseMethodHead, OperationLabelBulkDeleteResponseMethodOptions, OperationLabelBulkDeleteResponseMethodPut, OperationLabelBulkDeleteResponseMethodDelete, OperationLabelBulkDeleteResponseMethodConnect, OperationLabelBulkDeleteResponseMethodPatch, OperationLabelBulkDeleteResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelBulkDeleteResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelBulkDeleteResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelBulkDeleteResponseLabelJSON    `json:"-"`
}

// operationLabelBulkDeleteResponseLabelJSON contains the JSON metadata for the
// struct [OperationLabelBulkDeleteResponseLabel]
type operationLabelBulkDeleteResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkDeleteResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkDeleteResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelBulkDeleteResponseLabelsSource string

const (
	OperationLabelBulkDeleteResponseLabelsSourceUser    OperationLabelBulkDeleteResponseLabelsSource = "user"
	OperationLabelBulkDeleteResponseLabelsSourceManaged OperationLabelBulkDeleteResponseLabelsSource = "managed"
)

func (r OperationLabelBulkDeleteResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelBulkDeleteResponseLabelsSourceUser, OperationLabelBulkDeleteResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelBulkUpdateResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint" api:"required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host" api:"required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationLabelBulkUpdateResponseMethod `json:"method" api:"required"`
	// UUID.
	OperationID string                                  `json:"operation_id" api:"required"`
	Labels      []OperationLabelBulkUpdateResponseLabel `json:"labels"`
	JSON        operationLabelBulkUpdateResponseJSON    `json:"-"`
}

// operationLabelBulkUpdateResponseJSON contains the JSON metadata for the struct
// [OperationLabelBulkUpdateResponse]
type operationLabelBulkUpdateResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationLabelBulkUpdateResponseMethod string

const (
	OperationLabelBulkUpdateResponseMethodGet     OperationLabelBulkUpdateResponseMethod = "GET"
	OperationLabelBulkUpdateResponseMethodPost    OperationLabelBulkUpdateResponseMethod = "POST"
	OperationLabelBulkUpdateResponseMethodHead    OperationLabelBulkUpdateResponseMethod = "HEAD"
	OperationLabelBulkUpdateResponseMethodOptions OperationLabelBulkUpdateResponseMethod = "OPTIONS"
	OperationLabelBulkUpdateResponseMethodPut     OperationLabelBulkUpdateResponseMethod = "PUT"
	OperationLabelBulkUpdateResponseMethodDelete  OperationLabelBulkUpdateResponseMethod = "DELETE"
	OperationLabelBulkUpdateResponseMethodConnect OperationLabelBulkUpdateResponseMethod = "CONNECT"
	OperationLabelBulkUpdateResponseMethodPatch   OperationLabelBulkUpdateResponseMethod = "PATCH"
	OperationLabelBulkUpdateResponseMethodTrace   OperationLabelBulkUpdateResponseMethod = "TRACE"
)

func (r OperationLabelBulkUpdateResponseMethod) IsKnown() bool {
	switch r {
	case OperationLabelBulkUpdateResponseMethodGet, OperationLabelBulkUpdateResponseMethodPost, OperationLabelBulkUpdateResponseMethodHead, OperationLabelBulkUpdateResponseMethodOptions, OperationLabelBulkUpdateResponseMethodPut, OperationLabelBulkUpdateResponseMethodDelete, OperationLabelBulkUpdateResponseMethodConnect, OperationLabelBulkUpdateResponseMethodPatch, OperationLabelBulkUpdateResponseMethodTrace:
		return true
	}
	return false
}

type OperationLabelBulkUpdateResponseLabel struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source OperationLabelBulkUpdateResponseLabelsSource `json:"source" api:"required"`
	JSON   operationLabelBulkUpdateResponseLabelJSON    `json:"-"`
}

// operationLabelBulkUpdateResponseLabelJSON contains the JSON metadata for the
// struct [OperationLabelBulkUpdateResponseLabel]
type operationLabelBulkUpdateResponseLabelJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelBulkUpdateResponseLabel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelBulkUpdateResponseLabelJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type OperationLabelBulkUpdateResponseLabelsSource string

const (
	OperationLabelBulkUpdateResponseLabelsSourceUser    OperationLabelBulkUpdateResponseLabelsSource = "user"
	OperationLabelBulkUpdateResponseLabelsSourceManaged OperationLabelBulkUpdateResponseLabelsSource = "managed"
)

func (r OperationLabelBulkUpdateResponseLabelsSource) IsKnown() bool {
	switch r {
	case OperationLabelBulkUpdateResponseLabelsSourceUser, OperationLabelBulkUpdateResponseLabelsSourceManaged:
		return true
	}
	return false
}

type OperationLabelNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// List of managed label names.
	Managed param.Field[[]string] `json:"managed"`
	// List of user label names.
	User param.Field[[]string] `json:"user"`
}

func (r OperationLabelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelNewResponseEnvelope struct {
	Errors   Message                   `json:"errors" api:"required"`
	Messages Message                   `json:"messages" api:"required"`
	Result   OperationLabelNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success OperationLabelNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    operationLabelNewResponseEnvelopeJSON    `json:"-"`
}

// operationLabelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationLabelNewResponseEnvelope]
type operationLabelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationLabelNewResponseEnvelopeSuccess bool

const (
	OperationLabelNewResponseEnvelopeSuccessTrue OperationLabelNewResponseEnvelopeSuccess = true
)

func (r OperationLabelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationLabelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OperationLabelUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// List of managed label names. Omitting this property or passing an empty array
	// will result in all managed labels being removed from the operation
	Managed param.Field[[]string] `json:"managed"`
	// List of user label names. Omitting this property or passing an empty array will
	// result in all user labels being removed from the operation
	User param.Field[[]string] `json:"user"`
}

func (r OperationLabelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelUpdateResponseEnvelope struct {
	Errors   Message                      `json:"errors" api:"required"`
	Messages Message                      `json:"messages" api:"required"`
	Result   OperationLabelUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success OperationLabelUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    operationLabelUpdateResponseEnvelopeJSON    `json:"-"`
}

// operationLabelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [OperationLabelUpdateResponseEnvelope]
type operationLabelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationLabelUpdateResponseEnvelopeSuccess bool

const (
	OperationLabelUpdateResponseEnvelopeSuccessTrue OperationLabelUpdateResponseEnvelopeSuccess = true
)

func (r OperationLabelUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationLabelUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OperationLabelDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OperationLabelDeleteResponseEnvelope struct {
	Errors   Message                      `json:"errors" api:"required"`
	Messages Message                      `json:"messages" api:"required"`
	Result   OperationLabelDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success OperationLabelDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    operationLabelDeleteResponseEnvelopeJSON    `json:"-"`
}

// operationLabelDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OperationLabelDeleteResponseEnvelope]
type operationLabelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationLabelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationLabelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationLabelDeleteResponseEnvelopeSuccess bool

const (
	OperationLabelDeleteResponseEnvelopeSuccessTrue OperationLabelDeleteResponseEnvelopeSuccess = true
)

func (r OperationLabelDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationLabelDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OperationLabelBulkNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Operation IDs selector
	Selector param.Field[OperationLabelBulkNewParamsSelector] `json:"selector" api:"required"`
	Managed  param.Field[OperationLabelBulkNewParamsManaged]  `json:"managed"`
	User     param.Field[OperationLabelBulkNewParamsUser]     `json:"user"`
}

func (r OperationLabelBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Operation IDs selector
type OperationLabelBulkNewParamsSelector struct {
	Include param.Field[OperationLabelBulkNewParamsSelectorInclude] `json:"include" api:"required"`
}

func (r OperationLabelBulkNewParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelBulkNewParamsSelectorInclude struct {
	OperationIDs param.Field[[]string] `json:"operation_ids" api:"required"`
}

func (r OperationLabelBulkNewParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelBulkNewParamsManaged struct {
	// List of managed label names.
	Labels param.Field[[]string] `json:"labels"`
}

func (r OperationLabelBulkNewParamsManaged) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelBulkNewParamsUser struct {
	// List of user label names.
	Labels param.Field[[]string] `json:"labels"`
}

func (r OperationLabelBulkNewParamsUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelBulkDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OperationLabelBulkUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Managed labels to replace for all affected operations
	Managed param.Field[OperationLabelBulkUpdateParamsManaged] `json:"managed" api:"required"`
	// Operation IDs selector
	Selector param.Field[OperationLabelBulkUpdateParamsSelector] `json:"selector" api:"required"`
	// User labels to replace for all affected operations
	User param.Field[OperationLabelBulkUpdateParamsUser] `json:"user" api:"required"`
}

func (r OperationLabelBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Managed labels to replace for all affected operations
type OperationLabelBulkUpdateParamsManaged struct {
	// List of managed label names. Providing an empty array will result in all managed
	// labels being removed from all affected operations
	Labels param.Field[[]string] `json:"labels" api:"required"`
}

func (r OperationLabelBulkUpdateParamsManaged) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Operation IDs selector
type OperationLabelBulkUpdateParamsSelector struct {
	Include param.Field[OperationLabelBulkUpdateParamsSelectorInclude] `json:"include" api:"required"`
}

func (r OperationLabelBulkUpdateParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationLabelBulkUpdateParamsSelectorInclude struct {
	OperationIDs param.Field[[]string] `json:"operation_ids" api:"required"`
}

func (r OperationLabelBulkUpdateParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User labels to replace for all affected operations
type OperationLabelBulkUpdateParamsUser struct {
	// List of user label names. Providing an empty array will result in all user
	// labels being removed from all affected operations
	Labels param.Field[[]string] `json:"labels" api:"required"`
}

func (r OperationLabelBulkUpdateParamsUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
