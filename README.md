# Go API client for medium

Articles that matter on social publishing platform

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./medium"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.medium.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ImagesPost**](docs/DefaultApi.md#imagespost) | **Post** /images | Uploading an image
*DefaultApi* | [**PublicationsPublicationIdPostsPost**](docs/DefaultApi.md#publicationspublicationidpostspost) | **Post** /publications/{publicationId}/posts | Creating a post under a publication
*DefaultApi* | [**UsersAuthorIdPostsPost**](docs/DefaultApi.md#usersauthoridpostspost) | **Post** /users/{authorId}/posts | Creating a post
*PublicationsApi* | [**PublicationsPublicationIdContributorsGet**](docs/PublicationsApi.md#publicationspublicationidcontributorsget) | **Get** /publications/{publicationId}/contributors | Publication Contributors
*PublicationsApi* | [**UsersUserIdPublicationsGet**](docs/PublicationsApi.md#usersuseridpublicationsget) | **Get** /users/{userId}/publications | User Publications
*UserApi* | [**MeGet**](docs/UserApi.md#meget) | **Get** /me | User Profile


## Documentation For Models

 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Contributor](docs/Contributor.md)
 - [ContributorListData](docs/ContributorListData.md)
 - [Image](docs/Image.md)
 - [ImageData](docs/ImageData.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2001Data](docs/InlineResponse2001Data.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2002Data](docs/InlineResponse2002Data.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2003Data](docs/InlineResponse2003Data.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2004Data](docs/InlineResponse2004Data.md)
 - [InlineResponse200Data](docs/InlineResponse200Data.md)
 - [Post](docs/Post.md)
 - [PostData](docs/PostData.md)
 - [PostRequest](docs/PostRequest.md)
 - [Publication](docs/Publication.md)
 - [PublicationListData](docs/PublicationListData.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author



