// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SnippetContentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSnippetContentService] method
// instead.
type SnippetContentService struct {
	Options []option.RequestOption
}

// NewSnippetContentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnippetContentService(opts ...option.RequestOption) (r *SnippetContentService) {
	r = &SnippetContentService{}
	r.Options = opts
	return
}

// Snippet Content
func (r *SnippetContentService) Get(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/form-data")}, opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s/content", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
