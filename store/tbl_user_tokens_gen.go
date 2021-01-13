package store

////////////////////////////////////////////////////////////////////
//   Warning: this file is auto, please not edit.  自动生成请勿编辑
//   Author : ftfeng
//   主要包含：1、创建/删除/更新实例  2、更新实例 3、根据列获取所有 3、根据列分页获取
//   Create time : 2021-01-14 00:21:33
////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/golib/common/predicate"

	"time"
)

//  新增
func (store *Store) CreateTblUserToken(tblUserToken *models.TblUserToken) error {
	return store.db.Create(&tblUserToken).Error
}

//  删除
func (store *Store) DeleteTblUserToken(tblUserToken *models.TblUserToken) error {
	return store.db.Delete(&tblUserToken).Error
}

// 更新
func (store *Store) UpdateTblUserToken(tblUserToken *models.TblUserToken) error {
	return store.db.Save(&tblUserToken).Error
}

// GetBy主键
func (store *Store) GetOneTblUserTokenById(p predicate.Predicate, id uint) (*models.TblUserToken, error) {
	var model models.TblUserToken
	return &model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// list全部
func (store *Store) GetTblUserToken() ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Find(&model).Error
}

// list分页
func (store *Store) GetTblUserTokenByPage(limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值查询所有
func (store *Store) GetTblUserTokenById(p predicate.Predicate, id uint) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// 根据CreatedAt值查询所有
func (store *Store) GetTblUserTokenByCreatedAt(p predicate.Predicate, createdAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Find(&model).Error
}

// 根据UpdatedAt值查询所有
func (store *Store) GetTblUserTokenByUpdatedAt(p predicate.Predicate, updatedAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Find(&model).Error
}

// 根据DeletedAt值查询所有
func (store *Store) GetTblUserTokenByDeletedAt(p predicate.Predicate, deletedAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Find(&model).Error
}

// 根据UserName值查询所有
func (store *Store) GetTblUserTokenByUserName(p predicate.Predicate, userName string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Find(&model).Error
}

// 根据UserToken值查询所有
func (store *Store) GetTblUserTokenByUserToken(p predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_token %s ?", p), userToken).Find(&model).Error
}

// 根据Id值分页查询
func (store *Store) GetTblUserTokenByIdByPage(p predicate.Predicate, id uint, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据CreatedAt值分页查询
func (store *Store) GetTblUserTokenByCreatedAtByPage(p predicate.Predicate, createdAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UpdatedAt值分页查询
func (store *Store) GetTblUserTokenByUpdatedAtByPage(p predicate.Predicate, updatedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据DeletedAt值分页查询
func (store *Store) GetTblUserTokenByDeletedAtByPage(p predicate.Predicate, deletedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UserName值分页查询
func (store *Store) GetTblUserTokenByUserNameByPage(p predicate.Predicate, userName string, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UserToken值分页查询
func (store *Store) GetTblUserTokenByUserTokenByPage(p predicate.Predicate, userToken string, limit interface{}, offset interface{}) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_token %s ?", p), userToken).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值批量删除查询
func (store *Store) DelTblUserTokenById(p predicate.Predicate, id uint) error {
	return store.db.Where(fmt.Sprintf("id %s ?", p), id).Delete(models.TblUserToken{}).Error
}

// 根据CreatedAt值批量删除查询
func (store *Store) DelTblUserTokenByCreatedAt(p predicate.Predicate, createdAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Delete(models.TblUserToken{}).Error
}

// 根据UpdatedAt值批量删除查询
func (store *Store) DelTblUserTokenByUpdatedAt(p predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Delete(models.TblUserToken{}).Error
}

// 根据DeletedAt值批量删除查询
func (store *Store) DelTblUserTokenByDeletedAt(p predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Delete(models.TblUserToken{}).Error
}

// 根据UserName值批量删除查询
func (store *Store) DelTblUserTokenByUserName(p predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Delete(models.TblUserToken{}).Error
}

// 根据UserToken值批量删除查询
func (store *Store) DelTblUserTokenByUserToken(p predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("user_token %s ?", p), userToken).Delete(models.TblUserToken{}).Error
}

// 根据CreatedAt UpdatedAt值全部查询
func (store *Store) GetTblUserTokenByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Find(&model).Error
}

// 根据CreatedAt DeletedAt值全部查询
func (store *Store) GetTblUserTokenByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Find(&model).Error
}

// 根据CreatedAt UserName值全部查询
func (store *Store) GetTblUserTokenByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Find(&model).Error
}

// 根据CreatedAt UserToken值全部查询
func (store *Store) GetTblUserTokenByCreatedAtAndUserToken(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userToken).Find(&model).Error
}

// 根据UpdatedAt DeletedAt值全部查询
func (store *Store) GetTblUserTokenByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Find(&model).Error
}

// 根据UpdatedAt UserName值全部查询
func (store *Store) GetTblUserTokenByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Find(&model).Error
}

// 根据UpdatedAt UserToken值全部查询
func (store *Store) GetTblUserTokenByUpdatedAtAndUserToken(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userToken).Find(&model).Error
}

// 根据DeletedAt UserName值全部查询
func (store *Store) GetTblUserTokenByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Find(&model).Error
}

// 根据DeletedAt UserToken值全部查询
func (store *Store) GetTblUserTokenByDeletedAtAndUserToken(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userToken).Find(&model).Error
}

// 根据UserName UserToken值全部查询
func (store *Store) GetTblUserTokenByUserNameAndUserToken(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userToken).Find(&model).Error
}

// 根据CreatedAt UpdatedAt值批量删除查询
func (store *Store) DelTblUserTokenByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Delete(models.TblUserToken{}).Error
}

// 根据CreatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserTokenByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Delete(models.TblUserToken{}).Error
}

// 根据CreatedAt UserName值批量删除查询
func (store *Store) DelTblUserTokenByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Delete(models.TblUserToken{}).Error
}

// 根据CreatedAt UserToken值批量删除查询
func (store *Store) DelTblUserTokenByCreatedAtAndUserToken(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userToken).Delete(models.TblUserToken{}).Error
}

// 根据UpdatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserTokenByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Delete(models.TblUserToken{}).Error
}

// 根据UpdatedAt UserName值批量删除查询
func (store *Store) DelTblUserTokenByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Delete(models.TblUserToken{}).Error
}

// 根据UpdatedAt UserToken值批量删除查询
func (store *Store) DelTblUserTokenByUpdatedAtAndUserToken(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userToken).Delete(models.TblUserToken{}).Error
}

// 根据DeletedAt UserName值批量删除查询
func (store *Store) DelTblUserTokenByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Delete(models.TblUserToken{}).Error
}

// 根据DeletedAt UserToken值批量删除查询
func (store *Store) DelTblUserTokenByDeletedAtAndUserToken(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userToken).Delete(models.TblUserToken{}).Error
}

// 根据UserName UserToken值批量删除查询
func (store *Store) DelTblUserTokenByUserNameAndUserToken(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userToken).Delete(models.TblUserToken{}).Error
}
