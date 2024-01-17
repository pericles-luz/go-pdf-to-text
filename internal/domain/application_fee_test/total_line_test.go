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
	line.SetTotal(1100)
	line.SetFees(110)
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

func TestTotalLineMustAddLine(t *testing.T) {
	line := application_fee.NewTotalLine()
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	totalLine := application_fee.NewTotalLine()
	totalLine.Add(line)
	require.Equal(t, uint64(1000), totalLine.Main())
	require.Equal(t, uint64(100), totalLine.Interest())
	require.Equal(t, uint64(900), totalLine.Total())
	require.Equal(t, uint64(100), totalLine.Fees())
}

func TestTotalLineMustSubtractLine(t *testing.T) {
	line := application_fee.NewTotalLine()
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	totalLine := application_fee.NewTotalLine()
	totalLine.SetMain(1000)
	totalLine.SetInterest(100)
	totalLine.SetTotal(900)
	totalLine.SetFees(100)
	totalLine.Subtract(line)
	require.Equal(t, uint64(0), totalLine.Main())
	require.Equal(t, uint64(0), totalLine.Interest())
	require.Equal(t, uint64(0), totalLine.Total())
	require.Equal(t, uint64(0), totalLine.Fees())
}

func TestTotalLineMustAddTotalLine(t *testing.T) {
	line := application_fee.NewTotalLine()
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	totalLine := application_fee.NewTotalLine()
	totalLine.SetMain(1000)
	totalLine.SetInterest(100)
	totalLine.SetTotal(900)
	totalLine.SetFees(100)
	totalLine.Add(line)
	require.Equal(t, uint64(2000), totalLine.Main())
	require.Equal(t, uint64(200), totalLine.Interest())
	require.Equal(t, uint64(1800), totalLine.Total())
	require.Equal(t, uint64(200), totalLine.Fees())
}

func TestTotalLineMustSubtractTotalLine(t *testing.T) {
	line := application_fee.NewTotalLine()
	line.SetMain(1000)
	line.SetInterest(100)
	line.SetTotal(900)
	line.SetFees(100)
	totalLine := application_fee.NewTotalLine()
	totalLine.SetMain(1000)
	totalLine.SetInterest(100)
	totalLine.SetTotal(900)
	totalLine.SetFees(100)
	totalLine.Subtract(line)
	require.Equal(t, uint64(0), totalLine.Main())
	require.Equal(t, uint64(0), totalLine.Interest())
	require.Equal(t, uint64(0), totalLine.Total())
	require.Equal(t, uint64(0), totalLine.Fees())
}
