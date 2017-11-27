// Package allorm was auto-generated by github.com/posener/orm; DO NOT EDIT
package allorm

import (
	"context"

	"github.com/posener/orm/common"
	"github.com/posener/orm/example"
)

// AllCount is a struct for counting rows of type All
type AllCount struct {
	example.All
	Count int64
}

// SelectBuilder builds an SQL SELECT statement parameters
type SelectBuilder struct {
	params  common.SelectParams
	orm     *orm
	columns columns
}

// Where applies where conditions on the query
func (b *SelectBuilder) Where(where common.Where) *SelectBuilder {
	b.params.Where = where
	return b
}

// Limit applies rows limit on the query response
func (b *SelectBuilder) Limit(limit int64) *SelectBuilder {
	b.params.Page.Limit = limit
	return b
}

// Page applies rows offset and limit on the query response
func (b *SelectBuilder) Page(offset, limit int64) *SelectBuilder {
	b.params.Page.Offset = offset
	b.params.Page.Limit = limit
	return b
}

// SelectAuto adds Auto to the selected column of a query
func (b *SelectBuilder) SelectAuto() *SelectBuilder {
	b.columns.SelectAuto = true
	return b
}

// OrderByAuto set order to the query results according to column auto
func (b *SelectBuilder) OrderByAuto(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("auto", dir)
	return b
}

// GroupByAuto make the query group by column auto
func (b *SelectBuilder) GroupByAuto() *SelectBuilder {
	b.params.Groups.Add("auto")
	return b
}

// SelectNotNil adds NotNil to the selected column of a query
func (b *SelectBuilder) SelectNotNil() *SelectBuilder {
	b.columns.SelectNotNil = true
	return b
}

// OrderByNotNil set order to the query results according to column notnil
func (b *SelectBuilder) OrderByNotNil(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("notnil", dir)
	return b
}

// GroupByNotNil make the query group by column notnil
func (b *SelectBuilder) GroupByNotNil() *SelectBuilder {
	b.params.Groups.Add("notnil")
	return b
}

// SelectInt adds Int to the selected column of a query
func (b *SelectBuilder) SelectInt() *SelectBuilder {
	b.columns.SelectInt = true
	return b
}

// OrderByInt set order to the query results according to column int
func (b *SelectBuilder) OrderByInt(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("int", dir)
	return b
}

// GroupByInt make the query group by column int
func (b *SelectBuilder) GroupByInt() *SelectBuilder {
	b.params.Groups.Add("int")
	return b
}

// SelectInt8 adds Int8 to the selected column of a query
func (b *SelectBuilder) SelectInt8() *SelectBuilder {
	b.columns.SelectInt8 = true
	return b
}

// OrderByInt8 set order to the query results according to column int8
func (b *SelectBuilder) OrderByInt8(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("int8", dir)
	return b
}

// GroupByInt8 make the query group by column int8
func (b *SelectBuilder) GroupByInt8() *SelectBuilder {
	b.params.Groups.Add("int8")
	return b
}

// SelectInt16 adds Int16 to the selected column of a query
func (b *SelectBuilder) SelectInt16() *SelectBuilder {
	b.columns.SelectInt16 = true
	return b
}

// OrderByInt16 set order to the query results according to column int16
func (b *SelectBuilder) OrderByInt16(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("int16", dir)
	return b
}

// GroupByInt16 make the query group by column int16
func (b *SelectBuilder) GroupByInt16() *SelectBuilder {
	b.params.Groups.Add("int16")
	return b
}

// SelectInt32 adds Int32 to the selected column of a query
func (b *SelectBuilder) SelectInt32() *SelectBuilder {
	b.columns.SelectInt32 = true
	return b
}

// OrderByInt32 set order to the query results according to column int32
func (b *SelectBuilder) OrderByInt32(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("int32", dir)
	return b
}

