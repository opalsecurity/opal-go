# \AccessRulesAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessRule**](AccessRulesAPI.md#GetAccessRule) | **Get** /access-rules/{access_rule_id} | 
[**UpdateAccessRule**](AccessRulesAPI.md#UpdateAccessRule) | **Put** /access-rules/{access_rule_id} | 



## GetAccessRule

> AccessRuleCondition GetAccessRule(ctx, accessRuleId).Execute()





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
	accessRuleId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The access rule ID (group ID) of the access rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.GetAccessRule(context.Background(), accessRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.GetAccessRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRule`: AccessRuleCondition
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.GetAccessRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessRuleId** | **string** | The access rule ID (group ID) of the access rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRuleCondition**](AccessRuleCondition.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessRule

> AccessRuleCondition UpdateAccessRule(ctx, accessRuleId).AccessRuleCondition(accessRuleCondition).Execute()





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
	accessRuleId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The access rule ID (group ID) of the access rule.
	accessRuleCondition := *openapiclient.NewAccessRuleCondition("ACTIVE", *openapiclient.NewRuleClauses(*openapiclient.NewRuleConjunction([]openapiclient.RuleDisjunction{*openapiclient.NewRuleDisjunction([]openapiclient.TagSelector{*openapiclient.NewTagSelector("Key_example", "Value_example", "ConnectionId_example")})}))) // AccessRuleCondition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRulesAPI.UpdateAccessRule(context.Background(), accessRuleId).AccessRuleCondition(accessRuleCondition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRulesAPI.UpdateAccessRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessRule`: AccessRuleCondition
	fmt.Fprintf(os.Stdout, "Response from `AccessRulesAPI.UpdateAccessRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessRuleId** | **string** | The access rule ID (group ID) of the access rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRuleCondition** | [**AccessRuleCondition**](AccessRuleCondition.md) |  | 

### Return type

[**AccessRuleCondition**](AccessRuleCondition.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

