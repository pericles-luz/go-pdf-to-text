package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/pkg/convert"
)

func main() {
	processor := service_fee.NewSummary()
	// path := "/mnt/c/Users/peric/Downloads/PDFs"
	path := filepath.Join("/", "mnt", "c", "Users", "peric", "Downloads", "PDFs")
	err := convert.Walk(path, processor)
	if err != nil {
		log.Fatal(err)
	}
	for _, summary := range processor.Summaries() {
		fmt.Println(summary.LocalExecutionNumber(), summary.ExecutionNumber(), summary.MainProcess(), summary.CalculateTotal())
	}
	err = convert.GenerateSuccumbence(path, processor)
	if err != nil {
		log.Fatal(err)
	}
}
