// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RuleListBulkOperationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRuleListBulkOperationService]
// method instead.
type RuleListBulkOperationService struct {
	Options []option.RequestOption
}

// NewRuleListBulkOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRuleListBulkOperationService(opts ...option.RequestOption) (r *RuleListBulkOperationService) {
	r = &RuleListBulkOperationService{}
	r.Options = opts
	return
}

// Gets the current status of an asynchronous operation on a list.
//
// The `status` property can have one of the following values: `pending`,
// `running`, `completed`, or `failed`. If the status is `failed`, the `error`
// property will contain a message describing the error.
func (r *RuleListBulkOperationService) Get(ctx context.Context, accountIdentifier string, operationID string, opts ...option.RequestOption) (res *[]RuleListBulkOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListBulkOperationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/bulk_operations/%s", accountIdentifier, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleListBulkOperationGetResponse = interface{}

type RuleListBulkOperationGetResponseEnvelope struct {
	Errors   []RuleListBulkOperationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListBulkOperationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListBulkOperationGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListBulkOperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListBulkOperationGetResponseEnvelopeJSON    `json:"-"`
}

// ruleListBulkOperationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RuleListBulkOperationGetResponseEnvelope]
type ruleListBulkOperationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListBulkOperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListBulkOperationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleListBulkOperationGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    ruleListBulkOperationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListBulkOperationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [RuleListBulkOperationGetResponseEnvelopeErrors]
type ruleListBulkOperationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListBulkOperationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListBulkOperationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleListBulkOperationGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    ruleListBulkOperationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListBulkOperationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RuleListBulkOperationGetResponseEnvelopeMessages]
type ruleListBulkOperationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListBulkOperationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListBulkOperationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleListBulkOperationGetResponseEnvelopeSuccess bool

const (
	RuleListBulkOperationGetResponseEnvelopeSuccessTrue RuleListBulkOperationGetResponseEnvelopeSuccess = true
)
