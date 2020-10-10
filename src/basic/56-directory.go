package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"learn-golang/src/utils"
)

func main() {
	tmpdir := "tmpdir"
	err := os.Mkdir(tmpdir, 0755)
	utils.Check(err)

	defer os.RemoveAll(tmpdir)

	createEmptyFile := func(name string) {
		data := []byte("")
		utils.Check(ioutil.WriteFile(name, data, 0644))
	}

	createEmptyFile(tmpdir + "/file1")
	err = os.MkdirAll(tmpdir+"/parent/child", 0755)
	utils.Check(err)

	createEmptyFile(tmpdir + "/parent/file2")
	createEmptyFile(tmpdir + "/parent/file3")
	createEmptyFile(tmpdir + "/parent/child/file4")

	dir, err := ioutil.ReadDir(tmpdir + "/parent")
	utils.Check(err)

	fmt.Println("Listing tmpdir/parent")
	for _, entry := range dir {
		fmt.Println(" |-", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(tmpdir + "/parent/child")
	utils.Check(err)

	dir, err = ioutil.ReadDir(".")
	utils.Check(err)

	fmt.Println("Listing tmpdir/parent/child")
	for _, entry := range dir {
		fmt.Println(" |-", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	utils.Check(err)

	fmt.Println("Visiting tmpdir")
	err = filepath.Walk("tmpdir", visit)

}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" |-", path, info.IsDir())
	return nil
}
