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

// CollectOk returns the api response to the `/collectstd` endpoint
func CollectOk() []byte {
	return []byte(`
{
   "ptn":"99999166542651400095315364801168",
   "timestamp":"2022-10-10T18:28:34+00:00",
   "agentBalance":"247000.00",
   "receiptNumber":"999992624813740205",
   "veriCode":"f20873",
   "priceLocalCur":"1000.00",
   "priceSystemCur":"1000.00",
   "localCur":"XAF",
   "systemCur":"XAF",
   "trid": "999992624813740205",
   "pin":null,
   "status":"SUCCESS",
   "payItemDescr":null,
   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"
}
`)
}

// CollectPending returns the api response to the `/collectstd` endpoint for a pending transaction
func CollectPending() []byte {
	return []byte(`
{
   "ptn":"99999166542651400095315364801168",
   "timestamp":"2022-10-10T18:28:34+00:00",
   "agentBalance":"247000.00",
   "receiptNumber":"999992624813740205",
   "veriCode":"f20873",
   "priceLocalCur":"1000.00",
   "priceSystemCur":"1000.00",
   "localCur":"XAF",
   "systemCur":"XAF",
   "trid": "999992624813740205",
   "pin":null,
   "status":"PENDING",
   "payItemDescr":null,
   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"
}
`)
}

// VerifyOk returns the api response to the `/verifytx` endpoint
func VerifyOk() []byte {
	return []byte(`
[
	{
	   "ptn":"99999166542651400095315364801168",
	   "timestamp":"2022-10-10T18:28:34+00:00",
	   "agentBalance":"247000.00",
	   "receiptNumber":"999992624813740205",
	   "veriCode":"f20873",
	   "priceLocalCur":"1000.00",
	   "priceSystemCur":"1000.00",
	   "localCur":"XAF",
	   "systemCur":"XAF",
	   "trid": "999992624813740205",
	   "pin":null,
	   "status":"SUCCESS",
	   "payItemDescr":null,
	   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"
	}
]
`)
}

// VerifyInProcess returns the api response to the `/verifytx` endpoint for an in process transaction
func VerifyInProcess() []byte {
	return []byte(`
[
	{
	   "ptn":"99999166542651400095315364801168",
	   "timestamp":"2022-10-10T18:28:34+00:00",
	   "agentBalance":"247000.00",
	   "receiptNumber":"999992624813740205",
	   "veriCode":"f20873",
	   "priceLocalCur":"1000.00",
	   "priceSystemCur":"1000.00",
	   "localCur":"XAF",
	   "systemCur":"XAF",
	   "trid": "999992624813740205",
	   "pin":null,
	   "status":"INPROCESS",
	   "payItemDescr":null,
	   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"
	}
]
`)
}

// QuoteOk returns the api response to the `/quotestd` endpoint
func QuoteOk() []byte {
	return []byte(`
{
   "quoteId":"15380e55-6227-4a25-8e1f-23c8735ce242",
   "expiresAt":"2022-10-15T13:32:53+00:00",
   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1",
   "amountLocalCur":"100.00",
   "priceLocalCur":100,
   "priceSystemCur":100,
   "localCur":"XAF",
   "systemCur":"XAF",
   "promotion":null
}
`)
}

// VerifyEmpty returns the api response to the `/verifytx` of an empty transaction
func VerifyEmpty() []byte {
	return []byte(`[]`)
}
