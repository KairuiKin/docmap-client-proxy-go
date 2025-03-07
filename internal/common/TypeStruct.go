package common

import "errors"

type InterfaceSend interface {
	Send(msg *Message) *Message
}

var interfaceRegistry = make(map[string]interface{})

// RegisterInterface 注册接口实现
func RegisterInterface(name string, instance interface{}) {
	interfaceRegistry[name] = instance
}

// GetInterface 获取接口实现
func GetInterface(name string) (interface{}, error) {
	instance, exists := interfaceRegistry[name]
	if !exists {
		return nil, errors.New("interface not found: " + name)
	}
	return instance, nil
}

type Message struct {
	SourceID    string        // 消息来源插件 ID
	TargetID    string        // 消息目标插件 ID，空字符串表示广播
	Type        MessageType   // 消息类型：消息还是响应
	Destination MessageType   // 消息类型：内部消息定义
	Payload     interface{}   // 消息内容
	ResponseCh  chan *Message // 用于同步消息的响应通道
}
type MessageType int

const (
	MessageInfo   MessageType = iota // 同步消息
	MessageReport                    //响应
)

type ExchangeMode int

const (
	MODE_PIPE ExchangeMode = iota
	MODE_SOCKET
)

const (
	// 未定义属性
	UndefinedAttributeValue = "policy_attr_value_undefined"

	// 应用信息
	AppInfoImagePath = "image_path"

	// 事件信息
	EventInfoEventName = "event_name"
	EventInfoEventType = "event_type"

	// 资源信息 - 文件路径
	ResourceInfoFileSrc        = "file.source.path"
	ResourceInfoFileSrcSupport = "file.source.support"
	ResourceInfoFileDst        = "file.dst.path"
	ResourceInfoFileDstSupport = "file.dst.support"
	ResourceInfoFileDstAttrs   = "file.dst.attrs"

	// 资源信息 - 注册表
	ResourceInfoRegistryKeyPath   = "reg.key.path"
	ResourceInfoRegistryValueName = "reg.value.name"

	// 资源信息 - 插入的文件和标签
	ResourceInfoFileInserted            = "file.inserted.path"
	ResourceInfoTagInsertedSupport      = "tag.inserted.support"
	ResourceInfoTagInsertedFileID       = "tag.inserted.fileid"
	ResourceInfoTagInsertedFamilyID     = "tag.inserted.familyid"
	ResourceInfoTagInsertedEncLevel     = "tag.inserted.enclevel"
	ResourceInfoTagInsertedCategory     = "tag.inserted.category"
	ResourceInfoTagInsertedAllowOutdoor = "tag.inserted.allowoutdoor"

	// 资源信息 - 地址
	ResourceInfoAddrURL = "file.addr.url"
	ResourceInfoAddrIP  = "file.addr.ip"

	// 资源信息 - 源标签
	ResourceInfoTagSrcSupport           = "tag.src.support"
	ResourceInfoTagSrcFileID            = "tag.src.fileid"
	ResourceInfoTagSrcFamilyID          = "tag.src.familyid"
	ResourceInfoTagSrcEncLevel          = "tag.src.enclevel"
	ResourceInfoTagSrcCategory          = "tag.src.category"
	ResourceInfoTagSrcAllowOutdoor      = "tag.src.allowoutdoor"
	ResourceInfoTagSrcIPInnerNetEncrypt = "tag.src.ip.innernetencrypt"
	ResourceInfoTagSrcIPOuterNetEncrypt = "tag.src.ip.outernetencrypt"
	ResourceInfoTagSrcIPEnabled         = "tag.src.ip.enabled"

	// 资源信息 - 目标标签
	ResourceInfoTagDstSupport           = "tag.dst.support"
	ResourceInfoTagDstFileID            = "tag.dst.fileid"
	ResourceInfoTagDstFamilyID          = "tag.dst.familyid"
	ResourceInfoTagDstEncLevel          = "tag.dst.enclevel"
	ResourceInfoTagDstCategory          = "tag.dst.category"
	ResourceInfoTagDstAllowOutdoor      = "tag.dst.allowoutdoor"
	ResourceInfoTagDstIPInnerNetEncrypt = "tag.dst.ip.innernetencrypt"
	ResourceInfoTagDstIPOuterNetEncrypt = "tag.dst.ip.outernetencrypt"
	ResourceInfoTagDstIPEnabled         = "tag.dst.ip.enabled"

	// 资源信息 - 文档窗口
	ResourceInfoDocWindow = "doc.window"

	// 事件名称
	EventInfoEventNameNewFile             = "create"
	EventInfoEventNameCopy                = "copy"
	EventInfoEventNameSaveAs              = "saveas"
	EventInfoEventNameZip                 = "zip"
	EventInfoEventNameUnzip               = "unzip"
	EventInfoEventNameUpload              = "upload"
	EventInfoEventNameDownload            = "download"
	EventInfoEventNameEdit                = "edit"
	EventInfoEventNameMove                = "move"
	EventInfoEventNameRename              = "rename"
	EventInfoEventNameDelete              = "delete"
	EventInfoEventNameRecycle             = "recycle"
	EventInfoEventNameRestore             = "restore"
	EventInfoEventNameChangeAttr          = "change_attributes"
	EventInfoEventNameRun                 = "run"
	EventInfoEventNameKill                = "kill"
	EventInfoEventNameBurnCD              = "burncd"
	EventInfoEventNameFTPSend             = "ftpsend"
	EventInfoEventNameBeginEdit           = "beginedit"
	EventInfoEventNameEndEdit             = "endedit"
	EventInfoEventNameBeginRead           = "beginread"
	EventInfoEventNameEndRead             = "endread"
	EventInfoEventNameWriteTag            = "writetag"
	EventInfoEventNameCopyContentPrivate  = "_copycontent"
	EventInfoEventNamePasteContentPrivate = "_pastecontent"
	EventInfoEventNamePasteContent        = "pastecontent"
	EventInfoEventNameEncrypt             = "encrypt"
	EventInfoEventNameDecrypt             = "decrypt"
	EventInfoEventNameInsertFile          = "insertfile"
	EventInfoEventNameRenameFolder        = "renamefolder"
	EventInfoEventNameChangeRegistryValue = "changeregval"
	EventInfoEventNamePermissionEvent     = "permissionevent"
	EventInfoEventNamePrint               = "print"
	EventInfoEventNameOutdoor             = "outdoor"
	EventInfoEventNameMoveFolder          = "movefolder"
	EventInfoEventNameDeleteFolder        = "deletefolder"
	EventInfoEventNameRecycleFolder       = "recyclefolder"
	EventInfoEventNameRestoreFolder       = "restorefolder"

	// 自定义操作名称
	CustomActionNameIdentityID           = "identityID"
	CustomActionNameInheritSecurityLevel = "inheritSecurityLevel"
	CustomActionNameTag                  = "tag"
	CustomActionNameLog                  = "log"

	CustomActionNameFileIDNew    = "newFileID"
	CustomActionNameRedirectFile = "redirectFile"
)
