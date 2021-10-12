# goidealo

Here you can find our library for idealo. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.

## Generate access token for PWS2.0 (Partner Web Service)

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

## Read an offer

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

## Delete all existing offers

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