package application_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/stretchr/testify/require"
)

func TestLinhaMustValidate(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.NoError(t, linha.Validate())
}

func TestLinhaMustReturnErrMesAnoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrMesAnoInvalido.Error())
}

func TestLinhaMustReturnErrVencimentoBasicoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrVencimentoBasicoInvalido.Error())
}

// func TestLinhaMustReturnErrSomaInvalida(t *testing.T) {
// 	linha := application.NewLinha()
// 	linha.SetMesAno("fev/97")
// 	linha.SetVencimentoBasico("138,61")
// 	linha.SetTercoFerias("46,20")
// 	linha.SetPercentual("28,86%")
// 	linha.SetValorDevido("53,34")
// 	linha.SetIndiceCorrecao("2,56866867")
// 	linha.SetValorCorrigido("137,01")
// 	linha.SetJurosMora("131,8171%")
// 	linha.SetValorJurosMora("180,60")
// 	linha.SetTotalDevido("317,60")
// 	require.EqualError(t, linha.Validate(), application.ErrSomaInvalida.Error())
// }

func TestLinhaMustReturnErrPercentualInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrPercentualInvalido.Error())
}

func TestLinhaMustReturnErrValorDevidoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrValorDevidoInvalido.Error())
}

func TestLinhaMustReturnErrIndiceCorrecaoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrIndiceCorrecaoInvalido.Error())
}

func TestLinhaMustReturnErrValorCorrigidoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrValorCorrigidoInvalido.Error())
}

func TestLinhaMustReturnErrJurosMoraInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetValorJurosMora("180,60")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrJurosMoraInvalido.Error())
}

func TestLinhaMustReturnErrValorJurosMoraInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetTotalDevido("317,60")
	require.EqualError(t, linha.Validate(), application.ErrValorJurosMoraInvalido.Error())
}

func TestLinhaMustReturnErrTotalDevidoInvalido(t *testing.T) {
	linha := application.NewLinha()
	linha.SetMesAno("fev/97")
	linha.SetVencimentoBasico("138,61")
	linha.SetTercoFerias("46,20")
	linha.SetSoma("184,81")
	linha.SetPercentual("28,86%")
	linha.SetValorDevido("53,34")
	linha.SetIndiceCorrecao("2,56866867")
	linha.SetValorCorrigido("137,01")
	linha.SetJurosMora("131,8171%")
	linha.SetValorJurosMora("180,60")
	require.EqualError(t, linha.Validate(), application.ErrTotalDevidoInvalido.Error())
}
