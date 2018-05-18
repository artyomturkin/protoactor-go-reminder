// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: messages.proto

/*
	Package reminder is a generated protocol buffer package.

	It is generated from these files:
		messages.proto

	It has these top-level messages:
		Remind
		Reminded
		Snapshot
*/
package reminder

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/AsynkronIT/protoactor-go/actor"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Remind struct {
	Receiver *actor.PID                  `protobuf:"bytes,1,opt,name=Receiver" json:"Receiver,omitempty"`
	At       *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=At" json:"At,omitempty"`
	Message  *google_protobuf2.Any       `protobuf:"bytes,3,opt,name=Message" json:"Message,omitempty"`
}

func (m *Remind) Reset()                    { *m = Remind{} }
func (*Remind) ProtoMessage()               {}
func (*Remind) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{0} }

func (m *Remind) GetReceiver() *actor.PID {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Remind) GetAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *Remind) GetMessage() *google_protobuf2.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

type Reminded struct {
	At *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=At" json:"At,omitempty"`
}

func (m *Reminded) Reset()                    { *m = Reminded{} }
func (*Reminded) ProtoMessage()               {}
func (*Reminded) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{1} }

func (m *Reminded) GetAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

type Snapshot struct {
	Reminds []*Remind                   `protobuf:"bytes,1,rep,name=Reminds" json:"Reminds,omitempty"`
	At      *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=At" json:"At,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{2} }

func (m *Snapshot) GetReminds() []*Remind {
	if m != nil {
		return m.Reminds
	}
	return nil
}

func (m *Snapshot) GetAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.At
	}
	return nil
}

