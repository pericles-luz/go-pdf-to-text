package extract_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-pdf-to-text/internal/extract"
	"github.com/stretchr/testify/require"
)

func readOriginFile(t *testing.T) []string {
	lines, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/009-11804009-C.txt")
	require.NoError(t, err)
	return lines
}

func TestReadLinesFromFile(t *testing.T) {
	lines, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/lines.txt")
	require.NoError(t, err)
	require.Equal(t, []string{"line 1", "line 2", "line 3"}, lines)
}

func TestReadLinesFromFileMustReturnErrFileNotFound(t *testing.T) {
	_, err := extract.ReadLinesFromFile(utils.GetBaseDirectory("pdf") + "/not-found.txt")
	require.EqualError(t, err, "open "+utils.GetBaseDirectory("pdf")+"/not-found.txt: no such file or directory")
}

func TestReadLinesFromFileMustRead424Lines(t *testing.T) {
	lines := readOriginFile(t)
	require.Len(t, lines, 64)
}
