// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/testutil"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

func TestAPIV2EntityKYCGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
	)
	_, err := client.API.V2.Entities.KYC.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *dinariapisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV2EntityKYCSubmitWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
	)
	_, err := client.API.V2.Entities.KYC.Submit(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdk.APIV2EntityKYCSubmitParams{
			Data: dinariapisdk.KYCDataParam{
				CountryCode:        "SG",
				LastName:           "Doe",
				AddressCity:        dinariapisdk.String("San Francisco"),
				AddressPostalCode:  dinariapisdk.String("94111"),
				AddressStreet1:     dinariapisdk.String("123 Main St."),
				AddressStreet2:     dinariapisdk.String("Apt. 123"),
				AddressSubdivision: dinariapisdk.String("California"),
				BirthDate:          dinariapisdk.Time(time.Now()),
				Email:              dinariapisdk.String("johndoe@website.com"),
				FirstName:          dinariapisdk.String("John"),
				MiddleName:         dinariapisdk.String("x"),
				TaxIDNumber:        dinariapisdk.String("12-3456789"),
			},
			ProviderName: "x",
		},
	)
	if err != nil {
		var apierr *dinariapisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIV2EntityKYCUploadDocument(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dinariapisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
	)
	_, err := client.API.V2.Entities.KYC.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		dinariapisdk.APIV2EntityKYCUploadDocumentParams{
			EntityID:     "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			DocumentType: dinariapisdk.KYCDocumentTypeGovernmentID,
			File:         io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		},
	)
	if err != nil {
		var apierr *dinariapisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
