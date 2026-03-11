package stubs

// AccountGet is the response when getting the `/account` endpoint
func AccountGet() []byte {
	return []byte(`
{
    "balance": 234500.00,
    "currency": "XAF",
    "key": "6B352110-4716-11ED-963F-0800200C9A66",
    "agentId": "AGT-001-00234",
    "agentName": "John Doe",
    "agentAddress": "123 Main Street, Douala",
    "agentPhonenumber": "237699999999",
    "companyName": "Acme Payments Ltd",
    "companyAddress": "456 Business Ave, Yaounde",
    "companyPhonenumber": "237677777777",
    "limitMax": 5000000.00,
    "limitRemaining": 4765500.00
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
