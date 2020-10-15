package config

type Profiles interface {
	GetInclude() string
}

type defaultProfiles struct {
	Include string `json:"include"`
}

// Include 包含的配置文件
// 名称前缀为"application-"，格式为yml，如："application-xxx.yml"
// 多个文件名以逗号隔开，并省略掉前缀"application-"，如：dn, jpush, mysql
func (p defaultProfiles) GetInclude() string {
	return p.Include
}
