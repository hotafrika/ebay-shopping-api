package shopping

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// FindProductsRequest represents eBay FindProducts call request
//
// See more: https://developer.ebay.com/Devzone/shopping/docs/CallRef/FindProducts.html#Input
type FindProductsRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents FindProductsRequest"`
	RequestBasic
	RequestStandard
	CategoryID string `xml:"CategoryID,omitempty"`
	//DomainNames []string `xml:"DomainName"` 					// is not recommended for use
	//IncludeSelector string `xml:"IncludeSelector,omitempty"` 	// deprecated
	MaxEntries    int        `xml:"MaxEntries,omitempty"`
	PageNumber    int        `xml:"PageNumber,omitempty"`
	ProductID     *ProductID `xml:"ProductID,omitempty"`
	ProductSort   string     `xml:"ProductSort,omitempty"`
	QueryKeywords string     `xml:"QueryKeywords,omitempty"`
	SortOrder     string     `xml:"SortOrder,omitempty"`
}

// WithCategoryID adds categoryID to FindProductsRequest
//
// This field is included to restrict the catalog products that are returned.
// Only the catalog products associated with this category ID are returned.
// This field is generally used with the QueryKeywords field.
func (r *FindProductsRequest) WithCategoryID(categoryID string) *FindProductsRequest {
	r.CategoryID = categoryID
	return r
}

// WithPageNumber changes PageNumber in FindProductsRequest
// This field is used to control the page number of results to retrieve in the call.
// If this field is omitted, the first page of results is returned by default.
// You know that you have additional pages or results if the MoreResults field is returned as true.
//
// This field takes a positive integer value equal to or lower than the number of pages available.
// The total number of pages in the results set is shown in the ApproximatePages field of the response.
// Min: 1. Max: total number varies based on search criteria, but can be as high as 10,000+.
// Default: 1.
func (r *FindProductsRequest) WithPageNumber(limit int) *FindProductsRequest {
	if limit < 1 {
		limit = 1
	} else {
		if limit > 10000 {
			limit = 10000
		}
	}

	r.PageNumber = limit
	return r
}

// WithMaxEntries changes MaxEntries in FindProductsRequest
// This field is used to limit/control the maximum number of catalog products that are returned
// per page of data in a single call. This is generally used with string query searches using
// the QueryKeywords field.
//
// If this field is not used, its value defaults to '1', and only one catalog product is returned.
// The user may want to look at the TotalProducts field's value to see how many eBay catalog
// products matched the search criteria, and then the user may want to do another call, possibly
// refining/narrowing the search with a more precise query string in the QueryKeywords field,
// or perhaps with one or more DomainName filters.
//
// If the MoreResults field is returned as true, this indicates that more than one page of results
// are available based on the current search criteria, so the user will have to make additional
// calls to view additional pages of results, changing the PageNumber value as needed.
// Min: 1. Max: total number varies based on search criteria, but can be as high as 800+.
// Default: 1.
func (r *FindProductsRequest) WithMaxEntries(limit int) *FindProductsRequest {
	if limit < 1 {
		limit = 1
	} else {
		if limit > 10000 {
			limit = 10000
		}
	}

	r.MaxEntries = limit
	return r
}

// ProductID ...
//
// Use this field to find a catalog product (or products) associated with an
// eBay Product ID (ePID) or a Global Trade StatusItem Number (GTIN), such as a UPC, ISBN, or EAN.
// The product identifier is expressed as a string value, and the type of product identifier
// is expressed in the type attribute.
type ProductID struct {
	ProductIDCodeType string `xml:"type,attr"`
	ProductIDType     string `xml:",cdata"`
}

// WithProductID adds ProductID to request
//
// Use this field to find a catalog product (or products) associated with an
// eBay Product ID (ePID) or a Global Trade StatusItem Number (GTIN), such as a UPC, ISBN, or EAN.
// The product identifier is expressed as a string value, and the type of product identifier
// is expressed in the type attribute.
func (r *FindProductsRequest) WithProductID(productIDCodeType ProductIDCodeTypeOption, productID string) *FindProductsRequest {
	r.ProductID = &ProductID{
		ProductIDCodeType: string(productIDCodeType),
		ProductIDType:     productID,
	}
	return r
}

