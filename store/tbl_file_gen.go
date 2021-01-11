package store

////////////////////////////////////////////////////////////////////
//   Warning: this file is auto, please not edit.  自动生成请勿编辑
//   Author : ftfeng
//   主要包含：1、创建/删除/更新实例  2、更新实例 3、根据列获取所有 3、根据列分页获取
//   Create time : 2021-01-11 23:22:47
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

// 根据FileSha1值查询所有
func (store *Store) GetTblFileByFileSha1(p predicate.Predicate, fileSha1 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Find(&model).Error
}

// 根据FileName值查询所有
func (store *Store) GetTblFileByFileName(p predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Find(&model).Error
}

// 根据FileSize值查询所有
func (store *Store) GetTblFileByFileSize(p predicate.Predicate, fileSize int64) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Find(&model).Error
}

// 根据FileAddr值查询所有
func (store *Store) GetTblFileByFileAddr(p predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Find(&model).Error
}

// 根据CreateAt值查询所有
func (store *Store) GetTblFileByCreateAt(p predicate.Predicate, createAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ?", p), createAt).Find(&model).Error
}

// 根据UpdateAt值查询所有
func (store *Store) GetTblFileByUpdateAt(p predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("update_at %s ?", p), updateAt).Find(&model).Error
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

// 根据FileSha1值分页查询
func (store *Store) GetTblFileByFileSha1ByPage(p predicate.Predicate, fileSha1 string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileName值分页查询
func (store *Store) GetTblFileByFileNameByPage(p predicate.Predicate, fileName string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileSize值分页查询
func (store *Store) GetTblFileByFileSizeByPage(p predicate.Predicate, fileSize int64, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据FileAddr值分页查询
func (store *Store) GetTblFileByFileAddrByPage(p predicate.Predicate, fileAddr string, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据CreateAt值分页查询
func (store *Store) GetTblFileByCreateAtByPage(p predicate.Predicate, createAt time.Time, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ?", p), createAt).Limit(limit).Offset(offset).Find(&model).Error
}

// 根据UpdateAt值分页查询
func (store *Store) GetTblFileByUpdateAtByPage(p predicate.Predicate, updateAt time.Time, limit interface{}, offset interface{}) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("update_at %s ?", p), updateAt).Limit(limit).Offset(offset).Find(&model).Error
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

// 根据FileSha1值批量删除查询
func (store *Store) DelTblFileByFileSha1(p predicate.Predicate, fileSha1 string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ?", p), fileSha1).Delete(models.TblFile{}).Error
}

// 根据FileName值批量删除查询
func (store *Store) DelTblFileByFileName(p predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ?", p), fileName).Delete(models.TblFile{}).Error
}

// 根据FileSize值批量删除查询
func (store *Store) DelTblFileByFileSize(p predicate.Predicate, fileSize int64) error {
	return store.db.Where(fmt.Sprintf("file_size %s ?", p), fileSize).Delete(models.TblFile{}).Error
}

// 根据FileAddr值批量删除查询
func (store *Store) DelTblFileByFileAddr(p predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ?", p), fileAddr).Delete(models.TblFile{}).Error
}

// 根据CreateAt值批量删除查询
func (store *Store) DelTblFileByCreateAt(p predicate.Predicate, createAt time.Time) error {
	return store.db.Where(fmt.Sprintf("create_at %s ?", p), createAt).Delete(models.TblFile{}).Error
}

// 根据UpdateAt值批量删除查询
func (store *Store) DelTblFileByUpdateAt(p predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("update_at %s ?", p), updateAt).Delete(models.TblFile{}).Error
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

// 根据FileSha1 FileName值全部查询
func (store *Store) GetTblFileByFileSha1AndFileName(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileName string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileName).Find(&model).Error
}

// 根据FileSha1 FileSize值全部查询
func (store *Store) GetTblFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize int64) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Find(&model).Error
}

// 根据FileSha1 FileAddr值全部查询
func (store *Store) GetTblFileByFileSha1AndFileAddr(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileAddr).Find(&model).Error
}

// 根据FileSha1 CreateAt值全部查询
func (store *Store) GetTblFileByFileSha1AndCreateAt(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, createAt).Find(&model).Error
}

