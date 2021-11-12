# \UsersApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**User**](UsersApi.md#User) | **Get** /user | 



## User

> User User(ctx).UserId(userId).Email(email).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "32acc112-21ff-4669-91c2-21e27683eaa1" // string | The user ID of the user. (optional)
    email := "johndoe@domain.org" // string | The email of the user. If both user ID and email are provided, user ID will take precedence. If neither are provided, an error will occur. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.User(context.Background()).UserId(userId).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.User``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `User`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.User`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user ID of the user. | 
 **email** | **string** | The email of the user. If both user ID and email are provided, user ID will take precedence. If neither are provided, an error will occur. | 

### Return type

[**User**](User.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

