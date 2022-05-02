# UserData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Em** | **string** | A hashed email address in lower case using SHA-256 algorithm. | [optional] [default to null]
**Ph** | **string** | A hashed phone number using SHA-256 algorithm. Include only digits with country code, area code, and number. | [optional] [default to null]
**Ge** | **string** | A hashed gender (f or m) using SHA-256 algorithm. | [optional] [default to null]
**Db** | **string** | A hashed date of birth given as year, month, and day using SHA-256 algorithm | [optional] [default to null]
**Ln** | **string** | A hashed last name in lowercase using SHA-256 algorithm. | [optional] [default to null]
**Fn** | **string** | A hashed first name in lowercase using SHA-256 algorithm. | [optional] [default to null]
**Ct** | **string** | A hashed city in lower-case without spaces or punctuation using SHA-256 algorithm. | [optional] [default to null]
**Country** | **string** | A hashed two-letter country code in lowercase using SHA-256 algorithm. | [optional] [default to null]
**St** | **string** | A hashed two-letter state code in lowercase using SHA-256 algorithm. | [optional] [default to null]
**Zp** | **string** | A hashed zip code using SHA-256 algorithm. If you are in the United States, this is a five-digit zip code. For other locations, follow each country&#39;s standards. | [optional] [default to null]
**ExternalId** | **string** | Any unique ID from the advertiser, such as loyalty membership IDs, user IDs, and external cookie IDs. If External ID is being sent via other channels, it should be sent in the same format via Conversions API. Hashing external_id using SHA-256 algorithm is optional. | [optional] [default to null]
**ClientIpAddress** | **string** | The IP address of the browser corresponding to the event. | [optional] [default to null]
**ClientUserAgent** | **string** | The user agent for the browser corresponding to the event. | [optional] [default to null]
**Fbc** | **string** | The Facebook click ID value stored in the _fbc browser cookie under your domain. See Managing fbc and fbp Parameters for how to get this value (https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/fbp-and-fbc), or generate this value from a fbclid query parameter. | [optional] [default to null]
**Fbp** | **string** | The Facebook browser ID value stored in the _fbp browser cookie under your domain. See Managing fbc and fbp Parameters for how to get this value (https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/fbp-and-fbc). | [optional] [default to null]
**SubscriptionId** | **string** | The subscription ID for the user in this transaction. This is similar to the order ID for an individual product. | [optional] [default to null]
**FbLoginId** | **int64** | Do not hash. The ID issued by Facebook when a person first logs into an instance of an app. This is also known as App-Scoped ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


