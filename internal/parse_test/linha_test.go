package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

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
	require.Equal(t, uint64(256866867), line.IndiceCorrecao())
	require.Equal(t, uint64(13701), line.ValorCorrigido())
	require.Equal(t, uint64(1318171), line.JurosMora())
	require.Equal(t, uint64(18060), line.ValorJurosMora())
	require.Equal(t, uint64(31760), line.TotalDevido())
	require.Equal(t, uint64(2886), line.Percentual())
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
	require.Equal(t, uint64(243423518), line.IndiceCorrecao())
	require.Equal(t, uint64(13538), line.ValorCorrigido())
	require.Equal(t, uint64(1258171), line.JurosMora())
	require.Equal(t, uint64(17033), line.ValorJurosMora())
	require.Equal(t, uint64(30571), line.TotalDevido())
	require.Equal(t, uint64(2886), line.Percentual())
}
