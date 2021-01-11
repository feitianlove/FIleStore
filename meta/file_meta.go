package meta

import "github.com/feitianlove/FIleStore/models"

// 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var FileMetas map[string]FileMeta

func init() {
	FileMetas = make(map[string]FileMeta)
}

// 新增或者更新文件元信息
func UpdateFileMeta(fMeta FileMeta) {
	FileMetas[fMeta.FileSha1] = fMeta
}

// 获取文件元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return FileMetas[fileSha1]
}

//删除文件元信息

func RemoveFileMeta(fileSha1 string) {
	delete(FileMetas, fileSha1)
}

//UPFileMetaDB
func UPDateFileMetaDB(meta FileMeta) bool {
	return models.OnFileUploadFinished(meta.FileSha1, meta.FileName, meta.FileSize, meta.Location)
}
