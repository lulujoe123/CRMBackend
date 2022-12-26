# CRM Baclend

This project is to build an HTTP server with Go: the CRM Backend. <br/>
The application include serveal different RESTful enpoints for users to be able to perform CRUD operations. <br/>

## Installation

download the external 3rd party modules <br/>
github.com/gorilla/mux - Used as the http router go mod init go mod tidy


## Launch

go run main.go - run the HTTP server service
go test - run the unit tests

## Endpoints

Getting a single customer through a /customers/{id} path <br/>
Getting all customers through a the /customers path <br/>
Creating a customer through a /customers path <br/>
Updating a customer through a /customers/{id} path <br/>
Deleting a customer through a /customers/{id} path <br/>

