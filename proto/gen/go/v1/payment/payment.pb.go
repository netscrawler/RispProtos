// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: v1/payment/payment.proto

package paymentv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Статусы платежа
type PaymentStatus int32

const (
	PaymentStatus_PAYMENT_STATUS_UNSPECIFIED PaymentStatus = 0
	PaymentStatus_PAYMENT_STATUS_PENDING     PaymentStatus = 1 // Ожидает обработки
	PaymentStatus_PAYMENT_STATUS_SUCCEEDED   PaymentStatus = 2 // Успешно завершен
	PaymentStatus_PAYMENT_STATUS_FAILED      PaymentStatus = 3 // Ошибка оплаты
	PaymentStatus_PAYMENT_STATUS_REFUNDED    PaymentStatus = 4 // Возврат средств
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PAYMENT_STATUS_UNSPECIFIED",
		1: "PAYMENT_STATUS_PENDING",
		2: "PAYMENT_STATUS_SUCCEEDED",
		3: "PAYMENT_STATUS_FAILED",
		4: "PAYMENT_STATUS_REFUNDED",
	}
	PaymentStatus_value = map[string]int32{
		"PAYMENT_STATUS_UNSPECIFIED": 0,
		"PAYMENT_STATUS_PENDING":     1,
		"PAYMENT_STATUS_SUCCEEDED":   2,
		"PAYMENT_STATUS_FAILED":      3,
		"PAYMENT_STATUS_REFUNDED":    4,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_payment_payment_proto_enumTypes[0].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_v1_payment_payment_proto_enumTypes[0]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{0}
}

// Типы платежных методов
type PaymentMethodType int32

const (
	PaymentMethodType_PAYMENT_METHOD_UNSPECIFIED PaymentMethodType = 0
	PaymentMethodType_PAYMENT_METHOD_CARD        PaymentMethodType = 1 // Банковская карта
	PaymentMethodType_PAYMENT_METHOD_CASH        PaymentMethodType = 2 // Наличные
	PaymentMethodType_PAYMENT_METHOD_APPLE_PAY   PaymentMethodType = 3
	PaymentMethodType_PAYMENT_METHOD_GOOGLE_PAY  PaymentMethodType = 4
)

// Enum value maps for PaymentMethodType.
var (
	PaymentMethodType_name = map[int32]string{
		0: "PAYMENT_METHOD_UNSPECIFIED",
		1: "PAYMENT_METHOD_CARD",
		2: "PAYMENT_METHOD_CASH",
		3: "PAYMENT_METHOD_APPLE_PAY",
		4: "PAYMENT_METHOD_GOOGLE_PAY",
	}
	PaymentMethodType_value = map[string]int32{
		"PAYMENT_METHOD_UNSPECIFIED": 0,
		"PAYMENT_METHOD_CARD":        1,
		"PAYMENT_METHOD_CASH":        2,
		"PAYMENT_METHOD_APPLE_PAY":   3,
		"PAYMENT_METHOD_GOOGLE_PAY":  4,
	}
)

func (x PaymentMethodType) Enum() *PaymentMethodType {
	p := new(PaymentMethodType)
	*p = x
	return p
}

func (x PaymentMethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_payment_payment_proto_enumTypes[1].Descriptor()
}

func (PaymentMethodType) Type() protoreflect.EnumType {
	return &file_v1_payment_payment_proto_enumTypes[1]
}

func (x PaymentMethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethodType.Descriptor instead.
func (PaymentMethodType) EnumDescriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{1}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  *UUID             `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                                                                            // Идентификатор заказа
	Amount   float64           `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`                                                                                           // Сумма платежа
	Method   PaymentMethodType `protobuf:"varint,3,opt,name=method,proto3,enum=payment.PaymentMethodType" json:"method,omitempty"`                                                             // Способ оплаты
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Дополнительные данные
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentRequest) GetOrderId() *UUID {
	if x != nil {
		return x.OrderId
	}
	return nil
}

func (x *PaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentRequest) GetMethod() PaymentMethodType {
	if x != nil {
		return x.Method
	}
	return PaymentMethodType_PAYMENT_METHOD_UNSPECIFIED
}

func (x *PaymentRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"` // Уникальный идентификатор платежа
	Status    string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                        // Текущий статус
	Details   string                 `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`                      // Детализация (для ошибок)
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentResponse) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *PaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *PaymentResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type PaymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *PaymentStatusRequest) Reset() {
	*x = PaymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentStatusRequest) ProtoMessage() {}

