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
	ItemSpecifics      []NameValueList `xml:"ItemSpecifics>NameValueList"`
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

// GetMultipleItemsResponse is a response for GetMultipleItemsRequest
type GetMultipleItemsResponse struct {
	responseStandard
	Items []Item `xml:"Item"`
}

// Item contains details about the listing whose ID was specified in the request.
type Item struct {
	AutoPay                             bool                    `xml:"AutoPay"`
	AvailableForPickupDropOff           bool                    `xml:"AvailableForPickupDropOff"`
	BestOfferEnabled                    bool                    `xml:"BestOfferEnabled"`
	BuyItNowAvailable                   bool                    `xml:"BuyItNowAvailable"`
	EligibleForPickupDropOff            bool                    `xml:"EligibleForPickupDropOff"`
	GlobalShipping                      bool                    `xml:"GlobalShipping"`
	IgnoreQuantity                      bool                    `xml:"IgnoreQuantity"`
	IntegratedMerchantCreditCardEnabled bool                    `xml:"IntegratedMerchantCreditCardEnabled"`
	BidCount                            int                     `xml:"BidCount"`
	BusinessSellerDetails               BusinessSellerDetails   `xml:"BusinessSellerDetails"`
	BuyItNowPrice                       Price                   `xml:"BuyItNowPrice"`
	Charity                             Charity                 `xml:"Charity"`
	ConditionDescription                string                  `xml:"ConditionDescription"`
	ConditionDisplayName                string                  `xml:"ConditionDisplayName"`
	ConditionID                         int                     `xml:"ConditionID"`
	ConvertedBuyItNowPrice              Price                   `xml:"ConvertedBuyItNowPrice"`
	ConvertedCurrentPrice               Price                   `xml:"ConvertedCurrentPrice"`
	Country                             string                  `xml:"Country"`
	CurrentPrice                        Price                   `xml:"CurrentPrice"`
	Description                         string                  `xml:"Description"`
	DiscountPriceInfo                   DiscountPriceInfo       `xml:"DiscountPriceInfo"`
	EndTime                             string                  `xml:"EndTime"`
	ExcludeShipToLocations              []string                `xml:"ExcludeShipToLocation"`
	GalleryURL                          string                  `xml:"GalleryURL"`
	HandlingTime                        int                     `xml:"HandlingTime"`
	HighBidder                          User                    `xml:"HighBidder"`
	HitCount                            int64                   `xml:"HitCount"`
	ItemID                              string                  `xml:"ItemID"`
	ItemSpecifics                       []NameValueList         `xml:"ItemSpecifics>NameValueList"`
	ListingStatus                       string                  `xml:"ListingStatus"`
	ListingType                         string                  `xml:"ListingType"`
	Location                            string                  `xml:"Location"`
	LotSize                             int                     `xml:"LotSize"`
	MinimumToBid                        Price                   `xml:"MinimumToBid"`
	PaymentAllowedSites                 []string                `xml:"PaymentAllowedSite"`
	PaymentMethods                      []string                `xml:"PaymentMethods"`
	PictureURLs                         []string                `xml:"PictureURL"`
	PostalCode                          string                  `xml:"PostalCode"`
	PrimaryCategoryID                   string                  `xml:"PrimaryCategoryID"`
	PrimaryCategoryIDPath               string                  `xml:"PrimaryCategoryIDPath"`
	PrimaryCategoryName                 string                  `xml:"PrimaryCategoryName"`
	ProductID                           string                  `xml:"ProductID"`
	Quantity                            int                     `xml:"Quantity"`
	QuantityAvailableHint               string                  `xml:"QuantityAvailableHint"`
	MinimumRemnantSet                   int                     `xml:"QuantityInfo>MinimumRemnantSet"`
	QuantitySold                        int                     `xml:"QuantitySold"`
	QuantitySoldByPickupInStore         int                     `xml:"QuantitySoldByPickupInStore"`
	QuantityThreshold                   int                     `xml:"QuantityThreshold"`
	ReturnPolicy                        ReturnPolicy            `xml:"ReturnPolicy"`
	SecondaryCategoryID                 string                  `xml:"SecondaryCategoryID"`
	SecondaryCategoryIDPath             string                  `xml:"SecondaryCategoryIDPath"`
	SecondaryCategoryName               string                  `xml:"SecondaryCategoryName"`
	Seller                              Seller                  `xml:"Seller"`
	ShippingCostSummary                 ItemShippingCostSummary `xml:"ShippingCostSummary"`
	ShipToLocations                     []string                `xml:"ShipToLocations"`
	Site                                string                  `xml:"Site"`
	SKU                                 string                  `xml:"SKU"`
	StartTime                           string                  `xml:"StartTime"`
	Storefront                          Storefront              `xml:"Storefront"`
	Subtitle                            string                  `xml:"Subtitle"`
	TimeLeft                            string                  `xml:"TimeLeft"`
	Title                               string                  `xml:"Title"`
	ReserveMet                          bool                    `xml:"ReserveMet"`
	TopRatedListing                     bool                    `xml:"TopRatedListing"`
	UnitInfo                            UnitInfo                `xml:"UnitInfo"`
	Variations                          Variations              `xml:"Variations"`
	VhrAvailable                        string                  `xml:"VhrAvailable"`
	VhrUrl                              string                  `xml:"VhrUrl"`
	ViewItemURLForNaturalSearch         string                  `xml:"ViewItemURLForNaturalSearch"`
}

