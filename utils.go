package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 工具函数文件
func uintToByte(num uint64) []byte {
	var buffer bytes.Buffer

	// 将数据以二进制方式保存到buffer中
	// func Write(w io.Writer, order ByteOrder, data interface{}) error
	err := binary.Write(&buffer, binary.BigEndian /* 大端对齐 */, num)

	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
