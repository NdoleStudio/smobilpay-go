package stubs

// SubscriptionGet is the response for a subscription payment
func SubscriptionGet() []byte {
	return []byte(`
[
  {
    "serviceNumber": "500000001",
    "serviceid": "10042",
    "merchant": "CMENEOPREPAID",
    "payItemId": "string",
    "payItemDescr": "string",
    "amountType": "FIXED",
    "name": "string",
    "localCur": "str",
    "amountLocalCur": 0,
    "customerReference": "string",
    "customerName": "string",
    "customerNumber": "string",
    "startDate": "2019-08-24",
    "dueDate": "2019-08-24",
    "endDate": "2019-08-24",
    "optStrg": "string",
    "optNmb": 0
  }
]
`)
}
