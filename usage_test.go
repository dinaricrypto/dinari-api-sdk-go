// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/dinari-go"
	"github.com/stainless-sdks/dinari-go/internal/testutil"
	"github.com/stainless-sdks/dinari-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinari.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyID("My API Key ID"),
		option.WithAPISecretKey("My API Secret Key"),
	)
	stocks, err := client.V2.MarketData.Stocks.List(context.TODO(), dinari.V2MarketDataStockListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", stocks)
}
