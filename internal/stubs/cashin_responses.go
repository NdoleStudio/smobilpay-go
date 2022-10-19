package stubs

// CashinGetOk is the response when getting the `/cashin` endpoint
func CashinGetOk() []byte {
	return []byte(`
[
    {
        "serviceid": "30052",
        "merchant": "CMORANGEOM",
        "payItemId": "S-112-948-CMORANGEOM-30052-2006125104-1",
        "amountType": "CUSTOM",
        "localCur": "XAF",
        "name": "Custom Amount",
        "amountLocalCur": null,
        "description": "Customer amount",
        "payItemDescr": null,
        "optStrg": null,
        "optNmb": null
    }
]
`)
}

// CashinGetError is the cashout error with an invalid service
func CashinGetError() []byte {
	return []byte(`
{
    "devMsg": "Service unknown",
    "usrMsg": "Service unknown",
    "respCode": 40602,
    "link": "http://support.maviance.com/"
}
`)
}
