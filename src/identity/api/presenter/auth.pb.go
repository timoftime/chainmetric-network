// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: identity/api/presenter/auth.proto

package presenter

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash string `protobuf:"bytes,2,opt,name=passwordHash,proto3" json:"passwordHash,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_api_presenter_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_api_presenter_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_identity_api_presenter_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type SetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	PasswordHash string `protobuf:"bytes,2,opt,name=passwordHash,proto3" json:"passwordHash,omitempty"`
}

func (x *SetPasswordRequest) Reset() {
	*x = SetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_api_presenter_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordRequest) ProtoMessage() {}

func (x *SetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_api_presenter_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordRequest.ProtoReflect.Descriptor instead.
func (*SetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_identity_api_presenter_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SetPasswordRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SetPasswordRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretToken string `protobuf:"bytes,1,opt,name=secretToken,proto3" json:"secretToken,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_api_presenter_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_api_presenter_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_identity_api_presenter_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetSecretToken() string {
	if x != nil {
		return x.SecretToken
	}
	return ""
}

var File_identity_api_presenter_auth_proto protoreflect.FileDescriptor

var file_identity_api_presenter_auth_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x0c, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x08, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x63, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x2b, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x08, 0x52, 0x0c,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x30, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x46,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6d,
	0x6f, 0x74, 0x68, 0x2d, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_api_presenter_auth_proto_rawDescOnce sync.Once
	file_identity_api_presenter_auth_proto_rawDescData = file_identity_api_presenter_auth_proto_rawDesc
)

func file_identity_api_presenter_auth_proto_rawDescGZIP() []byte {
	file_identity_api_presenter_auth_proto_rawDescOnce.Do(func() {
		file_identity_api_presenter_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_api_presenter_auth_proto_rawDescData)
	})
	return file_identity_api_presenter_auth_proto_rawDescData
}

var file_identity_api_presenter_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_identity_api_presenter_auth_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),        // 0: chainmetric.identity.presenter.AuthRequest
	(*SetPasswordRequest)(nil), // 1: chainmetric.identity.presenter.SetPasswordRequest
	(*AuthResponse)(nil),       // 2: chainmetric.identity.presenter.AuthResponse
}
var file_identity_api_presenter_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identity_api_presenter_auth_proto_init() }
func file_identity_api_presenter_auth_proto_init() {
	if File_identity_api_presenter_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_api_presenter_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_identity_api_presenter_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordRequest); i {
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
		file_identity_api_presenter_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
			RawDescriptor: file_identity_api_presenter_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_identity_api_presenter_auth_proto_goTypes,
		DependencyIndexes: file_identity_api_presenter_auth_proto_depIdxs,
		MessageInfos:      file_identity_api_presenter_auth_proto_msgTypes,
	}.Build()
	File_identity_api_presenter_auth_proto = out.File
	file_identity_api_presenter_auth_proto_rawDesc = nil
	file_identity_api_presenter_auth_proto_goTypes = nil
	file_identity_api_presenter_auth_proto_depIdxs = nil
}
