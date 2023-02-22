/*
Opal API

Testing TagsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package opal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/opalsecurity/opal-go"
)

func Test_opal_TagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsApiService AddGroupTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var groupId string

		httpRes, err := apiClient.TagsApi.AddGroupTag(context.Background(), tagId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService AddResourceTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var resourceId string

		httpRes, err := apiClient.TagsApi.AddResourceTag(context.Background(), tagId, resourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService AddUserTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var userId string

		httpRes, err := apiClient.TagsApi.AddUserTag(context.Background(), tagId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService CreateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsApi.CreateTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsApi.GetTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsApi.GetTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService RemoveGroupTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var groupId string

		httpRes, err := apiClient.TagsApi.RemoveGroupTag(context.Background(), tagId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService RemoveResourceTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var resourceId string

		httpRes, err := apiClient.TagsApi.RemoveResourceTag(context.Background(), tagId, resourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService RemoveUserTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string
		var userId string

		httpRes, err := apiClient.TagsApi.RemoveUserTag(context.Background(), tagId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
