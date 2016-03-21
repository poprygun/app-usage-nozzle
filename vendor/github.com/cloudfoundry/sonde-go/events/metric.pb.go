// Code generated by protoc-gen-gogo.
// source: metric.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// / A ValueMetric indicates the value of a metric at an instant in time.
type ValueMetric struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,req,name=value" json:"value,omitempty"`
	Unit             *string  `protobuf:"bytes,3,req,name=unit" json:"unit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ValueMetric) Reset()                    { *m = ValueMetric{} }
func (m *ValueMetric) String() string            { return proto.CompactTextString(m) }
func (*ValueMetric) ProtoMessage()               {}
func (*ValueMetric) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

func (m *ValueMetric) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ValueMetric) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ValueMetric) GetUnit() string {
	if m != nil && m.Unit != nil {
		return *m.Unit
	}
	return ""
}

// / A CounterEvent represents the increment of a counter. It contains only the change in the value; it is the responsibility of downstream consumers to maintain the value of the counter.
type CounterEvent struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Delta            *uint64 `protobuf:"varint,2,req,name=delta" json:"delta,omitempty"`
	Total            *uint64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CounterEvent) Reset()                    { *m = CounterEvent{} }
func (m *CounterEvent) String() string            { return proto.CompactTextString(m) }
func (*CounterEvent) ProtoMessage()               {}
func (*CounterEvent) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{1} }

func (m *CounterEvent) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CounterEvent) GetDelta() uint64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *CounterEvent) GetTotal() uint64 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

// / A ContainerMetric records resource usage of an app in a container.
type ContainerMetric struct {
	ApplicationId    *string  `protobuf:"bytes,1,req,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32   `protobuf:"varint,2,req,name=instanceIndex" json:"instanceIndex,omitempty"`
	CpuPercentage    *float64 `protobuf:"fixed64,3,req,name=cpuPercentage" json:"cpuPercentage,omitempty"`
	MemoryBytes      *uint64  `protobuf:"varint,4,req,name=memoryBytes" json:"memoryBytes,omitempty"`
	DiskBytes        *uint64  `protobuf:"varint,5,req,name=diskBytes" json:"diskBytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ContainerMetric) Reset()                    { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string            { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()               {}
func (*ContainerMetric) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{2} }

func (m *ContainerMetric) GetApplicationId() string {
	if m != nil && m.ApplicationId != nil {
		return *m.ApplicationId
	}
	return ""
}

func (m *ContainerMetric) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *ContainerMetric) GetCpuPercentage() float64 {
	if m != nil && m.CpuPercentage != nil {
		return *m.CpuPercentage
	}
	return 0
}

func (m *ContainerMetric) GetMemoryBytes() uint64 {
	if m != nil && m.MemoryBytes != nil {
		return *m.MemoryBytes
	}
	return 0
}

func (m *ContainerMetric) GetDiskBytes() uint64 {
	if m != nil && m.DiskBytes != nil {
		return *m.DiskBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*ValueMetric)(nil), "events.ValueMetric")
	proto.RegisterType((*CounterEvent)(nil), "events.CounterEvent")
	proto.RegisterType((*ContainerMetric)(nil), "events.ContainerMetric")
}
func (m *ValueMetric) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ValueMetric) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.Value == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	} else {
		data[i] = 0x11
		i++
		i = encodeFixed64Metric(data, i, uint64(math.Float64bits(float64(*m.Value))))
	}
	if m.Unit == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("unit")
	} else {
		data[i] = 0x1a
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Unit)))
		i += copy(data[i:], *m.Unit)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CounterEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CounterEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.Delta == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("delta")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintMetric(data, i, uint64(*m.Delta))
	}
	if m.Total != nil {
		data[i] = 0x18
		i++
		i = encodeVarintMetric(data, i, uint64(*m.Total))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContainerMetric) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContainerMetric) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApplicationId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationId")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.ApplicationId)))
		i += copy(data[i:], *m.ApplicationId)
	}
	if m.InstanceIndex == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("instanceIndex")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintMetric(data, i, uint64(*m.InstanceIndex))
	}
	if m.CpuPercentage == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("cpuPercentage")
	} else {
		data[i] = 0x19
		i++
		i = encodeFixed64Metric(data, i, uint64(math.Float64bits(float64(*m.CpuPercentage))))
	}
	if m.MemoryBytes == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("memoryBytes")
	} else {
		data[i] = 0x20
		i++
		i = encodeVarintMetric(data, i, uint64(*m.MemoryBytes))
	}
	if m.DiskBytes == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("diskBytes")
	} else {
		data[i] = 0x28
		i++
		i = encodeVarintMetric(data, i, uint64(*m.DiskBytes))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Metric(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Metric(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetric(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ValueMetric) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Value != nil {
		n += 9
	}
	if m.Unit != nil {
		l = len(*m.Unit)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CounterEvent) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Delta != nil {
		n += 1 + sovMetric(uint64(*m.Delta))
	}
	if m.Total != nil {
		n += 1 + sovMetric(uint64(*m.Total))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContainerMetric) Size() (n int) {
	var l int
	_ = l
	if m.ApplicationId != nil {
		l = len(*m.ApplicationId)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.InstanceIndex != nil {
		n += 1 + sovMetric(uint64(*m.InstanceIndex))
	}
	if m.CpuPercentage != nil {
		n += 9
	}
	if m.MemoryBytes != nil {
		n += 1 + sovMetric(uint64(*m.MemoryBytes))
	}
	if m.DiskBytes != nil {
		n += 1 + sovMetric(uint64(*m.DiskBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetric(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValueMetric) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValueMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValueMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Value = &v2
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Unit = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("unit")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CounterEvent) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CounterEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CounterEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delta = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Total = &v
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("delta")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerMetric) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.ApplicationId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceIndex", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InstanceIndex = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuPercentage", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.CpuPercentage = &v2
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MemoryBytes = &v
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DiskBytes = &v
			hasFields[0] |= uint64(0x00000010)
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationId")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("instanceIndex")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("cpuPercentage")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("memoryBytes")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("diskBytes")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetric(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetric
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetric
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetric
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipMetric(data[start:])
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
	ErrInvalidLengthMetric = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetric   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorMetric = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x55, 0x4a, 0x82, 0x54, 0x37, 0x11, 0x22, 0x08, 0x29, 0xea, 0x80, 0x20, 0x53, 0x17, 0xd2,
	0x15, 0x31, 0x30, 0xa4, 0x02, 0xa9, 0x03, 0x12, 0x13, 0xbb, 0x6b, 0x5f, 0x83, 0x45, 0xe2, 0x8b,
	0x9c, 0x73, 0x45, 0x16, 0xbe, 0x8f, 0x91, 0x4f, 0x40, 0x7c, 0x09, 0xb6, 0xd3, 0x8d, 0xe1, 0xa4,
	0x7b, 0xef, 0xe9, 0xbd, 0x77, 0x36, 0x4b, 0x3b, 0x20, 0xa3, 0x44, 0xd5, 0x1b, 0x24, 0xcc, 0x4f,
	0xe1, 0x00, 0x9a, 0x86, 0xe5, 0x6d, 0xa3, 0xe8, 0xcd, 0xee, 0x2a, 0x81, 0xdd, 0xba, 0xc1, 0x06,
	0xd7, 0x41, 0xde, 0xd9, 0x7d, 0x40, 0x01, 0x84, 0x6d, 0xb2, 0x2d, 0x99, 0xb5, 0x4a, 0x4e, 0x7b,
	0x79, 0xc7, 0x16, 0xaf, 0xbc, 0xb5, 0xf0, 0x1c, 0x72, 0xf3, 0x94, 0xc5, 0x9a, 0x77, 0x50, 0x44,
	0xd7, 0xb3, 0xd5, 0x3c, 0xcf, 0x58, 0x72, 0xf0, 0x62, 0x31, 0x73, 0x30, 0xf2, 0xa2, 0xd5, 0x8a,
	0x8a, 0x13, 0x2f, 0x96, 0xf7, 0x2c, 0xdd, 0xa0, 0xd5, 0x04, 0xe6, 0xd1, 0x5f, 0xf1, 0xdf, 0x2a,
	0xa1, 0x25, 0x1e, 0xac, 0xb1, 0x87, 0x84, 0xc4, 0x5b, 0xe7, 0x8d, 0x56, 0x71, 0xf9, 0xc9, 0xce,
	0x36, 0xa8, 0x89, 0x2b, 0x0d, 0xe6, 0xd8, 0x7c, 0xc9, 0x32, 0xde, 0xf7, 0xad, 0x12, 0x9c, 0x14,
	0xea, 0xad, 0x3c, 0xe6, 0x38, 0x5a, 0xe9, 0x81, 0xb8, 0x16, 0xb0, 0xd5, 0x12, 0x3e, 0x42, 0x5e,
	0xe2, 0x69, 0xd1, 0xdb, 0x17, 0x30, 0xc2, 0x55, 0xf3, 0x06, 0xc2, 0x4d, 0x51, 0x7e, 0xc1, 0x16,
	0x1d, 0x74, 0x68, 0xc6, 0x7a, 0x24, 0x18, 0x8a, 0x38, 0x74, 0x9f, 0xb3, 0xb9, 0x54, 0xc3, 0xfb,
	0x44, 0x25, 0x9e, 0xaa, 0x1f, 0xbe, 0x7e, 0xaf, 0xa2, 0x6f, 0x37, 0x3f, 0x6e, 0xd8, 0x0d, 0x9a,
	0xa6, 0x12, 0x2d, 0x5a, 0xb9, 0x77, 0x0f, 0x92, 0x66, 0xac, 0xa4, 0xc1, 0x7e, 0x40, 0xd7, 0x57,
	0x4d, 0x3f, 0x5c, 0x67, 0xd3, 0x95, 0x4f, 0x5c, 0x90, 0x4b, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xa4, 0xe5, 0xfd, 0x87, 0x01, 0x00, 0x00,
}