package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

func TestTableMustReadEntireTable(t *testing.T) {
	lines := readOriginFile(t, "009-11804009-C.txt")
	calculo := application.NewCalculo()
	err := parse.Table(lines, calculo)
	require.NoError(t, err)
	require.NotNil(t, calculo.Linha("dez/95"))
	require.NotNil(t, calculo.Linha("jun/98"))
	require.Len(t, calculo.Table(), 31)
	require.Equal(t, uint64(882540), calculo.TotalDevido())
}

func TestTableMustReadEntireTable2(t *testing.T) {
	lines := readOriginFile(t, "004-01139401-C.txt")
	calculo := application.NewCalculo()
	err := parse.Table(lines, calculo)
	require.NoError(t, err)
	require.NotNil(t, calculo.Linha("jan/93"))
	require.NotNil(t, calculo.Linha("jun/98"))
	require.Len(t, calculo.Table(), 66)
	require.Equal(t, uint64(3065957), calculo.TotalDevido())
}
