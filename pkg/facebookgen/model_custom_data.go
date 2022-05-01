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

// An object that includes additional business data about the event which can be used for ads delivery optimization. If our predefined object properties don't suit your needs, you can include your own, custom properties. Custom properties can be used with both standard and custom events, and can help you further define custom audiences.
type CustomData struct {
	// A numeric value associated with this event. This could be a monetary value or a value in some other metric.
	Value float32 `json:"value,omitempty"`
	// The currency for the value specified, if applicable. Currency must be a valid ISO 4217 (https://en.wikipedia.org/wiki/ISO_4217) three digit currency code.
	Currency string `json:"currency,omitempty"`
	// The name of the page or product associated with the event.
	ContentName string `json:"content_name,omitempty"`
	// The category of the content associated with the event.
	ContentCategory string `json:"content_category,omitempty"`
	// The content IDs associated with the event, such as product SKUs for items in an AddToCart event: ['ABC123', 'XYZ789']. If content_type is a product, then your content IDs must be an array with a single string value. Otherwise, this array can contain any number of string values.
	ContentIds []string `json:"content_ids,omitempty"`
	// A list of Content objects that contain the product IDs associated with the event plus information about the products. id, quantity, and item_price are available fields.
	Contents []Content `json:"contents,omitempty"`
	// It should be set to 'product' or 'product_group'. Use 'product', if the keys you send represent products. Sent keys could be content_ids or contents. Use product_group, if the keys you send in content_ids represent product groups. Product groups are used to distinguish products that are identical but have variations such as color, material, size or pattern.
	ContentType string `json:"content_type,omitempty"`
	// The order ID for this transaction as a String.
	OrderId string `json:"order_id,omitempty"`
	// The predicted lifetime value of a conversion event, as a String.
	PredictedLtv float32 `json:"predicted_ltv,omitempty"`
	// Use only with InitiateCheckout events. The number of items that a user tries to buy during checkout.
	NumItems int32 `json:"num_items,omitempty"`
	// Use only with Search events. A search query made by a user.
	SearchString string `json:"search_string,omitempty"`
	// Use only with CompleteRegistration events. The status of the registration event
	Status string `json:"status,omitempty"`
}