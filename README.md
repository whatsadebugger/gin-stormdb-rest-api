# Address Book

to download the packages needed 

`go get -d ./... `

a fresh database called address.db will get created when you run the program and it will persist\
to run do a go build in the project folder. 

after you can run with `./gin-stormdb-rest-api`  

port number is 8080

check handlers_test.go to see how to call each route

available routes

```
POST   /address
GET    /address
GET    /address/:id
PUT    /address
DELETE /address/:id
POST   /address/upload
GET    /addressbook
```
