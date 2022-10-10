package stubs

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
   "trid":null,
   "pin":null,
   "status":"SUCCESS",
   "payItemDescr":null,
   "payItemId":"S-112-951-CMORANGE-20062-CM_ORANGE_VTU_CUSTOM-1"
}
`)
}