// GroupByInt32 make the query group by column int32
func (b *SelectBuilder) GroupByInt32() *SelectBuilder {
	b.params.Groups.Add("int32")
	return b
}

// SelectInt64 adds Int64 to the selected column of a query
func (b *SelectBuilder) SelectInt64() *SelectBuilder {
	b.columns.SelectInt64 = true
	return b
}

// OrderByInt64 set order to the query results according to column int64
func (b *SelectBuilder) OrderByInt64(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("int64", dir)
	return b
}

// GroupByInt64 make the query group by column int64
func (b *SelectBuilder) GroupByInt64() *SelectBuilder {
	b.params.Groups.Add("int64")
	return b
}

// SelectUInt adds UInt to the selected column of a query
func (b *SelectBuilder) SelectUInt() *SelectBuilder {
	b.columns.SelectUInt = true
	return b
}

// OrderByUInt set order to the query results according to column uint
func (b *SelectBuilder) OrderByUInt(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("uint", dir)
	return b
}

// GroupByUInt make the query group by column uint
func (b *SelectBuilder) GroupByUInt() *SelectBuilder {
	b.params.Groups.Add("uint")
	return b
}

// SelectUInt8 adds UInt8 to the selected column of a query
func (b *SelectBuilder) SelectUInt8() *SelectBuilder {
	b.columns.SelectUInt8 = true
	return b
}

// OrderByUInt8 set order to the query results according to column uint8
func (b *SelectBuilder) OrderByUInt8(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("uint8", dir)
	return b
}

// GroupByUInt8 make the query group by column uint8
func (b *SelectBuilder) GroupByUInt8() *SelectBuilder {
	b.params.Groups.Add("uint8")
	return b
}

// SelectUInt16 adds UInt16 to the selected column of a query
func (b *SelectBuilder) SelectUInt16() *SelectBuilder {
	b.columns.SelectUInt16 = true
	return b
}

// OrderByUInt16 set order to the query results according to column uint16
func (b *SelectBuilder) OrderByUInt16(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("uint16", dir)
	return b
}

// GroupByUInt16 make the query group by column uint16
func (b *SelectBuilder) GroupByUInt16() *SelectBuilder {
	b.params.Groups.Add("uint16")
	return b
}

// SelectUInt32 adds UInt32 to the selected column of a query
func (b *SelectBuilder) SelectUInt32() *SelectBuilder {
	b.columns.SelectUInt32 = true
	return b
}

// OrderByUInt32 set order to the query results according to column uint32
func (b *SelectBuilder) OrderByUInt32(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("uint32", dir)
	return b
}

// GroupByUInt32 make the query group by column uint32
func (b *SelectBuilder) GroupByUInt32() *SelectBuilder {
	b.params.Groups.Add("uint32")
	return b
}

// SelectUInt64 adds UInt64 to the selected column of a query
func (b *SelectBuilder) SelectUInt64() *SelectBuilder {
	b.columns.SelectUInt64 = true
	return b
}

// OrderByUInt64 set order to the query results according to column uint64
func (b *SelectBuilder) OrderByUInt64(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("uint64", dir)
	return b
}

// GroupByUInt64 make the query group by column uint64
func (b *SelectBuilder) GroupByUInt64() *SelectBuilder {
	b.params.Groups.Add("uint64")
	return b
}

// SelectTime adds Time to the selected column of a query
func (b *SelectBuilder) SelectTime() *SelectBuilder {
	b.columns.SelectTime = true
	return b
}

// OrderByTime set order to the query results according to column time
func (b *SelectBuilder) OrderByTime(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("time", dir)
	return b
}

// GroupByTime make the query group by column time
func (b *SelectBuilder) GroupByTime() *SelectBuilder {
	b.params.Groups.Add("time")
	return b
}

// SelectVarCharString adds VarCharString to the selected column of a query
func (b *SelectBuilder) SelectVarCharString() *SelectBuilder {
	b.columns.SelectVarCharString = true
	return b
}

