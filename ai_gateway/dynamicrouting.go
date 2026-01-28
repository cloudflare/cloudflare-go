// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/tidwall/gjson"
)

// DynamicRoutingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDynamicRoutingService] method instead.
type DynamicRoutingService struct {
	Options []option.RequestOption
}

// NewDynamicRoutingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDynamicRoutingService(opts ...option.RequestOption) (r *DynamicRoutingService) {
	r = &DynamicRoutingService{}
	r.Options = opts
	return
}

// Create a new AI Gateway Dynamic Route.
func (r *DynamicRoutingService) New(ctx context.Context, gatewayID string, params DynamicRoutingNewParams, opts ...option.RequestOption) (res *DynamicRoutingNewResponse, err error) {
	var env DynamicRoutingNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes", params.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an AI Gateway Dynamic Route.
func (r *DynamicRoutingService) Update(ctx context.Context, gatewayID string, id string, params DynamicRoutingUpdateParams, opts ...option.RequestOption) (res *DynamicRoutingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s", params.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all AI Gateway Dynamic Routes.
func (r *DynamicRoutingService) List(ctx context.Context, gatewayID string, query DynamicRoutingListParams, opts ...option.RequestOption) (res *DynamicRoutingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes", query.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an AI Gateway Dynamic Route.
func (r *DynamicRoutingService) Delete(ctx context.Context, gatewayID string, id string, body DynamicRoutingDeleteParams, opts ...option.RequestOption) (res *DynamicRoutingDeleteResponse, err error) {
	var env DynamicRoutingDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s", body.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new AI Gateway Dynamic Route Deployment.
func (r *DynamicRoutingService) NewDeployment(ctx context.Context, gatewayID string, id string, params DynamicRoutingNewDeploymentParams, opts ...option.RequestOption) (res *DynamicRoutingNewDeploymentResponse, err error) {
	var env DynamicRoutingNewDeploymentResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s/deployments", params.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new AI Gateway Dynamic Route Version.
func (r *DynamicRoutingService) NewVersion(ctx context.Context, gatewayID string, id string, params DynamicRoutingNewVersionParams, opts ...option.RequestOption) (res *DynamicRoutingNewVersionResponse, err error) {
	var env DynamicRoutingNewVersionResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s/versions", params.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an AI Gateway Dynamic Route.
func (r *DynamicRoutingService) Get(ctx context.Context, gatewayID string, id string, query DynamicRoutingGetParams, opts ...option.RequestOption) (res *DynamicRoutingGetResponse, err error) {
	var env DynamicRoutingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an AI Gateway Dynamic Route Version.
func (r *DynamicRoutingService) GetVersion(ctx context.Context, gatewayID string, id string, versionID string, query DynamicRoutingGetVersionParams, opts ...option.RequestOption) (res *DynamicRoutingGetVersionResponse, err error) {
	var env DynamicRoutingGetVersionResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s/versions/%s", query.AccountID, gatewayID, id, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all AI Gateway Dynamic Route Deployments.
func (r *DynamicRoutingService) ListDeployments(ctx context.Context, gatewayID string, id string, query DynamicRoutingListDeploymentsParams, opts ...option.RequestOption) (res *DynamicRoutingListDeploymentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s/deployments", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all AI Gateway Dynamic Route Versions.
func (r *DynamicRoutingService) ListVersions(ctx context.Context, gatewayID string, id string, query DynamicRoutingListVersionsParams, opts ...option.RequestOption) (res *DynamicRoutingListVersionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/routes/%s/versions", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DynamicRoutingNewResponse struct {
	ID         string                              `json:"id,required"`
	AccountTag string                              `json:"account_tag,required"`
	CreatedAt  time.Time                           `json:"created_at,required" format:"date-time"`
	Deployment DynamicRoutingNewResponseDeployment `json:"deployment,required"`
	Elements   []DynamicRoutingNewResponseElement  `json:"elements,required"`
	GatewayID  string                              `json:"gateway_id,required"`
	ModifiedAt time.Time                           `json:"modified_at,required" format:"date-time"`
	Name       string                              `json:"name,required"`
	Version    DynamicRoutingNewResponseVersion    `json:"version,required"`
	JSON       dynamicRoutingNewResponseJSON       `json:"-"`
}

// dynamicRoutingNewResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingNewResponse]
type dynamicRoutingNewResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Deployment  apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewResponseDeployment struct {
	CreatedAt    string                                  `json:"created_at,required"`
	DeploymentID string                                  `json:"deployment_id,required"`
	VersionID    string                                  `json:"version_id,required"`
	Comment      string                                  `json:"comment,nullable"`
	JSON         dynamicRoutingNewResponseDeploymentJSON `json:"-"`
}

// dynamicRoutingNewResponseDeploymentJSON contains the JSON metadata for the
// struct [DynamicRoutingNewResponseDeployment]
type dynamicRoutingNewResponseDeploymentJSON struct {
	CreatedAt    apijson.Field
	DeploymentID apijson.Field
	VersionID    apijson.Field
	Comment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingNewResponseElementsObjectOutput].
	Outputs interface{}                           `json:"outputs,required"`
	Type    DynamicRoutingNewResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewResponseElementsObjectProperties].
	Properties interface{}                          `json:"properties"`
	JSON       dynamicRoutingNewResponseElementJSON `json:"-"`
	union      DynamicRoutingNewResponseElementsUnion
}

// dynamicRoutingNewResponseElementJSON contains the JSON metadata for the struct
// [DynamicRoutingNewResponseElement]
type dynamicRoutingNewResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingNewResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingNewResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingNewResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingNewResponseElementsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject].
func (r DynamicRoutingNewResponseElement) AsUnion() DynamicRoutingNewResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject],
// [DynamicRoutingNewResponseElementsObject] or
// [DynamicRoutingNewResponseElementsObject].
type DynamicRoutingNewResponseElementsUnion interface {
	implementsDynamicRoutingNewResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingNewResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewResponseElementsObject{}),
		},
	)
}

type DynamicRoutingNewResponseElementsObject struct {
	ID      string                                         `json:"id,required"`
	Outputs DynamicRoutingNewResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingNewResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingNewResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingNewResponseElementsObjectJSON contains the JSON metadata for the
// struct [DynamicRoutingNewResponseElementsObject]
type dynamicRoutingNewResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingNewResponseElementsObject) implementsDynamicRoutingNewResponseElement() {}

type DynamicRoutingNewResponseElementsObjectOutputs struct {
	Next DynamicRoutingNewResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingNewResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingNewResponseElementsObjectOutputsJSON contains the JSON metadata
// for the struct [DynamicRoutingNewResponseElementsObjectOutputs]
type dynamicRoutingNewResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewResponseElementsObjectOutputsNext struct {
	ElementID string                                                 `json:"elementId,required"`
	JSON      dynamicRoutingNewResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingNewResponseElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct [DynamicRoutingNewResponseElementsObjectOutputsNext]
type dynamicRoutingNewResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewResponseElementsObjectType string

const (
	DynamicRoutingNewResponseElementsObjectTypeStart DynamicRoutingNewResponseElementsObjectType = "start"
)

func (r DynamicRoutingNewResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingNewResponseElementsType string

const (
	DynamicRoutingNewResponseElementsTypeStart       DynamicRoutingNewResponseElementsType = "start"
	DynamicRoutingNewResponseElementsTypeConditional DynamicRoutingNewResponseElementsType = "conditional"
	DynamicRoutingNewResponseElementsTypePercentage  DynamicRoutingNewResponseElementsType = "percentage"
	DynamicRoutingNewResponseElementsTypeRate        DynamicRoutingNewResponseElementsType = "rate"
	DynamicRoutingNewResponseElementsTypeModel       DynamicRoutingNewResponseElementsType = "model"
	DynamicRoutingNewResponseElementsTypeEnd         DynamicRoutingNewResponseElementsType = "end"
)

func (r DynamicRoutingNewResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewResponseElementsTypeStart, DynamicRoutingNewResponseElementsTypeConditional, DynamicRoutingNewResponseElementsTypePercentage, DynamicRoutingNewResponseElementsTypeRate, DynamicRoutingNewResponseElementsTypeModel, DynamicRoutingNewResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingNewResponseVersion struct {
	Active    DynamicRoutingNewResponseVersionActive `json:"active,required"`
	CreatedAt string                                 `json:"created_at,required"`
	Data      string                                 `json:"data,required"`
	VersionID string                                 `json:"version_id,required"`
	Comment   string                                 `json:"comment,nullable"`
	JSON      dynamicRoutingNewResponseVersionJSON   `json:"-"`
}

// dynamicRoutingNewResponseVersionJSON contains the JSON metadata for the struct
// [DynamicRoutingNewResponseVersion]
type dynamicRoutingNewResponseVersionJSON struct {
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseVersionJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewResponseVersionActive string

const (
	DynamicRoutingNewResponseVersionActiveTrue  DynamicRoutingNewResponseVersionActive = "true"
	DynamicRoutingNewResponseVersionActiveFalse DynamicRoutingNewResponseVersionActive = "false"
)

func (r DynamicRoutingNewResponseVersionActive) IsKnown() bool {
	switch r {
	case DynamicRoutingNewResponseVersionActiveTrue, DynamicRoutingNewResponseVersionActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingUpdateResponse struct {
	Route   DynamicRoutingUpdateResponseRoute `json:"route,required"`
	Success bool                              `json:"success,required"`
	JSON    dynamicRoutingUpdateResponseJSON  `json:"-"`
}

// dynamicRoutingUpdateResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingUpdateResponse]
type dynamicRoutingUpdateResponseJSON struct {
	Route       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRoute struct {
	ID         string                                      `json:"id,required"`
	AccountTag string                                      `json:"account_tag,required"`
	CreatedAt  time.Time                                   `json:"created_at,required" format:"date-time"`
	Deployment DynamicRoutingUpdateResponseRouteDeployment `json:"deployment,required"`
	Elements   []DynamicRoutingUpdateResponseRouteElement  `json:"elements,required"`
	GatewayID  string                                      `json:"gateway_id,required"`
	ModifiedAt time.Time                                   `json:"modified_at,required" format:"date-time"`
	Name       string                                      `json:"name,required"`
	Version    DynamicRoutingUpdateResponseRouteVersion    `json:"version,required"`
	JSON       dynamicRoutingUpdateResponseRouteJSON       `json:"-"`
}

// dynamicRoutingUpdateResponseRouteJSON contains the JSON metadata for the struct
// [DynamicRoutingUpdateResponseRoute]
type dynamicRoutingUpdateResponseRouteJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Deployment  apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRouteDeployment struct {
	CreatedAt    string                                          `json:"created_at,required"`
	DeploymentID string                                          `json:"deployment_id,required"`
	VersionID    string                                          `json:"version_id,required"`
	Comment      string                                          `json:"comment,nullable"`
	JSON         dynamicRoutingUpdateResponseRouteDeploymentJSON `json:"-"`
}

// dynamicRoutingUpdateResponseRouteDeploymentJSON contains the JSON metadata for
// the struct [DynamicRoutingUpdateResponseRouteDeployment]
type dynamicRoutingUpdateResponseRouteDeploymentJSON struct {
	CreatedAt    apijson.Field
	DeploymentID apijson.Field
	VersionID    apijson.Field
	Comment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRouteDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteDeploymentJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRouteElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingUpdateResponseRouteElementsObjectOutputs],
	// [map[string]DynamicRoutingUpdateResponseRouteElementsObjectOutput].
	Outputs interface{}                                   `json:"outputs,required"`
	Type    DynamicRoutingUpdateResponseRouteElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingUpdateResponseRouteElementsObjectProperties].
	Properties interface{}                                  `json:"properties"`
	JSON       dynamicRoutingUpdateResponseRouteElementJSON `json:"-"`
	union      DynamicRoutingUpdateResponseRouteElementsUnion
}

// dynamicRoutingUpdateResponseRouteElementJSON contains the JSON metadata for the
// struct [DynamicRoutingUpdateResponseRouteElement]
type dynamicRoutingUpdateResponseRouteElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingUpdateResponseRouteElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingUpdateResponseRouteElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingUpdateResponseRouteElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingUpdateResponseRouteElementsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject].
func (r DynamicRoutingUpdateResponseRouteElement) AsUnion() DynamicRoutingUpdateResponseRouteElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject],
// [DynamicRoutingUpdateResponseRouteElementsObject] or
// [DynamicRoutingUpdateResponseRouteElementsObject].
type DynamicRoutingUpdateResponseRouteElementsUnion interface {
	implementsDynamicRoutingUpdateResponseRouteElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingUpdateResponseRouteElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingUpdateResponseRouteElementsObject{}),
		},
	)
}

type DynamicRoutingUpdateResponseRouteElementsObject struct {
	ID      string                                                 `json:"id,required"`
	Outputs DynamicRoutingUpdateResponseRouteElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingUpdateResponseRouteElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingUpdateResponseRouteElementsObjectJSON    `json:"-"`
}

// dynamicRoutingUpdateResponseRouteElementsObjectJSON contains the JSON metadata
// for the struct [DynamicRoutingUpdateResponseRouteElementsObject]
type dynamicRoutingUpdateResponseRouteElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRouteElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingUpdateResponseRouteElementsObject) implementsDynamicRoutingUpdateResponseRouteElement() {
}

type DynamicRoutingUpdateResponseRouteElementsObjectOutputs struct {
	Next DynamicRoutingUpdateResponseRouteElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingUpdateResponseRouteElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingUpdateResponseRouteElementsObjectOutputsJSON contains the JSON
// metadata for the struct [DynamicRoutingUpdateResponseRouteElementsObjectOutputs]
type dynamicRoutingUpdateResponseRouteElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRouteElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRouteElementsObjectOutputsNext struct {
	ElementID string                                                         `json:"elementId,required"`
	JSON      dynamicRoutingUpdateResponseRouteElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingUpdateResponseRouteElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct
// [DynamicRoutingUpdateResponseRouteElementsObjectOutputsNext]
type dynamicRoutingUpdateResponseRouteElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRouteElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRouteElementsObjectType string

const (
	DynamicRoutingUpdateResponseRouteElementsObjectTypeStart DynamicRoutingUpdateResponseRouteElementsObjectType = "start"
)

func (r DynamicRoutingUpdateResponseRouteElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingUpdateResponseRouteElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingUpdateResponseRouteElementsType string

const (
	DynamicRoutingUpdateResponseRouteElementsTypeStart       DynamicRoutingUpdateResponseRouteElementsType = "start"
	DynamicRoutingUpdateResponseRouteElementsTypeConditional DynamicRoutingUpdateResponseRouteElementsType = "conditional"
	DynamicRoutingUpdateResponseRouteElementsTypePercentage  DynamicRoutingUpdateResponseRouteElementsType = "percentage"
	DynamicRoutingUpdateResponseRouteElementsTypeRate        DynamicRoutingUpdateResponseRouteElementsType = "rate"
	DynamicRoutingUpdateResponseRouteElementsTypeModel       DynamicRoutingUpdateResponseRouteElementsType = "model"
	DynamicRoutingUpdateResponseRouteElementsTypeEnd         DynamicRoutingUpdateResponseRouteElementsType = "end"
)

func (r DynamicRoutingUpdateResponseRouteElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingUpdateResponseRouteElementsTypeStart, DynamicRoutingUpdateResponseRouteElementsTypeConditional, DynamicRoutingUpdateResponseRouteElementsTypePercentage, DynamicRoutingUpdateResponseRouteElementsTypeRate, DynamicRoutingUpdateResponseRouteElementsTypeModel, DynamicRoutingUpdateResponseRouteElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingUpdateResponseRouteVersion struct {
	Active    DynamicRoutingUpdateResponseRouteVersionActive `json:"active,required"`
	CreatedAt string                                         `json:"created_at,required"`
	Data      string                                         `json:"data,required"`
	VersionID string                                         `json:"version_id,required"`
	Comment   string                                         `json:"comment,nullable"`
	JSON      dynamicRoutingUpdateResponseRouteVersionJSON   `json:"-"`
}

// dynamicRoutingUpdateResponseRouteVersionJSON contains the JSON metadata for the
// struct [DynamicRoutingUpdateResponseRouteVersion]
type dynamicRoutingUpdateResponseRouteVersionJSON struct {
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingUpdateResponseRouteVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingUpdateResponseRouteVersionJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateResponseRouteVersionActive string

const (
	DynamicRoutingUpdateResponseRouteVersionActiveTrue  DynamicRoutingUpdateResponseRouteVersionActive = "true"
	DynamicRoutingUpdateResponseRouteVersionActiveFalse DynamicRoutingUpdateResponseRouteVersionActive = "false"
)

func (r DynamicRoutingUpdateResponseRouteVersionActive) IsKnown() bool {
	switch r {
	case DynamicRoutingUpdateResponseRouteVersionActiveTrue, DynamicRoutingUpdateResponseRouteVersionActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingListResponse struct {
	Data    DynamicRoutingListResponseData `json:"data,required"`
	Success bool                           `json:"success,required"`
	JSON    dynamicRoutingListResponseJSON `json:"-"`
}

// dynamicRoutingListResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingListResponse]
type dynamicRoutingListResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseData struct {
	OrderBy          string                                `json:"order_by,required"`
	OrderByDirection string                                `json:"order_by_direction,required"`
	Page             float64                               `json:"page,required"`
	PerPage          float64                               `json:"per_page,required"`
	Routes           []DynamicRoutingListResponseDataRoute `json:"routes,required"`
	JSON             dynamicRoutingListResponseDataJSON    `json:"-"`
}

// dynamicRoutingListResponseDataJSON contains the JSON metadata for the struct
// [DynamicRoutingListResponseData]
type dynamicRoutingListResponseDataJSON struct {
	OrderBy          apijson.Field
	OrderByDirection apijson.Field
	Page             apijson.Field
	PerPage          apijson.Field
	Routes           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DynamicRoutingListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoute struct {
	ID         string                                         `json:"id,required"`
	AccountTag string                                         `json:"account_tag,required"`
	CreatedAt  time.Time                                      `json:"created_at,required" format:"date-time"`
	Deployment DynamicRoutingListResponseDataRoutesDeployment `json:"deployment,required"`
	Elements   []DynamicRoutingListResponseDataRoutesElement  `json:"elements,required"`
	GatewayID  string                                         `json:"gateway_id,required"`
	ModifiedAt time.Time                                      `json:"modified_at,required" format:"date-time"`
	Name       string                                         `json:"name,required"`
	Version    DynamicRoutingListResponseDataRoutesVersion    `json:"version,required"`
	JSON       dynamicRoutingListResponseDataRouteJSON        `json:"-"`
}

// dynamicRoutingListResponseDataRouteJSON contains the JSON metadata for the
// struct [DynamicRoutingListResponseDataRoute]
type dynamicRoutingListResponseDataRouteJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Deployment  apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRouteJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoutesDeployment struct {
	CreatedAt    string                                             `json:"created_at,required"`
	DeploymentID string                                             `json:"deployment_id,required"`
	VersionID    string                                             `json:"version_id,required"`
	Comment      string                                             `json:"comment,nullable"`
	JSON         dynamicRoutingListResponseDataRoutesDeploymentJSON `json:"-"`
}

// dynamicRoutingListResponseDataRoutesDeploymentJSON contains the JSON metadata
// for the struct [DynamicRoutingListResponseDataRoutesDeployment]
type dynamicRoutingListResponseDataRoutesDeploymentJSON struct {
	CreatedAt    apijson.Field
	DeploymentID apijson.Field
	VersionID    apijson.Field
	Comment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoutesDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRoutesDeploymentJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoutesElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingListResponseDataRoutesElementsObjectOutputs],
	// [map[string]DynamicRoutingListResponseDataRoutesElementsObjectOutput].
	Outputs interface{}                                      `json:"outputs,required"`
	Type    DynamicRoutingListResponseDataRoutesElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingListResponseDataRoutesElementsObjectProperties].
	Properties interface{}                                     `json:"properties"`
	JSON       dynamicRoutingListResponseDataRoutesElementJSON `json:"-"`
	union      DynamicRoutingListResponseDataRoutesElementsUnion
}

// dynamicRoutingListResponseDataRoutesElementJSON contains the JSON metadata for
// the struct [DynamicRoutingListResponseDataRoutesElement]
type dynamicRoutingListResponseDataRoutesElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingListResponseDataRoutesElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingListResponseDataRoutesElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingListResponseDataRoutesElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingListResponseDataRoutesElementsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject].
func (r DynamicRoutingListResponseDataRoutesElement) AsUnion() DynamicRoutingListResponseDataRoutesElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject],
// [DynamicRoutingListResponseDataRoutesElementsObject] or
// [DynamicRoutingListResponseDataRoutesElementsObject].
type DynamicRoutingListResponseDataRoutesElementsUnion interface {
	implementsDynamicRoutingListResponseDataRoutesElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingListResponseDataRoutesElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingListResponseDataRoutesElementsObject{}),
		},
	)
}

type DynamicRoutingListResponseDataRoutesElementsObject struct {
	ID      string                                                    `json:"id,required"`
	Outputs DynamicRoutingListResponseDataRoutesElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingListResponseDataRoutesElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingListResponseDataRoutesElementsObjectJSON    `json:"-"`
}

// dynamicRoutingListResponseDataRoutesElementsObjectJSON contains the JSON
// metadata for the struct [DynamicRoutingListResponseDataRoutesElementsObject]
type dynamicRoutingListResponseDataRoutesElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoutesElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRoutesElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingListResponseDataRoutesElementsObject) implementsDynamicRoutingListResponseDataRoutesElement() {
}

type DynamicRoutingListResponseDataRoutesElementsObjectOutputs struct {
	Next DynamicRoutingListResponseDataRoutesElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingListResponseDataRoutesElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingListResponseDataRoutesElementsObjectOutputsJSON contains the JSON
// metadata for the struct
// [DynamicRoutingListResponseDataRoutesElementsObjectOutputs]
type dynamicRoutingListResponseDataRoutesElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoutesElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRoutesElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoutesElementsObjectOutputsNext struct {
	ElementID string                                                            `json:"elementId,required"`
	JSON      dynamicRoutingListResponseDataRoutesElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingListResponseDataRoutesElementsObjectOutputsNextJSON contains the
// JSON metadata for the struct
// [DynamicRoutingListResponseDataRoutesElementsObjectOutputsNext]
type dynamicRoutingListResponseDataRoutesElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoutesElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRoutesElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoutesElementsObjectType string

const (
	DynamicRoutingListResponseDataRoutesElementsObjectTypeStart DynamicRoutingListResponseDataRoutesElementsObjectType = "start"
)

func (r DynamicRoutingListResponseDataRoutesElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingListResponseDataRoutesElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingListResponseDataRoutesElementsType string

const (
	DynamicRoutingListResponseDataRoutesElementsTypeStart       DynamicRoutingListResponseDataRoutesElementsType = "start"
	DynamicRoutingListResponseDataRoutesElementsTypeConditional DynamicRoutingListResponseDataRoutesElementsType = "conditional"
	DynamicRoutingListResponseDataRoutesElementsTypePercentage  DynamicRoutingListResponseDataRoutesElementsType = "percentage"
	DynamicRoutingListResponseDataRoutesElementsTypeRate        DynamicRoutingListResponseDataRoutesElementsType = "rate"
	DynamicRoutingListResponseDataRoutesElementsTypeModel       DynamicRoutingListResponseDataRoutesElementsType = "model"
	DynamicRoutingListResponseDataRoutesElementsTypeEnd         DynamicRoutingListResponseDataRoutesElementsType = "end"
)

func (r DynamicRoutingListResponseDataRoutesElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingListResponseDataRoutesElementsTypeStart, DynamicRoutingListResponseDataRoutesElementsTypeConditional, DynamicRoutingListResponseDataRoutesElementsTypePercentage, DynamicRoutingListResponseDataRoutesElementsTypeRate, DynamicRoutingListResponseDataRoutesElementsTypeModel, DynamicRoutingListResponseDataRoutesElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingListResponseDataRoutesVersion struct {
	Active    DynamicRoutingListResponseDataRoutesVersionActive `json:"active,required"`
	CreatedAt string                                            `json:"created_at,required"`
	Data      string                                            `json:"data,required"`
	VersionID string                                            `json:"version_id,required"`
	Comment   string                                            `json:"comment,nullable"`
	JSON      dynamicRoutingListResponseDataRoutesVersionJSON   `json:"-"`
}

// dynamicRoutingListResponseDataRoutesVersionJSON contains the JSON metadata for
// the struct [DynamicRoutingListResponseDataRoutesVersion]
type dynamicRoutingListResponseDataRoutesVersionJSON struct {
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListResponseDataRoutesVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListResponseDataRoutesVersionJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListResponseDataRoutesVersionActive string

const (
	DynamicRoutingListResponseDataRoutesVersionActiveTrue  DynamicRoutingListResponseDataRoutesVersionActive = "true"
	DynamicRoutingListResponseDataRoutesVersionActiveFalse DynamicRoutingListResponseDataRoutesVersionActive = "false"
)

func (r DynamicRoutingListResponseDataRoutesVersionActive) IsKnown() bool {
	switch r {
	case DynamicRoutingListResponseDataRoutesVersionActiveTrue, DynamicRoutingListResponseDataRoutesVersionActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingDeleteResponse struct {
	ID         string                                `json:"id,required"`
	AccountTag string                                `json:"account_tag,required"`
	CreatedAt  time.Time                             `json:"created_at,required" format:"date-time"`
	Elements   []DynamicRoutingDeleteResponseElement `json:"elements,required"`
	GatewayID  string                                `json:"gateway_id,required"`
	ModifiedAt time.Time                             `json:"modified_at,required" format:"date-time"`
	Name       string                                `json:"name,required"`
	JSON       dynamicRoutingDeleteResponseJSON      `json:"-"`
}

// dynamicRoutingDeleteResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingDeleteResponse]
type dynamicRoutingDeleteResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingDeleteResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingDeleteResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingDeleteResponseElementsObjectOutput].
	Outputs interface{}                              `json:"outputs,required"`
	Type    DynamicRoutingDeleteResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingDeleteResponseElementsObjectProperties].
	Properties interface{}                             `json:"properties"`
	JSON       dynamicRoutingDeleteResponseElementJSON `json:"-"`
	union      DynamicRoutingDeleteResponseElementsUnion
}

// dynamicRoutingDeleteResponseElementJSON contains the JSON metadata for the
// struct [DynamicRoutingDeleteResponseElement]
type dynamicRoutingDeleteResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingDeleteResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingDeleteResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingDeleteResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingDeleteResponseElementsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject].
func (r DynamicRoutingDeleteResponseElement) AsUnion() DynamicRoutingDeleteResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject],
// [DynamicRoutingDeleteResponseElementsObject] or
// [DynamicRoutingDeleteResponseElementsObject].
type DynamicRoutingDeleteResponseElementsUnion interface {
	implementsDynamicRoutingDeleteResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingDeleteResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingDeleteResponseElementsObject{}),
		},
	)
}

type DynamicRoutingDeleteResponseElementsObject struct {
	ID      string                                            `json:"id,required"`
	Outputs DynamicRoutingDeleteResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingDeleteResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingDeleteResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingDeleteResponseElementsObjectJSON contains the JSON metadata for
// the struct [DynamicRoutingDeleteResponseElementsObject]
type dynamicRoutingDeleteResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingDeleteResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingDeleteResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingDeleteResponseElementsObject) implementsDynamicRoutingDeleteResponseElement() {}

type DynamicRoutingDeleteResponseElementsObjectOutputs struct {
	Next DynamicRoutingDeleteResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingDeleteResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingDeleteResponseElementsObjectOutputsJSON contains the JSON metadata
// for the struct [DynamicRoutingDeleteResponseElementsObjectOutputs]
type dynamicRoutingDeleteResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingDeleteResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingDeleteResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingDeleteResponseElementsObjectOutputsNext struct {
	ElementID string                                                    `json:"elementId,required"`
	JSON      dynamicRoutingDeleteResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingDeleteResponseElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct [DynamicRoutingDeleteResponseElementsObjectOutputsNext]
type dynamicRoutingDeleteResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingDeleteResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingDeleteResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingDeleteResponseElementsObjectType string

const (
	DynamicRoutingDeleteResponseElementsObjectTypeStart DynamicRoutingDeleteResponseElementsObjectType = "start"
)

func (r DynamicRoutingDeleteResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingDeleteResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingDeleteResponseElementsType string

const (
	DynamicRoutingDeleteResponseElementsTypeStart       DynamicRoutingDeleteResponseElementsType = "start"
	DynamicRoutingDeleteResponseElementsTypeConditional DynamicRoutingDeleteResponseElementsType = "conditional"
	DynamicRoutingDeleteResponseElementsTypePercentage  DynamicRoutingDeleteResponseElementsType = "percentage"
	DynamicRoutingDeleteResponseElementsTypeRate        DynamicRoutingDeleteResponseElementsType = "rate"
	DynamicRoutingDeleteResponseElementsTypeModel       DynamicRoutingDeleteResponseElementsType = "model"
	DynamicRoutingDeleteResponseElementsTypeEnd         DynamicRoutingDeleteResponseElementsType = "end"
)

func (r DynamicRoutingDeleteResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingDeleteResponseElementsTypeStart, DynamicRoutingDeleteResponseElementsTypeConditional, DynamicRoutingDeleteResponseElementsTypePercentage, DynamicRoutingDeleteResponseElementsTypeRate, DynamicRoutingDeleteResponseElementsTypeModel, DynamicRoutingDeleteResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingNewDeploymentResponse struct {
	ID         string                                       `json:"id,required"`
	AccountTag string                                       `json:"account_tag,required"`
	CreatedAt  time.Time                                    `json:"created_at,required" format:"date-time"`
	Elements   []DynamicRoutingNewDeploymentResponseElement `json:"elements,required"`
	GatewayID  string                                       `json:"gateway_id,required"`
	ModifiedAt time.Time                                    `json:"modified_at,required" format:"date-time"`
	Name       string                                       `json:"name,required"`
	JSON       dynamicRoutingNewDeploymentResponseJSON      `json:"-"`
}

// dynamicRoutingNewDeploymentResponseJSON contains the JSON metadata for the
// struct [DynamicRoutingNewDeploymentResponse]
type dynamicRoutingNewDeploymentResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewDeploymentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewDeploymentResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewDeploymentResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewDeploymentResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingNewDeploymentResponseElementsObjectOutput].
	Outputs interface{}                                     `json:"outputs,required"`
	Type    DynamicRoutingNewDeploymentResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewDeploymentResponseElementsObjectProperties].
	Properties interface{}                                    `json:"properties"`
	JSON       dynamicRoutingNewDeploymentResponseElementJSON `json:"-"`
	union      DynamicRoutingNewDeploymentResponseElementsUnion
}

// dynamicRoutingNewDeploymentResponseElementJSON contains the JSON metadata for
// the struct [DynamicRoutingNewDeploymentResponseElement]
type dynamicRoutingNewDeploymentResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingNewDeploymentResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingNewDeploymentResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingNewDeploymentResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingNewDeploymentResponseElementsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject].
func (r DynamicRoutingNewDeploymentResponseElement) AsUnion() DynamicRoutingNewDeploymentResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject],
// [DynamicRoutingNewDeploymentResponseElementsObject] or
// [DynamicRoutingNewDeploymentResponseElementsObject].
type DynamicRoutingNewDeploymentResponseElementsUnion interface {
	implementsDynamicRoutingNewDeploymentResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingNewDeploymentResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewDeploymentResponseElementsObject{}),
		},
	)
}

