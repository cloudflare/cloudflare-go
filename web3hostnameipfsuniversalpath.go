// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// Web3HostnameIpfsUniversalPathService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWeb3HostnameIpfsUniversalPathService] method instead.
type Web3HostnameIpfsUniversalPathService struct {
	Options      []option.RequestOption
	ContentLists *Web3HostnameIpfsUniversalPathContentListService
}

// NewWeb3HostnameIpfsUniversalPathService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWeb3HostnameIpfsUniversalPathService(opts ...option.RequestOption) (r *Web3HostnameIpfsUniversalPathService) {
	r = &Web3HostnameIpfsUniversalPathService{}
	r.Options = opts
	r.ContentLists = NewWeb3HostnameIpfsUniversalPathContentListService(opts...)
	return
}
