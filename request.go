package shopping

// FindProductsRequest represents eBay FindProducts call request
//
// See more: https://developer.ebay.com/Devzone/shopping/docs/CallRef/FindProducts.html#Input
type FindProductsRequest struct {
	RequestBasic
	RequestStandard
	CategoryID string `xml:"CategoryID,omitempty"`
	//DomainNames []string `xml:"DomainName"` 					// is not recommended for use
	//IncludeSelector string `xml:"IncludeSelector,omitempty"` 	// deprecated
	MaxEntries    int       `xml:"MaxEntries,omitempty"`
	PageNumber    int       `xml:"PageNumber,omitempty"`
	ProductID     ProductID `xml:"ProductID,omitempty"`
	ProductSort   string    `xml:"ProductSort,omitempty"`
	QueryKeywords string    `xml:"QueryKeywords,omitempty"`
	SortOrder     string    `xml:"SortOrder,omitempty"`
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
// eBay Product ID (ePID) or a Global Trade Item Number (GTIN), such as a UPC, ISBN, or EAN.
// The product identifier is expressed as a string value, and the type of product identifier
// is expressed in the type attribute.
type ProductID struct {
	ProductIDCodeType string `xml:"type,attr"`
	ProductIDType     string `xml:",cdata"`
}

// WithProductID adds ProductID to request
//
// Use this field to find a catalog product (or products) associated with an
// eBay Product ID (ePID) or a Global Trade Item Number (GTIN), such as a UPC, ISBN, or EAN.
// The product identifier is expressed as a string value, and the type of product identifier
// is expressed in the type attribute.
func (r *FindProductsRequest) WithProductID(productIDCodeType ProductIDCodeTypeOption, productID string) *FindProductsRequest {
	r.ProductID.ProductIDCodeType = string(productIDCodeType)
	r.ProductID.ProductIDType = productID
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
// description, and/or Item Specifics, and it returns a list of matching catalog products.
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

/*
===================================================
*/

// GetCategoryInfoRequest is used to retrieve high-level data for a specified eBay category,
// including category Name, Category ID value, and full category path (category names and category IDs)
//
// See more: https://developer.ebay.com/Devzone/shopping/docs/CallRef/GetCategoryInfo.html#Input
type GetCategoryInfoRequest struct {
	RequestBasic
	RequestStandard
	CategoryID      string `xml:"CategoryID,omitempty"`
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
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
func (r *GetCategoryInfoRequest) WithIncludeSelector(selector IncludeSelectorGCIOption) *GetCategoryInfoRequest {
	r.IncludeSelector = string(selector)
	return r
}

/*
===================================================
*/

type GeteBayTimeRequest struct {
	RequestBasic
	// TODO
}

/*
===================================================
*/

type GetItemStatusRequest struct {
	RequestBasic
	// TODO
}

/*
===================================================
*/

type GetMultipleItemsRequest struct {
	RequestBasic
	// TODO
}

/*
===================================================
*/

type GetShippingCostsRequest struct {
	RequestBasic
	// TODO
}

/*
===================================================
*/

type GetSingleItemRequest struct {
	RequestBasic
	// TODO
}

/*
===================================================
*/

type GetUserProfileRequest struct {
	RequestBasic
	// TODO
}
