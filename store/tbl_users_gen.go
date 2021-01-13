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
func (store *Store) CreateTblUser(tblUser *models.TblUser) error {
	return store.db.Create(&tblUser).Error
}

//  删除
func (store *Store) DeleteTblUser(tblUser *models.TblUser) error {
	return store.db.Delete(&tblUser).Error
}

// 更新
func (store *Store) UpdateTblUser(tblUser *models.TblUser) error {
	return store.db.Save(&tblUser).Error
}

// GetBy主键
func (store *Store) GetOneTblUserById(p predicate.Predicate, id uint) (*models.TblUser, error) {
	var model models.TblUser
	return &model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// list全部
func (store *Store) GetTblUser() ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Find(&model).Error
}

// list分页
func (store *Store) GetTblUserByPage(limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值查询所有
func (store *Store) GetTblUserById(p predicate.Predicate, id uint) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// 根据CreatedAt值查询所有
func (store *Store) GetTblUserByCreatedAt(p predicate.Predicate, createdAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Find(&model).Error
}

// 根据UpdatedAt值查询所有
func (store *Store) GetTblUserByUpdatedAt(p predicate.Predicate, updatedAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Find(&model).Error
}

// 根据DeletedAt值查询所有
func (store *Store) GetTblUserByDeletedAt(p predicate.Predicate, deletedAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Find(&model).Error
}

// 根据UserName值查询所有
func (store *Store) GetTblUserByUserName(p predicate.Predicate, userName string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Find(&model).Error
}

// 根据UserPwd值查询所有
func (store *Store) GetTblUserByUserPwd(p predicate.Predicate, userPwd string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ?", p), userPwd).Find(&model).Error
}

// 根据Email值查询所有
func (store *Store) GetTblUserByEmail(p predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ?", p), email).Find(&model).Error
}

// 根据Phone值查询所有
func (store *Store) GetTblUserByPhone(p predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ?", p), phone).Find(&model).Error
}

// 根据EmailValidated值查询所有
func (store *Store) GetTblUserByEmailValidated(p predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ?", p), emailValidated).Find(&model).Error
}

// 根据PhoneValidate值查询所有
func (store *Store) GetTblUserByPhoneValidate(p predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validate %s ?", p), phoneValidate).Find(&model).Error
}

// 根据LastActive值查询所有
func (store *Store) GetTblUserByLastActive(p predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("last_active %s ?", p), lastActive).Find(&model).Error
}

// 根据Profile值查询所有
func (store *Store) GetTblUserByProfile(p predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("profile %s ?", p), profile).Find(&model).Error
}

// 根据Statu值查询所有
func (store *Store) GetTblUserByStatu(p predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Find(&model).Error
}