// BusinessSellerDetails  is returned if the seller of the item is registered on the eBay listing site as a
// Business Seller. This container consists of information related to the Business Seller's account.
// Not all eBay sites support Business Sellers.
type BusinessSellerDetails struct {
	AdditionalContactInformation string     `xml:"AdditionalContactInformation"`
	Address                      Address    `xml:"Address"`
	Email                        string     `xml:"Email"`
	Fax                          string     `xml:"Fax"`
	LegalInvoice                 bool       `xml:"LegalInvoice"`
	TermsAndConditions           string     `xml:"TermsAndConditions"`
	TradeRegistrationNumber      string     `xml:"TradeRegistrationNumber"`
	VATDetails                   VATDetails `xml:"VATDetails"`
}

// Address is used to provide details about a Business Seller's address.
type Address struct {
	CityName        string `xml:"CityName"`
	CompanyName     string `xml:"CompanyName"`
	CountryName     string `xml:"CountryName"`
	FirstName       string `xml:"FirstName"`
	LastName        string `xml:"LastName"`
	Name            string `xml:"Name"`
	Phone           string `xml:"Phone"`
	PostalCode      string `xml:"PostalCode"`
	StateOrProvince string `xml:"StateOrProvince"`
	Street1         string `xml:"Street1"`
	Street2         string `xml:"Street2"`
}

// VATDetails provides Value-Added Tax (VAT) details for the Business Seller, including the seller's VAT ID
// and the VAT percentage rate applicable to the item. VAT is similar to a sales and/or consumption tax,
// and it is only applicable to sellers selling on European sites.
type VATDetails struct {
	BusinessSeller       bool    `xml:"BusinessSeller"`
	RestrictedToBusiness bool    `xml:"RestrictedToBusiness"`
	VATID                string  `xml:"VATID"`
	VATPercent           float64 `xml:"VATPercent"`
	VATSite              string  `xml:"VATSite"`
}

// Charity is returned if any percentage of the sales proceeds is going to a nonprofit organization that is
// registered with eBay for Charity. This container consists of details related to the nonprofit charity
// organization, including the name, mission, and unique identifier of the charity, as well as the percentage
// rate of the sale proceeds that will go to the charity for each sale.
type Charity struct {
	CharityID       string  `xml:"CharityID"`
	CharityName     string  `xml:"CharityName"`
	CharityNumber   int     `xml:"CharityNumber"`
	DonationPercent float64 `xml:"DonationPercent"`
	LogoURL         string  `xml:"LogoURL"`
	Mission         string  `xml:"Mission"`
	Status          string  `xml:"Status"`
}

