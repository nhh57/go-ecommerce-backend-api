package model

const TableNameGoCrmUser = "go_crm_user"

// GoCrmUser Account

type GoCrmUser struct {
	UsrID            int32  `gorm:"column:usr_id;primaryKey;autoIncrement:true;comment:Account ID" json:"user_id"`
	UsrEmail         string `gorm:"column:usr_email;not null;comment:Email" json:"usr_email"`
	UsrPhone         string `gorm:"column:usr_phone;not null;comment:Phone" json:"usr_phone"`
	UsrUsername      string `gorm:"column:usr_username;not null;comment:Username" json:"usr_username"`
	UsrPassword      string `gorm:"column:usr_password;not null;comment:Password" json:"usr_password"`
	UsrCreatedAt     int32  `gorm:"column:usr_created_at;not null;comment:Creation Time" json:"usr_created_at"`
	UsrUpdatedAt     int32  `gorm:"column:usr_updated_at;not null;comment:Update Time" json:"usr_updated_at"`
	UsrCreateIPAt    string `gorm:"column:usr_create_ip_at;not null;comment:Creation IP" json:"usr_create_ip_at"`
	UsrLastLoginAt   int32  `gorm:"column:usr_last_login_at;not null;comment:Last Login Time" json:"usr_last_login_at"`
	UsrLastLoginIPAt string `gorm:"column:usr_last_login_ip_at;not null;comment:Last Login IP" json:"usr_last_login_ip_at"`
	UsrLoginTimes    int32  `gorm:"column:usr_login_times;not null;comment:Login Times" json:"usr_login_times"`
	UsrStatus        bool   `gorm:"column:usr_status;not null;comment:Status 1:enable, 0:disable, -1:delete" json:"usr_status"`
}
