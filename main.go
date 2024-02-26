package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/pericles-luz/go-pdf-to-text/pkg/convert"
)

func main() {
	start := time.Now()
	processor := service_fee.NewSummary()
	// path := "/mnt/c/Users/peric/Downloads/PDFs"
	path := filepath.Join("/", "mnt", "c", "Users", "peric", "Downloads", "PDFs")
	err := convert.Walk(path, processor)
	fmt.Println("fim do walk:", time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	for _, summary := range processor.Summaries() {
		fmt.Println(summary.LocalExecutionNumber(), summary.ExecutionNumber(), summary.MainProcess(), summary.CalculateTotal())
	}
	err = convert.GenerateSuccumbence(path, processor)
	fmt.Println("fim do generate:", time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
}