// DiscountPriceInfo provides information for an item that has a Strikethrough Price (STP)
// or a Minimum Advertised Price (MAP) discount pricing treatment. STP and MAP apply only to fixed-price listings.
// STP is available on the US, eBay Motors, UK, Germany, Canada (English and French), France,
// Italy, and Spain sites, while MAP is available only on the US site.
type DiscountPriceInfo struct {
	MinimumAdvertisedPrice         Price  `xml:"MinimumAdvertisedPrice"`
	MinimumAdvertisedPriceExposure string `xml:"MinimumAdvertisedPriceExposure"`
	OriginalRetailPrice            Price  `xml:"OriginalRetailPrice"`
	PricingTreatment               string `xml:"PricingTreatment"`
	SoldOffeBay                    bool   `xml:"SoldOffeBay"`
	SoldOneBay                     bool   `xml:"SoldOneBay"`
}

// User ...
type User struct {
	BasicUser
	UserAnonymized bool `xml:"UserAnonymized"`
}

// ReturnPolicy consists of details related to the seller's Return Policy, both for domestic and international
// buyers (if the seller ships internationally). If the seller does not accept returns, only the
// ReturnsAccepted field (or InternationalReturnsAccepted field for international buyers) is
// returned with a value of ReturnsNotAccepted.
type ReturnPolicy struct {
	Description                     string `xml:"Description"`
	InternationalRefund             string `xml:"InternationalRefund"`
	InternationalReturnsAccepted    string `xml:"InternationalReturnsAccepted"`
	InternationalReturnsWithin      string `xml:"InternationalReturnsWithin"`
	InternationalShippingCostPaidBy string `xml:"InternationalShippingCostPaidBy"`
	Refund                          string `xml:"Refund"`
	ReturnsAccepted                 string `xml:"ReturnsAccepted"`
	ReturnsWithin                   string `xml:"ReturnsWithin"`
	ShippingCostPaidBy              string `xml:"ShippingCostPaidBy"`
}

// Seller ...
type Seller struct {
	BasicUser
	TopRatedSeller bool `xml:"TopRatedSeller"`
}

// ItemShippingCostSummary returns a few details of the lowest-priced shipping service option that is available
// to the eBay user making the call. For Calculated shipping, the item's location and the destination location
// are considered when calculating the shipping cost.
type ItemShippingCostSummary struct {
	ListedShippingServiceCost float64 `xml:"ListedShippingServiceCost"`
	LocalPickup               bool    `xml:"LocalPickup"`
	ShippingServiceCost       float64 `xml:"ShippingServiceCost"`
	ShippingType              string  `xml:"ShippingType"`
}

// Storefront consists of the eBay seller's store name and the URL to the eBay store. This container
// is returned if the seller has an eBay Store subscription and the IncludeSelector field is included in
// the call request and set to Details.
type Storefront struct {
	StoreName string `xml:"StoreName"`
	StoreURL  string `xml:"StoreURL"`
}

// UnitInfo contains information about the weight, volume or other quantity measurement of a listed item so
// buyers can compare per-unit prices. The European Union requires listings for certain types of products
// to include the price per unit. eBay uses this information and the item's listed price to calculate
// and display the unit price on eBay EU sites.
type UnitInfo struct {
	UnitQuantity float64 `xml:"UnitQuantity"`
	UnitType     string  `xml:"UnitType"`
}

// Variations is only returned for multiple-variation listings, and it is required that the user include
// the IncludeSelector field in the call request, and set its value to Variations.
type Variations struct {
	Pictures              Pictures        `xml:"Pictures"`
	Variations            []Variation     `xml:"Variation"`
	VariationSpecificsSet []NameValueList `xml:"VariationSpecificsSet>NameValueList"`
}

