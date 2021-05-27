// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package v1

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type BinaryOp int8

const (
	BinaryOpEQ         BinaryOp = 0
	BinaryOpNE         BinaryOp = 1
	BinaryOpLT         BinaryOp = 2
	BinaryOpGT         BinaryOp = 3
	BinaryOpLE         BinaryOp = 4
	BinaryOpGE         BinaryOp = 5
	BinaryOpHAVING     BinaryOp = 6
	BinaryOpNOT_HAVING BinaryOp = 7
)

var EnumNamesBinaryOp = map[BinaryOp]string{
	BinaryOpEQ:         "EQ",
	BinaryOpNE:         "NE",
	BinaryOpLT:         "LT",
	BinaryOpGT:         "GT",
	BinaryOpLE:         "LE",
	BinaryOpGE:         "GE",
	BinaryOpHAVING:     "HAVING",
	BinaryOpNOT_HAVING: "NOT_HAVING",
}

var EnumValuesBinaryOp = map[string]BinaryOp{
	"EQ":         BinaryOpEQ,
	"NE":         BinaryOpNE,
	"LT":         BinaryOpLT,
	"GT":         BinaryOpGT,
	"LE":         BinaryOpLE,
	"GE":         BinaryOpGE,
	"HAVING":     BinaryOpHAVING,
	"NOT_HAVING": BinaryOpNOT_HAVING,
}

func (v BinaryOp) String() string {
	if s, ok := EnumNamesBinaryOp[v]; ok {
		return s
	}
	return "BinaryOp(" + strconv.FormatInt(int64(v), 10) + ")"
}

type TypedPair byte

const (
	TypedPairNONE    TypedPair = 0
	TypedPairIntPair TypedPair = 1
	TypedPairStrPair TypedPair = 2
)

var EnumNamesTypedPair = map[TypedPair]string{
	TypedPairNONE:    "NONE",
	TypedPairIntPair: "IntPair",
	TypedPairStrPair: "StrPair",
}

var EnumValuesTypedPair = map[string]TypedPair{
	"NONE":    TypedPairNONE,
	"IntPair": TypedPairIntPair,
	"StrPair": TypedPairStrPair,
}

func (v TypedPair) String() string {
	if s, ok := EnumNamesTypedPair[v]; ok {
		return s
	}
	return "TypedPair(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Sort int8

const (
	SortDESC Sort = 0
	SortASC  Sort = 1
)

var EnumNamesSort = map[Sort]string{
	SortDESC: "DESC",
	SortASC:  "ASC",
}

var EnumValuesSort = map[string]Sort{
	"DESC": SortDESC,
	"ASC":  SortASC,
}

func (v Sort) String() string {
	if s, ok := EnumNamesSort[v]; ok {
		return s
	}
	return "Sort(" + strconv.FormatInt(int64(v), 10) + ")"
}

type BinaryOps struct {
	_tab flatbuffers.Table
}

func GetRootAsBinaryOps(buf []byte, offset flatbuffers.UOffsetT) *BinaryOps {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BinaryOps{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBinaryOps(buf []byte, offset flatbuffers.UOffsetT) *BinaryOps {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BinaryOps{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BinaryOps) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BinaryOps) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BinaryOps) Ops(j int) BinaryOp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return BinaryOp(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *BinaryOps) OpsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BinaryOps) MutateOps(j int, n BinaryOp) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

func BinaryOpsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BinaryOpsAddOps(builder *flatbuffers.Builder, ops flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ops), 0)
}
func BinaryOpsStartOpsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BinaryOpsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type IntPair struct {
	_tab flatbuffers.Table
}

func GetRootAsIntPair(buf []byte, offset flatbuffers.UOffsetT) *IntPair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IntPair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIntPair(buf []byte, offset flatbuffers.UOffsetT) *IntPair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IntPair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IntPair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IntPair) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IntPair) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *IntPair) Values(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *IntPair) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *IntPair) MutateValues(j int, n int64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func IntPairStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func IntPairAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func IntPairAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(values), 0)
}
func IntPairStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func IntPairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type StrPair struct {
	_tab flatbuffers.Table
}

