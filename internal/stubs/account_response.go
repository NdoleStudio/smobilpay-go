package stubs

// AccountGet is the response when getting the `/account` endpoint
func AccountGet() []byte {
	return []byte(`
{
    "balance": 417986.82,
    "currency": "XAF",
    "key": "B835273F-0709-AE0C-E01D-B3411C452EEF",
    "agentId": "CM10714-38167",
    "agentName": "ACME S3P AGENT",
    "agentAddress": "Adjacent hospital of Buea ",
    "agentPhonenumber": "(+237) 673658376",
    "companyName": "ACME LTD",
    "companyAddress": "Adjacent Apostolic church Buea,526,Buea",
    "companyPhonenumber": "(+237) 673658376",
    "limitRemaining": 8950428,
    "limitMax": "10000000.00"
}
`)
}

// AccountGetError is the account error response
func AccountGetError() []byte {
	return []byte(`
{
    "devMsg": "Unauthorized",
    "usrMsg": "Request could not be authenticated",
    "respCode": 40100,
    "link": "http://support.maviance.com/"
}
`)
}
