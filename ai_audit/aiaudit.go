// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_audit

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AIAuditService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAuditService] method instead.
type AIAuditService struct {
	Options []option.RequestOption
	Robots  *RobotService
}

// NewAIAuditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIAuditService(opts ...option.RequestOption) (r *AIAuditService) {
	r = &AIAuditService{}
	r.Options = opts
	r.Robots = NewRobotService(opts...)
	return
}
