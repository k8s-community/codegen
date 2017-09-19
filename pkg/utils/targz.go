package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// tarGzArchivator creates *.tar.gz archives from files, dirs
type tarGzArchivator struct {
	SrcPath  string
	DestPath string

	tw *tar.Writer
}

// NewTarGzArchivator returns a new instance of TarGzArchivator
func newTarGzArchivator(srcPath, destPath string) *tarGzArchivator {
	return &tarGzArchivator{
		SrcPath:  srcPath,
		DestPath: destPath,
	}
}

// CreateTarGzArchive creates destPath.tar.gz archive
func CreateTarGzArchive(srcPath, destPath string) error {
	t := newTarGzArchivator(srcPath, destPath)
	return t.create()
}

// tarGzFile adds file in tar gz writer
func (t *tarGzArchivator) tarGzFile(relativePath string, fi os.FileInfo) error {
	fr, err := os.Open(t.SrcPath + "/" + relativePath)
	if err != nil {
		return fmt.Errorf("cannot open file: %s", err)
	}
	defer fr.Close()

	h := new(tar.Header)
	h.Name = relativePath
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = t.tw.WriteHeader(h)
	if err != nil {
		return fmt.Errorf("cannot write header to archive: %s", err)
	}

	_, err = io.Copy(t.tw, fr)
	if err != nil {
		return fmt.Errorf("cannot to copy file in archive: %s", err)
	}

	return nil
}

// tarGzDir adds dir in tar gz writer
func (t *tarGzArchivator) tarGzDir(relativePath string) error {
	directory, err := os.Open(t.SrcPath + "/" + relativePath)
	if err != nil {
		return fmt.Errorf("cannot open src dir: %s", err)
	}
	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return fmt.Errorf("cannot read src dir: %s", err)
	}

	for _, obj := range objects {
		srcObjectPath := relativePath + "/" + obj.Name()

		if obj.IsDir() {
			err = t.tarGzDir(srcObjectPath)
			if err != nil {
				return err
			}
		} else {
			err = t.tarGzFile(srcObjectPath, obj)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Create makes *.tar.gz
func (t *tarGzArchivator) create() error {
	info, err := os.Stat(t.SrcPath)
	if err != nil {
		return fmt.Errorf("cannot get info of src dir: %s", err)
	}

	fw, err := os.Create(t.DestPath)
	if err != nil {
		return fmt.Errorf("cannot create archive: %s", err)
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	t.tw = tar.NewWriter(gw)
	defer t.tw.Close()

	if info.IsDir() {
		return t.tarGzDir("")
	}

	// srcPath = dir of src file
	t.SrcPath = strings.TrimSuffix(t.SrcPath, "/"+info.Name())

	return t.tarGzFile(info.Name(), info)
}
