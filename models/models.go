package models

type TblFile struct {
	FileName string
	FileSha1 string
	FileSize int64
	FileAddr string
}

func OnFileUploadFinished(fileHash, fileName string, fileSize int64, fileAddr string) bool {
	//插入数据
	return true
}
