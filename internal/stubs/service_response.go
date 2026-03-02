package stubs

// ServiceGetAllOk returns the api response to the `GET /service` endpoint
func ServiceGetAllOk() []byte {
	return []byte(`
[
    {
        "serviceid": "20062",
        "merchant": "CMORANGE",
        "title": "ORANGE Recharge/Airtime",
        "description": "Recharge de credit de communication Orange",
        "category": "Airtime",
        "country": "CMR",
        "localCur": "XAF",
        "type": "TOPUP",
        "status": "Active",
        "isReqCustomerName": 0,
        "isReqCustomerAddress": 0,
        "isReqCustomerNumber": 0,
        "denomination": 50,
        "isReqServiceNumber": 1,
        "isVerifiable": false,
        "hint": [
            {
                "language": "en",
                "localText": "Enter phone number in the format 69#######. The Minimum amount for this service is 100 FCFA"
            },
            {
                "language": "fr",
                "localText": "Entrez le numéro de téléphone suivant le format 69#######"
            }
        ],
        "validationMask": "^(237)?((655|656|657|658|659|686|687|688|689|640)[0-9]{6}$|(69[0-9]{7})$)",
        "labelCustomerNumber": null,
        "labelServiceNumber": [
            {
                "localText": "ORANGE Number",
                "language": "en"
            },
            {
                "localText": "Numéro ORANGE",
                "language": "fr"
            }
        ]
    },
    {
        "serviceid": "90003",
        "merchant": "CMMTNMOMO",
        "title": "MTN MoMo Cash-In / Depot",
        "description": "The operation consists in recharging a MTN Mobile Money account with credits for cash",
        "category": "Cashin",
        "country": "CMR",
        "localCur": "XAF",
        "type": "CASHIN",
        "status": "Active",
        "isReqCustomerName": 0,
        "isReqCustomerAddress": 0,
        "isReqCustomerNumber": 0,
        "denomination": 1,
        "isReqServiceNumber": 1,
        "isVerifiable": false,
        "hint": null,
        "validationMask": "^(237|00237|\\+237)?((650|651|652|653|654|680|681|682|683|684)[0-9]{6}$|(67[0-9]{7})$)",
        "labelCustomerNumber": null,
        "labelServiceNumber": [
            {
                "localText": "MTN Number",
                "language": "en"
            },
            {
                "localText": "Numéro MTN",
                "language": "fr"
            }
        ]
    }
]
`)
}

// ServiceGetOk returns the api response to the `GET /service/{id}` endpoint
func ServiceGetOk() []byte {
	return []byte(`
{
    "serviceid": "20062",
    "merchant": "CMORANGE",
    "title": "ORANGE Recharge/Airtime",
    "description": "Recharge de credit de communication Orange",
    "category": "Airtime",
    "country": "CMR",
    "localCur": "XAF",
    "type": "TOPUP",
    "status": "Active",
    "isReqCustomerName": 0,
    "isReqCustomerAddress": 0,
    "isReqCustomerNumber": 0,
    "denomination": 50,
    "isReqServiceNumber": 1,
    "isVerifiable": false,
    "hint": [
        {
            "language": "en",
            "localText": "Enter phone number in the format 69#######. The Minimum amount for this service is 100 FCFA"
        }
    ],
    "validationMask": "^(237)?((655|656|657|658|659|686|687|688|689|640)[0-9]{6}$|(69[0-9]{7})$)",
    "labelCustomerNumber": null,
    "labelServiceNumber": [
        {
            "localText": "ORANGE Number",
            "language": "en"
        },
        {
            "localText": "Numéro ORANGE",
            "language": "fr"
        }
    ]
}
`)
}

// ServiceGetError returns an error response for the `/service` endpoint
func ServiceGetError() []byte {
	return []byte(`
{
    "devMsg": "Service unknown",
    "usrMsg": "Service unknown",
    "respCode": 40602,
    "link": "http://support.maviance.com/"
}
`)
}
