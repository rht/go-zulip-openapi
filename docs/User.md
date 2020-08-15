# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The Zulip API email address of the user or bot.  If you do not have permission to view the email address of the target user, this will be a fake email address that is usable for the Zulip API but nothing else.  | [optional] 
**IsBot** | **bool** | A boolean specifying whether the user is a bot or full account.  | [optional] 
**AvatarUrl** | Pointer to **string** | URL for the the user&#39;s avatar.  Will be &#x60;null&#x60; if the &#x60;client_gravatar&#x60; query parameter was set to &#x60;True&#x60; and the user&#39;s avatar is hosted by the Gravatar provider (i.e. the user has never uploaded an avatar).  | [optional] 
**AvatarVersion** | **int32** | Version for the the user&#39;s avatar.  Used for cache-busting requests for the user&#39;s avatar.  Clients generally shouldn&#39;t need to use this; most avatar URLs sent by Zulip will already end with &#x60;?v&#x3D;{avatar_version}&#x60;.  | [optional] 
**FullName** | **string** | Full name of the user or bot, used for all display purposes.  | [optional] 
**IsAdmin** | **bool** | A boolean specifying whether the user is an organization administrator.  | [optional] 
**IsOwner** | **bool** | A boolean specifying whether the user is an organization owner. If true, is_admin will also be true.  **Changes**: New in Zulip 3.0 (feature level 8).  | [optional] 
**BotType** | Pointer to **int32** | An integer describing the type of bot: * &#x60;null&#x60; if the user isn&#39;t a bot. * &#x60;1&#x60; for a &#x60;Generic&#x60; bot. * &#x60;2&#x60; for an &#x60;Incoming webhook&#x60; bot. * &#x60;3&#x60; for an &#x60;Outgoing webhook&#x60; bot. * &#x60;4&#x60; for an &#x60;Embedded&#x60; bot.  | [optional] 
**UserId** | **int32** | The unique ID of the user.  | [optional] 
**BotOwnerId** | Pointer to **int32** | If the user is a bot (i.e. &#x60;is_bot&#x60; is &#x60;True&#x60;), &#x60;bot_owner&#x60; is the user ID of the bot&#39;s owner (usually, whoever created the bot).  Will be null for legacy bots that do not have an owner.  **Changes**: New in Zulip 3.0 (feature level 1).  In previous versions, there was a &#x60;bot_owner&#x60; field containing the email address of the bot&#39;s owner.  | [optional] 
**IsActive** | **bool** | A boolean specifying whether the user account has been deactivated.  | [optional] 
**IsGuest** | **bool** | A boolean specifying whether the user is a guest user.  | [optional] 
**Timezone** | **string** | The time zone of the user.  | [optional] 
**DateJoined** | **string** | The time the user account was created.  | [optional] 
**DeliveryEmail** | **string** | The user&#39;s real email address.  This field is present only if [email address visibility](/help/restrict-visibility-of-email-addresses) is limited and you are an administrator with access to real email addresses under the configured policy.  | [optional] 
**ProfileData** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | A dictionary containing custom profile field data for the user. Each entry maps the integer ID of a custom profile field in the organization to a dictionary containing the user&#39;s data for that field.  Generally the data includes just a single &#x60;value&#x60; key; for those custom profile fields supporting markdown, a &#x60;rendered_value&#x60; key will also be present.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


