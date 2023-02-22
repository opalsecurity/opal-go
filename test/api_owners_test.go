/*
Opal API

Testing OwnersApiService

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

func Test_opal_OwnersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OwnersApiService CreateOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OwnersApi.CreateOwner(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService DeleteOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerId string

		httpRes, err := apiClient.OwnersApi.DeleteOwner(context.Background(), ownerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService GetOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerId string

		resp, httpRes, err := apiClient.OwnersApi.GetOwner(context.Background(), ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService GetOwnerFromName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerName string

		resp, httpRes, err := apiClient.OwnersApi.GetOwnerFromName(context.Background(), ownerName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService GetOwnerUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerId string

		resp, httpRes, err := apiClient.OwnersApi.GetOwnerUsers(context.Background(), ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService GetOwners", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OwnersApi.GetOwners(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService SetOwnerUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerId string

		resp, httpRes, err := apiClient.OwnersApi.SetOwnerUsers(context.Background(), ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnersApiService UpdateOwners", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OwnersApi.UpdateOwners(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
