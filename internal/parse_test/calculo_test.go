package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

func readOriginFile(t *testing.T) []string {
	lines, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/009-11804009-C.txt")
	require.NoError(t, err)
	return lines
}

func TestCalculoMustFindBaseData(t *testing.T) {
	lines := readOriginFile(t)
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

func TestCalculoMustFindDesagio35(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Desagio35(lines, calculo)
	require.NoError(t, err)
	require.Equal(t, uint64(134042), calculo.Desagio35().ValorCorrigido())
	require.Equal(t, uint64(174848), calculo.Desagio35().ValorJurosMora())
	require.Equal(t, uint64(308889), calculo.Desagio35().TotalDevido())
}

func TestCalculoMustFindTotalAposDesagio35(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.TotalAposDesagio35(lines, calculo)
	require.NoError(t, err)
	require.Equal(t, uint64(248934), calculo.TotalAposDesagio35().ValorCorrigido())
	require.Equal(t, uint64(324717), calculo.TotalAposDesagio35().ValorJurosMora())
	require.Equal(t, uint64(573651), calculo.TotalAposDesagio35().TotalDevido())
}

func TestCalculoMustFindTotal(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Total(lines, calculo)
	require.NoError(t, err)
	require.Equal(t, uint64(382976), calculo.Total().ValorCorrigido())
	require.Equal(t, uint64(499564), calculo.Total().ValorJurosMora())
	require.Equal(t, uint64(882540), calculo.Total().TotalDevido())
}

func TestCalculoMustFindLinha(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Linha(lines, 1, 15, calculo)
	require.NoError(t, err)
	line := calculo.Linha("fev/97")
	require.NotNil(t, line)
	require.Equal(t, uint64(13861), line.VencimentoBasico())
	require.Equal(t, uint64(18481), line.Soma())
	require.Equal(t, uint64(5334), line.ValorDevido())
}

func TestCalculoMustNotFindLinha(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Linha(lines, 1, 30, calculo)
	require.EqualError(t, err, application.ErrMesAnoNaoEncontrado.Error())
	err = parse.Linha(lines, 3, 1, calculo)
	require.EqualError(t, err, application.ErrMesAnoNaoEncontrado.Error())
}

func TestCalculoMustFindLinhaOnSecondPage(t *testing.T) {
	lines := readOriginFile(t)
	calculo := application.NewCalculo()
	err := parse.Linha(lines, 2, 3, calculo)
	require.NoError(t, err)
	line := calculo.Linha("mai/98")
	require.NotNil(t, line)
	require.Equal(t, uint64(14453), line.VencimentoBasico())
	require.Equal(t, uint64(19271), line.Soma())
	require.Equal(t, uint64(5562), line.ValorDevido())
}
