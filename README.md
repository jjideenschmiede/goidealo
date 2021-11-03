# goidealo

[![Go](https://github.com/jjideenschmiede/goidealo/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goidealo/actions/workflows/go.yml) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goidealo.svg)](https://github.com/jjideenschmiede/goidealo) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goidealo)](https://goreportcard.com/report/github.com/jjideenschmiede/goidealo) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/goidealo?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goidealo)

Here you can find our library for idealo. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.

## Install 

```console
go get github.com/jjideenschmiede/goidealo
```

## How to use?

### Technology Partner Header

To allow technology partners to transfer their headers as well, we have added another parameter to the request. Now you can set new values flexibly.

[Here](https://connect.idealo.de/pam/DE_idealo-PWS-API_Header.pdf) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    ClientId:       "",
    ClientPassword: "",
    Header:         map[string]string{},
}

// Add header
r.Header["Partner-Merchant-Solution"] = "J&J Afterialo"
r.Header["Partner-Merchant-Solution-Version"] = "1.0.0"
r.Header["Integration-Partner"] = "J&J Ideenschmiede GmbH"
r.Header["Interface-Version"] = "unbekannt"
```


### Generate access token for PWS2.0 (Partner Web Service)

To generate an access token for the pws api, you can use the following function. **Important! The bearer token must be renewed after 3600 seconds.**

[Here](https://import.idealo.com/docs/#authorization-token) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    ClientId:       "",
    ClientPassword: "",
}

// Get access token
accessToken, err := goidealo.PwsAccessToken(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(accessToken)
}
```

### Read an offer

With this function you can read an offer based on the sku. The following attributes are needed for this:

Also the shop id, the Sku and the and the access token is needed in the request struct. This must be generated before via the function: **goidealo.PwsAccessToken(r)**.

[Here](https://import.idealo.com/docs/#_get) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
}

// Get offer by sku
offer, err := goidealo.Offer(325081, "21-Lloyd-27-600-12-Hagen-UK6.5-Gr.40", r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(offer)
}
```

### Writes an offer

If you want to create a new offer, then this goes as follows. Some attributes are needed for this. You can see them in the following example. [Here](https://import.idealo.com/docs/#_put) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
}

// Set body
body := goidealo.OfferBody{
    Sku:   "21-Lloyd-27-600-12-Hagen-UK6.5-Gr.40",
    Title: "Lloyd Men Hagen 27-600-12 Herren Schuhe Derby Schnürer Wildleder Dunkelbraun",
    Price: "89.95",
    Url:   "http://www.shoes4friends.de",
    PaymentCosts: &goidealo.OfferBodyPaymentCosts{
        PAYPAL: "2.63",
    },
    DeliveryCosts: &goidealo.OfferBodyDeliveryCosts{
        DHL: "0.69",
    },
    BasePrice:                "",
    PackagingUnit:            5,
    VoucherCode:              "",
    BranchId:                 "lloyd",
    Brand:                    "Lloyd",
    Oens:                     nil,
    CategoryPath:             nil,
    Description:              "Lloyd Men Hagen 27-600-12 Herren Schuhe Derby Schnürer Wildleder Dunkelbraun",
    ImageUrls:                nil,
    Eans:                     nil,
    Hans:                     nil,
    Pzns:                     nil,
    Kbas:                     nil,
    MerchantName:             "",
    MerchantId:               "",
    DeliveryComment:          "",
    Delivery:                 "1-3 Werktage",
    MaxOrderProcessingTime:   1,
    FreeReturnDays:           30,
    Checkout:                 true,
    CheckoutLimitPerPeriod:   13,
    QuantityPerOrder:         2,
    MinimumPrice:             "",
    FulfillmentType:          "",
    TwoManHandlingFee:        "",
    DisposalFee:              "",
    Eec:                      "",
    EnergyLabels:             nil,
    Deposit:                  "",
    Size:                     "",
    Colour:                   "",
    Gender:                   "",
    Material:                 "",
    Replica:                  false,
    Used:                     false,
    Download:                 false,
    DynamicProductAttributes: nil,
}

// Add images
body.ImageUrls = append(body.ImageUrls, "https://www.marken-schuhmarkt.de/21-Lloyd-Herren-Schuhe-Derby-Halbschuhe-Wildleder-Hagen-27-600-12-braun-001.jpg")

// Add eans
body.Eans = append(body.Eans, "4032055134146")

// Update the timestamp
offer, err := goidealo.CreateOffer(325081, body, r)
if err != nil {
    fmt.Println(err)
} else {
	fmt.Println(offer)
}
```

### Partially updates an offer

If you want to revise an offer, you can do it as follows. [Here](https://import.idealo.com/docs/#_patch) you can find the description in the idealo documentation.

```go
// Set body
body := goidealo.OfferBody{
    Sku:   "21-Lloyd-27-600-12-Hagen-UK6.5-Gr.40",
    Title: "Lloyd Men Hagen 27-600-12 Herren Schuhe Derby Schnürer Wildleder Dunkelbraun",
    Price: "189.95",
}

// Update an offer
offer, err := goidealo.UpdateOffer(325081, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(offer)
}
```


### Delete an offer 

If you want to remove a special offer, you can do this with the following function. For this some attributes like shop id, sku and the access token are needed. [Here](https://import.idealo.com/docs/#_delete) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
AccessToken: "",
}

// Delete an offer
err := goidealo.DeleteOffer(325081, "21-Lloyd-27-600-12-Hagen-UK6.5-Gr.40", r)
if err != nil {
    fmt.Println(err)
}
```

### Delete all existing offers

If you want to remove all existing offers, then you can do this with the following function, for this the shop id and the access token are needed. [Here](https://import.idealo.com/docs/#_delete_all) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
}

// Delete all existing offers
err := goidealo.DeleteAllOffers(325081, r)
if err != nil {
    fmt.Println(err)
}
```

### Updating the import timestamp for all offers

If you want to update the timestamp of all products, you can do it as follows. You only need the store id and the access token. [Here](https://import.idealo.com/docs/#_update_timestamp) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
}

// Update the timestamp
err := goidealo.UpdateOffersTimestamp(325081, r)
if err != nil {
    fmt.Println(err)
}
```

### Generate access token for MOA2.0

To generate an access token for the moa api, you can use the following function. **Important! The bearer token must be renewed after 3600 seconds.** If the sandbox system is to be used, then this token is not required.

[Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-authorization-controller-it-client-login) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    ClientId:       "",
    ClientPassword: "",
    Sandbox:        false,
}

// Get access token
accessToken, err := goidealo.MoaAccessToken(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(accessToken)
}
```

### Get orders

To be able to read all orders, you can use the following function and populate it with url parameters. We have already set the url parameter pageSize. With a value of 250 orders. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-order-controller-it-get-orders) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Add url parameter
parameter := make(map[string]string)

parameter["status"] = "PROCESSING"
parameter["from"] = "2019-05-01T00:00:00Z"
parameter["pageNumber"] = "0"

// Read all orders
orders, err := goidealo.Orders(325081, parameter, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(orders)
}
```

### Get order

If you want to read out a specific order, this is done as follows. **For this you need the id of the shop and the id of the order.** [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-order-controller-it-custom-get-order) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Read an order
orders, err := goidealo.Order(325081, "00ALP766AM", r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(orders)
}
```

### Set Merchant Order Number

If you want to set the merchant order number, you can do this using the following function. For this you need the shop id and the idealo order id. In addition, the body in the given struct. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-order-controller-it-set-merchant-order-number) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Define body
body := goidealo.MerchantOrderNumberBody{
    MerchantOrderNumber: "012345777777",
}

// Set a merchant order number
merchantOrderNumber, err := goidealo.MerchantOrderNumber(325081, "00REAVQH3Y", body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(merchantOrderNumber)
}
```

### Set Fulfillment Information

If you want to assign the fulfillment to an order, you can do this with the following function. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-fulfillment-controller-it-set-fulfillment-information) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Define body
body := goidealo.FulfillmentInformationBody{
    Carrier:      "DHL",
    TrackingCode: []string{"example-tracking-number-1234"},
}
// Set a fulfillment information
fulfillmentInformation, err := goidealo.FulfillmentInformation(325081, "00REAVQH3Y", body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(fulfillmentInformation)
}
```

### Revoke order

If an order is to be revoked, then you can do it through the following api endpoint. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-revocation-controller-it-custom-revoke-order) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Define body
body := goidealo.RevokeOrderBody{
    Sku:               "21-Lloyd-27-600-12-Hagen-UK6.5-Gr.40",
    RemainingQuantity: 0,
    Reason:            "MERCHANT_DECLINE",
    Comment:           "Sorry, buddy.",
}

// Revoke an order
revokeOrder, err := goidealo.RevokeOrder(325081, "00ALP766AM", body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(revokeOrder)
}
```

### Refund order

If you want to make a refund, then this goes through the following route. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-refund-controller-it-refund-order) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Define body
body := goidealo.RefundOrderBody{
    RefundAmount: 1.99,
    Currency:     "EUR",
}

// Set a refund for an order
refundOrder, err := goidealo.RefundOrder(325081, "00ALP766AM", body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(refundOrder)
}
```

### Get refunds 

If you want to read out the refunds, this is done as follows. [Here](https://cdn.idealo.com/folder/Direktkauf/documentation/merchant-order-api-v2.html#resources-refund-controller-it-get-refunds) you can find the description in the idealo documentation.

```go
// Define request
r := goidealo.Request{
    AccessToken: "",
    Sandbox:     false,
}

// Get refunds
refundOrder, err := goidealo.Refunds(325081, "00ALP766AM", r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(refundOrder)
}
```