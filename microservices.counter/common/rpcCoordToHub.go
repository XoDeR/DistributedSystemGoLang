package common

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gizmo/web"
	"golang.org/x/net/context"
)

func (s *RPCService) AddNewItemWithTenant(ctx context.Context, r *AddNewItemWithTenantRequest) (*AddNewItemWithTenantResponse, error) {
	var (
		err error
		res *AddNewItemWithTenantResult
	)
	defer server.MonitorRPCRequest()(ctx, "AddNewItemWithTenant", err)

	res, err = s.client.AddNewItemWithTenant(r.ItemId, r.TenantId)
	if err != nil {
		return nil, err
	}
	return &AddNewItemWithTenantResponse{res}, nil
}

func (s *RPCService) GetItemCountByTenant(ctx context.Context, r *GetItemCountByTenantRequest) (*GetItemCountByTenantResponse, error) {
	var (
		err error
		res *GetItemCountByTenantResult
	)
	defer server.MonitorRPCRequest()(ctx, "GetItemCountByTenant", err)

	res, err = s.client.GetItemCountByTenant(r.TenantId)
	if err != nil {
		return nil, err
	}
	return &GetItemCountByTenantResponse{res}, nil
}

func (s *RPCService) AddNewItemWithTenantJSON(ctx context.Context, r *http.Request) (int, interface{}, error) {
	res, err := s.AddNewItemWithTenant(
		ctx,
		&AddNewItemWithTenantRequest{
			uint32(web.GetUInt64Var(r, "itemId")),
			uint32(web.GetUInt64Var(r, "tenantId"))})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, res.Result, nil
}

func (s *RPCService) GetItemCountByTenantJSON(ctx context.Context, r *http.Request) (int, interface{}, error) {
	res, err := s.GetItemCountByTenant(
		ctx,
		&GetItemCountByTenantRequest{
			uint32(web.GetUInt64Var(r, "tenantId")),
		})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, res.Result, nil
}
