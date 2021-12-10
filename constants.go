package shopping

const EbayShoppingAPIVersion = "1199"
const EbayRequestDataFormat = "XML"
const EbayResponseDataFormat = "XML"
const DefaultItemsPerPage = 100

type EbayEndpoint string

const (
	// EbayEndpointProduction is a production endpoint for Shopping API
	EbayEndpointProduction = "https://open.api.ebay.com/shopping"
	// EbayEndpointSandbox is a sandbox endpoint for Shopping API
	EbayEndpointSandbox = "https://open.api.sandbox.ebay.com/shopping"
)

type EbayOperation string

const (
	OperationFindProducts     EbayOperation = "FindProducts"
	OperationGetCategoryInfo  EbayOperation = "GetCategoryInfo"
	OperationGeteBayTime      EbayOperation = "GeteBayTime"
	OperationGetItemStatus    EbayOperation = "GetItemStatus"
	OperationGetMultipleItems EbayOperation = "GetMultipleItems"
	OperationGetShippingCosts EbayOperation = "GetShippingCosts"
	OperationGetSingleItem    EbayOperation = "GetSingleItem"
	OperationGetUserProfile   EbayOperation = "GetUserProfile"
)

type SiteID string

const (
	SiteIDEbayUS    SiteID = "0"
	SiteIDEbayENCA  SiteID = "2"
	SiteIDEbayGB    SiteID = "3"
	SiteIDEbayAU    SiteID = "15"
	SiteIDEbayAT    SiteID = "16"
	SiteIDEbayFRBE  SiteID = "23"
	SiteIDEbayFR    SiteID = "71"
	SiteIDEbayDE    SiteID = "77"
	SiteIDEbayMOTOR SiteID = "100"
	SiteIDEbayIT    SiteID = "101"
	SiteIDEbayNLBE  SiteID = "123"
	SiteIDEbayNL    SiteID = "146"
	SiteIDEbayES    SiteID = "186"
	SiteIDEbayCH    SiteID = "193"
	SiteIDEbayHK    SiteID = "201"
	SiteIDEbayIN    SiteID = "203"
	SiteIDEbayIE    SiteID = "205"
	SiteIDEbayMY    SiteID = "207"
	SiteIDEbayFRCA  SiteID = "210"
	SiteIDEbayPH    SiteID = "211"
	SiteIDEbaySG    SiteID = "216"
	SiteIDEbayPL    SiteID = "212"
	SiteIDEbayRU    SiteID = "215"
)

type ProductIDCodeTypeOption string

const (
	// ProductIDCodeTypeEAN indicates the product identifier type is an International Article Number
	// (also known as European Article Number). EAN values are typically 13 digits in length,
	// but some use only eight digits. EAN identify a wide variety of products.
	ProductIDCodeTypeEAN ProductIDCodeTypeOption = "EAN"

	// ProductIDCodeTypeISBN indicates the product identifier type is an International Standard Book
	// Number. ISBN values can be 10 characters (ISBN-10) or 13 characters (ISBN-13) in length,
	// and they identify books.
	ProductIDCodeTypeISBN ProductIDCodeTypeOption = "ISBN"

	// ProductIDCodeTypeMPN indicates the product identifier type is a Manufacturer Part Number.
	// Unlike ISBNs, EANs, and UPCs, an MPN value is not based on an international standard, but
	// is defined by the seller/manufacturer of the product. Technically, there is no maximum length
	// for an MPN, but eBay actually enforces a 65-character limit for MPN values.
	ProductIDCodeTypeMPN ProductIDCodeTypeOption = "MPN"

	// ProductIDCodeTypeReference indicates the product identifier type is an eBay Catalog product ID.
	// EPID is a commonly-used acronymn for an eBay Catalog product ID.
	ProductIDCodeTypeReference ProductIDCodeTypeOption = "Reference"

	// ProductIDCodeTypeUPC indicates the product identifier type is a Universal Product Code.
	// UPC values are 12 digits in length. UPC values identify a wide variety of products,
	// and are typically used in the US and Canada.
	ProductIDCodeTypeUPC ProductIDCodeTypeOption = "UPC"
)

type ProductSortOption string

const (
	ProductSortItemCount   ProductSortOption = "ItemCount"
	ProductSortPopularity  ProductSortOption = "Popularity"
	ProductSortRating      ProductSortOption = "Rating"
	ProductSortReviewCount ProductSortOption = "ReviewCount"
	ProductSortTitle       ProductSortOption = "Title"
)

type SortOrderOption string

const (
	SortOrderAscending  SortOrderOption = "Ascending"
	SortOrderDescending SortOrderOption = "Descending"
)

type IncludeSelectorGCIOption string

const (
	IncludeSelectorChildCategories IncludeSelectorGCIOption = "ChildCategories"
)
