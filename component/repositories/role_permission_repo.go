package repositories

//
//import (
//	"context"
//	"fmt"
//	"github.com/jmoiron/sqlx"
//	"github.com/lib/pq"
//	log "github.com/sirupsen/logrus"
//	"main/component/models"
//	"main/utils"
//	"strings"
//)
//
//type RolePermissionRepo interface {
//	FindPermissionsBy(ctx context.Context, filter models.RolePermissionFilter) ([]models.RolePermissionQueryResponse, error)
//}
//type rolePermissionRepo struct {
//	db *sqlx.DB
//}
//
//func NewRolePermissionRepo(db *sqlx.DB) RolePermissionRepo {
//	return &rolePermissionRepo{db: db}
//}
//
//func (r *rolePermissionRepo) FindPermissionsBy(ctx context.Context, filter models.RolePermissionFilter) ([]models.RolePermissionQueryResponse, error) {
//	query, args := r.buildQueryFindPermissionsBy(filter)
//	var response []models.RolePermissionQueryResponse
//	err := r.db.SelectContext(ctx, &response, query, args...)
//	if err != nil {
//		log.Errorf("Error get permission with filter %v, error %v", utils.LogFull(filter), err)
//		return nil, err
//	}
//
//	return response, nil
//}
//
//func (r *rolePermissionRepo) buildQueryFindPermissionsBy(filter models.RolePermissionFilter) (string, []interface{}) {
//	query := strings.Builder{}
//	args := make([]interface{}, 0)
//
//	query.WriteString(" select iug.user_id, irp.permission_code, irp.role_code, ig.org_id , ig.name as geo_name, ig.label as geo_label ")
//	query.WriteString(" from ims_role_permissions irp  ")
//	query.WriteString(" left join ims_roles ir on ir.role_code = irp.role_code ")
//	query.WriteString(" left join ims_users_roles iur on iur.role_id = ir.role_id ")
//	query.WriteString(" left join ims_users_geos iug on iug.user_id = iur.user_id ")
//	query.WriteString(" left join ims_geos ig on ig.org_id = iug.org_id ")
//	query.WriteString(" where 1=1 ")
//
//	i := 1 //index
//	if filter.UserId != 0 {
//		query.WriteString(fmt.Sprintf(" and iug.user_id = $%d ", i))
//		args = append(args, filter.UserId)
//		i++
//	}
//
//	if filter.RoleCodes != nil && len(filter.RoleCodes) > 0 {
//		query.WriteString(fmt.Sprintf(" and ir.role_code = any($%d) ", i))
//		args = append(args, pq.Array(filter.RoleCodes))
//		i++
//	}
//
//	if filter.PermissionCode != "" {
//		query.WriteString(fmt.Sprintf(" and irp.permission_code = $%d ", i))
//		args = append(args, filter.PermissionCode)
//		i++
//	}
//
//	if filter.OrgId != 0 {
//		query.WriteString(fmt.Sprintf(" and ig.org_id = $%d ", i))
//		args = append(args, filter.OrgId)
//		i++
//	}
//
//	if filter.GeoName != "" {
//		query.WriteString(fmt.Sprintf(" ig.name = $%d ", i))
//		args = append(args, filter.GeoName)
//		i++
//	}
//
//	if filter.GeoLabel != "" {
//		query.WriteString(fmt.Sprintf(" and ig.label = $%d ", i))
//		args = append(args, filter.GeoLabel)
//		i++
//	}
//
//	query.WriteString(" group by iug.user_id, irp.permission_code, irp.role_code, ig.org_id, ig.name, ig.label ")
//
//	return query.String(), args
//}
