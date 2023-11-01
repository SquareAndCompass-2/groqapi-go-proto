// Copyright 2023 Groq Inc.
// All Rights Reserved.
//
// Definition of the Shim Service
// This is the inverse of the Inference API Service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: public/llmcloud/inferencemanager/v1/inference_result_service.proto

package inferencemanagerv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InferenceResultServiceGetResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InferenceResultServiceGetResultResponse) Reset() {
	*x = InferenceResultServiceGetResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResultServiceGetResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResultServiceGetResultResponse) ProtoMessage() {}

func (x *InferenceResultServiceGetResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResultServiceGetResultResponse.ProtoReflect.Descriptor instead.
func (*InferenceResultServiceGetResultResponse) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescGZIP(), []int{0}
}

type ProcessingStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total number of tokens processed
	TokensProcessed uint32 `protobuf:"varint,1,opt,name=tokens_processed,json=tokensProcessed,proto3" json:"tokens_processed,omitempty"`
	// time in millis it took to process the query tokens
	ProcessTime *durationpb.Duration `protobuf:"bytes,2,opt,name=process_time,json=processTime,proto3" json:"process_time,omitempty"`
}

func (x *ProcessingStats) Reset() {
	*x = ProcessingStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingStats) ProtoMessage() {}

func (x *ProcessingStats) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingStats.ProtoReflect.Descriptor instead.
func (*ProcessingStats) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessingStats) GetTokensProcessed() uint32 {
	if x != nil {
		return x.TokensProcessed
	}
	return 0
}

func (x *ProcessingStats) GetProcessTime() *durationpb.Duration {
	if x != nil {
		return x.ProcessTime
	}
	return nil
}

type GenerationStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total number of tokens generated
	TokensGenerated uint32 `protobuf:"varint,1,opt,name=tokens_generated,json=tokensGenerated,proto3" json:"tokens_generated,omitempty"`
	// time in millis it took to generate the response tokens
	GenerateTime *durationpb.Duration `protobuf:"bytes,2,opt,name=generate_time,json=generateTime,proto3" json:"generate_time,omitempty"`
}

func (x *GenerationStats) Reset() {
	*x = GenerationStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerationStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationStats) ProtoMessage() {}

func (x *GenerationStats) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationStats.ProtoReflect.Descriptor instead.
func (*GenerationStats) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescGZIP(), []int{2}
}

func (x *GenerationStats) GetTokensGenerated() uint32 {
	if x != nil {
		return x.TokensGenerated
	}
	return 0
}

func (x *GenerationStats) GetGenerateTime() *durationpb.Duration {
	if x != nil {
		return x.GenerateTime
	}
	return nil
}

type InferenceResultServiceGetResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the newly generated token data
	Token uint32 `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	// the ID of the request
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// response address
	ResponseAddress string `protobuf:"bytes,3,opt,name=response_address,json=responseAddress,proto3" json:"response_address,omitempty"`
	// Statistics about the initial processing
	// Static per request once provided
	// Optional - sent at most once per stream
	ProcessingStats *ProcessingStats `protobuf:"bytes,4,opt,name=processing_stats,json=processingStats,proto3" json:"processing_stats,omitempty"`
	// Statistics about the inference generation
	GenerationStats *GenerationStats `protobuf:"bytes,5,opt,name=generation_stats,json=generationStats,proto3" json:"generation_stats,omitempty"`
	// True if generation is finished
	GenerationFinished bool `protobuf:"varint,6,opt,name=generation_finished,json=generationFinished,proto3" json:"generation_finished,omitempty"`
	// Model to use
	ModelId string `protobuf:"bytes,7,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
}

func (x *InferenceResultServiceGetResultRequest) Reset() {
	*x = InferenceResultServiceGetResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResultServiceGetResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResultServiceGetResultRequest) ProtoMessage() {}

