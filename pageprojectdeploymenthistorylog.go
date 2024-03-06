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
func (r *PageProjectDeploymentHistoryLogService) Get(ctx context.Context, projectName string, deploymentID string, query PageProjectDeploymentHistoryLogGetParams, opts ...option.RequestOption) (res *PageProjectDeploymentHistoryLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentHistoryLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", query.AccountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PageProjectDeploymentHistoryLogGetResponseUnknown],
// [PageProjectDeploymentHistoryLogGetResponseArray] or [shared.UnionString].
type PageProjectDeploymentHistoryLogGetResponse interface {
	ImplementsPageProjectDeploymentHistoryLogGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDeploymentHistoryLogGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PageProjectDeploymentHistoryLogGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDeploymentHistoryLogGetResponseArray []interface{}

func (r PageProjectDeploymentHistoryLogGetResponseArray) ImplementsPageProjectDeploymentHistoryLogGetResponse() {
}

type PageProjectDeploymentHistoryLogGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PageProjectDeploymentHistoryLogGetResponseEnvelope struct {
	Errors   []PageProjectDeploymentHistoryLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentHistoryLogGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentHistoryLogGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentHistoryLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentHistoryLogGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentHistoryLogGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [PageProjectDeploymentHistoryLogGetResponseEnvelope]
type pageProjectDeploymentHistoryLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDeploymentHistoryLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageProjectDeploymentHistoryLogGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentHistoryLogGetResponseEnvelopeErrors]
type pageProjectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDeploymentHistoryLogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageProjectDeploymentHistoryLogGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [PageProjectDeploymentHistoryLogGetResponseEnvelopeMessages]
type pageProjectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageProjectDeploymentHistoryLogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageProjectDeploymentHistoryLogGetResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue PageProjectDeploymentHistoryLogGetResponseEnvelopeSuccess = true
)
