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

func TestRecordingGetActiveRecordings(t *testing.T) {
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
	_, err := client.RealtimeKit.Recordings.GetActiveRecordings(
		context.TODO(),
		"app_id",
		"meeting_id",
		realtime_kit.RecordingGetActiveRecordingsParams{
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

func TestRecordingGetOneRecording(t *testing.T) {
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
	_, err := client.RealtimeKit.Recordings.GetOneRecording(
		context.TODO(),
		"app_id",
		"recording_id",
		realtime_kit.RecordingGetOneRecordingParams{
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

func TestRecordingGetRecordingsWithOptionalParams(t *testing.T) {
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
	_, err := client.RealtimeKit.Recordings.GetRecordings(
		context.TODO(),
		"app_id",
		realtime_kit.RecordingGetRecordingsParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			EndTime:   cloudflare.F(time.Now()),
			Expired:   cloudflare.F(true),
			MeetingID: cloudflare.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PageNo:    cloudflare.F(0.000000),
			PerPage:   cloudflare.F(0.000000),
			Search:    cloudflare.F("search"),
			SortBy:    cloudflare.F(realtime_kit.RecordingGetRecordingsParamsSortByInvokedTime),
			SortOrder: cloudflare.F(realtime_kit.RecordingGetRecordingsParamsSortOrderAsc),
			StartTime: cloudflare.F(time.Now()),
			Status:    cloudflare.F([]realtime_kit.RecordingGetRecordingsParamsStatus{realtime_kit.RecordingGetRecordingsParamsStatusInvoked}),
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

func TestRecordingPauseResumeStopRecording(t *testing.T) {
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
	_, err := client.RealtimeKit.Recordings.PauseResumeStopRecording(
		context.TODO(),
		"2a95132c15732412d22c1476fa83f27a",
		"recording_id",
		realtime_kit.RecordingPauseResumeStopRecordingParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Action:    cloudflare.F(realtime_kit.RecordingPauseResumeStopRecordingParamsActionStop),
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

func TestRecordingStartRecordingsWithOptionalParams(t *testing.T) {
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
	_, err := client.RealtimeKit.Recordings.StartRecordings(
		context.TODO(),
		"app_id",
		realtime_kit.RecordingStartRecordingsParams{
			AccountID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowMultipleRecordings: cloudflare.F(false),
			AudioConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsAudioConfig{
				Channel:    cloudflare.F(realtime_kit.RecordingStartRecordingsParamsAudioConfigChannelStereo),
				Codec:      cloudflare.F(realtime_kit.RecordingStartRecordingsParamsAudioConfigCodecAac),
				ExportFile: cloudflare.F(true),
			}),
			FileNamePrefix: cloudflare.F("string"),
			InteractiveConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsInteractiveConfig{
				Type: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsInteractiveConfigTypeId3),
			}),
			MaxSeconds: cloudflare.F(int64(60)),
			MeetingID:  cloudflare.F("97440c6a-140b-40a9-9499-b23fd7a3868a"),
			RealtimekitBucketConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsRealtimekitBucketConfig{
				Enabled: cloudflare.F(true),
			}),
			RtmpOutConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsRtmpOutConfig{
				RtmpURL: cloudflare.F("rtmp://a.rtmp.youtube.com/live2"),
			}),
			StorageConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsStorageConfig{
				Type:       cloudflare.F(realtime_kit.RecordingStartRecordingsParamsStorageConfigTypeAws),
				AccessKey:  cloudflare.F("access_key"),
				AuthMethod: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsStorageConfigAuthMethodKey),
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
			URL: cloudflare.F("https://example.com"),
			VideoConfig: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsVideoConfig{
				Codec:      cloudflare.F(realtime_kit.RecordingStartRecordingsParamsVideoConfigCodecH264),
				ExportFile: cloudflare.F(true),
				Height:     cloudflare.F(int64(720)),
				Watermark: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsVideoConfigWatermark{
					Position: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsVideoConfigWatermarkPositionLeftTop),
					Size: cloudflare.F(realtime_kit.RecordingStartRecordingsParamsVideoConfigWatermarkSize{
						Height: cloudflare.F(int64(1)),
						Width:  cloudflare.F(int64(1)),
					}),
					URL: cloudflare.F("http://example.com"),
				}),
				Width: cloudflare.F(int64(1280)),
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

func TestRecordingStartTrackRecordingWithOptionalParams(t *testing.T) {
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
	err := client.RealtimeKit.Recordings.StartTrackRecording(
		context.TODO(),
		"app_id",
		realtime_kit.RecordingStartTrackRecordingParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Layers: cloudflare.F(map[string]realtime_kit.RecordingStartTrackRecordingParamsLayers{
				"default": {
					FileNamePrefix: cloudflare.F("string"),
					Outputs: cloudflare.F([]realtime_kit.RecordingStartTrackRecordingParamsLayersOutput{{
						StorageConfig: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfig{
							Type:       cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAws),
							AccessKey:  cloudflare.F("access_key"),
							AuthMethod: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodKey),
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
						Type: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsTypeRealtimekitBucket),
					}}),
				},
				"default-video": {
					FileNamePrefix: cloudflare.F("string"),
					Outputs: cloudflare.F([]realtime_kit.RecordingStartTrackRecordingParamsLayersOutput{{
						StorageConfig: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfig{
							Type:       cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfigTypeAws),
							AccessKey:  cloudflare.F("access_key"),
							AuthMethod: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsStorageConfigAuthMethodKey),
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
						Type: cloudflare.F(realtime_kit.RecordingStartTrackRecordingParamsLayersOutputsTypeRealtimekitBucket),
					}}),
				},
			}),
			MeetingID:  cloudflare.F("string"),
			MaxSeconds: cloudflare.F(60.000000),
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
