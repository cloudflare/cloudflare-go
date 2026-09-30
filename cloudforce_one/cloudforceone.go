// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// CloudforceOneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudforceOneService] method instead.
type CloudforceOneService struct {
	Options       []option.RequestOption
	BinaryStorage *BinaryStorageService
	Requests      *RequestService
	Scans         *ScanService
	ThreatEvents  *ThreatEventService
	// Threat Signals API for managing threat intelligence feeds, articles, indicators,
	// and AI skills in Cloudforce One.
	//
	// ## Prerequisites
	//
	// 1. **API token** — requests must use an API token with Cloudforce One
	//    permissions; write operations (creating, editing, or deleting feeds, skills,
	//    and tags) require write access.
	// 2. **Plan limits** — access on the Free plan is limited; feed quotas and managed
	//    default skills apply.
	ThreatSignals *ThreatSignalService
}

// NewCloudforceOneService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudforceOneService(opts ...option.RequestOption) (r *CloudforceOneService) {
	r = &CloudforceOneService{}
	r.Options = opts
	r.BinaryStorage = NewBinaryStorageService(opts...)
	r.Requests = NewRequestService(opts...)
	r.Scans = NewScanService(opts...)
	r.ThreatEvents = NewThreatEventService(opts...)
	r.ThreatSignals = NewThreatSignalService(opts...)
	return
}
