package main

import (
	"BookLibrary/library"
	"BookLibrary/menu"
	"bufio"
	"fmt"
	"os"
)

func ModifyBook(Library *[]library.Book) {

	var modificationChoice string
	fmt.Printf("Select the index of the book you want to modify: \n")
	fmt.Printf("Press q to return to the main menu\n")

	DisplayLibrary(*Library)

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	choice := Scanner.Text()

	if choice == "q" {
		fmt.Printf("\n")
		fmt.Printf("Quitting... \n")
		return
	}

	for i, book := range *Library {
		if choice == fmt.Sprintf("%d", i) {
			fmt.Printf("\n")
			fmt.Printf("You have selected to modify the book: \n")
			fmt.Printf("Title: %v \n", book.Title)
			fmt.Printf("Author: %v \n", book.Author)
			fmt.Printf("Publishing Year: %v \n", book.Publishing_year)
			fmt.Printf("Price: %v \n", book.Price)
			fmt.Printf("\n")

			fmt.Printf("Please enter the detail to modify: \n")

			menu.DisplayBookModificationMenu()

			Scanner.Scan()
			modificationChoice = Scanner.Text()

			if modificationChoice == "5" {
				fmt.Printf("\n")
				fmt.Printf("Quitting... \n")
				return
			}

			if modificationChoice == "1" {
				fmt.Printf("Please enter the new title: \n")
				Scanner.Scan()
				newTitle := Scanner.Text()
				(*Library)[i].Title = newTitle
			}

			if modificationChoice == "2" {
				fmt.Printf("Please enter the new author: \n")
				Scanner.Scan()
				newAuthor := Scanner.Text()
				(*Library)[i].Author = newAuthor
			}

			if modificationChoice == "3" {
				fmt.Printf("Please enter the new publishing year: \n")
				Scanner.Scan()
				newPublishingYear := Scanner.Text()
				(*Library)[i].Publishing_year = newPublishingYear
			}

			if modificationChoice == "4" {
				fmt.Printf("Please enter the new price: \n")
				Scanner.Scan()
				newPrice := "$" + Scanner.Text()
				(*Library)[i].Price = newPrice
			}

		}
	}
}