// WithProductSort adds ProductSort to request
//
// This field allows the user to control the order in which the retrieved catalog products are
// displayed in the response. If this field is not included, the results are sorted by the catalog
// product's popularity.
//
// See the ProductSortCodeType definition to view the available sort values.
//
// This field can be used in conjunction with the SortOrder field. The SortOrder field controls
// whether catalog products are returned in ascending or descending order (according to the
// ProductSort value). If neither ProductSort nor SortOrder fields are used, catalog products
// are sorted by popularity in descending order.
// Default: Popularity.
func (r *FindProductsRequest) WithProductSort(sortBy ProductSortOption) *FindProductsRequest {
	r.ProductSort = string(sortBy)
	return r
}

// WithQueryKeywords is used to defined a query string using one or more keywords. When you use
// a keyword search, eBay searches the product catalog for matching words in the product title,
// description, and/or StatusItem Specifics, and it returns a list of matching catalog products.
//
// The query string must contain at least three alphanumeric characters.
// The words "and" and "or" are treated like any other word. Only use "and", "or", or "the" if you
// are searching for products containing these words. To use AND or OR logic, use eBay's standard
// search string modifiers. Wildcards (+, -, or *) are also supported. Be careful when using spaces
// before or after modifiers and wildcards.
// Some keyword queries can result in response times of 30 seconds or longer. If too many results
// are returned, you may want to refine the search by passing in more keywords and/or by using one
// or more DomainName filters. Using a CategoryID value is also an option, as this will return only
// catalog products associated with that eBay category.
//
// If you know your product's UPC, EAN, or ISBN, you may want to use the ProductID
// field instead of the QueryKeywords field.
// Max length: 350.
func (r *FindProductsRequest) WithQueryKeywords(query string) *FindProductsRequest {
	r.QueryKeywords = query
	return r
}

// WithSortOrder adds SortOrder to request
//
// This field is used to control whether catalog products are returned in ascending or descending
// order (according to the ProductSort value). If neither ProductSort nor SortOrder fields are used,
// catalog products are sorted by popularity in descending order.
// Default: Descending.
func (r *FindProductsRequest) WithSortOrder(sortBy SortOrderOption) *FindProductsRequest {
	r.SortOrder = string(sortBy)
	return r
}

