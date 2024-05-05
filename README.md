# AccessApi

 ### ***This is a very simple service that allows other services to request and check what the user who made the request has access to.***

## Quick Start

### Running the Service

*To run the service:*
```bash
docker-compose build
docker-compose up
```

### Database Management Made Easy

Open http://localhost:8080 after running the service to manage the database using Adminer.

### Admin Requests

#### Admin requests are served on port 7000:

Once the service is running, you will find scripts and request examples for each action that can be performed as an ***admin***.

#### Insert a User:

```bash
curl -d @insertRequest.json -H 'Content-Type:application/json' http://localhost:7000/insertUser
```

***insertRequest.json:***
```json
{
	"username":"admin",
	"services":{
		"archive_manager":{
                        "records":["record1","record2","record3","record4","record5","record0"]
                },
		"task_manager":{
                        "agents":["agent7","agent3","agent0"]
                }
	}	
}
```

    
#### Delete a User:

```bash
curl -d @deleteRequest.json -H 'Content-Type:application/json' http://localhost:7000/deleteUser
```

***deleteRequest.json:***
```json
{
	"username":"admin"
}
```

#### Add New Service Access to a User:

```bash
curl -d @addServicesRequest.json -H 'Content-Type:application/json' http://localhost:7000/addUserServices
```

***addServicesRequest.json:***
```json
{
	"username":"admin",
	"services":{
		"archive_manager":{
                        "records":["newrecord1","newrecord2","newrecord3"]
                },
		"task_manager":{
                        "agents":["newagent1","newagent2","newagent0"]
                }
	}	
}
```

#### Remove User Access for Some Services:

```bash
curl -d @removeServicesRequest.json -H 'Content-Type:application/json' http://localhost:7000/removeUserServices
```

***removeServicesRequest.json:***
```json
{
	"username":"admin",
	"services":{
		"archive_manager":{
                        "records":["oldrecord1","oldrecord2","oldrecord3"]
                },
		"task_manager":{
                        "agents":["oldagent1","oldagent2","oldagent0"]
                }
	}	
}
```

### Client Requests

#### Client requests are served on port 7001:

Once the service is running, you will find scripts and request examples for each action that can be performed as an ***client service***.

#### Check User Access

Before checking user access "client service" must get `TOKEN` from OAUTH2 server that is running on 9096 port.

This is shell script example how to get Token by `clientId` & `secret`

```bash
    curl -X POST   http://localhost:9096/token   -H 'Content-Type: application/x-www-form-urlencoded'   -d 'grant_type=client_credentials&client_id=000000&client_secret=999999'
```
Or you can check it by following the link: http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read

OAUTH2 Server Response:
```json
{
  "access_token": "YTHMNDFMNTITZTYYNS0ZNME2LWEYMZITOTDKMGNLZTG5MGI2",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

After getting token "client service" can make request to our "access service" on 7001 port:

Request example in bash script:

```bash
    curl -d @clientRequest.json -H 'Content-Type:application/json' http://localhost:7001/checkUserAccess
```

clientRequest.json
```json
{
    "username": "admin",
    "token":"ZDU2ZJG1ZJITZTC2MI0ZNZJJLTLJYZUTZWU3MZC2OGY1ZDZI"
}
```