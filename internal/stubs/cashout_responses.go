package stubs

// CashoutGetOk is the response when getting the `/cashout` endpoint
func CashoutGetOk() []byte {
	return []byte(`
[
    {
        "serviceid": "20053",
        "merchant": "MTNMOMO",
        "payItemId": "S-112-949-MTNMOMO-20053-200050001-1",
        "amountType": "CUSTOM",
        "localCur": "XAF",
        "name": "CASH-OUT",
        "amountLocalCur": null,
        "description": "Cash out a custom amount",
        "payItemDescr": null,
        "optStrg": null,
        "optNmb": null
    }
]
`)
}

// CashoutGetError is the cashout error with an invalid service
func CashoutGetError() []byte {
	return []byte(`
{
    "devMsg": "Service unknown",
    "usrMsg": "Service unknown",
    "respCode": 40602,
    "link": "http://support.maviance.com/"
}
`)
}
