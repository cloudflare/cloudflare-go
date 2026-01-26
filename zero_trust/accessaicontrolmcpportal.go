// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/tidwall/gjson"
)

// AccessAIControlMcpPortalService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessAIControlMcpPortalService] method instead.
type AccessAIControlMcpPortalService struct {
	Options []option.RequestOption
}

// NewAccessAIControlMcpPortalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAIControlMcpPortalService(opts ...option.RequestOption) (r *AccessAIControlMcpPortalService) {
	r = &AccessAIControlMcpPortalService{}
	r.Options = opts
	return
}

// Create a new MCP Portal
func (r *AccessAIControlMcpPortalService) New(ctx context.Context, params AccessAIControlMcpPortalNewParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalNewResponse, err error) {
	var env AccessAIControlMcpPortalNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a MCP Portal
func (r *AccessAIControlMcpPortalService) Update(ctx context.Context, id string, params AccessAIControlMcpPortalUpdateParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalUpdateResponse, err error) {
	var env AccessAIControlMcpPortalUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List MCP Portals
func (r *AccessAIControlMcpPortalService) List(ctx context.Context, params AccessAIControlMcpPortalListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessAIControlMcpPortalListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals", params.AccountID)
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

// List MCP Portals
func (r *AccessAIControlMcpPortalService) ListAutoPaging(ctx context.Context, params AccessAIControlMcpPortalListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessAIControlMcpPortalListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a MCP Portal
func (r *AccessAIControlMcpPortalService) Delete(ctx context.Context, id string, body AccessAIControlMcpPortalDeleteParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalDeleteResponse, err error) {
	var env AccessAIControlMcpPortalDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Read details of an MCP Portal
func (r *AccessAIControlMcpPortalService) Read(ctx context.Context, id string, query AccessAIControlMcpPortalReadParams, opts ...option.RequestOption) (res *AccessAIControlMcpPortalReadResponse, err error) {
	var env AccessAIControlMcpPortalReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/ai-controls/mcp/portals/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessAIControlMcpPortalNewResponse struct {
	// portal id
	ID          string    `json:"id,required"`
	Hostname    string    `json:"hostname,required"`
	Name        string    `json:"name,required"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	CreatedBy   string    `json:"created_by"`
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway bool                                    `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalNewResponseJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalNewResponse]
type accessAIControlMcpPortalNewResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateResponse struct {
	// portal id
	ID          string    `json:"id,required"`
	Hostname    string    `json:"hostname,required"`
	Name        string    `json:"name,required"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	CreatedBy   string    `json:"created_by"`
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway bool                                       `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalUpdateResponseJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalUpdateResponse]
type accessAIControlMcpPortalUpdateResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListResponse struct {
	// portal id
	ID          string    `json:"id,required"`
	Hostname    string    `json:"hostname,required"`
	Name        string    `json:"name,required"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	CreatedBy   string    `json:"created_by"`
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway bool                                     `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalListResponseJSON `json:"-"`
}

// accessAIControlMcpPortalListResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalListResponse]
type accessAIControlMcpPortalListResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalDeleteResponse struct {
	// portal id
	ID          string    `json:"id,required"`
	Hostname    string    `json:"hostname,required"`
	Name        string    `json:"name,required"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	CreatedBy   string    `json:"created_by"`
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modified_at" format:"date-time"`
	ModifiedBy  string    `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway bool                                       `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalDeleteResponseJSON `json:"-"`
}

// accessAIControlMcpPortalDeleteResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalDeleteResponse]
type accessAIControlMcpPortalDeleteResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponse struct {
	// portal id
	ID          string                                       `json:"id,required"`
	Hostname    string                                       `json:"hostname,required"`
	Name        string                                       `json:"name,required"`
	Servers     []AccessAIControlMcpPortalReadResponseServer `json:"servers,required"`
	CreatedAt   time.Time                                    `json:"created_at" format:"date-time"`
	CreatedBy   string                                       `json:"created_by"`
	Description string                                       `json:"description"`
	ModifiedAt  time.Time                                    `json:"modified_at" format:"date-time"`
	ModifiedBy  string                                       `json:"modified_by"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway bool                                     `json:"secure_web_gateway"`
	JSON             accessAIControlMcpPortalReadResponseJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseJSON contains the JSON metadata for the
// struct [AccessAIControlMcpPortalReadResponse]
type accessAIControlMcpPortalReadResponseJSON struct {
	ID               apijson.Field
	Hostname         apijson.Field
	Name             apijson.Field
	Servers          apijson.Field
	CreatedAt        apijson.Field
	CreatedBy        apijson.Field
	Description      apijson.Field
	ModifiedAt       apijson.Field
	ModifiedBy       apijson.Field
	SecureWebGateway apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServer struct {
	// server id
	ID                 string                                                                      `json:"id,required"`
	AuthType           AccessAIControlMcpPortalReadResponseServersAuthType                         `json:"auth_type,required"`
	Hostname           string                                                                      `json:"hostname,required" format:"uri"`
	Name               string                                                                      `json:"name,required"`
	Prompts            []map[string]interface{}                                                    `json:"prompts,required"`
	Tools              []map[string]interface{}                                                    `json:"tools,required"`
	UpdatedPrompts     []map[string]AccessAIControlMcpPortalReadResponseServersUpdatedPromptsUnion `json:"updated_prompts,required"`
	UpdatedTools       []map[string]AccessAIControlMcpPortalReadResponseServersUpdatedToolsUnion   `json:"updated_tools,required"`
	CreatedAt          time.Time                                                                   `json:"created_at" format:"date-time"`
	CreatedBy          string                                                                      `json:"created_by"`
	DefaultDisabled    bool                                                                        `json:"default_disabled"`
	Description        string                                                                      `json:"description,nullable"`
	Error              string                                                                      `json:"error"`
	LastSuccessfulSync time.Time                                                                   `json:"last_successful_sync" format:"date-time"`
	LastSynced         time.Time                                                                   `json:"last_synced" format:"date-time"`
	ModifiedAt         time.Time                                                                   `json:"modified_at" format:"date-time"`
	ModifiedBy         string                                                                      `json:"modified_by"`
	OnBehalf           bool                                                                        `json:"on_behalf"`
	Status             string                                                                      `json:"status"`
	JSON               accessAIControlMcpPortalReadResponseServerJSON                              `json:"-"`
}

// accessAIControlMcpPortalReadResponseServerJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalReadResponseServer]
type accessAIControlMcpPortalReadResponseServerJSON struct {
	ID                 apijson.Field
	AuthType           apijson.Field
	Hostname           apijson.Field
	Name               apijson.Field
	Prompts            apijson.Field
	Tools              apijson.Field
	UpdatedPrompts     apijson.Field
	UpdatedTools       apijson.Field
	CreatedAt          apijson.Field
	CreatedBy          apijson.Field
	DefaultDisabled    apijson.Field
	Description        apijson.Field
	Error              apijson.Field
	LastSuccessfulSync apijson.Field
	LastSynced         apijson.Field
	ModifiedAt         apijson.Field
	ModifiedBy         apijson.Field
	OnBehalf           apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseServerJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadResponseServersAuthType string

const (
	AccessAIControlMcpPortalReadResponseServersAuthTypeOAuth           AccessAIControlMcpPortalReadResponseServersAuthType = "oauth"
	AccessAIControlMcpPortalReadResponseServersAuthTypeBearer          AccessAIControlMcpPortalReadResponseServersAuthType = "bearer"
	AccessAIControlMcpPortalReadResponseServersAuthTypeUnauthenticated AccessAIControlMcpPortalReadResponseServersAuthType = "unauthenticated"
)

func (r AccessAIControlMcpPortalReadResponseServersAuthType) IsKnown() bool {
	switch r {
	case AccessAIControlMcpPortalReadResponseServersAuthTypeOAuth, AccessAIControlMcpPortalReadResponseServersAuthTypeBearer, AccessAIControlMcpPortalReadResponseServersAuthTypeUnauthenticated:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type AccessAIControlMcpPortalReadResponseServersUpdatedPromptsUnion interface {
	ImplementsAccessAIControlMcpPortalReadResponseServersUpdatedPromptsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAIControlMcpPortalReadResponseServersUpdatedPromptsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type AccessAIControlMcpPortalReadResponseServersUpdatedToolsUnion interface {
	ImplementsAccessAIControlMcpPortalReadResponseServersUpdatedToolsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAIControlMcpPortalReadResponseServersUpdatedToolsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessAIControlMcpPortalNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// portal id
	ID          param.Field[string] `json:"id,required"`
	Hostname    param.Field[string] `json:"hostname,required"`
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway param.Field[bool]                                      `json:"secure_web_gateway"`
	Servers          param.Field[[]AccessAIControlMcpPortalNewParamsServer] `json:"servers"`
}

func (r AccessAIControlMcpPortalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewParamsServer struct {
	// server id
	ServerID        param.Field[string]                                                  `json:"server_id,required"`
	DefaultDisabled param.Field[bool]                                                    `json:"default_disabled"`
	OnBehalf        param.Field[bool]                                                    `json:"on_behalf"`
	UpdatedPrompts  param.Field[[]AccessAIControlMcpPortalNewParamsServersUpdatedPrompt] `json:"updated_prompts"`
	UpdatedTools    param.Field[[]AccessAIControlMcpPortalNewParamsServersUpdatedTool]   `json:"updated_tools"`
}

func (r AccessAIControlMcpPortalNewParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewParamsServersUpdatedPrompt struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r AccessAIControlMcpPortalNewParamsServersUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewParamsServersUpdatedTool struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r AccessAIControlMcpPortalNewParamsServersUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalNewResponseEnvelope struct {
	Result  AccessAIControlMcpPortalNewResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    accessAIControlMcpPortalNewResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalNewResponseEnvelope]
type accessAIControlMcpPortalNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Description param.Field[string] `json:"description"`
	Hostname    param.Field[string] `json:"hostname"`
	Name        param.Field[string] `json:"name"`
	// Route outbound MCP traffic through Zero Trust Secure Web Gateway
	SecureWebGateway param.Field[bool]                                         `json:"secure_web_gateway"`
	Servers          param.Field[[]AccessAIControlMcpPortalUpdateParamsServer] `json:"servers"`
}

func (r AccessAIControlMcpPortalUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateParamsServer struct {
	// server id
	ServerID        param.Field[string]                                                     `json:"server_id,required"`
	DefaultDisabled param.Field[bool]                                                       `json:"default_disabled"`
	OnBehalf        param.Field[bool]                                                       `json:"on_behalf"`
	UpdatedPrompts  param.Field[[]AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt] `json:"updated_prompts"`
	UpdatedTools    param.Field[[]AccessAIControlMcpPortalUpdateParamsServersUpdatedTool]   `json:"updated_tools"`
}

func (r AccessAIControlMcpPortalUpdateParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r AccessAIControlMcpPortalUpdateParamsServersUpdatedPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateParamsServersUpdatedTool struct {
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r AccessAIControlMcpPortalUpdateParamsServersUpdatedTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAIControlMcpPortalUpdateResponseEnvelope struct {
	Result  AccessAIControlMcpPortalUpdateResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accessAIControlMcpPortalUpdateResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpPortalUpdateResponseEnvelope]
type accessAIControlMcpPortalUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name, hostname
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessAIControlMcpPortalListParams]'s query parameters as
// `url.Values`.
func (r AccessAIControlMcpPortalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessAIControlMcpPortalDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessAIControlMcpPortalDeleteResponseEnvelope struct {
	Result  AccessAIControlMcpPortalDeleteResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accessAIControlMcpPortalDeleteResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessAIControlMcpPortalDeleteResponseEnvelope]
type accessAIControlMcpPortalDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessAIControlMcpPortalReadParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessAIControlMcpPortalReadResponseEnvelope struct {
	Result  AccessAIControlMcpPortalReadResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    accessAIControlMcpPortalReadResponseEnvelopeJSON `json:"-"`
}

// accessAIControlMcpPortalReadResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessAIControlMcpPortalReadResponseEnvelope]
type accessAIControlMcpPortalReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAIControlMcpPortalReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAIControlMcpPortalReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
