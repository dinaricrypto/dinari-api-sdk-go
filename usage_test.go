// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKeyID("My API Key ID"),
		option.WithAPISecretKey("My API Secret Key"),
	)
	stocks, err := client.V2.MarketData.Stocks.List(context.TODO(), dinariapisdkgo.V2MarketDataStockListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", stocks)
}
