// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/lawrencefinn/Documents/sidecar/app/src/sidecar/pkg/aws/handler/s3/handler_interface.go

// Package mock_s3 is a generated GoMock package.
package s3

import (
	storage "cloud.google.com/go/storage"
	context "context"
	s3iface "github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
	gomock "github.com/golang/mock/gomock"
	viper "github.com/spf13/viper"
	reflect "reflect"
)

// MockGCPClient is a mock of GCPClient interface
type MockGCPClient struct {
	ctrl     *gomock.Controller
	recorder *MockGCPClientMockRecorder
}

// MockGCPClientMockRecorder is the mock recorder for MockGCPClient
type MockGCPClientMockRecorder struct {
	mock *MockGCPClient
}

// NewMockGCPClient creates a new mock instance
func NewMockGCPClient(ctrl *gomock.Controller) *MockGCPClient {
	mock := &MockGCPClient{ctrl: ctrl}
	mock.recorder = &MockGCPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCPClient) EXPECT() *MockGCPClientMockRecorder {
	return m.recorder
}

// Bucket mocks base method
func (m *MockGCPClient) Bucket(name string) *storage.BucketHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", name)
	ret0, _ := ret[0].(*storage.BucketHandle)
	return ret0
}

// Bucket indicates an expected call of Bucket
func (mr *MockGCPClientMockRecorder) Bucket(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockGCPClient)(nil).Bucket), name)
}

// MockGCPObject is a mock of GCPObject interface
type MockGCPObject struct {
	ctrl     *gomock.Controller
	recorder *MockGCPObjectMockRecorder
}

// MockGCPObjectMockRecorder is the mock recorder for MockGCPObject
type MockGCPObjectMockRecorder struct {
	mock *MockGCPObject
}

// NewMockGCPObject creates a new mock instance
func NewMockGCPObject(ctrl *gomock.Controller) *MockGCPObject {
	mock := &MockGCPObject{ctrl: ctrl}
	mock.recorder = &MockGCPObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCPObject) EXPECT() *MockGCPObjectMockRecorder {
	return m.recorder
}

// ACL mocks base method
func (m *MockGCPObject) ACL() *storage.ACLHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACL")
	ret0, _ := ret[0].(*storage.ACLHandle)
	return ret0
}

// ACL indicates an expected call of ACL
func (mr *MockGCPObjectMockRecorder) ACL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACL", reflect.TypeOf((*MockGCPObject)(nil).ACL))
}

// Generation mocks base method
func (m *MockGCPObject) Generation(gen int64) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generation", gen)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// Generation indicates an expected call of Generation
func (mr *MockGCPObjectMockRecorder) Generation(gen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generation", reflect.TypeOf((*MockGCPObject)(nil).Generation), gen)
}

// If mocks base method
func (m *MockGCPObject) If(conds storage.Conditions) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "If", conds)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// If indicates an expected call of If
func (mr *MockGCPObjectMockRecorder) If(conds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "If", reflect.TypeOf((*MockGCPObject)(nil).If), conds)
}

// Key mocks base method
func (m *MockGCPObject) Key(encryptionKey []byte) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key", encryptionKey)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockGCPObjectMockRecorder) Key(encryptionKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockGCPObject)(nil).Key), encryptionKey)
}

// Attrs mocks base method
func (m *MockGCPObject) Attrs(ctx context.Context) (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attrs", ctx)
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attrs indicates an expected call of Attrs
func (mr *MockGCPObjectMockRecorder) Attrs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockGCPObject)(nil).Attrs), ctx)
}

// Update mocks base method
func (m *MockGCPObject) Update(ctx context.Context, uattrs storage.ObjectAttrsToUpdate) (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, uattrs)
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGCPObjectMockRecorder) Update(ctx, uattrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGCPObject)(nil).Update), ctx, uattrs)
}

// BucketName mocks base method
func (m *MockGCPObject) BucketName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BucketName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BucketName indicates an expected call of BucketName
func (mr *MockGCPObjectMockRecorder) BucketName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BucketName", reflect.TypeOf((*MockGCPObject)(nil).BucketName))
}

// ObjectName mocks base method
func (m *MockGCPObject) ObjectName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ObjectName indicates an expected call of ObjectName
func (mr *MockGCPObjectMockRecorder) ObjectName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectName", reflect.TypeOf((*MockGCPObject)(nil).ObjectName))
}

// Delete mocks base method
func (m *MockGCPObject) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGCPObjectMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGCPObject)(nil).Delete), ctx)
}

// ReadCompressed mocks base method
func (m *MockGCPObject) ReadCompressed(compressed bool) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCompressed", compressed)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// ReadCompressed indicates an expected call of ReadCompressed
func (mr *MockGCPObjectMockRecorder) ReadCompressed(compressed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCompressed", reflect.TypeOf((*MockGCPObject)(nil).ReadCompressed), compressed)
}

