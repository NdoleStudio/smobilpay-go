package stubs

// PingOk returns the api response to the `/ping` endpoint
func PingOk() []byte {
	return []byte(`
{
    "time": "2022-10-08T14:31:10+00:00",
    "version": "2.2.0",
    "nonce": "95cdf110-4614-4d95-b6c2-f14fe01c4995",
    "key": "6B352110-4716-11ED-963F-0800200C9A66"
}
`)
}

// PingError returns an error api response to the `/ping` endpoint
func PingError() []byte {
	return []byte(`
{
    "devMsg": "Access token invalid",
    "usrMsg": "Authentication error when connecting to service",
    "respCode": 4009,
    "link": "http:\\/\\/support.maviance.com\\/"
}
`)
}
