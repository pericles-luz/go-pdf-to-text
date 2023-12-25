package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

func readOriginFile(t *testing.T, file string) []string {
	lines, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/" + file)
	require.NoError(t, err)
	return lines
}

func TestCalculoMustFindBaseData(t *testing.T) {
	lines := readOriginFile(t, "009-11804009-C.txt")
	calculo := application.NewCalculo()
	err := parse.CalculoBase(lines, calculo)
	require.NoError(t, err)
	require.Equal(t, "0001770-45.2013.4.05.8100", calculo.ProcessoNumero())
	require.Equal(t, "0006379-33.1997.4.05.8100", calculo.ProcessoPrincipal())
	require.Equal(t, "25/03/1997", calculo.Ajuizamento())
	require.Equal(t, "22/05/1997", calculo.Citacao())
	require.Equal(t, "01/01/2020", calculo.Calculo())
	require.Equal(t, "EDUARDO GOMES DE MEDEIROS", calculo.Exequente())
	require.Equal(t, "03410702709", calculo.Cpf())
	require.Equal(t, "11804009", calculo.IdUnica())
}
