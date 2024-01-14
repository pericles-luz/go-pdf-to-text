package service_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/service_fee"
	"github.com/stretchr/testify/require"
)

func TestSummaryMustParse(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/002-Honorários.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 1)
}

func TestSummaryMustNotParse(t *testing.T) {
	summary := service_fee.NewSummary()
	err := summary.Parse(utils.GetBaseDirectory("pdf") + "/009-11804009-C.txt")
	require.NoError(t, err)
	require.Len(t, summary.Summaries(), 0)
}
