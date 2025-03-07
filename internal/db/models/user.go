package models

// 用户管理
type User struct {
	Id            int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;comment:账号ID"`
	Name          string `json:"name" gorm:"type:varchar(150);default '';comment:账号名称;"`
	Email         string `json:"email" gorm:"type:varchar(150);default:'';not null;comment:邮箱;"`
	Phone         string `json:"phone" gorm:"type:varchar(150);default:'';not null;comment:手机号;"`
	Password      string `json:"password" gorm:"type:varchar(150);not null;comment:密码"`
	Subject       string `json:"subject" gorm:"type:varchar(100);not null;comment:客户所属公司"`
	MachineList   string `json:"machine_list" gorm:"type:varchar(200);not null;comment:客户机器列表"`
	MachineNumber int32  `json:"machine_number" gorm:"type:int;comment:客户机器数量"`
	AppKey        string `json:"app_key" gorm:"type:varchar(150);default '';comment:应用密钥"`
	Salt          string `json:"salt" gorm:"type:varchar(150);not null;comment:密码盐"`
	RoleId        int8   `json:"role_id" gorm:"type:tinyint unsigned;default:2;not null; comment:角色ID"`
	Description   string `json:"description" gorm:"type:varchar(200);comment:备注"`
	Status        int8   `json:"status" gorm:"type:tinyint unsigned;default:0;not null;comment:账号状态"`
	LoginIp       string `json:"login_ip" gorm:"type:varchar(50);default '';comment:上次登录IP"`
	LoginTime     int64  `json:"login_time" gorm:"type:bigint unsigned;default:0;not null;comment:上次登录时间"`
	CreateTime    int64  `json:"create_time" gorm:"type:bigint unsigned;default:0;not null;comment:注册时间"`
	UpdateTime    int64  `json:"update_time" gorm:"type:bigint unsigned;default:0;not null;comment:会员信息上次更新时间"`
}

var UserField = []string{
	"name",
	"password",
	"subject",
	"machine_list",
	"machine_number",
	"salt",
	"description",
	"status",
	"update_time",
}

type UserInfo struct {
	Id            int64  `json:"customer_id"`
	Name          string `json:"customer_name"`
	Password      string `json:"customer_password"`
	Subject       string `json:"customer_subject"`
	MachineList   string `json:"machine_list"`
	MachineNumber int32  `json:"machine_number"`
	RoleId        int8   `json:"role_id"`
	Description   string `json:"customer_desc"`
	Status        int8   `json:"customer_status"`
	LoginIp       string `json:"login_ip"`
	LoginTime     int64  `json:"login_time"`
	CreateTime    int64  `json:"create_time"`
	UpdateTime    int64  `json:"update_time"`
}

type UserIdAndName struct {
	Id   int64  `json:"customer_id"`
	Name string `json:"customer_name"`
}

// 获取表名
func (User) TableName() string {
	return "ip_user"
}

// 根据ID查询用户
func (model User) FindById() (user User, err error) {
	err = DB.Where("id = ? ", model.Id).Find(&user).Error
	return
}

// 根据用户名查询用户
func (model User) FindByName() (user User, err error) {
	err = DB.Where("name = ?", model.Name).Find(&user).Error
	return
}

// 创建用户
func (model User) Create() (user User, err error) {
	err = DB.Create(&model).Error
	return model, err
}

// 更新用户信息
func (model User) Update() (err error) {
	err = DB.Model(&model).Select(UserField).Updates(model).Error
	return
}

// 更新用户状态信息 -1 删除(注销账户) 0 禁用 1 正常
func (model User) UpdateStatus() error {
	return DB.Model(&model).Update("status", model.Status).Error
}

// 列表(模糊)查询
func (model User) SelectUserList(page, size int, name string) (users []UserInfo, total int64, err error) {
	tx := DB.Model(model).
		Select(`id, 
		name, 
		subject, 
		password, 
		machine_list, 
		machine_number, 
		role_id, 
		description, 
		status, 
		login_ip, 
		login_time, 
		create_time, 
		update_time`).
		Where("name != ?", "admin")
	if name != "" {
		tx = tx.Where("name like ?", "%"+name+"%")
	}
	err = tx.Count(&total).
		Offset((page - 1) * size).
		Limit(size).
		Scan(&users).Error
	return
}

// 删除用户
func (model User) Delete() error {
	return DB.Model(&model).Where("id = ?", model.Id).Delete(model).Error
}

// 查询所有用户信息
func (model User) FindAll() (users []UserInfo, err error) {
	err = DB.Instance.Raw("SELECT * FROM ip_user").Scan(&users).Error
	return
}
