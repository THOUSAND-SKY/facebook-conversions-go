/*
 * Facebook Conversions API (for Web)
 *
 * The Conversions API (for web) allows advertisers to send web events from their servers directly to Facebook. Conversions API events are linked to a pixel and are processed like browser pixel events. This means that these conversion events are used in measurement, reporting, and optimization in the same way as browser pixel events.
 *
 * API version: 1.0.0
 * Contact: web_signals_integrations@fb.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package facebookgen

// user_data is a set of identifiers Facebook can use for targeted attribution. You must provide at least one of the following user_data keys in your request
type UserData struct {
	// A hashed email address in lower case using SHA-256 algorithm.
	Em string `json:"em,omitempty"`
	// A hashed phone number using SHA-256 algorithm. Include only digits with country code, area code, and number.
	Ph string `json:"ph,omitempty"`
	// A hashed gender (f or m) using SHA-256 algorithm.
	Ge string `json:"ge,omitempty"`
	// A hashed date of birth given as year, month, and day using SHA-256 algorithm
	Db string `json:"db,omitempty"`
	// A hashed last name in lowercase using SHA-256 algorithm.
	Ln string `json:"ln,omitempty"`
	// A hashed first name in lowercase using SHA-256 algorithm.
	Fn string `json:"fn,omitempty"`
	// A hashed city in lower-case without spaces or punctuation using SHA-256 algorithm.
	Ct string `json:"ct,omitempty"`
	// A hashed two-letter country code in lowercase using SHA-256 algorithm.
	Country string `json:"country,omitempty"`
	// A hashed two-letter state code in lowercase using SHA-256 algorithm.
	St string `json:"st,omitempty"`
	// A hashed zip code using SHA-256 algorithm. If you are in the United States, this is a five-digit zip code. For other locations, follow each country's standards.
	Zp string `json:"zp,omitempty"`
	// Any unique ID from the advertiser, such as loyalty membership IDs, user IDs, and external cookie IDs. If External ID is being sent via other channels, it should be sent in the same format via Conversions API. Hashing external_id using SHA-256 algorithm is optional.
	ExternalId string `json:"external_id,omitempty"`
	// The IP address of the browser corresponding to the event.
	ClientIpAddress string `json:"client_ip_address,omitempty"`
	// The user agent for the browser corresponding to the event.
	ClientUserAgent string `json:"client_user_agent,omitempty"`
	// The Facebook click ID value stored in the _fbc browser cookie under your domain. See Managing fbc and fbp Parameters for how to get this value (https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/fbp-and-fbc), or generate this value from a fbclid query parameter.
	Fbc string `json:"fbc,omitempty"`
	// The Facebook browser ID value stored in the _fbp browser cookie under your domain. See Managing fbc and fbp Parameters for how to get this value (https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/fbp-and-fbc).
	Fbp string `json:"fbp,omitempty"`
	// The subscription ID for the user in this transaction. This is similar to the order ID for an individual product.
	SubscriptionId string `json:"subscription_id,omitempty"`
}