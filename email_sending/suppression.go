// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending

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

// SuppressionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuppressionService] method instead.
type SuppressionService struct {
	Options []option.RequestOption
}

// NewSuppressionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuppressionService(opts ...option.RequestOption) (r *SuppressionService) {
	r = &SuppressionService{}
	r.Options = opts
	return
}

// Creates a suppression for every sending domain of the account (default) or for
// one sending domain (`scope.type = sending_domain`). Creating an existing active
// suppression returns its identifier. If a mutable legacy zone-linked account row
// already exists, it is promoted without changing its identifier.
func (r *SuppressionService) New(ctx context.Context, params SuppressionNewParams, opts ...option.RequestOption) (res *SuppressionNewResponse, err error) {
	var env SuppressionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists every active Email Sending suppression owned by the account:
// sending-domain suppressions first, then account-wide suppressions (including
// legacy rows with internal zone memberships). Each group is newest first.
func (r *SuppressionService) List(ctx context.Context, params SuppressionListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[SuppressionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions", params.AccountID)
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

// Lists every active Email Sending suppression owned by the account:
// sending-domain suppressions first, then account-wide suppressions (including
// legacy rows with internal zone memberships). Each group is newest first.
func (r *SuppressionService) ListAutoPaging(ctx context.Context, params SuppressionListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[SuppressionListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the suppression, its note, and every legacy internal zone membership,
// allowing future delivery attempts to the address.
func (r *SuppressionService) Delete(ctx context.Context, suppressionID string, body SuppressionDeleteParams, opts ...option.RequestOption) (res *SuppressionDeleteResponse, err error) {
	var env SuppressionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", body.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates expiry or advisory note fields without changing legacy internal zone
// memberships. Scope cannot be changed.
func (r *SuppressionService) Edit(ctx context.Context, suppressionID string, params SuppressionEditParams, opts ...option.RequestOption) (res *SuppressionEditResponse, err error) {
	var env SuppressionEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", params.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets an Email Sending suppression owned by the account.
func (r *SuppressionService) Get(ctx context.Context, suppressionID string, query SuppressionGetParams, opts ...option.RequestOption) (res *SuppressionGetResponse, err error) {
	var env SuppressionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if suppressionID == "" {
		err = errors.New("missing required suppression_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/%s", query.AccountID, suppressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Imports up to 1,000 Email Sending suppressions in one request. Each item applies
// to every sending domain of the account (default) or to one sending domain.
func (r *SuppressionService) Import(ctx context.Context, params SuppressionImportParams, opts ...option.RequestOption) (res *SuppressionImportResponse, err error) {
	var env SuppressionImportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email/sending/suppressions/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SuppressionNewResponse struct {
	// The suppression's identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionNewResponseScope `json:"scope"`
	JSON  suppressionNewResponseJSON  `json:"-"`
}

// suppressionNewResponseJSON contains the JSON metadata for the struct
// [SuppressionNewResponse]
type suppressionNewResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionNewResponseScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionNewResponseScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                          `json:"value"`
	JSON  suppressionNewResponseScopeJSON `json:"-"`
	union SuppressionNewResponseScopeUnion
}

// suppressionNewResponseScopeJSON contains the JSON metadata for the struct
// [SuppressionNewResponseScope]
type suppressionNewResponseScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionNewResponseScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionNewResponseScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionNewResponseScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionNewResponseScopeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SuppressionNewResponseScopeType],
// [SuppressionNewResponseScopeObject].
func (r SuppressionNewResponseScope) AsUnion() SuppressionNewResponseScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionNewResponseScopeType] or
// [SuppressionNewResponseScopeObject].
type SuppressionNewResponseScopeUnion interface {
	implementsSuppressionNewResponseScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionNewResponseScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionNewResponseScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionNewResponseScopeObject{}),
		},
	)
}

type SuppressionNewResponseScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionNewResponseScopeTypeType `json:"type" api:"required"`
	JSON suppressionNewResponseScopeTypeJSON `json:"-"`
}

// suppressionNewResponseScopeTypeJSON contains the JSON metadata for the struct
// [SuppressionNewResponseScopeType]
type suppressionNewResponseScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponseScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionNewResponseScopeType) implementsSuppressionNewResponseScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionNewResponseScopeTypeType string

const (
	SuppressionNewResponseScopeTypeTypeAccount SuppressionNewResponseScopeTypeType = "account"
)

func (r SuppressionNewResponseScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionNewResponseScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionNewResponseScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionNewResponseScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                `json:"value" api:"required"`
	JSON  suppressionNewResponseScopeObjectJSON `json:"-"`
}

// suppressionNewResponseScopeObjectJSON contains the JSON metadata for the struct
// [SuppressionNewResponseScopeObject]
type suppressionNewResponseScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionNewResponseScopeObject) implementsSuppressionNewResponseScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionNewResponseScopeObjectType string

const (
	SuppressionNewResponseScopeObjectTypeSendingDomain SuppressionNewResponseScopeObjectType = "sending_domain"
)

func (r SuppressionNewResponseScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionNewResponseScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionListResponse struct {
	// Unique identifier for this suppression.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the suppression was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The suppressed email address.
	Email string `json:"email" api:"required" format:"email"`
	// When the suppression expires. Null for a permanent suppression.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool `json:"read_only" api:"required"`
	// Why the address is suppressed: `manual`, `complaint`, `hard_bounce`,
	// `soft_bounce`, or `policy`.
	Reason string `json:"reason" api:"required"`
	// Advisory note for this suppression, if any.
	Note string `json:"note" api:"nullable"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionListResponseScope `json:"scope"`
	JSON  suppressionListResponseJSON  `json:"-"`
}

// suppressionListResponseJSON contains the JSON metadata for the struct
// [SuppressionListResponse]
type suppressionListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionListResponseJSON) RawJSON() string {
	return r.raw
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionListResponseScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionListResponseScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                           `json:"value"`
	JSON  suppressionListResponseScopeJSON `json:"-"`
	union SuppressionListResponseScopeUnion
}

// suppressionListResponseScopeJSON contains the JSON metadata for the struct
// [SuppressionListResponseScope]
type suppressionListResponseScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionListResponseScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionListResponseScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionListResponseScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionListResponseScopeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SuppressionListResponseScopeType],
// [SuppressionListResponseScopeObject].
func (r SuppressionListResponseScope) AsUnion() SuppressionListResponseScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionListResponseScopeType] or
// [SuppressionListResponseScopeObject].
type SuppressionListResponseScopeUnion interface {
	implementsSuppressionListResponseScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionListResponseScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionListResponseScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionListResponseScopeObject{}),
		},
	)
}

type SuppressionListResponseScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionListResponseScopeTypeType `json:"type" api:"required"`
	JSON suppressionListResponseScopeTypeJSON `json:"-"`
}

// suppressionListResponseScopeTypeJSON contains the JSON metadata for the struct
// [SuppressionListResponseScopeType]
type suppressionListResponseScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionListResponseScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionListResponseScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionListResponseScopeType) implementsSuppressionListResponseScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionListResponseScopeTypeType string

const (
	SuppressionListResponseScopeTypeTypeAccount SuppressionListResponseScopeTypeType = "account"
)

func (r SuppressionListResponseScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionListResponseScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionListResponseScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionListResponseScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                 `json:"value" api:"required"`
	JSON  suppressionListResponseScopeObjectJSON `json:"-"`
}

// suppressionListResponseScopeObjectJSON contains the JSON metadata for the struct
// [SuppressionListResponseScopeObject]
type suppressionListResponseScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionListResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionListResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionListResponseScopeObject) implementsSuppressionListResponseScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionListResponseScopeObjectType string

const (
	SuppressionListResponseScopeObjectTypeSendingDomain SuppressionListResponseScopeObjectType = "sending_domain"
)

func (r SuppressionListResponseScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionListResponseScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionDeleteResponse struct {
	// The suppression's identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionDeleteResponseScope `json:"scope"`
	JSON  suppressionDeleteResponseJSON  `json:"-"`
}

// suppressionDeleteResponseJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponse]
type suppressionDeleteResponseJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionDeleteResponseScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionDeleteResponseScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                             `json:"value"`
	JSON  suppressionDeleteResponseScopeJSON `json:"-"`
	union SuppressionDeleteResponseScopeUnion
}

// suppressionDeleteResponseScopeJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponseScope]
type suppressionDeleteResponseScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionDeleteResponseScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionDeleteResponseScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionDeleteResponseScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionDeleteResponseScopeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SuppressionDeleteResponseScopeType],
// [SuppressionDeleteResponseScopeObject].
func (r SuppressionDeleteResponseScope) AsUnion() SuppressionDeleteResponseScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionDeleteResponseScopeType] or
// [SuppressionDeleteResponseScopeObject].
type SuppressionDeleteResponseScopeUnion interface {
	implementsSuppressionDeleteResponseScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionDeleteResponseScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionDeleteResponseScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionDeleteResponseScopeObject{}),
		},
	)
}

type SuppressionDeleteResponseScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionDeleteResponseScopeTypeType `json:"type" api:"required"`
	JSON suppressionDeleteResponseScopeTypeJSON `json:"-"`
}

// suppressionDeleteResponseScopeTypeJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponseScopeType]
type suppressionDeleteResponseScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponseScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionDeleteResponseScopeType) implementsSuppressionDeleteResponseScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionDeleteResponseScopeTypeType string

const (
	SuppressionDeleteResponseScopeTypeTypeAccount SuppressionDeleteResponseScopeTypeType = "account"
)

func (r SuppressionDeleteResponseScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionDeleteResponseScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionDeleteResponseScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionDeleteResponseScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                   `json:"value" api:"required"`
	JSON  suppressionDeleteResponseScopeObjectJSON `json:"-"`
}

// suppressionDeleteResponseScopeObjectJSON contains the JSON metadata for the
// struct [SuppressionDeleteResponseScopeObject]
type suppressionDeleteResponseScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionDeleteResponseScopeObject) implementsSuppressionDeleteResponseScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionDeleteResponseScopeObjectType string

const (
	SuppressionDeleteResponseScopeObjectTypeSendingDomain SuppressionDeleteResponseScopeObjectType = "sending_domain"
)

func (r SuppressionDeleteResponseScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionDeleteResponseScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionEditResponse struct {
	// Unique identifier for this suppression.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the suppression was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The suppressed email address.
	Email string `json:"email" api:"required" format:"email"`
	// When the suppression expires. Null for a permanent suppression.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool `json:"read_only" api:"required"`
	// Why the address is suppressed: `manual`, `complaint`, `hard_bounce`,
	// `soft_bounce`, or `policy`.
	Reason string `json:"reason" api:"required"`
	// Advisory note for this suppression, if any.
	Note string `json:"note" api:"nullable"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionEditResponseScope `json:"scope"`
	JSON  suppressionEditResponseJSON  `json:"-"`
}

// suppressionEditResponseJSON contains the JSON metadata for the struct
// [SuppressionEditResponse]
type suppressionEditResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseJSON) RawJSON() string {
	return r.raw
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionEditResponseScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionEditResponseScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                           `json:"value"`
	JSON  suppressionEditResponseScopeJSON `json:"-"`
	union SuppressionEditResponseScopeUnion
}

// suppressionEditResponseScopeJSON contains the JSON metadata for the struct
// [SuppressionEditResponseScope]
type suppressionEditResponseScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionEditResponseScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionEditResponseScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionEditResponseScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionEditResponseScopeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SuppressionEditResponseScopeType],
// [SuppressionEditResponseScopeObject].
func (r SuppressionEditResponseScope) AsUnion() SuppressionEditResponseScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionEditResponseScopeType] or
// [SuppressionEditResponseScopeObject].
type SuppressionEditResponseScopeUnion interface {
	implementsSuppressionEditResponseScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionEditResponseScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionEditResponseScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionEditResponseScopeObject{}),
		},
	)
}

type SuppressionEditResponseScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionEditResponseScopeTypeType `json:"type" api:"required"`
	JSON suppressionEditResponseScopeTypeJSON `json:"-"`
}

// suppressionEditResponseScopeTypeJSON contains the JSON metadata for the struct
// [SuppressionEditResponseScopeType]
type suppressionEditResponseScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponseScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionEditResponseScopeType) implementsSuppressionEditResponseScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionEditResponseScopeTypeType string

const (
	SuppressionEditResponseScopeTypeTypeAccount SuppressionEditResponseScopeTypeType = "account"
)

func (r SuppressionEditResponseScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionEditResponseScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionEditResponseScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionEditResponseScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                 `json:"value" api:"required"`
	JSON  suppressionEditResponseScopeObjectJSON `json:"-"`
}

// suppressionEditResponseScopeObjectJSON contains the JSON metadata for the struct
// [SuppressionEditResponseScopeObject]
type suppressionEditResponseScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionEditResponseScopeObject) implementsSuppressionEditResponseScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionEditResponseScopeObjectType string

const (
	SuppressionEditResponseScopeObjectTypeSendingDomain SuppressionEditResponseScopeObjectType = "sending_domain"
)

func (r SuppressionEditResponseScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionEditResponseScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionGetResponse struct {
	// Unique identifier for this suppression.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the suppression was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The suppressed email address.
	Email string `json:"email" api:"required" format:"email"`
	// When the suppression expires. Null for a permanent suppression.
	ExpiresAt time.Time `json:"expires_at" api:"required,nullable" format:"date-time"`
	// Whether clients may mutate this suppression. This is determined by the server
	// and must not be inferred from `reason`.
	ReadOnly bool `json:"read_only" api:"required"`
	// Why the address is suppressed: `manual`, `complaint`, `hard_bounce`,
	// `soft_bounce`, or `policy`.
	Reason string `json:"reason" api:"required"`
	// Advisory note for this suppression, if any.
	Note string `json:"note" api:"nullable"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionGetResponseScope `json:"scope"`
	JSON  suppressionGetResponseJSON  `json:"-"`
}

// suppressionGetResponseJSON contains the JSON metadata for the struct
// [SuppressionGetResponse]
type suppressionGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	ExpiresAt   apijson.Field
	ReadOnly    apijson.Field
	Reason      apijson.Field
	Note        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionGetResponseScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionGetResponseScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                          `json:"value"`
	JSON  suppressionGetResponseScopeJSON `json:"-"`
	union SuppressionGetResponseScopeUnion
}

// suppressionGetResponseScopeJSON contains the JSON metadata for the struct
// [SuppressionGetResponseScope]
type suppressionGetResponseScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionGetResponseScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionGetResponseScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionGetResponseScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionGetResponseScopeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SuppressionGetResponseScopeType],
// [SuppressionGetResponseScopeObject].
func (r SuppressionGetResponseScope) AsUnion() SuppressionGetResponseScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionGetResponseScopeType] or
// [SuppressionGetResponseScopeObject].
type SuppressionGetResponseScopeUnion interface {
	implementsSuppressionGetResponseScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionGetResponseScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionGetResponseScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionGetResponseScopeObject{}),
		},
	)
}

type SuppressionGetResponseScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionGetResponseScopeTypeType `json:"type" api:"required"`
	JSON suppressionGetResponseScopeTypeJSON `json:"-"`
}

// suppressionGetResponseScopeTypeJSON contains the JSON metadata for the struct
// [SuppressionGetResponseScopeType]
type suppressionGetResponseScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponseScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionGetResponseScopeType) implementsSuppressionGetResponseScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionGetResponseScopeTypeType string

const (
	SuppressionGetResponseScopeTypeTypeAccount SuppressionGetResponseScopeTypeType = "account"
)

func (r SuppressionGetResponseScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionGetResponseScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionGetResponseScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionGetResponseScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                `json:"value" api:"required"`
	JSON  suppressionGetResponseScopeObjectJSON `json:"-"`
}

// suppressionGetResponseScopeObjectJSON contains the JSON metadata for the struct
// [SuppressionGetResponseScopeObject]
type suppressionGetResponseScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponseScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionGetResponseScopeObject) implementsSuppressionGetResponseScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionGetResponseScopeObjectType string

const (
	SuppressionGetResponseScopeObjectTypeSendingDomain SuppressionGetResponseScopeObjectType = "sending_domain"
)

func (r SuppressionGetResponseScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionGetResponseScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionImportResponse struct {
	// Number of items dropped because their email address and scope repeated an
	// earlier item in this request. Counted once and excluded from `items`.
	Deduplicated int64 `json:"deduplicated" api:"required"`
	// Number of items that failed to import due to an unexpected error.
	Errors int64 `json:"errors" api:"required"`
	// Number of items with an invalid email address or sending domain.
	Invalid int64 `json:"invalid" api:"required"`
	// Per-item results, in the same order as the request body.
	Items []SuppressionImportResponseItem `json:"items" api:"required"`
	// Number of items successfully created or promoted.
	Processed int64 `json:"processed" api:"required"`
	// Number of items skipped because the existing suppression is not customer-managed
	// (for example, a read-only policy suppression).
	Skipped int64 `json:"skipped" api:"required"`
	// Total number of items in the request body, including duplicates.
	Total int64                         `json:"total" api:"required"`
	JSON  suppressionImportResponseJSON `json:"-"`
}

// suppressionImportResponseJSON contains the JSON metadata for the struct
// [SuppressionImportResponse]
type suppressionImportResponseJSON struct {
	Deduplicated apijson.Field
	Errors       apijson.Field
	Invalid      apijson.Field
	Items        apijson.Field
	Processed    apijson.Field
	Skipped      apijson.Field
	Total        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SuppressionImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportResponseItem struct {
	// Zero-based index of this item in the request body.
	Index int64 `json:"index" api:"required"`
	// Outcome for this item.
	Status SuppressionImportResponseItemsStatus `json:"status" api:"required"`
	// The created or promoted suppression's identifier. Present when `status` is
	// `processed`.
	ID string `json:"id" format:"uuid"`
	// The submitted email address for this item.
	Email string `json:"email" format:"email"`
	// Human-readable error message. Present when `status` is `invalid`, `error`, or
	// `skipped`.
	Error string `json:"error"`
	// Where the suppression applies: `account` for every sending domain of the
	// account, or `sending_domain` for one envelope MAIL FROM domain.
	Scope SuppressionImportResponseItemsScope `json:"scope"`
	JSON  suppressionImportResponseItemJSON   `json:"-"`
}

// suppressionImportResponseItemJSON contains the JSON metadata for the struct
// [SuppressionImportResponseItem]
type suppressionImportResponseItemJSON struct {
	Index       apijson.Field
	Status      apijson.Field
	ID          apijson.Field
	Email       apijson.Field
	Error       apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseItemJSON) RawJSON() string {
	return r.raw
}

// Outcome for this item.
type SuppressionImportResponseItemsStatus string

const (
	SuppressionImportResponseItemsStatusProcessed SuppressionImportResponseItemsStatus = "processed"
	SuppressionImportResponseItemsStatusInvalid   SuppressionImportResponseItemsStatus = "invalid"
	SuppressionImportResponseItemsStatusError     SuppressionImportResponseItemsStatus = "error"
	SuppressionImportResponseItemsStatusSkipped   SuppressionImportResponseItemsStatus = "skipped"
)

func (r SuppressionImportResponseItemsStatus) IsKnown() bool {
	switch r {
	case SuppressionImportResponseItemsStatusProcessed, SuppressionImportResponseItemsStatusInvalid, SuppressionImportResponseItemsStatusError, SuppressionImportResponseItemsStatusSkipped:
		return true
	}
	return false
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
type SuppressionImportResponseItemsScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionImportResponseItemsScopeType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                  `json:"value"`
	JSON  suppressionImportResponseItemsScopeJSON `json:"-"`
	union SuppressionImportResponseItemsScopeUnion
}

// suppressionImportResponseItemsScopeJSON contains the JSON metadata for the
// struct [SuppressionImportResponseItemsScope]
type suppressionImportResponseItemsScopeJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r suppressionImportResponseItemsScopeJSON) RawJSON() string {
	return r.raw
}

func (r *SuppressionImportResponseItemsScope) UnmarshalJSON(data []byte) (err error) {
	*r = SuppressionImportResponseItemsScope{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuppressionImportResponseItemsScopeUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SuppressionImportResponseItemsScopeType],
// [SuppressionImportResponseItemsScopeObject].
func (r SuppressionImportResponseItemsScope) AsUnion() SuppressionImportResponseItemsScopeUnion {
	return r.union
}

// Where the suppression applies: `account` for every sending domain of the
// account, or `sending_domain` for one envelope MAIL FROM domain.
//
// Union satisfied by [SuppressionImportResponseItemsScopeType] or
// [SuppressionImportResponseItemsScopeObject].
type SuppressionImportResponseItemsScopeUnion interface {
	implementsSuppressionImportResponseItemsScope()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuppressionImportResponseItemsScopeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionImportResponseItemsScopeType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuppressionImportResponseItemsScopeObject{}),
		},
	)
}

type SuppressionImportResponseItemsScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type SuppressionImportResponseItemsScopeTypeType `json:"type" api:"required"`
	JSON suppressionImportResponseItemsScopeTypeJSON `json:"-"`
}

// suppressionImportResponseItemsScopeTypeJSON contains the JSON metadata for the
// struct [SuppressionImportResponseItemsScopeType]
type suppressionImportResponseItemsScopeTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseItemsScopeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseItemsScopeTypeJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionImportResponseItemsScopeType) implementsSuppressionImportResponseItemsScope() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionImportResponseItemsScopeTypeType string

const (
	SuppressionImportResponseItemsScopeTypeTypeAccount SuppressionImportResponseItemsScopeTypeType = "account"
)

func (r SuppressionImportResponseItemsScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionImportResponseItemsScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionImportResponseItemsScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type SuppressionImportResponseItemsScopeObjectType `json:"type" api:"required"`
	// The sending domain: the domain part of the envelope MAIL FROM, lowercase,
	// without a trailing dot.
	Value string                                        `json:"value" api:"required"`
	JSON  suppressionImportResponseItemsScopeObjectJSON `json:"-"`
}

// suppressionImportResponseItemsScopeObjectJSON contains the JSON metadata for the
// struct [SuppressionImportResponseItemsScopeObject]
type suppressionImportResponseItemsScopeObjectJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseItemsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseItemsScopeObjectJSON) RawJSON() string {
	return r.raw
}

func (r SuppressionImportResponseItemsScopeObject) implementsSuppressionImportResponseItemsScope() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionImportResponseItemsScopeObjectType string

const (
	SuppressionImportResponseItemsScopeObjectTypeSendingDomain SuppressionImportResponseItemsScopeObjectType = "sending_domain"
)

func (r SuppressionImportResponseItemsScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionImportResponseItemsScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionNewParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The email address to suppress.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
	// Expiration timestamp for the suppression. Omit or set to null for a permanent
	// suppression that never expires.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Advisory note for this suppression. Not enforced or validated beyond length.
	Note param.Field[string] `json:"note"`
	// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
	// the recipient for every sending domain of the account.
	Scope param.Field[SuppressionNewParamsScopeUnion] `json:"scope"`
}

func (r SuppressionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
// the recipient for every sending domain of the account.
type SuppressionNewParamsScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type param.Field[SuppressionNewParamsScopeType] `json:"type" api:"required"`
	// The sending domain to suppress for: the domain part of the envelope MAIL FROM.
	// It is lowercased and trailing dots are removed. Internationalized domains must
	// use the ASCII (punycode) form. Ownership is not checked; a domain the account
	// does not send from never matches.
	Value param.Field[string] `json:"value"`
}

func (r SuppressionNewParamsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionNewParamsScope) implementsSuppressionNewParamsScopeUnion() {}

// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
// the recipient for every sending domain of the account.
//
// Satisfied by [email_sending.SuppressionNewParamsScopeType],
// [email_sending.SuppressionNewParamsScopeObject], [SuppressionNewParamsScope].
type SuppressionNewParamsScopeUnion interface {
	implementsSuppressionNewParamsScopeUnion()
}

type SuppressionNewParamsScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type param.Field[SuppressionNewParamsScopeTypeType] `json:"type" api:"required"`
}

func (r SuppressionNewParamsScopeType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionNewParamsScopeType) implementsSuppressionNewParamsScopeUnion() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionNewParamsScopeTypeType string

const (
	SuppressionNewParamsScopeTypeTypeAccount SuppressionNewParamsScopeTypeType = "account"
)

func (r SuppressionNewParamsScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionNewParamsScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionNewParamsScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type param.Field[SuppressionNewParamsScopeObjectType] `json:"type" api:"required"`
	// The sending domain to suppress for: the domain part of the envelope MAIL FROM.
	// It is lowercased and trailing dots are removed. Internationalized domains must
	// use the ASCII (punycode) form. Ownership is not checked; a domain the account
	// does not send from never matches.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r SuppressionNewParamsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionNewParamsScopeObject) implementsSuppressionNewParamsScopeUnion() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionNewParamsScopeObjectType string

const (
	SuppressionNewParamsScopeObjectTypeSendingDomain SuppressionNewParamsScopeObjectType = "sending_domain"
)

func (r SuppressionNewParamsScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionNewParamsScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionNewResponseEnvelope struct {
	Errors   []interface{}                      `json:"errors" api:"required"`
	Messages []interface{}                      `json:"messages" api:"required"`
	Result   SuppressionNewResponse             `json:"result" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	JSON     suppressionNewResponseEnvelopeJSON `json:"-"`
}

// suppressionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionNewResponseEnvelope]
type suppressionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionListParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque pagination cursor returned as `result_info.next_cursor`. It carries the
	// filters that produced it.
	Cursor param.Field[string] `query:"cursor"`
	// Exact email-address filter.
	Email param.Field[string] `query:"email" format:"email"`
	// Maximum number of suppressions to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter to suppressions with this reason.
	Reason param.Field[SuppressionListParamsReason] `query:"reason"`
	// Filter by scope: `account` returns only account-wide suppressions,
	// `sending_domain` only sending-domain suppressions. Omit to list both,
	// sending-domain suppressions first.
	ScopeType param.Field[SuppressionListParamsScopeType] `query:"scope_type"`
	// Exact sending-domain filter. Requires `scope_type=sending_domain`.
	ScopeValue param.Field[string] `query:"scope_value"`
	// A complete address is an exact match; a value ending in `@` matches that
	// username across every domain. Prefix searches may return short intermediate
	// pages while the bounded account scan advances.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SuppressionListParams]'s query parameters as `url.Values`.
func (r SuppressionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter to suppressions with this reason.
type SuppressionListParamsReason string

const (
	SuppressionListParamsReasonManual     SuppressionListParamsReason = "manual"
	SuppressionListParamsReasonComplaint  SuppressionListParamsReason = "complaint"
	SuppressionListParamsReasonHardBounce SuppressionListParamsReason = "hard_bounce"
	SuppressionListParamsReasonSoftBounce SuppressionListParamsReason = "soft_bounce"
	SuppressionListParamsReasonPolicy     SuppressionListParamsReason = "policy"
)

func (r SuppressionListParamsReason) IsKnown() bool {
	switch r {
	case SuppressionListParamsReasonManual, SuppressionListParamsReasonComplaint, SuppressionListParamsReasonHardBounce, SuppressionListParamsReasonSoftBounce, SuppressionListParamsReasonPolicy:
		return true
	}
	return false
}

// Filter by scope: `account` returns only account-wide suppressions,
// `sending_domain` only sending-domain suppressions. Omit to list both,
// sending-domain suppressions first.
type SuppressionListParamsScopeType string

const (
	SuppressionListParamsScopeTypeAccount       SuppressionListParamsScopeType = "account"
	SuppressionListParamsScopeTypeSendingDomain SuppressionListParamsScopeType = "sending_domain"
)

func (r SuppressionListParamsScopeType) IsKnown() bool {
	switch r {
	case SuppressionListParamsScopeTypeAccount, SuppressionListParamsScopeTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionDeleteParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SuppressionDeleteResponseEnvelope struct {
	Errors   []interface{}                         `json:"errors" api:"required"`
	Messages []interface{}                         `json:"messages" api:"required"`
	Result   SuppressionDeleteResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     suppressionDeleteResponseEnvelopeJSON `json:"-"`
}

// suppressionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionDeleteResponseEnvelope]
type suppressionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionEditParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// New expiry. Send `null` to make the suppression permanent; omit to leave it
	// unchanged.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Replacement advisory note. Send an empty string to clear it; omit to leave it
	// unchanged.
	Note param.Field[string] `json:"note"`
	// Not editable. Scope is fixed when the suppression is created; any value returns
	// 400 with code `scope_immutable`. Delete and recreate the suppression to change
	// it.
	Scope param.Field[interface{}] `json:"scope"`
}

func (r SuppressionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionEditResponseEnvelope struct {
	Errors   []interface{}                       `json:"errors" api:"required"`
	Messages []interface{}                       `json:"messages" api:"required"`
	Result   SuppressionEditResponse             `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     suppressionEditResponseEnvelopeJSON `json:"-"`
}

// suppressionEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionEditResponseEnvelope]
type suppressionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionGetParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SuppressionGetResponseEnvelope struct {
	Errors   []interface{}                      `json:"errors" api:"required"`
	Messages []interface{}                      `json:"messages" api:"required"`
	Result   SuppressionGetResponse             `json:"result" api:"required"`
	Success  bool                               `json:"success" api:"required"`
	JSON     suppressionGetResponseEnvelopeJSON `json:"-"`
}

// suppressionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionGetResponseEnvelope]
type suppressionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SuppressionImportParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Suppressions to import. Items with the same email address and scope are
	// deduplicated before processing.
	Items param.Field[[]SuppressionImportParamsItem] `json:"items" api:"required"`
}

func (r SuppressionImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuppressionImportParamsItem struct {
	// The email address to suppress.
	Email param.Field[string] `json:"email" api:"required"`
	// Expiration timestamp for the suppression. Omit or set to null for a permanent
	// suppression that never expires.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Advisory note for this suppression. Not enforced or validated beyond length.
	Note param.Field[string] `json:"note"`
	// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
	// the recipient for every sending domain of the account.
	Scope param.Field[SuppressionImportParamsItemsScopeUnion] `json:"scope"`
}

func (r SuppressionImportParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
// the recipient for every sending domain of the account.
type SuppressionImportParamsItemsScope struct {
	// Blocks the recipient for every sending domain of the account.
	Type param.Field[SuppressionImportParamsItemsScopeType] `json:"type" api:"required"`
	// The sending domain to suppress for: the domain part of the envelope MAIL FROM.
	// It is lowercased and trailing dots are removed. Internationalized domains must
	// use the ASCII (punycode) form. Ownership is not checked; a domain the account
	// does not send from never matches.
	Value param.Field[string] `json:"value"`
}

func (r SuppressionImportParamsItemsScope) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionImportParamsItemsScope) implementsSuppressionImportParamsItemsScopeUnion() {}

// Where the suppression applies. Omit for `{ "type": "account" }`, which blocks
// the recipient for every sending domain of the account.
//
// Satisfied by [email_sending.SuppressionImportParamsItemsScopeType],
// [email_sending.SuppressionImportParamsItemsScopeObject],
// [SuppressionImportParamsItemsScope].
type SuppressionImportParamsItemsScopeUnion interface {
	implementsSuppressionImportParamsItemsScopeUnion()
}

type SuppressionImportParamsItemsScopeType struct {
	// Blocks the recipient for every sending domain of the account.
	Type param.Field[SuppressionImportParamsItemsScopeTypeType] `json:"type" api:"required"`
}

func (r SuppressionImportParamsItemsScopeType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionImportParamsItemsScopeType) implementsSuppressionImportParamsItemsScopeUnion() {}

// Blocks the recipient for every sending domain of the account.
type SuppressionImportParamsItemsScopeTypeType string

const (
	SuppressionImportParamsItemsScopeTypeTypeAccount SuppressionImportParamsItemsScopeTypeType = "account"
)

func (r SuppressionImportParamsItemsScopeTypeType) IsKnown() bool {
	switch r {
	case SuppressionImportParamsItemsScopeTypeTypeAccount:
		return true
	}
	return false
}

type SuppressionImportParamsItemsScopeObject struct {
	// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
	Type param.Field[SuppressionImportParamsItemsScopeObjectType] `json:"type" api:"required"`
	// The sending domain to suppress for: the domain part of the envelope MAIL FROM.
	// It is lowercased and trailing dots are removed. Internationalized domains must
	// use the ASCII (punycode) form. Ownership is not checked; a domain the account
	// does not send from never matches.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r SuppressionImportParamsItemsScopeObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuppressionImportParamsItemsScopeObject) implementsSuppressionImportParamsItemsScopeUnion() {}

// Blocks the recipient only for mail whose envelope MAIL FROM uses `value`.
type SuppressionImportParamsItemsScopeObjectType string

const (
	SuppressionImportParamsItemsScopeObjectTypeSendingDomain SuppressionImportParamsItemsScopeObjectType = "sending_domain"
)

func (r SuppressionImportParamsItemsScopeObjectType) IsKnown() bool {
	switch r {
	case SuppressionImportParamsItemsScopeObjectTypeSendingDomain:
		return true
	}
	return false
}

type SuppressionImportResponseEnvelope struct {
	Errors   []interface{}                         `json:"errors" api:"required"`
	Messages []interface{}                         `json:"messages" api:"required"`
	Result   SuppressionImportResponse             `json:"result" api:"required"`
	Success  bool                                  `json:"success" api:"required"`
	JSON     suppressionImportResponseEnvelopeJSON `json:"-"`
}

// suppressionImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuppressionImportResponseEnvelope]
type suppressionImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuppressionImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r suppressionImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
