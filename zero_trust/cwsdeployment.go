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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CWSDeploymentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSDeploymentService] method instead.
type CWSDeploymentService struct {
	Options []option.RequestOption
}

// NewCWSDeploymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSDeploymentService(opts ...option.RequestOption) (r *CWSDeploymentService) {
	r = &CWSDeploymentService{}
	r.Options = opts
	return
}

// Creates a deployment with the specified policy sets, domain ACL, and basic-auth
// credentials.
func (r *CWSDeploymentService) New(ctx context.Context, params CWSDeploymentNewParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env CWSDeploymentNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/deployments", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the deployment's mutable configuration and resource attachments.
func (r *CWSDeploymentService) Update(ctx context.Context, deploymentID string, params CWSDeploymentUpdateParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env CWSDeploymentUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/deployments/%s", params.AccountID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of deployments in the account.
func (r *CWSDeploymentService) List(ctx context.Context, params CWSDeploymentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Deployment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/deployments", params.AccountID)
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

// Returns a paginated list of deployments in the account.
func (r *CWSDeploymentService) ListAutoPaging(ctx context.Context, params CWSDeploymentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Deployment] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the deployment and detaches its associated resources. The resources are
// not deleted.
func (r *CWSDeploymentService) Delete(ctx context.Context, deploymentID string, body CWSDeploymentDeleteParams, opts ...option.RequestOption) (res *CWSDeploymentDeleteResponse, err error) {
	var env CWSDeploymentDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/deployments/%s", body.AccountID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the deployment with the specified ID.
func (r *CWSDeploymentService) Get(ctx context.Context, deploymentID string, query CWSDeploymentGetParams, opts ...option.RequestOption) (res *Deployment, err error) {
	var env CWSDeploymentGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/deployments/%s", query.AccountID, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Deployment struct {
	// Server-assigned UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// IDs of basic-auth credentials attached to this deployment
	BasicAuthCredentialIDs []string  `json:"basic_auth_credential_ids" api:"required,nullable" format:"uuid"`
	CreatedAt              time.Time `json:"created_at" api:"required" format:"date-time"`
	// Default traffic decision when no policy set matches
	DefaultMode DeploymentDefaultMode `json:"default_mode" api:"required"`
	// Human-readable name; unique per account
	Name string `json:"name" api:"required"`
	// Action applied when an enforcing policy has a runtime Rego error
	OnErrorAction DeploymentOnErrorAction `json:"on_error_action" api:"required"`
	// IDs of policy sets attached to this deployment
	PolicySetIDs []string `json:"policy_set_ids" api:"required,nullable" format:"uuid"`
	// Whether OPA builtin failures are raised as runtime evaluation errors
	RegoRaiseError bool `json:"rego_raise_error" api:"required"`
	// Tags attached to this deployment
	Tags      map[string]string `json:"tags" api:"required"`
	UpdatedAt time.Time         `json:"updated_at" api:"required" format:"date-time"`
	// Optional free-form description.
	Description string `json:"description"`
	// ID of the attached domain ACL. Omitted when none.
	DomainACLID string         `json:"domain_acl_id" format:"uuid"`
	JSON        deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	ID                     apijson.Field
	BasicAuthCredentialIDs apijson.Field
	CreatedAt              apijson.Field
	DefaultMode            apijson.Field
	Name                   apijson.Field
	OnErrorAction          apijson.Field
	PolicySetIDs           apijson.Field
	RegoRaiseError         apijson.Field
	Tags                   apijson.Field
	UpdatedAt              apijson.Field
	Description            apijson.Field
	DomainACLID            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

// Default traffic decision when no policy set matches
type DeploymentDefaultMode string

const (
	DeploymentDefaultModeAllow DeploymentDefaultMode = "ALLOW"
	DeploymentDefaultModeBlock DeploymentDefaultMode = "BLOCK"
)

func (r DeploymentDefaultMode) IsKnown() bool {
	switch r {
	case DeploymentDefaultModeAllow, DeploymentDefaultModeBlock:
		return true
	}
	return false
}

// Action applied when an enforcing policy has a runtime Rego error
type DeploymentOnErrorAction string

const (
	DeploymentOnErrorActionBlock DeploymentOnErrorAction = "BLOCK"
	DeploymentOnErrorActionSkip  DeploymentOnErrorAction = "SKIP"
)

func (r DeploymentOnErrorAction) IsKnown() bool {
	switch r {
	case DeploymentOnErrorActionBlock, DeploymentOnErrorActionSkip:
		return true
	}
	return false
}

type CWSDeploymentDeleteResponse struct {
	ID   string                          `json:"id" api:"required" format:"uuid"`
	JSON cwsDeploymentDeleteResponseJSON `json:"-"`
}

// cwsDeploymentDeleteResponseJSON contains the JSON metadata for the struct
// [CWSDeploymentDeleteResponse]
type cwsDeploymentDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// IDs of basic-auth credentials to attach
	BasicAuthCredentialIDs param.Field[[]string] `json:"basic_auth_credential_ids" api:"required" format:"uuid"`
	// Default traffic mode: ALLOW or BLOCK
	DefaultMode param.Field[CWSDeploymentNewParamsDefaultMode] `json:"default_mode" api:"required"`
	// Name of the deployment
	Name param.Field[string] `json:"name" api:"required"`
	// IDs of policy sets to attach
	PolicySetIDs param.Field[[]string] `json:"policy_set_ids" api:"required" format:"uuid"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// ID of the domain ACL to attach (omit or null for none)
	DomainACLID param.Field[string] `json:"domain_acl_id" format:"uuid"`
	// Action for runtime Rego errors. Omit to use BLOCK.
	OnErrorAction param.Field[CWSDeploymentNewParamsOnErrorAction] `json:"on_error_action"`
	// Whether OPA builtin failures are raised as runtime evaluation errors. Omit to
	// use the default true value.
	RegoRaiseError param.Field[bool] `json:"rego_raise_error"`
	// Tags to attach. Keys are limited to 256 UTF-8 bytes and values to 1024 UTF-8
	// bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default traffic mode: ALLOW or BLOCK
type CWSDeploymentNewParamsDefaultMode string

const (
	CWSDeploymentNewParamsDefaultModeAllow CWSDeploymentNewParamsDefaultMode = "ALLOW"
	CWSDeploymentNewParamsDefaultModeBlock CWSDeploymentNewParamsDefaultMode = "BLOCK"
)

func (r CWSDeploymentNewParamsDefaultMode) IsKnown() bool {
	switch r {
	case CWSDeploymentNewParamsDefaultModeAllow, CWSDeploymentNewParamsDefaultModeBlock:
		return true
	}
	return false
}

// Action for runtime Rego errors. Omit to use BLOCK.
type CWSDeploymentNewParamsOnErrorAction string

const (
	CWSDeploymentNewParamsOnErrorActionBlock CWSDeploymentNewParamsOnErrorAction = "BLOCK"
	CWSDeploymentNewParamsOnErrorActionSkip  CWSDeploymentNewParamsOnErrorAction = "SKIP"
)

func (r CWSDeploymentNewParamsOnErrorAction) IsKnown() bool {
	switch r {
	case CWSDeploymentNewParamsOnErrorActionBlock, CWSDeploymentNewParamsOnErrorActionSkip:
		return true
	}
	return false
}

type CWSDeploymentNewResponseEnvelope struct {
	Errors   []CWSDeploymentNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDeploymentNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Deployment                                 `json:"result" api:"required"`
	Success  bool                                       `json:"success" api:"required"`
	JSON     cwsDeploymentNewResponseEnvelopeJSON       `json:"-"`
}

// cwsDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDeploymentNewResponseEnvelope]
type cwsDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSDeploymentNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDeploymentNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDeploymentNewResponseEnvelopeErrors]
type cwsDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsDeploymentNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDeploymentNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSDeploymentNewResponseEnvelopeErrorsSource]
type cwsDeploymentNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSDeploymentNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDeploymentNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSDeploymentNewResponseEnvelopeMessages]
type cwsDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsDeploymentNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDeploymentNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentNewResponseEnvelopeMessagesSource]
type cwsDeploymentNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Replacement set of basic-auth credential IDs
	BasicAuthCredentialIDs param.Field[[]string] `json:"basic_auth_credential_ids" api:"required" format:"uuid"`
	// Default traffic mode: ALLOW or BLOCK
	DefaultMode param.Field[CWSDeploymentUpdateParamsDefaultMode] `json:"default_mode" api:"required"`
	// Name of the deployment
	Name param.Field[string] `json:"name" api:"required"`
	// Replacement set of policy set IDs
	PolicySetIDs param.Field[[]string] `json:"policy_set_ids" api:"required" format:"uuid"`
	// Optional free-form description.
	Description param.Field[string] `json:"description"`
	// ID of the domain ACL to attach (omit or null to clear)
	DomainACLID param.Field[string] `json:"domain_acl_id" format:"uuid"`
	// Action for runtime Rego errors. Omit to reset to BLOCK.
	OnErrorAction param.Field[CWSDeploymentUpdateParamsOnErrorAction] `json:"on_error_action"`
	// Whether OPA builtin failures are raised as runtime evaluation errors. Omit to
	// reset to true.
	RegoRaiseError param.Field[bool] `json:"rego_raise_error"`
	// Replacement tag set. Omit or use {} to clear. Keys are limited to 256 UTF-8
	// bytes and values to 1024 UTF-8 bytes.
	Tags param.Field[map[string]string] `json:"tags"`
}

func (r CWSDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default traffic mode: ALLOW or BLOCK
type CWSDeploymentUpdateParamsDefaultMode string

const (
	CWSDeploymentUpdateParamsDefaultModeAllow CWSDeploymentUpdateParamsDefaultMode = "ALLOW"
	CWSDeploymentUpdateParamsDefaultModeBlock CWSDeploymentUpdateParamsDefaultMode = "BLOCK"
)

func (r CWSDeploymentUpdateParamsDefaultMode) IsKnown() bool {
	switch r {
	case CWSDeploymentUpdateParamsDefaultModeAllow, CWSDeploymentUpdateParamsDefaultModeBlock:
		return true
	}
	return false
}

// Action for runtime Rego errors. Omit to reset to BLOCK.
type CWSDeploymentUpdateParamsOnErrorAction string

const (
	CWSDeploymentUpdateParamsOnErrorActionBlock CWSDeploymentUpdateParamsOnErrorAction = "BLOCK"
	CWSDeploymentUpdateParamsOnErrorActionSkip  CWSDeploymentUpdateParamsOnErrorAction = "SKIP"
)

func (r CWSDeploymentUpdateParamsOnErrorAction) IsKnown() bool {
	switch r {
	case CWSDeploymentUpdateParamsOnErrorActionBlock, CWSDeploymentUpdateParamsOnErrorActionSkip:
		return true
	}
	return false
}

type CWSDeploymentUpdateResponseEnvelope struct {
	Errors   []CWSDeploymentUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDeploymentUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Deployment                                    `json:"result" api:"required"`
	Success  bool                                          `json:"success" api:"required"`
	JSON     cwsDeploymentUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsDeploymentUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSDeploymentUpdateResponseEnvelope]
type cwsDeploymentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSDeploymentUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDeploymentUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDeploymentUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDeploymentUpdateResponseEnvelopeErrors]
type cwsDeploymentUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsDeploymentUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDeploymentUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentUpdateResponseEnvelopeErrorsSource]
type cwsDeploymentUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSDeploymentUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDeploymentUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDeploymentUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSDeploymentUpdateResponseEnvelopeMessages]
type cwsDeploymentUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsDeploymentUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDeploymentUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentUpdateResponseEnvelopeMessagesSource]
type cwsDeploymentUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSDeploymentListParamsDirection] `query:"direction"`
	// Filter by deployment name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSDeploymentListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSDeploymentListParams]'s query parameters as
// `url.Values`.
func (r CWSDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSDeploymentListParamsDirection string

const (
	CWSDeploymentListParamsDirectionAsc  CWSDeploymentListParamsDirection = "asc"
	CWSDeploymentListParamsDirectionDesc CWSDeploymentListParamsDirection = "desc"
)

func (r CWSDeploymentListParamsDirection) IsKnown() bool {
	switch r {
	case CWSDeploymentListParamsDirectionAsc, CWSDeploymentListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSDeploymentListParamsOrder string

const (
	CWSDeploymentListParamsOrderName      CWSDeploymentListParamsOrder = "name"
	CWSDeploymentListParamsOrderCreatedAt CWSDeploymentListParamsOrder = "created_at"
	CWSDeploymentListParamsOrderUpdatedAt CWSDeploymentListParamsOrder = "updated_at"
)

func (r CWSDeploymentListParamsOrder) IsKnown() bool {
	switch r {
	case CWSDeploymentListParamsOrderName, CWSDeploymentListParamsOrderCreatedAt, CWSDeploymentListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSDeploymentDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSDeploymentDeleteResponseEnvelope struct {
	Errors   []CWSDeploymentDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDeploymentDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSDeploymentDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                          `json:"success" api:"required"`
	JSON     cwsDeploymentDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsDeploymentDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSDeploymentDeleteResponseEnvelope]
type cwsDeploymentDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSDeploymentDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDeploymentDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDeploymentDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDeploymentDeleteResponseEnvelopeErrors]
type cwsDeploymentDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsDeploymentDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDeploymentDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentDeleteResponseEnvelopeErrorsSource]
type cwsDeploymentDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSDeploymentDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDeploymentDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDeploymentDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSDeploymentDeleteResponseEnvelopeMessages]
type cwsDeploymentDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsDeploymentDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDeploymentDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentDeleteResponseEnvelopeMessagesSource]
type cwsDeploymentDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSDeploymentGetResponseEnvelope struct {
	Errors   []CWSDeploymentGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDeploymentGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   Deployment                                 `json:"result" api:"required"`
	Success  bool                                       `json:"success" api:"required"`
	JSON     cwsDeploymentGetResponseEnvelopeJSON       `json:"-"`
}

// cwsDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDeploymentGetResponseEnvelope]
type cwsDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSDeploymentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDeploymentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDeploymentGetResponseEnvelopeErrors]
type cwsDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsDeploymentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDeploymentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSDeploymentGetResponseEnvelopeErrorsSource]
type cwsDeploymentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSDeploymentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDeploymentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSDeploymentGetResponseEnvelopeMessages]
type cwsDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDeploymentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsDeploymentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDeploymentGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDeploymentGetResponseEnvelopeMessagesSource]
type cwsDeploymentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDeploymentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDeploymentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