// NewWriter mocks base method
func (m *MockGCPObject) NewWriter(ctx context.Context) *storage.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriter", ctx)
	ret0, _ := ret[0].(*storage.Writer)
	return ret0
}

// NewWriter indicates an expected call of NewWriter
func (mr *MockGCPObjectMockRecorder) NewWriter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriter", reflect.TypeOf((*MockGCPObject)(nil).NewWriter), ctx)
}

// NewReader mocks base method
func (m *MockGCPObject) NewReader(ctx context.Context) (GCPReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReader", ctx)
	ret0, _ := ret[0].(*storage.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReader indicates an expected call of NewReader
func (mr *MockGCPObjectMockRecorder) NewReader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReader", reflect.TypeOf((*MockGCPObject)(nil).NewReader), ctx)
}

// NewRangeReader mocks base method
func (m *MockGCPObject) NewRangeReader(ctx context.Context, offset, length int64) (GCPReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRangeReader", ctx, offset, length)
	ret0, _ := ret[0].(*storage.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRangeReader indicates an expected call of NewRangeReader
func (mr *MockGCPObjectMockRecorder) NewRangeReader(ctx, offset, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRangeReader", reflect.TypeOf((*MockGCPObject)(nil).NewRangeReader), ctx, offset, length)
}

// CopierFrom mocks base method
func (m *MockGCPObject) CopierFrom(src *storage.ObjectHandle) *storage.Copier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopierFrom", src)
	ret0, _ := ret[0].(*storage.Copier)
	return ret0
}

// CopierFrom indicates an expected call of CopierFrom
func (mr *MockGCPObjectMockRecorder) CopierFrom(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopierFrom", reflect.TypeOf((*MockGCPObject)(nil).CopierFrom), src)
}

// ComposerFrom mocks base method
func (m *MockGCPObject) ComposerFrom(srcs ...*storage.ObjectHandle) *storage.Composer {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range srcs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ComposerFrom", varargs...)
	ret0, _ := ret[0].(*storage.Composer)
	return ret0
}

// ComposerFrom indicates an expected call of ComposerFrom
func (mr *MockGCPObjectMockRecorder) ComposerFrom(srcs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComposerFrom", reflect.TypeOf((*MockGCPObject)(nil).ComposerFrom), srcs...)
}

// MockGCPBucket is a mock of GCPBucket interface
type MockGCPBucket struct {
	ctrl     *gomock.Controller
	recorder *MockGCPBucketMockRecorder
}

// MockGCPBucketMockRecorder is the mock recorder for MockGCPBucket
type MockGCPBucketMockRecorder struct {
	mock *MockGCPBucket
}

// NewMockGCPBucket creates a new mock instance
func NewMockGCPBucket(ctrl *gomock.Controller) *MockGCPBucket {
	mock := &MockGCPBucket{ctrl: ctrl}
	mock.recorder = &MockGCPBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCPBucket) EXPECT() *MockGCPBucketMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGCPBucket) Create(ctx context.Context, projectID string, attrs *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, projectID, attrs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGCPBucketMockRecorder) Create(ctx, projectID, attrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGCPBucket)(nil).Create), ctx, projectID, attrs)
}

// Delete mocks base method
func (m *MockGCPBucket) Delete(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGCPBucketMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGCPBucket)(nil).Delete), ctx)
}

// ACL mocks base method
func (m *MockGCPBucket) ACL() *storage.ACLHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACL")
	ret0, _ := ret[0].(*storage.ACLHandle)
	return ret0
}

// ACL indicates an expected call of ACL
func (mr *MockGCPBucketMockRecorder) ACL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACL", reflect.TypeOf((*MockGCPBucket)(nil).ACL))
}

// DefaultObjectACL mocks base method
func (m *MockGCPBucket) DefaultObjectACL() *storage.ACLHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultObjectACL")
	ret0, _ := ret[0].(*storage.ACLHandle)
	return ret0
}

// DefaultObjectACL indicates an expected call of DefaultObjectACL
func (mr *MockGCPBucketMockRecorder) DefaultObjectACL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultObjectACL", reflect.TypeOf((*MockGCPBucket)(nil).DefaultObjectACL))
}

// Object mocks base method
func (m *MockGCPBucket) Object(name string) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Object", name)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// Object indicates an expected call of Object
func (mr *MockGCPBucketMockRecorder) Object(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Object", reflect.TypeOf((*MockGCPBucket)(nil).Object), name)
}

// Attrs mocks base method
func (m *MockGCPBucket) Attrs(ctx context.Context) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attrs", ctx)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attrs indicates an expected call of Attrs
func (mr *MockGCPBucketMockRecorder) Attrs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockGCPBucket)(nil).Attrs), ctx)
}

// Update mocks base method
func (m *MockGCPBucket) Update(ctx context.Context, uattrs storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, uattrs)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGCPBucketMockRecorder) Update(ctx, uattrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGCPBucket)(nil).Update), ctx, uattrs)
}

