package client

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mjgrzybek/form3-interview-accountapi/client/internal/address"
	utils "github.com/mjgrzybek/form3-interview-accountapi/client/internal/utils"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/client"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/models"
	"github.com/stretchr/testify/assert"
)

var apiUrlForTests *url.URL

func init() {
	u, _ := url.Parse("http://blabla")
	apiUrlForTests = u
}

func TestRequestTimeout(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := utils.JoinPathUrl(*apiUrlForTests, "organisation", "accounts", FetchRequestsData.ID).String()

	httpmock.RegisterResponder("GET", endpoint,
		func(req *http.Request) (*http.Response, error) {
			<-time.After(10 * time.Second)
			return nil, nil
		},
	)

	t.Run("cancel request", func(t *testing.T) {
		client := client.NewClient(apiUrlForTests)

		ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
		defer cancelFunc()
		accountResponseData, err := client.AccountApi.Fetch(ctx, FetchRequestsData)

		assert.Nil(t, accountResponseData)
		assert.ErrorContains(t, err, "context deadline exceeded")
	})
}

func TestClientHandlingResponseCodes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	createUrl := func(accountId string) string {
		return utils.JoinPathUrl(*apiUrlForTests, "organisation", "accounts", accountId).String()
	}

	t.Run("status code 0 handling", func(t *testing.T) {
		const acccountId = "0_malformed"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(0, ``))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.EqualError(t, err, "0")
	})
	t.Run("status code 102 handling", func(t *testing.T) {
		const acccountId = "102_Processing"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(http.StatusProcessing, ``))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.EqualError(t, err, "102")
	})
	t.Run("status code 200 handling", func(t *testing.T) {
		const acccountId = "200_OK"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(http.StatusOK, `{}`))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.NoError(t, err)
	})
	t.Run("status code 301 handling", func(t *testing.T) {
		const acccountId = "301_Moved_Permanently"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(http.StatusMovedPermanently, ``))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.ErrorContains(t, err, "301 response missing Location header")
	})
	t.Run("status code 410 handling", func(t *testing.T) {
		const acccountId = "410_Gone"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(http.StatusGone, `{"error_message":"gone"}`))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.EqualError(t, err, "gone")
	})
	t.Run("status code 503 handling", func(t *testing.T) {
		const acccountId = "503_Service_Unavailable"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(http.StatusServiceUnavailable, ``))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.EqualError(t, err, "503")
	})
	t.Run("status code 1000 handling", func(t *testing.T) {
		const acccountId = "1000_malformed"
		httpmock.RegisterResponder("GET", createUrl(acccountId),
			httpmock.NewStringResponder(1000, ``))

		client := client.NewClient(apiUrlForTests)
		_, err := client.AccountApi.Fetch(context.TODO(), &models.AccountIdVersion{
			AccountId: models.AccountId{ID: acccountId},
			Version:   address.Of[int64](0),
		})
		assert.ErrorContains(t, err, "Unable to interpret httpResponse")
	})
}
