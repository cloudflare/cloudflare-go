// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RuleListService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRuleListService] method instead.
type RuleListService struct {
	Options        []option.RequestOption
	BulkOperations *RuleListBulkOperationService
	Items          *RuleListItemService
}

// NewRuleListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleListService(opts ...option.RequestOption) (r *RuleListService) {
	r = &RuleListService{}
	r.Options = opts
	r.BulkOperations = NewRuleListBulkOperationService(opts...)
	r.Items = NewRuleListItemService(opts...)
	return
}

// Creates a new list of the specified type.
func (r *RuleListService) New(ctx context.Context, accountID string, body RuleListNewParams, opts ...option.RequestOption) (res *[]RuleListNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the description of a list.
func (r *RuleListService) Update(ctx context.Context, accountID string, listID string, body RuleListUpdateParams, opts ...option.RequestOption) (res *[]RuleListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all lists in the account.
func (r *RuleListService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]RuleListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a specific list and all its items.
func (r *RuleListService) Delete(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *RuleListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a list.
func (r *RuleListService) Get(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *[]RuleListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleListGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleListNewResponse = interface{}

type RuleListUpdateResponse = interface{}

type RuleListListResponse struct {
	// The unique ID of the list.
	ID string `json:"id,required"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on,required"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind RuleListListResponseKind `json:"kind,required"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name,required"`
	// The number of items in the list.
	NumItems float64 `json:"num_items,required"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64                  `json:"num_referencing_filters"`
	JSON                  ruleListListResponseJSON `json:"-"`
}

// ruleListListResponseJSON contains the JSON metadata for the struct
// [RuleListListResponse]
type ruleListListResponseJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	Description           apijson.Field
	NumReferencingFilters apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RuleListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type RuleListListResponseKind string

const (
	RuleListListResponseKindIP       RuleListListResponseKind = "ip"
	RuleListListResponseKindRedirect RuleListListResponseKind = "redirect"
	RuleListListResponseKindHostname RuleListListResponseKind = "hostname"
	RuleListListResponseKindASN      RuleListListResponseKind = "asn"
)

type RuleListDeleteResponse struct {
	// The unique ID of the item in the List.
	ID   string                     `json:"id"`
	JSON ruleListDeleteResponseJSON `json:"-"`
}

// ruleListDeleteResponseJSON contains the JSON metadata for the struct
// [RuleListDeleteResponse]
type ruleListDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListGetResponse = interface{}

type RuleListNewParams struct {
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind param.Field[RuleListNewParamsKind] `json:"kind,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name param.Field[string] `json:"name,required"`
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r RuleListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type RuleListNewParamsKind string

const (
	RuleListNewParamsKindIP       RuleListNewParamsKind = "ip"
	RuleListNewParamsKindRedirect RuleListNewParamsKind = "redirect"
	RuleListNewParamsKindHostname RuleListNewParamsKind = "hostname"
	RuleListNewParamsKindASN      RuleListNewParamsKind = "asn"
)

type RuleListNewResponseEnvelope struct {
	Errors   []RuleListNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListNewResponseEnvelopeJSON    `json:"-"`
}

// ruleListNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListNewResponseEnvelope]
type ruleListNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    ruleListNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleListNewResponseEnvelopeErrors]
type ruleListNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    ruleListNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListNewResponseEnvelopeMessages]
type ruleListNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListNewResponseEnvelopeSuccess bool

const (
	RuleListNewResponseEnvelopeSuccessTrue RuleListNewResponseEnvelopeSuccess = true
)

type RuleListUpdateParams struct {
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r RuleListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListUpdateResponseEnvelope struct {
	Errors   []RuleListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleListUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListUpdateResponseEnvelope]
type ruleListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    ruleListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListUpdateResponseEnvelopeErrors]
type ruleListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    ruleListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListUpdateResponseEnvelopeMessages]
type ruleListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListUpdateResponseEnvelopeSuccess bool

const (
	RuleListUpdateResponseEnvelopeSuccessTrue RuleListUpdateResponseEnvelopeSuccess = true
)

type RuleListListResponseEnvelope struct {
	Errors   []RuleListListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListListResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListListResponseEnvelopeJSON    `json:"-"`
}

// ruleListListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListListResponseEnvelope]
type ruleListListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleListListResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleListListResponseEnvelopeErrors]
type ruleListListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListListResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    ruleListListResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListListResponseEnvelopeMessages]
type ruleListListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListListResponseEnvelopeSuccess bool

const (
	RuleListListResponseEnvelopeSuccessTrue RuleListListResponseEnvelopeSuccess = true
)

type RuleListDeleteResponseEnvelope struct {
	Errors   []RuleListDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleListDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleListDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListDeleteResponseEnvelope]
type ruleListDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    ruleListDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleListDeleteResponseEnvelopeErrors]
type ruleListDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    ruleListDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListDeleteResponseEnvelopeMessages]
type ruleListDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListDeleteResponseEnvelopeSuccess bool

const (
	RuleListDeleteResponseEnvelopeSuccessTrue RuleListDeleteResponseEnvelopeSuccess = true
)

type RuleListGetResponseEnvelope struct {
	Errors   []RuleListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleListGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListGetResponseEnvelopeJSON    `json:"-"`
}

// ruleListGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListGetResponseEnvelope]
type ruleListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    ruleListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleListGetResponseEnvelopeErrors]
type ruleListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleListGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    ruleListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleListGetResponseEnvelopeMessages]
type ruleListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RuleListGetResponseEnvelopeSuccess bool

const (
	RuleListGetResponseEnvelopeSuccessTrue RuleListGetResponseEnvelopeSuccess = true
)
