package shopping

// FindProductsResponse is response for FindProductsRequest
type FindProductsResponse struct {
	responseStandard
	ApproximatePages int `xml:"ApproximatePages"`
	//DomainHistogram DomainHistogram `xml:"DomainHistogram"`
	MoreResults   bool      `xml:"MoreResults"`
	PageNumber    int       `xml:"PageNumber"`
	TotalProducts int       `xml:"TotalProducts"`
	Products      []Product `xml:"Product"`
}

// Product is returned for each eBay catalog product that matches the search criteria.
// The Product container consists of specific data about the catalog product, including the product title,
// product identifiers (ePID and any GTIN value(s)), product aspects, a link to eBay product page,
// and links to stock photos (if any).
type Product struct {
	DetailsURL         string          `xml:"DetailsURL"`
	DisplayStockPhotos bool            `xml:"DisplayStockPhotos"`
	DomainName         string          `xml:"DomainName"`
	ItemSpecifics      []NameValueList `xml:"ItemSpecifics"`
	ProductIDs         []ProductID     `xml:"ProductID"`
	ProductState       string          `xml:"ProductState"`
	ReviewCount        int             `xml:"ReviewCount"`
	StockPhotoURL      string          `xml:"StockPhotoURL"`
	Title              string          `xml:"Title"`
}

// NameValueList is an array of StatusItem Specifics name-value pairs for an eBay Catalog product (if FindProducts is used)
// or StatusItem Specifics name-value pairs for a single-variation listing or individual variation within
// a multiple-variation listing (if GetSingleItem or GetMultipleItems is used).
type NameValueList struct {
	Name   string   `xml:"Name"`
	Values []string `xml:"Value"`
}

/*
===========================================================
*/

// GetCategoryInfoResponse represents response for GetCategoryInfoRequest
type GetCategoryInfoResponse struct {
	responseStandard
	CategoryArray   []Category `xml:"CategoryArray"`
	CategoryCount   int        `xml:"CategoryCount"`
	CategoryVersion string     `xml:"CategoryVersion"`
	UpdateTime      string     `xml:"UpdateTime"`
}

// Category consists of high-level details of a category, including its category ID value, full category path
// (by name and by category ID), its level in the eBay site's category hierarchy, category ID
// of its parent category, and a boolean value to indicate if it is a listing (leaf) category.
type Category struct {
	CategoryID       string `xml:"CategoryID"`
	CategoryIDPath   string `xml:"CategoryIDPath"`
	CategoryLevel    int    `xml:"CategoryLevel"`
	CategoryName     string `xml:"CategoryName"`
	CategoryNamePath string `xml:"CategoryNamePath"`
	CategoryParentID string `xml:"CategoryParentID"`
	LeafCategory     bool   `xml:"LeafCategory"`
}

/*
===========================================================
*/

//GeteBayTimeResponse is a response for GeteBayTimeRequest
type GeteBayTimeResponse struct {
	responseStandard
}

/*
===========================================================
*/

// GetItemStatusResponse is a response for GetItemStatusRequest
type GetItemStatusResponse struct {
	responseStandard
	Items []StatusItem `xml:"StatusItem"`
}

// StatusItem is returned for each ItemID value that was specified in the call request.
// One GetItemStatus call can retrieve up to 20 eBay listings.
type StatusItem struct {
	BidCount              int       `xml:"BidCount"`
	ConvertedCurrentPrice Price     `xml:"ConvertedCurrentPrice"`
	EndTime               string    `xml:"EndTime"`
	HighBidder            BasicUser `xml:"HighBidder"`
	ItemID                string    `xml:"ItemID"`
	ListingStatus         string    `xml:"ListingStatus"`
	TimeLeft              string    `xml:"TimeLeft"`
	ReserveMet            bool      `xml:"ReserveMet"`
	BuyItNowAvailable     bool      `xml:"BuyItNowAvailable"`
}

// Price ...
type Price struct {
	CurrencyID string  `xml:"currencyID,attr"`
	Value      float64 `xml:",cdata"`
}

// BasicUser is used to express the details for one eBay user.
type BasicUser struct {
	FeedbackPrivate    bool   `xml:"FeedbackPrivate"`
	FeedbackRatingStar string `xml:"FeedbackRatingStar"`
	FeedbackScore      int    `xml:"FeedbackScore"`
	UserID             string `xml:"UserID"`
}

/*
===========================================================
*/

type GetMultipleItemsResponse struct {
	responseStandard
}

/*
===========================================================
*/

// GetShippingCostsResponse is a response for GetShippingCostsRequest
type GetShippingCostsResponse struct {
	responseStandard
	PickUpInStoreDetails PickUpInStoreDetails `xml:"PickUpInStoreDetails"`
	ShippingCostSummary  ShippingCostSummary  `xml:"ShippingCostSummary"`
	ShippingDetails      ShippingDetails      `xml:"ShippingDetails"`
}

// PickUpInStoreDetails is only returned in GetShippingCosts if In-Store Pickup is available for the listing.
type PickUpInStoreDetails struct {
	AvailableForPickupInStore bool `xml:"AvailableForPickupInStore"`
	EligibleForPickupInStore  bool `xml:"EligibleForPickupInStore"`
}

