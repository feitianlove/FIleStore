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
func (store *Store) CreateTblUserFile(tblUserFile *models.TblUserFile) error {
	return store.db.Create(&tblUserFile).Error
}

//  删除
func (store *Store) DeleteTblUserFile(tblUserFile *models.TblUserFile) error {
	return store.db.Delete(&tblUserFile).Error
}

// 更新
func (store *Store) UpdateTblUserFile(tblUserFile *models.TblUserFile) error {
	return store.db.Save(&tblUserFile).Error
}

// GetBy主键
func (store *Store) GetOneTblUserFileById(p predicate.Predicate, id uint) (*models.TblUserFile, error) {
	var model models.TblUserFile
	return &model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// list全部
func (store *Store) GetTblUserFile() ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Find(&model).Error
}

// list分页
func (store *Store) GetTblUserFileByPage(limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值查询所有
func (store *Store) GetTblUserFileById(p predicate.Predicate, id uint) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// 根据CreatedAt值查询所有
func (store *Store) GetTblUserFileByCreatedAt(p predicate.Predicate, createdAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Find(&model).Error
}

// 根据UpdatedAt值查询所有
func (store *Store) GetTblUserFileByUpdatedAt(p predicate.Predicate, updatedAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Find(&model).Error
}

// 根据DeletedAt值查询所有
func (store *Store) GetTblUserFileByDeletedAt(p predicate.Predicate, deletedAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Find(&model).Error
}

// 根据UserName值查询所有
func (store *Store) GetTblUserFileByUserName(p predicate.Predicate, userName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Find(&model).Error
}

// 根据FileName值查询所有
func (store *Store) GetTblUserFileByFileName(p predicate.Predicate, fileName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Find(&model).Error
}

// 根据FileSha1值查询所有
func (store *Store) GetTblUserFileByFileSha1(p predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Find(&model).Error
}

// 根据FileSize值查询所有
func (store *Store) GetTblUserFileByFileSize(p predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Find(&model).Error
}

// 根据Statu值查询所有
func (store *Store) GetTblUserFileByStatu(p predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Find(&model).Error
}

// 根据Id值分页查询
func (store *Store) GetTblUserFileByIdByPage(p predicate.Predicate, id uint, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据CreatedAt值分页查询
func (store *Store) GetTblUserFileByCreatedAtByPage(p predicate.Predicate, createdAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UpdatedAt值分页查询
func (store *Store) GetTblUserFileByUpdatedAtByPage(p predicate.Predicate, updatedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据DeletedAt值分页查询
func (store *Store) GetTblUserFileByDeletedAtByPage(p predicate.Predicate, deletedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UserName值分页查询
func (store *Store) GetTblUserFileByUserNameByPage(p predicate.Predicate, userName string, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileName值分页查询
func (store *Store) GetTblUserFileByFileNameByPage(p predicate.Predicate, fileName string, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileSha1值分页查询
func (store *Store) GetTblUserFileByFileSha1ByPage(p predicate.Predicate, fileSha1 string, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileSize值分页查询
func (store *Store) GetTblUserFileByFileSizeByPage(p predicate.Predicate, fileSize string, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Statu值分页查询
func (store *Store) GetTblUserFileByStatuByPage(p predicate.Predicate, statu int, limit interface{}, offset interface{}) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值批量删除查询
func (store *Store) DelTblUserFileById(p predicate.Predicate, id uint) error {
	return store.db.Where(fmt.Sprintf("id %s ?", p), id).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt值批量删除查询
func (store *Store) DelTblUserFileByCreatedAt(p predicate.Predicate, createdAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAt(p predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt值批量删除查询
func (store *Store) DelTblUserFileByDeletedAt(p predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Delete(models.TblUserFile{}).Error
}

// 根据UserName值批量删除查询
func (store *Store) DelTblUserFileByUserName(p predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ?", p), userName).Delete(models.TblUserFile{}).Error
}

// 根据FileName值批量删除查询
func (store *Store) DelTblUserFileByFileName(p predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Delete(models.TblUserFile{}).Error
}

// 根据FileSha1值批量删除查询
func (store *Store) DelTblUserFileByFileSha1(p predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据FileSize值批量删除查询
func (store *Store) DelTblUserFileByFileSize(p predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Delete(models.TblUserFile{}).Error
}

// 根据Statu值批量删除查询
func (store *Store) DelTblUserFileByStatu(p predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("status %s ?", p), statu).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt UpdatedAt值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Find(&model).Error
}

// 根据CreatedAt DeletedAt值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Find(&model).Error
}

// 根据CreatedAt UserName值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Find(&model).Error
}

// 根据CreatedAt FileName值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndFileName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileName).Find(&model).Error
}

// 根据CreatedAt FileSha1值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndFileSha1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSha1).Find(&model).Error
}

// 根据CreatedAt FileSize值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndFileSize(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSize).Find(&model).Error
}

// 根据CreatedAt Statu值全部查询
func (store *Store) GetTblUserFileByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Find(&model).Error
}

// 根据UpdatedAt DeletedAt值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Find(&model).Error
}

// 根据UpdatedAt UserName值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Find(&model).Error
}

// 根据UpdatedAt FileName值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndFileName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileName).Find(&model).Error
}

// 根据UpdatedAt FileSha1值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndFileSha1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSha1).Find(&model).Error
}

// 根据UpdatedAt FileSize值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndFileSize(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSize).Find(&model).Error
}

// 根据UpdatedAt Statu值全部查询
func (store *Store) GetTblUserFileByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Find(&model).Error
}

