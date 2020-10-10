package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"learn-golang/src/utils"
)

func main() {
	file, err := ioutil.TempFile("", "sample")
	utils.Check(err)
	fmt.Println("Temp file name: ", file.Name())

	defer os.Remove(file.Name())

	_, err = file.Write([]byte{1, 2, 3, 4})
	utils.Check(err)

	dir, err := ioutil.TempDir("", "sampledir")
	utils.Check(err)
	fmt.Println("Temp dir name: ", dir)

	defer os.RemoveAll(dir)

	fname := filepath.Join(dir, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	utils.Check(err)
}

// Temp file name:  /var/folders/t7/zh_h0h0d5vn0c2xsr_1tfqnh0000gn/T/sample018699025
// Temp dir name:  /var/folders/t7/zh_h0h0d5vn0c2xsr_1tfqnh0000gn/T/sampledir214998844
