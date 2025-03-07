package event

import (
	"docmap-client-proxy-go/internal/common"
	__ "docmap-client-proxy-go/internal/proto"
	"github.com/djherbis/times"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"path/filepath"
	"time"
)

func GetFileTimesUsingLibrary(path string) (FileTimes, error) {
	t, err := times.Stat(path)
	if err != nil {
		return FileTimes{}, err
	}

	var ft FileTimes
	if t.HasBirthTime() {
		ft.CreateTime = uint64(t.BirthTime().UnixNano() / int64(time.Millisecond))
	} else {
		// 如果没有创建时间，使用修改时间作为替代
		ft.CreateTime = uint64(t.ModTime().UnixNano() / int64(time.Millisecond))
	}
	ft.ModifyTime = uint64(t.ModTime().UnixNano() / int64(time.Millisecond))
	return ft, nil
}

func GetFileInfo(path, entityIdentity string) (*__.FileInfo, error) {

	// 获取文件名和扩展名
	filename := filepath.Base(path)
	ext := GetFileExt(path)

	// 获取分类类型
	classification := GetClassificationTypeFromExtName(ext)

	// 获取驱动器类型
	driveType := GetDriveTypeFunc(path)

	// 获取安全级别
	securityLevel := GetSecurityLevel()

	var size uint64
	var sha1Value string
	var fileTimes FileTimes
	fileInfo, err := os.Stat(path)
	if err == nil {
		// 获取创建时间和修改时间
		fileTimes, err = GetFileTimesUsingLibrary(path)
		if err != nil {
			common.Logger.Printf("获取文件时间失败: %v", err)
			fileTimes.CreateTime = 0
			fileTimes.ModifyTime = uint64(fileInfo.ModTime().UnixNano() / int64(time.Millisecond))
		}

		// 获取文件大小
		size = uint64(fileInfo.Size())

		// 计算 SHA1
		sha1Value, err = CalculateSHA1(path)

		if err != nil {
			common.Logger.Printf("计算文件 SHA1 失败: %v", err)
			sha1Value = ""
		}
	}
	// 构建 FileInfo protobuf 消息
	fi := &__.FileInfo{
		EntityIdentity:     wrapperspb.String(entityIdentity),
		Path:               wrapperspb.String(path),
		Filename:           wrapperspb.String(filename),
		DriveType:          wrapperspb.String(driveType),
		SecurityLevel:      wrapperspb.String(securityLevel),
		ClassificationType: []*wrapperspb.StringValue{wrapperspb.String(classification)},
		CreateTime:         wrapperspb.UInt64(fileTimes.CreateTime),
		ModifyTime:         wrapperspb.UInt64(fileTimes.ModifyTime),
		Sha1:               wrapperspb.String(sha1Value),
		Size:               wrapperspb.UInt64(size),
		Ext:                wrapperspb.String(ext),
		EncryptChannel:     wrapperspb.String(""), // 设置为空字符串
		Encrypted:          wrapperspb.Int32(0),   // 设置为 false (0)
	}

	return fi, nil
}