type DynamicRoutingNewDeploymentResponseElementsObject struct {
	ID      string                                                   `json:"id,required"`
	Outputs DynamicRoutingNewDeploymentResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingNewDeploymentResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingNewDeploymentResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingNewDeploymentResponseElementsObjectJSON contains the JSON metadata
// for the struct [DynamicRoutingNewDeploymentResponseElementsObject]
type dynamicRoutingNewDeploymentResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewDeploymentResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewDeploymentResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingNewDeploymentResponseElementsObject) implementsDynamicRoutingNewDeploymentResponseElement() {
}

type DynamicRoutingNewDeploymentResponseElementsObjectOutputs struct {
	Next DynamicRoutingNewDeploymentResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingNewDeploymentResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingNewDeploymentResponseElementsObjectOutputsJSON contains the JSON
// metadata for the struct
// [DynamicRoutingNewDeploymentResponseElementsObjectOutputs]
type dynamicRoutingNewDeploymentResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewDeploymentResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewDeploymentResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewDeploymentResponseElementsObjectOutputsNext struct {
	ElementID string                                                           `json:"elementId,required"`
	JSON      dynamicRoutingNewDeploymentResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingNewDeploymentResponseElementsObjectOutputsNextJSON contains the
// JSON metadata for the struct
// [DynamicRoutingNewDeploymentResponseElementsObjectOutputsNext]
type dynamicRoutingNewDeploymentResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewDeploymentResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewDeploymentResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewDeploymentResponseElementsObjectType string

const (
	DynamicRoutingNewDeploymentResponseElementsObjectTypeStart DynamicRoutingNewDeploymentResponseElementsObjectType = "start"
)

func (r DynamicRoutingNewDeploymentResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewDeploymentResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingNewDeploymentResponseElementsType string

const (
	DynamicRoutingNewDeploymentResponseElementsTypeStart       DynamicRoutingNewDeploymentResponseElementsType = "start"
	DynamicRoutingNewDeploymentResponseElementsTypeConditional DynamicRoutingNewDeploymentResponseElementsType = "conditional"
	DynamicRoutingNewDeploymentResponseElementsTypePercentage  DynamicRoutingNewDeploymentResponseElementsType = "percentage"
	DynamicRoutingNewDeploymentResponseElementsTypeRate        DynamicRoutingNewDeploymentResponseElementsType = "rate"
	DynamicRoutingNewDeploymentResponseElementsTypeModel       DynamicRoutingNewDeploymentResponseElementsType = "model"
	DynamicRoutingNewDeploymentResponseElementsTypeEnd         DynamicRoutingNewDeploymentResponseElementsType = "end"
)

func (r DynamicRoutingNewDeploymentResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewDeploymentResponseElementsTypeStart, DynamicRoutingNewDeploymentResponseElementsTypeConditional, DynamicRoutingNewDeploymentResponseElementsTypePercentage, DynamicRoutingNewDeploymentResponseElementsTypeRate, DynamicRoutingNewDeploymentResponseElementsTypeModel, DynamicRoutingNewDeploymentResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingNewVersionResponse struct {
	ID         string                                    `json:"id,required"`
	AccountTag string                                    `json:"account_tag,required"`
	CreatedAt  time.Time                                 `json:"created_at,required" format:"date-time"`
	Elements   []DynamicRoutingNewVersionResponseElement `json:"elements,required"`
	GatewayID  string                                    `json:"gateway_id,required"`
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	Name       string                                    `json:"name,required"`
	JSON       dynamicRoutingNewVersionResponseJSON      `json:"-"`
}

// dynamicRoutingNewVersionResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingNewVersionResponse]
type dynamicRoutingNewVersionResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewVersionResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewVersionResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewVersionResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingNewVersionResponseElementsObjectOutput].
	Outputs interface{}                                  `json:"outputs,required"`
	Type    DynamicRoutingNewVersionResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingNewVersionResponseElementsObjectProperties].
	Properties interface{}                                 `json:"properties"`
	JSON       dynamicRoutingNewVersionResponseElementJSON `json:"-"`
	union      DynamicRoutingNewVersionResponseElementsUnion
}

// dynamicRoutingNewVersionResponseElementJSON contains the JSON metadata for the
// struct [DynamicRoutingNewVersionResponseElement]
type dynamicRoutingNewVersionResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingNewVersionResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingNewVersionResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingNewVersionResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingNewVersionResponseElementsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject].
func (r DynamicRoutingNewVersionResponseElement) AsUnion() DynamicRoutingNewVersionResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject],
// [DynamicRoutingNewVersionResponseElementsObject] or
// [DynamicRoutingNewVersionResponseElementsObject].
type DynamicRoutingNewVersionResponseElementsUnion interface {
	implementsDynamicRoutingNewVersionResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingNewVersionResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingNewVersionResponseElementsObject{}),
		},
	)
}

