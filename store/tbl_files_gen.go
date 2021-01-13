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
func (store *Store) CreateTblFile(tblFile *models.TblFile) error {
	return store.db.Create(&tblFile).Error
}

//  删除
func (store *Store) DeleteTblFile(tblFile *models.TblFile) error {
	return store.db.Delete(&tblFile).Error
}

// 更新
func (store *Store) UpdateTblFile(tblFile *models.TblFile) error {
	return store.db.Save(&tblFile).Error
}

// GetBy主键
func (store *Store) GetOneTblFileById(p predicate.Predicate, id uint) (*models.TblFile, error) {
	var model models.TblFile
	return &model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// list全部
func (store *Store) GetTblFile() ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Find(&model).Error
}

// list分页
func (store *Store) GetTblFileByPage(limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值查询所有
func (store *Store) GetTblFileById(p predicate.Predicate, id uint) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Find(&model).Error
}

// 根据CreatedAt值查询所有
func (store *Store) GetTblFileByCreatedAt(p predicate.Predicate, createdAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Find(&model).Error
}

// 根据UpdatedAt值查询所有
func (store *Store) GetTblFileByUpdatedAt(p predicate.Predicate, updatedAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Find(&model).Error
}

// 根据DeletedAt值查询所有
func (store *Store) GetTblFileByDeletedAt(p predicate.Predicate, deletedAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Find(&model).Error
}

// 根据FileName值查询所有
func (store *Store) GetTblFileByFileName(p predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Find(&model).Error
}

// 根据FileSha1值查询所有
func (store *Store) GetTblFileByFileSha1(p predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Find(&model).Error
}

// 根据FileSize值查询所有
func (store *Store) GetTblFileByFileSize(p predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Find(&model).Error
}

// 根据FileAddr值查询所有
func (store *Store) GetTblFileByFileAddr(p predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Find(&model).Error
}

// 根据Statu值查询所有
func (store *Store) GetTblFileByStatu(p predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Find(&model).Error
}

// 根据Ext1值查询所有
func (store *Store) GetTblFileByExt1(p predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("ext1 %s ?", p), ext1).Find(&model).Error
}

// 根据Ext2值查询所有
func (store *Store) GetTblFileByExt2(p predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("ext2 %s ?", p), ext2).Find(&model).Error
}

// 根据Id值分页查询
func (store *Store) GetTblFileByIdByPage(p predicate.Predicate, id uint, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("id %s ?", p), id).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据CreatedAt值分页查询
func (store *Store) GetTblFileByCreatedAtByPage(p predicate.Predicate, createdAt time.Time, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UpdatedAt值分页查询
func (store *Store) GetTblFileByUpdatedAtByPage(p predicate.Predicate, updatedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据DeletedAt值分页查询
func (store *Store) GetTblFileByDeletedAtByPage(p predicate.Predicate, deletedAt time.Time, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileName值分页查询
func (store *Store) GetTblFileByFileNameByPage(p predicate.Predicate, fileName string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileSha1值分页查询
func (store *Store) GetTblFileByFileSha1ByPage(p predicate.Predicate, fileSha1 string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileSize值分页查询
func (store *Store) GetTblFileByFileSizeByPage(p predicate.Predicate, fileSize string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileAddr值分页查询
func (store *Store) GetTblFileByFileAddrByPage(p predicate.Predicate, fileAddr string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Statu值分页查询
func (store *Store) GetTblFileByStatuByPage(p predicate.Predicate, statu int, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("status %s ?", p), statu).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Ext1值分页查询
func (store *Store) GetTblFileByExt1ByPage(p predicate.Predicate, ext1 int, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("ext1 %s ?", p), ext1).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Ext2值分页查询
func (store *Store) GetTblFileByExt2ByPage(p predicate.Predicate, ext2 string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("ext2 %s ?", p), ext2).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据Id值批量删除查询
func (store *Store) DelTblFileById(p predicate.Predicate, id uint) error {
	return store.db.Where(fmt.Sprintf("id %s ?", p), id).Delete(models.TblFile{}).Error
}

// 根据CreatedAt值批量删除查询
func (store *Store) DelTblFileByCreatedAt(p predicate.Predicate, createdAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ?", p), createdAt).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt值批量删除查询
func (store *Store) DelTblFileByUpdatedAt(p predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ?", p), updatedAt).Delete(models.TblFile{}).Error
}

// 根据DeletedAt值批量删除查询
func (store *Store) DelTblFileByDeletedAt(p predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ?", p), deletedAt).Delete(models.TblFile{}).Error
}

// 根据FileName值批量删除查询
func (store *Store) DelTblFileByFileName(p predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Delete(models.TblFile{}).Error
}

// 根据FileSha1值批量删除查询
func (store *Store) DelTblFileByFileSha1(p predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Delete(models.TblFile{}).Error
}

// 根据FileSize值批量删除查询
func (store *Store) DelTblFileByFileSize(p predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Delete(models.TblFile{}).Error
}

// 根据FileAddr值批量删除查询
func (store *Store) DelTblFileByFileAddr(p predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Delete(models.TblFile{}).Error
}

// 根据Statu值批量删除查询
func (store *Store) DelTblFileByStatu(p predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("status %s ?", p), statu).Delete(models.TblFile{}).Error
}

// 根据Ext1值批量删除查询
func (store *Store) DelTblFileByExt1(p predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("ext1 %s ?", p), ext1).Delete(models.TblFile{}).Error
}

// 根据Ext2值批量删除查询
func (store *Store) DelTblFileByExt2(p predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("ext2 %s ?", p), ext2).Delete(models.TblFile{}).Error
}

// 根据CreatedAt UpdatedAt值全部查询
func (store *Store) GetTblFileByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Find(&model).Error
}

// 根据CreatedAt DeletedAt值全部查询
func (store *Store) GetTblFileByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Find(&model).Error
}

// 根据CreatedAt FileName值全部查询
func (store *Store) GetTblFileByCreatedAtAndFileName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileName).Find(&model).Error
}

// 根据CreatedAt FileSha1值全部查询
func (store *Store) GetTblFileByCreatedAtAndFileSha1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSha1).Find(&model).Error
}

