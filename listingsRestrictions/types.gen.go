// Package listingsRestrictions provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package listingsRestrictions

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// Link defines model for Link.
type Link struct {

	// The URI of the related resource.
	Resource string `json:"resource"`

	// The title of the related resource.
	Title *string `json:"title,omitempty"`

	// The media type of the related resource.
	Type *string `json:"type,omitempty"`

	// The HTTP verb used to interact with the related resource.
	Verb string `json:"verb"`
}

// Reason defines model for Reason.
type Reason struct {

	// A list of path forward links that may allow Selling Partners to remove the restriction.
	Links *[]Link `json:"links,omitempty"`

	// A message describing the reason for the restriction.
	Message string `json:"message"`

	// A code indicating why the listing is restricted.
	ReasonCode *string `json:"reasonCode,omitempty"`
}

// Restriction defines model for Restriction.
type Restriction struct {

	// The condition that applies to the restriction.
	ConditionType *string `json:"conditionType,omitempty"`

	// A marketplace identifier. Identifies the Amazon marketplace where the restriction is enforced.
	MarketplaceId string `json:"marketplaceId"`

	// A list of reasons for the restriction.
	Reasons *[]Reason `json:"reasons,omitempty"`
}

// RestrictionList defines model for RestrictionList.
type RestrictionList struct {
	Restrictions []Restriction `json:"restrictions"`
}

// GetListingsRestrictionsParams defines parameters for GetListingsRestrictions.
type GetListingsRestrictionsParams struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	Asin string `json:"asin"`

	// The condition used to filter restrictions.
	ConditionType *string `json:"conditionType,omitempty"`

	// A selling partner identifier, such as a merchant account.
	SellerId string `json:"sellerId"`

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A locale for reason text localization. When not provided, the default language code of the first marketplace is used. Examples: "en_US", "fr_CA", "fr_FR". Localized messages default to "en_US" when a localization is not available in the specified locale.
	ReasonLocale *string `json:"reasonLocale,omitempty"`
}
