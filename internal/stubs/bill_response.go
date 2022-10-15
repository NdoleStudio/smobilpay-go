package stubs

// BillGet is the response for a bill payment
func BillGet() []byte {
	return []byte(`
[
    {
        "payItemId": "S-112-950-ENEO-10039-8d12d9869ae74225b9fe69da5c1f03a4-4a9309e5d7224da3a788c6a63c9ab876",
        "serviceNumber": "500000001",
        "serviceid": "10039",
        "merchant": "ENEO",
        "amountType": "FIXED",
        "localCur": "XAF",
        "amountLocalCur": "200.00",
        "billNumber": "500000001",
        "customerNumber": "QA829294",
        "billMonth": "07",
        "billYear": "2009",
        "billDate": "2009-07-28T00:00:00+00:00",
        "billDueDate": "2025-08-11T00:00:00+00:00",
        "payItemDescr": "COMMODITY - USAGE INVOICE",
        "billType": "REGULAR",
        "penaltyAmount": "0.00",
        "payOrder": 0,
        "optStrg": null,
        "optNmb": null
    }
]
`)
}

// BillEmpty is the response when there is no response
func BillEmpty() []byte {
	return []byte(`[]`)
}
