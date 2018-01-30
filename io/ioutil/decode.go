package ioutil

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// ReadUint8 读取uint8
func ReadUint8(r io.Reader) (uint8, error) {

	buffer := make([]byte, 1)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	if read != 1 {
		return 0, fmt.Errorf("read %d != 1", read)
	}

	return uint8(buffer[0]), nil
}

// ReadUint16 读取uint16
func ReadUint16(r io.Reader) (uint16, error) {

	buffer := make([]byte, 2)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	if read != 2 {
		return 0, fmt.Errorf("read %d != 2", read)
	}

	return binary.BigEndian.Uint16(buffer), nil
}

// ReadUint32 读取uint32
func ReadUint32(r io.Reader) (uint32, error) {

	buffer := make([]byte, 4)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	if read != 4 {
		return 0, fmt.Errorf("read %d != 4", read)
	}

	return binary.BigEndian.Uint32(buffer), nil
}

// ReadUint64 读取uint64
func ReadUint64(r io.Reader) (uint64, error) {

	buffer := make([]byte, 8)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	if read != 8 {
		return 0, fmt.Errorf("read %d != 8", read)
	}

	return binary.BigEndian.Uint64(buffer), nil
}

// ReadString 读取字符串
func ReadString(r io.Reader) (string, error) {

	size, err := ReadUint64(r)
	if err != nil {
		return "", err
	}

	if size == 0 {
		return "", nil
	}

	buffer := make([]byte, int(size))
	read, err := r.Read(buffer)
	if err != nil {
		return "", err
	}

	if read != int(size) {
		return "", fmt.Errorf("read %d < %d", read, int(size))
	}

	return string(buffer), nil
}

// ReadTime 读取时间
func ReadTime(r io.Reader) (time.Time, error) {

	unix, err := ReadUint64(r)
	if err != nil {
		return time.Time{}, err
	}
	value := time.Unix(int64(unix), 0)

	locationName, err := ReadString(r)
	if err != nil {
		return time.Time{}, err
	}

	location, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, err
	}

	return value.In(location), nil
}
