// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/realtime_kit"
)

func TestMeetingNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.New(
		context.TODO(),
		"app_id",
		realtime_kit.MeetingNewParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AIConfig: cloudflare.F(realtime_kit.MeetingNewParamsAIConfig{
				Summarization: cloudflare.F(realtime_kit.MeetingNewParamsAIConfigSummarization{
					SummaryType: cloudflare.F(realtime_kit.MeetingNewParamsAIConfigSummarizationSummaryTypeGeneral),
					TextFormat:  cloudflare.F(realtime_kit.MeetingNewParamsAIConfigSummarizationTextFormatPlainText),
					WordLimit:   cloudflare.F(int64(150)),
				}),
				Transcription: cloudflare.F(realtime_kit.MeetingNewParamsAIConfigTranscription{
					Keywords:        cloudflare.F([]string{"string"}),
					Language:        cloudflare.F(realtime_kit.MeetingNewParamsAIConfigTranscriptionLanguageEnUs),
					ProfanityFilter: cloudflare.F(true),
				}),
			}),
			LiveStreamOnStart: cloudflare.F(true),
			PersistChat:       cloudflare.F(true),
			RecordOnStart:     cloudflare.F(true),
			RecordingConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfig{
				AudioConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigAudioConfig{
					Channel:    cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigAudioConfigChannelMono),
					Codec:      cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigAudioConfigCodecMP3),
					ExportFile: cloudflare.F(true),
				}),
				FileNamePrefix: cloudflare.F("file_name_prefix"),
				LiveStreamingConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigLiveStreamingConfig{
					RtmpURL: cloudflare.F("rtmp://a.rtmp.youtube.com/live2"),
				}),
				MaxSeconds: cloudflare.F(60.000000),
				RealtimekitBucketConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigRealtimekitBucketConfig{
					Enabled: cloudflare.F(true),
				}),
				StorageConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigStorageConfig{
					Type:       cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigStorageConfigTypeAws),
					AccessKey:  cloudflare.F("access_key"),
					AuthMethod: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigStorageConfigAuthMethodKey),
					Bucket:     cloudflare.F("bucket"),
					Host:       cloudflare.F("host"),
					Password:   cloudflare.F("password"),
					Path:       cloudflare.F("path"),
					Port:       cloudflare.F(0.000000),
					PrivateKey: cloudflare.F("private_key"),
					Region:     cloudflare.F("us-east-1"),
					Secret:     cloudflare.F("secret"),
					Username:   cloudflare.F("username"),
				}),
				VideoConfig: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigVideoConfig{
					Codec:      cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigVideoConfigCodecH264),
					ExportFile: cloudflare.F(true),
					Height:     cloudflare.F(int64(720)),
					Watermark: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigVideoConfigWatermark{
						Position: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigVideoConfigWatermarkPositionLeftTop),
						Size: cloudflare.F(realtime_kit.MeetingNewParamsRecordingConfigVideoConfigWatermarkSize{
							Height: cloudflare.F(int64(1)),
							Width:  cloudflare.F(int64(1)),
						}),
						URL: cloudflare.F("https://example.com"),
					}),
					Width: cloudflare.F(int64(1280)),
				}),
			}),
			SessionKeepAliveTimeInSecs: cloudflare.F(60.000000),
			SummarizeOnEnd:             cloudflare.F(true),
			Title:                      cloudflare.F("title"),
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

func TestMeetingAddParticipantWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.AddParticipant(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.MeetingAddParticipantParams{
			AccountID:           cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CustomParticipantID: cloudflare.F("custom_participant_id"),
			PresetName:          cloudflare.F("preset_name"),
			Name:                cloudflare.F("Mary Sue"),
			Picture:             cloudflare.F("https://i.imgur.com/test.jpg"),
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

func TestMeetingDeleteMeetingParticipant(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.DeleteMeetingParticipant(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"participant_id",
		realtime_kit.MeetingDeleteMeetingParticipantParams{
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

func TestMeetingEditParticipantWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.EditParticipant(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"participant_id",
		realtime_kit.MeetingEditParticipantParams{
			AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:       cloudflare.F("Jane Doe"),
			Picture:    cloudflare.F("https://example.com"),
			PresetName: cloudflare.F("preset_name"),
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

func TestMeetingGetWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.Get(
		context.TODO(),
		"app_id",
		realtime_kit.MeetingGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndTime:   cloudflare.F(time.Now()),
			PageNo:    cloudflare.F(0.000000),
			PerPage:   cloudflare.F(0.000000),
			Search:    cloudflare.F("search"),
			StartTime: cloudflare.F(time.Now()),
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

func TestMeetingGetMeetingByIDWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.GetMeetingByID(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.MeetingGetMeetingByIDParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Name:      cloudflare.F("name"),
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

func TestMeetingGetMeetingParticipant(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.GetMeetingParticipant(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"participant_id",
		realtime_kit.MeetingGetMeetingParticipantParams{
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

func TestMeetingGetMeetingParticipantsWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.GetMeetingParticipants(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.MeetingGetMeetingParticipantsParams{
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

func TestMeetingRefreshParticipantToken(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.RefreshParticipantToken(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"participant_id",
		realtime_kit.MeetingRefreshParticipantTokenParams{
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

func TestMeetingReplaceMeetingByIDWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.ReplaceMeetingByID(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.MeetingReplaceMeetingByIDParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AIConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfig{
				Summarization: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfigSummarization{
					SummaryType: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral),
					TextFormat:  cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfigSummarizationTextFormatPlainText),
					WordLimit:   cloudflare.F(int64(150)),
				}),
				Transcription: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfigTranscription{
					Keywords:        cloudflare.F([]string{"string"}),
					Language:        cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsAIConfigTranscriptionLanguageEnUs),
					ProfanityFilter: cloudflare.F(true),
				}),
			}),
			LiveStreamOnStart: cloudflare.F(true),
			PersistChat:       cloudflare.F(true),
			RecordOnStart:     cloudflare.F(true),
			RecordingConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfig{
				AudioConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfig{
					Channel:    cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigChannelMono),
					Codec:      cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigAudioConfigCodecMP3),
					ExportFile: cloudflare.F(true),
				}),
				FileNamePrefix: cloudflare.F("file_name_prefix"),
				LiveStreamingConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigLiveStreamingConfig{
					RtmpURL: cloudflare.F("rtmp://a.rtmp.youtube.com/live2"),
				}),
				MaxSeconds: cloudflare.F(60.000000),
				RealtimekitBucketConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigRealtimekitBucketConfig{
					Enabled: cloudflare.F(true),
				}),
				StorageConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfig{
					Type:       cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigTypeAws),
					AccessKey:  cloudflare.F("access_key"),
					AuthMethod: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigStorageConfigAuthMethodKey),
					Bucket:     cloudflare.F("bucket"),
					Host:       cloudflare.F("host"),
					Password:   cloudflare.F("password"),
					Path:       cloudflare.F("path"),
					Port:       cloudflare.F(0.000000),
					PrivateKey: cloudflare.F("private_key"),
					Region:     cloudflare.F("us-east-1"),
					Secret:     cloudflare.F("secret"),
					Username:   cloudflare.F("username"),
				}),
				VideoConfig: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfig{
					Codec:      cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigCodecH264),
					ExportFile: cloudflare.F(true),
					Height:     cloudflare.F(int64(720)),
					Watermark: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermark{
						Position: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkPositionLeftTop),
						Size: cloudflare.F(realtime_kit.MeetingReplaceMeetingByIDParamsRecordingConfigVideoConfigWatermarkSize{
							Height: cloudflare.F(int64(1)),
							Width:  cloudflare.F(int64(1)),
						}),
						URL: cloudflare.F("https://example.com"),
					}),
					Width: cloudflare.F(int64(1280)),
				}),
			}),
			SessionKeepAliveTimeInSecs: cloudflare.F(60.000000),
			SummarizeOnEnd:             cloudflare.F(true),
			Title:                      cloudflare.F("title"),
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

func TestMeetingUpdateMeetingByIDWithOptionalParams(t *testing.T) {
	t.Skip("TODO: HTTP 401 from prism, support api tokens")
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
	_, err := client.RealtimeKit.Meetings.UpdateMeetingByID(
		context.TODO(),
		"app_id",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		realtime_kit.MeetingUpdateMeetingByIDParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AIConfig: cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfig{
				Summarization: cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfigSummarization{
					SummaryType: cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfigSummarizationSummaryTypeGeneral),
					TextFormat:  cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfigSummarizationTextFormatPlainText),
					WordLimit:   cloudflare.F(int64(150)),
				}),
				Transcription: cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfigTranscription{
					Keywords:        cloudflare.F([]string{"string"}),
					Language:        cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsAIConfigTranscriptionLanguageEnUs),
					ProfanityFilter: cloudflare.F(true),
				}),
			}),
			LiveStreamOnStart:          cloudflare.F(true),
			PersistChat:                cloudflare.F(true),
			RecordOnStart:              cloudflare.F(true),
			SessionKeepAliveTimeInSecs: cloudflare.F(60.000000),
			Status:                     cloudflare.F(realtime_kit.MeetingUpdateMeetingByIDParamsStatusInactive),
			SummarizeOnEnd:             cloudflare.F(true),
			Title:                      cloudflare.F("title"),
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
