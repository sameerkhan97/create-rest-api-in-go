# create-rest-api-in-go

If you are writing any form of web application, then you are most likely interfacing with 1 or more REST APIs in order to populate
the dynamic parts of your application and to perform tasks such as updating or deleting data within a database.

In this repository we are going to build a REST API that exposes GET, POST, DELETE and PUT endpoints that will subsequently allow 
you to perform the full range of CRUD operations.

In order to keep this simple and focus on the basic concepts, we won’t be interacting with any backend database technologies to store
the articles that we’ll be playing with. However, we will be writing this REST API in such a way that it will be easy to update the  
functions we will be defining so that they make subsequent calls to a database to perform any necessary CRUD operations 

## Topics : 
* [Creating a demo http server in GO](https://github.com/sameerkhan97/create-rest-api-in-go/blob/master/creatingHttpServer.go)
* [Simple read pertion on simple routes](https://github.com/sameerkhan97/create-rest-api-in-go/blob/master/readOperationWithoutMuxRouter.go)
* [Read operation using mux router](https://github.com/sameerkhan97/create-rest-api-in-go/blob/master/readOperationUsingRouter.go)
* [Create operation](https://github.com/sameerkhan97/create-rest-api-in-go/blob/master/createOperation.go)
* [Delete operation](https://github.com/sameerkhan97/create-rest-api-in-go/blob/master/deleteOperation.go)


## instructions to run code:
 * prerequisite: - You will need Go version 1.11+ installed on your development machine.
 STEPS:-
 1] Create a folder on your root directory
 2] Create a go file in the folder and paste code from repository
 3] Open the file in your IDE.
 4] In console Terminal run 'go mod init' command and then run 'go mod download' command 
 5] Now run the go file using 'go run filename.go' command
 6] Allow the admin permission for access.
 7] Once this is done head over to your browser and head to 'http://localhost:5060/world' (you can choose your own adress instead of 5060.
 
