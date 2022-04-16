package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Books struct {
	Books []Book `json:"booklist"`
}

type Book struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Pages     int    `json:"pages"`
	Stocks    int    `json:"stocks"`
	Price     int    `json:"price"`
	StockCode string `json:"stockCode"`
	ISBN      string `json:"ISBN"`
	Author    Author `json:"author"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var (
	command  = flag.String("command", "", "List of Commands")
	name     = flag.String("name", "", "Book Name")
	id       = flag.Int("id", 0, "Book ID")
	quantity = flag.Int("quantity", 0, "Book Quantity ")
)

var usage = `Usage: [commands...] [params...]
Options:
-command
-name
-id
-quantity
`

var books Books

func main() {
	flag.Usage = func() {
		fmt.Print(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	readFile("booklist.json")

	switch *command {
	case "list":
		list(books)
	case "search":
		search(*name, books)
	case "get":
		get(*id, books)
	case "delete":
		delete(*id, books)
	case "buy":
		buy(*id, *quantity, books)
	default:
		fmt.Print("Wrong Command!")
		flag.Usage()
	}

}

func readFile(input string) {
	jsonFile, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &books)
}

// list: Lists all elements in the list.
func list(books Books) {
	for _, i := range books.Books {
		fmt.Println(i.Name)
	}
}

// search: Checks the book you entered whether exist in the book list.
func search(book string, books Books) {
	for _, i := range books.Books {
		if book == i.Name {
			fmt.Println(strings.Title(i.Name))
			return
		}
	}
	fmt.Println("The book you typed could not be found.")
	return
}

// get: Returns book that you typed book's id
func get(id int, books Books) {
	for _, i := range books.Books {
		if id == i.Id {
			fmt.Println(strings.Title(i.Name))
			return
		}
	}
	fmt.Println("The book id you typed could not be found.")
	return
}

// delete: deletes book that you typed book's id
func delete(id int, books Books) {
	for _, i := range books.Books {
		if id == i.Id {
			books.Books = append(books.Books[:i.Id-1], books.Books[i.Id:]...)
			fmt.Println(strings.Title(i.Name), "is deleted")
			return
		}
	}
	fmt.Println("The book id you typed could not be found.")
	return

}

// buy: Returns the latest information about books after purchasing
func buy(id int, quantity int, books Books) {
	for _, i := range books.Books {
		if id == i.Id {
			if i.Stocks >= quantity && quantity > 0 {
				i.Stocks -= quantity
				fmt.Println(i.Id)
				fmt.Println(i.Name)
				fmt.Println(i.Pages)
				fmt.Println(i.Stocks)
				fmt.Println(i.StockCode)
				fmt.Println(i.ISBN)
				fmt.Println(i.Author.Id)
				fmt.Println(i.Author.Name)
				return
			}

			fmt.Printf("We have just %d sample of this book left.", i.Stocks)
			return
		}

	}
}
