// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *SendMailRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMailRequest[number], err)
}

func (x *SendMailRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = MailType(v)
	return offset, nil
}

func (x *SendMailRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.To, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMailRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Subject, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMailRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Body, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EmailMessage[number], err)
}

func (x *EmailMessage) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.To, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ContentType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Subject, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Body, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.From, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailMessage) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.ReplyTo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMailRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SendMailRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetType()))
	return offset
}

func (x *SendMailRequest) fastWriteField2(buf []byte) (offset int) {
	if x.To == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTo())
	return offset
}

func (x *SendMailRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Subject == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSubject())
	return offset
}

func (x *SendMailRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Body == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetBody())
	return offset
}

func (x *EmailMessage) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *EmailMessage) fastWriteField1(buf []byte) (offset int) {
	if x.To == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTo())
	return offset
}

func (x *EmailMessage) fastWriteField2(buf []byte) (offset int) {
	if x.ContentType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetContentType())
	return offset
}

func (x *EmailMessage) fastWriteField3(buf []byte) (offset int) {
	if x.Subject == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSubject())
	return offset
}

func (x *EmailMessage) fastWriteField4(buf []byte) (offset int) {
	if x.Body == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetBody())
	return offset
}

func (x *EmailMessage) fastWriteField5(buf []byte) (offset int) {
	if x.From == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetFrom())
	return offset
}

func (x *EmailMessage) fastWriteField6(buf []byte) (offset int) {
	if x.ReplyTo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetReplyTo())
	return offset
}

func (x *SendMailRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SendMailRequest) sizeField1() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetType()))
	return n
}

func (x *SendMailRequest) sizeField2() (n int) {
	if x.To == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTo())
	return n
}

func (x *SendMailRequest) sizeField3() (n int) {
	if x.Subject == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSubject())
	return n
}

func (x *SendMailRequest) sizeField4() (n int) {
	if x.Body == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetBody())
	return n
}

func (x *EmailMessage) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *EmailMessage) sizeField1() (n int) {
	if x.To == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTo())
	return n
}

func (x *EmailMessage) sizeField2() (n int) {
	if x.ContentType == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetContentType())
	return n
}

func (x *EmailMessage) sizeField3() (n int) {
	if x.Subject == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSubject())
	return n
}

func (x *EmailMessage) sizeField4() (n int) {
	if x.Body == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetBody())
	return n
}

func (x *EmailMessage) sizeField5() (n int) {
	if x.From == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetFrom())
	return n
}

func (x *EmailMessage) sizeField6() (n int) {
	if x.ReplyTo == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetReplyTo())
	return n
}

var fieldIDToName_SendMailRequest = map[int32]string{
	1: "Type",
	2: "To",
	3: "Subject",
	4: "Body",
}

var fieldIDToName_EmailMessage = map[int32]string{
	1: "To",
	2: "ContentType",
	3: "Subject",
	4: "Body",
	5: "From",
	6: "ReplyTo",
}

var _ = emptypb.File_google_protobuf_empty_proto
