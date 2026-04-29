// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// Moves multiple messages to a specified mailbox folder (Inbox, JunkEmail,
// DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires
// active integration.
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

// Moves multiple messages to a specified mailbox folder (Inbox, JunkEmail,
// DeletedItems, RecoverableItemsDeletions, or RecoverableItemsPurges). Requires
// active integration.
func (r *InvestigateMoveService) BulkAutoPaging(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InvestigateMoveBulkResponse] {
	return pagination.NewSinglePageAutoPager(r.Bulk(ctx, params, opts...))
}

type InvestigateMoveBulkResponse struct {
	// Whether the operation succeeded
	Success bool `json:"success" api:"required"`
	// When the move operation completed (UTC)
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Deprecated, use `completed_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	CompletedTimestamp time.Time `json:"completed_timestamp" format:"date-time"`
	// Destination folder for the message
	Destination string `json:"destination" api:"nullable"`
	// Number of items moved. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	ItemCount int64 `json:"item_count"`
	// Message identifier
	MessageID string `json:"message_id" api:"nullable"`
	// Type of operation performed
	Operation string `json:"operation" api:"nullable"`
	// Recipient email address
	Recipient string `json:"recipient" api:"nullable"`
	// Operation status
	Status string                          `json:"status" api:"nullable"`
	JSON   investigateMoveBulkResponseJSON `json:"-"`
}

// investigateMoveBulkResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveBulkResponse]
type investigateMoveBulkResponseJSON struct {
	Success            apijson.Field
	CompletedAt        apijson.Field
	CompletedTimestamp apijson.Field
	Destination        apijson.Field
	ItemCount          apijson.Field
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

type InvestigateMoveBulkParams struct {
	// Identifier.
	AccountID   param.Field[string]                               `path:"account_id" api:"required"`
	Destination param.Field[InvestigateMoveBulkParamsDestination] `json:"destination" api:"required"`
	// List of message IDs to move
	IDs param.Field[[]string] `json:"ids"`
	// Deprecated, use `ids` instead. End of life: November 1, 2026. List of message
	// IDs to move.
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