type DynamicRoutingNewVersionResponseElementsObject struct {
	ID      string                                                `json:"id,required"`
	Outputs DynamicRoutingNewVersionResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingNewVersionResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingNewVersionResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingNewVersionResponseElementsObjectJSON contains the JSON metadata
// for the struct [DynamicRoutingNewVersionResponseElementsObject]
type dynamicRoutingNewVersionResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewVersionResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewVersionResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingNewVersionResponseElementsObject) implementsDynamicRoutingNewVersionResponseElement() {
}

type DynamicRoutingNewVersionResponseElementsObjectOutputs struct {
	Next DynamicRoutingNewVersionResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingNewVersionResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingNewVersionResponseElementsObjectOutputsJSON contains the JSON
// metadata for the struct [DynamicRoutingNewVersionResponseElementsObjectOutputs]
type dynamicRoutingNewVersionResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewVersionResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewVersionResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewVersionResponseElementsObjectOutputsNext struct {
	ElementID string                                                        `json:"elementId,required"`
	JSON      dynamicRoutingNewVersionResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingNewVersionResponseElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct
// [DynamicRoutingNewVersionResponseElementsObjectOutputsNext]
type dynamicRoutingNewVersionResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewVersionResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewVersionResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewVersionResponseElementsObjectType string

const (
	DynamicRoutingNewVersionResponseElementsObjectTypeStart DynamicRoutingNewVersionResponseElementsObjectType = "start"
)

func (r DynamicRoutingNewVersionResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewVersionResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingNewVersionResponseElementsType string

const (
	DynamicRoutingNewVersionResponseElementsTypeStart       DynamicRoutingNewVersionResponseElementsType = "start"
	DynamicRoutingNewVersionResponseElementsTypeConditional DynamicRoutingNewVersionResponseElementsType = "conditional"
	DynamicRoutingNewVersionResponseElementsTypePercentage  DynamicRoutingNewVersionResponseElementsType = "percentage"
	DynamicRoutingNewVersionResponseElementsTypeRate        DynamicRoutingNewVersionResponseElementsType = "rate"
	DynamicRoutingNewVersionResponseElementsTypeModel       DynamicRoutingNewVersionResponseElementsType = "model"
	DynamicRoutingNewVersionResponseElementsTypeEnd         DynamicRoutingNewVersionResponseElementsType = "end"
)

func (r DynamicRoutingNewVersionResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewVersionResponseElementsTypeStart, DynamicRoutingNewVersionResponseElementsTypeConditional, DynamicRoutingNewVersionResponseElementsTypePercentage, DynamicRoutingNewVersionResponseElementsTypeRate, DynamicRoutingNewVersionResponseElementsTypeModel, DynamicRoutingNewVersionResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingGetResponse struct {
	ID         string                              `json:"id,required"`
	AccountTag string                              `json:"account_tag,required"`
	CreatedAt  time.Time                           `json:"created_at,required" format:"date-time"`
	Deployment DynamicRoutingGetResponseDeployment `json:"deployment,required"`
	Elements   []DynamicRoutingGetResponseElement  `json:"elements,required"`
	GatewayID  string                              `json:"gateway_id,required"`
	ModifiedAt time.Time                           `json:"modified_at,required" format:"date-time"`
	Name       string                              `json:"name,required"`
	Version    DynamicRoutingGetResponseVersion    `json:"version,required"`
	JSON       dynamicRoutingGetResponseJSON       `json:"-"`
}

// dynamicRoutingGetResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingGetResponse]
type dynamicRoutingGetResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Deployment  apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetResponseDeployment struct {
	CreatedAt    string                                  `json:"created_at,required"`
	DeploymentID string                                  `json:"deployment_id,required"`
	VersionID    string                                  `json:"version_id,required"`
	Comment      string                                  `json:"comment,nullable"`
	JSON         dynamicRoutingGetResponseDeploymentJSON `json:"-"`
}

// dynamicRoutingGetResponseDeploymentJSON contains the JSON metadata for the
// struct [DynamicRoutingGetResponseDeployment]
type dynamicRoutingGetResponseDeploymentJSON struct {
	CreatedAt    apijson.Field
	DeploymentID apijson.Field
	VersionID    apijson.Field
	Comment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseDeploymentJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingGetResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingGetResponseElementsObjectOutput].
	Outputs interface{}                           `json:"outputs,required"`
	Type    DynamicRoutingGetResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingGetResponseElementsObjectProperties].
	Properties interface{}                          `json:"properties"`
	JSON       dynamicRoutingGetResponseElementJSON `json:"-"`
	union      DynamicRoutingGetResponseElementsUnion
}

// dynamicRoutingGetResponseElementJSON contains the JSON metadata for the struct
// [DynamicRoutingGetResponseElement]
type dynamicRoutingGetResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingGetResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingGetResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingGetResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingGetResponseElementsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject].
func (r DynamicRoutingGetResponseElement) AsUnion() DynamicRoutingGetResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject],
// [DynamicRoutingGetResponseElementsObject] or
// [DynamicRoutingGetResponseElementsObject].
type DynamicRoutingGetResponseElementsUnion interface {
	implementsDynamicRoutingGetResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingGetResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetResponseElementsObject{}),
		},
	)
}

