// Package catalog provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package catalog

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// BrandRefinement defines model for BrandRefinement.
type BrandRefinement struct {

	// Brand name. For display and can be used as a search refinement.
	BrandName string `json:"brandName"`

	// The estimated number of results that would still be returned if refinement key applied.
	NumberOfResults int `json:"numberOfResults"`
}

// ClassificationRefinement defines model for ClassificationRefinement.
type ClassificationRefinement struct {

	// Identifier for the classification that can be used for search refinement purposes.
	ClassificationId string `json:"classificationId"`

	// Display name for the classification.
	DisplayName string `json:"displayName"`

	// The estimated number of results that would still be returned if refinement key applied.
	NumberOfResults int `json:"numberOfResults"`
}

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
type ErrorList struct {
	Errors []Error `json:"errors"`
}

// Item defines model for Item.
type Item struct {

	// Amazon Standard Identification Number (ASIN) is the unique identifier for an item in the Amazon catalog.
	Asin ItemAsin `json:"asin"`

	// A JSON object that contains structured item attribute data keyed by attribute name. Catalog item attributes are available only to brand owners and conform to the related product type definitions available in the Selling Partner API for Product Type Definitions.
	Attributes *ItemAttributes `json:"attributes,omitempty"`

	// Identifiers associated with the item in the Amazon catalog, such as UPC and EAN identifiers.
	Identifiers *ItemIdentifiers `json:"identifiers,omitempty"`

	// Images for an item in the Amazon catalog. All image variants are provided to brand owners. Otherwise, a thumbnail of the "MAIN" image variant is provided.
	Images *ItemImages `json:"images,omitempty"`

	// Product types associated with the Amazon catalog item.
	ProductTypes *ItemProductTypes `json:"productTypes,omitempty"`

	// Sales ranks of an Amazon catalog item.
	SalesRanks *ItemSalesRanks `json:"salesRanks,omitempty"`

	// Summary details of an Amazon catalog item.
	Summaries *ItemSummaries `json:"summaries,omitempty"`

	// Variation details by marketplace for an Amazon catalog item (variation relationships).
	Variations *ItemVariations `json:"variations,omitempty"`

	// Vendor details associated with an Amazon catalog item. Vendor details are available to vendors only.
	VendorDetails *ItemVendorDetails `json:"vendorDetails,omitempty"`
}

// ItemAsin defines model for ItemAsin.
type ItemAsin string

// ItemAttributes defines model for ItemAttributes.
type ItemAttributes struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ItemIdentifier defines model for ItemIdentifier.
type ItemIdentifier struct {

	// Identifier.
	Identifier string `json:"identifier"`

	// Type of identifier, such as UPC, EAN, or ISBN.
	IdentifierType string `json:"identifierType"`
}

// ItemIdentifiers defines model for ItemIdentifiers.
type ItemIdentifiers []ItemIdentifiersByMarketplace

