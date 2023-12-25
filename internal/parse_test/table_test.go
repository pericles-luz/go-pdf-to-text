package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

func TestTableMustReadEntireTable(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Table(lines, calculo)
	require.NoError(t, err)
	require.NotNil(t, calculo.Linha("dez/95"))
	require.NotNil(t, calculo.Linha("jun/98"))
	require.Len(t, calculo.Table(), 31)
}
