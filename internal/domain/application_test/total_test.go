package application_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application"
	"github.com/stretchr/testify/require"
)

func TestTotalMustValidate(t *testing.T) {
	total := application.NewTotal()
	total.SetValorCorrigido("R$ 1.000,00")
	total.SetValorJurosMora("R$ 1.000,00")
	total.SetTotalDevido("R$ 2.000,00")
	require.NoError(t, total.Validate())
}

func TestTotalMustReturnErrValorCorrigidoInvalido(t *testing.T) {
	total := application.NewTotal()
	total.SetValorJurosMora("R$ 1.000,00")
	total.SetTotalDevido("R$ 2.000,00")
	require.EqualError(t, total.Validate(), application.ErrValorCorrigidoInvalido.Error())
}

func TestTotalMustReturnErrValorJurosMoraInvalido(t *testing.T) {
	total := application.NewTotal()
	total.SetValorCorrigido("R$ 1.000,00")
	total.SetTotalDevido("R$ 2.000,00")
	require.EqualError(t, total.Validate(), application.ErrValorJurosMoraInvalido.Error())
}

func TestTotalMustReturnErrTotalDevidoInvalido(t *testing.T) {
	total := application.NewTotal()
	total.SetValorCorrigido("R$ 1.000,00")
	total.SetValorJurosMora("R$ 1.000,00")
	require.EqualError(t, total.Validate(), application.ErrTotalDevidoInvalido.Error())
}
