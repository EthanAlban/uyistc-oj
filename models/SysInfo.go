package models

// SysInfo 保存系统通知
type SysInfo struct {
	Iid     int    `orm:"column(iid);pk"`
	UId     *User  `orm:"column(uid);rel(fk)"`
	Content string `orm:"column(content)"`
	Read    int    `orm:"column(read)"`
}

// TableName 获取对应数据库表名.
func (i *SysInfo) TableName() string {
	return "sysinfo"
}

// TableEngine 获取数据使用的引擎.
func (i *SysInfo) TableEngine() string {
	return "INNODB"
}

func NewSysInfo() *SysInfo {
	return &SysInfo{}
}

// GetSysInfoList 获取分页的系统通知（未读消息）
func (i *SysInfo) GetSysInfoList(limit, offset int) *[]SysInfo {
	var sysinfoList []SysInfo
	O.QueryTable("sysinfo").Limit(limit, offset).Filter("read", 0).All(&sysinfoList)
	return &sysinfoList
}

// SetSysInfoReadedById 设置系统通知已读
func (i *SysInfo) SetSysInfoReadedById(Iid int) bool {
	sysinfo := SysInfo{Iid: Iid}
	if O.Read(&sysinfo, "Iid") == nil {
		sysinfo.Read = 1
		if _, err := O.Update(&sysinfo); err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