// Pictures contains a set of pictures that correspond to one of the variation specifics, such as 'Color'.
type Pictures struct {
	VariationSpecificName        string              `xml:"VariationSpecificName"`
	VariationSpecificPictureSets []VarSpecificPicSet `xml:"VariationSpecificPictureSet"`
}

// VarSpecificPicSet is returned for each product variation for which there are one or more pictures available,
// helping buyers distinguish between the different variations in the listing.
type VarSpecificPicSet struct {
	PictureURLs            []string `xml:"PictureURL"`
	VariationSpecificValue string   `xml:"VariationSpecificValue"`
}

// Variation Contains data that distinguishes one variation from another. For example, if the items vary by
// color and size, each Variation node specifies a combination of one of those colors and sizes.
// The quantity and price for each variation is also shown in the Variation container
type Variation struct {
	DiscountPriceInfo  DiscountPriceInfo `xml:"DiscountPriceInfo"`
	ProductID          string            `xml:"ProductID"`
	Quantity           int               `xml:"Quantity"`
	SellingStatus      SellingStatus     `xml:"SellingStatus"`
	SKU                string            `xml:"SKU"`
	StartPrice         float64           `xml:"StartPrice"`
	VariationSpecifics []NameValueList   `xml:"VariationSpecifics>NameValueList"`
}

