// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPageProjectDeploymentRetryService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentRetryService] method instead.
type AccountPageProjectDeploymentRetryService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentRetryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDeploymentRetryService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentRetryService) {
	r = &AccountPageProjectDeploymentRetryService{}
	r.Options = opts
	return
}

// Retry a previous deployment.
func (r *AccountPageProjectDeploymentRetryService) PagesDeploymentRetryDeployment(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *DeploymentNewDeployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
