package common

const (
	MessageLogin                 MessageType = iota // 登入
	MessageAnonymousLogin                           // 匿名登录
	MessageLogout                                   // 登出
	MessageStatOuterAudit                           // 外发申请
	MessageOuterAudit                               // 外发审批
	MessageGetConfig                                // 获取配置
	MessageSendLog                                  // 发送日志
	MessageGetUpgradeInformation                    // 获取升级信息
	MessageDownloadUrlFile                          // 下载指定的文件
	MessageGetFileCategories                        // 获取文件类别范围
	MessageGetFileLevels                            // 获取文件密级范围
)

type SecurityLevelRequest struct {
	FilePath      string `json:"filePath"`
	SecurityLevel string `json:"securityLevel"`
}
