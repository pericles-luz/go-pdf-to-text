package convert

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
)

type processFile interface {
	ProcessFile(path string) error
}

// Walk is a function that walks through a directory and process each file
// if a directory is found, it will walk through it recursively
func Walk(path string, p processFile) error {
	if !utils.FileExists(path) {
		return application.ErrArquivoNaoEncontrado
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			err = Walk(filepath.Join(path, file.Name()), p)
			if err != nil {
				return err
			}
			continue
		}
		err = p.ProcessFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("error processing file(%s): %s\n", filepath.Join(path, file.Name()), err.Error())
		}
	}
	return nil
}
