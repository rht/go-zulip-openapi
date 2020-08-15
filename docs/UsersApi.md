# \UsersApi

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users | 
[**CreateUserGroup**](UsersApi.md#CreateUserGroup) | **Post** /user_groups/create | 
[**DeactivateMyAccount**](UsersApi.md#DeactivateMyAccount) | **Delete** /users/me | 
[**DeactivateUser**](UsersApi.md#DeactivateUser) | **Delete** /users/{user_id} | 
[**GetAttachments**](UsersApi.md#GetAttachments) | **Get** /attachments | 
[**GetOwnUser**](UsersApi.md#GetOwnUser) | **Get** /users/me | 
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{user_id} | 
[**GetUserGroups**](UsersApi.md#GetUserGroups) | **Get** /user_groups | 
[**GetUserPresence**](UsersApi.md#GetUserPresence) | **Get** /users/{email}/presence | 
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users | 
[**ReactivateUser**](UsersApi.md#ReactivateUser) | **Post** /users/{user_id}/reactivate | 
[**RemoveUserGroup**](UsersApi.md#RemoveUserGroup) | **Delete** /user_groups/{group_id} | 
[**SetTypingStatus**](UsersApi.md#SetTypingStatus) | **Post** /typing | 
[**UpdateNotificationSettings**](UsersApi.md#UpdateNotificationSettings) | **Patch** /settings/notifications | 
[**UpdateUser**](UsersApi.md#UpdateUser) | **Patch** /users/{user_id} | 
[**UpdateUserGroup**](UsersApi.md#UpdateUserGroup) | **Patch** /user_groups/{group_id} | 



## CreateUser

> JsonSuccess CreateUser(ctx, email, password, fullName, shortName)



{!api-admin-only.md!}  Create a new user account via the API.  `POST {{ api_url }}/v1/users` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| The email address of the new user.  | 
**password** | **string**| The password of the new user.  | 
**fullName** | **string**| The full name of the new user.  | 
**shortName** | **string**| The short name of the new user.  Not user-visible.  | 

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


## CreateUserGroup

> JsonSuccess CreateUserGroup(ctx, name, description, members)



Create a new [user group](/help/user-groups).  `POST {{ api_url }}/v1/user_groups/create` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the user group.  | 
**description** | **string**| The description of the user group.  | 
**members** | [**[]int32**](int32.md)| An array containing the user IDs of the initial members for the new user group.  | 

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


## DeactivateMyAccount

> JsonSuccess DeactivateMyAccount(ctx, )



Delete the requesting user from the realm. 

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


## DeactivateUser

> JsonSuccess DeactivateUser(ctx, userId)



{!api-admin-only.md!}  [Deactivates a user](https://zulip.com/help/deactivate-or-reactivate-a-user) given their user ID.  `DELETE {{ api_url }}/v1/users/{user_id}` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The target user&#39;s ID.  | 

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


## GetAttachments

> JsonSuccess GetAttachments(ctx, )



Fetch metadata on files uploaded by the requesting user.  `GET {{ api_url }}/v1/attachments` 

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


## GetOwnUser

> JsonSuccess GetOwnUser(ctx, )



Get basic data about the user/bot that requests this endpoint.  `GET {{ api_url }}/v1/users/me` 

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


## GetUser

> JsonSuccess GetUser(ctx, userId, optional)



Fetch details for a single user in the organization.  `GET {{ api_url }}/v1/users/{user_id}`  You can also fetch details on [all users in the organization](/api/get-users).  *This endpoint is new in Zulip Server 3.0 (feature level 1).* 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The target user&#39;s ID.  | 
 **optional** | ***GetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientGravatar** | **optional.Bool**| Whether the client supports computing gravatars URLs.  If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar.  This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  | [default to false]
 **includeCustomProfileFields** | **optional.Bool**| Whether the client wants [custom profile field](/help/add-custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0.  Previous versions do no offer these data via the API.  | [default to false]

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


## GetUserGroups

> JsonSuccess GetUserGroups(ctx, )



Fetches all of the user groups in the organization.  `GET {{ api_url }}/v1/user_groups` 

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


## GetUserPresence

> JsonSuccess GetUserPresence(ctx, email)



Get the presence status for a specific user.  This endpoint is most useful for embedding data about a user's presence status in other sites (E.g. an employee directory).  Full Zulip clients like mobile/desktop apps will want to use the main presence endpoint, which returns data for all active users in the organization, instead.  `GET {{ api_url }}/v1/users/{email}/presence`  See [Zulip's developer documentation](https://zulip.readthedocs.io/en/latest/subsystems/presence.html) for details on the data model for presence in Zulip. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| The email address of the user whose presence you want to fetch.  | 

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


## GetUsers

> JsonSuccess GetUsers(ctx, optional)



Retrieve details on all users in the organization.  Optionally includes values of [custom profile field](/help/add-custom-profile-fields).  `GET {{ api_url }}/v1/users` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGravatar** | **optional.Bool**| Whether the client supports computing gravatars URLs.  If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar.  This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  | [default to false]
 **includeCustomProfileFields** | **optional.Bool**| Whether the client wants [custom profile field](/help/add-custom-profile-fields) data to be included in the response.  **Changes**: New in Zulip 2.1.0.  Previous versions do no offer these data via the API.  | [default to false]

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


## ReactivateUser

> JsonSuccess ReactivateUser(ctx, userId)



{!api-admin-only.md!}  [Reactivates a user](https://zulip.com/help/deactivate-or-reactivate-a-user) given their user ID.  `POST {{ api_url }}/v1/users/{user_id}/reactivate` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The target user&#39;s ID.  | 

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


## RemoveUserGroup

> JsonSuccess RemoveUserGroup(ctx, groupId)



Delete a [user group](/help/user-groups).  `DELETE {{ api_url }}/v1/user_groups/{group_id}` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32**| The ID of the target user group.  | 

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


## SetTypingStatus

> JsonSuccess SetTypingStatus(ctx, op, to)



Send an event indicating that the user has started or stopped typing on their client.  See [the typing notification docs](https://zulip.readthedocs.io/en/latest/subsystems/typing-indicators.html) for details on Zulip's typing notifications protocol.  `POST {{ api_url }}/v1/typing` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**op** | **string**| Whether the user has started (&#x60;start&#x60;) or stopped (&#x60;stop&#x60;) to type.  | 
**to** | [**[]int32**](int32.md)| The user_ids of the recipients of the message being typed. Typing notifications are only supported for private messages. Send a JSON-encoded list of user_ids. (Use a list even if there is only one recipient.).  **Changes**: Before Zulip 2.0, this parameter accepted only a JSON-encoded list of email addresses.  Support for the email address-based format was removed in Zulip 3.0 (feature level 11).  | 

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


## UpdateNotificationSettings

> JsonSuccess UpdateNotificationSettings(ctx, optional)



This endpoint is used to edit the user's global notification settings. See [this endpoint](/api/update-subscription-settings) for per-stream notification settings.  `PATCH {{ api_url }}/v1/settings/notifications` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateNotificationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNotificationSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableStreamDesktopNotifications** | **optional.Bool**| Enable visual desktop notifications for stream messages.  | 
 **enableStreamEmailNotifications** | **optional.Bool**| Enable email notifications for stream messages.  | 
 **enableStreamPushNotifications** | **optional.Bool**| Enable mobile notifications for stream messages.  | 
 **enableStreamAudibleNotifications** | **optional.Bool**| Enable audible desktop notifications for stream messages.  | 
 **notificationSound** | **optional.String**| Notification sound name.  | 
 **enableDesktopNotifications** | **optional.Bool**| Enable visual desktop notifications for private messages and @-mentions.  | 
 **enableSounds** | **optional.Bool**| Enable audible desktop notifications for private messages and @-mentions.  | 
 **enableOfflineEmailNotifications** | **optional.Bool**| Enable email notifications for private messages and @-mentions received when the user is offline.  | 
 **enableOfflinePushNotifications** | **optional.Bool**| Enable mobile notification for private messages and @-mentions received when the user is offline.  | 
 **enableOnlinePushNotifications** | **optional.Bool**| Enable mobile notification for private messages and @-mentions received when the user is online.  | 
 **enableDigestEmails** | **optional.Bool**| Enable digest emails when the user is away.  | 
 **enableLoginEmails** | **optional.Bool**| Enable email notifications for new logins to account.  | 
 **messageContentInEmailNotifications** | **optional.Bool**| Include the message&#39;s content in missed messages email notifications.  | 
 **pmContentInDesktopNotifications** | **optional.Bool**| Include content of private messages in desktop notifications.  | 
 **wildcardMentionsNotify** | **optional.Bool**| Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  | 
 **desktopIconCountDisplay** | **optional.Int32**| Unread count summary (appears in desktop sidebar and browser tab)  * 1 - All unreads * 2 - Private messages and mentions * 3 - None  | 
 **realmNameInNotifications** | **optional.Bool**| Include organization name in subject of missed message emails.  | 
 **presenceEnabled** | **optional.Bool**| Display the presence status to other users when online.  | 

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


## UpdateUser

> JsonSuccess UpdateUser(ctx, userId, optional)



{!api-admin-only.md!}  Administrative endpoint to update the details of another user in the organization.  `PATCH {{ api_url }}/v1/users/{user_id}`  Supports everything an administrator can do to edit details of another user's account, including editing full name, [role](/help/roles-and-permissions), and [custom profile fields](/help/add-custom-profile-fields). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32**| The target user&#39;s ID.  | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullName** | **optional.String**| The user&#39;s full name.  | 
 **role** | **optional.Int32**| New [role](/help/roles-and-permissions) for the user.  Roles are encoded as:  * Organization owner: 100 * Organization administrator: 200 * Member: 400 * Guest: 600  Only organization owners can add or remove the owner role.  The owner role cannot be removed from the only organization owner.  **Changes**: New in Zulip 3.0 (feature level 8), replacing the previous pair of &#x60;is_admin&#x60; and &#x60;is_guest&#x60; boolean parameters.  | 
 **profileData** | [**optional.Interface of []map[string]interface{}**](map[string]interface{}.md)| A dictionary containing the to be updated custom profile field data for the user.  | 

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


## UpdateUserGroup

> JsonSuccess UpdateUserGroup(ctx, groupId, name, description)



Update the name or description of a [user group](/help/user-groups).  `PATCH {{ api_url }}/v1/user_groups/{group_id}` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32**| The ID of the target user group.  | 
**name** | **string**| The new name of the group.  | 
**description** | **string**| The new description of the group.  | 

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

