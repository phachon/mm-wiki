package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var Zipx = NewZipx()

type zipx struct {
	lock sync.Mutex
}

func NewZipx() *zipx {
	return &zipx{}
}

type CompressFileInfo struct {
	File       string
	PrefixPath string
	osFile     *os.File
}

func (z *zipx) PackFile(files []*CompressFileInfo, dest string) error {

	z.lock.Lock()
	defer z.lock.Unlock()

	// create dest file
	destDir := filepath.Dir(dest)
	err := os.RemoveAll(destDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(destDir, 0777)
	if err != nil {
		return err
	}
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()

	// file handle
	for _, f := range files {
		f3, err := os.Open(f.File)
		if err != nil {
			continue
		}
		f.osFile = f3
	}
	defer func() {
		for _, f := range files {
			if f.osFile == nil {
				continue
			}
			_ = f.osFile.Close()
		}
	}()

	// zip writer
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := z.compress(file.osFile, file.PrefixPath, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *zipx) Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := z.compress(file, "images/1/4", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *zipx) compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = z.compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func (z *zipx) DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(z.getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func (z *zipx) getDir(path string) string {
	return z.subString(path, 0, strings.LastIndex(path, "/"))
}

func (z *zipx) subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// srcFile could be a single file or a directory
func (z *zipx) Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	_ = filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
