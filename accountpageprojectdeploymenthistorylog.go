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

// AccountPageProjectDeploymentHistoryLogService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPageProjectDeploymentHistoryLogService] method instead.
type AccountPageProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentHistoryLogService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountPageProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentHistoryLogService) {
	r = &AccountPageProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *AccountPageProjectDeploymentHistoryLogService) PagesDeploymentGetDeploymentLogs(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse struct {
	Errors   []AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseError   `json:"errors"`
	Messages []AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessage `json:"messages"`
	Result   interface{}                                                                             `json:"result"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseSuccess `json:"success"`
	JSON    accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse]
type accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseError struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseErrorJSON `json:"-"`
}

// accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseError]
type accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessage struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessageJSON `json:"-"`
}

// accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessage]
type accountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseSuccess bool

const (
	AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseSuccessTrue AccountPageProjectDeploymentHistoryLogPagesDeploymentGetDeploymentLogsResponseSuccess = true
)
