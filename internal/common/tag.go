package common

import (
	"path/filepath"
	"strings"
)

const (
	MessageGetFileID    MessageType = iota // 获取文件ID
	MessageGetFileLevel                    // 获取文件安全级别
	MessageSetFileID                       // 设置文件ID
	MessageSetFileLevel                    // 设置文件安全级别
	MessageDeleteFile                      // 设置文件安全级别
)

type Secret struct {
	Level string `yaml:"level" json:"level"`
	Code  string `yaml:"code" json:"code"`
	Name  string `yaml:"name" json:"name"`
}

type DocumentClassification struct {
	FilePath  string `json:"file_path"`
	FileLevel Secret `json:"secret"`
}

type IdentityWriteTask struct {
	FilePath string
	Identity FileIdentity
}

// FileIdentity 封装了 FamilyID 和 IndividualID
type FileIdentity struct {
	Family   string
	Identity string
	FileHash string
}

// Message 消息结构体
type SetIdentityRequest struct {
	FilePath       string `json:"filepath"`
	ParentFilePath string `json:"parent_filepath"`
	IsMove         int    `json:"is_move"`
}

var AllowedExtensions = []string{
	// 文档格式
	"wps", "et", "txt", "dps", "pdf",
	// 演示和表格格式
	"ppt", "pptx", "doc", "docx", "xls", "xlsx",
	// 版式文件格式
	"ofd", "sep", "gd",
	// 压缩文件格式
	"rar", "zip", "7z", "tar", "gz",
	// 图像文件格式
	"jpg", "jpeg", "png", "svg", "gif",
	// 音频/视频文件格式
	"mp3", "mp4", "wav", "wma", "avi",
}

// IsExtensionAllowed 判断后缀名是否在白名单内
func IsExtensionAllowed(filePath string) bool {
	// 获取带点的后缀名，比如 ".txt"
	ext := filepath.Ext(filePath)
	// 去掉点变成 "txt"
	ext = strings.TrimPrefix(ext, ".")

	// 遍历 allowedExtensions，看是否包含该后缀
	for _, allowed := range AllowedExtensions {
		if strings.EqualFold(ext, allowed) {
			// strings.EqualFold 不区分大小写比较
			// 如果希望区分大小写可改用 ext == allowed
			return true
		}
	}
	return false
}
