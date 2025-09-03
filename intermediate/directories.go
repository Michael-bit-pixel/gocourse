package basics

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		//fmt.Println(err)
	}
}

func main() {

	// err := os.Mkdir("subdir",0750)
	// checkError(err)
	// checkError(os.Mkdir("subdir1", 0750))

	// defer os.RemoveAll("subdir1")

	// os.WriteFile("subdir1/file", []byte(""), 0750)

	// checkError(os.MkdirAll("subdir/parent/child", 0750))
	// checkError(os.MkdirAll("subdir/parent/child1", 0750))
	// checkError(os.MkdirAll("subdir/parent/child2", 0750))
	// checkError(os.MkdirAll("subdir/parent/child3", 0750))
	// os.WriteFile("subdir/parent/file", []byte(""), 0750)
	// os.WriteFile("subdir/parent/child/file", []byte(""), 0750)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.Type().IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))
	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	result1, err := os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result1 {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))
	dir, err = os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir

	// pathfile := "subdir"
	// fmt.Println("Walking Director")
	// err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
	// 	if err != nil {
	// 		fmt.Println("Error", err)
	// 		return err
	// 	}

	// 	fmt.Println(path)
	// 	return nil
	// })

	// checkError(err)
	// checkError(os.Chdir("gocourse"))
	// checkError(os.RemoveAll("./subdir"))
}
