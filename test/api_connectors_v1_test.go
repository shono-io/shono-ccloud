/*
Confluent Cloud APIs

Testing ConnectorsV1ApiService

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

func Test_openapi_ConnectorsV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectorsV1ApiService CreateConnectv1Connector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.CreateConnectv1Connector(context.Background(), environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService CreateOrUpdateConnectv1ConnectorConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorName string
		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.CreateOrUpdateConnectv1ConnectorConfig(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService DeleteConnectv1Connector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorName string
		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.DeleteConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService GetConnectv1ConnectorConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorName string
		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.GetConnectv1ConnectorConfig(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService ListConnectv1Connectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.ListConnectv1Connectors(context.Background(), environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService ListConnectv1ConnectorsWithExpansions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.ListConnectv1ConnectorsWithExpansions(context.Background(), environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsV1ApiService ReadConnectv1Connector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorName string
		var environmentId string
		var kafkaClusterId string

		resp, httpRes, err := apiClient.ConnectorsV1Api.ReadConnectv1Connector(context.Background(), connectorName, environmentId, kafkaClusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
