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

type ResponseErrorError struct {
	Code string `json:"code,omitempty"`
	Messages string `json:"messages,omitempty"`
	Type_ string `json:"type,omitempty"`
	FbtraceId string `json:"fbtrace_id,omitempty"`
	ErrorSubcode string `json:"error_subcode,omitempty"`
	IsTransient string `json:"is_transient,omitempty"`
	ErrorUserTitle string `json:"error_user_title,omitempty"`
	ErrorUserMsg string `json:"error_user_msg,omitempty"`
}