// GetPage executes FindProductsRequest for page #
// Valid pages # 1 - 10000+
func (r *FindProductsRequest) GetPage(page int) (FindProductsResponse, error) {
	if page < 1 {
		page = 1
	}
	r.WithPageNumber(page)
	body, err := r.getBody()
	if err != nil {
		return FindProductsResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return FindProductsResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return FindProductsResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := FindProductsResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return FindProductsResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// Execute executes FindProductsRequest for the first page
func (r *FindProductsRequest) Execute() (FindProductsResponse, error) {
	return r.GetPage(1)
}

// GetBody return FindProductsRequest body as XML
func (r *FindProductsRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *FindProductsRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetCategoryInfoRequest is used to retrieve high-level data for a specified eBay category,
// including category Name, Category ID value, and full category path (category names and category IDs)
//
// See more: https://developer.ebay.com/Devzone/shopping/docs/CallRef/GetCategoryInfo.html#Input
type GetCategoryInfoRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoRequest"`
	RequestBasic
	RequestStandard
	CategoryID         string              `xml:"CategoryID,omitempty"`
	IncludeSelectorMap map[string]struct{} `xml:"-"`
	IncludeSelector    string              `xml:"IncludeSelector,omitempty"`
}

// WithCategoryID adds categoryID to GetCategoryInfoRequest
//
// In this required field, the user specifies the unique identifier of an eBay category.
// Detailed information is returned for this category.
//
// If a user wanted to see all Level 1 (L1) categories for an eBay site, a value of -1 is
// passed into this field, and the user also includes the IncludeSelector field and
// sets its value to ChildCategories.
func (r *GetCategoryInfoRequest) WithCategoryID(categoryID string) *GetCategoryInfoRequest {
	r.CategoryID = categoryID
	return r
}

// WithIncludeSelector adds IncludeSelector to GetCategoryInfoRequest
//
// This field is included and its value is set to ChildCategories if the user wishes to retrieve all of
// the specified category's children categories (one category level down in eBay categorical hierarchy).
// If the specified category is a leaf category (and has no children), this filter has no effect on the output.
func (r *GetCategoryInfoRequest) WithIncludeSelector(selectors ...IncludeSelectorGCIOption) *GetCategoryInfoRequest {
	if len(r.IncludeSelectorMap) == 0 {
		r.IncludeSelectorMap = make(map[string]struct{})
	}
	for _, s := range selectors {
		r.IncludeSelectorMap[string(s)] = struct{}{}
	}
	var is []string
	for k := range r.IncludeSelectorMap {
		is = append(is, k)
	}
	r.IncludeSelector = strings.Join(is, ",")
	return r
}

// Execute executes GetCategoryInfoRequest
func (r *GetCategoryInfoRequest) Execute() (GetCategoryInfoResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetCategoryInfoResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetCategoryInfoResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetCategoryInfoResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetCategoryInfoResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetCategoryInfoResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetCategoryInfoRequest body as XML
func (r *GetCategoryInfoRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetCategoryInfoRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

type GeteBayTimeRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GeteBayTimeRequest"`
	RequestBasic
}

// Execute executes GeteBayTimeRequest
func (r *GeteBayTimeRequest) Execute() (GeteBayTimeResponse, error) {
	// TODO check content type
	res, err := r.Client.R().Post(r.URL)
	if err != nil {
		return GeteBayTimeResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GeteBayTimeResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GeteBayTimeResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GeteBayTimeResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GeteBayTimeRequest body as XML
func (r *GeteBayTimeRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetItemStatusRequest can be used by sellers and buyers who want to know the current status of any eBay listing.
// All retrieved listings will show listing status, fixed price (or highest bid price for auctions),
// and scheduled end time of listing.
type GetItemStatusRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetItemStatusRequest"`
	RequestBasic
	ItemIDMap map[string]struct{} `xml:"-"`
	ItemIDs   []string            `xml:"ItemID,omitempty"`
}

// WithItemID adds items IDs to request
// The unique identifier of the eBay listing to retrieve. You can retrieve the status of up to 20 listings per call,
// and a separate ItemID field is required for each listing.
func (r *GetItemStatusRequest) WithItemID(itemIDs ...string) *GetItemStatusRequest {
	if len(r.ItemIDMap) == 0 {
		r.ItemIDMap = make(map[string]struct{})
	}
	for _, id := range itemIDs {
		r.ItemIDMap[id] = struct{}{}
	}
	var ids []string
	for k := range r.ItemIDMap {
		ids = append(ids, k)
	}
	r.ItemIDs = ids
	return r
}

// Execute executes GetItemStatusRequest
func (r *GetItemStatusRequest) Execute() (GetItemStatusResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetItemStatusResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetItemStatusResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetItemStatusResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetItemStatusResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetItemStatusResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetItemStatusRequest body as XML
func (r *GetItemStatusRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetItemStatusRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetMultipleItemsRequest retrieve much of the information that is visible on a listing's
// View StatusItem page on the eBay site, such as title and prices.
type GetMultipleItemsRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetMultipleItemsRequest"`
	RequestBasic
	IncludeSelectorMap map[string]struct{} `xml:"-"`
	IncludeSelector    string              `xml:"IncludeSelector,omitempty"`
	ItemIDs            []string            `xml:"ItemID,omitempty"`
}

// WithIncludeSelector adds selector options to request
func (r *GetMultipleItemsRequest) WithIncludeSelector(selectors ...IncludeSelectorGMIOption) *GetMultipleItemsRequest {
	if len(r.IncludeSelectorMap) == 0 {
		r.IncludeSelectorMap = make(map[string]struct{})
	}
	for _, s := range selectors {
		r.IncludeSelectorMap[string(s)] = struct{}{}
	}
	var is []string
	for k := range r.IncludeSelectorMap {
		is = append(is, k)
	}
	r.IncludeSelector = strings.Join(is, ",")
	return r
}

// WithItemID adds items IDs to request
// The uniqe ID that identifies the listing for which to retrieve the data.
// You can provide a maximum of 20 ItemID values.
// Max length: 19 (Note: The eBay database specifies 38. Currently, StatusItem IDs are usually 9 to 12 digits).
func (r *GetMultipleItemsRequest) WithItemID(itemIDs ...string) *GetMultipleItemsRequest {
	r.ItemIDs = append(r.ItemIDs, itemIDs...)
	return r
}

// WithItemIDs replaces items IDs in request
// The uniqe ID that identifies the listing for which to retrieve the data.
// You can provide a maximum of 20 ItemID values.
// Max length: 19 (Note: The eBay database specifies 38. Currently, StatusItem IDs are usually 9 to 12 digits).
func (r *GetMultipleItemsRequest) WithItemIDs(itemIDs ...string) *GetMultipleItemsRequest {
	r.ItemIDs = itemIDs
	return r
}

// Execute executes GetMultipleItemsRequest
func (r *GetMultipleItemsRequest) Execute() (GetMultipleItemsResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetMultipleItemsResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetMultipleItemsResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetMultipleItemsResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetMultipleItemsResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetMultipleItemsResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetMultipleItemsRequest body as XML
func (r *GetMultipleItemsRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetMultipleItemsRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetShippingCostsRequest is used to retrieve the estimated shipping cost to ship an active item
// to a specified destination country and postal code. Any user can make this call on an active
// listing (does not have to be the seller or buyer/bidder).
// It pertains to all shipping types, including flat-rate and calculated.
type GetShippingCostsRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetShippingCostsRequest"`
	RequestBasic
	// DestinationCountryCode from https://developer.ebay.com/Devzone/shopping/docs/CallRef/types/CountryCodeType.html
	DestinationCountryCode string `xml:"DestinationCountryCode,omitempty"`
	DestinationPostalCode  string `xml:"DestinationPostalCode,omitempty"`
	IncludeDetails         bool   `xml:"IncludeDetails,omitempty"`
	ItemID                 string `xml:"ItemID"`
	QuantitySold           int    `xml:"QuantitySold,omitempty"`
}

// WithDestinationCountryCode adds code from
// Destination country code. If DestinationCountryCode is US, postal code is required and represents US zip code.
// https://developer.ebay.com/Devzone/shopping/docs/CallRef/types/CountryCodeType.html
func (r *GetShippingCostsRequest) WithDestinationCountryCode(code string) *GetShippingCostsRequest {
	r.DestinationCountryCode = code
	return r
}

// WithDestinationPostalCode adds postal code
// Destination country postal code (or zip code, for US). Ignored if no country code is provided.
func (r *GetShippingCostsRequest) WithDestinationPostalCode(code string) *GetShippingCostsRequest {
	r.DestinationPostalCode = code
	return r
}

// WithIncludeDetails allows the user to return  the ShippingDetails in the response.
func (r *GetShippingCostsRequest) WithIncludeDetails(includeDetails bool) *GetShippingCostsRequest {
	r.IncludeDetails = includeDetails
	return r
}

// WithItemID adds itemID to request
// The item ID that uniquely identifies the listing for which to retrieve the data.
// Max length: 19 (Note: The eBay database specifies 38. Currently, StatusItem IDs are usually 9 to 12 digits).
func (r *GetShippingCostsRequest) WithItemID(itemID string) *GetShippingCostsRequest {
	r.ItemID = itemID
	return r
}

// WithQuantitySold adds QuantitySold to request
// Quantity of items sold to a single buyer and to be shipped together.
func (r *GetShippingCostsRequest) WithQuantitySold(sold int) *GetShippingCostsRequest {
	r.QuantitySold = sold
	return r
}

// Execute executes GetShippingCostsRequest
func (r *GetShippingCostsRequest) Execute() (GetShippingCostsResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetShippingCostsResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetShippingCostsResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetShippingCostsResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetShippingCostsResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetShippingCostsResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetShippingCostsRequest body as XML
func (r *GetShippingCostsRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetShippingCostsRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetSingleItemRequest retrieves publicly visible details about one listing on eBay.
// This gives you most of the data that eBay shows to the general public on the
// View StatusItem page (title, description, basic price information, and other details).
type GetSingleItemRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetSingleItemRequest"`
	RequestBasic
	IncludeSelector    string              `xml:"IncludeSelector,omitempty"`
	IncludeSelectorMap map[string]struct{} `xml:"-"`
	ItemID             string              `xml:"ItemID"`
	VariationSKU       string              `xml:"VariationSKU,omitempty"`
	VariationSpecifics *VariationSpecifics `xml:"VariationSpecifics,omitempty"`
}

type VariationSpecifics struct {
	NameValueList []NameValueListType `xml:"NameValueList,omitempty"`
}

// NameValueListType  is an array of StatusItem Specifics name-value pairs for an eBay Catalog product
// (if FindProducts is used) or StatusItem Specifics name-value pairs for a single-variation listing or individual
// variation within a multiple-variation listing (if GetSingleItem or GetMultipleItems is used).
type NameValueListType struct {
	Name   string   `xml:"Name,omitempty"`
	Values []string `xml:"Value,omitempty"`
}

// WithIncludeSelector adds selector options to request
func (r *GetSingleItemRequest) WithIncludeSelector(selectors ...IncludeSelectorGSIOption) *GetSingleItemRequest {
	if len(r.IncludeSelectorMap) == 0 {
		r.IncludeSelectorMap = make(map[string]struct{})
	}
	for _, s := range selectors {
		r.IncludeSelectorMap[string(s)] = struct{}{}
	}
	var is []string
	for k := range r.IncludeSelectorMap {
		is = append(is, k)
	}
	r.IncludeSelector = strings.Join(is, ",")
	return r
}

// WithItemID adds itemID to request
// The item ID that uniquely identifies the listing for which to retrieve the data.
// Max length: 19 (Note: The eBay database specifies 38. Currently, StatusItem IDs are usually 9 to 12 digits).
func (r *GetSingleItemRequest) WithItemID(itemID string) *GetSingleItemRequest {
	r.ItemID = itemID
	return r
}

// WithVariationSKU adds VariationSKU to request
// Variation-level SKU that uniquely identifies a variation within the listing identified by ItemID.
// Only applicable when the seller included variation-level SKU (Variation.SKU) values.
// Retrieves all the usual listing fields, but limits the variations content to the specified variation.
// If not specified, the response includes all variations.
func (r *GetSingleItemRequest) WithVariationSKU(variationSKU string) *GetSingleItemRequest {
	r.VariationSKU = variationSKU
	return r
}

// WithVariationSpecifics adds additional StatusItem Specific name-value pairs
func (r *GetSingleItemRequest) WithVariationSpecifics(name string, values ...string) *GetSingleItemRequest {
	if r.VariationSpecifics == nil {
		r.VariationSpecifics = &VariationSpecifics{}
	}
	r.VariationSpecifics.NameValueList = append(r.VariationSpecifics.NameValueList, NameValueListType{
		Name:   name,
		Values: values,
	})
	return r
}

// Execute executes GetSingleItemRequest
func (r *GetSingleItemRequest) Execute() (GetSingleItemResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetSingleItemResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetSingleItemResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetSingleItemResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetSingleItemResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetSingleItemResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetSingleItemRequest body as XML
func (r *GetSingleItemRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetSingleItemRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

/*
===================================================
*/

// GetUserProfileRequest ...
type GetUserProfileRequest struct {
	XMLName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetUserProfileRequest"`
	RequestBasic
	IncludeSelector    string              `xml:"IncludeSelector,omitempty"`
	IncludeSelectorMap map[string]struct{} `xml:"-"`
	UserID             string              `xml:"UserID"`
}

// WithIncludeSelector adds selector options to request
func (r *GetUserProfileRequest) WithIncludeSelector(selectors ...IncludeSelectorGUPOption) *GetUserProfileRequest {
	if len(r.IncludeSelectorMap) == 0 {
		r.IncludeSelectorMap = make(map[string]struct{})
	}
	for _, s := range selectors {
		r.IncludeSelectorMap[string(s)] = struct{}{}
	}
	var is []string
	for k := range r.IncludeSelectorMap {
		is = append(is, k)
	}
	r.IncludeSelector = strings.Join(is, ",")
	return r
}

// WithUserID adds userID to request
// An eBay user ID is input into this field to retrieve information about that eBay user.
func (r *GetUserProfileRequest) WithUserID(userID string) *GetUserProfileRequest {
	r.UserID = userID
	return r
}

// Execute executes GetUserProfileRequest
func (r *GetUserProfileRequest) Execute() (GetUserProfileResponse, error) {
	body, err := r.getBody()
	if err != nil {
		return GetUserProfileResponse{}, fmt.Errorf("unable to serialize req body: %w", err)
	}
	// TODO check content type
	res, err := r.Client.R().SetBody(body).Post(r.URL)
	if err != nil {
		return GetUserProfileResponse{}, fmt.Errorf("sending req: %w", err)
	}
	if res.StatusCode() != 200 {
		return GetUserProfileResponse{}, fmt.Errorf("status code %d: %s", res.StatusCode(), res.String())
	}
	ar := GetUserProfileResponse{}
	err = xml.Unmarshal(res.Body(), &ar)
	fmt.Println(string(res.Body()))
	if err != nil {
		return GetUserProfileResponse{}, fmt.Errorf("parsing response body: %w", err)
	}
	return ar, nil
}

// GetBody return GetUserProfileRequest body as XML
func (r *GetUserProfileRequest) GetBody() ([]byte, error) {
	b, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}

func (r *GetUserProfileRequest) getBody() ([]byte, error) {
	b, err := xml.Marshal(r)
	if err != nil {
		return b, err
	}
	return append([]byte(xml.Header), b...), err
}
