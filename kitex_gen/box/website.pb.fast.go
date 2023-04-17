// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package box

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *GetWebsiteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetWebsiteRequest[number], err)
}

func (x *GetWebsiteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Key, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetWebsiteResponse[number], err)
}

func (x *GetWebsiteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Key, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.AvatarIcon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Background, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Language, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetWebsiteResponse) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.AllowAnonymous, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetWebsiteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetWebsiteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Key == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKey())
	return offset
}

func (x *GetWebsiteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *GetWebsiteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Key == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKey())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField4(buf []byte) (offset int) {
	if x.AvatarIcon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetAvatarIcon())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField5(buf []byte) (offset int) {
	if x.Background == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetBackground())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField6(buf []byte) (offset int) {
	if x.Language == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetLanguage())
	return offset
}

func (x *GetWebsiteResponse) fastWriteField7(buf []byte) (offset int) {
	if !x.AllowAnonymous {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.GetAllowAnonymous())
	return offset
}

func (x *GetWebsiteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetWebsiteRequest) sizeField1() (n int) {
	if x.Key == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetKey())
	return n
}

func (x *GetWebsiteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *GetWebsiteResponse) sizeField1() (n int) {
	if x.Key == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetKey())
	return n
}

func (x *GetWebsiteResponse) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *GetWebsiteResponse) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *GetWebsiteResponse) sizeField4() (n int) {
	if x.AvatarIcon == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetAvatarIcon())
	return n
}

func (x *GetWebsiteResponse) sizeField5() (n int) {
	if x.Background == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetBackground())
	return n
}

func (x *GetWebsiteResponse) sizeField6() (n int) {
	if x.Language == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetLanguage())
	return n
}

func (x *GetWebsiteResponse) sizeField7() (n int) {
	if !x.AllowAnonymous {
		return n
	}
	n += fastpb.SizeBool(7, x.GetAllowAnonymous())
	return n
}

var fieldIDToName_GetWebsiteRequest = map[int32]string{
	1: "Key",
}

var fieldIDToName_GetWebsiteResponse = map[int32]string{
	1: "Key",
	2: "Name",
	3: "Description",
	4: "AvatarIcon",
	5: "Background",
	6: "Language",
	7: "AllowAnonymous",
}