// OrderByVarCharString set order to the query results according to column varcharstring
func (b *SelectBuilder) OrderByVarCharString(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("varcharstring", dir)
	return b
}

// GroupByVarCharString make the query group by column varcharstring
func (b *SelectBuilder) GroupByVarCharString() *SelectBuilder {
	b.params.Groups.Add("varcharstring")
	return b
}

// SelectVarCharByte adds VarCharByte to the selected column of a query
func (b *SelectBuilder) SelectVarCharByte() *SelectBuilder {
	b.columns.SelectVarCharByte = true
	return b
}

// OrderByVarCharByte set order to the query results according to column varcharbyte
func (b *SelectBuilder) OrderByVarCharByte(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("varcharbyte", dir)
	return b
}

// GroupByVarCharByte make the query group by column varcharbyte
func (b *SelectBuilder) GroupByVarCharByte() *SelectBuilder {
	b.params.Groups.Add("varcharbyte")
	return b
}

// SelectString adds String to the selected column of a query
func (b *SelectBuilder) SelectString() *SelectBuilder {
	b.columns.SelectString = true
	return b
}

// OrderByString set order to the query results according to column string
func (b *SelectBuilder) OrderByString(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("string", dir)
	return b
}

// GroupByString make the query group by column string
func (b *SelectBuilder) GroupByString() *SelectBuilder {
	b.params.Groups.Add("string")
	return b
}

// SelectBytes adds Bytes to the selected column of a query
func (b *SelectBuilder) SelectBytes() *SelectBuilder {
	b.columns.SelectBytes = true
	return b
}

// OrderByBytes set order to the query results according to column bytes
func (b *SelectBuilder) OrderByBytes(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("bytes", dir)
	return b
}

// GroupByBytes make the query group by column bytes
func (b *SelectBuilder) GroupByBytes() *SelectBuilder {
	b.params.Groups.Add("bytes")
	return b
}

// SelectBool adds Bool to the selected column of a query
func (b *SelectBuilder) SelectBool() *SelectBuilder {
	b.columns.SelectBool = true
	return b
}

// OrderByBool set order to the query results according to column bool
func (b *SelectBuilder) OrderByBool(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("bool", dir)
	return b
}

// GroupByBool make the query group by column bool
func (b *SelectBuilder) GroupByBool() *SelectBuilder {
	b.params.Groups.Add("bool")
	return b
}

// SelectPInt adds PInt to the selected column of a query
func (b *SelectBuilder) SelectPInt() *SelectBuilder {
	b.columns.SelectPInt = true
	return b
}

// OrderByPInt set order to the query results according to column pint
func (b *SelectBuilder) OrderByPInt(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pint", dir)
	return b
}

// GroupByPInt make the query group by column pint
func (b *SelectBuilder) GroupByPInt() *SelectBuilder {
	b.params.Groups.Add("pint")
	return b
}

// SelectPInt8 adds PInt8 to the selected column of a query
func (b *SelectBuilder) SelectPInt8() *SelectBuilder {
	b.columns.SelectPInt8 = true
	return b
}

// OrderByPInt8 set order to the query results according to column pint8
func (b *SelectBuilder) OrderByPInt8(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pint8", dir)
	return b
}

// GroupByPInt8 make the query group by column pint8
func (b *SelectBuilder) GroupByPInt8() *SelectBuilder {
	b.params.Groups.Add("pint8")
	return b
}

// SelectPInt16 adds PInt16 to the selected column of a query
func (b *SelectBuilder) SelectPInt16() *SelectBuilder {
	b.columns.SelectPInt16 = true
	return b
}

// OrderByPInt16 set order to the query results according to column pint16
func (b *SelectBuilder) OrderByPInt16(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pint16", dir)
	return b
}

// GroupByPInt16 make the query group by column pint16
func (b *SelectBuilder) GroupByPInt16() *SelectBuilder {
	b.params.Groups.Add("pint16")
	return b
}

