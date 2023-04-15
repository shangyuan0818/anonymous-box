// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Comment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 8:
		offset, err = x.fastReadField8(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Comment[number], err)
}

func (x *Comment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.WebsiteRefer, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	var v base.Timestamp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CreatedAt = &v
	return offset, nil
}

func (x *Comment) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v base.Timestamp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UpdatedAt = &v
	return offset, nil
}

func (x *ListCommentsRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentsRequest[number], err)
}

func (x *ListCommentsRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ListCommentsRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.WebsiteRefer, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ListCommentsRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v base.Pagination
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Pagination = &v
	return offset, nil
}

func (x *ListCommentsResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentsResponse[number], err)
}

func (x *ListCommentsResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadSint64(buf, _type)
	return offset, err
}

func (x *ListCommentsResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comments = append(x.Comments, &v)
	return offset, nil
}

func (x *GetCommentRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCommentRequest[number], err)
}

func (x *GetCommentRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetCommentRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *DeleteCommentRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteCommentRequest[number], err)
}

func (x *DeleteCommentRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *DeleteCommentRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Comment) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Comment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Comment) fastWriteField2(buf []byte) (offset int) {
	if x.WebsiteRefer == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetWebsiteRefer())
	return offset
}

func (x *Comment) fastWriteField3(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetName())
	return offset
}

func (x *Comment) fastWriteField4(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetEmail())
	return offset
}

func (x *Comment) fastWriteField5(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetUrl())
	return offset
}

func (x *Comment) fastWriteField6(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetContent())
	return offset
}

func (x *Comment) fastWriteField7(buf []byte) (offset int) {
	if x.CreatedAt == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 7, x.GetCreatedAt())
	return offset
}

func (x *Comment) fastWriteField8(buf []byte) (offset int) {
	if x.UpdatedAt == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 8, x.GetUpdatedAt())
	return offset
}

func (x *ListCommentsRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListCommentsRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *ListCommentsRequest) fastWriteField2(buf []byte) (offset int) {
	if x.WebsiteRefer == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetWebsiteRefer())
	return offset
}

func (x *ListCommentsRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Pagination == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetPagination())
	return offset
}

func (x *ListCommentsResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListCommentsResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteSint64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *ListCommentsResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Comments == nil {
		return offset
	}
	for i := range x.GetComments() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetComments()[i])
	}
	return offset
}

func (x *GetCommentRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetCommentRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetCommentRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetId())
	return offset
}

func (x *DeleteCommentRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteCommentRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *DeleteCommentRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetId())
	return offset
}

func (x *Comment) Size() (n int) {
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
	n += x.sizeField8()
	return n
}

func (x *Comment) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *Comment) sizeField2() (n int) {
	if x.WebsiteRefer == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetWebsiteRefer())
	return n
}

func (x *Comment) sizeField3() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetName())
	return n
}

func (x *Comment) sizeField4() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetEmail())
	return n
}

func (x *Comment) sizeField5() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetUrl())
	return n
}

func (x *Comment) sizeField6() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetContent())
	return n
}

func (x *Comment) sizeField7() (n int) {
	if x.CreatedAt == nil {
		return n
	}
	n += fastpb.SizeMessage(7, x.GetCreatedAt())
	return n
}

func (x *Comment) sizeField8() (n int) {
	if x.UpdatedAt == nil {
		return n
	}
	n += fastpb.SizeMessage(8, x.GetUpdatedAt())
	return n
}

func (x *ListCommentsRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListCommentsRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *ListCommentsRequest) sizeField2() (n int) {
	if x.WebsiteRefer == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetWebsiteRefer())
	return n
}

func (x *ListCommentsRequest) sizeField3() (n int) {
	if x.Pagination == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetPagination())
	return n
}

func (x *ListCommentsResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListCommentsResponse) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeSint64(1, x.GetTotal())
	return n
}

func (x *ListCommentsResponse) sizeField2() (n int) {
	if x.Comments == nil {
		return n
	}
	for i := range x.GetComments() {
		n += fastpb.SizeMessage(2, x.GetComments()[i])
	}
	return n
}

func (x *GetCommentRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetCommentRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *GetCommentRequest) sizeField2() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetId())
	return n
}

func (x *DeleteCommentRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteCommentRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *DeleteCommentRequest) sizeField2() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetId())
	return n
}

var fieldIDToName_Comment = map[int32]string{
	1: "Id",
	2: "WebsiteRefer",
	3: "Name",
	4: "Email",
	5: "Url",
	6: "Content",
	7: "CreatedAt",
	8: "UpdatedAt",
}

var fieldIDToName_ListCommentsRequest = map[int32]string{
	1: "UserId",
	2: "WebsiteRefer",
	3: "Pagination",
}

var fieldIDToName_ListCommentsResponse = map[int32]string{
	1: "Total",
	2: "Comments",
}

var fieldIDToName_GetCommentRequest = map[int32]string{
	1: "UserId",
	2: "Id",
}

var fieldIDToName_DeleteCommentRequest = map[int32]string{
	1: "UserId",
	2: "Id",
}

var _ = base.File_idl_base_empty_proto
var _ = base.File_idl_base_timestamp_proto
var _ = base.File_idl_base_pagination_proto
