package client

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	internal "github.com/mjgrzybek/form3-interview-accountapi/client/internal/utils"
	"github.com/mjgrzybek/form3-interview-accountapi/client/pkg/client"
	"github.com/stretchr/testify/assert"
)

func TestRequestTimeout(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	apiUrl, _ := url.Parse("http://blabla")

	httpmock.RegisterResponder("GET", internal.JoinPathUrl(*apiUrl, "organisation", "accounts", FetchRequestsData.ID).String(),
		func(req *http.Request) (*http.Response, error) {
			<-time.After(10 * time.Second)
			return nil, nil
		},
	)

	t.Run("create unique entry", func(t *testing.T) {
		client := client.NewClient(apiUrl)

		ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
		defer cancelFunc()
		accountResponseData, err := client.AccountApi.Fetch(ctx, FetchRequestsData)

		assert.Nil(t, accountResponseData)
		assert.ErrorContains(t, err, "context deadline exceeded")
	})
}
