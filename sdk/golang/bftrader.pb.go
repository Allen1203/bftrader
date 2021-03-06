// Code generated by protoc-gen-go.
// source: bftrader.proto
// DO NOT EDIT!

/*
Package bftrader is a generated protocol buffer package.

It is generated from these files:
	bftrader.proto

It has these top-level messages:
	BfVoid
	BfTickData
	BfTradeData
	BfOrderData
	BfPositionData
	BfAccountData
	BfErrorData
	BfLogData
	BfContractData
	BfSendOrderReq
	BfSendOrderResp
	BfCancelOrderReq
	BfConnectReq
	BfConnectResp
	BfGetContractReq
	BfPingData
*/
package bftrader

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 方向常量
type BfDirection int32

const (
	BfDirection_DIRECTION_UNKNOWN BfDirection = 0
	BfDirection_DIRECTION_LONG    BfDirection = 1
	BfDirection_DIRECTION_SHORT   BfDirection = 2
	BfDirection_DIRECTION_NET     BfDirection = 3
)

var BfDirection_name = map[int32]string{
	0: "DIRECTION_UNKNOWN",
	1: "DIRECTION_LONG",
	2: "DIRECTION_SHORT",
	3: "DIRECTION_NET",
}
var BfDirection_value = map[string]int32{
	"DIRECTION_UNKNOWN": 0,
	"DIRECTION_LONG":    1,
	"DIRECTION_SHORT":   2,
	"DIRECTION_NET":     3,
}

func (x BfDirection) String() string {
	return proto.EnumName(BfDirection_name, int32(x))
}
func (BfDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 开平常量
type BfOffset int32

const (
	BfOffset_OFFSET_UNKNOWN        BfOffset = 0
	BfOffset_OFFSET_OPEN           BfOffset = 1
	BfOffset_OFFSET_CLOSE          BfOffset = 2
	BfOffset_OFFSET_CLOSETODAY     BfOffset = 3
	BfOffset_OFFSET_CLOSEYESTERDAY BfOffset = 4
)

var BfOffset_name = map[int32]string{
	0: "OFFSET_UNKNOWN",
	1: "OFFSET_OPEN",
	2: "OFFSET_CLOSE",
	3: "OFFSET_CLOSETODAY",
	4: "OFFSET_CLOSEYESTERDAY",
}
var BfOffset_value = map[string]int32{
	"OFFSET_UNKNOWN":        0,
	"OFFSET_OPEN":           1,
	"OFFSET_CLOSE":          2,
	"OFFSET_CLOSETODAY":     3,
	"OFFSET_CLOSEYESTERDAY": 4,
}

func (x BfOffset) String() string {
	return proto.EnumName(BfOffset_name, int32(x))
}
func (BfOffset) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 状态常量
type BfStatus int32

const (
	BfStatus_STATUS_UNKNOWN    BfStatus = 0
	BfStatus_STATUS_NOTTRADED  BfStatus = 1
	BfStatus_STATUS_PARTTRADED BfStatus = 2
	BfStatus_STATUS_ALLTRADED  BfStatus = 3
	BfStatus_STATUS_CANCELLED  BfStatus = 4
)

var BfStatus_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_NOTTRADED",
	2: "STATUS_PARTTRADED",
	3: "STATUS_ALLTRADED",
	4: "STATUS_CANCELLED",
}
var BfStatus_value = map[string]int32{
	"STATUS_UNKNOWN":    0,
	"STATUS_NOTTRADED":  1,
	"STATUS_PARTTRADED": 2,
	"STATUS_ALLTRADED":  3,
	"STATUS_CANCELLED":  4,
}

