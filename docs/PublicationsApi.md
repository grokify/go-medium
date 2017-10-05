# \PublicationsApi

All URIs are relative to *https://api.medium.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicationsPublicationIdContributorsGet**](PublicationsApi.md#PublicationsPublicationIdContributorsGet) | **Get** /publications/{publicationId}/contributors | Publication Contributors
[**UsersUserIdPublicationsGet**](PublicationsApi.md#UsersUserIdPublicationsGet) | **Get** /users/{userId}/publications | User Publications


# **PublicationsPublicationIdContributorsGet**
> InlineResponse2002 PublicationsPublicationIdContributorsGet($publicationId)

Publication Contributors

Fetching contributors for a publication


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicationId** | **string**| The id of the publication being queried | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPublicationsGet**
> InlineResponse2001 UsersUserIdPublicationsGet($userId)

User Publications

Listing the user’s publications


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user being queried | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

