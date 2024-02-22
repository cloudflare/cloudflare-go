// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// Web3HostnameIPFSUniversalPathService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWeb3HostnameIPFSUniversalPathService] method instead.
type Web3HostnameIPFSUniversalPathService struct {
	Options      []option.RequestOption
	ContentLists *Web3HostnameIPFSUniversalPathContentListService
}

// NewWeb3HostnameIPFSUniversalPathService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWeb3HostnameIPFSUniversalPathService(opts ...option.RequestOption) (r *Web3HostnameIPFSUniversalPathService) {
	r = &Web3HostnameIPFSUniversalPathService{}
	r.Options = opts
	r.ContentLists = NewWeb3HostnameIPFSUniversalPathContentListService(opts...)
	return
}
