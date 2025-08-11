// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestV2AccountOrderRequestGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestGetParams{
			AccountID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestListParams{
			Page:     dinariapisdkgo.Int(1),
			PageSize: dinariapisdkgo.Int(1),
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestNewLimitBuyWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.NewLimitBuy(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestNewLimitBuyParams{
			CreateLimitBuyOrderInput: dinariapisdkgo.CreateLimitBuyOrderInputParam{
				AssetQuantity:      0,
				LimitPrice:         0,
				StockID:            "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				RecipientAccountID: dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestNewLimitSellWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.NewLimitSell(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestNewLimitSellParams{
			CreateLimitSellOrderInput: dinariapisdkgo.CreateLimitSellOrderInputParam{
				AssetQuantity:       0,
				LimitPrice:          0,
				StockID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PaymentTokenAddress: dinariapisdkgo.String("payment_token_address"),
				RecipientAccountID:  dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestNewMarketBuyWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.NewMarketBuy(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestNewMarketBuyParams{
			CreateMarketBuyOrderInput: dinariapisdkgo.CreateMarketBuyOrderInputParam{
				PaymentAmount:      0,
				StockID:            "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				RecipientAccountID: dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestNewMarketSellWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.NewMarketSell(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestNewMarketSellParams{
			CreateMarketSellOrderInput: dinariapisdkgo.CreateMarketSellOrderInputParam{
				AssetQuantity:       0,
				StockID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PaymentTokenAddress: dinariapisdkgo.String("payment_token_address"),
				RecipientAccountID:  dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			},
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2AccountOrderRequestGetFeeQuoteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.V2.Accounts.OrderRequests.GetFeeQuote(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderRequestGetFeeQuoteParams{
			OrderSide:            dinariapisdkgo.OrderSideBuy,
			OrderType:            dinariapisdkgo.OrderTypeMarket,
			StockID:              "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AssetTokenQuantity:   dinariapisdkgo.Float(0),
			ChainID:              dinariapisdkgo.ChainEip155_1,
			LimitPrice:           dinariapisdkgo.Float(0),
			PaymentTokenAddress:  dinariapisdkgo.String("payment_token_address"),
			PaymentTokenQuantity: dinariapisdkgo.Float(0),
		},
	)
	if err != nil {
		var apierr *dinariapisdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
