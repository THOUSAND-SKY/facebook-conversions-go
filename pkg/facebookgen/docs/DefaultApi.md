# \DefaultApi

All URIs are relative to *https://graph.facebook.com/v13.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PixelIdEventsPost**](DefaultApi.md#PixelIdEventsPost) | **Post** /{pixelId}/events | 


# **PixelIdEventsPost**
> ResponseSuccess PixelIdEventsPost(ctx, body, pixelId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventRequest**](EventRequest.md)| Facebook Conversions API (for Web) post request | 
  **pixelId** | **string**|  | 

### Return type

[**ResponseSuccess**](response_success.md)

### Authorization

[facebook_api_key](../README.md#facebook_api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