// 根据FileSha1 UpdateAt值全部查询
func (store *Store) GetTblFileByFileSha1AndUpdateAt(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, updateAt).Find(&model).Error
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

// 根据FileName FileSize值全部查询
func (store *Store) GetTblFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize int64) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Find(&model).Error
}

// 根据FileName FileAddr值全部查询
func (store *Store) GetTblFileByFileNameAndFileAddr(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileAddr).Find(&model).Error
}

// 根据FileName CreateAt值全部查询
func (store *Store) GetTblFileByFileNameAndCreateAt(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, createAt).Find(&model).Error
}

// 根据FileName UpdateAt值全部查询
func (store *Store) GetTblFileByFileNameAndUpdateAt(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, updateAt).Find(&model).Error
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

// 根据FileSize FileAddr值全部查询
func (store *Store) GetTblFileByFileSizeAndFileAddr(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, fileAddr).Find(&model).Error
}

// 根据FileSize CreateAt值全部查询
func (store *Store) GetTblFileByFileSizeAndCreateAt(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, createAt).Find(&model).Error
}

// 根据FileSize UpdateAt值全部查询
func (store *Store) GetTblFileByFileSizeAndUpdateAt(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, updateAt).Find(&model).Error
}

// 根据FileSize Statu值全部查询
func (store *Store) GetTblFileByFileSizeAndStatu(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Find(&model).Error
}

// 根据FileSize Ext1值全部查询
func (store *Store) GetTblFileByFileSizeAndExt1(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext1).Find(&model).Error
}

// 根据FileSize Ext2值全部查询
func (store *Store) GetTblFileByFileSizeAndExt2(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext2).Find(&model).Error
}

// 根据FileAddr CreateAt值全部查询
func (store *Store) GetTblFileByFileAddrAndCreateAt(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, createAt).Find(&model).Error
}

// 根据FileAddr UpdateAt值全部查询
func (store *Store) GetTblFileByFileAddrAndUpdateAt(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, updateAt).Find(&model).Error
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

// 根据CreateAt UpdateAt值全部查询
func (store *Store) GetTblFileByCreateAtAndUpdateAt(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, updateAt).Find(&model).Error
}

// 根据CreateAt Statu值全部查询
func (store *Store) GetTblFileByCreateAtAndStatu(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, statu).Find(&model).Error
}

// 根据CreateAt Ext1值全部查询
func (store *Store) GetTblFileByCreateAtAndExt1(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, ext1).Find(&model).Error
}

// 根据CreateAt Ext2值全部查询
func (store *Store) GetTblFileByCreateAtAndExt2(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, ext2).Find(&model).Error
}

// 根据UpdateAt Statu值全部查询
func (store *Store) GetTblFileByUpdateAtAndStatu(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, statu).Find(&model).Error
}

// 根据UpdateAt Ext1值全部查询
func (store *Store) GetTblFileByUpdateAtAndExt1(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, ext1).Find(&model).Error
}

// 根据UpdateAt Ext2值全部查询
func (store *Store) GetTblFileByUpdateAtAndExt2(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) ([]*models.TblFile, error) {
	var model []*models.TblFile
	return model, store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, ext2).Find(&model).Error
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

// 根据FileSha1 FileName值批量删除查询
func (store *Store) DelTblFileByFileSha1AndFileName(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileName string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileName).Delete(models.TblFile{}).Error
}

// 根据FileSha1 FileSize值批量删除查询
func (store *Store) DelTblFileByFileSha1AndFileSize(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileSize int64) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileSize).Delete(models.TblFile{}).Error
}

// 根据FileSha1 FileAddr值批量删除查询
func (store *Store) DelTblFileByFileSha1AndFileAddr(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileSha1 CreateAt值批量删除查询
func (store *Store) DelTblFileByFileSha1AndCreateAt(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, createAt).Delete(models.TblFile{}).Error
}

// 根据FileSha1 UpdateAt值批量删除查询
func (store *Store) DelTblFileByFileSha1AndUpdateAt(p predicate.Predicate, fileSha1 string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_sha1 %s ? %s  %s ?", p, c, p2), fileSha1, updateAt).Delete(models.TblFile{}).Error
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

// 根据FileName FileSize值批量删除查询
func (store *Store) DelTblFileByFileNameAndFileSize(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileSize int64) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileSize).Delete(models.TblFile{}).Error
}