// SelectPInt32 adds PInt32 to the selected column of a query
func (b *SelectBuilder) SelectPInt32() *SelectBuilder {
	b.columns.SelectPInt32 = true
	return b
}

// OrderByPInt32 set order to the query results according to column pint32
func (b *SelectBuilder) OrderByPInt32(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pint32", dir)
	return b
}

// GroupByPInt32 make the query group by column pint32
func (b *SelectBuilder) GroupByPInt32() *SelectBuilder {
	b.params.Groups.Add("pint32")
	return b
}

// SelectPInt64 adds PInt64 to the selected column of a query
func (b *SelectBuilder) SelectPInt64() *SelectBuilder {
	b.columns.SelectPInt64 = true
	return b
}

// OrderByPInt64 set order to the query results according to column pint64
func (b *SelectBuilder) OrderByPInt64(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pint64", dir)
	return b
}

// GroupByPInt64 make the query group by column pint64
func (b *SelectBuilder) GroupByPInt64() *SelectBuilder {
	b.params.Groups.Add("pint64")
	return b
}

// SelectPUInt adds PUInt to the selected column of a query
func (b *SelectBuilder) SelectPUInt() *SelectBuilder {
	b.columns.SelectPUInt = true
	return b
}

// OrderByPUInt set order to the query results according to column puint
func (b *SelectBuilder) OrderByPUInt(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("puint", dir)
	return b
}

// GroupByPUInt make the query group by column puint
func (b *SelectBuilder) GroupByPUInt() *SelectBuilder {
	b.params.Groups.Add("puint")
	return b
}

// SelectPUInt8 adds PUInt8 to the selected column of a query
func (b *SelectBuilder) SelectPUInt8() *SelectBuilder {
	b.columns.SelectPUInt8 = true
	return b
}

// OrderByPUInt8 set order to the query results according to column puint8
func (b *SelectBuilder) OrderByPUInt8(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("puint8", dir)
	return b
}

// GroupByPUInt8 make the query group by column puint8
func (b *SelectBuilder) GroupByPUInt8() *SelectBuilder {
	b.params.Groups.Add("puint8")
	return b
}

// SelectPUInt16 adds PUInt16 to the selected column of a query
func (b *SelectBuilder) SelectPUInt16() *SelectBuilder {
	b.columns.SelectPUInt16 = true
	return b
}

// OrderByPUInt16 set order to the query results according to column puint16
func (b *SelectBuilder) OrderByPUInt16(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("puint16", dir)
	return b
}

// GroupByPUInt16 make the query group by column puint16
func (b *SelectBuilder) GroupByPUInt16() *SelectBuilder {
	b.params.Groups.Add("puint16")
	return b
}

// SelectPUInt32 adds PUInt32 to the selected column of a query
func (b *SelectBuilder) SelectPUInt32() *SelectBuilder {
	b.columns.SelectPUInt32 = true
	return b
}

// OrderByPUInt32 set order to the query results according to column puint32
func (b *SelectBuilder) OrderByPUInt32(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("puint32", dir)
	return b
}

// GroupByPUInt32 make the query group by column puint32
func (b *SelectBuilder) GroupByPUInt32() *SelectBuilder {
	b.params.Groups.Add("puint32")
	return b
}

// SelectPUInt64 adds PUInt64 to the selected column of a query
func (b *SelectBuilder) SelectPUInt64() *SelectBuilder {
	b.columns.SelectPUInt64 = true
	return b
}

// OrderByPUInt64 set order to the query results according to column puint64
func (b *SelectBuilder) OrderByPUInt64(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("puint64", dir)
	return b
}

// GroupByPUInt64 make the query group by column puint64
func (b *SelectBuilder) GroupByPUInt64() *SelectBuilder {
	b.params.Groups.Add("puint64")
	return b
}

