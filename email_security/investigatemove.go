// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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

// Move a message
func (r *InvestigateMoveService) New(ctx context.Context, postfixID string, params InvestigateMoveNewParams, opts ...option.RequestOption) (res *[]InvestigateMoveNewResponse, err error) {
	var env InvestigateMoveNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/move", params.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Move multiple messages
func (r *InvestigateMoveService) Bulk(ctx context.Context, params InvestigateMoveBulkParams, opts ...option.RequestOption) (res *[]InvestigateMoveBulkResponse, err error) {
	var env InvestigateMoveBulkResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/move", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateMoveNewResponse struct {
	CompletedTimestamp time.Time                      `json:"completed_timestamp,required" format:"date-time"`
	Destination        string                         `json:"destination,required"`
	ItemCount          int64                          `json:"item_count,required"`
	MessageID          string                         `json:"message_id,required"`
	Operation          string                         `json:"operation,required"`
	Recipient          string                         `json:"recipient,required"`
	Status             string                         `json:"status,required"`
	JSON               investigateMoveNewResponseJSON `json:"-"`
}

// investigateMoveNewResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveNewResponse]
type investigateMoveNewResponseJSON struct {
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

func (r *InvestigateMoveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateMoveBulkResponse struct {
	CompletedTimestamp time.Time                       `json:"completed_timestamp,required" format:"date-time"`
	Destination        string                          `json:"destination,required"`
	ItemCount          int64                           `json:"item_count,required"`
	MessageID          string                          `json:"message_id,required"`
	Operation          string                          `json:"operation,required"`
	Recipient          string                          `json:"recipient,required"`
	Status             string                          `json:"status,required"`
	JSON               investigateMoveBulkResponseJSON `json:"-"`
}

// investigateMoveBulkResponseJSON contains the JSON metadata for the struct
// [InvestigateMoveBulkResponse]
type investigateMoveBulkResponseJSON struct {
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

type InvestigateMoveNewParams struct {
	// Account Identifier
	AccountID   param.Field[string]                              `path:"account_id,required"`
	Destination param.Field[InvestigateMoveNewParamsDestination] `json:"destination,required"`
}

func (r InvestigateMoveNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   []InvestigateMoveNewResponse           `json:"result,required"`
	Success  bool                                   `json:"success,required"`
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
	AccountID   param.Field[string]                               `path:"account_id,required"`
	Destination param.Field[InvestigateMoveBulkParamsDestination] `json:"destination,required"`
	PostfixIDs  param.Field[[]string]                             `json:"postfix_ids,required"`
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

type InvestigateMoveBulkResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   []InvestigateMoveBulkResponse           `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     investigateMoveBulkResponseEnvelopeJSON `json:"-"`
}

// investigateMoveBulkResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateMoveBulkResponseEnvelope]
type investigateMoveBulkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateMoveBulkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateMoveBulkResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
