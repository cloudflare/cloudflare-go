// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPageProjectDeploymentRollbackService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPageProjectDeploymentRollbackService] method instead.
type AccountPageProjectDeploymentRollbackService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentRollbackService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountPageProjectDeploymentRollbackService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentRollbackService) {
	r = &AccountPageProjectDeploymentRollbackService{}
	r.Options = opts
	return
}

// Rollback the production deployment to a previous deploy. You can only rollback
// to succesful builds on production.
func (r *AccountPageProjectDeploymentRollbackService) PagesDeploymentRollbackDeployment(ctx context.Context, accountIdentifier string, projectName string, deploymentIdentifier string, opts ...option.RequestOption) (res *DeploymentResponseDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", accountIdentifier, projectName, deploymentIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
