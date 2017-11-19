package ioutil

import (
	"encoding/binary"
	"io"
)

// ReadUint8 读取uint8
func ReadUint8(r io.Reader) (uint8, error) {

	buffer := make([]byte, 1)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	return uint8(buffer[0]), nil
}

// ReadUint16 读取uint16
func ReadUint16(r io.Reader) (uint16, error) {

	buffer := make([]byte, 2)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(buffer), nil
}

// ReadUint32 读取uint32
func ReadUint32(r io.Reader) (uint32, error) {

	buffer := make([]byte, 4)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(buffer), nil
}

// ReadUint64 读取uint64
func ReadUint64(r io.Reader) (uint64, error) {

	buffer := make([]byte, 8)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(buffer), nil
}

// ReadString 读取字符串
func ReadString(r io.Reader) (string, error) {

	size, err := ReadUint64(r)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, size)
	_, err = r.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