// 根据Id值分页查询
func (store *Store) GetTblUserByIdByPage(p predicate.Predicate, id uint, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据CreatedAt值分页查询
func (store *Store) GetTblUserByCreatedAtByPage(p predicate.Predicate, createdAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UpdatedAt值分页查询
func (store *Store) GetTblUserByUpdatedAtByPage(p predicate.Predicate, updatedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据DeletedAt值分页查询
func (store *Store) GetTblUserByDeletedAtByPage(p predicate.Predicate, deletedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UserName值分页查询
func (store *Store) GetTblUserByUserNameByPage(p predicate.Predicate, userName string, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UserPwd值分页查询
func (store *Store) GetTblUserByUserPwdByPage(p predicate.Predicate, userPwd string, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ?", p), userPwd).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Email值分页查询
func (store *Store) GetTblUserByEmailByPage(p predicate.Predicate, email string, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ?", p), email).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Phone值分页查询
func (store *Store) GetTblUserByPhoneByPage(p predicate.Predicate, phone string, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ?", p), phone).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据EmailValidated值分页查询
func (store *Store) GetTblUserByEmailValidatedByPage(p predicate.Predicate, emailValidated int, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ?", p), emailValidated).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据PhoneValidate值分页查询
func (store *Store) GetTblUserByPhoneValidateByPage(p predicate.Predicate, phoneValidate int, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validate %s ?", p), phoneValidate).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据LastActive值分页查询
func (store *Store) GetTblUserByLastActiveByPage(p predicate.Predicate, lastActive time.Time, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("last_active %s ?", p), lastActive).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Profile值分页查询
func (store *Store) GetTblUserByProfileByPage(p predicate.Predicate, profile string, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("profile %s ?", p), profile).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Statu值分页查询
func (store *Store) GetTblUserByStatuByPage(p predicate.Predicate, statu int, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值批量删除查询
func (store *Store) DelTblUserById(p predicate.Predicate, id uint) error {
	return store.db.Where(fmt.Sprintf("id %s ?", p), id).Delete(models.TblUser{}).Error
}

// 根据CreatedAt值批量删除查询
func (store *Store) DelTblUserByCreatedAt(p predicate.Predicate, createdAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt值批量删除查询
func (store *Store) DelTblUserByUpdatedAt(p predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Delete(models.TblUser{}).Error
}

// 根据DeletedAt值批量删除查询
func (store *Store) DelTblUserByDeletedAt(p predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Delete(models.TblUser{}).Error
}

// 根据UserName值批量删除查询
func (store *Store) DelTblUserByUserName(p predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Delete(models.TblUser{}).Error
}

// 根据UserPwd值批量删除查询
func (store *Store) DelTblUserByUserPwd(p predicate.Predicate, userPwd string) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ?", p), userPwd).Delete(models.TblUser{}).Error
}

// 根据Email值批量删除查询
func (store *Store) DelTblUserByEmail(p predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("email %s ?", p), email).Delete(models.TblUser{}).Error
}

// 根据Phone值批量删除查询
func (store *Store) DelTblUserByPhone(p predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("phone %s ?", p), phone).Delete(models.TblUser{}).Error
}

// 根据EmailValidated值批量删除查询
func (store *Store) DelTblUserByEmailValidated(p predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ?", p), emailValidated).Delete(models.TblUser{}).Error
}

// 根据PhoneValidate值批量删除查询
func (store *Store) DelTblUserByPhoneValidate(p predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("phone_validate %s ?", p), phoneValidate).Delete(models.TblUser{}).Error
}

// 根据LastActive值批量删除查询
func (store *Store) DelTblUserByLastActive(p predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("last_active %s ?", p), lastActive).Delete(models.TblUser{}).Error
}

// 根据Profile值批量删除查询
func (store *Store) DelTblUserByProfile(p predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("profile %s ?", p), profile).Delete(models.TblUser{}).Error
}

// 根据Statu值批量删除查询
func (store *Store) DelTblUserByStatu(p predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("status %s ?", p), statu).Delete(models.TblUser{}).Error
}

// 根据CreatedAt UpdatedAt值全部查询
func (store *Store) GetTblUserByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Find(&model).Error
}

// 根据CreatedAt DeletedAt值全部查询
func (store *Store) GetTblUserByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Find(&model).Error
}

// 根据CreatedAt UserName值全部查询
func (store *Store) GetTblUserByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Find(&model).Error
}

// 根据CreatedAt UserPwd值全部查询
func (store *Store) GetTblUserByCreatedAtAndUserPwd(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userPwd).Find(&model).Error
}

// 根据CreatedAt Email值全部查询
func (store *Store) GetTblUserByCreatedAtAndEmail(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, email).Find(&model).Error
}

// 根据CreatedAt Phone值全部查询
func (store *Store) GetTblUserByCreatedAtAndPhone(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, phone).Find(&model).Error
}

// 根据CreatedAt EmailValidated值全部查询
func (store *Store) GetTblUserByCreatedAtAndEmailValidated(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, emailValidated).Find(&model).Error
}

// 根据CreatedAt PhoneValidate值全部查询
func (store *Store) GetTblUserByCreatedAtAndPhoneValidate(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, phoneValidate).Find(&model).Error
}

// 根据CreatedAt LastActive值全部查询
func (store *Store) GetTblUserByCreatedAtAndLastActive(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, lastActive).Find(&model).Error
}

// 根据CreatedAt Profile值全部查询
func (store *Store) GetTblUserByCreatedAtAndProfile(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, profile).Find(&model).Error
}

// 根据CreatedAt Statu值全部查询
func (store *Store) GetTblUserByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Find(&model).Error
}

// 根据UpdatedAt DeletedAt值全部查询
func (store *Store) GetTblUserByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Find(&model).Error
}

// 根据UpdatedAt UserName值全部查询
func (store *Store) GetTblUserByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Find(&model).Error
}

