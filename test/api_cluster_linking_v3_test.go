/*
Confluent Cloud APIs

Testing ClusterLinkingV3ApiService

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

func Test_openapi_ClusterLinkingV3ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterLinkingV3ApiService CreateKafkaLink", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		httpRes, err := apiClient.ClusterLinkingV3Api.CreateKafkaLink(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService CreateKafkaMirrorTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		httpRes, err := apiClient.ClusterLinkingV3Api.CreateKafkaMirrorTopic(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService DeleteKafkaLink", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		httpRes, err := apiClient.ClusterLinkingV3Api.DeleteKafkaLink(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService DeleteKafkaLinkConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string
		var configName string

		httpRes, err := apiClient.ClusterLinkingV3Api.DeleteKafkaLinkConfig(context.Background(), clusterId, linkName, configName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService GetKafkaLink", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.GetKafkaLink(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService GetKafkaLinkConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string
		var configName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.GetKafkaLinkConfigs(context.Background(), clusterId, linkName, configName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService ListKafkaLinkConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.ListKafkaLinkConfigs(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService ListKafkaLinks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.ListKafkaLinks(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService ListKafkaMirrorTopics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.ListKafkaMirrorTopics(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService ListKafkaMirrorTopicsUnderLink", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.ListKafkaMirrorTopicsUnderLink(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService ReadKafkaMirrorTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string
		var mirrorTopicName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.ReadKafkaMirrorTopic(context.Background(), clusterId, linkName, mirrorTopicName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaLinkConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string
		var configName string

		httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaLinkConfig(context.Background(), clusterId, linkName, configName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaLinkConfigBatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaLinkConfigBatch(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaMirrorTopicsFailover", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsFailover(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaMirrorTopicsPause", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPause(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaMirrorTopicsPromote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsPromote(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterLinkingV3ApiService UpdateKafkaMirrorTopicsResume", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string
		var linkName string

		resp, httpRes, err := apiClient.ClusterLinkingV3Api.UpdateKafkaMirrorTopicsResume(context.Background(), clusterId, linkName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