// 根据CreatedAt FileSize值全部查询
func (store *Store) GetTblFileByCreatedAtAndFileSize(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSize).Find(&model).Error
}

// 根据CreatedAt FileAddr值全部查询
func (store *Store) GetTblFileByCreatedAtAndFileAddr(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileAddr).Find(&model).Error
}

// 根据CreatedAt Statu值全部查询
func (store *Store) GetTblFileByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Find(&model).Error
}

// 根据CreatedAt Ext1值全部查询
func (store *Store) GetTblFileByCreatedAtAndExt1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, ext1).Find(&model).Error
}

// 根据CreatedAt Ext2值全部查询
func (store *Store) GetTblFileByCreatedAtAndExt2(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, ext2).Find(&model).Error
}

// 根据UpdatedAt DeletedAt值全部查询
func (store *Store) GetTblFileByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Find(&model).Error
}

// 根据UpdatedAt FileName值全部查询
func (store *Store) GetTblFileByUpdatedAtAndFileName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileName).Find(&model).Error
}

// 根据UpdatedAt FileSha1值全部查询
func (store *Store) GetTblFileByUpdatedAtAndFileSha1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSha1).Find(&model).Error
}

// 根据UpdatedAt FileSize值全部查询
func (store *Store) GetTblFileByUpdatedAtAndFileSize(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSize).Find(&model).Error
}

// 根据UpdatedAt FileAddr值全部查询
func (store *Store) GetTblFileByUpdatedAtAndFileAddr(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileAddr).Find(&model).Error
}

// 根据UpdatedAt Statu值全部查询
func (store *Store) GetTblFileByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Find(&model).Error
}

// 根据UpdatedAt Ext1值全部查询
func (store *Store) GetTblFileByUpdatedAtAndExt1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, ext1).Find(&model).Error
}

// 根据UpdatedAt Ext2值全部查询
func (store *Store) GetTblFileByUpdatedAtAndExt2(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, ext2).Find(&model).Error
}

// 根据DeletedAt FileName值全部查询
func (store *Store) GetTblFileByDeletedAtAndFileName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileName).Find(&model).Error
}

// 根据DeletedAt FileSha1值全部查询
func (store *Store) GetTblFileByDeletedAtAndFileSha1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSha1).Find(&model).Error
}

// 根据DeletedAt FileSize值全部查询
func (store *Store) GetTblFileByDeletedAtAndFileSize(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSize).Find(&model).Error
}

// 根据DeletedAt FileAddr值全部查询
func (store *Store) GetTblFileByDeletedAtAndFileAddr(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileAddr).Find(&model).Error
}

// 根据DeletedAt Statu值全部查询
func (store *Store) GetTblFileByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Find(&model).Error
}

// 根据DeletedAt Ext1值全部查询
func (store *Store) GetTblFileByDeletedAtAndExt1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, ext1).Find(&model).Error
}

// 根据DeletedAt Ext2值全部查询
func (store *Store) GetTblFileByDeletedAtAndExt2(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, ext2).Find(&model).Error
}

// 根据FileName FileSha1值全部查询
func (store *Store) GetTblFileByFileNameAndFileSha1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSha1).Find(&model).Error
}

// 根据FileName FileSize值全部查询
func (store *Store) GetTblFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Find(&model).Error
}

// 根据FileName FileAddr值全部查询
func (store *Store) GetTblFileByFileNameAndFileAddr(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileAddr).Find(&model).Error
}