// SellingStatus shows the quantity sold for the variation, including the quantity that is sold through
// 'In-Store Pickup' (a logistics option that is only available to a select number of large US sellers
// with 'brick and mortar' stores). The SellingStatus container is returned for each item variation,
// even if the quantity sold value is '0'.
type SellingStatus struct {
	QuantitySold                int `xml:"QuantitySold"`
	QuantitySoldByPickupInStore int `xml:"QuantitySoldByPickupInStore"`
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
	TaxTable                            []TaxJurisdiction       `xml:"TaxTable>TaxJurisdiction"`
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

// GetSingleItemResponse is a response for GetSingleItemRequest
type GetSingleItemResponse struct {
	responseStandard
	Item ItemExtended `xml:"Item"`
}

// ItemExtended contains details about the listing whose ID was specified in the request.
type ItemExtended struct {
	Item
	ItemCompatibilityCount int             `xml:"ItemCompatibilityCount"`
	ItemCompatibilityList  []Compatibility `xml:"ItemCompatibilityList>Compatibility"`
}

// Compatibility is returned for each motor vehicle that is compatible with the motor vehicle part or accessory.
type Compatibility struct {
	CompatibilityNotes string `xml:"CompatibilityNotes"`
	NameValueLists     []NameValueList
}

/*
===========================================================
*/

// GetUserProfileResponse is a response of GetUserProfileRequest
type GetUserProfileResponse struct {
	responseStandard
	FeedbackDetails []FeedbackDetail `xml:"FeedbackDetails"`
	FeedbackHistory FeedbackHistory  `xml:"FeedbackHistory"`
	User            UserProfile      `xml:"User"`
}

// FeedbackDetail consists of detailed information about one Feedback entry for the specified eBay user.
type FeedbackDetail struct {
	CommentingUser      string  `xml:"CommentingUser"`
	CommentingUserScore int     `xml:"CommentingUserScore"`
	CommentText         string  `xml:"CommentText"`
	CommentTime         string  `xml:"CommentTime"`
	CommentType         string  `xml:"CommentType"`
	FeedbackID          string  `xml:"FeedbackID"`
	FeedbackRatingStar  string  `xml:"FeedbackRatingStar"`
	FeedbackResponse    string  `xml:"FeedbackResponse"`
	FollowUp            string  `xml:"FollowUp"`
	ItemID              string  `xml:"ItemID"`
	ItemPrice           float64 `xml:"ItemPrice"`
	ItemTitle           string  `xml:"ItemTitle"`
	Role                string  `xml:"Role"`
	TransactionID       string  `xml:"TransactionID"`
	CommentReplaced     bool    `xml:"CommentReplaced"`
	Countable           bool    `xml:"Countable"`
	FollowUpReplaced    bool    `xml:"FollowUpReplaced"`
	ResponseReplaced    bool    `xml:"ResponseReplaced"`
}

// FeedbackHistory consists of numerous statistical data about the specified eBay user's Feedback history,
// including counts of Positive, Neutral, and Negative Feedback entries for predefined time periods
// (last week, last month, last 6 months, and last year). For the FeedbackHistory container to be returned,
// the user must include the IncludeSelector field in the request and set its value to FeedbackHistory.
type FeedbackHistory struct {
	AverageRatingDetails                  []AverageRatingDetail `xml:"AverageRatingDetails"`
	BidRetractionFeedbackPeriods          []FeedbackPeriod      `xml:"BidRetractionFeedbackPeriods"`
	NegativeFeedbackPeriods               []FeedbackPeriod      `xml:"NegativeFeedbackPeriods"`
	NeutralCommentCountFromSuspendedUsers int64                 `xml:"NeutralCommentCountFromSuspendedUsers"`
	NeutralFeedbackPeriods                []FeedbackPeriod      `xml:"NeutralFeedbackPeriods"`
	PositiveFeedbackPeriods               []FeedbackPeriod      `xml:"PositiveFeedbackPeriods"`
	TotalFeedbackPeriods                  []FeedbackPeriod      `xml:"TotalFeedbackPeriods"`
	UniqueNegativeFeedbackCount           int64                 `xml:"UniqueNegativeFeedbackCount"`
	UniqueNeutralFeedbackCount            int64                 `xml:"UniqueNeutralFeedbackCount"`
	UniquePositiveFeedbackCount           int64                 `xml:"UniquePositiveFeedbackCount"`
}

// AverageRatingDetail shows the seller's current rating for the Detailed Seller Rating type (specified in the
// RatingDetail field), as well as the total count that this seller has been rated for
// this particular Detailed Seller Rating type.
type AverageRatingDetail struct {
	Rating       float64 `xml:"Rating"`
	RatingCount  int64   `xml:"RatingCount"`
	RatingDetail string  `xml:"RatingDetail"`
}

// FeedbackPeriod shows the cumulative number of all Feedback entries (shown in Count field) for the specified
// time period (shown in PeriodInDays field).
type FeedbackPeriod struct {
	Count        int64 `xml:"Count"`
	PeriodInDays int   `xml:"PeriodInDays"`
}

// UserProfile consists of various details about the eBay user, including Feedback rating, Seller Level,
// link to profile page, and other information. This container is always returned, but more fields will be returned
// under this container if the user includes the IncludeSelector field in the request and sets its value to Details.
type UserProfile struct {
	BasicUser
	AboutMeURL          string `xml:"AboutMeURL"`
	FeedbackDetailsURL  bool   `xml:"FeedbackDetailsURL"`
	MyWorldLargeImage   string `xml:"MyWorldLargeImage"`
	MyWorldSmallImage   string `xml:"MyWorldSmallImage"`
	MyWorldURL          string `xml:"MyWorldURL"`
	NewUser             bool   `xml:"NewUser"`
	RegistrationDate    string `xml:"RegistrationDate"`
	RegistrationSite    string `xml:"RegistrationSite"`
	ReviewsAndGuidesURL string `xml:"ReviewsAndGuidesURL"`
	SellerBusinessType  string `xml:"SellerBusinessType"`
	SellerItemsURL      string `xml:"SellerItemsURL"`
	SellerLevel         string `xml:"SellerLevel"`
	Status              string `xml:"Status"`
	StoreName           string `xml:"StoreName"`
	StoreURL            string `xml:"StoreURL"`
	TopRatedSeller      bool   `xml:"TopRatedSeller"`
}
