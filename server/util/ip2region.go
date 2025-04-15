package util

import (
	"fmt"
	"os"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var cBuff []byte

// LoadContent load the whole xdb content from the specified file handle
func LoadContent(handle *os.File) ([]byte, error) {
	// get file size
	fi, err := handle.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	size := fi.Size()

	// seek to the head of the file
	_, err = handle.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("seek to get xdb file length: %w", err)
	}

	var buff = make([]byte, size)
	rLen, err := handle.Read(buff)
	if err != nil {
		return nil, err
	}

	if rLen != len(buff) {
		return nil, fmt.Errorf("incomplete read: readed bytes should be %d", len(buff))
	}

	return buff, nil
}

// LoadContentFromFile load the whole xdb content from the specified db file path
func LoadContentFromFile(dbFile string) ([]byte, error) {
	handle, err := os.OpenFile(dbFile, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("open xdb file `%s`: %w", dbFile, err)
	}

	defer func() {
		_ = handle.Close()
	}()

	cBuff, err := LoadContent(handle)
	if err != nil {
		return nil, err
	}

	return cBuff, nil
}

func InitIp2Region() (err error) {
	//从 dbPath 加载整个 xdb 到内存
	dbPath := "./ip2region.xdb"
	cBuff, err = LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}
	return
}

func Ip2Region(list_ip []string) ([]string, error) {
	// 用全局的 cBuff 创建完全基于内存的查询对象。
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		logrus.Errorf("failed to create searcher with content: %s\n", err)
		return nil, err
	}
	defer searcher.Close()

	list_region := make([]string, len(list_ip))
	for i, ip := range list_ip {
		list_region[i], err = searcher.SearchByStr(ip)
		if err != nil {
			logrus.Errorf("failed to SearchIP(%s): %s\n", ip, err)
			list_region[i] = "未知归属地"
		}
	}
	return list_region, nil
}
