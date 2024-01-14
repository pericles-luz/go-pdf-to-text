package service_fee

import (
	"os/exec"
	"path/filepath"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
)

type Summary struct {
	Summary   *application_fee.Summary
	summaries []*application_fee.Summary
	lines     []string
}

func NewSummary() *Summary {
	return &Summary{}
}

func (s *Summary) Parse(path string) error {
	if err := s.loadFile(path); err != nil {
		return err
	}
	summary := application_fee.NewSummary()
	if !summary.IsFeesFile(s.lines) {
		return nil
	}
	if err := summary.Parse(s.lines); err != nil {
		return err
	}
	if err := summary.Validate(); err != nil {
		return err
	}
	s.AddSummary(summary)
	return nil
}

func (s *Summary) loadFile(path string) error {
	if !utils.FileExists(path) {
		return application.ErrArquivoNaoEncontrado
	}
	lines, err := extract.ReadLinesFromFile(path)
	if err != nil {
		return err
	}
	s.lines = lines
	return nil
}

func (s *Summary) GenerateTextFile(path string) error {
	if !utils.FileExists(path) {
		return application.ErrArquivoNaoEncontrado
	}
	if filepath.Ext(path) != ".pdf" {
		return application.ErrArquivoInvalido
	}
	// log.Println("pdftotext", "-layout", "-nopgbrk", path, path[:len(path)-4]+".txt")
	if utils.FileExists(path[:len(path)-4] + ".txt") {
		// log.Println("Arquivo já existe")
		return nil
	}
	err := exec.Command("pdftotext", "-layout", "-nopgbrk", path, path[:len(path)-4]+".txt").Run()
	if err != nil {
		return err
	}
	return nil
}

func (s *Summary) ProcessFile(path string) error {
	if !utils.FileExists(path) {
		return application.ErrArquivoNaoEncontrado
	}
	if filepath.Ext(path) == ".txt" {
		return s.Parse(path)
	}
	if filepath.Ext(path) == ".pdf" {
		return s.GenerateTextFile(path)
	}
	return nil
}

func (s *Summary) Summaries() []*application_fee.Summary {
	return s.summaries
}

func (s *Summary) AddSummary(summary *application_fee.Summary) {
	s.summaries = append(s.summaries, summary)
}
