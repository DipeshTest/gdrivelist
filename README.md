---
title: List GDrive Files
weight: 1
---

# List GDrive Files
This activity lists the files from Gdrive account. This activity is built by Team AllStars

## Installation

### Flogo CLI
```bash
flogo install github.com/DipeshTest/gdrivelist
```

### Third-party libraries used
- #### package drive - "google.golang.org/api/drive/v3":
	Package drive provides access to the Drive API. For more details, check https://developers.google.com/drive/
- #### package googleapi - "google.golang.org/api/googleapi":
	Package googleapi contains the common code shared by all Google API libraries.
	
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
			]
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
Please refer activity_test.go


## Response Codes
### Google Drive List
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful and the response body contains the representation requested. *If the file is not present, the response count would be 0|
|105 |INVALID INPUT| Access token is not specified.|
|400 |SERVER ERROR| Invalid values specified for either pageSize, nextPageToken or Timeout.|
|401 |AUTHENTICATION ERROR| Invalid Access Token.|

Note - Please refer link - https://developers.google.com/doubleclick-search/v2/standard-error-responses for other standard error responses from Google 