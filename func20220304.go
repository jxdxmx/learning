package main

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"log"
)

// 问题1：SafeModeException  解决办法：# ../../bin/hadoop dfsadmin -safemode leave
//root@k8s66mgr:~/learning/git-learning# go run .
//{MD5-of-0MD5-of-512CRC32C 0000020000000000000000004b8b791405539be4efd80842cc8e4b0a00000000 28}
//2022/03/05 13:24:58 Unable to create test directory /user1/name_test1.txt:RemoteException: SafeModeException [org.apache.hadoop.hdfs.server.namenode.SafeModeException]
//[Cannot create directory /user1/name_test1.txt. Name node is in safe mode.
//The reported blocks 40 has reached the threshold 0.9990 of total blocks 40. The number of live datanodes 1 has reached the minimum number 0. In safe mode extension. Safe mode will be turned off automatically in 1 seconds.]
//exit status 1
//sh-4.1# ../../bin/hadoop dfsadmin -safemode leave
//DEPRECATED: Use of this script to execute hdfs command is deprecated.
//Instead use the hdfs command for it.
//
//Safe mode is OFF

//问题2：AccessControlException 解决办法，修改权限 sh-4.1# ./hadoop fs -chmod -R 777 /user1
//root@k8s66mgr:~/learning/git-learning# go run .
//{MD5-of-0MD5-of-512CRC32C 0000020000000000000000004b8b791405539be4efd80842cc8e4b0a00000000 28}
//2022/03/05 13:29:29 Unable to create test directory /user1/name_test1.txt:RemoteException: AccessControlException [org.apache.hadoop.security.AccessControlException]
//[Permission denied: user=hadoop, access=WRITE, inode="/user1/name_test1.txt":root:supergroup:drwxr-xr-x]
//exit status 1
//sh-4.1# ./hadoop fs -ls /user1
//Found 2 items
//-rw-r--r--   1 root supergroup         11 2022-03-04 00:01 /user1/name1.txt
//-rw-r--r--   1 root supergroup          0 2022-03-03 22:51 /user1/name2.txt
//sh-4.1# ./hadoop fs -chmod -R 777 /user1
//sh-4.1# ./hadoop fs -ls /user1
//Found 2 items
//-rwxrwxrwx   1 root supergroup         11 2022-03-04 00:01 /user1/name1.txt
//-rwxrwxrwx   1 root supergroup          0 2022-03-03 22:51 /user1/name2.txt
//sh-4.1#

// 终于成功！！！【代码要运行在192.168.47.66_m上】
//root@k8s66mgr:~/learning/git-learning# go run .
//{MD5-of-0MD5-of-512CRC32C 0000020000000000000000004b8b791405539be4efd80842cc8e4b0a00000000 28}
//2022/03/05 13:33:41 HDFS Path  /user1/name_test1.txt  created.
//root@k8s66mgr:~/learning/git-learning#

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