// 根据UpdatedAt UserPwd值全部查询
func (store *Store) GetTblUserByUpdatedAtAndUserPwd(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userPwd).Find(&model).Error
}

// 根据UpdatedAt Email值全部查询
func (store *Store) GetTblUserByUpdatedAtAndEmail(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, email).Find(&model).Error
}

// 根据UpdatedAt Phone值全部查询
func (store *Store) GetTblUserByUpdatedAtAndPhone(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, phone).Find(&model).Error
}

// 根据UpdatedAt EmailValidated值全部查询
func (store *Store) GetTblUserByUpdatedAtAndEmailValidated(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, emailValidated).Find(&model).Error
}

// 根据UpdatedAt PhoneValidate值全部查询
func (store *Store) GetTblUserByUpdatedAtAndPhoneValidate(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, phoneValidate).Find(&model).Error
}

// 根据UpdatedAt LastActive值全部查询
func (store *Store) GetTblUserByUpdatedAtAndLastActive(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, lastActive).Find(&model).Error
}

// 根据UpdatedAt Profile值全部查询
func (store *Store) GetTblUserByUpdatedAtAndProfile(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, profile).Find(&model).Error
}

// 根据UpdatedAt Statu值全部查询
func (store *Store) GetTblUserByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Find(&model).Error
}

// 根据DeletedAt UserName值全部查询
func (store *Store) GetTblUserByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Find(&model).Error
}

// 根据DeletedAt UserPwd值全部查询
func (store *Store) GetTblUserByDeletedAtAndUserPwd(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userPwd).Find(&model).Error
}

// 根据DeletedAt Email值全部查询
func (store *Store) GetTblUserByDeletedAtAndEmail(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, email).Find(&model).Error
}

// 根据DeletedAt Phone值全部查询
func (store *Store) GetTblUserByDeletedAtAndPhone(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, phone).Find(&model).Error
}

// 根据DeletedAt EmailValidated值全部查询
func (store *Store) GetTblUserByDeletedAtAndEmailValidated(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, emailValidated).Find(&model).Error
}

// 根据DeletedAt PhoneValidate值全部查询
func (store *Store) GetTblUserByDeletedAtAndPhoneValidate(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, phoneValidate).Find(&model).Error
}

// 根据DeletedAt LastActive值全部查询
func (store *Store) GetTblUserByDeletedAtAndLastActive(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, lastActive).Find(&model).Error
}

// 根据DeletedAt Profile值全部查询
func (store *Store) GetTblUserByDeletedAtAndProfile(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, profile).Find(&model).Error
}

// 根据DeletedAt Statu值全部查询
func (store *Store) GetTblUserByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Find(&model).Error
}

// 根据UserName UserPwd值全部查询
func (store *Store) GetTblUserByUserNameAndUserPwd(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userPwd).Find(&model).Error
}

// 根据UserName Email值全部查询
func (store *Store) GetTblUserByUserNameAndEmail(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, email).Find(&model).Error
}

// 根据UserName Phone值全部查询
func (store *Store) GetTblUserByUserNameAndPhone(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phone).Find(&model).Error
}

// 根据UserName EmailValidated值全部查询
func (store *Store) GetTblUserByUserNameAndEmailValidated(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, emailValidated).Find(&model).Error
}

// 根据UserName PhoneValidate值全部查询
func (store *Store) GetTblUserByUserNameAndPhoneValidate(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phoneValidate).Find(&model).Error
}

// 根据UserName LastActive值全部查询
func (store *Store) GetTblUserByUserNameAndLastActive(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, lastActive).Find(&model).Error
}

// 根据UserName Profile值全部查询
func (store *Store) GetTblUserByUserNameAndProfile(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, profile).Find(&model).Error
}

// 根据UserName Statu值全部查询
func (store *Store) GetTblUserByUserNameAndStatu(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, statu).Find(&model).Error
}

// 根据UserPwd Email值全部查询
func (store *Store) GetTblUserByUserPwdAndEmail(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, email string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, email).Find(&model).Error
}

// 根据UserPwd Phone值全部查询
func (store *Store) GetTblUserByUserPwdAndPhone(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phone).Find(&model).Error
}

// 根据UserPwd EmailValidated值全部查询
func (store *Store) GetTblUserByUserPwdAndEmailValidated(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, emailValidated).Find(&model).Error
}