func (x *PaymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentStatusRequest.ProtoReflect.Descriptor instead.
func (*PaymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentStatusRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type PaymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       PaymentStatus          `protobuf:"varint,1,opt,name=status,proto3,enum=payment.PaymentStatus" json:"status,omitempty"`
	ErrorMessage string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // Причина ошибки, если есть
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PaymentStatusResponse) Reset() {
	*x = PaymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentStatusResponse) ProtoMessage() {}

func (x *PaymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentStatusResponse.ProtoReflect.Descriptor instead.
func (*PaymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentStatusResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
}

func (x *PaymentStatusResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *PaymentStatusResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string  `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Amount    float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"` // Сумма для возврата
	Reason    string  `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`   // Причина возврата
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *RefundRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *RefundRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RefundRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type RefundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefundId    string                 `protobuf:"bytes,1,opt,name=refund_id,json=refundId,proto3" json:"refund_id,omitempty"`
	Status      PaymentStatus          `protobuf:"varint,2,opt,name=status,proto3,enum=payment.PaymentStatus" json:"status,omitempty"`
	ProcessedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=processed_at,json=processedAt,proto3" json:"processed_at,omitempty"`
}

func (x *RefundResponse) Reset() {
	*x = RefundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundResponse) ProtoMessage() {}

func (x *RefundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundResponse.ProtoReflect.Descriptor instead.
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{6}
}

func (x *RefundResponse) GetRefundId() string {
	if x != nil {
		return x.RefundId
	}
	return ""
}

func (x *RefundResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
}

func (x *RefundResponse) GetProcessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ProcessedAt
	}
	return nil
}

