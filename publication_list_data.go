/* 
 * Medium API
 *
 * Articles that matter on social publishing platform
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package medium

type PublicationListData struct {

	// a list of publication objects within a data envelope
	Data []InlineResponse2001Data `json:"data,omitempty"`
}
