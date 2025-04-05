package utils

import (
    "fmt"
    "time"
    "strconv"
    "hash/crc32"
    "encoding/binary"
)

var DocId = NewDocumentId()

func NewDocumentId() *docId {
	return &docId{}
}

type docId struct{}

// get new id
func (f *docId) Create() string {
    var ts = time.Now().UnixNano() /1e6
    var tsStr = strconv.FormatInt(ts,10)
    return NewDocumentId().CreateId(tsStr)
}

// get new id
func (f *docId) createDocId() string {
    var ts = time.Now().UnixNano() /1e6
    var tsStr = strconv.FormatInt(ts,10)
    return NewDocumentId().CreateId(tsStr)
}

func (f *docId) CreateId(val string) string {
    var crc32q = crc32.MakeTable(crc32.IEEE)
    var id = crc32.Checksum([]byte(val), crc32q)
    var idStr = strconv.FormatInt(int64(id), 16)
    return idStr
}

func int64ToBytes(i int64) []byte {
    var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}

func bytesToInt64(buf []byte) int64 {
    return int64(binary.BigEndian.Uint64(buf))
}

func main() {

    fmt.Println("idStr: ", NewDocumentId().Create())
    fmt.Println("idStr1 1611729480020: ", NewDocumentId().CreateId("1611729480020"))
    fmt.Println("idStr2 1611729480021: ", NewDocumentId().CreateId("1611729480021"))

}
