# Body1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the post. Note that this title is used for SEO and when rendering the post as a listing, but will not appear in the actual post—for that, the title must be specified in the content field as well. Titles longer than 100 characters will be ignored. In that case, a title will be synthesized from the first content in the post when it is published. | [optional] [default to null]
**ContentFormat** | **string** | The format of the \&quot;content\&quot; field. There are two valid values, \&quot;html\&quot;, and \&quot;markdown\&quot; | [optional] [default to null]
**Content** | **string** | The body of the post, in a valid, semantic, HTML fragment, or Markdown. Further markups may be supported in the future. For a full list of accepted HTML tags, see here. If you want your title to appear on the post page, you must also include it as part of the post content. | [optional] [default to null]
**Tags** | **string** | Tags to classify the post. Only the first three will be used. Tags longer than 25 characters will be ignored. | [optional] [default to null]
**CanonicalUrl** | **string** | The original home of this content, if it was originally published elsewhere. | [optional] [default to null]
**PublishStatus** | **string** | The status of the post. Valid values are “public”, “draft”, or “unlisted”. The default is “public”. | [optional] [default to null]
**License** | **string** | The license of the post. Valid values are “all-rights-reserved”, “cc-40-by”, “cc-40-by-sa”, “cc-40-by-nd”, “cc-40-by-nc”, “cc-40-by-nc-nd”, “cc-40-by-nc-sa”, “cc-40-zero”, “public-domain”. The default is “all-rights-reserved”. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