// ShippingCostSummary returns a few details of the lowest-priced shipping service option that is
// available to the shipping destination specified in the call request.
type ShippingCostSummary struct {
	ImportCharge              Price  `xml:"ImportCharge"`
	InsuranceCost             Price  `xml:"InsuranceCost"`
	InsuranceOption           string `xml:"InsuranceOption"`
	ListedShippingServiceCost Price  `xml:"ListedShippingServiceCost"`
	ShippingServiceCost       Price  `xml:"ShippingServiceCost"`
	ShippingServiceName       string `xml:"ShippingServiceName"`
	ShippingType              string `xml:"ShippingType"`
}

// ShippingDetails consists of shipping details related to the specified item and specified shipping destination.
// This container is only returned if the IncludeDetails field is included and set to true in the call request.
type ShippingDetails struct {
	CODCost                             Price                   `xml:"CODCost"`
	ExcludeShipToLocations              []string                `xml:"ExcludeShipToLocation"`
	InsuranceCost                       Price                   `xml:"InsuranceCost"`
	InsuranceOption                     string                  `xml:"InsuranceOption"`
	InternationalInsuranceCost          Price                   `xml:"InternationalInsuranceCost"`
	InternationalInsuranceOption        string                  `xml:"InternationalInsuranceOption"`
	InternationalShippingServiceOptions []IntShipServiceOption  `xml:"InternationalShippingServiceOption"`
	SalesTax                            SalesTax                `xml:"SalesTax"`
	ShippingRateErrorMessage            string                  `xml:"ShippingRateErrorMessage"`
	ShippingServiceOptions              []ShippingServiceOption `xml:"ShippingServiceOption"`
	TaxTable                            TaxTable                `xml:"TaxTable"`
}

// IntShipServiceOption consists of detailed information for an international shipping service option that is
// available to an international buyer located at the shipping destination specified in the call request.
type IntShipServiceOption struct {
	EstimatedDeliveryMaxTime      string   `xml:"EstimatedDeliveryMaxTime"`
	EstimatedDeliveryMinTime      string   `xml:"EstimatedDeliveryMinTime"`
	ImportCharge                  Price    `xml:"ImportCharge"`
	ShippingServiceAdditionalCost Price    `xml:"ShippingServiceAdditionalCost"`
	ShippingServiceCost           Price    `xml:"ShippingServiceCost"`
	ShippingServiceCutOffTime     string   `xml:"ShippingServiceCutOffTime"`
	ShippingServiceName           string   `xml:"ShippingServiceName"`
	ShippingServicePriority       int      `xml:"ShippingServicePriority"`
	ShipsTo                       []string `xml:"ShipsTo"`
}

// SalesTax is used to express sales tax details for the shipping destination.
type SalesTax struct {
	SalesTaxAmount        Price   `xml:"SalesTaxAmount"`
	SalesTaxPercent       float64 `xml:"SalesTaxPercent"`
	SalesTaxState         string  `xml:"SalesTaxState"`
	ShippingIncludedInTax bool    `xml:"ShippingIncludedInTax"`
}

// ShippingServiceOption  consists of detailed information for a domestic shipping
// service option that is available to a buyer located at the shipping destination specified in
// the call request. A ShippingServiceOption container is returned for each available domestic shipping
// service option. A seller can specify up to four domestic shipping service options in an eBay listing.
type ShippingServiceOption struct {
	EstimatedDeliveryMaxTime      string   `xml:"EstimatedDeliveryMaxTime"`
	EstimatedDeliveryMinTime      string   `xml:"EstimatedDeliveryMinTime"`
	ExpeditedService              bool     `xml:"ExpeditedService"`
	FastAndFree                   bool     `xml:"FastAndFree"`
	LogisticPlanType              string   `xml:"LogisticPlanType"`
	ShippingInsuranceCost         Price    `xml:"ShippingInsuranceCost"`
	ShippingServiceAdditionalCost Price    `xml:"ShippingServiceAdditionalCost"`
	ShippingServiceCost           Price    `xml:"ShippingServiceCost"`
	ShippingServiceCutOffTime     string   `xml:"ShippingServiceCutOffTime"`
	ShippingServiceName           string   `xml:"ShippingServiceName"`
	ShippingServicePriority       int      `xml:"ShippingServicePriority"`
	ShippingSurcharge             Price    `xml:"ShippingSurcharge"`
	ShippingTimeMax               int      `xml:"ShippingTimeMax"`
	ShippingTimeMin               int      `xml:"ShippingTimeMin"`
	ShipsTo                       []string `xml:"ShipsTo"`
}

// TaxTable consists of an array of TaxJurisdiction containers; one returned for each tax jurisdiction where
// the seller has defined a sales tax rate (using the Sales Tax Table UI in My eBay).
// The TaxTable container is returned as an empty element if no sales tax rates have been set up for
// any jurisdictions.
type TaxTable struct {
	TaxJurisdictions []TaxJurisdiction `xml:"TaxJurisdiction"`
}

// TaxJurisdiction ...
type TaxJurisdiction struct {
	JurisdictionID        string  `xml:"JurisdictionID"`
	SalesTaxPercent       float64 `xml:"SalesTaxPercent"`
	ShippingIncludedInTax bool    `xml:"ShippingIncludedInTax"`
}

/*
===========================================================
*/

type GetSingleItemResponse struct {
	responseStandard
}

/*
===========================================================
*/

type GetUserProfileResponse struct {
	responseStandard
}