func init() {
	proto.RegisterType((*Remind)(nil), "reminder.Remind")
	proto.RegisterType((*Reminded)(nil), "reminder.Reminded")
	proto.RegisterType((*Snapshot)(nil), "reminder.Snapshot")
}
func (this *Remind) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Remind)
	if !ok {
		that2, ok := that.(Remind)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Receiver.Equal(that1.Receiver) {
		return false
	}
	if !this.At.Equal(that1.At) {
		return false
	}
	if !this.Message.Equal(that1.Message) {
		return false
	}
	return true
}
func (this *Reminded) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Reminded)
	if !ok {
		that2, ok := that.(Reminded)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.At.Equal(that1.At) {
		return false
	}
	return true
}
func (this *Snapshot) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Snapshot)
	if !ok {
		that2, ok := that.(Snapshot)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Reminds) != len(that1.Reminds) {
		return false
	}
	for i := range this.Reminds {
		if !this.Reminds[i].Equal(that1.Reminds[i]) {
			return false
		}
	}
	if !this.At.Equal(that1.At) {
		return false
	}
	return true
}
func (this *Remind) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&reminder.Remind{")
	if this.Receiver != nil {
		s = append(s, "Receiver: "+fmt.Sprintf("%#v", this.Receiver)+",\n")
	}
	if this.At != nil {
		s = append(s, "At: "+fmt.Sprintf("%#v", this.At)+",\n")
	}
	if this.Message != nil {
		s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Reminded) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&reminder.Reminded{")
	if this.At != nil {
		s = append(s, "At: "+fmt.Sprintf("%#v", this.At)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Snapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&reminder.Snapshot{")
	if this.Reminds != nil {
		s = append(s, "Reminds: "+fmt.Sprintf("%#v", this.Reminds)+",\n")
	}
	if this.At != nil {
		s = append(s, "At: "+fmt.Sprintf("%#v", this.At)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessages(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Remind) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Remind) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Receiver != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(m.Receiver.Size()))
		n1, err := m.Receiver.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.At != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(m.At.Size()))
		n2, err := m.At.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Message != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(m.Message.Size()))
		n3, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *Reminded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reminded) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.At != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(m.At.Size()))
		n4, err := m.At.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Reminds) > 0 {
		for _, msg := range m.Reminds {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMessages(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.At != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(m.At.Size()))
		n5, err := m.At.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Remind) Size() (n int) {
	var l int
	_ = l
	if m.Receiver != nil {
		l = m.Receiver.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.At != nil {
		l = m.At.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *Reminded) Size() (n int) {
	var l int
	_ = l
	if m.At != nil {
		l = m.At.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Reminds) > 0 {
		for _, e := range m.Reminds {
			l = e.Size()
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	if m.At != nil {
		l = m.At.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Remind) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Remind{`,
		`Receiver:` + strings.Replace(fmt.Sprintf("%v", this.Receiver), "PID", "actor.PID", 1) + `,`,
		`At:` + strings.Replace(fmt.Sprintf("%v", this.At), "Timestamp", "google_protobuf1.Timestamp", 1) + `,`,
		`Message:` + strings.Replace(fmt.Sprintf("%v", this.Message), "Any", "google_protobuf2.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Reminded) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Reminded{`,
		`At:` + strings.Replace(fmt.Sprintf("%v", this.At), "Timestamp", "google_protobuf1.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Snapshot{`,
		`Reminds:` + strings.Replace(fmt.Sprintf("%v", this.Reminds), "Remind", "Remind", 1) + `,`,
		`At:` + strings.Replace(fmt.Sprintf("%v", this.At), "Timestamp", "google_protobuf1.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessages(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Remind) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Remind: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Remind: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receiver == nil {
				m.Receiver = &actor.PID{}
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field At", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.At == nil {
				m.At = &google_protobuf1.Timestamp{}
			}
			if err := m.At.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &google_protobuf2.Any{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Reminded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reminded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reminded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field At", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.At == nil {
				m.At = &google_protobuf1.Timestamp{}
			}
			if err := m.At.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reminds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reminds = append(m.Reminds, &Remind{})
			if err := m.Reminds[len(m.Reminds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field At", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.At == nil {
				m.At = &google_protobuf1.Timestamp{}
			}
			if err := m.At.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessages
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessages
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessages(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessages = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("messages.proto", fileDescriptorMessages) }

var fileDescriptorMessages = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4a, 0xcd, 0xcd,
	0xcc, 0x4b, 0x49, 0x2d, 0x92, 0xe2, 0x4e, 0x4c, 0x2e, 0xc9, 0x2f, 0x82, 0x08, 0x4b, 0xc9, 0xa7,
	0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99, 0xb9, 0xa9,
	0xc5, 0x25, 0x89, 0xb9, 0x05, 0x50, 0x05, 0x92, 0xe8, 0x0a, 0x12, 0xf3, 0x2a, 0x21, 0x52, 0x4a,
	0x3d, 0x8c, 0x5c, 0x6c, 0x41, 0x60, 0x53, 0x85, 0xd4, 0xb8, 0x38, 0x82, 0x52, 0x93, 0x53, 0x33,
	0xcb, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x20, 0xd6, 0x04, 0x78,
	0xba, 0x04, 0xc1, 0xe5, 0x84, 0xb4, 0xb8, 0x98, 0x1c, 0x4b, 0x24, 0x98, 0xc0, 0x2a, 0xa4, 0xf4,
	0x20, 0x46, 0xeb, 0xc1, 0x8c, 0xd6, 0x0b, 0x81, 0xd9, 0x1d, 0xc4, 0xe4, 0x58, 0x22, 0xa4, 0xc7,
	0xc5, 0xee, 0x0b, 0xf1, 0x83, 0x04, 0x33, 0x58, 0x83, 0x08, 0x86, 0x06, 0xc7, 0xbc, 0xca, 0x20,
	0x98, 0x22, 0x25, 0x33, 0x90, 0x1b, 0xc0, 0x7e, 0x4c, 0x21, 0xc5, 0x1e, 0xa5, 0x24, 0x2e, 0x8e,
	0xe0, 0xbc, 0xc4, 0x82, 0xe2, 0x8c, 0xfc, 0x12, 0x21, 0x2d, 0x2e, 0x76, 0x88, 0x19, 0xc5, 0x12,
	0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x02, 0x7a, 0xb0, 0x70, 0xd3, 0x83, 0x48, 0x04, 0xc1, 0x14,
	0x90, 0x62, 0x87, 0x93, 0xce, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28,
	0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x9b, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x1c,
	0x3d, 0x8b, 0xc4, 0x01, 0x00, 0x00,
}