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

func TestV2AccountOrderStockEip155GetFeeQuoteWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Accounts.Orders.Stocks.Eip155.GetFeeQuote(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderStockEip155GetFeeQuoteParams{
			ChainID:              dinariapisdkgo.ChainEip155_1,
			OrderSide:            dinariapisdkgo.OrderSideBuy,
			OrderTif:             dinariapisdkgo.OrderTifDay,
			OrderType:            dinariapisdkgo.OrderTypeMarket,
			PaymentToken:         "payment_token",
			AssetTokenQuantity:   dinariapisdkgo.Float(0),
			ClientOrderID:        dinariapisdkgo.String("client_order_id"),
			LimitPrice:           dinariapisdkgo.Float(0),
			PaymentTokenQuantity: dinariapisdkgo.Float(0),
			StockID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TokenID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestV2AccountOrderStockEip155PrepareOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.V2.Accounts.Orders.Stocks.Eip155.PrepareOrder(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdkgo.V2AccountOrderStockEip155PrepareOrderParams{
			ChainID:              dinariapisdkgo.ChainEip155_1,
			OrderSide:            dinariapisdkgo.OrderSideBuy,
			OrderTif:             dinariapisdkgo.OrderTifDay,
			OrderType:            dinariapisdkgo.OrderTypeMarket,
			PaymentToken:         "payment_token",
			AssetTokenQuantity:   dinariapisdkgo.Float(0),
			ClientOrderID:        dinariapisdkgo.String("client_order_id"),
			LimitPrice:           dinariapisdkgo.Float(0),
			PaymentTokenQuantity: dinariapisdkgo.Float(0),
			StockID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TokenID:              dinariapisdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
