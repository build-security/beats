package beater

import (
	"github.com/elastic/beats/v7/libbeat/logp"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// FileSystemFetcher implement the Fetcher interface
// The FileSystemFetcher meant to fetch file/directories from the file system and ship it
// to the Kubebeat
type FileSystemFetcher struct {
	inputFilePatterns []string // Files and directories paths for the fetcher to extract info from
}

const (
	FileSystemType = "file-system"
)

// FileSystemResource represents a struct for a system resource data
// This struct is being used by the fileSystemFetcher when
type FileSystemResource struct {
	ID       string `json:"id"`
	FileName string `json:"filename"`
	FileMode string `json:"mode"`
	Gid      string `json:"gid"`
	Uid      string `json:"uid"`
	Path     string `json:"path"`
}

func NewFileFetcher(filesPaths []string) Fetcher {
	return &FileSystemFetcher{
		inputFilePatterns: filesPaths,
	}
}

func (f *FileSystemFetcher) Fetch() ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	// Input files might contain glob pattern
	for _, filePattern := range f.inputFilePatterns {
		matchedFiles, err := Glob(filePattern)
		if err != nil {
			logp.Err("Failed to find matched glob for %s, error - %+v", filePattern, err)
		}
		for _, file := range matchedFiles {
			resourceInfo := f.fetchSystemResource(file)
			results = append(results, FetcherResult{FileSystemType, resourceInfo})
		}
	}
	return results, nil
}

func (f *FileSystemFetcher) fetchSystemResource(filePath string) ResourceInfo {

	info, err := os.Stat(filePath)
	if err != nil {
		logp.Err("Failed to fetch %s, error - %+v", filePath, err)
		return ResourceInfo{}
	}
	resourceInfo := FromFileInfo(info, filePath)

	return resourceInfo
}

func FromFileInfo(info os.FileInfo, path string) ResourceInfo {

	if info == nil {
		return ResourceInfo{}
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		logp.Err("Not a syscall.Stat_t")
		return ResourceInfo{}
	}

	uid := stat.Uid
	gid := stat.Gid
	u := strconv.FormatUint(uint64(uid), 10)
	g := strconv.FormatUint(uint64(gid), 10)
	usr, _ := user.LookupId(u)
	group, _ := user.LookupGroupId(g)
	mod := strconv.FormatUint(uint64(info.Mode().Perm()), 8)
	id := strconv.FormatUint(uint64(stat.Ino), 10)

	data := FileSystemResource{
		FileName: info.Name(),
		FileMode: mod,
		Uid:      usr.Name,
		Gid:      group.Name,
		Path:     path,
	}

	return ResourceInfo{id, data}
}

func (f *FileSystemFetcher) Stop() {
}
