package service_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service"
	"github.com/stretchr/testify/require"
)

func TestCalculoMustParse(t *testing.T) {
	calculo := service.NewCalculo()
	err := calculo.Parse(utils.GetBaseDirectory("pdf") + "/009-11804009-C.txt")
	require.NoError(t, err)
}

func TestCalculoWithAdvantagesMustParse(t *testing.T) {
	calculo := service.NewCalculo()
	err := calculo.Parse(utils.GetBaseDirectory("pdf") + "/004-01139401-C.txt")
	require.NoError(t, err)
}
