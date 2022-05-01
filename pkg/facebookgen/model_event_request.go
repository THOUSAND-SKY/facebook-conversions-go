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

type EventRequest struct {
	// An array of Server Event objects.
	Data []Event `json:"data"`
	// Code used to verify that your server events are received correctly by Facebook. Use this code to test your server events in the Test Events feature in Events Manager. See Test Events Tool (https://developers.facebook.com/docs/marketing-api/conversions-api/using-the-api#testEvents) for an example.
	TestEventCode string `json:"test_event_code,omitempty"`
	// Partner agent string.
	PartnerAgent string `json:"partner_agent"`
}