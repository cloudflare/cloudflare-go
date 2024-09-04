// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
  "github.com/cloudflare/cloudflare-go/v2/option"
)

// PageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
Options []option.RequestOption
Projects *ProjectService
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r *PageService) {
  r = &PageService{}
  r.Options = opts
  r.Projects = NewProjectService(opts...)
  return
}
