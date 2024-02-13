// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *PageProjectDeploymentHistoryLogService) PagesDeploymentGetDeploymentLogs(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseUnknown],
// [PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseArray]
// or [shared.UnionString].
type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse interface {
	ImplementsPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseArray []interface{}

func (r PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseArray) ImplementsPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse() {
}

type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelope struct {
	Errors   []PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelope]
type pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrors struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrors]
type pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessages struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessages]
type pageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeSuccess bool

const (
	PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeSuccessTrue PageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseEnvelopeSuccess = true
)
