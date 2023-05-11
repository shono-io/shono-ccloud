/*
Confluent Cloud APIs

Testing JwksIamV2ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_JwksIamV2ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JwksIamV2ApiService RefreshIamV2JsonWebKeySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.JwksIamV2Api.RefreshIamV2JsonWebKeySet(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}