// 根据UserPwd PhoneValidate值全部查询
func (store *Store) GetTblUserByUserPwdAndPhoneValidate(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phoneValidate).Find(&model).Error
}

// 根据UserPwd LastActive值全部查询
func (store *Store) GetTblUserByUserPwdAndLastActive(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, lastActive).Find(&model).Error
}

// 根据UserPwd Profile值全部查询
func (store *Store) GetTblUserByUserPwdAndProfile(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, profile).Find(&model).Error
}

// 根据UserPwd Statu值全部查询
func (store *Store) GetTblUserByUserPwdAndStatu(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, statu).Find(&model).Error
}

// 根据Email Phone值全部查询
func (store *Store) GetTblUserByEmailAndPhone(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phone string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phone).Find(&model).Error
}

// 根据Email EmailValidated值全部查询
func (store *Store) GetTblUserByEmailAndEmailValidated(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, emailValidated).Find(&model).Error
}

// 根据Email PhoneValidate值全部查询
func (store *Store) GetTblUserByEmailAndPhoneValidate(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phoneValidate).Find(&model).Error
}

// 根据Email LastActive值全部查询
func (store *Store) GetTblUserByEmailAndLastActive(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, lastActive).Find(&model).Error
}

// 根据Email Profile值全部查询
func (store *Store) GetTblUserByEmailAndProfile(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, profile).Find(&model).Error
}

// 根据Email Statu值全部查询
func (store *Store) GetTblUserByEmailAndStatu(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, statu).Find(&model).Error
}

// 根据Phone EmailValidated值全部查询
func (store *Store) GetTblUserByPhoneAndEmailValidated(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, emailValidated).Find(&model).Error
}

// 根据Phone PhoneValidate值全部查询
func (store *Store) GetTblUserByPhoneAndPhoneValidate(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, phoneValidate).Find(&model).Error
}

// 根据Phone LastActive值全部查询
func (store *Store) GetTblUserByPhoneAndLastActive(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, lastActive).Find(&model).Error
}

// 根据Phone Profile值全部查询
func (store *Store) GetTblUserByPhoneAndProfile(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, profile).Find(&model).Error
}

// 根据Phone Statu值全部查询
func (store *Store) GetTblUserByPhoneAndStatu(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, statu).Find(&model).Error
}

// 根据EmailValidated PhoneValidate值全部查询
func (store *Store) GetTblUserByEmailValidatedAndPhoneValidate(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, phoneValidate).Find(&model).Error
}

// 根据EmailValidated LastActive值全部查询
func (store *Store) GetTblUserByEmailValidatedAndLastActive(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, lastActive).Find(&model).Error
}

// 根据EmailValidated Profile值全部查询
func (store *Store) GetTblUserByEmailValidatedAndProfile(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, profile).Find(&model).Error
}

// 根据EmailValidated Statu值全部查询
func (store *Store) GetTblUserByEmailValidatedAndStatu(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, statu).Find(&model).Error
}

// 根据PhoneValidate LastActive值全部查询
func (store *Store) GetTblUserByPhoneValidateAndLastActive(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, lastActive).Find(&model).Error
}

// 根据PhoneValidate Profile值全部查询
func (store *Store) GetTblUserByPhoneValidateAndProfile(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, profile).Find(&model).Error
}

// 根据PhoneValidate Statu值全部查询
func (store *Store) GetTblUserByPhoneValidateAndStatu(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, statu).Find(&model).Error
}

// 根据LastActive Profile值全部查询
func (store *Store) GetTblUserByLastActiveAndProfile(p predicate.Predicate, lastActive time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("last_active %s ? %s  %s ?", p, c, p2), lastActive, profile).Find(&model).Error
}

// 根据LastActive Statu值全部查询
func (store *Store) GetTblUserByLastActiveAndStatu(p predicate.Predicate, lastActive time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("last_active %s ? %s  %s ?", p, c, p2), lastActive, statu).Find(&model).Error
}

// 根据Profile Statu值全部查询
func (store *Store) GetTblUserByProfileAndStatu(p predicate.Predicate, profile string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("profile %s ? %s  %s ?", p, c, p2), profile, statu).Find(&model).Error
}

// 根据CreatedAt UpdatedAt值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Delete(models.TblUser{}).Error
}

// 根据CreatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Delete(models.TblUser{}).Error
}

// 根据CreatedAt UserName值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Delete(models.TblUser{}).Error
}

