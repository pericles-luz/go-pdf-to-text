package main

import (
	"fmt"
	"log"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/pkg/convert"
)

func main() {
	processor := service_fee.NewSummary()
	err := convert.Walk("/mnt/c/Users/peric/Downloads/PDFs", processor)
	if err != nil {
		log.Fatal(err)
	}
	for _, summary := range processor.Summaries() {
		fmt.Println(summary.ExecutionNumber())
		fmt.Println(summary.MainProcess())
		fmt.Println(summary.Total().Total())
	}
}