// SelectPTime adds PTime to the selected column of a query
func (b *SelectBuilder) SelectPTime() *SelectBuilder {
	b.columns.SelectPTime = true
	return b
}

// OrderByPTime set order to the query results according to column ptime
func (b *SelectBuilder) OrderByPTime(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("ptime", dir)
	return b
}

// GroupByPTime make the query group by column ptime
func (b *SelectBuilder) GroupByPTime() *SelectBuilder {
	b.params.Groups.Add("ptime")
	return b
}

// SelectPVarCharString adds PVarCharString to the selected column of a query
func (b *SelectBuilder) SelectPVarCharString() *SelectBuilder {
	b.columns.SelectPVarCharString = true
	return b
}

// OrderByPVarCharString set order to the query results according to column pvarcharstring
func (b *SelectBuilder) OrderByPVarCharString(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pvarcharstring", dir)
	return b
}

// GroupByPVarCharString make the query group by column pvarcharstring
func (b *SelectBuilder) GroupByPVarCharString() *SelectBuilder {
	b.params.Groups.Add("pvarcharstring")
	return b
}

// SelectPVarCharByte adds PVarCharByte to the selected column of a query
func (b *SelectBuilder) SelectPVarCharByte() *SelectBuilder {
	b.columns.SelectPVarCharByte = true
	return b
}

// OrderByPVarCharByte set order to the query results according to column pvarcharbyte
func (b *SelectBuilder) OrderByPVarCharByte(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pvarcharbyte", dir)
	return b
}

// GroupByPVarCharByte make the query group by column pvarcharbyte
func (b *SelectBuilder) GroupByPVarCharByte() *SelectBuilder {
	b.params.Groups.Add("pvarcharbyte")
	return b
}

// SelectPString adds PString to the selected column of a query
func (b *SelectBuilder) SelectPString() *SelectBuilder {
	b.columns.SelectPString = true
	return b
}

// OrderByPString set order to the query results according to column pstring
func (b *SelectBuilder) OrderByPString(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pstring", dir)
	return b
}

// GroupByPString make the query group by column pstring
func (b *SelectBuilder) GroupByPString() *SelectBuilder {
	b.params.Groups.Add("pstring")
	return b
}

// SelectPBytes adds PBytes to the selected column of a query
func (b *SelectBuilder) SelectPBytes() *SelectBuilder {
	b.columns.SelectPBytes = true
	return b
}

// OrderByPBytes set order to the query results according to column pbytes
func (b *SelectBuilder) OrderByPBytes(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pbytes", dir)
	return b
}

// GroupByPBytes make the query group by column pbytes
func (b *SelectBuilder) GroupByPBytes() *SelectBuilder {
	b.params.Groups.Add("pbytes")
	return b
}

// SelectPBool adds PBool to the selected column of a query
func (b *SelectBuilder) SelectPBool() *SelectBuilder {
	b.columns.SelectPBool = true
	return b
}

// OrderByPBool set order to the query results according to column pbool
func (b *SelectBuilder) OrderByPBool(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("pbool", dir)
	return b
}

// GroupByPBool make the query group by column pbool
func (b *SelectBuilder) GroupByPBool() *SelectBuilder {
	b.params.Groups.Add("pbool")
	return b
}

// SelectSelect adds Select to the selected column of a query
func (b *SelectBuilder) SelectSelect() *SelectBuilder {
	b.columns.SelectSelect = true
	return b
}

// OrderBySelect set order to the query results according to column select
func (b *SelectBuilder) OrderBySelect(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("select", dir)
	return b
}

// GroupBySelect make the query group by column select
func (b *SelectBuilder) GroupBySelect() *SelectBuilder {
	b.params.Groups.Add("select")
	return b
}

// Context sets the context for the SQL query
func (b *SelectBuilder) Context(ctx context.Context) *SelectBuilder {
	b.params.Ctx = ctx
	return b
}
