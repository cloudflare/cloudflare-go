// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/realtime_kit"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

func TestPresetNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Presets.New(
		context.TODO(),
		"app_id",
		realtime_kit.PresetNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Config: cloudflare.F(realtime_kit.PresetNewParamsConfig{
				MaxScreenshareCount: cloudflare.F(int64(0)),
				MaxVideoStreams: cloudflare.F(realtime_kit.PresetNewParamsConfigMaxVideoStreams{
					Desktop: cloudflare.F(int64(0)),
					Mobile:  cloudflare.F(int64(0)),
				}),
				Media: cloudflare.F(realtime_kit.PresetNewParamsConfigMedia{
					Screenshare: cloudflare.F(realtime_kit.PresetNewParamsConfigMediaScreenshare{
						FrameRate: cloudflare.F(int64(0)),
						Quality:   cloudflare.F(realtime_kit.PresetNewParamsConfigMediaScreenshareQualityHD),
					}),
					Video: cloudflare.F(realtime_kit.PresetNewParamsConfigMediaVideo{
						FrameRate: cloudflare.F(int64(30)),
						Quality:   cloudflare.F(realtime_kit.PresetNewParamsConfigMediaVideoQualityHD),
					}),
					Audio: cloudflare.F(realtime_kit.PresetNewParamsConfigMediaAudio{
						EnableHighBitrate: cloudflare.F(true),
						EnableStereo:      cloudflare.F(true),
					}),
				}),
				ViewType: cloudflare.F(realtime_kit.PresetNewParamsConfigViewTypeGroupCall),
			}),
			Name: cloudflare.F("name"),
			UI: cloudflare.F(realtime_kit.PresetNewParamsUI{
				DesignTokens: cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokens{
					BorderRadius: cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensBorderRadiusRounded),
					BorderWidth:  cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensBorderWidthThin),
					Colors: cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensColors{
						Background: cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensColorsBackground{
							Number600:  cloudflare.F("600"),
							Number700:  cloudflare.F("700"),
							Number800:  cloudflare.F("800"),
							Number900:  cloudflare.F("900"),
							Number1000: cloudflare.F("1000"),
						}),
						Brand: cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensColorsBrand{
							Number300: cloudflare.F("300"),
							Number400: cloudflare.F("400"),
							Number500: cloudflare.F("500"),
							Number600: cloudflare.F("600"),
							Number700: cloudflare.F("700"),
						}),
						Danger:      cloudflare.F("danger"),
						Success:     cloudflare.F("success"),
						Text:        cloudflare.F("text"),
						TextOnBrand: cloudflare.F("text_on_brand"),
						VideoBg:     cloudflare.F("video_bg"),
						Warning:     cloudflare.F("warning"),
					}),
					Logo:        cloudflare.F("logo"),
					SpacingBase: cloudflare.F(0.000000),
					Theme:       cloudflare.F(realtime_kit.PresetNewParamsUIDesignTokensThemeDark),
				}),
				ConfigDiff: cloudflare.F[any](map[string]interface{}{}),
			}),
			Permissions: cloudflare.F(realtime_kit.PresetNewParamsPermissions{
				AcceptWaitingRequests:           cloudflare.F(true),
				CanAcceptProductionRequests:     cloudflare.F(true),
				CanChangeParticipantPermissions: cloudflare.F(true),
				CanEditDisplayName:              cloudflare.F(true),
				CanLivestream:                   cloudflare.F(true),
				CanRecord:                       cloudflare.F(true),
				CanSpotlight:                    cloudflare.F(true),
				Chat: cloudflare.F(realtime_kit.PresetNewParamsPermissionsChat{
					Private: cloudflare.F(realtime_kit.PresetNewParamsPermissionsChatPrivate{
						CanReceive: cloudflare.F(true),
						CanSend:    cloudflare.F(true),
						Files:      cloudflare.F(true),
						Text:       cloudflare.F(true),
					}),
					Public: cloudflare.F(realtime_kit.PresetNewParamsPermissionsChatPublic{
						CanSend: cloudflare.F(true),
						Files:   cloudflare.F(true),
						Text:    cloudflare.F(true),
					}),
				}),
				ConnectedMeetings: cloudflare.F(realtime_kit.PresetNewParamsPermissionsConnectedMeetings{
					CanAlterConnectedMeetings:  cloudflare.F(true),
					CanSwitchConnectedMeetings: cloudflare.F(true),
					CanSwitchToParentMeeting:   cloudflare.F(true),
				}),
				DisableParticipantAudio:         cloudflare.F(true),
				DisableParticipantScreensharing: cloudflare.F(true),
				DisableParticipantVideo:         cloudflare.F(true),
				HiddenParticipant:               cloudflare.F(true),
				KickParticipant:                 cloudflare.F(true),
				Media: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMedia{
					Audio: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaAudio{
						CanProduce: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaAudioCanProduceAllowed),
					}),
					Screenshare: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaScreenshare{
						CanProduce: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaScreenshareCanProduceAllowed),
					}),
					Video: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaVideo{
						CanProduce: cloudflare.F(realtime_kit.PresetNewParamsPermissionsMediaVideoCanProduceAllowed),
					}),
				}),
				PinParticipant: cloudflare.F(true),
				Plugins: cloudflare.F(realtime_kit.PresetNewParamsPermissionsPlugins{
					CanClose:      cloudflare.F(true),
					CanEditConfig: cloudflare.F(true),
					CanStart:      cloudflare.F(true),
					Config:        cloudflare.F[realtime_kit.PresetNewParamsPermissionsPluginsConfigUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
				}),
				Polls: cloudflare.F(realtime_kit.PresetNewParamsPermissionsPolls{
					CanCreate: cloudflare.F(true),
					CanView:   cloudflare.F(true),
					CanVote:   cloudflare.F(true),
				}),
				RecorderType:        cloudflare.F(realtime_kit.PresetNewParamsPermissionsRecorderTypeRecorder),
				ShowParticipantList: cloudflare.F(true),
				WaitingRoomType:     cloudflare.F(realtime_kit.PresetNewParamsPermissionsWaitingRoomTypeSkip),
				IsRecorder:          cloudflare.F(true),
			}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPresetUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Presets.Update(
		context.TODO(),
		"app_id",
		"preset_id",
		realtime_kit.PresetUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Config: cloudflare.F(realtime_kit.PresetUpdateParamsConfig{
				MaxScreenshareCount: cloudflare.F(int64(0)),
				MaxVideoStreams: cloudflare.F(realtime_kit.PresetUpdateParamsConfigMaxVideoStreams{
					Desktop: cloudflare.F(int64(0)),
					Mobile:  cloudflare.F(int64(0)),
				}),
				Media: cloudflare.F(realtime_kit.PresetUpdateParamsConfigMedia{
					Screenshare: cloudflare.F(realtime_kit.PresetUpdateParamsConfigMediaScreenshare{
						FrameRate: cloudflare.F(int64(0)),
						Quality:   cloudflare.F(realtime_kit.PresetUpdateParamsConfigMediaScreenshareQualityHD),
					}),
					Video: cloudflare.F(realtime_kit.PresetUpdateParamsConfigMediaVideo{
						FrameRate: cloudflare.F(int64(30)),
						Quality:   cloudflare.F(realtime_kit.PresetUpdateParamsConfigMediaVideoQualityHD),
					}),
				}),
				ViewType: cloudflare.F(realtime_kit.PresetUpdateParamsConfigViewTypeGroupCall),
			}),
			Name: cloudflare.F("name"),
			Permissions: cloudflare.F(realtime_kit.PresetUpdateParamsPermissions{
				AcceptWaitingRequests:           cloudflare.F(true),
				CanAcceptProductionRequests:     cloudflare.F(true),
				CanChangeParticipantPermissions: cloudflare.F(true),
				CanEditDisplayName:              cloudflare.F(true),
				CanLivestream:                   cloudflare.F(true),
				CanRecord:                       cloudflare.F(true),
				CanSpotlight:                    cloudflare.F(true),
				Chat: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsChat{
					Private: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsChatPrivate{
						CanReceive: cloudflare.F(true),
						CanSend:    cloudflare.F(true),
						Files:      cloudflare.F(true),
						Text:       cloudflare.F(true),
					}),
					Public: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsChatPublic{
						CanSend: cloudflare.F(true),
						Files:   cloudflare.F(true),
						Text:    cloudflare.F(true),
					}),
				}),
				ConnectedMeetings: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsConnectedMeetings{
					CanAlterConnectedMeetings:  cloudflare.F(true),
					CanSwitchConnectedMeetings: cloudflare.F(true),
					CanSwitchToParentMeeting:   cloudflare.F(true),
				}),
				DisableParticipantAudio:         cloudflare.F(true),
				DisableParticipantScreensharing: cloudflare.F(true),
				DisableParticipantVideo:         cloudflare.F(true),
				HiddenParticipant:               cloudflare.F(true),
				IsRecorder:                      cloudflare.F(true),
				KickParticipant:                 cloudflare.F(true),
				Media: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMedia{
					Audio: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaAudio{
						CanProduce: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaAudioCanProduceAllowed),
					}),
					Screenshare: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaScreenshare{
						CanProduce: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaScreenshareCanProduceAllowed),
					}),
					Video: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaVideo{
						CanProduce: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsMediaVideoCanProduceAllowed),
					}),
				}),
				PinParticipant: cloudflare.F(true),
				Plugins: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsPlugins{
					CanClose:      cloudflare.F(true),
					CanEditConfig: cloudflare.F(true),
					CanStart:      cloudflare.F(true),
					Config:        cloudflare.F[realtime_kit.PresetUpdateParamsPermissionsPluginsConfigUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
				}),
				Polls: cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsPolls{
					CanCreate: cloudflare.F(true),
					CanView:   cloudflare.F(true),
					CanVote:   cloudflare.F(true),
				}),
				RecorderType:        cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsRecorderTypeRecorder),
				ShowParticipantList: cloudflare.F(true),
				WaitingRoomType:     cloudflare.F(realtime_kit.PresetUpdateParamsPermissionsWaitingRoomTypeSkip),
			}),
			UI: cloudflare.F(realtime_kit.PresetUpdateParamsUI{
				ConfigDiff: cloudflare.F[any](map[string]interface{}{}),
				DesignTokens: cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokens{
					BorderRadius: cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensBorderRadiusRounded),
					BorderWidth:  cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensBorderWidthThin),
					Colors: cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensColors{
						Background: cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensColorsBackground{
							Number600:  cloudflare.F("600"),
							Number700:  cloudflare.F("700"),
							Number800:  cloudflare.F("800"),
							Number900:  cloudflare.F("900"),
							Number1000: cloudflare.F("1000"),
						}),
						Brand: cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensColorsBrand{
							Number300: cloudflare.F("300"),
							Number400: cloudflare.F("400"),
							Number500: cloudflare.F("500"),
							Number600: cloudflare.F("600"),
							Number700: cloudflare.F("700"),
						}),
						Danger:      cloudflare.F("danger"),
						Success:     cloudflare.F("success"),
						Text:        cloudflare.F("text"),
						TextOnBrand: cloudflare.F("text_on_brand"),
						VideoBg:     cloudflare.F("video_bg"),
						Warning:     cloudflare.F("warning"),
					}),
					Logo:        cloudflare.F("logo"),
					SpacingBase: cloudflare.F(0.000000),
					Theme:       cloudflare.F(realtime_kit.PresetUpdateParamsUIDesignTokensThemeDark),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPresetDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Presets.Delete(
		context.TODO(),
		"app_id",
		"preset_id",
		realtime_kit.PresetDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPresetGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Presets.Get(
		context.TODO(),
		"app_id",
		realtime_kit.PresetGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			PageNo:    cloudflare.F(0.000000),
			PerPage:   cloudflare.F(0.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPresetGetPresetByID(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.RealtimeKit.Presets.GetPresetByID(
		context.TODO(),
		"app_id",
		"preset_id",
		realtime_kit.PresetGetPresetByIDParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
