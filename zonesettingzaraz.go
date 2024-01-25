// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingZarazService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingZarazService] method
// instead.
type ZoneSettingZarazService struct {
	Options       []option.RequestOption
	Config        *ZoneSettingZarazConfigService
	Default       *ZoneSettingZarazDefaultService
	Export        *ZoneSettingZarazExportService
	History       *ZoneSettingZarazHistoryService
	ConfigHistory *ZoneSettingZarazConfigHistoryService
	Publish       *ZoneSettingZarazPublishService
}

// NewZoneSettingZarazService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazService(opts ...option.RequestOption) (r *ZoneSettingZarazService) {
	r = &ZoneSettingZarazService{}
	r.Options = opts
	r.Config = NewZoneSettingZarazConfigService(opts...)
	r.Default = NewZoneSettingZarazDefaultService(opts...)
	r.Export = NewZoneSettingZarazExportService(opts...)
	r.History = NewZoneSettingZarazHistoryService(opts...)
	r.ConfigHistory = NewZoneSettingZarazConfigHistoryService(opts...)
	r.Publish = NewZoneSettingZarazPublishService(opts...)
	return
}
