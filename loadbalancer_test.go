// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.New(context.TODO(), cloudflare.LoadBalancerNewParams{
		ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
		DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
		FallbackPool: cloudflare.F[any](map[string]interface{}{}),
		Name:         cloudflare.F("www.example.com"),
		AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerNewParamsAdaptiveRouting{
			FailoverAcrossPools: cloudflare.F(true),
		}),
		CountryPools: cloudflare.F[any](map[string]interface{}{
			"GB": map[string]interface{}{
				"0": "abd90f38ced07c2e2f4df50b1f61d4194",
			},
			"US": map[string]interface{}{
				"0": "de90f38ced07c2e2f4df50b1f61d4194",
				"1": "00920f38ce07c2e2f4df50b1f61d4194",
			},
		}),
		Description: cloudflare.F("Load Balancer for www.example.com"),
		LocationStrategy: cloudflare.F(cloudflare.LoadBalancerNewParamsLocationStrategy{
			Mode:      cloudflare.F(cloudflare.LoadBalancerNewParamsLocationStrategyModeResolverIP),
			PreferEcs: cloudflare.F(cloudflare.LoadBalancerNewParamsLocationStrategyPreferEcsAlways),
		}),
		PopPools: cloudflare.F[any](map[string]interface{}{
			"LAX": map[string]interface{}{
				"0": "de90f38ced07c2e2f4df50b1f61d4194",
				"1": "9290f38c5d07c2e2f4df57b1f61d4196",
			},
			"LHR": map[string]interface{}{
				"0": "abd90f38ced07c2e2f4df50b1f61d4194",
				"1": "f9138c5d07c2e2f4df57b1f61d4196",
			},
			"SJC": map[string]interface{}{
				"0": "00920f38ce07c2e2f4df50b1f61d4194",
			},
		}),
		Proxied: cloudflare.F(true),
		RandomSteering: cloudflare.F(cloudflare.LoadBalancerNewParamsRandomSteering{
			DefaultWeight: cloudflare.F(0.200000),
			PoolWeights: cloudflare.F[any](map[string]interface{}{
				"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
				"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
			}),
		}),
		RegionPools: cloudflare.F[any](map[string]interface{}{
			"ENAM": map[string]interface{}{
				"0": "00920f38ce07c2e2f4df50b1f61d4194",
			},
			"WNAM": map[string]interface{}{
				"0": "de90f38ced07c2e2f4df50b1f61d4194",
				"1": "9290f38c5d07c2e2f4df57b1f61d4196",
			},
		}),
		Rules: cloudflare.F([]cloudflare.LoadBalancerNewParamsRule{{
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F[any](map[string]interface{}{
					"GB": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					},
					"US": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F[any](map[string]interface{}{}),
				LocationStrategy: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
				}),
				PopPools: cloudflare.F[any](map[string]interface{}{
					"LAX": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
					"LHR": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						"1": "f9138c5d07c2e2f4df57b1f61d4196",
					},
					"SJC": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				RandomSteering: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesRandomSteering{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F[any](map[string]interface{}{
						"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
						"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
					}),
				}),
				RegionPools: cloudflare.F[any](map[string]interface{}{
					"ENAM": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
					"WNAM": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
				}),
				SessionAffinity: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F[any](map[string]interface{}{
					"GB": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					},
					"US": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F[any](map[string]interface{}{}),
				LocationStrategy: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
				}),
				PopPools: cloudflare.F[any](map[string]interface{}{
					"LAX": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
					"LHR": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						"1": "f9138c5d07c2e2f4df57b1f61d4196",
					},
					"SJC": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				RandomSteering: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesRandomSteering{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F[any](map[string]interface{}{
						"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
						"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
					}),
				}),
				RegionPools: cloudflare.F[any](map[string]interface{}{
					"ENAM": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
					"WNAM": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
				}),
				SessionAffinity: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F[any](map[string]interface{}{
					"GB": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					},
					"US": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F[any](map[string]interface{}{}),
				LocationStrategy: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
				}),
				PopPools: cloudflare.F[any](map[string]interface{}{
					"LAX": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
					"LHR": map[string]interface{}{
						"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						"1": "f9138c5d07c2e2f4df57b1f61d4196",
					},
					"SJC": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
				}),
				RandomSteering: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesRandomSteering{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F[any](map[string]interface{}{
						"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
						"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
					}),
				}),
				RegionPools: cloudflare.F[any](map[string]interface{}{
					"ENAM": map[string]interface{}{
						"0": "00920f38ce07c2e2f4df50b1f61d4194",
					},
					"WNAM": map[string]interface{}{
						"0": "de90f38ced07c2e2f4df50b1f61d4194",
						"1": "9290f38c5d07c2e2f4df57b1f61d4196",
					},
				}),
				SessionAffinity: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}}),
		SessionAffinity: cloudflare.F(cloudflare.LoadBalancerNewParamsSessionAffinityCookie),
		SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerNewParamsSessionAffinityAttributes{
			DrainDuration:        cloudflare.F(100.000000),
			Headers:              cloudflare.F([]string{"x"}),
			RequireAllHeaders:    cloudflare.F(true),
			Samesite:             cloudflare.F(cloudflare.LoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto),
			Secure:               cloudflare.F(cloudflare.LoadBalancerNewParamsSessionAffinityAttributesSecureAuto),
			ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
		}),
		SessionAffinityTTL: cloudflare.F(1800.000000),
		SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerNewParamsSteeringPolicyDynamicLatency),
		TTL:                cloudflare.F(30.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.LoadBalancerUpdateParams{
			ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
			DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerUpdateParamsAdaptiveRouting{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F[any](map[string]interface{}{
				"GB": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
				},
				"US": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Description: cloudflare.F("Load Balancer for www.example.com"),
			Enabled:     cloudflare.F(true),
			LocationStrategy: cloudflare.F(cloudflare.LoadBalancerUpdateParamsLocationStrategy{
				Mode:      cloudflare.F(cloudflare.LoadBalancerUpdateParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(cloudflare.LoadBalancerUpdateParamsLocationStrategyPreferEcsAlways),
			}),
			PopPools: cloudflare.F[any](map[string]interface{}{
				"LAX": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
				"LHR": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					"1": "f9138c5d07c2e2f4df57b1f61d4196",
				},
				"SJC": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRandomSteering{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F[any](map[string]interface{}{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F[any](map[string]interface{}{
				"ENAM": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
				"WNAM": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
			}),
			Rules: cloudflare.F([]cloudflare.LoadBalancerUpdateParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(cloudflare.LoadBalancerUpdateParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerUpdateParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(cloudflare.LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(cloudflare.LoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerUpdateParamsSteeringPolicyDynamicLatency),
			TTL:                cloudflare.F(30.000000),
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

func TestLoadBalancerList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.List(context.TODO(), cloudflare.LoadBalancerListParams{
		ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.LoadBalancerDeleteParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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

func TestLoadBalancerEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.LoadBalancerEditParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerEditParamsAdaptiveRouting{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F[any](map[string]interface{}{
				"GB": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
				},
				"US": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			Description:  cloudflare.F("Load Balancer for www.example.com"),
			Enabled:      cloudflare.F(true),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			LocationStrategy: cloudflare.F(cloudflare.LoadBalancerEditParamsLocationStrategy{
				Mode:      cloudflare.F(cloudflare.LoadBalancerEditParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(cloudflare.LoadBalancerEditParamsLocationStrategyPreferEcsAlways),
			}),
			Name: cloudflare.F("www.example.com"),
			PopPools: cloudflare.F[any](map[string]interface{}{
				"LAX": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
				"LHR": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					"1": "f9138c5d07c2e2f4df57b1f61d4196",
				},
				"SJC": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(cloudflare.LoadBalancerEditParamsRandomSteering{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F[any](map[string]interface{}{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F[any](map[string]interface{}{
				"ENAM": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
				"WNAM": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
			}),
			Rules: cloudflare.F([]cloudflare.LoadBalancerEditParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(cloudflare.LoadBalancerEditParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(cloudflare.LoadBalancerEditParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(cloudflare.LoadBalancerEditParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(cloudflare.LoadBalancerEditParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(cloudflare.LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(cloudflare.LoadBalancerEditParamsSteeringPolicyDynamicLatency),
			TTL:                cloudflare.F(30.000000),
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

func TestLoadBalancerGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.LoadBalancerGetParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
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