// 根据FileName FileAddr值批量删除查询
func (store *Store) DelTblFileByFileNameAndFileAddr(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileName CreateAt值批量删除查询
func (store *Store) DelTblFileByFileNameAndCreateAt(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, createAt).Delete(models.TblFile{}).Error
}

// 根据FileName UpdateAt值批量删除查询
func (store *Store) DelTblFileByFileNameAndUpdateAt(p predicate.Predicate, fileName string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_name %s ? %s  %s ?", p, c, p2), fileName, updateAt).Delete(models.TblFile{}).Error
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

// 根据FileSize FileAddr值批量删除查询
func (store *Store) DelTblFileByFileSizeAndFileAddr(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, fileAddr string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, fileAddr).Delete(models.TblFile{}).Error
}

// 根据FileSize CreateAt值批量删除查询
func (store *Store) DelTblFileByFileSizeAndCreateAt(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, createAt).Delete(models.TblFile{}).Error
}

// 根据FileSize UpdateAt值批量删除查询
func (store *Store) DelTblFileByFileSizeAndUpdateAt(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, updateAt).Delete(models.TblFile{}).Error
}

// 根据FileSize Statu值批量删除查询
func (store *Store) DelTblFileByFileSizeAndStatu(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, statu).Delete(models.TblFile{}).Error
}

// 根据FileSize Ext1值批量删除查询
func (store *Store) DelTblFileByFileSizeAndExt1(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext1).Delete(models.TblFile{}).Error
}

// 根据FileSize Ext2值批量删除查询
func (store *Store) DelTblFileByFileSizeAndExt2(p predicate.Predicate, fileSize int64, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("file_size %s ? %s  %s ?", p, c, p2), fileSize, ext2).Delete(models.TblFile{}).Error
}

// 根据FileAddr CreateAt值批量删除查询
func (store *Store) DelTblFileByFileAddrAndCreateAt(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, createAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, createAt).Delete(models.TblFile{}).Error
}

// 根据FileAddr UpdateAt值批量删除查询
func (store *Store) DelTblFileByFileAddrAndUpdateAt(p predicate.Predicate, fileAddr string, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("file_addr %s ? %s  %s ?", p, c, p2), fileAddr, updateAt).Delete(models.TblFile{}).Error
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

// 根据CreateAt UpdateAt值批量删除查询
func (store *Store) DelTblFileByCreateAtAndUpdateAt(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, updateAt time.Time) error {
	return store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, updateAt).Delete(models.TblFile{}).Error
}

// 根据CreateAt Statu值批量删除查询
func (store *Store) DelTblFileByCreateAtAndStatu(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, statu).Delete(models.TblFile{}).Error
}

// 根据CreateAt Ext1值批量删除查询
func (store *Store) DelTblFileByCreateAtAndExt1(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, ext1).Delete(models.TblFile{}).Error
}

// 根据CreateAt Ext2值批量删除查询
func (store *Store) DelTblFileByCreateAtAndExt2(p predicate.Predicate, createAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("create_at %s ? %s  %s ?", p, c, p2), createAt, ext2).Delete(models.TblFile{}).Error
}

// 根据UpdateAt Statu值批量删除查询
func (store *Store) DelTblFileByUpdateAtAndStatu(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, statu int) error {
	return store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, statu).Delete(models.TblFile{}).Error
}

// 根据UpdateAt Ext1值批量删除查询
func (store *Store) DelTblFileByUpdateAtAndExt1(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext1 int) error {
	return store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, ext1).Delete(models.TblFile{}).Error
}

// 根据UpdateAt Ext2值批量删除查询
func (store *Store) DelTblFileByUpdateAtAndExt2(p predicate.Predicate, updateAt time.Time, c predicate.Conjunction, p2 predicate.Predicate, ext2 string) error {
	return store.db.Where(fmt.Sprintf("update_at %s ? %s  %s ?", p, c, p2), updateAt, ext2).Delete(models.TblFile{}).Error
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
