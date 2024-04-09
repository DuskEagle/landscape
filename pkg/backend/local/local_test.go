package local

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLocalBackend(t *testing.T) {
	dir := t.TempDir()
	_, err := NewLocalBackend(dir)
	require.ErrorContains(t, err, "was a directory, expected a file")

	_, err = NewLocalBackend(dir + "somesuffix/file.txt")
	require.Error(t, err, "calling on a directory that doesn't exist should error")

	file := dir + "/file.txt"
	_, err = NewLocalBackend(file)
	require.NoError(t, err, "calling NewLocalBackend on non-existent file should work")
	_, err = os.Stat(file)
	require.NoError(t, err, "file.txt should now exist")

	_, err = NewLocalBackend(file)
	require.NoError(t, err, "calling NewLocalBackend on existing file should work")
}
