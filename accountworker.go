// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWorkerService] method
// instead.
type AccountWorkerService struct {
	Options         []option.RequestOption
	AccountSettings *AccountWorkerAccountSettingService
	Deployments     *AccountWorkerDeploymentService
	Domains         *AccountWorkerDomainService
	DurableObjects  *AccountWorkerDurableObjectService
	Queues          *AccountWorkerQueueService
	Scripts         *AccountWorkerScriptService
	Subdomains      *AccountWorkerSubdomainService
}

// NewAccountWorkerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWorkerService(opts ...option.RequestOption) (r *AccountWorkerService) {
	r = &AccountWorkerService{}
	r.Options = opts
	r.AccountSettings = NewAccountWorkerAccountSettingService(opts...)
	r.Deployments = NewAccountWorkerDeploymentService(opts...)
	r.Domains = NewAccountWorkerDomainService(opts...)
	r.DurableObjects = NewAccountWorkerDurableObjectService(opts...)
	r.Queues = NewAccountWorkerQueueService(opts...)
	r.Scripts = NewAccountWorkerScriptService(opts...)
	r.Subdomains = NewAccountWorkerSubdomainService(opts...)
	return
}
