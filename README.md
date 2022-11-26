## Companies Manager REST API
A microservice of company data

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup Instructions](#setup-instructions)
* [Endpoints and Calls](#endpoints-and-calls)

## General info
The purpose of this project is to create a microservice that handles data about companies. To achieve that functions of Insertion, Editing, Retrieval and Deletion
of company data in the set database are supported. To communicate data to the service REST API functions are used.

## Technologies
Technologies used for this project and required are:
* Docker version 20.10.16
* Docker Compose version v2.6.0
* Go version 1.17
* Postgres version 12.2
* Task version v3.18.0
* git version 2.34.1.windows.1

## Setup Instructions

Open a terminal and exexute the command:
```
git clone https://github.com/melinaco4/companies-manager
```
!Make sure docker is running

Inside the companies-manager directory execute the command:
```
task build
```
And in the same directory then the command
```
task run
```

You are ready to go!

## Endpoints and Calls

The following are the REST API calls that are supported:

### Create a new Company

#### To create a company send a POST request to the endpoint:
```
http://localhost:8080/api/company
```
The Headers of the request should be:
```
* Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.rBU55ySwV_E-ltd7EcNAL3laGC0HjbsSOT39FcSysyI
* Content-Type: application/json
```

And the Body of the request should be:
```
{
  "name":<Required field: enter a name of type string with at least 2 letters and maximum 15>,
  "description": <Not Required field: enter a description of type string>,
  "amountofemployees": <Required field: enter an integer>,
  "registered": <Required Field: a boolean type, true or false>,
  "type": <Required Field: enter a type, should of type string and one of those options:Corporations | NonProfit | Cooperative | Sole Proprietorship>
}
```

### Get one Company

#### To get one company send a GET request to the endpoint:

```
http://localhost:8080/api/company/<id>
```
Where <id> represents an id of uuid type.

The Headers of the request should be:
```
* Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.rBU55ySwV_E-ltd7EcNAL3laGC0HjbsSOT39FcSysyI
```

### Update data of a Company

#### To update a company send a PATCH request to the endpoint:

```
http://localhost:8080/api/company/<id>
```
Where <id> represents an id of uuid type.

The Headers of the request should be:
```
* Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.rBU55ySwV_E-ltd7EcNAL3laGC0HjbsSOT39FcSysyI
* Content-Type: application/json
```

And the Body of the request should be:
```
{
  "name":<Required field: enter a name of type string with at least 2 letters and maximum 15>,
  "description": <if the description changed, enter the new description of type string>,
  "amountofemployees": <if the amount of employees changed, enter the new integer>,
  "registered": <if the registered parameter changed, enter the new one of type boolean, true or false>,
  "type": <if the type changed, enter the type, should of type string and one of those options:Corporations | NonProfit | Cooperative | Sole Proprietorship>
}
```

### Delete a Company

#### To delete a company send a DELETE request to the endpoint:

```
http://localhost:8080/api/company/<id>
```
Where <id> represents an id of uuid type.

The Headers of the request should be:
```
* Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.rBU55ySwV_E-ltd7EcNAL3laGC0HjbsSOT39FcSysyI
```

