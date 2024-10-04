# \DepthApi

All URIs are relative to *https://stage.garantex.biz/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV2Depth**](DepthApi.md#GetV2Depth) | **Get** /v2/depth | 


# **GetV2Depth**
> GetV2Depth(ctx, market)


Get depth or specified market. Both asks and bids are sorted from highest price to lowest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **market** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

