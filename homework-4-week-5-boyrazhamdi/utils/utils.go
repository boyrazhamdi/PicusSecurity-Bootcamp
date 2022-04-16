package utils

import (
	"encoding/csv"
	"fmt"
	"homework-3-week-5-boyrazhamdi/author"
	"homework-3-week-5-boyrazhamdi/book"
	"os"
	"strconv"
)

func BookCsvToStruct(path string) []book.Book {
	var newbook book.Book
	var books []book.Book
	// Creating a connection via csv files
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while csv read operation ")
	}
	//creating a *csv.reader
	reader := csv.NewReader(csvFile)
	//reading csv files via reader and method of reader
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while reading csv file ")
	}
	//in csv file, each comma represent a column, with this knowledge every row has 12 column and each column parse where they are belong.
	for i, each := range csvData {
		if i > 0 {

			bookid, _ := strconv.Atoi(each[0])
			newbook.ID = bookid
			newbook.Name = each[1]
			newbook.PageNumber, _ = strconv.Atoi(each[2])
			newbook.Stock, _ = strconv.Atoi(each[3])
			newbook.Price, _ = strconv.ParseFloat(each[4], 64)
			newbook.StockCode, _ = strconv.Atoi(each[5])
			newbook.ISBN = each[6]
			newbook.AuthorID, _ = strconv.Atoi(each[7])
			newbook.Deleted, _ = strconv.ParseBool(each[9])
			books = append(books, newbook)
		}

	}

	return books
}
func AuthorCsvToStruct(path string) []author.Author {
	var newauthor author.Author
	var authors []author.Author
	// Creating a connection via csv files
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while csv read operation ")
	}
	//creating a *csv.reader
	reader := csv.NewReader(csvFile)
	//reading csv files via reader and method of reader
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while reading csv file ")
	}
	//in csv file, each comma represent a column, with this knowledge every row has 12 column and each column parse where they are belong.
	for i, each := range csvData {
		if i > 0 {

			authorid, _ := strconv.Atoi(each[0])
			newauthor.ID = authorid
			newauthor.Name = each[1]
			authors = append(authors, newauthor)
		}

	}

	return authors
}