// 根据CreatedAt UserPwd值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndUserPwd(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userPwd).Delete(models.TblUser{}).Error
}

// 根据CreatedAt Email值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndEmail(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, email).Delete(models.TblUser{}).Error
}

// 根据CreatedAt Phone值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndPhone(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, phone).Delete(models.TblUser{}).Error
}

// 根据CreatedAt EmailValidated值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndEmailValidated(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, emailValidated).Delete(models.TblUser{}).Error
}

// 根据CreatedAt PhoneValidate值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndPhoneValidate(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据CreatedAt LastActive值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndLastActive(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, lastActive).Delete(models.TblUser{}).Error
}

// 根据CreatedAt Profile值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndProfile(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, profile).Delete(models.TblUser{}).Error
}

// 根据CreatedAt Statu值批量删除查询
func (store *Store) DelTblUserByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt UserName值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt UserPwd值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndUserPwd(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userPwd).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt Email值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndEmail(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, email).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt Phone值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndPhone(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, phone).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt EmailValidated值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndEmailValidated(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, emailValidated).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt PhoneValidate值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndPhoneValidate(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt LastActive值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndLastActive(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, lastActive).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt Profile值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndProfile(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, profile).Delete(models.TblUser{}).Error
}

// 根据UpdatedAt Statu值批量删除查询
func (store *Store) DelTblUserByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Delete(models.TblUser{}).Error
}

// 根据DeletedAt UserName值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Delete(models.TblUser{}).Error
}

// 根据DeletedAt UserPwd值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndUserPwd(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userPwd).Delete(models.TblUser{}).Error
}

// 根据DeletedAt Email值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndEmail(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, email).Delete(models.TblUser{}).Error
}

// 根据DeletedAt Phone值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndPhone(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, phone).Delete(models.TblUser{}).Error
}

// 根据DeletedAt EmailValidated值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndEmailValidated(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, emailValidated).Delete(models.TblUser{}).Error
}

// 根据DeletedAt PhoneValidate值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndPhoneValidate(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据DeletedAt LastActive值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndLastActive(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, lastActive).Delete(models.TblUser{}).Error
}

// 根据DeletedAt Profile值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndProfile(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, profile).Delete(models.TblUser{}).Error
}

// 根据DeletedAt Statu值批量删除查询
func (store *Store) DelTblUserByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Delete(models.TblUser{}).Error
}

// 根据UserName UserPwd值批量删除查询
func (store *Store) DelTblUserByUserNameAndUserPwd(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, userPwd string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, userPwd).Delete(models.TblUser{}).Error
}

// 根据UserName Email值批量删除查询
func (store *Store) DelTblUserByUserNameAndEmail(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, email).Delete(models.TblUser{}).Error
}

// 根据UserName Phone值批量删除查询
func (store *Store) DelTblUserByUserNameAndPhone(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phone).Delete(models.TblUser{}).Error
}

// 根据UserName EmailValidated值批量删除查询
func (store *Store) DelTblUserByUserNameAndEmailValidated(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, emailValidated).Delete(models.TblUser{}).Error
}

// 根据UserName PhoneValidate值批量删除查询
func (store *Store) DelTblUserByUserNameAndPhoneValidate(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据UserName LastActive值批量删除查询
func (store *Store) DelTblUserByUserNameAndLastActive(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, lastActive).Delete(models.TblUser{}).Error
}

// 根据UserName Profile值批量删除查询
func (store *Store) DelTblUserByUserNameAndProfile(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, profile).Delete(models.TblUser{}).Error
}

// 根据UserName Statu值批量删除查询
func (store *Store) DelTblUserByUserNameAndStatu(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, statu).Delete(models.TblUser{}).Error
}

// 根据UserPwd Email值批量删除查询
func (store *Store) DelTblUserByUserPwdAndEmail(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, email string) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, email).Delete(models.TblUser{}).Error
}

// 根据UserPwd Phone值批量删除查询
func (store *Store) DelTblUserByUserPwdAndPhone(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phone).Delete(models.TblUser{}).Error
}

// 根据UserPwd EmailValidated值批量删除查询
func (store *Store) DelTblUserByUserPwdAndEmailValidated(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, emailValidated).Delete(models.TblUser{}).Error
}

