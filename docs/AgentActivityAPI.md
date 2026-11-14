# \AgentActivityAPI

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnthropicInferenceHook**](AgentActivityAPI.md#AnthropicInferenceHook) | **Post** /inference-hooks/anthropic | Anthropic inference hook verdict
[**GetAgentActivityClientSettings**](AgentActivityAPI.md#GetAgentActivityClientSettings) | **Get** /agent-activity/client-settings | Get agent activity client settings
[**GetAgentKillSwitch**](AgentActivityAPI.md#GetAgentKillSwitch) | **Get** /agent-kill-switch | Get the agent kill switch
[**ReportAgentActivity**](AgentActivityAPI.md#ReportAgentActivity) | **Post** /agent-activity | Report agent activity
[**ReportMcpConfig**](AgentActivityAPI.md#ReportMcpConfig) | **Post** /agent-activity/mcp-config | Report MCP server configuration
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


## GetAgentActivityClientSettings

> AgentActivityClientSettings GetAgentActivityClientSettings(ctx).Execute()

Get agent activity client settings



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
	resp, r, err := apiClient.AgentActivityAPI.GetAgentActivityClientSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.GetAgentActivityClientSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentActivityClientSettings`: AgentActivityClientSettings
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.GetAgentActivityClientSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentActivityClientSettingsRequest struct via the builder pattern


### Return type

[**AgentActivityClientSettings**](AgentActivityClientSettings.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

> AgentActivityIngestResult ReportAgentActivity(ctx).AgentActivityIngestRequest(agentActivityIngestRequest).XOpalActor(xOpalActor).XOpalHost(xOpalHost).XOpalContent(xOpalContent).Execute()

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
	xOpalActor := "xOpalActor_example" // string | The identity the reporting client claims this activity belongs to - an email or an Anthropic user id. **Unverified**: a Claude Code hook payload contains no user identity, so this has to be injected by whoever deploys the hook (via the http handler's `headers` map and `allowedEnvVars` interpolation), and whoever controls the machine controls the header. Used only when the API key itself does not identify a developer. When the email belongs to a person in the organization, that person is recorded as who the activity was reported for (always as a claim), and the activity is attributed to the one Claude agent they own, if they own exactly one. The API key's own email is never treated this way: a key belongs to whoever created it, which with a shared key is not who is at the keyboard. Omitting it is valid. When an event's `opal_reported` object carries other emails, Opal also tries them, in the order `configured_email`, `claude_account_email`, `git_email`, after this header and skipping any that repeat it. The first one that resolves to an agent wins. (optional)
	xOpalHost := "xOpalHost_example" // string | The name of the machine that sent this request, as that machine reports it (`hostname -s`). **Unverified**: two machines can share a name, and whoever controls the machine controls the header. Stored with each event, so a reader can tell which machine reported it. When an event's own `opal_reported.hostname` names a different machine, this header wins and the body's value is kept beside it. Omitting it is valid. (optional)
	xOpalContent := "xOpalContent_example" // string | `1` when the sender collects conversation content: the prompt of a `UserPromptSubmit` event, the `tool_response` of a `PostToolUse` event and the `last_assistant_message` of a `Stop` or `SubagentStop` event. Without it (or an event's own `opal_content`) those fields are read and discarded, as before. An `http` hook sets it in its `headers` map, because it cannot change the payload. Opal stores the content only when the organization stores content as well. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.ReportAgentActivity(context.Background()).AgentActivityIngestRequest(agentActivityIngestRequest).XOpalActor(xOpalActor).XOpalHost(xOpalHost).XOpalContent(xOpalContent).Execute()
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
 **xOpalActor** | **string** | The identity the reporting client claims this activity belongs to - an email or an Anthropic user id. **Unverified**: a Claude Code hook payload contains no user identity, so this has to be injected by whoever deploys the hook (via the http handler&#39;s &#x60;headers&#x60; map and &#x60;allowedEnvVars&#x60; interpolation), and whoever controls the machine controls the header. Used only when the API key itself does not identify a developer. When the email belongs to a person in the organization, that person is recorded as who the activity was reported for (always as a claim), and the activity is attributed to the one Claude agent they own, if they own exactly one. The API key&#39;s own email is never treated this way: a key belongs to whoever created it, which with a shared key is not who is at the keyboard. Omitting it is valid. When an event&#39;s &#x60;opal_reported&#x60; object carries other emails, Opal also tries them, in the order &#x60;configured_email&#x60;, &#x60;claude_account_email&#x60;, &#x60;git_email&#x60;, after this header and skipping any that repeat it. The first one that resolves to an agent wins. | 
 **xOpalHost** | **string** | The name of the machine that sent this request, as that machine reports it (&#x60;hostname -s&#x60;). **Unverified**: two machines can share a name, and whoever controls the machine controls the header. Stored with each event, so a reader can tell which machine reported it. When an event&#39;s own &#x60;opal_reported.hostname&#x60; names a different machine, this header wins and the body&#39;s value is kept beside it. Omitting it is valid. | 
 **xOpalContent** | **string** | &#x60;1&#x60; when the sender collects conversation content: the prompt of a &#x60;UserPromptSubmit&#x60; event, the &#x60;tool_response&#x60; of a &#x60;PostToolUse&#x60; event and the &#x60;last_assistant_message&#x60; of a &#x60;Stop&#x60; or &#x60;SubagentStop&#x60; event. Without it (or an event&#39;s own &#x60;opal_content&#x60;) those fields are read and discarded, as before. An &#x60;http&#x60; hook sets it in its &#x60;headers&#x60; map, because it cannot change the payload. Opal stores the content only when the organization stores content as well. | 

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


## ReportMcpConfig

> McpConfigReportResult ReportMcpConfig(ctx).McpConfigReport(mcpConfigReport).XOpalActor(xOpalActor).XOpalHost(xOpalHost).Execute()

Report MCP server configuration



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
	mcpConfigReport := *openapiclient.NewMcpConfigReport([]openapiclient.McpConfigServer{*openapiclient.NewMcpConfigServer("Name_example")}) // McpConfigReport | 
	xOpalActor := "xOpalActor_example" // string | The identity the reporting client claims - an email, or an operating-system user name. **Unverified**, used the same way as on `/agent-activity`. (optional)
	xOpalHost := "xOpalHost_example" // string | The name of the machine that sent the report, as that machine reports it (`hostname -s`). **Unverified.** A server this machine listed before and this report leaves out is marked removed for the reporting person, unless another of their machines still lists it. Omitting it is valid: the reports without one count as one machine. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.ReportMcpConfig(context.Background()).McpConfigReport(mcpConfigReport).XOpalActor(xOpalActor).XOpalHost(xOpalHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.ReportMcpConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportMcpConfig`: McpConfigReportResult
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.ReportMcpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportMcpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcpConfigReport** | [**McpConfigReport**](McpConfigReport.md) |  | 
 **xOpalActor** | **string** | The identity the reporting client claims - an email, or an operating-system user name. **Unverified**, used the same way as on &#x60;/agent-activity&#x60;. | 
 **xOpalHost** | **string** | The name of the machine that sent the report, as that machine reports it (&#x60;hostname -s&#x60;). **Unverified.** A server this machine listed before and this report leaves out is marked removed for the reporting person, unless another of their machines still lists it. Omitting it is valid: the reports without one count as one machine. | 

### Return type

[**McpConfigReportResult**](McpConfigReportResult.md)

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

