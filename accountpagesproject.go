// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPagesProjectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountPagesProjectService]
// method instead.
type AccountPagesProjectService struct {
	Options []option.RequestOption
}

// NewAccountPagesProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPagesProjectService(opts ...option.RequestOption) (r *AccountPagesProjectService) {
	r = &AccountPagesProjectService{}
	r.Options = opts
	return
}

// Purge all cached build artifacts for a Pages project
func (r *AccountPagesProjectService) PurgeBuildCache(ctx context.Context, accountIdentifier string, projectName string, opts ...option.RequestOption) (res *AccountPagesProjectPurgeBuildCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", accountIdentifier, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountPagesProjectPurgeBuildCacheResponse = interface{}