func GetRootAsStrPair(buf []byte, offset flatbuffers.UOffsetT) *StrPair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StrPair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStrPair(buf []byte, offset flatbuffers.UOffsetT) *StrPair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StrPair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StrPair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StrPair) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StrPair) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *StrPair) Values(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *StrPair) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func StrPairStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StrPairAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func StrPairAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(values), 0)
}
func StrPairStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StrPairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Pair struct {
	_tab flatbuffers.Table
}

func GetRootAsPair(buf []byte, offset flatbuffers.UOffsetT) *Pair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Pair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPair(buf []byte, offset flatbuffers.UOffsetT) *Pair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Pair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Pair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Pair) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Pair) PairType() TypedPair {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return TypedPair(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Pair) MutatePairType(n TypedPair) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Pair) Pair(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func PairStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PairAddPairType(builder *flatbuffers.Builder, pairType TypedPair) {
	builder.PrependByteSlot(0, byte(pairType), 0)
}
func PairAddPair(builder *flatbuffers.Builder, pair flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pair), 0)
}
func PairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type PairQuery struct {
	_tab flatbuffers.Table
}

func GetRootAsPairQuery(buf []byte, offset flatbuffers.UOffsetT) *PairQuery {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PairQuery{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPairQuery(buf []byte, offset flatbuffers.UOffsetT) *PairQuery {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PairQuery{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PairQuery) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PairQuery) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PairQuery) Ops(obj *BinaryOps) *BinaryOps {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BinaryOps)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PairQuery) Condition(obj *Pair) *Pair {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Pair)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PairQueryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PairQueryAddOps(builder *flatbuffers.Builder, ops flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ops), 0)
}
func PairQueryAddCondition(builder *flatbuffers.Builder, condition flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(condition), 0)
}
func PairQueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type QueryOrder struct {
	_tab flatbuffers.Table
}

func GetRootAsQueryOrder(buf []byte, offset flatbuffers.UOffsetT) *QueryOrder {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &QueryOrder{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsQueryOrder(buf []byte, offset flatbuffers.UOffsetT) *QueryOrder {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &QueryOrder{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *QueryOrder) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *QueryOrder) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *QueryOrder) KeyName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *QueryOrder) Sort() Sort {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Sort(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *QueryOrder) MutateSort(n Sort) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func QueryOrderStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func QueryOrderAddKeyName(builder *flatbuffers.Builder, keyName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(keyName), 0)
}
func QueryOrderAddSort(builder *flatbuffers.Builder, sort Sort) {
	builder.PrependInt8Slot(1, int8(sort), 0)
}
func QueryOrderEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Entity struct {
	_tab flatbuffers.Table
}

func GetRootAsEntity(buf []byte, offset flatbuffers.UOffsetT) *Entity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Entity{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntity(buf []byte, offset flatbuffers.UOffsetT) *Entity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Entity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Entity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Entity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Entity) EntityId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entity) TimestampNanoseconds() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entity) MutateTimestampNanoseconds(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *Entity) DataBinary(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Entity) DataBinaryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Entity) DataBinaryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entity) MutateDataBinary(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Entity) Fields(obj *Pair, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Entity) FieldsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func EntityAddEntityId(builder *flatbuffers.Builder, entityId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entityId), 0)
}
func EntityAddTimestampNanoseconds(builder *flatbuffers.Builder, timestampNanoseconds uint64) {
	builder.PrependUint64Slot(1, timestampNanoseconds, 0)
}
func EntityAddDataBinary(builder *flatbuffers.Builder, dataBinary flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dataBinary), 0)
}
func EntityStartDataBinaryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func EntityAddFields(builder *flatbuffers.Builder, fields flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(fields), 0)
}
func EntityStartFieldsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type TracesResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsTracesResponse(buf []byte, offset flatbuffers.UOffsetT) *TracesResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TracesResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTracesResponse(buf []byte, offset flatbuffers.UOffsetT) *TracesResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TracesResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TracesResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TracesResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TracesResponse) Entities(obj *Entity, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TracesResponse) EntitiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TracesResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TracesResponseAddEntities(builder *flatbuffers.Builder, entities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entities), 0)
}
func TracesResponseStartEntitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TracesResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Projection struct {
	_tab flatbuffers.Table
}

