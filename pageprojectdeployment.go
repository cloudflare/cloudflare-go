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

// PageProjectDeploymentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageProjectDeploymentService]
// method instead.
type PageProjectDeploymentService struct {
	Options []option.RequestOption
	History *PageProjectDeploymentHistoryService
}

// NewPageProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentService(opts ...option.RequestOption) (r *PageProjectDeploymentService) {
	r = &PageProjectDeploymentService{}
	r.Options = opts
	r.History = NewPageProjectDeploymentHistoryService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *PageProjectDeploymentService) New(ctx context.Context, projectName string, params PageProjectDeploymentNewParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of project deployments.
func (r *PageProjectDeploymentService) List(ctx context.Context, projectName string, query PageProjectDeploymentListParams, opts ...option.RequestOption) (res *[]PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a deployment.
func (r *PageProjectDeploymentService) Delete(ctx context.Context, projectName string, deploymentID string, body PageProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *PageProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetch information about a deployment.
func (r *PageProjectDeploymentService) Get(ctx context.Context, projectName string, deploymentID string, query PageProjectDeploymentGetParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry a previous deployment.
func (r *PageProjectDeploymentService) Retry(ctx context.Context, projectName string, deploymentID string, body PageProjectDeploymentRetryParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentRetryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *PageProjectDeploymentService) Rollback(ctx context.Context, projectName string, deploymentID string, body PageProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *PagesDeployments, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentRollbackResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", body.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageProjectDeploymentDeleteResponse = interface{}

type PageProjectDeploymentNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r PageProjectDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageProjectDeploymentNewResponseEnvelope struct {
	Errors   []PageProjectDeploymentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentNewResponseEnvelope]
type pageProjectDeploymentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentNewResponseEnvelopeErrors]
type pageProjectDeploymentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageProjectDeploymentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentNewResponseEnvelopeMessages]
type pageProjectDeploymentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentNewResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentNewResponseEnvelopeSuccessTrue PageProjectDeploymentNewResponseEnvelopeSuccess = true
)

type PageProjectDeploymentListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentListResponseEnvelope struct {
	Errors   []PageProjectDeploymentListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PagesDeployments                                  `json:"result,required"`
	// Whether the API call was successful
	Success    PageProjectDeploymentListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageProjectDeploymentListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageProjectDeploymentListResponseEnvelopeJSON       `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentListResponseEnvelope]
type pageProjectDeploymentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    pageProjectDeploymentListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentListResponseEnvelopeErrors]
type pageProjectDeploymentListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    pageProjectDeploymentListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentListResponseEnvelopeMessages]
type pageProjectDeploymentListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentListResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentListResponseEnvelopeSuccessTrue PageProjectDeploymentListResponseEnvelopeSuccess = true
)

type PageProjectDeploymentListResponseEnvelopeResultInfo struct {
	Count      interface{}                                             `json:"count"`
	Page       interface{}                                             `json:"page"`
	PerPage    interface{}                                             `json:"per_page"`
	TotalCount interface{}                                             `json:"total_count"`
	JSON       pageProjectDeploymentListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageProjectDeploymentListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [PageProjectDeploymentListResponseEnvelopeResultInfo]
type pageProjectDeploymentListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentGetResponseEnvelope struct {
	Errors   []PageProjectDeploymentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDeploymentGetResponseEnvelope]
type pageProjectDeploymentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentGetResponseEnvelopeErrors]
type pageProjectDeploymentGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageProjectDeploymentGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDeploymentGetResponseEnvelopeMessages]
type pageProjectDeploymentGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentGetResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentGetResponseEnvelopeSuccessTrue PageProjectDeploymentGetResponseEnvelopeSuccess = true
)

type PageProjectDeploymentRetryParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentRetryResponseEnvelope struct {
	Errors   []PageProjectDeploymentRetryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentRetryResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                     `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentRetryResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentRetryResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentRetryResponseEnvelopeJSON contains the JSON metadata for
// the struct [PageProjectDeploymentRetryResponseEnvelope]
type pageProjectDeploymentRetryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRetryResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageProjectDeploymentRetryResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentRetryResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageProjectDeploymentRetryResponseEnvelopeErrors]
type pageProjectDeploymentRetryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRetryResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    pageProjectDeploymentRetryResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentRetryResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [PageProjectDeploymentRetryResponseEnvelopeMessages]
type pageProjectDeploymentRetryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRetryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentRetryResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentRetryResponseEnvelopeSuccessTrue PageProjectDeploymentRetryResponseEnvelopeSuccess = true
)

type PageProjectDeploymentRollbackParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentRollbackResponseEnvelope struct {
	Errors   []PageProjectDeploymentRollbackResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentRollbackResponseEnvelopeMessages `json:"messages,required"`
	Result   PagesDeployments                                        `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentRollbackResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentRollbackResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentRollbackResponseEnvelopeJSON contains the JSON metadata for
// the struct [PageProjectDeploymentRollbackResponseEnvelope]
type pageProjectDeploymentRollbackResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRollbackResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    pageProjectDeploymentRollbackResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentRollbackResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [PageProjectDeploymentRollbackResponseEnvelopeErrors]
type pageProjectDeploymentRollbackResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentRollbackResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    pageProjectDeploymentRollbackResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentRollbackResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [PageProjectDeploymentRollbackResponseEnvelopeMessages]
type pageProjectDeploymentRollbackResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentRollbackResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentRollbackResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentRollbackResponseEnvelopeSuccessTrue PageProjectDeploymentRollbackResponseEnvelopeSuccess = true
)
