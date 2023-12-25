package parse_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/pericles-luz/go-pdf-to-text/internal/parse"
	"github.com/stretchr/testify/require"
)

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