// ItemIdentifiersByMarketplace defines model for ItemIdentifiersByMarketplace.
type ItemIdentifiersByMarketplace struct {

	// Identifiers associated with the item in the Amazon catalog for the indicated Amazon marketplace.
	Identifiers []ItemIdentifier `json:"identifiers"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

// ItemImage defines model for ItemImage.
type ItemImage struct {

	// Height of the image in pixels.
	Height int `json:"height"`

	// Link, or URL, for the image.
	Link string `json:"link"`

	// Variant of the image, such as MAIN or PT01.
	Variant string `json:"variant"`

	// Width of the image in pixels.
	Width int `json:"width"`
}

// ItemImages defines model for ItemImages.
type ItemImages []ItemImagesByMarketplace

// ItemImagesByMarketplace defines model for ItemImagesByMarketplace.
type ItemImagesByMarketplace struct {

	// Images for an item in the Amazon catalog for the indicated Amazon marketplace.
	Images []ItemImage `json:"images"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

// ItemProductTypeByMarketplace defines model for ItemProductTypeByMarketplace.
type ItemProductTypeByMarketplace struct {

	// Amazon marketplace identifier.
	MarketplaceId *string `json:"marketplaceId,omitempty"`

	// Name of the product type associated with the Amazon catalog item.
	ProductType *string `json:"productType,omitempty"`
}

// ItemProductTypes defines model for ItemProductTypes.
type ItemProductTypes []ItemProductTypeByMarketplace

// ItemSalesRank defines model for ItemSalesRank.
type ItemSalesRank struct {

	// Corresponding Amazon retail website link, or URL, for the sales rank.
	Link *string `json:"link,omitempty"`

	// Sales rank value.
	Rank int `json:"rank"`

	// Title, or name, of the sales rank.
	Title string `json:"title"`
}

// ItemSalesRanks defines model for ItemSalesRanks.
type ItemSalesRanks []ItemSalesRanksByMarketplace

// ItemSalesRanksByMarketplace defines model for ItemSalesRanksByMarketplace.
type ItemSalesRanksByMarketplace struct {

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Sales ranks of an Amazon catalog item for an Amazon marketplace.
	Ranks []ItemSalesRank `json:"ranks"`
}

// ItemSearchResults defines model for ItemSearchResults.
type ItemSearchResults struct {

	// A list of items from the Amazon catalog.
	Items []Item `json:"items"`

	// The estimated total number of products matched by the search query (only results up to the page count limit will be returned per request regardless of the number found).
	//
	// Note: The maximum number of items (ASINs) that can be returned and paged through is 1000.
	NumberOfResults int `json:"numberOfResults"`

	// When a request produces a response that exceeds the pageSize, pagination occurs. This means the response is divided into individual pages. To retrieve the next page or the previous page, you must pass the nextToken value or the previousToken value as the pageToken parameter in the next request. When you receive the last page, there will be no nextToken key in the pagination object.
	Pagination Pagination `json:"pagination"`

	// Search refinements.
	Refinements Refinements `json:"refinements"`
}

// ItemSummaries defines model for ItemSummaries.
type ItemSummaries []ItemSummaryByMarketplace

// ItemSummaryByMarketplace defines model for ItemSummaryByMarketplace.
type ItemSummaryByMarketplace struct {

	// Name of the brand associated with an Amazon catalog item.
	BrandName *string `json:"brandName,omitempty"`

	// Identifier of the browse node associated with an Amazon catalog item.
	BrowseNode *string `json:"browseNode,omitempty"`

	// Name of the color associated with an Amazon catalog item.
	ColorName *string `json:"colorName,omitempty"`

	// Name, or title, associated with an Amazon catalog item.
	ItemName *string `json:"itemName,omitempty"`

	// Name of the manufacturer associated with an Amazon catalog item.
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Model number associated with an Amazon catalog item.
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Name of the size associated with an Amazon catalog item.
	SizeName *string `json:"sizeName,omitempty"`

	// Name of the style associated with an Amazon catalog item.
	StyleName *string `json:"styleName,omitempty"`
}

// ItemVariations defines model for ItemVariations.
type ItemVariations []ItemVariationsByMarketplace

// ItemVariationsByMarketplace defines model for ItemVariationsByMarketplace.
type ItemVariationsByMarketplace struct {

	// Identifiers (ASINs) of the related items.
	Asins []string `json:"asins"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Type of variation relationship of the Amazon catalog item in the request to the related item(s): "PARENT" or "CHILD".
	VariationType string `json:"variationType"`
}

// ItemVendorDetails defines model for ItemVendorDetails.
type ItemVendorDetails []ItemVendorDetailsByMarketplace

// ItemVendorDetailsByMarketplace defines model for ItemVendorDetailsByMarketplace.
type ItemVendorDetailsByMarketplace struct {

	// Brand code associated with an Amazon catalog item.
	BrandCode *string `json:"brandCode,omitempty"`

	// Product category associated with an Amazon catalog item.
	CategoryCode *string `json:"categoryCode,omitempty"`

	// Manufacturer code associated with an Amazon catalog item.
	ManufacturerCode *string `json:"manufacturerCode,omitempty"`

	// Parent vendor code of the manufacturer code.
	ManufacturerCodeParent *string `json:"manufacturerCodeParent,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Product group associated with an Amazon catalog item.
	ProductGroup *string `json:"productGroup,omitempty"`

	// Replenishment category associated with an Amazon catalog item.
	ReplenishmentCategory *string `json:"replenishmentCategory,omitempty"`

	// Product subcategory associated with an Amazon catalog item.
	SubcategoryCode *string `json:"subcategoryCode,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {

	// A token that can be used to fetch the next page.
	NextToken *string `json:"nextToken,omitempty"`

	// A token that can be used to fetch the previous page.
	PreviousToken *string `json:"previousToken,omitempty"`
}

// Refinements defines model for Refinements.
type Refinements struct {

	// Brand search refinements.
	Brands []BrandRefinement `json:"brands"`

	// Classification search refinements.
	Classifications []ClassificationRefinement `json:"classifications"`
}

// SearchCatalogItemsParams defines parameters for SearchCatalogItems.
type SearchCatalogItemsParams struct {

	// A comma-delimited list of words or item identifiers to search the Amazon catalog for.
	Keywords []string `json:"keywords"`

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: summaries.
	IncludedData *[]string `json:"includedData,omitempty"`

	// A comma-delimited list of brand names to limit the search to.
	BrandNames *[]string `json:"brandNames,omitempty"`

	// A comma-delimited list of classification identifiers to limit the search to.
	ClassificationIds *[]string `json:"classificationIds,omitempty"`

	// Number of results to be returned per page.
	PageSize *int `json:"pageSize,omitempty"`

	// A token to fetch a certain page when there are multiple pages worth of results.
	PageToken *string `json:"pageToken,omitempty"`

	// The language the keywords are provided in. Defaults to the primary locale of the marketplace.
	KeywordsLocale *string `json:"keywordsLocale,omitempty"`

	// Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace.
	Locale *string `json:"locale,omitempty"`
}

// GetCatalogItemParams defines parameters for GetCatalogItem.
type GetCatalogItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers. Data sets in the response contain data only for the specified marketplaces.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: summaries.
	IncludedData *[]string `json:"includedData,omitempty"`

	// Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace.
	Locale *string `json:"locale,omitempty"`
}

// Getter for additional properties for ItemAttributes. Returns the specified
// element and whether it was found
func (a ItemAttributes) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ItemAttributes
func (a *ItemAttributes) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a *ItemAttributes) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a ItemAttributes) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