func GetRootAsProjection(buf []byte, offset flatbuffers.UOffsetT) *Projection {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Projection{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProjection(buf []byte, offset flatbuffers.UOffsetT) *Projection {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Projection{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Projection) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Projection) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Projection) KeyNames(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Projection) KeyNamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProjectionStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ProjectionAddKeyNames(builder *flatbuffers.Builder, keyNames flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(keyNames), 0)
}
func ProjectionStartKeyNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProjectionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type RangeQuery struct {
	_tab flatbuffers.Table
}

func GetRootAsRangeQuery(buf []byte, offset flatbuffers.UOffsetT) *RangeQuery {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RangeQuery{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRangeQuery(buf []byte, offset flatbuffers.UOffsetT) *RangeQuery {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RangeQuery{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RangeQuery) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RangeQuery) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RangeQuery) Begin() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RangeQuery) MutateBegin(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *RangeQuery) End() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RangeQuery) MutateEnd(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func RangeQueryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RangeQueryAddBegin(builder *flatbuffers.Builder, begin uint64) {
	builder.PrependUint64Slot(0, begin, 0)
}
func RangeQueryAddEnd(builder *flatbuffers.Builder, end uint64) {
	builder.PrependUint64Slot(1, end, 0)
}
func RangeQueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type EntityCriteria struct {
	_tab flatbuffers.Table
}

func GetRootAsEntityCriteria(buf []byte, offset flatbuffers.UOffsetT) *EntityCriteria {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EntityCriteria{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntityCriteria(buf []byte, offset flatbuffers.UOffsetT) *EntityCriteria {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EntityCriteria{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EntityCriteria) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EntityCriteria) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EntityCriteria) Metatdata(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *EntityCriteria) TimestampNanoseconds(obj *RangeQuery) *RangeQuery {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RangeQuery)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *EntityCriteria) Offset() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntityCriteria) MutateOffset(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *EntityCriteria) Limit() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntityCriteria) MutateLimit(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *EntityCriteria) OrderBy(obj *QueryOrder) *QueryOrder {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(QueryOrder)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *EntityCriteria) Fields(obj *PairQuery, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EntityCriteria) FieldsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *EntityCriteria) Projection(obj *Projection) *Projection {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Projection)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func EntityCriteriaStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func EntityCriteriaAddMetatdata(builder *flatbuffers.Builder, metatdata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metatdata), 0)
}
func EntityCriteriaAddTimestampNanoseconds(builder *flatbuffers.Builder, timestampNanoseconds flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(timestampNanoseconds), 0)
}
func EntityCriteriaAddOffset(builder *flatbuffers.Builder, offset uint32) {
	builder.PrependUint32Slot(2, offset, 0)
}
func EntityCriteriaAddLimit(builder *flatbuffers.Builder, limit uint32) {
	builder.PrependUint32Slot(3, limit, 0)
}
func EntityCriteriaAddOrderBy(builder *flatbuffers.Builder, orderBy flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(orderBy), 0)
}
func EntityCriteriaAddFields(builder *flatbuffers.Builder, fields flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(fields), 0)
}
func EntityCriteriaStartFieldsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EntityCriteriaAddProjection(builder *flatbuffers.Builder, projection flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(projection), 0)
}
func EntityCriteriaEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
