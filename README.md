# goidealo

Here you can find our library for idealo. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.

## Generate access token for PWS2.0 (Partner Web Service)

To generate an access token for the pws api, you can use the following function. **Important! The bearer token must be renewed after 3600 seconds.**

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