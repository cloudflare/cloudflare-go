// File generated from our OpenAPI spec by Stainless.

package pages

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ProjectDeploymentHistoryLogService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewProjectDeploymentHistoryLogService] method instead.
type ProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentHistoryLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *ProjectDeploymentHistoryLogService) {
	r = &ProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *ProjectDeploymentHistoryLogService) Get(ctx context.Context, projectName string, deploymentID string, query ProjectDeploymentHistoryLogGetParams, opts ...option.RequestOption) (res *ProjectDeploymentHistoryLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProjectDeploymentHistoryLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [pages.ProjectDeploymentHistoryLogGetResponseUnknown],
// [pages.ProjectDeploymentHistoryLogGetResponseArray] or [shared.UnionString].
type ProjectDeploymentHistoryLogGetResponse interface {
	ImplementsPagesProjectDeploymentHistoryLogGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentHistoryLogGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDeploymentHistoryLogGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDeploymentHistoryLogGetResponseArray []interface{}

func (r ProjectDeploymentHistoryLogGetResponseArray) ImplementsPagesProjectDeploymentHistoryLogGetResponse() {
}

type ProjectDeploymentHistoryLogGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeploymentHistoryLogGetResponseEnvelope struct {
	Errors   []ProjectDeploymentHistoryLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeploymentHistoryLogGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeploymentHistoryLogGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ProjectDeploymentHistoryLogGetResponseEnvelope]
type projectDeploymentHistoryLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ProjectDeploymentHistoryLogGetResponseEnvelopeErrors]
type projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeploymentHistoryLogGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ProjectDeploymentHistoryLogGetResponseEnvelopeMessages]
type projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess = true
)
