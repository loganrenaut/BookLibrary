package main

import (
	"BookLibrary/library"
	"fmt"
	"os"
	"text/tabwriter"
)

func DisplayLibrary(Library []library.Book) {
	fmt.Printf("\n")
	fmt.Printf("Current library contents: \n")
	fmt.Printf("\n")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "Index\tTitle\tAuthor\tPublishing Year\tPrice")
	for i, book := range Library {
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n", i, book.Title, book.Author, book.Publishing_year, book.Price)
	}
	writer.Flush()
}
