//go:build integration

package main

import (
	"norma_core/target/generated-sources/protobuf/gremlin/protobuf_unittest"
	"norma_core/target/generated-sources/protobuf/gremlin/protobuf_unittest_import"
	"testing"
)

func populateGoldenMessage(msg *protobuf_unittest.TestAllTypes) {
	msg.OptionalInt32 = 101
	msg.OptionalInt64 = 102
	msg.OptionalUint32 = 103
	msg.OptionalUint64 = 104
	msg.OptionalSint32 = 105
	msg.OptionalSint64 = 106
	msg.OptionalFixed32 = 107
	msg.OptionalFixed64 = 108
	msg.OptionalSfixed32 = 109
	msg.OptionalSfixed64 = 110
	msg.OptionalFloat = 111.0
	msg.OptionalDouble = 112.0
	msg.OptionalBool = true
	msg.OptionalString = "115"
	msg.OptionalBytes = []byte("116")

	msg.OptionalNestedMessage = &protobuf_unittest.TestAllTypes_NestedMessage{
		Bb: 118,
	}
	msg.OptionalForeignMessage = &protobuf_unittest.ForeignMessage{
		C: 119,
	}
	msg.OptionalImportMessage = &protobuf_unittest_import.ImportMessage{
		D: 120,
	}
	msg.OptionalNestedEnum = protobuf_unittest.TestAllTypes_BAZ
	msg.OptionalForeignEnum = protobuf_unittest.ForeignEnum_FOREIGN_BAZ
	msg.OptionalImportEnum = protobuf_unittest_import.ImportEnum_IMPORT_BAZ
	msg.OptionalStringPiece = "124"
	msg.OptionalCord = "125"
	msg.OptionalPublicImportMessage = &protobuf_unittest_import.PublicImportMessage{
		E: 126,
	}
	msg.OptionalLazyMessage = &protobuf_unittest.TestAllTypes_NestedMessage{
		Bb: 127,
	}
	msg.OptionalUnverifiedLazyMessage = &protobuf_unittest.TestAllTypes_NestedMessage{
		Bb: 128,
	}
	msg.RepeatedInt32 = []int32{201, 301}
	msg.RepeatedInt64 = []int64{202, 302}
	msg.RepeatedUint32 = []uint32{203, 303}
	msg.RepeatedUint64 = []uint64{204, 304}
	msg.RepeatedSint32 = []int32{205, 305}
	msg.RepeatedSint64 = []int64{206, 306}
	msg.RepeatedFixed32 = []uint32{207, 307}
	msg.RepeatedFixed64 = []uint64{208, 308}
	msg.RepeatedSfixed32 = []int32{209, 309}
	msg.RepeatedSfixed64 = []int64{210, 310}
	msg.RepeatedFloat = []float32{211.0, 311.0}
	msg.RepeatedDouble = []float64{212.0, 312.0}
	msg.RepeatedBool = []bool{true, false}
	msg.RepeatedString = []string{"215", "315"}
	msg.RepeatedBytes = [][]byte{[]byte("216"), []byte("316")}
	msg.RepeatedNestedMessage = []*protobuf_unittest.TestAllTypes_NestedMessage{
		{Bb: 218},
		{Bb: 318},
	}
	msg.RepeatedForeignMessage = []*protobuf_unittest.ForeignMessage{
		{C: 219},
		{C: 319},
	}
	msg.RepeatedImportMessage = []*protobuf_unittest_import.ImportMessage{
		{D: 220},
		{D: 320},
	}
	msg.RepeatedNestedEnum = []protobuf_unittest.TestAllTypes_NestedEnum{
		protobuf_unittest.TestAllTypes_BAR,
		protobuf_unittest.TestAllTypes_BAZ,
	}
	msg.RepeatedForeignEnum = []protobuf_unittest.ForeignEnum{
		protobuf_unittest.ForeignEnum_FOREIGN_BAR,
		protobuf_unittest.ForeignEnum_FOREIGN_BAZ,
	}
	msg.RepeatedImportEnum = []protobuf_unittest_import.ImportEnum{
		protobuf_unittest_import.ImportEnum_IMPORT_BAR,
		protobuf_unittest_import.ImportEnum_IMPORT_BAZ,
	}
	msg.RepeatedStringPiece = []string{"224", "324"}
	msg.RepeatedCord = []string{"225", "325"}
	msg.RepeatedLazyMessage = []*protobuf_unittest.TestAllTypes_NestedMessage{
		{Bb: 227},
		{Bb: 327},
	}
	msg.DefaultInt32 = 401
	msg.DefaultInt64 = 402
	msg.DefaultUint32 = 403
	msg.DefaultUint64 = 404
	msg.DefaultSint32 = 405
	msg.DefaultSint64 = 406
	msg.DefaultFixed32 = 407
	msg.DefaultFixed64 = 408
	msg.DefaultSfixed32 = 409
	msg.DefaultSfixed64 = 410
	msg.DefaultFloat = 411.0
	msg.DefaultDouble = 412.0
	msg.DefaultBool = false
	msg.DefaultString = "415"
	msg.DefaultBytes = []byte("416")
	msg.DefaultNestedEnum = protobuf_unittest.TestAllTypes_FOO
	msg.DefaultForeignEnum = protobuf_unittest.ForeignEnum_FOREIGN_FOO
	msg.DefaultImportEnum = protobuf_unittest_import.ImportEnum_IMPORT_FOO
	msg.DefaultStringPiece = "424"
	msg.DefaultCord = "425"
	msg.OneofUint32 = 601
}

func updateGoldenMessage(msg *protobuf_unittest.TestAllTypes, i int) {
	msg.OptionalInt32 = 101 + int32(i)
	msg.OptionalNestedMessage.Bb = 118 + int32(i)
	msg.RepeatedNestedMessage[0].Bb = 218 + int32(i)
	msg.RepeatedNestedMessage[1].Bb = 318 + int32(i)
	msg.OneofUint32 = 601 + uint32(i)
}

func BenchmarkMarshal(b *testing.B) {
	msg := &protobuf_unittest.TestAllTypes{}
	populateGoldenMessage(msg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updateGoldenMessage(msg, i)
		_ = msg.Marshal()
	}
}
