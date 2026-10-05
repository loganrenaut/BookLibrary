package main

import (
	"BookLibrary/library"
	"BookLibrary/menu"
	"bufio"
	"fmt"
	"os"
)

func Interface() {

	var choice string
	var additionalBook library.Book
	Scanner := bufio.NewScanner(os.Stdin)

	book1 := library.Book{
		Title:           "The Great Gatsby",
		Author:          "F. Scott Fitzgerald",
		Publishing_year: "1925",
		Price:           "$10.99",
	}
	library.Library = append(library.Library, book1)
	book2 := library.Book{
		Title:           "To Kill a Mockingbird",
		Author:          "Harper Lee",
		Publishing_year: "1960",
		Price:           "$12.99",
	}
	library.Library = append(library.Library, book2)
	book3 := library.Book{
		Title:           "1984",
		Author:          "George Orwell",
		Publishing_year: "1948",
		Price:           "$11.99",
	}
	library.Library = append(library.Library, book3)
	book4 := library.Book{
		Title:           "Pride and Prejudice",
		Author:          "Jane Austen",
		Publishing_year: "1813",
		Price:           "$10.99",
	}
	library.Library = append(library.Library, book4)
	book5 := library.Book{
		Title:           "The Catcher in the Rye",
		Author:          "J.D. Salinger",
		Publishing_year: "1951",
		Price:           "$10.99",
	}
	library.Library = append(library.Library, book5)

	for choice != "6" {

		menu.DisplayMenu()

		Scanner.Scan()

		choice = Scanner.Text()

		switch choice {
		case "1":
			fmt.Printf("\n")
			fmt.Printf("You have selected to add a book \n")
			additionalBook = CreateBook()
			StoreBook(additionalBook, &library.Library)
			DisplayBook(additionalBook)
			fmt.Printf("\n")
			break

		case "2":
			fmt.Printf("\n")
			fmt.Printf("You have selected to delete a book \n")
			DeleteBook(&library.Library)
			fmt.Printf("\n")
			break

		case "3":
			fmt.Printf("\n")
			fmt.Printf("You have selected to modify a book \n")
			ModifyBook(&library.Library)
			fmt.Printf("\n")
			break

		case "4":
			fmt.Printf("\n")
			fmt.Printf("You have selected to view the library \n")
			fmt.Printf("\n")
			DisplayLibrary(library.Library)
			break

		case "5":
			fmt.Printf("\n")
			fmt.Printf("You have selected to quit \n")
			fmt.Printf("\n")
			return

		default:
			fmt.Printf("\n")
			fmt.Printf("Invalid selection, please try again \n")
			fmt.Printf("\n")
			menu.DisplayMenu()
			Scanner.Scan()
			choice = Scanner.Text()
		}
	}

}
