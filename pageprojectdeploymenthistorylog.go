// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// PageProjectDeploymentHistoryLogService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewPageProjectDeploymentHistoryLogService] method instead.
type PageProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewPageProjectDeploymentHistoryLogService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPageProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *PageProjectDeploymentHistoryLogService) {
	r = &PageProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *PageProjectDeploymentHistoryLogService) List(ctx context.Context, projectName string, deploymentID string, query PageProjectDeploymentHistoryLogListParams, opts ...option.RequestOption) (res *PageProjectDeploymentHistoryLogListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentHistoryLogListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PageProjectDeploymentHistoryLogListResponseUnknown],
// [PageProjectDeploymentHistoryLogListResponseArray] or [shared.UnionString].
type PageProjectDeploymentHistoryLogListResponse interface {
	ImplementsPageProjectDeploymentHistoryLogListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDeploymentHistoryLogListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDeploymentHistoryLogListResponseArray []interface{}

func (r PageProjectDeploymentHistoryLogListResponseArray) ImplementsPageProjectDeploymentHistoryLogListResponse() {
}

type PageProjectDeploymentHistoryLogListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentHistoryLogListResponseEnvelope struct {
	Errors   []PageProjectDeploymentHistoryLogListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentHistoryLogListResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentHistoryLogListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentHistoryLogListResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentHistoryLogListResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentHistoryLogListResponseEnvelopeJSON contains the JSON
// metadata for the struct [PageProjectDeploymentHistoryLogListResponseEnvelope]
type pageProjectDeploymentHistoryLogListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentHistoryLogListResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentHistoryLogListResponseEnvelopeErrors]
type pageProjectDeploymentHistoryLogListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentHistoryLogListResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [PageProjectDeploymentHistoryLogListResponseEnvelopeMessages]
type pageProjectDeploymentHistoryLogListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentHistoryLogListResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentHistoryLogListResponseEnvelopeSuccessTrue PageProjectDeploymentHistoryLogListResponseEnvelopeSuccess = true
)
