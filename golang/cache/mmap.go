package main

import (
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func do1() {
	n := 1000
	t := int(unsafe.Sizeof(0)) * n

	file, err := os.Create("a.txt")
	if err != nil {
		log.Println(err)
		return
	}

	offset, err := file.Seek(int64(t-1), os.SEEK_SET)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("offset:", offset)

	nbytes, err := file.Write([]byte(" "))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("nbytes:", nbytes)

	mdata, err := syscall.Mmap(int(file.Fd()), 0, int(t), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		return
	}

	err = file.Close()
	if err != nil {
		log.Println(err)
		return
	}

	pdata := (*[1000]int)(unsafe.Pointer(&mdata))
	for i := 0; i < n; i++ {
		pdata[i] = i * i
	}

	//log.Println(*pdata)

	err = syscall.Munmap(mdata)
	if err != nil {
		log.Println(err)
		return
	}
}

func do2() {
	// create file and write data and close
	file, err := os.Create("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	nbytes, err := file.Write([]byte("123456789"))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("nbytes:", nbytes)
	if err := file.Close(); err != nil {
		log.Println(err)
		return
	}

	// open file
	file, err = os.OpenFile("a.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	content := make([]byte, 128)
	nbytes, err = file.Read(content)
	if err != nil {
		log.Println(err)
		return
	}
	content = content[:nbytes]
	log.Println("content:", string(content))
	fileinfo, _ := file.Stat()
	log.Println("file size:", fileinfo.Size())

	mdata, err := syscall.Mmap(int(file.Fd()), 0, int(fileinfo.Size())+1, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("mdata:", mdata)
	log.Println("len mdata:", len(mdata))

	mdata[4] = 'a'

	if err := syscall.Munmap(mdata); err != nil {
		log.Println(err)
		return
	}

	offset, err := file.Seek(0, os.SEEK_SET)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("offset:", offset)

	content = make([]byte, 128)
	nbytes, err = file.Read(content)
	if err != nil {
		log.Println(err)
		return
	}
	content = content[:nbytes]
	log.Println("content:", string(content))

	if err := file.Close(); err != nil {
		log.Println(err)
		return
	}
	os.Remove("a.txt")
}

func do3() {
	file, err := os.Create("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	defer os.Remove("a.txt")

	mdata, err := syscall.Mmap(int(file.Fd()), 0, 4096, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		return
	}

	//log.Println("mdata:", mdata)
	log.Println("len mdata:", len(mdata))

	//空文件没有内存页,访问会报错
	//mdata[0] = 'a'

	if err := syscall.Munmap(mdata); err != nil {
		log.Println(err)
		return
	}

	file.Seek(0, os.SEEK_SET)
	content := make([]byte, 10)
	nbytes, err := file.Read(content)
	if err != nil {
		log.Println(err)
		return
	}
	content = content[:nbytes]
	log.Println("content:", string(content))
}

func do4() {
	file, err := os.Create("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	file.Write([]byte("1234567890"))
	fileinfo, _ := file.Stat()
	log.Println("file size:", fileinfo.Size())

	mdata, err := syscall.Mmap(int(file.Fd()), 0, 4096, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	//一页4096,超出一页报错
	//mdata, err := syscall.Mmap(int(file.Fd()), 0, 4096+1, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		return
	}

	//log.Println("mdata:", mdata)
	log.Println("len mdata:", len(mdata))

	mdata[4] = 'a'
	mdata[9] = 'b'
	mdata[len(mdata)-1] = 'c'

	if err := syscall.Munmap(mdata); err != nil {
		log.Println(err)
		return
	}

	file.Seek(0, os.SEEK_SET)
	content := make([]byte, 10)
	nbytes, err := file.Read(content)
	if err != nil {
		log.Println(err)
		return
	}
	content = content[:nbytes]
	log.Println("content:", string(content))

	if err := file.Close(); err != nil {
		log.Println(err)
		return
	}
	os.Remove("a.txt")
}

func do5() {
	file, err := os.Create("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	file.Write([]byte("1234567890"))

	mdata, err := syscall.Mmap(int(file.Fd()), 0, 4096, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Println(err)
		return
	}
	// 内存页直接映射磁盘,即使文件关闭也能读写
	file.Close()

	log.Println("len mdata:", len(mdata))
	mdata[0] = 'a'

	if err := syscall.Munmap(mdata); err != nil {
		log.Println(err)
		return
	}

	content, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("content:", string(content))
	os.Remove("a.txt")
}

func do6() {
	mdata, err := syscall.Mmap(-1, 0, 16, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("len mdata:", len(mdata))
	mdata[0] = 'a'
	log.Println("mdata:", mdata)

	if err := syscall.Munmap(mdata); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do6()
}
