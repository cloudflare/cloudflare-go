// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ListBulkOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListBulkOperationService] method instead.
type ListBulkOperationService struct {
	Options []option.RequestOption
}

// NewListBulkOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewListBulkOperationService(opts ...option.RequestOption) (r *ListBulkOperationService) {
	r = &ListBulkOperationService{}
	r.Options = opts
	return
}

// Gets the current status of an asynchronous operation on a list.
//
// The `status` property can have one of the following values: `pending`,
// `running`, `completed`, or `failed`. If the status is `failed`, the `error`
// property will contain a message describing the error.
func (r *ListBulkOperationService) Get(ctx context.Context, operationID string, query ListBulkOperationGetParams, opts ...option.RequestOption) (res *ListBulkOperationGetResponse, err error) {
	var env ListBulkOperationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/bulk_operations/%s", query.AccountID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ListBulkOperationGetResponse struct {
	// The unique operation ID of the asynchronous action.
	ID string `json:"id,required"`
	// The current status of the asynchronous operation.
	Status ListBulkOperationGetResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp of when the operation was completed.
	Completed string `json:"completed"`
	// A message describing the error when the status is `failed`.
	Error string                           `json:"error"`
	JSON  listBulkOperationGetResponseJSON `json:"-"`
}

// listBulkOperationGetResponseJSON contains the JSON metadata for the struct
// [ListBulkOperationGetResponse]
type listBulkOperationGetResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Completed   apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListBulkOperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listBulkOperationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The current status of the asynchronous operation.
type ListBulkOperationGetResponseStatus string

const (
	ListBulkOperationGetResponseStatusPending   ListBulkOperationGetResponseStatus = "pending"
	ListBulkOperationGetResponseStatusRunning   ListBulkOperationGetResponseStatus = "running"
	ListBulkOperationGetResponseStatusCompleted ListBulkOperationGetResponseStatus = "completed"
	ListBulkOperationGetResponseStatusFailed    ListBulkOperationGetResponseStatus = "failed"
)

func (r ListBulkOperationGetResponseStatus) IsKnown() bool {
	switch r {
	case ListBulkOperationGetResponseStatusPending, ListBulkOperationGetResponseStatusRunning, ListBulkOperationGetResponseStatusCompleted, ListBulkOperationGetResponseStatusFailed:
		return true
	}
	return false
}

type ListBulkOperationGetParams struct {
	// Defines an identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ListBulkOperationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   ListBulkOperationGetResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListBulkOperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    listBulkOperationGetResponseEnvelopeJSON    `json:"-"`
}

// listBulkOperationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ListBulkOperationGetResponseEnvelope]
type listBulkOperationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListBulkOperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listBulkOperationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListBulkOperationGetResponseEnvelopeSuccess bool

const (
	ListBulkOperationGetResponseEnvelopeSuccessTrue ListBulkOperationGetResponseEnvelopeSuccess = true
)

func (r ListBulkOperationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ListBulkOperationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