// 根据FileName Statu值全部查询
func (store *Store) GetTblFileByFileNameAndStatu(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, statu).Find(&model).Error
}

// 根据FileName Ext1值全部查询
func (store *Store) GetTblFileByFileNameAndExt1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, ext1).Find(&model).Error
}

// 根据FileName Ext2值全部查询
func (store *Store) GetTblFileByFileNameAndExt2(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, ext2).Find(&model).Error
}

// 根据FileSha1 FileSize值全部查询
func (store *Store) GetTblFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Find(&model).Error
}

// 根据FileSha1 FileAddr值全部查询
func (store *Store) GetTblFileByFileSha1AndFileAddr(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileAddr).Find(&model).Error
}

// 根据FileSha1 Statu值全部查询
func (store *Store) GetTblFileByFileSha1AndStatu(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, statu).Find(&model).Error
}

// 根据FileSha1 Ext1值全部查询
func (store *Store) GetTblFileByFileSha1AndExt1(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, ext1).Find(&model).Error
}

// 根据FileSha1 Ext2值全部查询
func (store *Store) GetTblFileByFileSha1AndExt2(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, ext2).Find(&model).Error
}

// 根据FileSize FileAddr值全部查询
func (store *Store) GetTblFileByFileSizeAndFileAddr(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, fileAddr).Find(&model).Error
}

// 根据FileSize Statu值全部查询
func (store *Store) GetTblFileByFileSizeAndStatu(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Find(&model).Error
}

// 根据FileSize Ext1值全部查询
func (store *Store) GetTblFileByFileSizeAndExt1(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext1).Find(&model).Error
}

// 根据FileSize Ext2值全部查询
func (store *Store) GetTblFileByFileSizeAndExt2(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext2).Find(&model).Error
}

// 根据FileAddr Statu值全部查询
func (store *Store) GetTblFileByFileAddrAndStatu(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, statu).Find(&model).Error
}

// 根据FileAddr Ext1值全部查询
func (store *Store) GetTblFileByFileAddrAndExt1(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, ext1).Find(&model).Error
}

// 根据FileAddr Ext2值全部查询
func (store *Store) GetTblFileByFileAddrAndExt2(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, ext2).Find(&model).Error
}

// 根据Statu Ext1值全部查询
func (store *Store) GetTblFileByStatuAndExt1(p predicate.Predicate, statu int, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("status %s ? %s  %s ?", p, c, p2), statu, ext1).Find(&model).Error
}

// 根据Statu Ext2值全部查询
func (store *Store) GetTblFileByStatuAndExt2(p predicate.Predicate, statu int, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("status %s ? %s  %s ?", p, c, p2), statu, ext2).Find(&model).Error
}

// 根据Ext1 Ext2值全部查询
func (store *Store) GetTblFileByExt1AndExt2(p predicate.Predicate, ext1 int, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("ext1 %s ? %s  %s ?", p, c, p2), ext1, ext2).Find(&model).Error
}

// 根据CreatedAt UpdatedAt值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndUpdatedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updatedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, updatedAt).Delete(models.TblFile{}).Error
}

// 根据CreatedAt DeletedAt值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndDeletedAt(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, deletedAt).Delete(models.TblFile{}).Error
}

// 根据CreatedAt FileName值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndFileName(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileName).Delete(models.TblFile{}).Error
}

// 根据CreatedAt FileSha1值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndFileSha1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSha1).Delete(models.TblFile{}).Error
}

// 根据CreatedAt FileSize值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndFileSize(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileSize).Delete(models.TblFile{}).Error
}

// 根据CreatedAt FileAddr值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndFileAddr(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, fileAddr).Delete(models.TblFile{}).Error
}

// 根据CreatedAt Statu值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndStatu(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, statu).Delete(models.TblFile{}).Error
}

// 根据CreatedAt Ext1值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndExt1(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, ext1).Delete(models.TblFile{}).Error
}

// 根据CreatedAt Ext2值批量删除查询
func (store *Store) DelTblFileByCreatedAtAndExt2(p predicate.Predicate, createdAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("created_at %s ? %s  %s ?", p, c, p2), createdAt, ext2).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt DeletedAt值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndDeletedAt(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, deletedAt time.Time) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, deletedAt).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt FileName值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndFileName(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileName).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt FileSha1值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndFileSha1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSha1).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt FileSize值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndFileSize(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileSize).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt FileAddr值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndFileAddr(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, fileAddr).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt Statu值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndStatu(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, statu).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt Ext1值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndExt1(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, ext1).Delete(models.TblFile{}).Error
}

