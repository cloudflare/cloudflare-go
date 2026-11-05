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

// CWSDomainACLService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSDomainACLService] method instead.
type CWSDomainACLService struct {
	Options []option.RequestOption
}

// NewCWSDomainACLService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSDomainACLService(opts ...option.RequestOption) (r *CWSDomainACLService) {
	r = &CWSDomainACLService{}
	r.Options = opts
	return
}

// Creates a domain access control list with its default action and entries.
func (r *CWSDomainACLService) New(ctx context.Context, params CWSDomainACLNewParams, opts ...option.RequestOption) (res *DomainACL, err error) {
	var env CWSDomainACLNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/domain-acls", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the domain access control list's name, default action, and entries.
func (r *CWSDomainACLService) Update(ctx context.Context, domainACLID string, params CWSDomainACLUpdateParams, opts ...option.RequestOption) (res *DomainACL, err error) {
	var env CWSDomainACLUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainACLID == "" {
		err = errors.New("missing required domain_acl_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/domain-acls/%s", params.AccountID, domainACLID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of domain access control lists in the account.
func (r *CWSDomainACLService) List(ctx context.Context, params CWSDomainACLListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[DomainACL], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/domain-acls", params.AccountID)
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

// Returns a paginated list of domain access control lists in the account.
func (r *CWSDomainACLService) ListAutoPaging(ctx context.Context, params CWSDomainACLListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[DomainACL] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the domain ACL if it is not attached to a deployment.
func (r *CWSDomainACLService) Delete(ctx context.Context, domainACLID string, body CWSDomainACLDeleteParams, opts ...option.RequestOption) (res *CWSDomainACLDeleteResponse, err error) {
	var env CWSDomainACLDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainACLID == "" {
		err = errors.New("missing required domain_acl_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/domain-acls/%s", body.AccountID, domainACLID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the domain access control list with the specified ID.
func (r *CWSDomainACLService) Get(ctx context.Context, domainACLID string, query CWSDomainACLGetParams, opts ...option.RequestOption) (res *DomainACL, err error) {
	var env CWSDomainACLGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if domainACLID == "" {
		err = errors.New("missing required domain_acl_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/domain-acls/%s", query.AccountID, domainACLID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DomainACL struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Action applied when no entry matches the request domain
	DefaultAction DomainACLDefaultAction `json:"default_action" api:"required"`
	// Per-domain action overrides. Each domain appears at most once.
	Entries []DomainACLEntry `json:"entries" api:"required,nullable"`
	// Human-readable name; unique per account
	Name      string        `json:"name" api:"required"`
	UpdatedAt time.Time     `json:"updated_at" api:"required" format:"date-time"`
	JSON      domainACLJSON `json:"-"`
}

// domainACLJSON contains the JSON metadata for the struct [DomainACL]
type domainACLJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	DefaultAction apijson.Field
	Entries       apijson.Field
	Name          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DomainACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainACLJSON) RawJSON() string {
	return r.raw
}

// Action applied when no entry matches the request domain
type DomainACLDefaultAction string

const (
	DomainACLDefaultActionAllow DomainACLDefaultAction = "ALLOW"
	DomainACLDefaultActionBlock DomainACLDefaultAction = "BLOCK"
)

func (r DomainACLDefaultAction) IsKnown() bool {
	switch r {
	case DomainACLDefaultActionAllow, DomainACLDefaultActionBlock:
		return true
	}
	return false
}

type DomainACLEntry struct {
	// Per-domain decision overriding the parent ACL's default_action
	Action DomainACLEntriesAction `json:"action" api:"required"`
	// Hostname or domain pattern, capped at 255 UTF-8 bytes
	Domain string             `json:"domain" api:"required"`
	JSON   domainACLEntryJSON `json:"-"`
}

// domainACLEntryJSON contains the JSON metadata for the struct [DomainACLEntry]
type domainACLEntryJSON struct {
	Action      apijson.Field
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainACLEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainACLEntryJSON) RawJSON() string {
	return r.raw
}

// Per-domain decision overriding the parent ACL's default_action
type DomainACLEntriesAction string

const (
	DomainACLEntriesActionAllow DomainACLEntriesAction = "ALLOW"
	DomainACLEntriesActionBlock DomainACLEntriesAction = "BLOCK"
)

func (r DomainACLEntriesAction) IsKnown() bool {
	switch r {
	case DomainACLEntriesActionAllow, DomainACLEntriesActionBlock:
		return true
	}
	return false
}

type CWSDomainACLDeleteResponse struct {
	ID   string                         `json:"id" api:"required" format:"uuid"`
	JSON cwsDomainACLDeleteResponseJSON `json:"-"`
}

// cwsDomainACLDeleteResponseJSON contains the JSON metadata for the struct
// [CWSDomainACLDeleteResponse]
type cwsDomainACLDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Default action when no entry matches
	DefaultAction param.Field[CWSDomainACLNewParamsDefaultAction] `json:"default_action" api:"required"`
	// Domain entries for this ACL. Each domain must appear at most once.
	Entries param.Field[[]CWSDomainACLNewParamsEntry] `json:"entries" api:"required"`
	// Name of the domain ACL
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CWSDomainACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default action when no entry matches
type CWSDomainACLNewParamsDefaultAction string

const (
	CWSDomainACLNewParamsDefaultActionAllow CWSDomainACLNewParamsDefaultAction = "ALLOW"
	CWSDomainACLNewParamsDefaultActionBlock CWSDomainACLNewParamsDefaultAction = "BLOCK"
)

func (r CWSDomainACLNewParamsDefaultAction) IsKnown() bool {
	switch r {
	case CWSDomainACLNewParamsDefaultActionAllow, CWSDomainACLNewParamsDefaultActionBlock:
		return true
	}
	return false
}

type CWSDomainACLNewParamsEntry struct {
	// Per-domain decision overriding the parent ACL's default_action
	Action param.Field[CWSDomainACLNewParamsEntriesAction] `json:"action" api:"required"`
	// Hostname or domain pattern, capped at 255 UTF-8 bytes
	Domain param.Field[string] `json:"domain" api:"required"`
}

func (r CWSDomainACLNewParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-domain decision overriding the parent ACL's default_action
type CWSDomainACLNewParamsEntriesAction string

const (
	CWSDomainACLNewParamsEntriesActionAllow CWSDomainACLNewParamsEntriesAction = "ALLOW"
	CWSDomainACLNewParamsEntriesActionBlock CWSDomainACLNewParamsEntriesAction = "BLOCK"
)

func (r CWSDomainACLNewParamsEntriesAction) IsKnown() bool {
	switch r {
	case CWSDomainACLNewParamsEntriesActionAllow, CWSDomainACLNewParamsEntriesActionBlock:
		return true
	}
	return false
}

type CWSDomainACLNewResponseEnvelope struct {
	Errors   []CWSDomainACLNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDomainACLNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   DomainACL                                 `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsDomainACLNewResponseEnvelopeJSON       `json:"-"`
}

// cwsDomainACLNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDomainACLNewResponseEnvelope]
type cwsDomainACLNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSDomainACLNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDomainACLNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDomainACLNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDomainACLNewResponseEnvelopeErrors]
type cwsDomainACLNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsDomainACLNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDomainACLNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSDomainACLNewResponseEnvelopeErrorsSource]
type cwsDomainACLNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSDomainACLNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDomainACLNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDomainACLNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSDomainACLNewResponseEnvelopeMessages]
type cwsDomainACLNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsDomainACLNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDomainACLNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSDomainACLNewResponseEnvelopeMessagesSource]
type cwsDomainACLNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Default action when no entry matches
	DefaultAction param.Field[CWSDomainACLUpdateParamsDefaultAction] `json:"default_action" api:"required"`
	// Replacement set of domain entries. Each domain must appear at most once.
	Entries param.Field[[]CWSDomainACLUpdateParamsEntry] `json:"entries" api:"required"`
	// Name of the domain ACL
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CWSDomainACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default action when no entry matches
type CWSDomainACLUpdateParamsDefaultAction string

const (
	CWSDomainACLUpdateParamsDefaultActionAllow CWSDomainACLUpdateParamsDefaultAction = "ALLOW"
	CWSDomainACLUpdateParamsDefaultActionBlock CWSDomainACLUpdateParamsDefaultAction = "BLOCK"
)

func (r CWSDomainACLUpdateParamsDefaultAction) IsKnown() bool {
	switch r {
	case CWSDomainACLUpdateParamsDefaultActionAllow, CWSDomainACLUpdateParamsDefaultActionBlock:
		return true
	}
	return false
}

type CWSDomainACLUpdateParamsEntry struct {
	// Per-domain decision overriding the parent ACL's default_action
	Action param.Field[CWSDomainACLUpdateParamsEntriesAction] `json:"action" api:"required"`
	// Hostname or domain pattern, capped at 255 UTF-8 bytes
	Domain param.Field[string] `json:"domain" api:"required"`
}

func (r CWSDomainACLUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-domain decision overriding the parent ACL's default_action
type CWSDomainACLUpdateParamsEntriesAction string

const (
	CWSDomainACLUpdateParamsEntriesActionAllow CWSDomainACLUpdateParamsEntriesAction = "ALLOW"
	CWSDomainACLUpdateParamsEntriesActionBlock CWSDomainACLUpdateParamsEntriesAction = "BLOCK"
)

func (r CWSDomainACLUpdateParamsEntriesAction) IsKnown() bool {
	switch r {
	case CWSDomainACLUpdateParamsEntriesActionAllow, CWSDomainACLUpdateParamsEntriesActionBlock:
		return true
	}
	return false
}

type CWSDomainACLUpdateResponseEnvelope struct {
	Errors   []CWSDomainACLUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDomainACLUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   DomainACL                                    `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsDomainACLUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsDomainACLUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDomainACLUpdateResponseEnvelope]
type cwsDomainACLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSDomainACLUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDomainACLUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDomainACLUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDomainACLUpdateResponseEnvelopeErrors]
type cwsDomainACLUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsDomainACLUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDomainACLUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSDomainACLUpdateResponseEnvelopeErrorsSource]
type cwsDomainACLUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSDomainACLUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDomainACLUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDomainACLUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSDomainACLUpdateResponseEnvelopeMessages]
type cwsDomainACLUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsDomainACLUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDomainACLUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDomainACLUpdateResponseEnvelopeMessagesSource]
type cwsDomainACLUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Filter to domain ACLs attached to this deployment
	DeploymentID param.Field[string] `query:"deployment_id" format:"uuid"`
	// Sort direction (default: asc)
	Direction param.Field[CWSDomainACLListParamsDirection] `query:"direction"`
	// Filter by domain ACL name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSDomainACLListParamsOrder] `query:"order"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSDomainACLListParams]'s query parameters as `url.Values`.
func (r CWSDomainACLListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSDomainACLListParamsDirection string

const (
	CWSDomainACLListParamsDirectionAsc  CWSDomainACLListParamsDirection = "asc"
	CWSDomainACLListParamsDirectionDesc CWSDomainACLListParamsDirection = "desc"
)

func (r CWSDomainACLListParamsDirection) IsKnown() bool {
	switch r {
	case CWSDomainACLListParamsDirectionAsc, CWSDomainACLListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSDomainACLListParamsOrder string

const (
	CWSDomainACLListParamsOrderName      CWSDomainACLListParamsOrder = "name"
	CWSDomainACLListParamsOrderCreatedAt CWSDomainACLListParamsOrder = "created_at"
	CWSDomainACLListParamsOrderUpdatedAt CWSDomainACLListParamsOrder = "updated_at"
)

func (r CWSDomainACLListParamsOrder) IsKnown() bool {
	switch r {
	case CWSDomainACLListParamsOrderName, CWSDomainACLListParamsOrderCreatedAt, CWSDomainACLListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSDomainACLDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSDomainACLDeleteResponseEnvelope struct {
	Errors   []CWSDomainACLDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDomainACLDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSDomainACLDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     cwsDomainACLDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsDomainACLDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDomainACLDeleteResponseEnvelope]
type cwsDomainACLDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSDomainACLDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDomainACLDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDomainACLDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDomainACLDeleteResponseEnvelopeErrors]
type cwsDomainACLDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsDomainACLDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDomainACLDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSDomainACLDeleteResponseEnvelopeErrorsSource]
type cwsDomainACLDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	Source  CWSDomainACLDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDomainACLDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDomainACLDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSDomainACLDeleteResponseEnvelopeMessages]
type cwsDomainACLDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    cwsDomainACLDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDomainACLDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSDomainACLDeleteResponseEnvelopeMessagesSource]
type cwsDomainACLDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSDomainACLGetResponseEnvelope struct {
	Errors   []CWSDomainACLGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSDomainACLGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   DomainACL                                 `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     cwsDomainACLGetResponseEnvelopeJSON       `json:"-"`
}

// cwsDomainACLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSDomainACLGetResponseEnvelope]
type cwsDomainACLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	Source  CWSDomainACLGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsDomainACLGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsDomainACLGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSDomainACLGetResponseEnvelopeErrors]
type cwsDomainACLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    cwsDomainACLGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsDomainACLGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSDomainACLGetResponseEnvelopeErrorsSource]
type cwsDomainACLGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	Source  CWSDomainACLGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsDomainACLGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsDomainACLGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSDomainACLGetResponseEnvelopeMessages]
type cwsDomainACLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSDomainACLGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    cwsDomainACLGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsDomainACLGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [CWSDomainACLGetResponseEnvelopeMessagesSource]
type cwsDomainACLGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSDomainACLGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsDomainACLGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
