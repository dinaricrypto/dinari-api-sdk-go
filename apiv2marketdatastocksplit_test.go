// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/dinari-go"
	"github.com/stainless-sdks/dinari-go/internal/testutil"
	"github.com/stainless-sdks/dinari-go/option"
)

func TestAPIV2MarketDataStockSplitGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinari.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.V2.MarketData.Stocks.Splits.Get(
		context.TODO(),
		"stock_id",
		dinari.APIV2MarketDataStockSplitGetParams{
			Page:     dinari.Int(1),
			PageSize: dinari.Int(1),
		},
	)
	if err != nil {
		var apierr *dinari.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV2MarketDataStockSplitListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinari.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.V2.MarketData.Stocks.Splits.List(context.TODO(), dinari.APIV2MarketDataStockSplitListParams{
		Page:     dinari.Int(1),
		PageSize: dinari.Int(1),
	})
	if err != nil {
		var apierr *dinari.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
