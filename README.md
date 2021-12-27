![example workflow](https://github.com/hotafrika/ebay-shopping-api/actions/workflows/autotests.yml/badge.svg)

# ebay-shopping-api
Golang client for eBay Shopping API

For using eBay shopping API you need to get Application Token. You can make it using [this repo.](github.com/hotafrika/ebay-common)

### Example
```go
package main

import (
	"fmt"
	"github.com/hotafrika/ebay-common/auth"
	"github.com/hotafrika/ebay-shopping-api"
)

func main() {
	authService := auth.NewService().WithScopes(auth.ScopeCredentialCommon).
		WithCredentials("myAppID", "myAppSecret")

	token, err := authService.GetAppToken()
	if err != nil {
		panic(err)
	}
	shoppingService := shopping.NewService(token.Token)
	r := shoppingService.NewGetSingleItemRequest()
	r.WithItemID("123")
	r.WithVariationSpecifics("param1", "one", "two")
	r.WithVariationSpecifics("param2", "one2", "two2")
	
	res, err := r.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

```
