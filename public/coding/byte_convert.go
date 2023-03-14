package coding

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

// binary convert 字节码转换

// TODO: 后续有时间测试、优化并完善，目前应该还有问题

func ByteToBool(i byte) bool {
	return i == 1
}

func BytesToUint8(data []byte) uint8 {
	_ = data[0]
	return data[0]
}

func BytesToUint16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

func BytesToUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func BytesToUint64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func BytesToInt16(data []byte) int16 {
	return int16(BytesToUint16(data))
}

func BytesToInt32(data []byte) int32 {
	return int32(BytesToUint16(data))
}

func BytesToInt64(data []byte) int64 {
	return int64(BytesToUint64(data))
}

func BytesToInt(data []byte) int {
	switch len(data) {
	case 2:
		return int(binary.LittleEndian.Uint16(data))
	case 4:
		return int(binary.LittleEndian.Uint32(data))
	case 8:
		return int(binary.LittleEndian.Uint64(data))
	}
	return 0
}

func BytesToIntBig(data []byte) int {
	switch len(data) {
	case 2:
		return int(binary.BigEndian.Uint16(data))
	case 4:
		return int(binary.BigEndian.Uint32(data))
	case 8:
		return int(binary.BigEndian.Uint64(data))
	}
	return 0
}

// PutToBytes 整数转换成字节码
func PutToBytes(data interface{}) []byte {
	var buf []byte
	switch data.(type) {
	case int:
		buf = make([]byte, 8)
		target, _ := data.(int)
		binary.LittleEndian.PutUint64(buf, uint64(target))
	case int16:
		buf = make([]byte, 2)
		target, _ := data.(int16)
		binary.LittleEndian.PutUint16(buf, uint16(target))
	case int32:
		buf = make([]byte, 4)
		target, _ := data.(int32)
		binary.LittleEndian.PutUint32(buf, uint32(target))
	case int64:
		buf = make([]byte, 8)
		target, _ := data.(int64)
		binary.LittleEndian.PutUint64(buf, uint64(target))
	case uint8:
		buf = make([]byte, 1)
		target, _ := data.(uint8)
		buf[0] = target
	case uint16:
		buf = make([]byte, 2)
		target, _ := data.(uint16)
		binary.LittleEndian.PutUint16(buf, target)
	case uint32:
		buf = make([]byte, 4)
		target, _ := data.(uint32)
		binary.LittleEndian.PutUint32(buf, target)
	case uint64:
		buf = make([]byte, 8)
		target, _ := data.(uint64)
		binary.LittleEndian.PutUint64(buf, target)
	case uint:
		buf = make([]byte, 8)
		target, _ := data.(uint)
		binary.LittleEndian.PutUint64(buf, uint64(target))
	}
	return buf
}

// PutToBytesBig 整数转换成字节码
func PutToBytesBig(data interface{}) []byte {
	var buf []byte
	switch data.(type) {
	case int:
		buf = make([]byte, 8)
		target, _ := data.(int)
		binary.BigEndian.PutUint64(buf, uint64(target))
	case int16:
		buf = make([]byte, 2)
		target, _ := data.(int16)
		binary.BigEndian.PutUint16(buf, uint16(target))
	case int32:
		buf = make([]byte, 4)
		target, _ := data.(int32)
		binary.BigEndian.PutUint32(buf, uint32(target))
	case int64:
		buf = make([]byte, 8)
		target, _ := data.(int64)
		binary.BigEndian.PutUint64(buf, uint64(target))
	case uint:
		buf = make([]byte, 8)
		target, _ := data.(uint)
		binary.BigEndian.PutUint64(buf, uint64(target))
	case uint16:
		buf = make([]byte, 2)
		target, _ := data.(uint16)
		binary.BigEndian.PutUint16(buf, target)
	case uint32:
		buf = make([]byte, 4)
		target, _ := data.(uint32)
		binary.BigEndian.PutUint32(buf, target)
	case uint64:
		buf = make([]byte, 8)
		target, _ := data.(uint64)
		binary.BigEndian.PutUint64(buf, target)
	}
	return buf
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func Float64ToBytesBig(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Float32ToBytesBig(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}

func FotmatToBytes(content interface{}) ([]byte, error) {
	contentBuffer := make([]byte, 0)
	var err error
	switch content.(type) {
	case nil:

	case string:
		dataString, _ := content.(string)
		contentBuffer = []byte(dataString)
		break
	case []byte:
		contentBuffer, _ = content.([]byte)
	case byte:
		byte, _ := content.(byte)
		contentBuffer = append(contentBuffer, byte)
	case bool:
		dataBool, _ := content.(bool)
		if dataBool {
			contentBuffer = append(contentBuffer, 1)
		} else {
			contentBuffer = append(contentBuffer, 0)
		}
	case int, int16, int32, int64, uint, uint16, uint32, uint64:
		contentBuffer = []byte(fmt.Sprintf("%d", content))
		break
	case float32:
		dataFloat, _ := content.(float32)
		contentBuffer = Float32ToBytesBig(dataFloat)
	default:
		if reflect.TypeOf(content).Kind() == reflect.Ptr {
			if reflect.ValueOf(content).IsNil() {

			} else {
				contentBuffer, err = json.Marshal(content)
			}
		} else if reflect.TypeOf(content).Kind() == reflect.Slice {
			if reflect.ValueOf(content).Len() == 0 {
				contentBuffer = []byte("[]")
			} else {
				contentBuffer, err = json.Marshal(content)
			}
		} else {
			contentBuffer, err = json.Marshal(content)
		}
		break
	}
	return contentBuffer, err
}
