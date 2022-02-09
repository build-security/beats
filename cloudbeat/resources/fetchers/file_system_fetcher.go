package fetchers

import (
	"context"
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/elastic/beats/v7/libbeat/logp"
)

// FileSystemFetcher implement the Fetcher interface
// The FileSystemFetcher meant to fetch file/directories from the file system and ship it
// to the Cloudbeat
type FileSystemFetcher struct {
	cfg FileFetcherConfig
}

type FileFetcherConfig struct {
	BaseFetcherConfig
	Patterns []string `config:"patterns"` // Files and directories paths for the fetcher to extract info from
}

const (
	FileSystemType = "file-system"
)

func NewFileFetcher(cfg FileFetcherConfig) Fetcher {
	return &FileSystemFetcher{
		cfg: cfg,
	}
}

func (f *FileSystemFetcher) Fetch(ctx context.Context) ([]PolicyResource, error) {
	results := make([]PolicyResource, 0)

	// Input files might contain glob pattern
	for _, filePattern := range f.cfg.Patterns {
		matchedFiles, err := Glob(filePattern)
		if err != nil {
			logp.Err("Failed to find matched glob for %s, error - %+v", filePattern, err)
		}
		for _, file := range matchedFiles {
			resource := f.fetchSystemResource(file)
			results = append(results, resource)
		}
	}
	return results, nil
}

func (f *FileSystemFetcher) fetchSystemResource(filePath string) FileSystemResource {

	info, err := os.Stat(filePath)
	if err != nil {
		logp.Err("Failed to fetch %s, error - %+v", filePath, err)
		return FileSystemResource{}
	}
	resourceInfo := FromFileInfo(info, filePath)

	return resourceInfo
}

func FromFileInfo(info os.FileInfo, path string) FileSystemResource {

	if info == nil {
		return FileSystemResource{}
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		logp.Err("Not a syscall.Stat_t")
		return FileSystemResource{}
	}

	uid := stat.Uid
	gid := stat.Gid
	u := strconv.FormatUint(uint64(uid), 10)
	g := strconv.FormatUint(uint64(gid), 10)
	usr, _ := user.LookupId(u)
	group, _ := user.LookupGroupId(g)
	mod := strconv.FormatUint(uint64(info.Mode().Perm()), 8)
	inode := strconv.FormatUint(uint64(stat.Ino), 10)

	data := FileSystemResource{
		FileName: info.Name(),
		FileMode: mod,
		Uid:      usr.Name,
		Gid:      group.Name,
		Path:     path,
		Inode:    inode,
	}

	return data
}

func (f *FileSystemFetcher) Stop() {
}

func (res FileSystemResource) GetID() string {
	return res.Inode
}
