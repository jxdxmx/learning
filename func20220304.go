package main

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"log"
)

func f20220304() {
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: "192.168.47.66:50070", User: "hadoop"})
	if err != nil {
		log.Fatal(err)
	}
	checksum, err := fs.GetFileChecksum(gowfs.Path{Name: "/user1/name1.txt"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checksum)

	createTestDir(fs, "/user1/name_test1.txt")
}

func createTestDir(fs *gowfs.FileSystem, hdfsPath string) {
	path := gowfs.Path{Name: hdfsPath}
	ok, err := fs.MkDirs(path, 0744)
	if err != nil || !ok {
		log.Fatal("Unable to create test directory ", hdfsPath, ":", err)
	}
	log.Println("HDFS Path ", path.Name, " created.")
}

//type Student struct {
//	name string
//	age  int
//}
//
//func f20220304() {
//	var student = new(Student)
//	student.name = "xu"
//	fmt.Println(student)
//}
