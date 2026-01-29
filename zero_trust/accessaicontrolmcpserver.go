// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// AccessAIControlMcpServerService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessAIControlMcpServerService] method instead.
type AccessAIControlMcpServerService struct {
	Options []option.RequestOption
}

// NewAccessAIControlMcpServerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAIControlMcpServerService(opts ...option.RequestOption) (r *AccessAIControlMcpServerService) {
	r = &AccessAIControlMcpServerService{}
	r.Options = opts
	return
}

// Create a new MCP Server
func (r *AccessAIControlMcpServerService) New(ctx context.Context, params AccessAIControlMcpServerNewParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerNewResponse, err error) {
	var env AccessAIControlMcpServerNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a MCP Server
func (r *AccessAIControlMcpServerService) Update(ctx context.Context, id string, params AccessAIControlMcpServerUpdateParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerUpdateResponse, err error) {
	var env AccessAIControlMcpServerUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List MCP Servers
func (r *AccessAIControlMcpServerService) List(ctx context.Context, params AccessAIControlMcpServerListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessAIControlMcpServerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers", params.AccountID)
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

// List MCP Servers
func (r *AccessAIControlMcpServerService) ListAutoPaging(ctx context.Context, params AccessAIControlMcpServerListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessAIControlMcpServerListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a MCP Server
func (r *AccessAIControlMcpServerService) Delete(ctx context.Context, id string, body AccessAIControlMcpServerDeleteParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerDeleteResponse, err error) {
	var env AccessAIControlMcpServerDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read the details of a MCP Server
func (r *AccessAIControlMcpServerService) Read(ctx context.Context, id string, query AccessAIControlMcpServerReadParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerReadResponse, err error) {
	var env AccessAIControlMcpServerReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Sync MCP Server Capabilities
func (r *AccessAIControlMcpServerService) Sync(ctx context.Context, id string, body AccessAIControlMcpServerSyncParams, opts ...option.RequestOption) (res *AccessAIControlMcpServerSyncResponse, err error) {
	var env AccessAIControlMcpServerSyncResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/servers/%s/sync", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessAIControlMcpServerNewResponse struct {
	// server id
	ID                 string                                      `json:"id,required"`
	AuthType           AccessAIControlMcpServerNewResponseAuthType `json:"auth_type,required"`
	Hostname           string                                      `json:"hostname,required" format:"uri"`
	Name               string                                      `json:"name,required"`
	Prompts            []map[string]interface{}                    `json:"prompts,required"`
	Tools              []map[string]interface{}                    `json:"tools,required"`
	CreatedAt          time.Time                                   `json:"created_at" format:"date-time"`
	CreatedBy          string                                      `json:"created_by"`
	Description        string                                      `json:"description,nullable"`
	Error              string                                      `json:"error"`
	LastSuccessfulSync time.Time                                   `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                   `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                   `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                      `json:"modified_by"`
	Status             string                                      `json:"status"`
	JSON               accessAIControlMcpServerNewResponseJSON     `json:"-"`
}

// accessAIControlMcpServerNewResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerNewResponse]
type accessAIControlMcpServerNewResponseJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerNewResponseAuthType string

const (
	AccessAIControlMcpServerNewResponseAuthTypeOAuth           AccessAIControlMcpServerNewResponseAuthType = "oauth"
	AccessAIControlMcpServerNewResponseAuthTypeBearer          AccessAIControlMcpServerNewResponseAuthType = "bearer"
	AccessAIControlMcpServerNewResponseAuthTypeUnauthenticated AccessAIControlMcpServerNewResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerNewResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewResponseAuthTypeOAuth, AccessAIControlMcpServerNewResponseAuthTypeBearer, AccessAIControlMcpServerNewResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerUpdateResponse struct {
	// server id
	ID                 string                                         `json:"id,required"`
	AuthType           AccessAIControlMcpServerUpdateResponseAuthType `json:"auth_type,required"`
	Hostname           string                                         `json:"hostname,required" format:"uri"`
	Name               string                                         `json:"name,required"`
	Prompts            []map[string]interface{}                       `json:"prompts,required"`
	Tools              []map[string]interface{}                       `json:"tools,required"`
	CreatedAt          time.Time                                      `json:"created_at" format:"date-time"`
	CreatedBy          string                                         `json:"created_by"`
	Description        string                                         `json:"description,nullable"`
	Error              string                                         `json:"error"`
	LastSuccessfulSync time.Time                                      `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                      `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                      `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                         `json:"modified_by"`
	Status             string                                         `json:"status"`
	JSON               accessAIControlMcpServerUpdateResponseJSON     `json:"-"`
}

// accessAIControlMcpServerUpdateResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerUpdateResponse]
type accessAIControlMcpServerUpdateResponseJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateResponseAuthType string

const (
	AccessAIControlMcpServerUpdateResponseAuthTypeOAuth           AccessAIControlMcpServerUpdateResponseAuthType = "oauth"
	AccessAIControlMcpServerUpdateResponseAuthTypeBearer          AccessAIControlMcpServerUpdateResponseAuthType = "bearer"
	AccessAIControlMcpServerUpdateResponseAuthTypeUnauthenticated AccessAIControlMcpServerUpdateResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerUpdateResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerUpdateResponseAuthTypeOAuth, AccessAIControlMcpServerUpdateResponseAuthTypeBearer, AccessAIControlMcpServerUpdateResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerListResponse struct {
	// server id
	ID                 string                                       `json:"id,required"`
	AuthType           AccessAIControlMcpServerListResponseAuthType `json:"auth_type,required"`
	Hostname           string                                       `json:"hostname,required" format:"uri"`
	Name               string                                       `json:"name,required"`
	Prompts            []map[string]interface{}                     `json:"prompts,required"`
	Tools              []map[string]interface{}                     `json:"tools,required"`
	CreatedAt          time.Time                                    `json:"created_at" format:"date-time"`
	CreatedBy          string                                       `json:"created_by"`
	Description        string                                       `json:"description,nullable"`
	Error              string                                       `json:"error"`
	LastSuccessfulSync time.Time                                    `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                    `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                    `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                       `json:"modified_by"`
	Status             string                                       `json:"status"`
	JSON               accessAIControlMcpServerListResponseJSON     `json:"-"`
}

// accessAIControlMcpServerListResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerListResponse]
type accessAIControlMcpServerListResponseJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpServerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListResponseAuthType string

const (
	AccessAIControlMcpServerListResponseAuthTypeOAuth           AccessAIControlMcpServerListResponseAuthType = "oauth"
	AccessAIControlMcpServerListResponseAuthTypeBearer          AccessAIControlMcpServerListResponseAuthType = "bearer"
	AccessAIControlMcpServerListResponseAuthTypeUnauthenticated AccessAIControlMcpServerListResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerListResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerListResponseAuthTypeOAuth, AccessAIControlMcpServerListResponseAuthTypeBearer, AccessAIControlMcpServerListResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerDeleteResponse struct {
	// server id
	ID                 string                                         `json:"id,required"`
	AuthType           AccessAIControlMcpServerDeleteResponseAuthType `json:"auth_type,required"`
	Hostname           string                                         `json:"hostname,required" format:"uri"`
	Name               string                                         `json:"name,required"`
	Prompts            []map[string]interface{}                       `json:"prompts,required"`
	Tools              []map[string]interface{}                       `json:"tools,required"`
	CreatedAt          time.Time                                      `json:"created_at" format:"date-time"`
	CreatedBy          string                                         `json:"created_by"`
	Description        string                                         `json:"description,nullable"`
	Error              string                                         `json:"error"`
	LastSuccessfulSync time.Time                                      `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                      `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                      `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                         `json:"modified_by"`
	Status             string                                         `json:"status"`
	JSON               accessAIControlMcpServerDeleteResponseJSON     `json:"-"`
}

// accessAIControlMcpServerDeleteResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerDeleteResponse]
type accessAIControlMcpServerDeleteResponseJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerDeleteResponseAuthType string

const (
	AccessAIControlMcpServerDeleteResponseAuthTypeOAuth           AccessAIControlMcpServerDeleteResponseAuthType = "oauth"
	AccessAIControlMcpServerDeleteResponseAuthTypeBearer          AccessAIControlMcpServerDeleteResponseAuthType = "bearer"
	AccessAIControlMcpServerDeleteResponseAuthTypeUnauthenticated AccessAIControlMcpServerDeleteResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerDeleteResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerDeleteResponseAuthTypeOAuth, AccessAIControlMcpServerDeleteResponseAuthTypeBearer, AccessAIControlMcpServerDeleteResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerReadResponse struct {
	// server id
	ID                 string                                       `json:"id,required"`
	AuthType           AccessAIControlMcpServerReadResponseAuthType `json:"auth_type,required"`
	Hostname           string                                       `json:"hostname,required" format:"uri"`
	Name               string                                       `json:"name,required"`
	Prompts            []map[string]interface{}                     `json:"prompts,required"`
	Tools              []map[string]interface{}                     `json:"tools,required"`
	CreatedAt          time.Time                                    `json:"created_at" format:"date-time"`
	CreatedBy          string                                       `json:"created_by"`
	Description        string                                       `json:"description,nullable"`
	Error              string                                       `json:"error"`
	LastSuccessfulSync time.Time                                    `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                    `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                    `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                       `json:"modified_by"`
	Status             string                                       `json:"status"`
	JSON               accessAIControlMcpServerReadResponseJSON     `json:"-"`
}

// accessAIControlMcpServerReadResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpServerReadResponse]
type accessAIControlMcpServerReadResponseJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadResponseAuthType string

const (
	AccessAIControlMcpServerReadResponseAuthTypeOAuth           AccessAIControlMcpServerReadResponseAuthType = "oauth"
	AccessAIControlMcpServerReadResponseAuthTypeBearer          AccessAIControlMcpServerReadResponseAuthType = "bearer"
	AccessAIControlMcpServerReadResponseAuthTypeUnauthenticated AccessAIControlMcpServerReadResponseAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerReadResponseAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerReadResponseAuthTypeOAuth, AccessAIControlMcpServerReadResponseAuthTypeBearer, AccessAIControlMcpServerReadResponseAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerSyncResponse = interface{}

type AccessAIControlMcpServerNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// server id
	ID              param.Field[string]                                    `json:"id,required"`
	AuthType        param.Field[AccessAIControlMcpServerNewParamsAuthType] `json:"auth_type,required"`
	Hostname        param.Field[string]                                    `json:"hostname,required" format:"uri"`
	Name            param.Field[string]                                    `json:"name,required"`
	AuthCredentials param.Field[string]                                    `json:"auth_credentials"`
	Description     param.Field[string]                                    `json:"description"`
}

func (r AccessAIControlMcpServerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerNewParamsAuthType string

const (
	AccessAIControlMcpServerNewParamsAuthTypeOAuth           AccessAIControlMcpServerNewParamsAuthType = "oauth"
	AccessAIControlMcpServerNewParamsAuthTypeBearer          AccessAIControlMcpServerNewParamsAuthType = "bearer"
	AccessAIControlMcpServerNewParamsAuthTypeUnauthenticated AccessAIControlMcpServerNewParamsAuthType = "unauthenticated"
)

func (r AccessAIControlMcpServerNewParamsAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpServerNewParamsAuthTypeOAuth, AccessAIControlMcpServerNewParamsAuthTypeBearer, AccessAIControlMcpServerNewParamsAuthTypeUnauthenticated:
		return true
	}
	return false
}

type AccessAIControlMcpServerNewResponseEnvelope struct {
	Result  AccessAIControlMcpServerNewResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    accessAIControlMcpServerNewResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerNewResponseEnvelope]
type accessAIControlMcpServerNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerUpdateParams struct {
	AccountID       param.Field[string] `path:"account_id,required"`
	AuthCredentials param.Field[string] `json:"auth_credentials"`
	Description     param.Field[string] `json:"description"`
	Name            param.Field[string] `json:"name"`
}

func (r AccessAIControlMcpServerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpServerUpdateResponseEnvelope struct {
	Result  AccessAIControlMcpServerUpdateResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accessAIControlMcpServerUpdateResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerUpdateResponseEnvelope]
type accessAIControlMcpServerUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessAIControlMcpServerListParams]'s query parameters as
// `url.Values`.
func (r AccessAIControlMcpServerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessAIControlMcpServerDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessAIControlMcpServerDeleteResponseEnvelope struct {
	Result  AccessAIControlMcpServerDeleteResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accessAIControlMcpServerDeleteResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpServerDeleteResponseEnvelope]
type accessAIControlMcpServerDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerReadParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessAIControlMcpServerReadResponseEnvelope struct {
	Result  AccessAIControlMcpServerReadResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    accessAIControlMcpServerReadResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerReadResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerReadResponseEnvelope]
type accessAIControlMcpServerReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpServerSyncParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessAIControlMcpServerSyncResponseEnvelope struct {
	Result  AccessAIControlMcpServerSyncResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    accessAIControlMcpServerSyncResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpServerSyncResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpServerSyncResponseEnvelope]
type accessAIControlMcpServerSyncResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpServerSyncResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpServerSyncResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
