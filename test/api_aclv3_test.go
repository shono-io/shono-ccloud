/*
Confluent Cloud APIs

Testing ACLV3ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/shono-io/shono-ccloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_ACLV3ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ACLV3ApiService BatchCreateKafkaAcls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		httpRes, err := apiClient.ACLV3Api.BatchCreateKafkaAcls(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ACLV3ApiService CreateKafkaAcls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		httpRes, err := apiClient.ACLV3Api.CreateKafkaAcls(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ACLV3ApiService DeleteKafkaAcls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ACLV3Api.DeleteKafkaAcls(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ACLV3ApiService GetKafkaAcls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ACLV3Api.GetKafkaAcls(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