// Конфигурация поведения заглушки
type MockConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Процент успешных платежей (0-100)
	SuccessRate int32 `protobuf:"varint,1,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	// Фиксированный ответ (если задан, success_rate игнорируется)
	FixedStatus *PaymentStatus `protobuf:"varint,2,opt,name=fixed_status,json=fixedStatus,proto3,enum=payment.PaymentStatus,oneof" json:"fixed_status,omitempty"`
	// Задержка ответа в миллисекундах
	ResponseDelayMs int32 `protobuf:"varint,3,opt,name=response_delay_ms,json=responseDelayMs,proto3" json:"response_delay_ms,omitempty"`
	// Регулярное выражение для ошибок по метаданным
	ErrorPattern string `protobuf:"bytes,4,opt,name=error_pattern,json=errorPattern,proto3" json:"error_pattern,omitempty"` // Пример: ".*TEST_FAIL.*"
	// Автоматическое время ответа
	AutoApproveRefunds bool `protobuf:"varint,5,opt,name=auto_approve_refunds,json=autoApproveRefunds,proto3" json:"auto_approve_refunds,omitempty"` // Авто-подтверждение возвратов
}

func (x *MockConfig) Reset() {
	*x = MockConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MockConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MockConfig) ProtoMessage() {}

func (x *MockConfig) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MockConfig.ProtoReflect.Descriptor instead.
func (*MockConfig) Descriptor() ([]byte, []int) {
	return file_v1_payment_payment_proto_rawDescGZIP(), []int{7}
}

func (x *MockConfig) GetSuccessRate() int32 {
	if x != nil {
		return x.SuccessRate
	}
	return 0
}

func (x *MockConfig) GetFixedStatus() PaymentStatus {
	if x != nil && x.FixedStatus != nil {
		return *x.FixedStatus
	}
	return PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
}

func (x *MockConfig) GetResponseDelayMs() int32 {
	if x != nil {
		return x.ResponseDelayMs
	}
	return 0
}

func (x *MockConfig) GetErrorPattern() string {
	if x != nil {
		return x.ErrorPattern
	}
	return ""
}

func (x *MockConfig) GetAutoApproveRefunds() bool {
	if x != nil {
		return x.AutoApproveRefunds
	}
	return false
}

var File_v1_payment_payment_proto protoreflect.FileDescriptor

var file_v1_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x86, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x41, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x35, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0xa7, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x02, 0x0a, 0x0a, 0x4d, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xa1,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x41,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x45, 0x44,
	0x10, 0x04, 0x2a, 0xa2, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x41,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x50, 0x50,
	0x4c, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x5f, 0x50, 0x41, 0x59, 0x10, 0x04, 0x32, 0xee, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x63, 0x6b, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x63, 0x6b,
	0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x6e, 0x65, 0x74, 0x73,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_payment_payment_proto_rawDescOnce sync.Once
	file_v1_payment_payment_proto_rawDescData = file_v1_payment_payment_proto_rawDesc
)

func file_v1_payment_payment_proto_rawDescGZIP() []byte {
	file_v1_payment_payment_proto_rawDescOnce.Do(func() {
		file_v1_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_payment_payment_proto_rawDescData)
	})
	return file_v1_payment_payment_proto_rawDescData
}

var file_v1_payment_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_payment_payment_proto_goTypes = []interface{}{
	(PaymentStatus)(0),            // 0: payment.PaymentStatus
	(PaymentMethodType)(0),        // 1: payment.PaymentMethodType
	(*UUID)(nil),                  // 2: payment.UUID
	(*PaymentRequest)(nil),        // 3: payment.PaymentRequest
	(*PaymentResponse)(nil),       // 4: payment.PaymentResponse
	(*PaymentStatusRequest)(nil),  // 5: payment.PaymentStatusRequest
	(*PaymentStatusResponse)(nil), // 6: payment.PaymentStatusResponse
	(*RefundRequest)(nil),         // 7: payment.RefundRequest
	(*RefundResponse)(nil),        // 8: payment.RefundResponse
	(*MockConfig)(nil),            // 9: payment.MockConfig
	nil,                           // 10: payment.PaymentRequest.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_v1_payment_payment_proto_depIdxs = []int32{
	2,  // 0: payment.PaymentRequest.order_id:type_name -> payment.UUID
	1,  // 1: payment.PaymentRequest.method:type_name -> payment.PaymentMethodType
	10, // 2: payment.PaymentRequest.metadata:type_name -> payment.PaymentRequest.MetadataEntry
	11, // 3: payment.PaymentResponse.created_at:type_name -> google.protobuf.Timestamp
	0,  // 4: payment.PaymentStatusResponse.status:type_name -> payment.PaymentStatus
	11, // 5: payment.PaymentStatusResponse.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 6: payment.RefundResponse.status:type_name -> payment.PaymentStatus
	11, // 7: payment.RefundResponse.processed_at:type_name -> google.protobuf.Timestamp
	0,  // 8: payment.MockConfig.fixed_status:type_name -> payment.PaymentStatus
	3,  // 9: payment.PaymentService.CreatePayment:input_type -> payment.PaymentRequest
	5,  // 10: payment.PaymentService.GetPaymentStatus:input_type -> payment.PaymentStatusRequest
	7,  // 11: payment.PaymentService.ProcessRefund:input_type -> payment.RefundRequest
	9,  // 12: payment.PaymentService.SetMockBehavior:input_type -> payment.MockConfig
	12, // 13: payment.PaymentService.ResetMockBehavior:input_type -> google.protobuf.Empty
	4,  // 14: payment.PaymentService.CreatePayment:output_type -> payment.PaymentResponse
	6,  // 15: payment.PaymentService.GetPaymentStatus:output_type -> payment.PaymentStatusResponse
	8,  // 16: payment.PaymentService.ProcessRefund:output_type -> payment.RefundResponse
	12, // 17: payment.PaymentService.SetMockBehavior:output_type -> google.protobuf.Empty
	12, // 18: payment.PaymentService.ResetMockBehavior:output_type -> google.protobuf.Empty
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_v1_payment_payment_proto_init() }
func file_v1_payment_payment_proto_init() {
	if File_v1_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_payment_payment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MockConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_v1_payment_payment_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_payment_payment_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_payment_payment_proto_goTypes,
		DependencyIndexes: file_v1_payment_payment_proto_depIdxs,
		EnumInfos:         file_v1_payment_payment_proto_enumTypes,
		MessageInfos:      file_v1_payment_payment_proto_msgTypes,
	}.Build()
	File_v1_payment_payment_proto = out.File
	file_v1_payment_payment_proto_rawDesc = nil
	file_v1_payment_payment_proto_goTypes = nil
	file_v1_payment_payment_proto_depIdxs = nil
}
