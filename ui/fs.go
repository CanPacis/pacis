package ui

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func MergeFS(filesystems ...fs.FS) fs.FS {
	return MergedFS{filesystems: filesystems}
}

type MergedFS struct {
	filesystems []fs.FS
}

// Open opens the named file.
func (mfs MergedFS) Open(name string) (fs.File, error) {
	for _, fs := range mfs.filesystems {
		file, err := fs.Open(name)
		if err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

// ReadDir reads from the directory, and produces a DirEntry array of different
// directories.
//
// It iterates through all different filesystems that exist in the mfs MergeFS
// filesystem slice and it identifies overlapping directories that exist in different
// filesystems
func (mfs MergedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	dirsMap := make(map[string]fs.DirEntry)
	notExistCount := 0
	for _, filesystem := range mfs.filesystems {
		dir, err := fs.ReadDir(filesystem, name)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				notExistCount++
				continue
			}
			return nil, err
		}
		for _, v := range dir {
			if _, ok := dirsMap[v.Name()]; !ok {
				dirsMap[v.Name()] = v
			}
		}
		continue
	}
	if len(mfs.filesystems) == notExistCount {
		return nil, fs.ErrNotExist
	}
	dirs := make([]fs.DirEntry, 0, len(dirsMap))

	for _, value := range dirsMap {
		dirs = append(dirs, value)
	}

	return dirs, nil
}

type defaultFS struct {
	prefix string
	fs     fs.FS
}

func (fs defaultFS) Open(name string) (fs.File, error) {
	if fs.fs == nil {
		return os.Open(name)
	}
	return fs.fs.Open(name)
}

func MustSubFS(currentFs fs.FS, fsRoot string) fs.FS {
	subFs, err := subFS(currentFs, fsRoot)
	if err != nil {
		panic(fmt.Errorf("can not create sub FS, invalid root given, err: %w", err))
	}
	return subFs
}

func subFS(currentFs fs.FS, root string) (fs.FS, error) {
	root = filepath.ToSlash(filepath.Clean(root)) // note: fs.FS operates only with slashes. `ToSlash` is necessary for Windows
	if dFS, ok := currentFs.(*defaultFS); ok {
		// we need to make exception for `defaultFS` instances as it interprets root prefix differently from fs.FS.
		// fs.Fs.Open does not like relative paths ("./", "../") and absolute paths at all but prior echo.Filesystem we
		// were able to use paths like `./myfile.log`, `/etc/hosts` and these would work fine with `os.Open` but not with fs.Fs
		if !filepath.IsAbs(root) {
			root = filepath.Join(dFS.prefix, root)
		}
		return &defaultFS{
			prefix: root,
			fs:     os.DirFS(root),
		}, nil
	}
	return fs.Sub(currentFs, root)
}
