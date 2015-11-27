package copyutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src string, dst string) (err error) {
	sfh, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sfh.Close()
	dfh, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dfh.Close()
	_, err = io.Copy(dfh, sfh)
	if err == nil {
		srcinfo, err := os.Stat(src)
		if err != nil {
			err = os.Chmod(dst, srcinfo.Mode())
		}
	}
	return
}

func CopyDir(src string, dst string) (err error) {
	// get properties of src dir
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	// create dst dir
	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return err
	}
	directory, _ := os.Open(src)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sfp := filepath.Join(src, obj.Name())
		dfp := filepath.Join(dst, obj.Name())
		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}
