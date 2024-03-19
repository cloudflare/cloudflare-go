// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DispatchNamespaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDispatchNamespaceService] method
// instead.
type DispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *DispatchNamespaceScriptService
}

// NewDispatchNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceService(opts ...option.RequestOption) (r *DispatchNamespaceService) {
	r = &DispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewDispatchNamespaceScriptService(opts...)
	return
}

// Create a new Workers for Platforms namespace.
func (r *DispatchNamespaceService) New(ctx context.Context, params DispatchNamespaceNewParams, opts ...option.RequestOption) (res *DispatchNamespaceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of Workers for Platforms namespaces.
func (r *DispatchNamespaceService) List(ctx context.Context, query DispatchNamespaceListParams, opts ...option.RequestOption) (res *[]DispatchNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Delete(ctx context.Context, dispatchNamespace string, body DispatchNamespaceDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", body.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Get(ctx context.Context, dispatchNamespace string, query DispatchNamespaceGetParams, opts ...option.RequestOption) (res *DispatchNamespaceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", query.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceNewResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string                           `json:"namespace_name"`
	JSON          dispatchNamespaceNewResponseJSON `json:"-"`
}

// dispatchNamespaceNewResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceNewResponse]
type dispatchNamespaceNewResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceListResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string                            `json:"namespace_name"`
	JSON          dispatchNamespaceListResponseJSON `json:"-"`
}

// dispatchNamespaceListResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceListResponse]
type dispatchNamespaceListResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponse = interface{}

type DispatchNamespaceGetResponse struct {
	// Identifier
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string                           `json:"namespace_name"`
	JSON          dispatchNamespaceGetResponseJSON `json:"-"`
}

// dispatchNamespaceGetResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceGetResponse]
type dispatchNamespaceGetResponseJSON struct {
	CreatedBy     apijson.Field
	CreatedOn     apijson.Field
	ModifiedBy    apijson.Field
	ModifiedOn    apijson.Field
	NamespaceID   apijson.Field
	NamespaceName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the dispatch namespace
	Name param.Field[string] `json:"name"`
}

func (r DispatchNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceNewResponseEnvelope struct {
	Errors   []DispatchNamespaceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceNewResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceNewResponseEnvelope]
type dispatchNamespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dispatchNamespaceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceNewResponseEnvelopeErrors]
type dispatchNamespaceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dispatchNamespaceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatchNamespaceNewResponseEnvelopeMessages]
type dispatchNamespaceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceNewResponseEnvelopeSuccess bool

const (
	DispatchNamespaceNewResponseEnvelopeSuccessTrue DispatchNamespaceNewResponseEnvelopeSuccess = true
)

type DispatchNamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceListResponseEnvelope struct {
	Errors   []DispatchNamespaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DispatchNamespaceListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceListResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceListResponseEnvelope]
type dispatchNamespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dispatchNamespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceListResponseEnvelopeErrors]
type dispatchNamespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dispatchNamespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatchNamespaceListResponseEnvelopeMessages]
type dispatchNamespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceListResponseEnvelopeSuccess bool

const (
	DispatchNamespaceListResponseEnvelopeSuccessTrue DispatchNamespaceListResponseEnvelopeSuccess = true
)

type DispatchNamespaceDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DispatchNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceDeleteResponseEnvelope]
type dispatchNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceDeleteResponseEnvelopeErrors]
type dispatchNamespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DispatchNamespaceDeleteResponseEnvelopeMessages]
type dispatchNamespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceDeleteResponseEnvelopeSuccessTrue DispatchNamespaceDeleteResponseEnvelopeSuccess = true
)

type DispatchNamespaceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceGetResponseEnvelope struct {
	Errors   []DispatchNamespaceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceGetResponseEnvelope]
type dispatchNamespaceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dispatchNamespaceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceGetResponseEnvelopeErrors]
type dispatchNamespaceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dispatchNamespaceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatchNamespaceGetResponseEnvelopeMessages]
type dispatchNamespaceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceGetResponseEnvelopeSuccessTrue DispatchNamespaceGetResponseEnvelopeSuccess = true
)
