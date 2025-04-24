// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestAPIV2AccountWalletExternalConnect(t *testing.T) {
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
	_, err := client.API.V2.Accounts.Wallet.External.Connect(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinari.APIV2AccountWalletExternalConnectParams{
			ChainID:       0,
			Nonce:         "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Signature:     "0xeaF12bD1DfFd",
			WalletAddress: "wallet_address",
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

func TestAPIV2AccountWalletExternalGetNonce(t *testing.T) {
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
	_, err := client.API.V2.Accounts.Wallet.External.GetNonce(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinari.APIV2AccountWalletExternalGetNonceParams{
			WalletAddress: "wallet_address",
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