// 根据UpdatedAt Ext2值批量删除查询
func (store *Store) DelTblFileByUpdatedAtAndExt2(p predicate.Predicate, updatedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("updated_at %s ? %s  %s ?", p, c, p2), updatedAt, ext2).Delete(models.TblFile{}).Error
}

// 根据DeletedAt FileName值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndFileName(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileName).Delete(models.TblFile{}).Error
}

// 根据DeletedAt FileSha1值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndFileSha1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSha1).Delete(models.TblFile{}).Error
}

// 根据DeletedAt FileSize值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndFileSize(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileSize).Delete(models.TblFile{}).Error
}

// 根据DeletedAt FileAddr值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndFileAddr(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, fileAddr).Delete(models.TblFile{}).Error
}

// 根据DeletedAt Statu值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndStatu(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, statu).Delete(models.TblFile{}).Error
}

// 根据DeletedAt Ext1值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndExt1(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, ext1).Delete(models.TblFile{}).Error
}

// 根据DeletedAt Ext2值批量删除查询
func (store *Store) DelTblFileByDeletedAtAndExt2(p predicate.Predicate, deletedAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("deleted_at %s ? %s  %s ?", p, c, p2), deletedAt, ext2).Delete(models.TblFile{}).Error
}

// 根据FileName FileSha1值批量删除查询
func (store *Store) DelTblFileByFileNameAndFileSha1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSha1).Delete(models.TblFile{}).Error
}

// 根据FileName FileSize值批量删除查询
func (store *Store) DelTblFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Delete(models.TblFile{}).Error
}

// 根据FileName FileAddr值批量删除查询
func (store *Store) DelTblFileByFileNameAndFileAddr(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileName Statu值批量删除查询
func (store *Store) DelTblFileByFileNameAndStatu(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, statu).Delete(models.TblFile{}).Error
}

// 根据FileName Ext1值批量删除查询
func (store *Store) DelTblFileByFileNameAndExt1(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, ext1).Delete(models.TblFile{}).Error
}

// 根据FileName Ext2值批量删除查询
func (store *Store) DelTblFileByFileNameAndExt2(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, ext2).Delete(models.TblFile{}).Error
}

// 根据FileSha1 FileSize值批量删除查询
func (store *Store) DelTblFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Delete(models.TblFile{}).Error
}

// 根据FileSha1 FileAddr值批量删除查询
func (store *Store) DelTblFileByFileSha1AndFileAddr(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileSha1 Statu值批量删除查询
func (store *Store) DelTblFileByFileSha1AndStatu(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, statu).Delete(models.TblFile{}).Error
}

// 根据FileSha1 Ext1值批量删除查询
func (store *Store) DelTblFileByFileSha1AndExt1(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, ext1).Delete(models.TblFile{}).Error
}

// 根据FileSha1 Ext2值批量删除查询
func (store *Store) DelTblFileByFileSha1AndExt2(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, ext2).Delete(models.TblFile{}).Error
}

// 根据FileSize FileAddr值批量删除查询
func (store *Store) DelTblFileByFileSizeAndFileAddr(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileSize Statu值批量删除查询
func (store *Store) DelTblFileByFileSizeAndStatu(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Delete(models.TblFile{}).Error
}

// 根据FileSize Ext1值批量删除查询
func (store *Store) DelTblFileByFileSizeAndExt1(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext1).Delete(models.TblFile{}).Error
}

// 根据FileSize Ext2值批量删除查询
func (store *Store) DelTblFileByFileSizeAndExt2(p predicate.Predicate, fileSize string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext2).Delete(models.TblFile{}).Error
}

// 根据FileAddr Statu值批量删除查询
func (store *Store) DelTblFileByFileAddrAndStatu(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, statu).Delete(models.TblFile{}).Error
}

// 根据FileAddr Ext1值批量删除查询
func (store *Store) DelTblFileByFileAddrAndExt1(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, ext1).Delete(models.TblFile{}).Error
}

// 根据FileAddr Ext2值批量删除查询
func (store *Store) DelTblFileByFileAddrAndExt2(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, ext2).Delete(models.TblFile{}).Error
}

// 根据Statu Ext1值批量删除查询
func (store *Store) DelTblFileByStatuAndExt1(p predicate.Predicate, statu int, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("status %s ? %s  %s ?", p, c, p2), statu, ext1).Delete(models.TblFile{}).Error
}

// 根据Statu Ext2值批量删除查询
func (store *Store) DelTblFileByStatuAndExt2(p predicate.Predicate, statu int, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("status %s ? %s  %s ?", p, c, p2), statu, ext2).Delete(models.TblFile{}).Error
}

// 根据Ext1 Ext2值批量删除查询
func (store *Store) DelTblFileByExt1AndExt2(p predicate.Predicate, ext1 int, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("ext1 %s ? %s  %s ?", p, c, p2), ext1, ext2).Delete(models.TblFile{}).Error
}
