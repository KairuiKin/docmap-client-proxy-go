package event

import (
	"crypto/sha1"
	__ "docmap-client-proxy-go/internal/proto"
	"docmap-client-proxy-go/internal/proto/docmap/exchange/model"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileTimes 结构体用于存储文件的创建时间和修改时间
type FileTimes struct {
	CreateTime uint64
	ModifyTime uint64
}

// 获取文件扩展名
func GetFileExt(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

// 根据文件扩展名获取分类类型（示例实现，需根据实际需求调整）
func GetClassificationTypeFromExtName(ext string) string {
	switch ext {
	case ".txt", ".md":
		return "Text"
	case ".jpg", ".png", ".gif":
		return "Image"
	case ".pdf":
		return "Document"
	default:
		return "OTHER"
	}
}

// 计算文件的 SHA1 哈希值
func CalculateSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 获取驱动器类型
func GetDriveTypeFunc(path string) string {
	return GetDriveType(path)
}

// 获取安全级别（示例设置为 "common2"）
func GetSecurityLevel() string {
	return "common2"
}

// 获取驱动器类型（平台特定实现）
func GetDriveType(path string) string {
	// 您需要根据之前的说明为不同平台提供实现
	// 这里仅作为示例，返回 "LocalDisk"
	return "Harddisk"
}

func (p *EventPlugin) eventCreate(event *model.EventCreate) (*__.Create, error) {
	path := event.GetDestinationFile()
	var createFile *__.Create
	var err error
	createFile = &__.Create{}
	createFile.DestinationFile, err = GetFileInfo(path, "")
	return createFile, err
}
