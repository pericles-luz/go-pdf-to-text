package application_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/stretchr/testify/require"
)

func TestTotalLineMustValidate(t *testing.T) {
	line := application_fee.NewTotalLine()
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	err := line.Validate()
	require.NoError(t, err)
}

func TestTotalLineMustParse(t *testing.T) {
	line := application_fee.NewTotalLine()
	err := line.Parse("TOTAL GERAL                                                                                      1.903.713,12       2.465.680,23         4.369.393,36          436.939,34")
	require.NoError(t, err)
	require.Equal(t, uint64(190371312), line.Main())
	require.Equal(t, uint64(246568023), line.Interest())
	require.Equal(t, uint64(436939336), line.Total())
	require.Equal(t, uint64(43693934), line.Fees())
}
