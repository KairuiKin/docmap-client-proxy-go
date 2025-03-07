package common

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func DMRandomUuid() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// 设置 UUID 版本为 4 （随机生成）
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // 0100 xxxx - Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // 10xx xxxx - Variant 10

	uuidStr := fmt.Sprintf("%08x%04x%04x%04x%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])

	// 移除所有的 "-" 符号
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")

	return uuidStr, nil
}

// GetClassificationType 根据应用程序文件名返回分类类型
func GetClassificationType(appFileName string) string {
	// 定义各分类的应用程序列表
	officeApps := []string{
		"winword.exe", "excel.exe", "powerpnt.exe", "wps.exe", "wpspdf.exe",
		"et.exe", "wpp.exe", "foxitpdfreader.exe", "acrobat.exe", "foxitphantom.exe",
		"foxitphantompdf.exe",
	}
	imageEditorApps := []string{
		"photoshop.exe",
	}
	browserApps := []string{
		"chrome.exe", "iexplore.exe", "firefox.exe", "qqbrowser.exe",
		"360se.exe", "msedge.exe", "wxworkweb.exe",
	}
	imApps := []string{
		"wechat.exe", "dingtalk.exe", "qq.exe", "wxwork.exe",
		"wemeetapp.exe", "feishu.exe", "wwmapp.exe",
	}
	zipApps := []string{
		"winrar.exe", "winzip.exe", "winzip32.exe", "winzip64.exe", "7zg.exe",
		"7zfm.exe", "haozip.exe", "360zip.exe", "kuaizip.exe", "multimediaopt.exe",
	}
	mailApps := []string{
		"outlook.exe", "foxmail.exe", "wemail.exe", "wemailnode.exe",
	}

	// 定义分类名称和对应的应用程序列表
	categories := []struct {
		name string
		apps []string
	}{
		{"OFFICE", officeApps},
		{"IMAGE EDITOR", imageEditorApps},
		{"BROWSER", browserApps},
		{"IM", imApps},
		{"ZIP", zipApps},
		{"MAIL", mailApps},
	}

	// 将应用程序名称转换为小写以实现不区分大小写的比较
	lowerName := strings.ToLower(appFileName)

	// 遍历分类并查找匹配
	for _, category := range categories {
		for _, app := range category.apps {
			if lowerName == strings.ToLower(app) {
				return category.name
			}
		}
	}

	// 如果没有匹配到任何分类，返回 "OTHER"
	return "OTHER"
}
