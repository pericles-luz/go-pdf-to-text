package convert_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service"
	"github.com/pericles-luz/go-pdf-to-text/pkg/convert"
	"github.com/stretchr/testify/require"
)

func TestWalkMustProcessDirectory(t *testing.T) {
	err := convert.Walk(utils.GetBaseDirectory("pdf"), &service.Calculo{})
	require.NoError(t, err)
}

func TestWalkMustProcessPDFDirectory(t *testing.T) {
	t.Skip("This test takes too long to run")
	err := convert.Walk("/mnt/c/Users/peric/Downloads/PDFs", &service.Calculo{})
	require.NoError(t, err)
}
