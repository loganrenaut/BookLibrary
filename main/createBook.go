package main

import (
	"BookLibrary/library"
	"bufio"
	"fmt"
	"os"
)

var Scanner = bufio.NewScanner(os.Stdin)

func CreateBook() library.Book {

	fmt.Printf("Please enter the title of the book: \n")
	Scanner.Scan()
	title := Scanner.Text()

	fmt.Printf("Please enter the author of the book: \n")
	Scanner.Scan()
	author := Scanner.Text()

	fmt.Printf("Please enter the year the book was published: \n")
	Scanner.Scan()
	publishing_year := Scanner.Text()

	fmt.Printf("Please enter the price of the book: \n")
	Scanner.Scan()
	price := "$" + Scanner.Text()

	book := library.Book{
		Title:           title,
		Author:          author,
		Publishing_year: publishing_year,
		Price:           price,
	}
	return book
}

func StoreBook(book library.Book, Library *[]library.Book) {
	*Library = append(*Library, book)
}

func DisplayBook(book library.Book) {
	fmt.Printf("\n")
	fmt.Printf("Book added details: \n")
	fmt.Printf("Title: %v \n", book.Title)
	fmt.Printf("Author: %v \n", book.Author)
	fmt.Printf("Publishing Year: %v \n", book.Publishing_year)
	fmt.Printf("Price: %v \n", book.Price)
}
