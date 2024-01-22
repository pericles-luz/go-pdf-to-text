package application_fee_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/domain/application_fee"
	"github.com/stretchr/testify/require"
)

func TestMustHavensMustParseFromJsonFile(t *testing.T) {
	list := application_fee.NewMustHaveList()
	err := list.FromJSONFile(utils.GetBaseDirectory("assets") + "/mustHave.json")
	require.NoError(t, err)
}
