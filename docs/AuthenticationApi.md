# \AuthenticationApi

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevFetchApiKey**](AuthenticationApi.md#DevFetchApiKey) | **Post** /dev_fetch_api_key | 
[**FetchApiKey**](AuthenticationApi.md#FetchApiKey) | **Post** /fetch_api_key | 



## DevFetchApiKey

> ApiKeyResponse DevFetchApiKey(ctx, username)



For easy testing of mobile apps and other clients and against Zulip development servers, we support fetching a Zulip API key for any user on the development server without authentication (so that they can implement analogues of the one-click login process available for Zulip development servers on the web).  **Note:** This endpoint is only available on Zulip development servers; for obvious security reasons it will always return an error in a Zulip production server.  `POST {{ api_url }}/v1/dev_fetch_api_key` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| The email address for the user that owns the API key.  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchApiKey

> ApiKeyResponse FetchApiKey(ctx, username, password)



Given a username and password, fetch the user's API key.  Used to authenticate the mobile and terminal apps when the server has EmailAuthBackend or LDAPAuthBackend enabled. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| The username to be used for authentication (typically, the email address, but it could be an LDAP username).  | 
**password** | **string**| The user&#39;s Zulip password (or LDAP password, if LDAP authentication is in use).  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

