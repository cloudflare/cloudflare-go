// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// InvestigateMoveService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateMoveService] method instead.
type InvestigateMoveService struct {
	Options []option.RequestOption
}

// NewInvestigateMoveService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateMoveService(opts ...option.RequestOption) (r *InvestigateMoveService) {
	r = &InvestigateMoveService{}
	r.Options = opts
	return
}

// Moves a single email message to a different folder or changes its quarantine
// status.
func (r *InvestigateMoveService) New(ctx context.Context, postfixID string, params InvestigateMoveNewParams, opts ...option.RequestOption) (res *[]InvestigateMoveNewResponse, err error) {
	var env InvestigateMoveNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/move", params.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Maximum batch size: 1000 messages per request
func (r *InvestigateMoveService) Bulk(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) (res *pagination.SinglePage[InvestigateMoveBulkResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/move", params.AccountID)
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

// Maximum batch size: 1000 messages per request
func (r *InvestigateMoveService) BulkAutoPaging(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InvestigateMoveBulkResponse] {
	return pagination.NewSinglePageAutoPager(r.Bulk(ctx, params, opts...))
}

type InvestigateMoveNewResponse struct {
	// Deprecated, use `completed_at` instead
	//
	// Deprecated: deprecated
	CompletedTimestamp time.Time `json:"completed_timestamp" api:"required" format:"date-time"`
	// Deprecated: deprecated
	ItemCount   int64                          `json:"item_count" api:"required"`
	Success     bool                           `json:"success" api:"required"`
	CompletedAt time.Time                      `json:"completed_at" format:"date-time"`
	Destination string                         `json:"destination" api:"nullable"`
	MessageID   string                         `json:"message_id" api:"nullable"`
	Operation   string                         `json:"operation" api:"nullable"`
	Recipient   string                         `json:"recipient" api:"nullable"`
	Status      string                         `json:"status" api:"nullable"`
	JSON        investigateMoveNewResponseJSON `json:"-"`
}

// investigateMoveNewResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveNewResponse]
type investigateMoveNewResponseJSON struct {
	CompletedTimestamp apijson.Field
	ItemCount          apijson.Field
	Success            apijson.Field
	CompletedAt        apijson.Field
	Destination        apijson.Field
	MessageID          apijson.Field
	Operation          apijson.Field
	Recipient          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateMoveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveBulkResponse struct {
	// Deprecated, use `completed_at` instead
	//
	// Deprecated: deprecated
	CompletedTimestamp time.Time `json:"completed_timestamp" api:"required" format:"date-time"`
	// Deprecated: deprecated
	ItemCount   int64                           `json:"item_count" api:"required"`
	Success     bool                            `json:"success" api:"required"`
	CompletedAt time.Time                       `json:"completed_at" format:"date-time"`
	Destination string                          `json:"destination" api:"nullable"`
	MessageID   string                          `json:"message_id" api:"nullable"`
	Operation   string                          `json:"operation" api:"nullable"`
	Recipient   string                          `json:"recipient" api:"nullable"`
	Status      string                          `json:"status" api:"nullable"`
	JSON        investigateMoveBulkResponseJSON `json:"-"`
}

// investigateMoveBulkResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveBulkResponse]
type investigateMoveBulkResponseJSON struct {
	CompletedTimestamp apijson.Field
	ItemCount          apijson.Field
	Success            apijson.Field
	CompletedAt        apijson.Field
	Destination        apijson.Field
	MessageID          apijson.Field
	Operation          apijson.Field
	Recipient          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestigateMoveBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveBulkResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveNewParams struct {
	// Account Identifier
	AccountID   param.Field[string]                              `path:"account_id" api:"required"`
	Destination param.Field[InvestigateMoveNewParamsDestination] `json:"destination" api:"required"`
	// When true, search the submissions datastore only. When false or omitted, search
	// the regular datastore only.
	Submission param.Field[bool] `query:"submission"`
}

func (r InvestigateMoveNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InvestigateMoveNewParams]'s query parameters as
// `url.Values`.
func (r InvestigateMoveNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InvestigateMoveNewParamsDestination string

const (
	InvestigateMoveNewParamsDestinationInbox                     InvestigateMoveNewParamsDestination = "Inbox"
	InvestigateMoveNewParamsDestinationJunkEmail                 InvestigateMoveNewParamsDestination = "JunkEmail"
	InvestigateMoveNewParamsDestinationDeletedItems              InvestigateMoveNewParamsDestination = "DeletedItems"
	InvestigateMoveNewParamsDestinationRecoverableItemsDeletions InvestigateMoveNewParamsDestination = "RecoverableItemsDeletions"
	InvestigateMoveNewParamsDestinationRecoverableItemsPurges    InvestigateMoveNewParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateMoveNewParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateMoveNewParamsDestinationInbox, InvestigateMoveNewParamsDestinationJunkEmail, InvestigateMoveNewParamsDestinationDeletedItems, InvestigateMoveNewParamsDestinationRecoverableItemsDeletions, InvestigateMoveNewParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateMoveNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                  `json:"messages" api:"required"`
	Result   []InvestigateMoveNewResponse           `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	JSON     investigateMoveNewResponseEnvelopeJSON `json:"-"`
}

// investigateMoveNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateMoveNewResponseEnvelope]
type investigateMoveNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateMoveNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveBulkParams struct {
	// Account Identifier
	AccountID   param.Field[string]                               `path:"account_id" api:"required"`
	Destination param.Field[InvestigateMoveBulkParamsDestination] `json:"destination" api:"required"`
	// List of message IDs to move.
	IDs param.Field[[]string] `json:"ids"`
	// Deprecated: Use `ids` instead. List of message IDs to move.
	PostfixIDs param.Field[[]string] `json:"postfix_ids"`
}

func (r InvestigateMoveBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateMoveBulkParamsDestination string

const (
	InvestigateMoveBulkParamsDestinationInbox                     InvestigateMoveBulkParamsDestination = "Inbox"
	InvestigateMoveBulkParamsDestinationJunkEmail                 InvestigateMoveBulkParamsDestination = "JunkEmail"
	InvestigateMoveBulkParamsDestinationDeletedItems              InvestigateMoveBulkParamsDestination = "DeletedItems"
	InvestigateMoveBulkParamsDestinationRecoverableItemsDeletions InvestigateMoveBulkParamsDestination = "RecoverableItemsDeletions"
	InvestigateMoveBulkParamsDestinationRecoverableItemsPurges    InvestigateMoveBulkParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateMoveBulkParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateMoveBulkParamsDestinationInbox, InvestigateMoveBulkParamsDestinationJunkEmail, InvestigateMoveBulkParamsDestinationDeletedItems, InvestigateMoveBulkParamsDestinationRecoverableItemsDeletions, InvestigateMoveBulkParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}
