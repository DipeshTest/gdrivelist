---
title: List GDrive Files
weight: 1
---

# Counter
This activity lists the files from Gdrive account.

## Installation
### Flogo Web
This activity is built by Team AllStars
### Flogo CLI
```bash
flogo add activity github.com/DipeshTest/gdrivelist
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
		"name": "accessToken",
		"type": "string",
		"required": true
	},
	{
		"name": "fileName",
		"type": "string"
	},
	{
		"name": "orderBy",
		"type": "string",
		"allowed": [
			"createdTime",
			"modifiedByMeTime",
			"modifiedTime",
			"name",
			"recency",
			"starred",
			"viewedByMeTime",
			"sharedWithMeTime",
			"quotaBytesUsed"
	},
	{
		"name": "pageSize",
		"type": "int",
		"value": 50
	},
	{
		"name": "nextPageToken",
		"type": "string"

	},
	{
		"name": "timeout",
		"type": "string",
		"value": "120"
	}
  ],
  "outputs": [
    {
      "name": "statusCode",
      "type": "string"
    },
    {
      "name": "message",
      "type": "any"
    },
    {
      "name": "fileCount",
      "type": "int"
    },
    {
      "name": "nextPageToken",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessToken | True     | The access token for your account |         
| fileName   | False    | Name of the file to search |
| orderBy    | False     | Key to sort list of files |  
| pageSize   | False     | The maximum number of files to return per page. Values must be within the range: [1, 1000], default is 50 |
| nextPageToken | False  | The token for continuing a previous list request on the next page. This should be set to the value of 'nextPageToken' from the previous response | 
| timeout       | False    | Timeout value for the activity, default is 120 seconds|

## Examples
### Increment
The below example for a sample list:

```json
{
	"accessToken": "ya29.GlurBW7n5A2Fk_rstX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv9cnhsUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqFR8",
	"fileName": "MyFile.pdf",
	"pageSize": "50",
	"nextPageToken":"~!!~AI9FV7QPNqNJgQnM3rJ7F3mJ-jesiPQQAmsi1GisbVc-bR0uqIRH5onEFtOWn7og0rjo9-qntk0u8JCnKsiyyrRbpUKrdyFqpLEKEtFGkOaJWGPxQCIKgmMG9GstFEh5zRL4zgZwxRtGK__tYgJARsvBAHmHw7iAPRODS97FK6_kk-xNDTVREPyH3ZBlYyz4mcqnNg4W4syWyRXscUes8IjsJXlEfwzZoV_4ejd21UTDfP_9EuyKpdQPcbI-GWLVqhESJoQkMOqhQ46eHEHu1qTQIHZJGnc7wg6HOxoynvaysf1ZDF4PYeHKIzdFBAnq_0gHBi3i-ciYGXRy8NueS3e_ZLX4COjT_5lUjHzixAgriGooLa_BkApHOOl6ACjY8d6yd_30TpGx-394pOqXzxdsMVBEXwD_BaxFUUnqI46JTs1jVSzTSdCaXhK_8sso_MOMw5KdnxqZq",
	"timeout": "120"
}
```


## Response Codes
### Google Drive Create
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful and the response body contains the representation requested. *If the file is not present, the response count would be 0|
|105 |INVALID INPUT| Access token is not specified.|
|400 |SERVER ERROR| Invalid values specified for either pageSize, nextPageToken or Timeout|
|401 |AUTHENTICATION ERROR| Invalid Access Token.|
