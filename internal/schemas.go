package internal

import (
	"context"

	omega_catalogv1 "github.com/zcking/omega-catalog/gen/go/omega_catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Impl) CreateSchema(ctx context.Context, req *omega_catalogv1.CreateSchemaRequest) (*omega_catalogv1.CreateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}

func (i *Impl) ListSchemas(ctx context.Context, req *omega_catalogv1.ListSchemasRequest) (*omega_catalogv1.ListSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}

func (i *Impl) GetSchema(ctx context.Context, req *omega_catalogv1.GetSchemaRequest) (*omega_catalogv1.GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}

func (i *Impl) UpdateSchema(ctx context.Context, req *omega_catalogv1.UpdateSchemaRequest) (*omega_catalogv1.UpdateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchema not implemented")
}

func (i *Impl) DeleteSchema(ctx context.Context, req *omega_catalogv1.DeleteSchemaRequest) (*omega_catalogv1.DeleteSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchema not implemented")
}

func (i *Impl) CreateTable(ctx context.Context, req *omega_catalogv1.CreateTableRequest) (*omega_catalogv1.CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet	")
}
func (i *Impl) ListTables(ctx context.Context, req *omega_catalogv1.ListTablesRequest) (*omega_catalogv1.ListTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) GetTable(ctx context.Context, req *omega_catalogv1.GetTableRequest) (*omega_catalogv1.GetTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) DeleteTable(ctx context.Context, req *omega_catalogv1.DeleteTableRequest) (*omega_catalogv1.DeleteTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) CreateVolume(ctx context.Context, req *omega_catalogv1.CreateVolumeRequest) (*omega_catalogv1.CreateVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) ListVolumes(ctx context.Context, req *omega_catalogv1.ListVolumesRequest) (*omega_catalogv1.ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) GetVolume(ctx context.Context, req *omega_catalogv1.GetVolumeRequest) (*omega_catalogv1.GetVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) UpdateVolume(ctx context.Context, req *omega_catalogv1.UpdateVolumeRequest) (*omega_catalogv1.UpdateVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) DeleteVolume(ctx context.Context, req *omega_catalogv1.DeleteVolumeRequest) (*omega_catalogv1.DeleteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) CreateFunction(ctx context.Context, req *omega_catalogv1.CreateFunctionRequest) (*omega_catalogv1.CreateFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) ListFunctions(ctx context.Context, req *omega_catalogv1.ListFunctionsRequest) (*omega_catalogv1.ListFunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) GetFunction(ctx context.Context, req *omega_catalogv1.GetFunctionRequest) (*omega_catalogv1.GetFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
func (i *Impl) DeleteFunction(ctx context.Context, req *omega_catalogv1.DeleteFunctionRequest) (*omega_catalogv1.DeleteFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented yet")
}
