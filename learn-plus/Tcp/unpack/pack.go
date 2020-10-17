package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_hander ="12345678"

// 编码
func Encode(bytesByffer io.Writer, content string)  error {
	// msg_handle + content_len + content
	// 8 + 4 + len
	if err := binary.Write(bytesByffer,binary.BigEndian,[]byte(Msg_hander)); err != nil {  // 使用二进制编码
		return err
	}

	clen := int32(len([]byte(content)))  // 转化为32位，4字节的长度
	if err := binary.Write(bytesByffer,binary.BigEndian,clen); err != nil {
		return err
	}

	if err := binary.Write(bytesByffer,binary.BigEndian,[]byte(content)); err != nil {
		return err
	}

	return nil
}

// 解码
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte,len(Msg_hander))
	if _, err  = io.ReadFull(bytesBuffer,MagicBuf); err != nil {
		return nil ,err
	}
	if string(MagicBuf) != Msg_hander {
		return nil , errors.New("msg_hander error")
	}

	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer,lengthBuf); err != nil {
		return nil , errors.New("length error")
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer,bodyBuf);err != nil {
		return nil ,err
	}
	return bodyBuf, nil
}