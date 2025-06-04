package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Zip(src, dst string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	wr := zip.NewWriter(dstFile)
	defer wr.Close()
	return zipFile(src, "", wr)
}

func zipFile(src, relativePath string, zipw *zip.Writer) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if relativePath != "" {
		relativePath = filepath.Join(relativePath, info.Name())
	} else {
		relativePath = info.Name()
	}

	if info.IsDir() {
		if _, err := zipw.Create(relativePath + "/"); err != nil {
			return err
		}

		files, err := os.ReadDir(src)
		if err != nil {
			return err
		}
		for _, file := range files {
			srcPath := filepath.Join(src, file.Name())
			if err := zipFile(srcPath, relativePath, zipw); err != nil {
				return err
			}
		}
		return nil
	}

	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	w, err := zipw.Create(relativePath)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, file); err != nil {
		return err
	}
	return nil
}

func Unzip(src, dst string) error {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	for _, file := range zipReader.File {
		if err := unzipFile(dst, file); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(dst string, zipFile *zip.File) error {
	dstPath := filepath.Join(dst, zipFile.Name)
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}
	if zipFile.FileInfo().IsDir() {
		if err := os.Mkdir(dstPath, 0755); err != nil {
			return err
		}
		return nil
	}
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	srcFile, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return nil
}

func test1() {
	// 密码
	err := Zip("/data/wangdaoan/ugreen/zero/dir01", "/data/wangdaoan/ugreen/zero/dir01.zip")
	if err != nil {
		log.Fatal(err)
	}
	err = Unzip("/data/wangdaoan/ugreen/zero/dir01.zip", "/data/wangdaoan/ugreen/zero/test03")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test1()
}
