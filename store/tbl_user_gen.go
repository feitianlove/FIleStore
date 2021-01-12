package store

////////////////////////////////////////////////////////////////////
//   Warning: this file is auto, please not edit.  自动生成请勿编辑
//   Author : ftfeng
//   主要包含：1、创建/删除/更新实例  2、更新实例 3、根据列获取所有 3、根据列分页获取
//   Create time : 2021-01-12 23:53:05
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

// 根据PhoneValidated值查询所有
func (store *Store) GetTblUserByPhoneValidated(p predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ?", p), phoneValidated).Find(&model).Error
}

// 根据SignupAt值查询所有
func (store *Store) GetTblUserBySignupAt(p predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("signup_at %s ?", p), signupAt).Find(&model).Error
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

// 根据PhoneValidated值分页查询
func (store *Store) GetTblUserByPhoneValidatedByPage(p predicate.Predicate, phoneValidated int, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ?", p), phoneValidated).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据SignupAt值分页查询
func (store *Store) GetTblUserBySignupAtByPage(p predicate.Predicate, signupAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("signup_at %s ?", p), signupAt).Limit(limit).Offset(offset).Find(&model).Error
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

// 根据PhoneValidated值批量删除查询
func (store *Store) DelTblUserByPhoneValidated(p predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("phone_validated %s ?", p), phoneValidated).Delete(models.TblUser{}).Error
}

// 根据SignupAt值批量删除查询
func (store *Store) DelTblUserBySignupAt(p predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("signup_at %s ?", p), signupAt).Delete(models.TblUser{}).Error
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

// 根据UserName PhoneValidated值全部查询
func (store *Store) GetTblUserByUserNameAndPhoneValidated(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phoneValidated).Find(&model).Error
}

// 根据UserName SignupAt值全部查询
func (store *Store) GetTblUserByUserNameAndSignupAt(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, signupAt).Find(&model).Error
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

// 根据UserPwd PhoneValidated值全部查询
func (store *Store) GetTblUserByUserPwdAndPhoneValidated(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phoneValidated).Find(&model).Error
}

// 根据UserPwd SignupAt值全部查询
func (store *Store) GetTblUserByUserPwdAndSignupAt(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, signupAt).Find(&model).Error
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

// 根据Email PhoneValidated值全部查询
func (store *Store) GetTblUserByEmailAndPhoneValidated(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phoneValidated).Find(&model).Error
}

// 根据Email SignupAt值全部查询
func (store *Store) GetTblUserByEmailAndSignupAt(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, signupAt).Find(&model).Error
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

// 根据Phone PhoneValidated值全部查询
func (store *Store) GetTblUserByPhoneAndPhoneValidated(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, phoneValidated).Find(&model).Error
}

// 根据Phone SignupAt值全部查询
func (store *Store) GetTblUserByPhoneAndSignupAt(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, signupAt).Find(&model).Error
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

// 根据EmailValidated PhoneValidated值全部查询
func (store *Store) GetTblUserByEmailValidatedAndPhoneValidated(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, phoneValidated).Find(&model).Error
}

// 根据EmailValidated SignupAt值全部查询
func (store *Store) GetTblUserByEmailValidatedAndSignupAt(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, signupAt).Find(&model).Error
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

// 根据PhoneValidated SignupAt值全部查询
func (store *Store) GetTblUserByPhoneValidatedAndSignupAt(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, signupAt).Find(&model).Error
}

// 根据PhoneValidated LastActive值全部查询
func (store *Store) GetTblUserByPhoneValidatedAndLastActive(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, lastActive).Find(&model).Error
}

// 根据PhoneValidated Profile值全部查询
func (store *Store) GetTblUserByPhoneValidatedAndProfile(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, profile).Find(&model).Error
}

// 根据PhoneValidated Statu值全部查询
func (store *Store) GetTblUserByPhoneValidatedAndStatu(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, statu).Find(&model).Error
}

// 根据SignupAt LastActive值全部查询
func (store *Store) GetTblUserBySignupAtAndLastActive(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, lastActive).Find(&model).Error
}

// 根据SignupAt Profile值全部查询
func (store *Store) GetTblUserBySignupAtAndProfile(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, profile).Find(&model).Error
}