// 根据UserPwd PhoneValidate值批量删除查询
func (store *Store) DelTblUserByUserPwdAndPhoneValidate(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据UserPwd LastActive值批量删除查询
func (store *Store) DelTblUserByUserPwdAndLastActive(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, lastActive).Delete(models.TblUser{}).Error
}

// 根据UserPwd Profile值批量删除查询
func (store *Store) DelTblUserByUserPwdAndProfile(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, profile).Delete(models.TblUser{}).Error
}

// 根据UserPwd Statu值批量删除查询
func (store *Store) DelTblUserByUserPwdAndStatu(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, statu).Delete(models.TblUser{}).Error
}

// 根据Email Phone值批量删除查询
func (store *Store) DelTblUserByEmailAndPhone(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phone string) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phone).Delete(models.TblUser{}).Error
}

// 根据Email EmailValidated值批量删除查询
func (store *Store) DelTblUserByEmailAndEmailValidated(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, emailValidated).Delete(models.TblUser{}).Error
}

// 根据Email PhoneValidate值批量删除查询
func (store *Store) DelTblUserByEmailAndPhoneValidate(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据Email LastActive值批量删除查询
func (store *Store) DelTblUserByEmailAndLastActive(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, lastActive).Delete(models.TblUser{}).Error
}

// 根据Email Profile值批量删除查询
func (store *Store) DelTblUserByEmailAndProfile(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, profile).Delete(models.TblUser{}).Error
}

// 根据Email Statu值批量删除查询
func (store *Store) DelTblUserByEmailAndStatu(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, statu).Delete(models.TblUser{}).Error
}

// 根据Phone EmailValidated值批量删除查询
func (store *Store) DelTblUserByPhoneAndEmailValidated(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, emailValidated int) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, emailValidated).Delete(models.TblUser{}).Error
}

// 根据Phone PhoneValidate值批量删除查询
func (store *Store) DelTblUserByPhoneAndPhoneValidate(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据Phone LastActive值批量删除查询
func (store *Store) DelTblUserByPhoneAndLastActive(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, lastActive).Delete(models.TblUser{}).Error
}

// 根据Phone Profile值批量删除查询
func (store *Store) DelTblUserByPhoneAndProfile(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, profile).Delete(models.TblUser{}).Error
}

// 根据Phone Statu值批量删除查询
func (store *Store) DelTblUserByPhoneAndStatu(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, statu).Delete(models.TblUser{}).Error
}

// 根据EmailValidated PhoneValidate值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndPhoneValidate(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, phoneValidate int) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, phoneValidate).Delete(models.TblUser{}).Error
}

// 根据EmailValidated LastActive值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndLastActive(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, lastActive).Delete(models.TblUser{}).Error
}

// 根据EmailValidated Profile值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndProfile(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, profile).Delete(models.TblUser{}).Error
}

// 根据EmailValidated Statu值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndStatu(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, statu).Delete(models.TblUser{}).Error
}

// 根据PhoneValidate LastActive值批量删除查询
func (store *Store) DelTblUserByPhoneValidateAndLastActive(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, lastActive).Delete(models.TblUser{}).Error
}

// 根据PhoneValidate Profile值批量删除查询
func (store *Store) DelTblUserByPhoneValidateAndProfile(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, profile).Delete(models.TblUser{}).Error
}

// 根据PhoneValidate Statu值批量删除查询
func (store *Store) DelTblUserByPhoneValidateAndStatu(p predicate.Predicate, phoneValidate int, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("phone_validate %s ? %s  %s ?", p, c, p2), phoneValidate, statu).Delete(models.TblUser{}).Error
}

// 根据LastActive Profile值批量删除查询
func (store *Store) DelTblUserByLastActiveAndProfile(p predicate.Predicate, lastActive time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("last_active %s ? %s  %s ?", p, c, p2), lastActive, profile).Delete(models.TblUser{}).Error
}

// 根据LastActive Statu值批量删除查询
func (store *Store) DelTblUserByLastActiveAndStatu(p predicate.Predicate, lastActive time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("last_active %s ? %s  %s ?", p, c, p2), lastActive, statu).Delete(models.TblUser{}).Error
}

// 根据Profile Statu值批量删除查询
func (store *Store) DelTblUserByProfileAndStatu(p predicate.Predicate, profile string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("profile %s ? %s  %s ?", p, c, p2), profile, statu).Delete(models.TblUser{}).Error
}
