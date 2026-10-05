package main

import (
	"BookLibrary/library"
	"bufio"
	"fmt"
	"os"
)

var deletedBooks = []library.Book{}
var choice string

func DeleteBook(Library *[]library.Book) {

	fmt.Printf("Select the index of the book you want to delete: \n")
	fmt.Printf("Press q to return to the main menu\n")

	DisplayLibrary(*Library)

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	choice = Scanner.Text()

	if choice == "q" {
		fmt.Printf("\n")
		fmt.Printf("Quitting... \n")
		return
	}

	for i, book := range *Library {
		if choice == fmt.Sprintf("%d", i) {
			deletedBooks = append(deletedBooks, book)
			*Library = append((*Library)[:i], (*Library)[i+1:]...)
			fmt.Printf("\n")
			fmt.Printf("Book deleted successfully \n")
			fmt.Printf("\n")
			DisplayLibrary(*Library)
		}
	}

}