func (x *InferenceResultServiceGetResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResultServiceGetResultRequest.ProtoReflect.Descriptor instead.
func (*InferenceResultServiceGetResultRequest) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescGZIP(), []int{3}
}

func (x *InferenceResultServiceGetResultRequest) GetToken() uint32 {
	if x != nil {
		return x.Token
	}
	return 0
}

func (x *InferenceResultServiceGetResultRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *InferenceResultServiceGetResultRequest) GetResponseAddress() string {
	if x != nil {
		return x.ResponseAddress
	}
	return ""
}

func (x *InferenceResultServiceGetResultRequest) GetProcessingStats() *ProcessingStats {
	if x != nil {
		return x.ProcessingStats
	}
	return nil
}

func (x *InferenceResultServiceGetResultRequest) GetGenerationStats() *GenerationStats {
	if x != nil {
		return x.GenerationStats
	}
	return nil
}

func (x *InferenceResultServiceGetResultRequest) GetGenerationFinished() bool {
	if x != nil {
		return x.GenerationFinished
	}
	return false
}

func (x *InferenceResultServiceGetResultRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

var File_public_llmcloud_inferencemanager_v1_inference_result_service_proto protoreflect.FileDescriptor

var file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDesc = []byte{
	0x0a, 0x42, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x27, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x7c, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3e,
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x96,
	0x03, 0x0a, 0x26, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x10, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x32, 0xc5, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xaa, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x4b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4c, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0xca, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x2e,
	0x67, 0x72, 0x6f, 0x71, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6c,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x4c, 0x49, 0xaa, 0x02, 0x23, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4c, 0x6c, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x5c, 0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x30, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x4c, 0x6c, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x4c, 0x6c,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescOnce sync.Once
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescData = file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDesc
)

func file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescGZIP() []byte {
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescOnce.Do(func() {
		file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescData)
	})
	return file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDescData
}

var file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_goTypes = []interface{}{
	(*InferenceResultServiceGetResultResponse)(nil), // 0: public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultResponse
	(*ProcessingStats)(nil),                         // 1: public.llmcloud.inferencemanager.v1.ProcessingStats
	(*GenerationStats)(nil),                         // 2: public.llmcloud.inferencemanager.v1.GenerationStats
	(*InferenceResultServiceGetResultRequest)(nil),  // 3: public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultRequest
	(*durationpb.Duration)(nil),                     // 4: google.protobuf.Duration
}
var file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_depIdxs = []int32{
	4, // 0: public.llmcloud.inferencemanager.v1.ProcessingStats.process_time:type_name -> google.protobuf.Duration
	4, // 1: public.llmcloud.inferencemanager.v1.GenerationStats.generate_time:type_name -> google.protobuf.Duration
	1, // 2: public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultRequest.processing_stats:type_name -> public.llmcloud.inferencemanager.v1.ProcessingStats
	2, // 3: public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultRequest.generation_stats:type_name -> public.llmcloud.inferencemanager.v1.GenerationStats
	3, // 4: public.llmcloud.inferencemanager.v1.InferenceResultService.GetResult:input_type -> public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultRequest
	0, // 5: public.llmcloud.inferencemanager.v1.InferenceResultService.GetResult:output_type -> public.llmcloud.inferencemanager.v1.InferenceResultServiceGetResultResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_init() }
func file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_init() {
	if File_public_llmcloud_inferencemanager_v1_inference_result_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResultServiceGetResultResponse); i {
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
		file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingStats); i {
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
		file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerationStats); i {
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
		file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResultServiceGetResultRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_goTypes,
		DependencyIndexes: file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_depIdxs,
		MessageInfos:      file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_msgTypes,
	}.Build()
	File_public_llmcloud_inferencemanager_v1_inference_result_service_proto = out.File
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_rawDesc = nil
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_goTypes = nil
	file_public_llmcloud_inferencemanager_v1_inference_result_service_proto_depIdxs = nil
}
