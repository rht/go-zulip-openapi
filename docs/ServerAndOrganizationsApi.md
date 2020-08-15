# \ServerAndOrganizationsApi

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLinkifier**](ServerAndOrganizationsApi.md#AddLinkifier) | **Post** /realm/filters | 
[**GetCustomEmoji**](ServerAndOrganizationsApi.md#GetCustomEmoji) | **Get** /realm/emoji | 
[**GetLinkifiers**](ServerAndOrganizationsApi.md#GetLinkifiers) | **Get** /realm/filters | 
[**GetServerSettings**](ServerAndOrganizationsApi.md#GetServerSettings) | **Get** /server_settings | 
[**RemoveLinkifier**](ServerAndOrganizationsApi.md#RemoveLinkifier) | **Delete** /realm/filters/{filter_id} | 
[**UploadCustomEmoji**](ServerAndOrganizationsApi.md#UploadCustomEmoji) | **Post** /realm/emoji/{emoji_name} | 



## AddLinkifier

> JsonSuccess AddLinkifier(ctx, pattern, urlFormatString)



Configure [linkifiers](/help/add-a-custom-linkification-filter), regular expression patterns that are automatically linkified when they appear in messages and topics.  `POST {{ api_url }}/v1/realm/filters` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pattern** | **string**| The [Python regular expression](https://docs.python.org/3/howto/regex.html) that should trigger the linkifier.  | 
**urlFormatString** | **string**| The URL used for the link. If you used named groups for the &#x60;pattern&#x60;, you can insert their content here with &#x60;%(name_of_the_capturing_group)s&#x60;.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEmoji

> JsonSuccess GetCustomEmoji(ctx, )



Get all the custom emoji in the user's organization.  `GET {{ api_url }}/v1/realm/emoji` 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkifiers

> JsonSuccess GetLinkifiers(ctx, )



List all of an organization's configured [linkifiers](/help/add-a-custom-linkification-filter), regular expression patterns that are automatically linkified when they appear in messages and topics.  `GET {{ api_url }}/v1/realm/filters` 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerSettings

> JsonSuccess GetServerSettings(ctx, )



Fetch global settings for a Zulip server.  `GET {{ api_url }}/v1/server_settings`  **Note:** this endpoint does not require any authentication at all, and you can use it to check:  * If this is a Zulip server, and if so, what version of Zulip it's running. * What a Zulip client (e.g. a mobile app or [zulip-terminal](https://github.com/zulip/zulip-terminal/)) needs to know in order to display a login prompt for the server (e.g. what authentication methods are available). 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLinkifier

> JsonSuccess RemoveLinkifier(ctx, filterId)



Remove [linkifiers](/help/add-a-custom-linkification-filter), regular expression patterns that are automatically linkified when they appear in messages and topics.  `DELETE {{ api_url }}/v1/realm/filters/{filter_id}` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32**| The ID of the filter that you want to remove.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCustomEmoji

> JsonSuccess UploadCustomEmoji(ctx, emojiName, optional)



This endpoint is used to upload a custom emoji for use in the user's organization.  Access to this endpoint depends on the [organization's configuration](https://zulip.com/help/only-allow-admins-to-add-emoji).  `POST {{ api_url }}/v1/realm/emoji/{emoji_name}` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emojiName** | **string**| The name that should be associated with the uploaded emoji image/gif.  | 
 **optional** | ***UploadCustomEmojiOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadCustomEmojiOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

