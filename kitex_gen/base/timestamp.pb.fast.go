// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package base

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Timestamp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Timestamp[number], err)
}

func (x *Timestamp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Seconds, offset, err = fastpb.ReadSint64(buf, _type)
	return offset, err
}

func (x *Timestamp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Nanos, offset, err = fastpb.ReadSint32(buf, _type)
	return offset, err
}

func (x *Timestamp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Timestamp) fastWriteField1(buf []byte) (offset int) {
	if x.Seconds == 0 {
		return offset
	}
	offset += fastpb.WriteSint64(buf[offset:], 1, x.GetSeconds())
	return offset
}

func (x *Timestamp) fastWriteField2(buf []byte) (offset int) {
	if x.Nanos == 0 {
		return offset
	}
	offset += fastpb.WriteSint32(buf[offset:], 2, x.GetNanos())
	return offset
}

func (x *Timestamp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Timestamp) sizeField1() (n int) {
	if x.Seconds == 0 {
		return n
	}
	n += fastpb.SizeSint64(1, x.GetSeconds())
	return n
}

func (x *Timestamp) sizeField2() (n int) {
	if x.Nanos == 0 {
		return n
	}
	n += fastpb.SizeSint32(2, x.GetNanos())
	return n
}

var fieldIDToName_Timestamp = map[int32]string{
	1: "Seconds",
	2: "Nanos",
}
