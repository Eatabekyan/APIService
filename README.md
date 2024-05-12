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
                        "access":["record1","record2","record3","record4","record5","record0"]
                },
		"task_manager":{
                        "access":["agent7","agent3","agent0"]
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
                        "access":["newrecord1","newrecord2","newrecord3"]
                },
		"task_manager":{
                        "access":["newagent1","newagent2","newagent0"]
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
                        "access":["oldrecord1","oldrecord2","oldrecord3"]
                },
		"task_manager":{
                        "access":["oldagent1","oldagent2","oldagent0"]
                }
	}	
}
```

### Client Requests

#### Client requests are served on port 7001:

Once the service is running, you will find scripts and request examples for each action that can be performed as a ***client service***.

#### Check User Access

Before checking user access, the "client service" must obtain a `TOKEN` from the OAuth2 server running on port 9096.

Here is a shell script example of how to obtain a Token using `clientId` and `secret`:

```bash
curl -X POST   http://localhost:9096/token   -H 'Content-Type: application/x-www-form-urlencoded'   -d 'grant_type=client_credentials&client_id=000000&client_secret=999999'
```

Alternatively, you can obtain the token by following this link: [http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read](http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read)

*I've decided to use OAuth2 because OpenId is a wrapper over it.*

OAuth2 Server Response:
```json
{
  "access_token": "YTHMNDFMNTITZTYYNS0ZNME2LWEYMZITOTDKMGNLZTG5MGI2",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

After obtaining the token, the "client service" can make a request to our "access service" on port 7001:

Request example in a bash script:

```bash
curl -d @clientRequest.json -H 'Content-Type:application/json' http://localhost:7001/checkUserAccess
```

***clientRequest.json***:
```json
{
    "username": "admin",
    "token":"ZDU2ZJG1ZJITZTC2MI0ZNZJJLTLJYZUTZWU3MZC2OGY1ZDZI"
}
```



`Important:`
        After running project all the logs of accessing service are written in */logs/service_logs.log*

###  Mocks for Task Manager & Archive Manager

##### To test service work you can send requests to mocks

#### Task Manager Mock

***Script to send request as user that wants access to some agent:***
```bash
curl -d @task_manager_request.json -H 'Content-Type:application/json' http://localhost:5050/requestTask
```

***task_manager_request.json***:
```json
{
	"user": "admin",
	"agent": "some_agent"
}
```

**Response:** "Access Accepted" or "Access Denied".

#### Archive Manager Mock

***Script to send request as user that wants access to some record:***
```bash
curl -d @archive_manager_request.json -H 'Content-Type:application/json' http://localhost:6060/requestArchive
```

***archive_manager_request.json***:
```json
{
	"user": "admin",
	"record": "record"
}
```

**Response:** "Access Accepted" or "Access Denied".

