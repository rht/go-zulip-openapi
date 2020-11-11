/*
 * Zulip REST API
 *
 * Powerful open source group chat 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InvalidMessageError struct for InvalidMessageError
type InvalidMessageError struct {
	Result map[string]interface{} `json:"result"`
	Msg map[string]interface{} `json:"msg"`
	// The raw content of the message. 
	RawContent string `json:"raw_content,omitempty"`
}
