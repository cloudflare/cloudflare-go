// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// DomService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomService] method instead.
type DomService struct {
	Options []option.RequestOption
}

// NewDomService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomService(opts ...option.RequestOption) (r *DomService) {
	r = &DomService{}
	r.Options = opts
	return
}

// Returns a plain text response, with the scan's DOM content as rendered by
// Chrome.
func (r *DomService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/dom/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
