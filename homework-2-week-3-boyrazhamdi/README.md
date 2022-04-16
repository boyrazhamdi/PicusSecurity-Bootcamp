## Homework | Week 3 | 

## Description
- In this project you can some queries in the booklist. For example, list, search, get, delete and buy.
- Project is written in go language.
- I faced with some problems in code like getting the input.
- I hope to implement some error handling and exiting the program and using it again at the request of the user in the future.

## How to Run Project
There are functions in project and explained below.
#### list command
Lists all elements in the list.
```
go run main.go list
```
#### search command 
Checks the movie you entered whether exist in the film list.
```
go run main.go search <filmName>
go run main.go search Lord of the Ring: The Return of 
```

#### get command
Returns book that you typed book's id
```
go run main.go get <bookID>
go run main.go get 5
```
#### delete command 
Deletes book that you typed book's id
```
go run main.go delete <bookID>
go run main.go delete 5
```
#### buy command 
Returns the latest information about books after purchasing
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```
