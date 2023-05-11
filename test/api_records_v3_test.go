/*
Confluent Cloud APIs

Testing RecordsV3ApiService

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

func Test_openapi_RecordsV3ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RecordsV3ApiService ProduceRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterId string
		var topicName string

		resp, httpRes, err := apiClient.RecordsV3Api.ProduceRecord(context.Background(), clusterId, topicName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
