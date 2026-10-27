# \AgentActivityAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnthropicInferenceHook**](AgentActivityAPI.md#AnthropicInferenceHook) | **Post** /inference-hooks/anthropic | Anthropic inference hook verdict
[**GetAgentKillSwitch**](AgentActivityAPI.md#GetAgentKillSwitch) | **Get** /agent-kill-switch | Get the agent kill switch
[**ReportAgentActivity**](AgentActivityAPI.md#ReportAgentActivity) | **Post** /agent-activity | Report agent activity
[**SetAgentKillSwitch**](AgentActivityAPI.md#SetAgentKillSwitch) | **Put** /agent-kill-switch | Set the agent kill switch



## AnthropicInferenceHook

> AnthropicInferenceHookVerdict AnthropicInferenceHook(ctx).RequestBody(requestBody).Execute()

Anthropic inference hook verdict



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.AnthropicInferenceHook(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.AnthropicInferenceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnthropicInferenceHook`: AnthropicInferenceHookVerdict
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.AnthropicInferenceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnthropicInferenceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AnthropicInferenceHookVerdict**](AnthropicInferenceHookVerdict.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentKillSwitch

> AgentKillSwitch GetAgentKillSwitch(ctx).Execute()

Get the agent kill switch



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.GetAgentKillSwitch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.GetAgentKillSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentKillSwitch`: AgentKillSwitch
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.GetAgentKillSwitch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentKillSwitchRequest struct via the builder pattern


### Return type

[**AgentKillSwitch**](AgentKillSwitch.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportAgentActivity

> AgentActivityIngestResult ReportAgentActivity(ctx).AgentActivityIngestRequest(agentActivityIngestRequest).XOpalActor(xOpalActor).Execute()

Report agent activity



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
	agentActivityIngestRequest := *openapiclient.NewAgentActivityIngestRequest() // AgentActivityIngestRequest | 
	xOpalActor := "xOpalActor_example" // string | The identity the reporting client claims this activity belongs to - an email or an Anthropic user id. **Unverified**: a Claude Code hook payload contains no user identity, so this has to be injected by whoever deploys the hook (via the http handler's `headers` map and `allowedEnvVars` interpolation), and whoever controls the machine controls the header. Used only when the API key itself does not identify a developer. When the email belongs to a person in the organization, that person is recorded as who the activity was reported for (always as a claim), and the activity is attributed to the one Claude agent they own, if they own exactly one. The API key's own email is never treated this way: a key belongs to whoever created it, which with a shared key is not who is at the keyboard. Omitting it is valid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.ReportAgentActivity(context.Background()).AgentActivityIngestRequest(agentActivityIngestRequest).XOpalActor(xOpalActor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.ReportAgentActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportAgentActivity`: AgentActivityIngestResult
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.ReportAgentActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportAgentActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentActivityIngestRequest** | [**AgentActivityIngestRequest**](AgentActivityIngestRequest.md) |  | 
 **xOpalActor** | **string** | The identity the reporting client claims this activity belongs to - an email or an Anthropic user id. **Unverified**: a Claude Code hook payload contains no user identity, so this has to be injected by whoever deploys the hook (via the http handler&#39;s &#x60;headers&#x60; map and &#x60;allowedEnvVars&#x60; interpolation), and whoever controls the machine controls the header. Used only when the API key itself does not identify a developer. When the email belongs to a person in the organization, that person is recorded as who the activity was reported for (always as a claim), and the activity is attributed to the one Claude agent they own, if they own exactly one. The API key&#39;s own email is never treated this way: a key belongs to whoever created it, which with a shared key is not who is at the keyboard. Omitting it is valid. | 

### Return type

[**AgentActivityIngestResult**](AgentActivityIngestResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAgentKillSwitch

> AgentKillSwitch SetAgentKillSwitch(ctx).SetAgentKillSwitchRequest(setAgentKillSwitchRequest).Execute()

Set the agent kill switch



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
	setAgentKillSwitchRequest := *openapiclient.NewSetAgentKillSwitchRequest(false) // SetAgentKillSwitchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.SetAgentKillSwitch(context.Background()).SetAgentKillSwitchRequest(setAgentKillSwitchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.SetAgentKillSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAgentKillSwitch`: AgentKillSwitch
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.SetAgentKillSwitch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAgentKillSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setAgentKillSwitchRequest** | [**SetAgentKillSwitchRequest**](SetAgentKillSwitchRequest.md) |  | 

### Return type

[**AgentKillSwitch**](AgentKillSwitch.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

