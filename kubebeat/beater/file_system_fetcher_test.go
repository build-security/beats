package beater

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileFetcherFetchASingleFile(t *testing.T) {
	directoryName := "test-outer-dir"
	files := []string{"file.txt"}
	dir := createDirectoriesWithFiles(t, "", directoryName, files)
	defer os.RemoveAll(dir)

	filePaths := []string{filepath.Join(dir, files[0])}
	fileFetcher := NewFileFetcher(filePaths)
	results, err := fileFetcher.Fetch()

	if err != nil {
		assert.Fail(t, "Fetcher was not able to fetch files from FS", err)
	}
	assert.Equal(t, len(results), 1)
	result := results[0].(FileSystemResourceData)

	assert.Equal(t, files[0], result.FileName)
	assert.Equal(t, "600", result.FileMode)
}

func TestFileFetcherFetchDirectoryOnly(t *testing.T) {
	directoryName := "test-outer-dir"
	files := []string{"file.txt"}
	dir := createDirectoriesWithFiles(t, "", directoryName, files)
	defer os.RemoveAll(dir)

	filePaths := []string{filepath.Join(dir)}
	fileFetcher := NewFileFetcher(filePaths)
	results, err := fileFetcher.Fetch()

	if err != nil {
		assert.Fail(t, "Fetcher was not able to fetch files from FS", err)
	}
	assert.Equal(t, len(results), 1)
	result := results[0].(FileSystemResourceData)

	expectedResult := filepath.Base(dir)
	assert.Equal(t, expectedResult, result.FileName)
}


func TestFileFetcherFetchOuterDirectoryOnly(t *testing.T) {
	outerDirectoryName := "test-outer-dir"
	outerFiles := []string{"output.txt"}
	outerDir := createDirectoriesWithFiles(t, "", outerDirectoryName, outerFiles)
	defer os.RemoveAll(outerDir)

	innerDirectoryName := "test-inner-dir"
	innerFiles := []string{"innerFolderFile.txt"}
	innerDir := createDirectoriesWithFiles(t, outerDir, innerDirectoryName, innerFiles)

	path := []string{outerDir + "/*"}
	fileFetcher := NewFileFetcher(path)
	results, err := fileFetcher.Fetch()

	if err != nil {
		assert.Fail(t, "Fetcher was not able to fetch files from FS", err)
	}

	assert.Equal(t, len(results), 2)

	//All inner files should exist in the final result
	expectedResult := []string{"output.txt", filepath.Base(innerDir)}
	for i := 0; i < len(results); i++ {
		fileSystemDataResources := results[i].(FileSystemResourceData)
		assert.Contains(t, expectedResult, fileSystemDataResources.FileName)
	}
}

func TestFileFetcherFetchDirectoryRecursively(t *testing.T) {
	outerDirectoryName := "test-outer-dir"
	outerFiles := []string{"output.txt"}
	outerDir := createDirectoriesWithFiles(t, "", outerDirectoryName, outerFiles)
	defer os.RemoveAll(outerDir)

	innerDirectoryName := "test-inner-dir"
	innerFiles := []string{"innerFolderFile.txt"}
	innerDir := createDirectoriesWithFiles(t, outerDir, innerDirectoryName, innerFiles)

	innerInnerDirectoryName := "test-inner-inner-dir"
	innerInnerFiles := []string{"innerInnerFolderFile.txt"}
	innerInnerDir := createDirectoriesWithFiles(t, innerDir, innerInnerDirectoryName, innerInnerFiles)

	path := []string{outerDir + "/**"}
	fileFetcher := NewFileFetcher(path)
	results, err := fileFetcher.Fetch()

	if err != nil {
		assert.Fail(t, "Fetcher was not able to fetch files from FS", err)
	}

	assert.Equal(t, len(results), 6)

	directories := []string{filepath.Base(outerDir), filepath.Base(innerDir), filepath.Base(innerInnerDir)}
	allFilesName := append(append(append(innerFiles, directories...), outerFiles...),innerInnerFiles...)

	//All inner files should exist in the final result
	for i := 0; i < len(results); i++ {

		fileSystemDataResources := results[i].(FileSystemResourceData)
		assert.Contains(t, allFilesName, fileSystemDataResources.FileName)
	}
}

// This function creates a new directory with files inside and returns the path of the new directory
func createDirectoriesWithFiles(t *testing.T, dirPath string, dirName string, filesToWriteInDirectory []string) string {
	dirPath, err := ioutil.TempDir(dirPath, dirName)
	if err != nil {
		t.Fatal(err)
	}
	for _, fileName := range filesToWriteInDirectory {
		file := filepath.Join(dirPath, fileName)
		if err := ioutil.WriteFile(file, []byte("test txt\n"), 0600); err != nil {
			assert.Fail(t, "Could not able to write a new file", err)
		}
	}
	return dirPath
}
