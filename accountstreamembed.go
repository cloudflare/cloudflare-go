// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStreamEmbedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamEmbedService] method
// instead.
type AccountStreamEmbedService struct {
	Options []option.RequestOption
}

// NewAccountStreamEmbedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamEmbedService(opts ...option.RequestOption) (r *AccountStreamEmbedService) {
	r = &AccountStreamEmbedService{}
	r.Options = opts
	return
}

// Fetches an HTML code snippet to embed a video in a web page delivered through
// Cloudflare. On success, returns an HTML fragment for use on web pages to display
// a video. On failure, returns a JSON response body.
func (r *AccountStreamEmbedService) List(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamEmbedListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/embed", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamEmbedListResponse = interface{}