// If mocks base method
func (m *MockGCPBucket) If(conds storage.BucketConditions) *storage.BucketHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "If", conds)
	ret0, _ := ret[0].(*storage.BucketHandle)
	return ret0
}

// If indicates an expected call of If
func (mr *MockGCPBucketMockRecorder) If(conds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "If", reflect.TypeOf((*MockGCPBucket)(nil).If), conds)
}

// UserProject mocks base method
func (m *MockGCPBucket) UserProject(projectID string) *storage.BucketHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserProject", projectID)
	ret0, _ := ret[0].(*storage.BucketHandle)
	return ret0
}

// UserProject indicates an expected call of UserProject
func (mr *MockGCPBucketMockRecorder) UserProject(projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserProject", reflect.TypeOf((*MockGCPBucket)(nil).UserProject), projectID)
}

// LockRetentionPolicy mocks base method
func (m *MockGCPBucket) LockRetentionPolicy(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockRetentionPolicy", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockRetentionPolicy indicates an expected call of LockRetentionPolicy
func (mr *MockGCPBucketMockRecorder) LockRetentionPolicy(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockRetentionPolicy", reflect.TypeOf((*MockGCPBucket)(nil).LockRetentionPolicy), ctx)
}

// Objects mocks base method
func (m *MockGCPBucket) Objects(ctx context.Context, q *storage.Query) *storage.ObjectIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Objects", ctx, q)
	ret0, _ := ret[0].(*storage.ObjectIterator)
	return ret0
}

// Objects indicates an expected call of Objects
func (mr *MockGCPBucketMockRecorder) Objects(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Objects", reflect.TypeOf((*MockGCPBucket)(nil).Objects), ctx, q)
}

// MockHandlerInterface is a mock of HandlerInterface interface
type MockHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerInterfaceMockRecorder
}

// MockHandlerInterfaceMockRecorder is the mock recorder for MockHandlerInterface
type MockHandlerInterfaceMockRecorder struct {
	mock *MockHandlerInterface
}

// NewMockHandlerInterface creates a new mock instance
func NewMockHandlerInterface(ctrl *gomock.Controller) *MockHandlerInterface {
	mock := &MockHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandlerInterface) EXPECT() *MockHandlerInterfaceMockRecorder {
	return m.recorder
}

// GetS3Client mocks base method
func (m *MockHandlerInterface) GetS3Client() s3iface.S3API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetS3Client")
	ret0, _ := ret[0].(s3iface.S3API)
	return ret0
}

// GetS3Client indicates an expected call of GetS3Client
func (mr *MockHandlerInterfaceMockRecorder) GetS3Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetS3Client", reflect.TypeOf((*MockHandlerInterface)(nil).GetS3Client))
}

// GetGCPClient mocks base method
func (m *MockHandlerInterface) GetGCPClient() GCPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGCPClient")
	ret0, _ := ret[0].(GCPClient)
	return ret0
}

// GetGCPClient indicates an expected call of GetGCPClient
func (mr *MockHandlerInterfaceMockRecorder) GetGCPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGCPClient", reflect.TypeOf((*MockHandlerInterface)(nil).GetGCPClient))
}

// GetContext mocks base method
func (m *MockHandlerInterface) GetContext() *context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(*context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext
func (mr *MockHandlerInterfaceMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockHandlerInterface)(nil).GetContext))
}

// GetConfig mocks base method
func (m *MockHandlerInterface) GetConfig() *viper.Viper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*viper.Viper)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockHandlerInterfaceMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockHandlerInterface)(nil).GetConfig))
}

// SetS3Client mocks base method
func (m *MockHandlerInterface) SetS3Client(s3Client s3iface.S3API) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetS3Client", s3Client)
}

// SetS3Client indicates an expected call of SetS3Client
func (mr *MockHandlerInterfaceMockRecorder) SetS3Client(s3Client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetS3Client", reflect.TypeOf((*MockHandlerInterface)(nil).SetS3Client), s3Client)
}

// SetGCPClient mocks base method
func (m *MockHandlerInterface) SetGCPClient(gcpClient GCPClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGCPClient", gcpClient)
}

// SetGCPClient indicates an expected call of SetGCPClient
func (mr *MockHandlerInterfaceMockRecorder) SetGCPClient(gcpClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGCPClient", reflect.TypeOf((*MockHandlerInterface)(nil).SetGCPClient), gcpClient)
}

// SetContext mocks base method
func (m *MockHandlerInterface) SetContext(context *context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContext", context)
}

// SetContext indicates an expected call of SetContext
func (mr *MockHandlerInterfaceMockRecorder) SetContext(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContext", reflect.TypeOf((*MockHandlerInterface)(nil).SetContext), context)
}

// SetConfig mocks base method
func (m *MockHandlerInterface) SetConfig(config *viper.Viper) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfig", config)
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockHandlerInterfaceMockRecorder) SetConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockHandlerInterface)(nil).SetConfig), config)
}
