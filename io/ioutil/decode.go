package ioutil

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// ReadUint8 读取uint8
func ReadUint8(r io.Reader) (uint8, int, error) {

	buffer := make([]byte, 1)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, 0, err
	}

	if read != 1 {
		return 0, 0, fmt.Errorf("read %d != 1", read)
	}

	return uint8(buffer[0]), read, nil
}

// ReadUint16 读取uint16
func ReadUint16(r io.Reader) (uint16, int, error) {

	buffer := make([]byte, 2)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, 0, err
	}

	if read != 2 {
		return 0, 0, fmt.Errorf("read %d != 2", read)
	}

	return binary.BigEndian.Uint16(buffer), read, nil
}

// ReadUint32 读取uint32
func ReadUint32(r io.Reader) (uint32, int, error) {

	buffer := make([]byte, 4)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, 0, err
	}

	if read != 4 {
		return 0, 0, fmt.Errorf("read %d != 4", read)
	}

	return binary.BigEndian.Uint32(buffer), read, nil
}

// ReadUint64 读取uint64
func ReadUint64(r io.Reader) (uint64, int, error) {

	buffer := make([]byte, 8)
	read, err := r.Read(buffer)
	if err != nil {
		return 0, 0, err
	}

	if read != 8 {
		return 0, 0, fmt.Errorf("read %d != 8", read)
	}

	return binary.BigEndian.Uint64(buffer), read, nil
}

// ReadString 读取字符串
func ReadString(r io.Reader) (string, int, error) {

	size, size1, err := ReadUint32(r)
	if err != nil {
		return "", 0, err
	}

	if size == 0 {
		return "", 0, nil
	}

	buffer := make([]byte, int(size))
	size2, err := r.Read(buffer)
	if err != nil {
		return "", 0, err
	}

	if size2 != int(size) {
		return "", 0, fmt.Errorf("read %d < %d", size2, size)
	}

	return string(buffer), size1 + size2, nil
}

// ReadTime 读取时间
func ReadTime(r io.Reader) (time.Time, int, error) {

	unix, size1, err := ReadUint64(r)
	if err != nil {
		return time.Time{}, 0, err
	}
	value := time.Unix(int64(unix), 0)

	locationName, size2, err := ReadString(r)
	if err != nil {
		return time.Time{}, 0, err
	}

	location, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, 0, err
	}

	return value.In(location), size1 + size2, nil
}