type DynamicRoutingGetResponseElementsObject struct {
	ID      string                                         `json:"id,required"`
	Outputs DynamicRoutingGetResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingGetResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingGetResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingGetResponseElementsObjectJSON contains the JSON metadata for the
// struct [DynamicRoutingGetResponseElementsObject]
type dynamicRoutingGetResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingGetResponseElementsObject) implementsDynamicRoutingGetResponseElement() {}

type DynamicRoutingGetResponseElementsObjectOutputs struct {
	Next DynamicRoutingGetResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingGetResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingGetResponseElementsObjectOutputsJSON contains the JSON metadata
// for the struct [DynamicRoutingGetResponseElementsObjectOutputs]
type dynamicRoutingGetResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetResponseElementsObjectOutputsNext struct {
	ElementID string                                                 `json:"elementId,required"`
	JSON      dynamicRoutingGetResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingGetResponseElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct [DynamicRoutingGetResponseElementsObjectOutputsNext]
type dynamicRoutingGetResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetResponseElementsObjectType string

const (
	DynamicRoutingGetResponseElementsObjectTypeStart DynamicRoutingGetResponseElementsObjectType = "start"
)

func (r DynamicRoutingGetResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingGetResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingGetResponseElementsType string

const (
	DynamicRoutingGetResponseElementsTypeStart       DynamicRoutingGetResponseElementsType = "start"
	DynamicRoutingGetResponseElementsTypeConditional DynamicRoutingGetResponseElementsType = "conditional"
	DynamicRoutingGetResponseElementsTypePercentage  DynamicRoutingGetResponseElementsType = "percentage"
	DynamicRoutingGetResponseElementsTypeRate        DynamicRoutingGetResponseElementsType = "rate"
	DynamicRoutingGetResponseElementsTypeModel       DynamicRoutingGetResponseElementsType = "model"
	DynamicRoutingGetResponseElementsTypeEnd         DynamicRoutingGetResponseElementsType = "end"
)

func (r DynamicRoutingGetResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingGetResponseElementsTypeStart, DynamicRoutingGetResponseElementsTypeConditional, DynamicRoutingGetResponseElementsTypePercentage, DynamicRoutingGetResponseElementsTypeRate, DynamicRoutingGetResponseElementsTypeModel, DynamicRoutingGetResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingGetResponseVersion struct {
	Active    DynamicRoutingGetResponseVersionActive `json:"active,required"`
	CreatedAt string                                 `json:"created_at,required"`
	Data      string                                 `json:"data,required"`
	VersionID string                                 `json:"version_id,required"`
	Comment   string                                 `json:"comment,nullable"`
	JSON      dynamicRoutingGetResponseVersionJSON   `json:"-"`
}

// dynamicRoutingGetResponseVersionJSON contains the JSON metadata for the struct
// [DynamicRoutingGetResponseVersion]
type dynamicRoutingGetResponseVersionJSON struct {
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseVersionJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetResponseVersionActive string

const (
	DynamicRoutingGetResponseVersionActiveTrue  DynamicRoutingGetResponseVersionActive = "true"
	DynamicRoutingGetResponseVersionActiveFalse DynamicRoutingGetResponseVersionActive = "false"
)

func (r DynamicRoutingGetResponseVersionActive) IsKnown() bool {
	switch r {
	case DynamicRoutingGetResponseVersionActiveTrue, DynamicRoutingGetResponseVersionActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingGetVersionResponse struct {
	ID         string                                    `json:"id,required"`
	AccountTag string                                    `json:"account_tag,required"`
	Active     DynamicRoutingGetVersionResponseActive    `json:"active,required"`
	CreatedAt  string                                    `json:"created_at,required"`
	Data       string                                    `json:"data,required"`
	Elements   []DynamicRoutingGetVersionResponseElement `json:"elements,required"`
	GatewayID  string                                    `json:"gateway_id,required"`
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	Name       string                                    `json:"name,required"`
	VersionID  string                                    `json:"version_id,required"`
	Comment    string                                    `json:"comment,nullable"`
	JSON       dynamicRoutingGetVersionResponseJSON      `json:"-"`
}

// dynamicRoutingGetVersionResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingGetVersionResponse]
type dynamicRoutingGetVersionResponseJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	Elements    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetVersionResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetVersionResponseActive string

const (
	DynamicRoutingGetVersionResponseActiveTrue  DynamicRoutingGetVersionResponseActive = "true"
	DynamicRoutingGetVersionResponseActiveFalse DynamicRoutingGetVersionResponseActive = "false"
)

func (r DynamicRoutingGetVersionResponseActive) IsKnown() bool {
	switch r {
	case DynamicRoutingGetVersionResponseActiveTrue, DynamicRoutingGetVersionResponseActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingGetVersionResponseElement struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [DynamicRoutingGetVersionResponseElementsObjectOutputs],
	// [map[string]DynamicRoutingGetVersionResponseElementsObjectOutput].
	Outputs interface{}                                  `json:"outputs,required"`
	Type    DynamicRoutingGetVersionResponseElementsType `json:"type,required"`
	// This field can have the runtime type of
	// [DynamicRoutingGetVersionResponseElementsObjectProperties].
	Properties interface{}                                 `json:"properties"`
	JSON       dynamicRoutingGetVersionResponseElementJSON `json:"-"`
	union      DynamicRoutingGetVersionResponseElementsUnion
}

// dynamicRoutingGetVersionResponseElementJSON contains the JSON metadata for the
// struct [DynamicRoutingGetVersionResponseElement]
type dynamicRoutingGetVersionResponseElementJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	Properties  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dynamicRoutingGetVersionResponseElementJSON) RawJSON() string {
	return r.raw
}

func (r *DynamicRoutingGetVersionResponseElement) UnmarshalJSON(data []byte) (err error) {
	*r = DynamicRoutingGetVersionResponseElement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DynamicRoutingGetVersionResponseElementsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject].
func (r DynamicRoutingGetVersionResponseElement) AsUnion() DynamicRoutingGetVersionResponseElementsUnion {
	return r.union
}

// Union satisfied by [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject],
// [DynamicRoutingGetVersionResponseElementsObject] or
// [DynamicRoutingGetVersionResponseElementsObject].
type DynamicRoutingGetVersionResponseElementsUnion interface {
	implementsDynamicRoutingGetVersionResponseElement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DynamicRoutingGetVersionResponseElementsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DynamicRoutingGetVersionResponseElementsObject{}),
		},
	)
}

type DynamicRoutingGetVersionResponseElementsObject struct {
	ID      string                                                `json:"id,required"`
	Outputs DynamicRoutingGetVersionResponseElementsObjectOutputs `json:"outputs,required"`
	Type    DynamicRoutingGetVersionResponseElementsObjectType    `json:"type,required"`
	JSON    dynamicRoutingGetVersionResponseElementsObjectJSON    `json:"-"`
}

// dynamicRoutingGetVersionResponseElementsObjectJSON contains the JSON metadata
// for the struct [DynamicRoutingGetVersionResponseElementsObject]
type dynamicRoutingGetVersionResponseElementsObjectJSON struct {
	ID          apijson.Field
	Outputs     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetVersionResponseElementsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetVersionResponseElementsObjectJSON) RawJSON() string {
	return r.raw
}

func (r DynamicRoutingGetVersionResponseElementsObject) implementsDynamicRoutingGetVersionResponseElement() {
}

type DynamicRoutingGetVersionResponseElementsObjectOutputs struct {
	Next DynamicRoutingGetVersionResponseElementsObjectOutputsNext `json:"next,required"`
	JSON dynamicRoutingGetVersionResponseElementsObjectOutputsJSON `json:"-"`
}

// dynamicRoutingGetVersionResponseElementsObjectOutputsJSON contains the JSON
// metadata for the struct [DynamicRoutingGetVersionResponseElementsObjectOutputs]
type dynamicRoutingGetVersionResponseElementsObjectOutputsJSON struct {
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetVersionResponseElementsObjectOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetVersionResponseElementsObjectOutputsJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetVersionResponseElementsObjectOutputsNext struct {
	ElementID string                                                        `json:"elementId,required"`
	JSON      dynamicRoutingGetVersionResponseElementsObjectOutputsNextJSON `json:"-"`
}

// dynamicRoutingGetVersionResponseElementsObjectOutputsNextJSON contains the JSON
// metadata for the struct
// [DynamicRoutingGetVersionResponseElementsObjectOutputsNext]
type dynamicRoutingGetVersionResponseElementsObjectOutputsNextJSON struct {
	ElementID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetVersionResponseElementsObjectOutputsNext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetVersionResponseElementsObjectOutputsNextJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetVersionResponseElementsObjectType string

const (
	DynamicRoutingGetVersionResponseElementsObjectTypeStart DynamicRoutingGetVersionResponseElementsObjectType = "start"
)

func (r DynamicRoutingGetVersionResponseElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingGetVersionResponseElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingGetVersionResponseElementsType string

const (
	DynamicRoutingGetVersionResponseElementsTypeStart       DynamicRoutingGetVersionResponseElementsType = "start"
	DynamicRoutingGetVersionResponseElementsTypeConditional DynamicRoutingGetVersionResponseElementsType = "conditional"
	DynamicRoutingGetVersionResponseElementsTypePercentage  DynamicRoutingGetVersionResponseElementsType = "percentage"
	DynamicRoutingGetVersionResponseElementsTypeRate        DynamicRoutingGetVersionResponseElementsType = "rate"
	DynamicRoutingGetVersionResponseElementsTypeModel       DynamicRoutingGetVersionResponseElementsType = "model"
	DynamicRoutingGetVersionResponseElementsTypeEnd         DynamicRoutingGetVersionResponseElementsType = "end"
)

func (r DynamicRoutingGetVersionResponseElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingGetVersionResponseElementsTypeStart, DynamicRoutingGetVersionResponseElementsTypeConditional, DynamicRoutingGetVersionResponseElementsTypePercentage, DynamicRoutingGetVersionResponseElementsTypeRate, DynamicRoutingGetVersionResponseElementsTypeModel, DynamicRoutingGetVersionResponseElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingListDeploymentsResponse struct {
	Data    DynamicRoutingListDeploymentsResponseData `json:"data,required"`
	Success bool                                      `json:"success,required"`
	JSON    dynamicRoutingListDeploymentsResponseJSON `json:"-"`
}

// dynamicRoutingListDeploymentsResponseJSON contains the JSON metadata for the
// struct [DynamicRoutingListDeploymentsResponse]
type dynamicRoutingListDeploymentsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListDeploymentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListDeploymentsResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListDeploymentsResponseData struct {
	Deployments      []DynamicRoutingListDeploymentsResponseDataDeployment `json:"deployments,required"`
	OrderBy          string                                                `json:"order_by,required"`
	OrderByDirection string                                                `json:"order_by_direction,required"`
	Page             float64                                               `json:"page,required"`
	PerPage          float64                                               `json:"per_page,required"`
	JSON             dynamicRoutingListDeploymentsResponseDataJSON         `json:"-"`
}

// dynamicRoutingListDeploymentsResponseDataJSON contains the JSON metadata for the
// struct [DynamicRoutingListDeploymentsResponseData]
type dynamicRoutingListDeploymentsResponseDataJSON struct {
	Deployments      apijson.Field
	OrderBy          apijson.Field
	OrderByDirection apijson.Field
	Page             apijson.Field
	PerPage          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DynamicRoutingListDeploymentsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListDeploymentsResponseDataJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListDeploymentsResponseDataDeployment struct {
	CreatedAt    string                                                  `json:"created_at,required"`
	DeploymentID string                                                  `json:"deployment_id,required"`
	VersionID    string                                                  `json:"version_id,required"`
	Comment      string                                                  `json:"comment,nullable"`
	JSON         dynamicRoutingListDeploymentsResponseDataDeploymentJSON `json:"-"`
}

// dynamicRoutingListDeploymentsResponseDataDeploymentJSON contains the JSON
// metadata for the struct [DynamicRoutingListDeploymentsResponseDataDeployment]
type dynamicRoutingListDeploymentsResponseDataDeploymentJSON struct {
	CreatedAt    apijson.Field
	DeploymentID apijson.Field
	VersionID    apijson.Field
	Comment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DynamicRoutingListDeploymentsResponseDataDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListDeploymentsResponseDataDeploymentJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListVersionsResponse struct {
	Data    DynamicRoutingListVersionsResponseData `json:"data,required"`
	Success bool                                   `json:"success,required"`
	JSON    dynamicRoutingListVersionsResponseJSON `json:"-"`
}

// dynamicRoutingListVersionsResponseJSON contains the JSON metadata for the struct
// [DynamicRoutingListVersionsResponse]
type dynamicRoutingListVersionsResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListVersionsResponseJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListVersionsResponseData struct {
	OrderBy          string                                          `json:"order_by,required"`
	OrderByDirection string                                          `json:"order_by_direction,required"`
	Page             float64                                         `json:"page,required"`
	PerPage          float64                                         `json:"per_page,required"`
	Versions         []DynamicRoutingListVersionsResponseDataVersion `json:"versions,required"`
	JSON             dynamicRoutingListVersionsResponseDataJSON      `json:"-"`
}

// dynamicRoutingListVersionsResponseDataJSON contains the JSON metadata for the
// struct [DynamicRoutingListVersionsResponseData]
type dynamicRoutingListVersionsResponseDataJSON struct {
	OrderBy          apijson.Field
	OrderByDirection apijson.Field
	Page             apijson.Field
	PerPage          apijson.Field
	Versions         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DynamicRoutingListVersionsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListVersionsResponseDataJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListVersionsResponseDataVersion struct {
	Active    DynamicRoutingListVersionsResponseDataVersionsActive `json:"active,required"`
	CreatedAt string                                               `json:"created_at,required"`
	Data      string                                               `json:"data,required"`
	VersionID string                                               `json:"version_id,required"`
	Comment   string                                               `json:"comment,nullable"`
	JSON      dynamicRoutingListVersionsResponseDataVersionJSON    `json:"-"`
}

// dynamicRoutingListVersionsResponseDataVersionJSON contains the JSON metadata for
// the struct [DynamicRoutingListVersionsResponseDataVersion]
type dynamicRoutingListVersionsResponseDataVersionJSON struct {
	Active      apijson.Field
	CreatedAt   apijson.Field
	Data        apijson.Field
	VersionID   apijson.Field
	Comment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingListVersionsResponseDataVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingListVersionsResponseDataVersionJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListVersionsResponseDataVersionsActive string

const (
	DynamicRoutingListVersionsResponseDataVersionsActiveTrue  DynamicRoutingListVersionsResponseDataVersionsActive = "true"
	DynamicRoutingListVersionsResponseDataVersionsActiveFalse DynamicRoutingListVersionsResponseDataVersionsActive = "false"
)

func (r DynamicRoutingListVersionsResponseDataVersionsActive) IsKnown() bool {
	switch r {
	case DynamicRoutingListVersionsResponseDataVersionsActiveTrue, DynamicRoutingListVersionsResponseDataVersionsActiveFalse:
		return true
	}
	return false
}

type DynamicRoutingNewParams struct {
	AccountID param.Field[string]                                `path:"account_id,required"`
	Elements  param.Field[[]DynamicRoutingNewParamsElementUnion] `json:"elements,required"`
	Name      param.Field[string]                                `json:"name,required"`
}

func (r DynamicRoutingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewParamsElement struct {
	ID         param.Field[string]                              `json:"id,required"`
	Outputs    param.Field[interface{}]                         `json:"outputs,required"`
	Type       param.Field[DynamicRoutingNewParamsElementsType] `json:"type,required"`
	Properties param.Field[interface{}]                         `json:"properties"`
}

func (r DynamicRoutingNewParamsElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DynamicRoutingNewParamsElement) implementsDynamicRoutingNewParamsElementUnion() {}

// Satisfied by [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [ai_gateway.DynamicRoutingNewParamsElementsObject],
// [DynamicRoutingNewParamsElement].
type DynamicRoutingNewParamsElementUnion interface {
	implementsDynamicRoutingNewParamsElementUnion()
}

type DynamicRoutingNewParamsElementsObject struct {
	ID      param.Field[string]                                       `json:"id,required"`
	Outputs param.Field[DynamicRoutingNewParamsElementsObjectOutputs] `json:"outputs,required"`
	Type    param.Field[DynamicRoutingNewParamsElementsObjectType]    `json:"type,required"`
}

func (r DynamicRoutingNewParamsElementsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DynamicRoutingNewParamsElementsObject) implementsDynamicRoutingNewParamsElementUnion() {}

type DynamicRoutingNewParamsElementsObjectOutputs struct {
	Next param.Field[DynamicRoutingNewParamsElementsObjectOutputsNext] `json:"next,required"`
}

func (r DynamicRoutingNewParamsElementsObjectOutputs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewParamsElementsObjectOutputsNext struct {
	ElementID param.Field[string] `json:"elementId,required"`
}

func (r DynamicRoutingNewParamsElementsObjectOutputsNext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewParamsElementsObjectType string

const (
	DynamicRoutingNewParamsElementsObjectTypeStart DynamicRoutingNewParamsElementsObjectType = "start"
)

func (r DynamicRoutingNewParamsElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewParamsElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingNewParamsElementsType string

const (
	DynamicRoutingNewParamsElementsTypeStart       DynamicRoutingNewParamsElementsType = "start"
	DynamicRoutingNewParamsElementsTypeConditional DynamicRoutingNewParamsElementsType = "conditional"
	DynamicRoutingNewParamsElementsTypePercentage  DynamicRoutingNewParamsElementsType = "percentage"
	DynamicRoutingNewParamsElementsTypeRate        DynamicRoutingNewParamsElementsType = "rate"
	DynamicRoutingNewParamsElementsTypeModel       DynamicRoutingNewParamsElementsType = "model"
	DynamicRoutingNewParamsElementsTypeEnd         DynamicRoutingNewParamsElementsType = "end"
)

func (r DynamicRoutingNewParamsElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewParamsElementsTypeStart, DynamicRoutingNewParamsElementsTypeConditional, DynamicRoutingNewParamsElementsTypePercentage, DynamicRoutingNewParamsElementsTypeRate, DynamicRoutingNewParamsElementsTypeModel, DynamicRoutingNewParamsElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingNewResponseEnvelope struct {
	Result  DynamicRoutingNewResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    dynamicRoutingNewResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DynamicRoutingNewResponseEnvelope]
type dynamicRoutingNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r DynamicRoutingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DynamicRoutingDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DynamicRoutingDeleteResponseEnvelope struct {
	Result  DynamicRoutingDeleteResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    dynamicRoutingDeleteResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DynamicRoutingDeleteResponseEnvelope]
type dynamicRoutingDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewDeploymentParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Comment   param.Field[string] `json:"comment,required"`
	VersionID param.Field[string] `json:"version_id,required"`
}

func (r DynamicRoutingNewDeploymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewDeploymentResponseEnvelope struct {
	Result  DynamicRoutingNewDeploymentResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    dynamicRoutingNewDeploymentResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingNewDeploymentResponseEnvelopeJSON contains the JSON metadata for
// the struct [DynamicRoutingNewDeploymentResponseEnvelope]
type dynamicRoutingNewDeploymentResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewDeploymentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewDeploymentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingNewVersionParams struct {
	AccountID param.Field[string]                                       `path:"account_id,required"`
	Comment   param.Field[string]                                       `json:"comment,required"`
	Elements  param.Field[[]DynamicRoutingNewVersionParamsElementUnion] `json:"elements,required"`
}

func (r DynamicRoutingNewVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewVersionParamsElement struct {
	ID         param.Field[string]                                     `json:"id,required"`
	Outputs    param.Field[interface{}]                                `json:"outputs,required"`
	Type       param.Field[DynamicRoutingNewVersionParamsElementsType] `json:"type,required"`
	Properties param.Field[interface{}]                                `json:"properties"`
}

func (r DynamicRoutingNewVersionParamsElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DynamicRoutingNewVersionParamsElement) implementsDynamicRoutingNewVersionParamsElementUnion() {
}

// Satisfied by [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [ai_gateway.DynamicRoutingNewVersionParamsElementsObject],
// [DynamicRoutingNewVersionParamsElement].
type DynamicRoutingNewVersionParamsElementUnion interface {
	implementsDynamicRoutingNewVersionParamsElementUnion()
}

type DynamicRoutingNewVersionParamsElementsObject struct {
	ID      param.Field[string]                                              `json:"id,required"`
	Outputs param.Field[DynamicRoutingNewVersionParamsElementsObjectOutputs] `json:"outputs,required"`
	Type    param.Field[DynamicRoutingNewVersionParamsElementsObjectType]    `json:"type,required"`
}

func (r DynamicRoutingNewVersionParamsElementsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DynamicRoutingNewVersionParamsElementsObject) implementsDynamicRoutingNewVersionParamsElementUnion() {
}

type DynamicRoutingNewVersionParamsElementsObjectOutputs struct {
	Next param.Field[DynamicRoutingNewVersionParamsElementsObjectOutputsNext] `json:"next,required"`
}

func (r DynamicRoutingNewVersionParamsElementsObjectOutputs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewVersionParamsElementsObjectOutputsNext struct {
	ElementID param.Field[string] `json:"elementId,required"`
}

func (r DynamicRoutingNewVersionParamsElementsObjectOutputsNext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DynamicRoutingNewVersionParamsElementsObjectType string

const (
	DynamicRoutingNewVersionParamsElementsObjectTypeStart DynamicRoutingNewVersionParamsElementsObjectType = "start"
)

func (r DynamicRoutingNewVersionParamsElementsObjectType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewVersionParamsElementsObjectTypeStart:
		return true
	}
	return false
}

type DynamicRoutingNewVersionParamsElementsType string

const (
	DynamicRoutingNewVersionParamsElementsTypeStart       DynamicRoutingNewVersionParamsElementsType = "start"
	DynamicRoutingNewVersionParamsElementsTypeConditional DynamicRoutingNewVersionParamsElementsType = "conditional"
	DynamicRoutingNewVersionParamsElementsTypePercentage  DynamicRoutingNewVersionParamsElementsType = "percentage"
	DynamicRoutingNewVersionParamsElementsTypeRate        DynamicRoutingNewVersionParamsElementsType = "rate"
	DynamicRoutingNewVersionParamsElementsTypeModel       DynamicRoutingNewVersionParamsElementsType = "model"
	DynamicRoutingNewVersionParamsElementsTypeEnd         DynamicRoutingNewVersionParamsElementsType = "end"
)

func (r DynamicRoutingNewVersionParamsElementsType) IsKnown() bool {
	switch r {
	case DynamicRoutingNewVersionParamsElementsTypeStart, DynamicRoutingNewVersionParamsElementsTypeConditional, DynamicRoutingNewVersionParamsElementsTypePercentage, DynamicRoutingNewVersionParamsElementsTypeRate, DynamicRoutingNewVersionParamsElementsTypeModel, DynamicRoutingNewVersionParamsElementsTypeEnd:
		return true
	}
	return false
}

type DynamicRoutingNewVersionResponseEnvelope struct {
	Result  DynamicRoutingNewVersionResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    dynamicRoutingNewVersionResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingNewVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [DynamicRoutingNewVersionResponseEnvelope]
type dynamicRoutingNewVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingNewVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingNewVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DynamicRoutingGetResponseEnvelope struct {
	Result  DynamicRoutingGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    dynamicRoutingGetResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DynamicRoutingGetResponseEnvelope]
type dynamicRoutingGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingGetVersionParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DynamicRoutingGetVersionResponseEnvelope struct {
	Result  DynamicRoutingGetVersionResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    dynamicRoutingGetVersionResponseEnvelopeJSON `json:"-"`
}

// dynamicRoutingGetVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [DynamicRoutingGetVersionResponseEnvelope]
type dynamicRoutingGetVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DynamicRoutingGetVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dynamicRoutingGetVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DynamicRoutingListDeploymentsParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DynamicRoutingListVersionsParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