// 根据SignupAt Statu值全部查询
func (store *Store) GetTblUserBySignupAtAndStatu(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUser, error) {
	var model []*models.TblUser
	return model, store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, statu).Find(&model).Error
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

// 根据UserName PhoneValidated值批量删除查询
func (store *Store) DelTblUserByUserNameAndPhoneValidated(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, phoneValidated).Delete(models.TblUser{}).Error
}

// 根据UserName SignupAt值批量删除查询
func (store *Store) DelTblUserByUserNameAndSignupAt(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, signupAt).Delete(models.TblUser{}).Error
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

// 根据UserPwd PhoneValidated值批量删除查询
func (store *Store) DelTblUserByUserPwdAndPhoneValidated(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, phoneValidated).Delete(models.TblUser{}).Error
}

// 根据UserPwd SignupAt值批量删除查询
func (store *Store) DelTblUserByUserPwdAndSignupAt(p predicate.Predicate, userPwd string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("user_pwd %s ? %s  %s ?", p, c, p2), userPwd, signupAt).Delete(models.TblUser{}).Error
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

// 根据Email PhoneValidated值批量删除查询
func (store *Store) DelTblUserByEmailAndPhoneValidated(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, phoneValidated).Delete(models.TblUser{}).Error
}

// 根据Email SignupAt值批量删除查询
func (store *Store) DelTblUserByEmailAndSignupAt(p predicate.Predicate, email string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("email %s ? %s  %s ?", p, c, p2), email, signupAt).Delete(models.TblUser{}).Error
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

// 根据Phone PhoneValidated值批量删除查询
func (store *Store) DelTblUserByPhoneAndPhoneValidated(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, phoneValidated).Delete(models.TblUser{}).Error
}

// 根据Phone SignupAt值批量删除查询
func (store *Store) DelTblUserByPhoneAndSignupAt(p predicate.Predicate, phone string, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("phone %s ? %s  %s ?", p, c, p2), phone, signupAt).Delete(models.TblUser{}).Error
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

// 根据EmailValidated PhoneValidated值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndPhoneValidated(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, phoneValidated int) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, phoneValidated).Delete(models.TblUser{}).Error
}

// 根据EmailValidated SignupAt值批量删除查询
func (store *Store) DelTblUserByEmailValidatedAndSignupAt(p predicate.Predicate, emailValidated int, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("email_validated %s ? %s  %s ?", p, c, p2), emailValidated, signupAt).Delete(models.TblUser{}).Error
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

// 根据PhoneValidated SignupAt值批量删除查询
func (store *Store) DelTblUserByPhoneValidatedAndSignupAt(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, signupAt time.Time) error {
	return store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, signupAt).Delete(models.TblUser{}).Error
}

// 根据PhoneValidated LastActive值批量删除查询
func (store *Store) DelTblUserByPhoneValidatedAndLastActive(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, lastActive).Delete(models.TblUser{}).Error
}

// 根据PhoneValidated Profile值批量删除查询
func (store *Store) DelTblUserByPhoneValidatedAndProfile(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, profile).Delete(models.TblUser{}).Error
}

// 根据PhoneValidated Statu值批量删除查询
func (store *Store) DelTblUserByPhoneValidatedAndStatu(p predicate.Predicate, phoneValidated int, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("phone_validated %s ? %s  %s ?", p, c, p2), phoneValidated, statu).Delete(models.TblUser{}).Error
}

// 根据SignupAt LastActive值批量删除查询
func (store *Store) DelTblUserBySignupAtAndLastActive(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, lastActive time.Time) error {
	return store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, lastActive).Delete(models.TblUser{}).Error
}

// 根据SignupAt Profile值批量删除查询
func (store *Store) DelTblUserBySignupAtAndProfile(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, profile string) error {
	return store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, profile).Delete(models.TblUser{}).Error
}

// 根据SignupAt Statu值批量删除查询
func (store *Store) DelTblUserBySignupAtAndStatu(p predicate.Predicate, signupAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("signup_at %s ? %s  %s ?", p, c, p2), signupAt, statu).Delete(models.TblUser{}).Error
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
