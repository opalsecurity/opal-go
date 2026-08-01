# \CampaignsAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /campaigns | 
[**EndCampaign**](CampaignsAPI.md#EndCampaign) | **Post** /campaigns/{campaign_id}/end | End campaign
[**GetCampaign**](CampaignsAPI.md#GetCampaign) | **Get** /campaigns/{campaign_id} | Get campaign by ID
[**GetCampaigns**](CampaignsAPI.md#GetCampaigns) | **Get** /campaigns | 
[**StartCampaign**](CampaignsAPI.md#StartCampaign) | **Post** /campaigns/{campaign_id}/start | Start campaign
[**StopCampaign**](CampaignsAPI.md#StopCampaign) | **Post** /campaigns/{campaign_id}/stop | Stop campaign
[**UpdateCampaign**](CampaignsAPI.md#UpdateCampaign) | **Put** /campaigns/{campaign_id} | Update campaign



## CreateCampaign

> Campaign CreateCampaign(ctx).CreateCampaignInfo(createCampaignInfo).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	createCampaignInfo := *openapiclient.NewCreateCampaignInfo("Q3 Access Review", *openapiclient.NewCreateCampaignConfigurationInfo(*openapiclient.NewOpalAccessPathQueryBody())) // CreateCampaignInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaign(context.Background()).CreateCampaignInfo(createCampaignInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCampaignInfo** | [**CreateCampaignInfo**](CreateCampaignInfo.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndCampaign

> Campaign EndCampaign(ctx, campaignId).Execute()

End campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	campaignId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | The ID of the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.EndCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.EndCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.EndCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, campaignId).Execute()

Get campaign by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	campaignId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | The ID of the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> PaginatedCampaignsList GetCampaigns(ctx).Cursor(cursor).PageSize(pageSize).Name(name).Status(status).CreatedAtAfter(createdAtAfter).CreatedAtBefore(createdAtBefore).StartedAtAfter(startedAtAfter).StartedAtBefore(startedAtBefore).EndedAtAfter(endedAtAfter).EndedAtBefore(endedAtBefore).StoppedAtAfter(stoppedAtAfter).StoppedAtBefore(stoppedAtBefore).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
	pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)
	name := "Q3 Access Review" // string | Campaign name to filter by. Returns campaigns whose names contain this substring (case-insensitive). (optional)
	status := openapiclient.CampaignStatusEnum("DRAFT") // CampaignStatusEnum | Filter by campaign status. Status is derived from lifecycle timestamps and review progress. (optional)
	createdAtAfter := time.Now() // time.Time | Include campaigns created after this timestamp (exclusive). ISO 8601 format. (optional)
	createdAtBefore := time.Now() // time.Time | Include campaigns created before this timestamp (exclusive). ISO 8601 format. (optional)
	startedAtAfter := time.Now() // time.Time | Include campaigns started after this timestamp (exclusive). ISO 8601 format. (optional)
	startedAtBefore := time.Now() // time.Time | Include campaigns started before this timestamp (exclusive). ISO 8601 format. (optional)
	endedAtAfter := time.Now() // time.Time | Include campaigns ended after this timestamp (exclusive). ISO 8601 format. (optional)
	endedAtBefore := time.Now() // time.Time | Include campaigns ended before this timestamp (exclusive). ISO 8601 format. (optional)
	stoppedAtAfter := time.Now() // time.Time | Include campaigns stopped after this timestamp (exclusive). ISO 8601 format. (optional)
	stoppedAtBefore := time.Now() // time.Time | Include campaigns stopped before this timestamp (exclusive). ISO 8601 format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaigns(context.Background()).Cursor(cursor).PageSize(pageSize).Name(name).Status(status).CreatedAtAfter(createdAtAfter).CreatedAtBefore(createdAtBefore).StartedAtAfter(startedAtAfter).StartedAtBefore(startedAtBefore).EndedAtAfter(endedAtAfter).EndedAtBefore(endedAtBefore).StoppedAtAfter(stoppedAtAfter).StoppedAtBefore(stoppedAtBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: PaginatedCampaignsList
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **name** | **string** | Campaign name to filter by. Returns campaigns whose names contain this substring (case-insensitive). | 
 **status** | [**CampaignStatusEnum**](CampaignStatusEnum.md) | Filter by campaign status. Status is derived from lifecycle timestamps and review progress. | 
 **createdAtAfter** | **time.Time** | Include campaigns created after this timestamp (exclusive). ISO 8601 format. | 
 **createdAtBefore** | **time.Time** | Include campaigns created before this timestamp (exclusive). ISO 8601 format. | 
 **startedAtAfter** | **time.Time** | Include campaigns started after this timestamp (exclusive). ISO 8601 format. | 
 **startedAtBefore** | **time.Time** | Include campaigns started before this timestamp (exclusive). ISO 8601 format. | 
 **endedAtAfter** | **time.Time** | Include campaigns ended after this timestamp (exclusive). ISO 8601 format. | 
 **endedAtBefore** | **time.Time** | Include campaigns ended before this timestamp (exclusive). ISO 8601 format. | 
 **stoppedAtAfter** | **time.Time** | Include campaigns stopped after this timestamp (exclusive). ISO 8601 format. | 
 **stoppedAtBefore** | **time.Time** | Include campaigns stopped before this timestamp (exclusive). ISO 8601 format. | 

### Return type

[**PaginatedCampaignsList**](PaginatedCampaignsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCampaign

> Campaign StartCampaign(ctx, campaignId).Execute()

Start campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	campaignId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | The ID of the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.StartCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.StartCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.StartCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCampaign

> Campaign StopCampaign(ctx, campaignId).StopCampaignRequest(stopCampaignRequest).Execute()

Stop campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	campaignId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | The ID of the campaign.
	stopCampaignRequest := *openapiclient.NewStopCampaignRequest() // StopCampaignRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.StopCampaign(context.Background(), campaignId).StopCampaignRequest(stopCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.StopCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.StopCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopCampaignRequest** | [**StopCampaignRequest**](StopCampaignRequest.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, campaignId).UpdateCampaignInfo(updateCampaignInfo).Execute()

Update campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opalsecurity/opal-go"
)

func main() {
	campaignId := "f454d283-ca87-4a8a-bdbb-df212eca5353" // string | The ID of the campaign.
	updateCampaignInfo := *openapiclient.NewUpdateCampaignInfo() // UpdateCampaignInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.UpdateCampaign(context.Background(), campaignId).UpdateCampaignInfo(updateCampaignInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.UpdateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignInfo** | [**UpdateCampaignInfo**](UpdateCampaignInfo.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

