package stubs

// TopupGetPackagesOk returns the api responses to /topup
func TopupGetPackagesOk() []byte {
	return []byte(`
[
   {
      "serviceid":"2000",
      "merchant":"CMENEOPREPAID",
      "payItemId":"S-112-974-CMENEOPREPAID-2000-10010-1",
      "amountType":"CUSTOM",
      "localCur":"XAF",
      "name":"Custom Amount",
      "amountLocalCur":null,
      "description":"",
      "payItemDescr":null,
      "optStrg":null,
      "optNmb":null
   },
   {
      "serviceid":"20062",
      "merchant":"CMORANGE",
      "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1",
      "amountType":"CUSTOM",
      "localCur":"XAF",
      "name":"Airtime Custom Amount",
      "amountLocalCur":null,
      "description":"Airtime Custom Topup",
      "payItemDescr":null,
      "optStrg":null,
      "optNmb":null
   }
]
`)
}

// TopupCollectOk returns the api response to the `/collectstd` of a topup transaction
func TopupCollectOk() []byte {
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

// TopupVerifyOk returns the api response to the `/verifytx` of a topup transaction
func TopupVerifyOk() []byte {
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

// TopupQuoteOk returns the api response to the `quotestd` endpoint
func TopupQuoteOk() []byte {
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

// TopupVerifyEmpty returns the api response to the `/verifytx` of a topup transaction
func TopupVerifyEmpty() []byte {
	return []byte(`[]`)
}
