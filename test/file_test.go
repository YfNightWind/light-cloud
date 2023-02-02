package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片的大小
const chunkSize = 100 * 1024 * 1024 // 100MB 测试用
// const chunkSize = 1024 * 1024 // 1MB

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	// 数据应向上取整，否则会出现缺失分片
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	// 声明一个数组，放置分片
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0) // 从0开始
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)

		// 创建分片文件
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		// 读取chunk文件
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}

		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	// 获取第一个文件的信息
	file1, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	// 获取第二个文件的信息
	file2, err := os.OpenFile("test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
