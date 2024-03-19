// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WorkerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWorkerService] method instead.
type WorkerService struct {
	Options         []option.RequestOption
	AI              *AIService
	Scripts         *ScriptService
	Filters         *FilterService
	Routes          *RouteService
	AccountSettings *AccountSettingService
	Deployments     *DeploymentService
	Domains         *DomainService
	Subdomains      *SubdomainService
	Services        *ServiceService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	r.AI = NewAIService(opts...)
	r.Scripts = NewScriptService(opts...)
	r.Filters = NewFilterService(opts...)
	r.Routes = NewRouteService(opts...)
	r.AccountSettings = NewAccountSettingService(opts...)
	r.Deployments = NewDeploymentService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Subdomains = NewSubdomainService(opts...)
	r.Services = NewServiceService(opts...)
	return
}