func (x BfStatus) String() string {
	return proto.EnumName(BfStatus_name, int32(x))
}
func (BfStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 合约类型常量
type BfProduct int32

const (
	BfProduct_PRODUCT_UNKNOWN BfProduct = 0
	BfProduct_PRODUCT_EQUITY  BfProduct = 1
	BfProduct_PRODUCT_FUTURES BfProduct = 2
)

var BfProduct_name = map[int32]string{
	0: "PRODUCT_UNKNOWN",
	1: "PRODUCT_EQUITY",
	2: "PRODUCT_FUTURES",
}
var BfProduct_value = map[string]int32{
	"PRODUCT_UNKNOWN": 0,
	"PRODUCT_EQUITY":  1,
	"PRODUCT_FUTURES": 2,
}

func (x BfProduct) String() string {
	return proto.EnumName(BfProduct_name, int32(x))
}
func (BfProduct) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 价格类型常量
type BfPriceType int32

const (
	BfPriceType_PRICETYPE_UNKONWN     BfPriceType = 0
	BfPriceType_PRICETYPE_LIMITPRICE  BfPriceType = 1
	BfPriceType_PRICETYPE_MARKETPRICE BfPriceType = 2
)

var BfPriceType_name = map[int32]string{
	0: "PRICETYPE_UNKONWN",
	1: "PRICETYPE_LIMITPRICE",
	2: "PRICETYPE_MARKETPRICE",
}
var BfPriceType_value = map[string]int32{
	"PRICETYPE_UNKONWN":     0,
	"PRICETYPE_LIMITPRICE":  1,
	"PRICETYPE_MARKETPRICE": 2,
}

func (x BfPriceType) String() string {
	return proto.EnumName(BfPriceType_name, int32(x))
}
func (BfPriceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 空参数
type BfVoid struct {
}

func (m *BfVoid) Reset()                    { *m = BfVoid{} }
func (m *BfVoid) String() string            { return proto.CompactTextString(m) }
func (*BfVoid) ProtoMessage()               {}
func (*BfVoid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Tick行情数据类
type BfTickData struct {
	// 代码相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 成交数据
	ActionDate   string  `protobuf:"bytes,3,opt,name=actionDate" json:"actionDate,omitempty"`
	TickTime     string  `protobuf:"bytes,4,opt,name=tickTime" json:"tickTime,omitempty"`
	LastPrice    float64 `protobuf:"fixed64,5,opt,name=lastPrice" json:"lastPrice,omitempty"`
	Volume       int32   `protobuf:"varint,6,opt,name=volume" json:"volume,omitempty"`
	OpenInterest float64 `protobuf:"fixed64,7,opt,name=openInterest" json:"openInterest,omitempty"`
	LastVolume   int32   `protobuf:"varint,8,opt,name=lastVolume" json:"lastVolume,omitempty"`
	// 常规行情
	OpenPrice     float64 `protobuf:"fixed64,9,opt,name=openPrice" json:"openPrice,omitempty"`
	HighPrice     float64 `protobuf:"fixed64,10,opt,name=highPrice" json:"highPrice,omitempty"`
	LowPrice      float64 `protobuf:"fixed64,11,opt,name=lowPrice" json:"lowPrice,omitempty"`
	PreClosePrice float64 `protobuf:"fixed64,12,opt,name=preClosePrice" json:"preClosePrice,omitempty"`
	UpperLimit    float64 `protobuf:"fixed64,13,opt,name=upperLimit" json:"upperLimit,omitempty"`
	LowerLimit    float64 `protobuf:"fixed64,14,opt,name=lowerLimit" json:"lowerLimit,omitempty"`
	// x档行情
	BidPrice1  float64 `protobuf:"fixed64,15,opt,name=bidPrice1" json:"bidPrice1,omitempty"`
	AskPrice1  float64 `protobuf:"fixed64,16,opt,name=askPrice1" json:"askPrice1,omitempty"`
	BidVolume1 int32   `protobuf:"varint,17,opt,name=bidVolume1" json:"bidVolume1,omitempty"`
	AskVolume1 int32   `protobuf:"varint,18,opt,name=askVolume1" json:"askVolume1,omitempty"`
}

func (m *BfTickData) Reset()                    { *m = BfTickData{} }
func (m *BfTickData) String() string            { return proto.CompactTextString(m) }
func (*BfTickData) ProtoMessage()               {}
func (*BfTickData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 成交数据类
type BfTradeData struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	TradeId  string `protobuf:"bytes,3,opt,name=tradeId" json:"tradeId,omitempty"`
	// 对于ctp/lts是frontid.sessioni.orderref
	// ctp/lts的trade里面没有frontid+sessionid
	// 内部通过sysOrderId是做了映射trade到Order
	BfOrderId string `protobuf:"bytes,4,opt,name=bfOrderId" json:"bfOrderId,omitempty"`
	// 成交相关
	Direction BfDirection `protobuf:"varint,5,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset    BfOffset    `protobuf:"varint,6,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
	Price     float64     `protobuf:"fixed64,7,opt,name=price" json:"price,omitempty"`
	Volume    int32       `protobuf:"varint,8,opt,name=volume" json:"volume,omitempty"`
	TradeDate string      `protobuf:"bytes,9,opt,name=tradeDate" json:"tradeDate,omitempty"`
	TradeTime string      `protobuf:"bytes,10,opt,name=tradeTime" json:"tradeTime,omitempty"`
}

func (m *BfTradeData) Reset()                    { *m = BfTradeData{} }
func (m *BfTradeData) String() string            { return proto.CompactTextString(m) }
func (*BfTradeData) ProtoMessage()               {}
func (*BfTradeData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 订单数据类
type BfOrderData struct {
	// 代码编号相关
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange  string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	BfOrderId string `protobuf:"bytes,3,opt,name=bfOrderId" json:"bfOrderId,omitempty"`
	// 报单相关
	Direction    BfDirection `protobuf:"varint,4,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset       BfOffset    `protobuf:"varint,5,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
	Price        float64     `protobuf:"fixed64,6,opt,name=price" json:"price,omitempty"`
	TotalVolume  int32       `protobuf:"varint,7,opt,name=totalVolume" json:"totalVolume,omitempty"`
	TradedVolume int32       `protobuf:"varint,8,opt,name=tradedVolume" json:"tradedVolume,omitempty"`
	Status       BfStatus    `protobuf:"varint,9,opt,name=status,enum=bftrader.BfStatus" json:"status,omitempty"`
	InsertDate   string      `protobuf:"bytes,10,opt,name=insertDate" json:"insertDate,omitempty"`
	InsertTime   string      `protobuf:"bytes,11,opt,name=insertTime" json:"insertTime,omitempty"`
	CancelTime   string      `protobuf:"bytes,12,opt,name=cancelTime" json:"cancelTime,omitempty"`
}

func (m *BfOrderData) Reset()                    { *m = BfOrderData{} }
func (m *BfOrderData) String() string            { return proto.CompactTextString(m) }
func (*BfOrderData) ProtoMessage()               {}
func (*BfOrderData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 持仓数据类
type BfPositionData struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 持仓相关
	Direction  BfDirection `protobuf:"varint,3,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Position   int32       `protobuf:"varint,4,opt,name=position" json:"position,omitempty"`
	Frozen     int32       `protobuf:"varint,5,opt,name=frozen" json:"frozen,omitempty"`
	Price      float64     `protobuf:"fixed64,6,opt,name=price" json:"price,omitempty"`
	YdPosition int32       `protobuf:"varint,7,opt,name=ydPosition" json:"ydPosition,omitempty"`
}

func (m *BfPositionData) Reset()                    { *m = BfPositionData{} }
func (m *BfPositionData) String() string            { return proto.CompactTextString(m) }
func (*BfPositionData) ProtoMessage()               {}
func (*BfPositionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 账户数据类
type BfAccountData struct {
	// 账号代码相关
	AccountId string `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	// 数值相关
	PreBalance     float64 `protobuf:"fixed64,2,opt,name=preBalance" json:"preBalance,omitempty"`
	Balance        float64 `protobuf:"fixed64,3,opt,name=balance" json:"balance,omitempty"`
	Available      float64 `protobuf:"fixed64,4,opt,name=available" json:"available,omitempty"`
	Commission     float64 `protobuf:"fixed64,5,opt,name=commission" json:"commission,omitempty"`
	FrozenMargin   float64 `protobuf:"fixed64,6,opt,name=frozenMargin" json:"frozenMargin,omitempty"`
	CloseProfit    float64 `protobuf:"fixed64,7,opt,name=closeProfit" json:"closeProfit,omitempty"`
	PositionProfit float64 `protobuf:"fixed64,8,opt,name=positionProfit" json:"positionProfit,omitempty"`
}

func (m *BfAccountData) Reset()                    { *m = BfAccountData{} }
func (m *BfAccountData) String() string            { return proto.CompactTextString(m) }
func (*BfAccountData) ProtoMessage()               {}
func (*BfAccountData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 错误数据类
type BfErrorData struct {
	ErrorId        int32  `protobuf:"varint,1,opt,name=errorId" json:"errorId,omitempty"`
	ErrorMsg       string `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
	AdditionalInfo string `protobuf:"bytes,3,opt,name=additionalInfo" json:"additionalInfo,omitempty"`
}

func (m *BfErrorData) Reset()                    { *m = BfErrorData{} }
func (m *BfErrorData) String() string            { return proto.CompactTextString(m) }
func (*BfErrorData) ProtoMessage()               {}
func (*BfErrorData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 日志数据类
type BfLogData struct {
	When    string `protobuf:"bytes,1,opt,name=when" json:"when,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *BfLogData) Reset()                    { *m = BfLogData{} }
func (m *BfLogData) String() string            { return proto.CompactTextString(m) }
func (*BfLogData) ProtoMessage()               {}
func (*BfLogData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 合约详细信息类
type BfContractData struct {
	// 代码编号相关
	Symbol         string    `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange       string    `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Name           string    `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ProductClass   BfProduct `protobuf:"varint,4,opt,name=productClass,enum=bftrader.BfProduct" json:"productClass,omitempty"`
	VolumeMultiple int32     `protobuf:"varint,5,opt,name=volumeMultiple" json:"volumeMultiple,omitempty"`
	PriceTick      float64   `protobuf:"fixed64,6,opt,name=priceTick" json:"priceTick,omitempty"`
	MaxLimit       int32     `protobuf:"varint,7,opt,name=maxLimit" json:"maxLimit,omitempty"`
	MinLimit       int32     `protobuf:"varint,8,opt,name=minLimit" json:"minLimit,omitempty"`
	MaxMarket      int32     `protobuf:"varint,9,opt,name=maxMarket" json:"maxMarket,omitempty"`
	MinMartet      int32     `protobuf:"varint,10,opt,name=minMartet" json:"minMartet,omitempty"`
}

func (m *BfContractData) Reset()                    { *m = BfContractData{} }
func (m *BfContractData) String() string            { return proto.CompactTextString(m) }
func (*BfContractData) ProtoMessage()               {}
func (*BfContractData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 发单时传入的对象类
type BfSendOrderReq struct {
	// 代码编号相关
	Symbol    string      `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange  string      `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	Price     float64     `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
	Volume    int32       `protobuf:"varint,4,opt,name=volume" json:"volume,omitempty"`
	PriceType BfPriceType `protobuf:"varint,5,opt,name=priceType,enum=bftrader.BfPriceType" json:"priceType,omitempty"`
	Direction BfDirection `protobuf:"varint,6,opt,name=direction,enum=bftrader.BfDirection" json:"direction,omitempty"`
	Offset    BfOffset    `protobuf:"varint,7,opt,name=offset,enum=bftrader.BfOffset" json:"offset,omitempty"`
}

func (m *BfSendOrderReq) Reset()                    { *m = BfSendOrderReq{} }
func (m *BfSendOrderReq) String() string            { return proto.CompactTextString(m) }
func (*BfSendOrderReq) ProtoMessage()               {}
func (*BfSendOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 发单本地返回的对象类
type BfSendOrderResp struct {
	BfOrderId string `protobuf:"bytes,1,opt,name=bfOrderId" json:"bfOrderId,omitempty"`
}

func (m *BfSendOrderResp) Reset()                    { *m = BfSendOrderResp{} }
func (m *BfSendOrderResp) String() string            { return proto.CompactTextString(m) }
func (*BfSendOrderResp) ProtoMessage()               {}
func (*BfSendOrderResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// 撤单时传入的对象类
type BfCancelOrderReq struct {
	// 代码编号相关
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange  string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	BfOrderId string `protobuf:"bytes,3,opt,name=bfOrderId" json:"bfOrderId,omitempty"`
}

func (m *BfCancelOrderReq) Reset()                    { *m = BfCancelOrderReq{} }
func (m *BfCancelOrderReq) String() string            { return proto.CompactTextString(m) }
func (*BfCancelOrderReq) ProtoMessage()               {}
func (*BfCancelOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// 连接时传入的对象类
type BfConnectReq struct {
	// 唯一英文代号
	ProxyId string `protobuf:"bytes,1,opt,name=proxyId" json:"proxyId,omitempty"`
	// 服务地址和端口
	ProxyIp   string `protobuf:"bytes,2,opt,name=proxyIp" json:"proxyIp,omitempty"`
	ProxyPort int32  `protobuf:"varint,3,opt,name=proxyPort" json:"proxyPort,omitempty"`
	// 角色类别，可多选
	// 注意：OnTradeWillBegin OnGotContracts OnPing 都必须实现
	TickHandler  bool `protobuf:"varint,4,opt,name=tickHandler" json:"tickHandler,omitempty"`
	TradeHandler bool `protobuf:"varint,5,opt,name=tradeHandler" json:"tradeHandler,omitempty"`
	LogHandler   bool `protobuf:"varint,6,opt,name=logHandler" json:"logHandler,omitempty"`
	// tickHandler相关
	Symbol   string `protobuf:"bytes,7,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,8,opt,name=exchange" json:"exchange,omitempty"`
}

func (m *BfConnectReq) Reset()                    { *m = BfConnectReq{} }
func (m *BfConnectReq) String() string            { return proto.CompactTextString(m) }
func (*BfConnectReq) ProtoMessage()               {}
func (*BfConnectReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// 连接时返回的对象类
type BfConnectResp struct {
	ErrorCode int32  `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
	ErrorMsg  string `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *BfConnectResp) Reset()                    { *m = BfConnectResp{} }
func (m *BfConnectResp) String() string            { return proto.CompactTextString(m) }
func (*BfConnectResp) ProtoMessage()               {}
func (*BfConnectResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// 获取合约信息传入的对象类
type BfGetContractReq struct {
	// 代码编号相关
	Symbol   string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	// 获取全部合约时候可以用index来枚举，从1开始，返回空结束
	Index int32 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	// 过滤条件，指定是否要求已经订阅
	Subscribled bool `protobuf:"varint,4,opt,name=subscribled" json:"subscribled,omitempty"`
}

func (m *BfGetContractReq) Reset()                    { *m = BfGetContractReq{} }
func (m *BfGetContractReq) String() string            { return proto.CompactTextString(m) }
func (*BfGetContractReq) ProtoMessage()               {}
func (*BfGetContractReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// Ping/Pong检测传入的对象类
// 同ctp一样，5秒一次
type BfPingData struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *BfPingData) Reset()                    { *m = BfPingData{} }
func (m *BfPingData) String() string            { return proto.CompactTextString(m) }
func (*BfPingData) ProtoMessage()               {}
func (*BfPingData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*BfVoid)(nil), "bftrader.BfVoid")
	proto.RegisterType((*BfTickData)(nil), "bftrader.BfTickData")
	proto.RegisterType((*BfTradeData)(nil), "bftrader.BfTradeData")
	proto.RegisterType((*BfOrderData)(nil), "bftrader.BfOrderData")
	proto.RegisterType((*BfPositionData)(nil), "bftrader.BfPositionData")
	proto.RegisterType((*BfAccountData)(nil), "bftrader.BfAccountData")
	proto.RegisterType((*BfErrorData)(nil), "bftrader.BfErrorData")
	proto.RegisterType((*BfLogData)(nil), "bftrader.BfLogData")
	proto.RegisterType((*BfContractData)(nil), "bftrader.BfContractData")
	proto.RegisterType((*BfSendOrderReq)(nil), "bftrader.BfSendOrderReq")
	proto.RegisterType((*BfSendOrderResp)(nil), "bftrader.BfSendOrderResp")
	proto.RegisterType((*BfCancelOrderReq)(nil), "bftrader.BfCancelOrderReq")
	proto.RegisterType((*BfConnectReq)(nil), "bftrader.BfConnectReq")
	proto.RegisterType((*BfConnectResp)(nil), "bftrader.BfConnectResp")
	proto.RegisterType((*BfGetContractReq)(nil), "bftrader.BfGetContractReq")
	proto.RegisterType((*BfPingData)(nil), "bftrader.BfPingData")
	proto.RegisterEnum("bftrader.BfDirection", BfDirection_name, BfDirection_value)
	proto.RegisterEnum("bftrader.BfOffset", BfOffset_name, BfOffset_value)
	proto.RegisterEnum("bftrader.BfStatus", BfStatus_name, BfStatus_value)
	proto.RegisterEnum("bftrader.BfProduct", BfProduct_name, BfProduct_value)
	proto.RegisterEnum("bftrader.BfPriceType", BfPriceType_name, BfPriceType_value)
}

var fileDescriptor0 = []byte{
	// 1366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xae, 0x7e, 0x2d, 0xad, 0x65, 0x59, 0x61, 0x92, 0x42, 0x2d, 0x8a, 0x22, 0x20, 0x8a, 0xa0,
	0xd0, 0x21, 0x45, 0x92, 0x43, 0xd1, 0xa3, 0x7e, 0xe8, 0x44, 0x88, 0x2c, 0xaa, 0x14, 0x95, 0xc2,
	0xa7, 0x96, 0x22, 0x97, 0x32, 0x61, 0x8a, 0x54, 0x49, 0x2a, 0xb6, 0x7b, 0xe8, 0xb9, 0xf7, 0xbe,
	0x44, 0x8f, 0x7d, 0x88, 0xde, 0xfa, 0x18, 0x7d, 0x8b, 0x9e, 0xba, 0x33, 0xbb, 0x4b, 0x2e, 0x8d,
	0xc4, 0x4d, 0x9c, 0x9b, 0xe7, 0xfb, 0x86, 0xb3, 0xbb, 0xdf, 0x7c, 0x3b, 0x2b, 0x93, 0xee, 0xda,
	0xcf, 0x12, 0xc7, 0xa3, 0xc9, 0x93, 0x5d, 0x12, 0x67, 0xb1, 0xd6, 0x92, 0xb1, 0xde, 0x22, 0xcd,
	0x91, 0xff, 0x3a, 0x0e, 0x3c, 0xfd, 0xcf, 0x3a, 0x21, 0x23, 0xdf, 0x0e, 0xdc, 0x8b, 0x89, 0x93,
	0x39, 0xda, 0xa7, 0xa4, 0x99, 0x5e, 0x6f, 0xd7, 0x71, 0xd8, 0xaf, 0x3c, 0xaa, 0x7c, 0xdd, 0xb6,
	0x44, 0xa4, 0x7d, 0x4e, 0x5a, 0xf4, 0xca, 0x3d, 0x77, 0xa2, 0x0d, 0xed, 0x57, 0x91, 0xc9, 0x63,
	0xed, 0x4b, 0x42, 0x1c, 0x37, 0x0b, 0xe2, 0x88, 0x55, 0xa0, 0xfd, 0x1a, 0xb2, 0x0a, 0x02, 0xdf,
	0x66, 0xac, 0xbe, 0x1d, 0x6c, 0x69, 0xbf, 0xce, 0xbf, 0x95, 0xb1, 0xf6, 0x05, 0x69, 0x87, 0x4e,
	0x9a, 0x2d, 0x92, 0xc0, 0xa5, 0xfd, 0x06, 0x23, 0x2b, 0x56, 0x01, 0xc0, 0x6e, 0xde, 0xc4, 0xe1,
	0x9e, 0x7d, 0xd7, 0x64, 0x54, 0xc3, 0x12, 0x91, 0xa6, 0x93, 0x4e, 0xbc, 0xa3, 0xd1, 0x34, 0xca,
	0x68, 0x42, 0xd3, 0xac, 0x7f, 0x80, 0x1f, 0x96, 0x30, 0xd8, 0x15, 0x14, 0x7a, 0xcd, 0xbf, 0x6f,
	0xe1, 0xf7, 0x0a, 0x02, 0x2b, 0x43, 0x3e, 0x5f, 0xb9, 0xcd, 0x57, 0xce, 0x01, 0x60, 0xcf, 0x83,
	0xcd, 0x39, 0x67, 0x09, 0x67, 0x73, 0x00, 0x4e, 0x14, 0xc6, 0x97, 0x9c, 0x3c, 0x44, 0x32, 0x8f,
	0xb5, 0xaf, 0xc8, 0xd1, 0x2e, 0xa1, 0xe3, 0x30, 0x4e, 0x29, 0x4f, 0xe8, 0x60, 0x42, 0x19, 0x84,
	0xdd, 0xed, 0x77, 0x3b, 0x9a, 0xcc, 0x82, 0x6d, 0x90, 0xf5, 0x8f, 0x30, 0x45, 0x41, 0x70, 0xf7,
	0xf1, 0xa5, 0xe4, 0xbb, 0x9c, 0x2f, 0x10, 0xd8, 0xdf, 0x3a, 0xf0, 0xb0, 0xd6, 0xd3, 0xfe, 0x31,
	0xdf, 0x5f, 0x0e, 0x00, 0xeb, 0xa4, 0x17, 0x82, 0xed, 0x71, 0x36, 0x07, 0xa0, 0x36, 0x4b, 0xe5,
	0x32, 0x3c, 0xed, 0xdf, 0xe3, 0xca, 0x14, 0x08, 0xf6, 0x33, 0xbd, 0x90, 0xbc, 0xc6, 0xf9, 0x02,
	0xd1, 0xff, 0xaa, 0x92, 0x43, 0x66, 0x19, 0x70, 0xd2, 0x9d, 0x3d, 0xd3, 0x27, 0x07, 0x68, 0xc5,
	0xa9, 0x27, 0x0c, 0x23, 0x43, 0x3c, 0x99, 0x6f, 0x26, 0xcc, 0xa5, 0x8c, 0xe3, 0x76, 0x29, 0x00,
	0xed, 0x39, 0x69, 0x7b, 0x41, 0x42, 0xd1, 0x5c, 0xe8, 0x97, 0xee, 0xb3, 0x87, 0x4f, 0x72, 0x9b,
	0x8f, 0xfc, 0x89, 0x24, 0xad, 0x22, 0x4f, 0x1b, 0x90, 0x66, 0xec, 0xfb, 0x29, 0xcd, 0xd0, 0x46,
	0xdd, 0x67, 0x9a, 0xfa, 0x85, 0x89, 0x8c, 0x25, 0x32, 0xb4, 0x07, 0xa4, 0xb1, 0xc3, 0xb6, 0x71,
	0x4f, 0xf1, 0x40, 0x31, 0x62, 0xab, 0x64, 0x44, 0xb6, 0xd9, 0x4c, 0xe8, 0xc0, 0x4d, 0xc4, 0x36,
	0x9b, 0x03, 0x39, 0x8b, 0xce, 0x27, 0x0a, 0x0b, 0x80, 0xfe, 0x47, 0x0d, 0x64, 0xc4, 0x83, 0xdd,
	0x59, 0xc6, 0x92, 0x58, 0xb5, 0x5b, 0xc5, 0xaa, 0x7f, 0xb0, 0x58, 0x8d, 0xf7, 0x17, 0xab, 0xa9,
	0x8a, 0xf5, 0x88, 0x1c, 0x66, 0x71, 0xe6, 0x84, 0xe2, 0xea, 0x1d, 0xa0, 0x62, 0x2a, 0x04, 0xf7,
	0x17, 0x4b, 0x7a, 0xa5, 0xdb, 0x59, 0xc2, 0x60, 0x1f, 0x69, 0xe6, 0x64, 0xfb, 0x14, 0x75, 0xbd,
	0xb1, 0x8f, 0x25, 0x32, 0x96, 0xc8, 0x00, 0xc7, 0x06, 0x51, 0x4a, 0x93, 0x0c, 0xfb, 0xc0, 0x95,
	0x56, 0x90, 0x82, 0xc7, 0x4e, 0x1c, 0xaa, 0x3c, 0x4e, 0x21, 0xc6, 0xbb, 0x4e, 0xe4, 0xd2, 0x10,
	0xf9, 0x0e, 0xe7, 0x0b, 0x44, 0xff, 0xa7, 0x42, 0xba, 0x23, 0x7f, 0x11, 0xa7, 0x81, 0x18, 0x6a,
	0x77, 0xeb, 0x56, 0xa9, 0x1f, 0xb5, 0xf7, 0xec, 0x07, 0x2b, 0xb8, 0x13, 0x0b, 0x63, 0x0f, 0x1b,
	0x56, 0x1e, 0xc3, 0x26, 0xfc, 0x24, 0xfe, 0x85, 0xf2, 0xab, 0xc0, 0x6c, 0xc9, 0xa3, 0x77, 0xf4,
	0x85, 0x9d, 0xf2, 0xda, 0x93, 0x87, 0x10, 0x6d, 0x51, 0x10, 0xfd, 0xf7, 0x2a, 0x39, 0x1a, 0xf9,
	0x43, 0xd7, 0x8d, 0xf7, 0x51, 0x86, 0x87, 0x84, 0x39, 0xc2, 0x43, 0x66, 0x2f, 0x7e, 0xce, 0x02,
	0x80, 0x7a, 0x6c, 0xa8, 0x8d, 0x9c, 0x10, 0x84, 0xc2, 0xc3, 0xb2, 0x19, 0x55, 0x20, 0x70, 0xc7,
	0xd7, 0x82, 0xac, 0x21, 0x29, 0x43, 0xac, 0xfb, 0xc6, 0x09, 0x42, 0x67, 0x1d, 0xf2, 0x27, 0x01,
	0xe6, 0x93, 0x04, 0xb0, 0x1b, 0xf1, 0x76, 0x1b, 0xa4, 0xa9, 0xbc, 0xe4, 0xac, 0x6e, 0x81, 0x80,
	0x7b, 0xf8, 0x39, 0x4f, 0x9d, 0x64, 0x13, 0x44, 0xe2, 0x90, 0x25, 0x0c, 0x3c, 0xe8, 0xf2, 0x69,
	0x1b, 0xfb, 0x81, 0x7c, 0x20, 0x54, 0x48, 0x7b, 0x4c, 0xba, 0x52, 0x47, 0x91, 0xd4, 0xc2, 0xa4,
	0x1b, 0xa8, 0x7e, 0x01, 0xb7, 0xd4, 0x48, 0x92, 0x98, 0xdf, 0x52, 0x76, 0x28, 0x0a, 0x81, 0x10,
	0xa4, 0x61, 0xc9, 0x10, 0x3b, 0x0f, 0x7f, 0x9e, 0xa6, 0x9b, 0xbc, 0xf3, 0x22, 0x86, 0xc5, 0x1c,
	0xcf, 0xc3, 0xb2, 0x4e, 0x38, 0x8d, 0xfc, 0x58, 0x5c, 0xd6, 0x1b, 0xa8, 0xfe, 0x1d, 0x69, 0x8f,
	0xfc, 0x59, 0xbc, 0xc1, 0xa5, 0x34, 0x52, 0xbf, 0x3c, 0x67, 0xbd, 0xe5, 0xc2, 0xe3, 0xdf, 0xb0,
	0xfc, 0x96, 0xa6, 0xa9, 0x93, 0xbb, 0x4b, 0x86, 0xfa, 0xdf, 0x55, 0xf0, 0xe8, 0x38, 0x8e, 0x98,
	0x9d, 0xdc, 0xec, 0xce, 0x1e, 0x65, 0x8b, 0x46, 0xce, 0x56, 0x3e, 0xe3, 0xf8, 0xb7, 0xf6, 0x2d,
	0xe9, 0xb0, 0x1f, 0x10, 0xde, 0xde, 0xcd, 0xc6, 0xec, 0x01, 0x4d, 0xc5, 0x28, 0xb9, 0xaf, 0x5a,
	0x77, 0xc1, 0x79, 0xab, 0x94, 0x08, 0xc7, 0xe6, 0x83, 0xf2, 0x74, 0x1f, 0x66, 0xc1, 0x2e, 0xa4,
	0xc2, 0xa7, 0x37, 0x50, 0xf0, 0x03, 0x5a, 0x14, 0x7e, 0x86, 0x88, 0x76, 0x16, 0x00, 0x6c, 0x77,
	0xeb, 0x5c, 0xf1, 0x97, 0x90, 0xbb, 0x36, 0x8f, 0x91, 0x0b, 0x22, 0xce, 0xb5, 0x04, 0x27, 0x62,
	0xa8, 0xca, 0xf2, 0x98, 0x21, 0x2e, 0xd8, 0x30, 0x6b, 0x23, 0x59, 0x00, 0xc8, 0x06, 0x60, 0x97,
	0x8c, 0xb1, 0x44, 0xb0, 0x12, 0xd0, 0x7f, 0x43, 0x35, 0x97, 0x34, 0xf2, 0x70, 0x98, 0x5a, 0xf4,
	0xe7, 0x3b, 0xa9, 0x99, 0x5f, 0xc4, 0xda, 0xdb, 0x5f, 0x93, 0x7a, 0xe9, 0x35, 0x79, 0x2e, 0x65,
	0xb8, 0xde, 0xd1, 0xb7, 0x3d, 0x6e, 0x0b, 0x49, 0x5a, 0x45, 0x5e, 0x79, 0xa8, 0x34, 0x3f, 0x78,
	0xc8, 0x1f, 0xfc, 0xdf, 0x90, 0xd7, 0xbf, 0x21, 0xc7, 0x25, 0x25, 0xd2, 0x5d, 0xf9, 0xd9, 0xa9,
	0xdc, 0x78, 0x76, 0x74, 0x8f, 0xf4, 0x98, 0x11, 0x71, 0x7a, 0x7e, 0x94, 0x78, 0xb7, 0x3e, 0x6e,
	0xfa, 0xbf, 0x15, 0xd2, 0x41, 0xbf, 0x47, 0xec, 0x4c, 0xb0, 0x04, 0xbb, 0x1a, 0xcc, 0x7c, 0x57,
	0xd7, 0xf9, 0x96, 0x64, 0x58, 0x30, 0x3b, 0x79, 0x69, 0x44, 0xc8, 0x8d, 0xc7, 0xfe, 0x5c, 0xc4,
	0x49, 0x86, 0x4b, 0x34, 0xac, 0x02, 0xc0, 0x87, 0x8c, 0x19, 0xf0, 0xa5, 0x13, 0x79, 0x21, 0x4d,
	0xb0, 0x59, 0x2d, 0x4b, 0x85, 0xf2, 0x87, 0x4c, 0xa6, 0x34, 0x30, 0xa5, 0x84, 0xf1, 0x9f, 0x72,
	0x1b, 0x99, 0xd1, 0xc4, 0x0c, 0x05, 0x51, 0xa4, 0x39, 0x78, 0xa7, 0x34, 0xad, 0xb2, 0x34, 0xfa,
	0x14, 0x26, 0x75, 0x7e, 0x76, 0xde, 0x11, 0x1c, 0x36, 0xe3, 0xd8, 0xa3, 0x62, 0x30, 0x15, 0xc0,
	0x6d, 0xa3, 0x49, 0xff, 0x15, 0xba, 0xf5, 0x82, 0x66, 0x72, 0x72, 0x7c, 0x84, 0xd5, 0x83, 0xc8,
	0xa3, 0x57, 0x42, 0x46, 0x1e, 0x80, 0x84, 0xe9, 0x7e, 0x9d, 0xba, 0x49, 0xc0, 0x26, 0xbb, 0x27,
	0x25, 0x54, 0x20, 0xfd, 0x31, 0xfc, 0xff, 0xb1, 0x08, 0xa2, 0x8d, 0x1c, 0xaf, 0x72, 0xbe, 0x55,
	0x4a, 0xf3, 0x6d, 0xf0, 0x13, 0xcc, 0xe1, 0xdc, 0xcc, 0xda, 0x43, 0x72, 0x6f, 0x32, 0xb5, 0x8c,
	0xb1, 0x3d, 0x35, 0xe7, 0x3f, 0xae, 0xe6, 0xaf, 0xe6, 0xe6, 0x0f, 0xf3, 0xde, 0x27, 0x6c, 0x7c,
	0x75, 0x0b, 0x78, 0x66, 0xce, 0x5f, 0xf4, 0x2a, 0xda, 0x7d, 0x72, 0x5c, 0x60, 0xcb, 0x97, 0xa6,
	0x65, 0xf7, 0xaa, 0xda, 0x3d, 0x72, 0x54, 0x80, 0x73, 0xc3, 0xee, 0xd5, 0x06, 0x29, 0x69, 0x49,
	0xf3, 0x43, 0x1d, 0xf3, 0xe4, 0x64, 0x69, 0xd8, 0x4a, 0xed, 0x63, 0x72, 0x28, 0x30, 0x73, 0x61,
	0xcc, 0x59, 0xe1, 0x1e, 0xe9, 0x08, 0x60, 0x3c, 0x33, 0x97, 0x06, 0xab, 0xca, 0x76, 0xa5, 0x22,
	0xb6, 0x39, 0x19, 0x9e, 0xf5, 0x6a, 0xda, 0x67, 0xe4, 0xa1, 0x0a, 0x9f, 0x19, 0x4b, 0xdb, 0xb0,
	0x80, 0xaa, 0x0f, 0x2e, 0x61, 0x51, 0xfe, 0x73, 0x06, 0x16, 0x5d, 0xda, 0x43, 0x7b, 0xb5, 0x54,
	0x16, 0x7d, 0x40, 0x7a, 0x02, 0x9b, 0x9b, 0xb6, 0x6d, 0x0d, 0x27, 0xc6, 0x84, 0xad, 0xcc, 0xd6,
	0x11, 0xe8, 0x62, 0x68, 0x49, 0xb8, 0xaa, 0x24, 0x0f, 0x67, 0x33, 0x81, 0xd6, 0x14, 0x74, 0x3c,
	0x9c, 0x8f, 0x8d, 0xd9, 0x8c, 0xa1, 0xf5, 0xc1, 0x14, 0x9e, 0x1a, 0x31, 0xb6, 0x41, 0xa2, 0x85,
	0x65, 0x4e, 0x56, 0x63, 0xbb, 0xac, 0xa5, 0x04, 0x8d, 0xef, 0x57, 0x53, 0xfb, 0x8c, 0x6b, 0x29,
	0xb1, 0x93, 0x95, 0xbd, 0xb2, 0x8c, 0x65, 0xaf, 0x3a, 0x38, 0x83, 0xd6, 0xe4, 0xc3, 0x09, 0x36,
	0xb7, 0xb0, 0xa6, 0x63, 0xc3, 0x3e, 0x5b, 0x18, 0x50, 0xce, 0x9c, 0x63, 0xb9, 0x3e, 0x79, 0x50,
	0xc0, 0xb3, 0xe9, 0xe9, 0xd4, 0xc6, 0x90, 0x15, 0x65, 0xf2, 0x14, 0xcc, 0xe9, 0xd0, 0x7a, 0x65,
	0x08, 0xaa, 0xba, 0x6e, 0xe2, 0x7f, 0xae, 0xcf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xb4,
	0x89, 0x8e, 0xcb, 0x0e, 0x00, 0x00,
}
