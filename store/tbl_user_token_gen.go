package store

////////////////////////////////////////////////////////////////////
//   Warning: this file is auto, please not edit.  自动生成请勿编辑
//   Author : ftfeng
//   主要包含：1、创建/删除/更新实例  2、更新实例 3、根据列获取所有 3、根据列分页获取
//   Create time : 2021-01-13 00:50:13
////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"github.com/feitianlove/FIleStore/models"
	"github.com/feitianlove/golib/common/predicate"
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

// 根据UserName值批量删除查询
func (store *Store) DelTblUserTokenByUserName(p predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Delete(models.TblUserToken{}).Error
}

// 根据UserToken值批量删除查询
func (store *Store) DelTblUserTokenByUserToken(p predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("user_token %s ?", p), userToken).Delete(models.TblUserToken{}).Error
}

// 根据UserName UserToken值全部查询
func (store *Store) GetTblUserTokenByUserNameAndUserToken(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userToken string) ([]*models.TblUserToken, error) {
	var model []*models.TblUserToken
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userToken).Find(&model).Error
}

// 根据UserName UserToken值批量删除查询
func (store *Store) DelTblUserTokenByUserNameAndUserToken(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userToken string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userToken).Delete(models.TblUserToken{}).Error
}
