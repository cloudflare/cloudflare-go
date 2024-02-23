// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWorkerService] method instead.
type WorkerService struct {
	Options             []option.RequestOption
	AI                  *WorkerAIService
	Scripts             *WorkerScriptService
	Filters             *WorkerFilterService
	Routes              *WorkerRouteService
	AccountSettings     *WorkerAccountSettingService
	Deployments         *WorkerDeploymentService
	Domains             *WorkerDomainService
	DurableObjects      *WorkerDurableObjectService
	Queues              *WorkerQueueService
	Subdomains          *WorkerSubdomainService
	DeploymentsByScript *WorkerDeploymentsByScriptService
	Services            *WorkerServiceService
	Script              *WorkerScriptService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	r.AI = NewWorkerAIService(opts...)
	r.Scripts = NewWorkerScriptService(opts...)
	r.Filters = NewWorkerFilterService(opts...)
	r.Routes = NewWorkerRouteService(opts...)
	r.AccountSettings = NewWorkerAccountSettingService(opts...)
	r.Deployments = NewWorkerDeploymentService(opts...)
	r.Domains = NewWorkerDomainService(opts...)
	r.DurableObjects = NewWorkerDurableObjectService(opts...)
	r.Queues = NewWorkerQueueService(opts...)
	r.Subdomains = NewWorkerSubdomainService(opts...)
	r.DeploymentsByScript = NewWorkerDeploymentsByScriptService(opts...)
	r.Services = NewWorkerServiceService(opts...)
	r.Script = NewWorkerScriptService(opts...)
	return
}