// 根据DeletedAt UserName值全部查询
func (store *Store) GetTblUserFileByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Find(&model).Error
}

// 根据DeletedAt FileName值全部查询
func (store *Store) GetTblUserFileByDeletedAtAndFileName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileName).Find(&model).Error
}

// 根据DeletedAt FileSha1值全部查询
func (store *Store) GetTblUserFileByDeletedAtAndFileSha1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSha1).Find(&model).Error
}

// 根据DeletedAt FileSize值全部查询
func (store *Store) GetTblUserFileByDeletedAtAndFileSize(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSize).Find(&model).Error
}

// 根据DeletedAt Statu值全部查询
func (store *Store) GetTblUserFileByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Find(&model).Error
}

// 根据UserName FileName值全部查询
func (store *Store) GetTblUserFileByUserNameAndFileName(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileName).Find(&model).Error
}

// 根据UserName FileSha1值全部查询
func (store *Store) GetTblUserFileByUserNameAndFileSha1(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileSha1).Find(&model).Error
}

// 根据UserName FileSize值全部查询
func (store *Store) GetTblUserFileByUserNameAndFileSize(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileSize).Find(&model).Error
}

// 根据UserName Statu值全部查询
func (store *Store) GetTblUserFileByUserNameAndStatu(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, statu).Find(&model).Error
}

// 根据FileName FileSha1值全部查询
func (store *Store) GetTblUserFileByFileNameAndFileSha1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSha1).Find(&model).Error
}

// 根据FileName FileSize值全部查询
func (store *Store) GetTblUserFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Find(&model).Error
}

// 根据FileName Statu值全部查询
func (store *Store) GetTblUserFileByFileNameAndStatu(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, statu).Find(&model).Error
}

// 根据FileSha1 FileSize值全部查询
func (store *Store) GetTblUserFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Find(&model).Error
}

// 根据FileSha1 Statu值全部查询
func (store *Store) GetTblUserFileByFileSha1AndStatu(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, statu).Find(&model).Error
}

// 根据FileSize Statu值全部查询
func (store *Store) GetTblUserFileByFileSizeAndStatu(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblUserFile, error) {
	var model []*models.TblUserFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Find(&model).Error
}

// 根据CreatedAt UpdatedAt值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt UserName值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndUserName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, userName).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt FileName值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndFileName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileName).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt FileSha1值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndFileSha1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt FileSize值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndFileSize(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据CreatedAt Statu值批量删除查询
func (store *Store) DelTblUserFileByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt DeletedAt值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt UserName值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndUserName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, userName).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt FileName值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndFileName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileName).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt FileSha1值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndFileSha1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt FileSize值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndFileSize(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据UpdatedAt Statu值批量删除查询
func (store *Store) DelTblUserFileByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt UserName值批量删除查询
func (store *Store) DelTblUserFileByDeletedAtAndUserName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, userName string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, userName).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt FileName值批量删除查询
func (store *Store) DelTblUserFileByDeletedAtAndFileName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileName).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt FileSha1值批量删除查询
func (store *Store) DelTblUserFileByDeletedAtAndFileSha1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt FileSize值批量删除查询
func (store *Store) DelTblUserFileByDeletedAtAndFileSize(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据DeletedAt Statu值批量删除查询
func (store *Store) DelTblUserFileByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Delete(models.TblUserFile{}).Error
}

// 根据UserName FileName值批量删除查询
func (store *Store) DelTblUserFileByUserNameAndFileName(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileName).Delete(models.TblUserFile{}).Error
}

// 根据UserName FileSha1值批量删除查询
func (store *Store) DelTblUserFileByUserNameAndFileSha1(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据UserName FileSize值批量删除查询
func (store *Store) DelTblUserFileByUserNameAndFileSize(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据UserName Statu值批量删除查询
func (store *Store) DelTblUserFileByUserNameAndStatu(p predicate.Predicate, userName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("user_name %s ? %s  %s ?", p, c, p2), userName, statu).Delete(models.TblUserFile{}).Error
}

// 根据FileName FileSha1值批量删除查询
func (store *Store) DelTblUserFileByFileNameAndFileSha1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSha1).Delete(models.TblUserFile{}).Error
}

// 根据FileName FileSize值批量删除查询
func (store *Store) DelTblUserFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据FileName Statu值批量删除查询
func (store *Store) DelTblUserFileByFileNameAndStatu(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, statu).Delete(models.TblUserFile{}).Error
}

// 根据FileSha1 FileSize值批量删除查询
func (store *Store) DelTblUserFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Delete(models.TblUserFile{}).Error
}

// 根据FileSha1 Statu值批量删除查询
func (store *Store) DelTblUserFileByFileSha1AndStatu(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, statu).Delete(models.TblUserFile{}).Error
}

// 根据FileSize Statu值批量删除查询
func (store *Store) DelTblUserFileByFileSizeAndStatu(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Delete(models.TblUserFile{}).Error
}
