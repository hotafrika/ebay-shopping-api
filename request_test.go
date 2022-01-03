package shopping

import (
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestFindProductsRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *FindProductsRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewFindProductsRequest()
	request1.WithQueryKeywords("Harry Potter")
	request1.WithMaxEntries(2)
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 FindProductsRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "findproducts", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 FindProductsRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1, req2)
		})

	}
}

func TestGetCategoryInfoRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetCategoryInfoRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetCategoryInfoRequest()
	request1.WithCategoryID("279")
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewGetCategoryInfoRequest()
	request2.WithCategoryID("-1")
	request2.WithIncludeSelector(IncludeSelectorChildCategories)
	tests = append(tests, tcase{
		filename: "HighLevelCategories.xml",
		request:  request2,
	})

	request3 := service.NewGetCategoryInfoRequest()
	request3.WithCategoryID("-1")
	tests = append(tests, tcase{
		filename: "CategoryVersionUpdateTime.xml",
		request:  request3,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetCategoryInfoRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "categoryinfo", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetCategoryInfoRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1, req2)
		})

	}
}

func TestGeteBayTimeRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GeteBayTimeRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGeteBayTimeRequest()
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GeteBayTimeRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "ebaytime", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GeteBayTimeRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1, req2)
		})

	}
}

func TestGetItemStatusRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetItemStatusRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetItemStatusRequest()
	request1.WithItemID("1**********1", "2**********9", "3**********4", "3**********5", "3**********7")
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetItemStatusRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "itemstatus", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetItemStatusRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1.RequestBasic, req2.RequestBasic)
			assert.EqualValues(t, req1.ItemIDs, req2.ItemIDs)
		})

	}
}

func TestGetMultipleItemsRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetMultipleItemsRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetMultipleItemsRequest()
	request1.WithItemID("1**********7", "2**********0", "9********3")
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetMultipleItemsRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "multipleitems", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetMultipleItemsRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1.RequestBasic, req2.RequestBasic)
			assert.EqualValues(t, req1.ItemIDs, req2.ItemIDs)
			assert.Equal(t, req1.IncludeSelector, req2.IncludeSelector)
		})

	}
}

func TestGetShippingCostsRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetShippingCostsRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetShippingCostsRequest()
	request1.WithItemID("1**********1")
	request1.WithDestinationCountryCode("US")
	request1.WithDestinationPostalCode("9***8")
	request1.WithIncludeDetails(true)
	request1.WithQuantitySold(1)
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetShippingCostsRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "shippingcosts", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetShippingCostsRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1, req2)
		})

	}
}

func TestGetSingleItemRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetSingleItemRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetSingleItemRequest()
	request1.WithItemID("1**********1")
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewGetSingleItemRequest()
	request2.WithItemID("1**********2")
	request2.WithIncludeSelector(IncludeSelectorSIDescription)
	tests = append(tests, tcase{
		filename: "ProductDetailsItemCondition.xml",
		request:  request2,
	})

	request3 := service.NewGetSingleItemRequest()
	request3.WithItemID("1**********1")
	request3.WithIncludeSelector(IncludeSelectorSITextShippingCosts)
	tests = append(tests, tcase{
		filename: "IncludingShippingCostSummary.xml",
		request:  request3,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetSingleItemRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "singleitem", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetSingleItemRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1.RequestBasic, req2.RequestBasic)
			assert.Equal(t, req1.ItemID, req2.ItemID)
			assert.Equal(t, req1.IncludeSelector, req2.IncludeSelector)
			assert.Equal(t, req1.VariationSKU, req2.VariationSKU)
			assert.EqualValues(t, req1.VariationSpecifics, req2.VariationSpecifics)
		})

	}
}

func TestGetUserProfileRequest_GetBody(t *testing.T) {
	type tcase struct {
		filename string
		request  *GetUserProfileRequest
	}

	var tests []tcase

	service := NewService("")

	request1 := service.NewGetUserProfileRequest()
	request1.WithUserID("h***************r")
	tests = append(tests, tcase{
		filename: "Basic.xml",
		request:  request1,
	})

	request2 := service.NewGetUserProfileRequest()
	request2.WithUserID("h***************r")
	request2.WithIncludeSelector(IncludeSelectorUPDetails)
	tests = append(tests, tcase{
		filename: "ReturningDetailedUserInformation.xml",
		request:  request2,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			var req1 GetUserProfileRequest
			err = xml.Unmarshal(got, &req1)
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "xml", "userprofile", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			var req2 GetUserProfileRequest
			err = xml.Unmarshal(buf.Bytes(), &req2)
			if !assert.NoError(t, err) {
				return
			}

			//fmt.Println(string(got))
			assert.Equal(t, req1.RequestBasic, req2.RequestBasic)
			assert.Equal(t, req1.UserID, req2.UserID)
			assert.Equal(t, req1.IncludeSelector, req2.IncludeSelector)
		})

	}
}
