package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawFilmList := []string{"The Godfather", "The Wizard of Oz", "Citizen Kane", "The Shawshank Redemption", "Pulp Fiction", "Casablanca", "The Godfather: Part II", "E.T. The Extra-Terrestrial", "2001: A Space Odyssey", "Schindler's List", "Star Wars", "Back to the Future", "Raiders of the Lost Ark", "Forrest Gump", "Gone With the Wind", "To Kill a Mockingbird", "Apocalypse Now", "Annie Hall", "Goodfellas", "It's a Wonderful Life"}

	var filmList []string
	for i := 0; i < len(rawFilmList); i++ {
		filmList = append(filmList, strings.ToLower(rawFilmList[i]))
	}

	if len(os.Args) > 1 {
		input1 := os.Args[1]
		if input1 == "list" {
			list(rawFilmList)
		} else if input1 == "search" {
			filmName := make([]string, 0)
			for i := 2; i < len(os.Args); i++ {
				filmName = append(filmName, strings.ToLower(os.Args[i]))
			}
			filmName1 := strings.Join(filmName, " ")
			search(filmName1, filmList)
		}
	} else {
		fmt.Println("You missed an argument!")
	}

}

// list: Lists all elements in the list.
func list(bookList []string) {
	for i := 0; i < len(bookList); i++ {
		println(bookList[i])
	}
}

// search: Checks the movie you entered whether exist in the film list.
func search(film string, filmList []string) {
	for _, i := range filmList {
		if film == i {
			fmt.Println(strings.Title(i))
			return
		}
	}
	fmt.Println("The movie you typed could not be found.")
	return
}
