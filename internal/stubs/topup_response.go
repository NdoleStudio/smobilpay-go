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
