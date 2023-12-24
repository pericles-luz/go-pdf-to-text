package application_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/stretchr/testify/require"
)

func linha(t *testing.T) *application.Linha {
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
	return linha
}

func total(t *testing.T) *application.Total {
	total := application.NewTotal()
	total.SetValorCorrigido("137,01")
	total.SetValorJurosMora("180,60")
	total.SetTotalDevido("317,61")
	require.NoError(t, total.Validate())
	return total
}

func desagio35(t *testing.T) *application.Total {
	desagio35 := application.NewTotal()
	desagio35.SetValorCorrigido("47,95")
	desagio35.SetValorJurosMora("63,21")
	desagio35.SetTotalDevido("111,16")
	require.NoError(t, desagio35.Validate())
	return desagio35
}

func totalAposDesagio35(t *testing.T) *application.Total {
	totalAposDesagio35 := application.NewTotal()
	totalAposDesagio35.SetValorCorrigido("89,06")
	totalAposDesagio35.SetValorJurosMora("117,39")
	totalAposDesagio35.SetTotalDevido("206,45")
	require.NoError(t, totalAposDesagio35.Validate())
	return totalAposDesagio35
}

func TestCalculoMustValidate(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.NoError(t, calculo.Validate())
}

func TestCalculoMustReturnErrProcessoNumeroInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrProcessoNumeroInvalido.Error())
}

func TestCalculoMustReturnErrProcessoPrincipalInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrProcessoPrincipalInvalido.Error())
}

func TestCalculoMustReturnErrAjuizamentoInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrAjuizamentoInvalido.Error())
}

func TestCalculoMustReturnErrCitacaoInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrCitacaoInvalido.Error())
}

func TestCalculoMustReturnErrDataCalculoInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrDataCalculoInvalido.Error())
}

func TestCalculoMustReturnErrExequenteInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrExequenteInvalido.Error())
}

func TestCalculoMustReturnErrCpfInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrCpfInvalido.Error())
}

func TestCalculoMustReturnErrIdUnicaInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrIdUnicaInvalido.Error())
}

func TestCalculoMustReturnErrTabelaInvalida(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrTabelaInvalida.Error())
}

func TestCalculoMustReturnErrTotalInvalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal(" 0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("034.107.027-09")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetDesagio35(desagio35(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrTotalInvalido.Error())
}

func TestCalculoMustReturnErrDesagio35Invalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal("0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("03410702709")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetTotalAposDesagio35(totalAposDesagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrDesagio35Invalido.Error())
}

func TestCalculoMustReturnErrTotalAposDesagio35Invalido(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	calculo.SetProcessoPrincipal("0006379-33.1997.4.05.8100")
	calculo.SetAjuizamento("01/01/2000")
	calculo.SetCitacao("01/01/2000")
	calculo.SetCalculo("01/01/2000")
	calculo.SetExequente("EXEQUENTE")
	calculo.SetCpf("03410702709")
	calculo.SetIdUnica("123456789")
	calculo.AddLinha(linha(t))
	calculo.SetTotal(total(t))
	calculo.SetDesagio35(desagio35(t))
	require.EqualError(t, calculo.Validate(), application.ErrTotalAposDesagio35Invalido.Error())
}

func TestCalculoMustReturnErrLinhaJaExistente(t *testing.T) {
	calculo := application.NewCalculo()
	calculo.SetProcessoNumero("0001770-45.2013.4.05.8100")
	linha := linha(t)
	calculo.AddLinha(linha)
	require.EqualError(t, calculo.AddLinha(linha), application.ErrLinhaJaExistente.Error())
}
