/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SenderSupplied
type SenderSupplied struct {
	// FormatVersion 30
	FormatVersion string `json:"formatVersion"`
	// UserRequestCorrelation
	UserRequestCorrelation string `json:"userRequestCorrelation"`
	// Identifies if test or production.  * `T` - Test * `P` - Production
	TestProductionCode string `json:"testProductionCode"`
	// MessageDuplicationCode  * ` ` - Original Message * `R` - Retrieval of an original message * `P` - Resend
	MessageDuplicationCode string `json:"messageDuplicationCode"`
}
