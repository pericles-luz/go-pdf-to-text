package application_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/stretchr/testify/require"
)

func TestLineMustValidate(t *testing.T) {
	line := application_fee.NewLine()
	line.SetSequence(1)
	line.SetCPF("00000000000")
	line.SetName("PERICLES LUZ")
	line.SetUniqueID("00000000000")
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	line.SetStatus("OK")
	err := line.Validate()
	require.NoError(t, err)
}

func TestLineMustParseWithNoValue(t *testing.T) {
	line := application_fee.NewLine()
	err := line.Parse("2     RONALDO ASSUNCAO JACOMINI                       176.757.826-15         1495275                         1.000,00               100,00           900,00                90,00   Não Consta CPF na Lista")
	require.NoError(t, err)
	require.Equal(t, uint16(2), line.Sequence())
	require.Equal(t, "17675782615", line.CPF())
	require.Equal(t, "RONALDO ASSUNCAO JACOMINI", line.Name())
	require.Equal(t, "1495275", line.UniqueID())
	require.Equal(t, uint64(100000), line.Main())
	require.Equal(t, uint64(10000), line.Interest())
	require.Equal(t, uint64(90000), line.Total())
	require.Equal(t, uint64(9000), line.Fees())
	require.Equal(t, "Não Consta CPF na Lista", line.Status())
}

func TestLineMustParseWithValue(t *testing.T) {
	line := application_fee.NewLine()
	err := line.Parse("4     RONALDO DE CARVALHO PEREZ                       145.396.428-20         11818506                   42.303,73          54.073,76      96.377,50            9.637,75           cálculo")
	require.NoError(t, err)
	require.Equal(t, uint16(4), line.Sequence())
	require.Equal(t, "14539642820", line.CPF())
	require.Equal(t, "RONALDO DE CARVALHO PEREZ", line.Name())
	require.Equal(t, "11818506", line.UniqueID())
	require.Equal(t, uint64(4230373), line.Main())
	require.Equal(t, uint64(5407376), line.Interest())
	require.Equal(t, uint64(9637750), line.Total())
	require.Equal(t, uint64(963775), line.Fees())
	require.Equal(t, "cálculo", line.Status())
}
