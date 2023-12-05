package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcube/v2/ioc/config/cache"
)

func CacheWrapper(i *impl) *Decorator {
	return &Decorator{i}
}

type Decorator struct {
	*impl
}

func (d *Decorator) DescribeRole(ctx context.Context, req *role.DescribeRoleRequest) (
	*role.Role, error) {
	if req.Id != "" {
		ins := role.NewDefaultRole()
		if err := cache.C().Get(ctx, req.Id, ins); err != nil {
			d.log.Warn().Msgf("get %s from cache error, %s", req.Id, err)
		} else {
			d.log.Info().Msgf("get %s from cache", ins.Meta.Id)
			return ins, nil
		}
	}

	ins, err := d.impl.DescribeRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.Id != "" {
		if err := cache.C().Set(ctx, req.Id, ins); err != nil {
			d.log.Warn().Msgf("set %s to cache error, %s", req.Id, err)
		} else {
			d.log.Info().Msgf("set %s to cache", ins.Meta.Id)
		}
	}
	return ins, nil
}

func (d *Decorator) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (
	*role.Role, error) {
	ins, err := d.impl.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Del(ctx, req.Id); err != nil {
		d.log.Info().Msgf("delete %s to cache error, %s", req.Id, err)
	}
	return ins, nil
}

func (d *Decorator) AddPermissionToRole(ctx context.Context, req *role.AddPermissionToRoleRequest) (
	*role.Role, error) {
	ins, err := d.impl.AddPermissionToRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Del(ctx, req.RoleId); err != nil {
		d.log.Info().Msgf("delete %s to cache error, %s", req.RoleId, err)
	}
	return ins, nil
}

func (d *Decorator) RemovePermissionFromRole(ctx context.Context, req *role.RemovePermissionFromRoleRequest) (
	*role.Role, error) {
	ins, err := d.impl.RemovePermissionFromRole(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Del(ctx, req.RoleId); err != nil {
		d.log.Info().Msgf("delete %s to cache error, %s", req.RoleId, err)
	}
	return ins, nil
}

func (d *Decorator) UpdatePermission(ctx context.Context, req *role.UpdatePermissionRequest) (
	*role.Role, error) {
	ins, err := d.impl.UpdatePermission(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := cache.C().Del(ctx, ins.Meta.Id); err != nil {
		d.log.Info().Msgf("delete %s to cache error, %s", req.Id, err)
	}
	return ins, err
}
