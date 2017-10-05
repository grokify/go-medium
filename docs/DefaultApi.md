# \DefaultApi

All URIs are relative to *https://api.medium.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesPost**](DefaultApi.md#ImagesPost) | **Post** /images | Uploading an image
[**PublicationsPublicationIdPostsPost**](DefaultApi.md#PublicationsPublicationIdPostsPost) | **Post** /publications/{publicationId}/posts | Creating a post under a publication
[**UsersAuthorIdPostsPost**](DefaultApi.md#UsersAuthorIdPostsPost) | **Post** /users/{authorId}/posts | Creating a post


# **ImagesPost**
> InlineResponse2004 ImagesPost()

Uploading an image

Medium will automatically side-load any images specified by the src attribute on an <img> tag in post content when creating a post. However, if you are building a desktop integration and have local image files that you wish to send, you may use the images endpoint.


### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicationsPublicationIdPostsPost**
> InlineResponse2003 PublicationsPublicationIdPostsPost($publicationId, $body)

Creating a post under a publication

This API allows creating a post and associating it with a publication on Medium


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicationId** | **string**| ID of the publication this post was created under. This matches the publication ID requested in the URL of the request | 
 **body** | [**Body1**](Body1.md)| Create post JSON body | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersAuthorIdPostsPost**
> InlineResponse2003 UsersAuthorIdPostsPost($authorId, $body)

Creating a post

Creates a post on the authenticated user’s profile.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **string**| The userId of the post’s author | 
 **body** | [**Body**](Body.md)| Create post JSON body | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

