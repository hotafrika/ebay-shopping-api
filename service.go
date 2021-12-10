package shopping

import (
	"github.com/go-resty/resty/v2"
	"time"
)

// Service represents Ebay Shopping API service
type Service struct {
	version   string
	endpoint  string
	siteID    string
	xIAFToken string
	timeout   time.Duration
}

// NewService creates new Ebay Shopping service
// API version according to EbayShoppingAPIVersion (1199)
// Default endpoint: EbayEndpointProduction (https://open.api.ebay.com/shopping)
// Default GlobalID: SiteIDEbayUS (0)
// Default Page Limit: DefaultItemsPerPage (100)
// Default timeout for requests: 10 seconds
func NewService(xIAFToken string) *Service {
	s := &Service{
		version:   EbayShoppingAPIVersion,
		xIAFToken: xIAFToken,
		timeout:   10 * time.Second,
	}
	s.WithEndpoint(EbayEndpointProduction)
	s.WithSiteID(SiteIDEbayUS)
	return s
}

// WithEndpoint changes endpoint for service
// You can add your own endpoint (for tests purposes)
func (s *Service) WithEndpoint(endpoint string) *Service {
	s.endpoint = endpoint
	return s
}

// WithSiteID changes site for search
func (s *Service) WithSiteID(siteID SiteID) *Service {
	s.siteID = string(siteID)
	return s
}

// WithTimeout changes default timeout for search requests
func (s *Service) WithTimeout(timeout time.Duration) *Service {
	s.timeout = timeout
	return s
}

// WithToken changes IAFToken for service
func (s *Service) WithToken(xIAFToken string) *Service {
	s.xIAFToken = xIAFToken
	return s
}

// creates new http client (resty)
func (s *Service) newHTTPClient() *resty.Client {
	return resty.New().
		SetHeader("X-EBAY-API-VERSION", s.version).
		SetHeader("X-EBAY-API-IAF-TOKEN", s.xIAFToken).
		SetHeader("X-EBAY-API-REQUEST-ENCODING", EbayRequestDataFormat).
		SetHeader("X-EBAY-API-RESPONSE-ENCODING", EbayResponseDataFormat).
		SetHeader("X-EBAY-API-SITE-ID", s.siteID).
		SetTimeout(s.timeout)
}

// NewFindProductsRequest creates new FindProductsRequest
func (s *Service) NewFindProductsRequest() *FindProductsRequest {
	req := FindProductsRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationFindProducts))
	req.URL = s.endpoint
	return &req
}

// NewGetCategoryInfoRequest creates new GetCategoryInfoRequest
func (s *Service) NewGetCategoryInfoRequest() *GetCategoryInfoRequest {
	req := GetCategoryInfoRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetCategoryInfo))
	req.URL = s.endpoint
	return &req
}

// NewGetCategoryInfoRequestWithCategory creates new GetCategoryInfoRequest
func (s *Service) NewGetCategoryInfoRequestWithCategory(categoryID string) *GetCategoryInfoRequest {
	req := GetCategoryInfoRequest{
		CategoryID: categoryID,
	}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetCategoryInfo))
	req.URL = s.endpoint
	return &req
}

// NewGeteBayTimeRequest creates new GeteBayTimeRequest
func (s *Service) NewGeteBayTimeRequest() *GeteBayTimeRequest {
	req := GeteBayTimeRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGeteBayTime))
	req.URL = s.endpoint
	return &req
}

// NewGetItemStatusRequest creates new GetItemStatusRequest
func (s *Service) NewGetItemStatusRequest() *GetItemStatusRequest {
	req := GetItemStatusRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetItemStatus))
	req.URL = s.endpoint
	return &req
}

// NewGetMultipleItemsRequest creates new GetMultipleItemsRequest
func (s *Service) NewGetMultipleItemsRequest() *GetMultipleItemsRequest {
	req := GetMultipleItemsRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetMultipleItems))
	req.URL = s.endpoint
	return &req
}

// NewGetShippingCostsRequest creates new GetShippingCostsRequest
func (s *Service) NewGetShippingCostsRequest() *GetShippingCostsRequest {
	req := GetShippingCostsRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetShippingCosts))
	req.URL = s.endpoint
	return &req
}

// NewGetSingleItemRequest creates new GetSingleItemRequest
func (s *Service) NewGetSingleItemRequest() *GetSingleItemRequest {
	req := GetSingleItemRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetSingleItem))
	req.URL = s.endpoint
	return &req
}

// NewGetUserProfileRequest creates new GetUserProfileRequest
func (s *Service) NewGetUserProfileRequest() *GetUserProfileRequest {
	req := GetUserProfileRequest{}
	req.Client = s.newHTTPClient().
		SetHeader("X-EBAY-API-CALL-NAME", string(OperationGetUserProfile))
	req.URL = s.endpoint
	return &req
}
