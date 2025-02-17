package config

type Setting struct {
	Name            string `yaml:"name" json:"name"`
	Enable          bool   `yaml:"enable" json:"enable"`
	Order           int    `yaml:"order" json:"order"`
	Title           string `yaml:"title" json:"title"`
	Prompt          string `yaml:"prompt" json:"prompt"`
	Slogan          string `yaml:"slogan" json:"slogan"`
	AccessKeyId     string `yaml:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" json:"access_key_secret"`
	AgentKey        string `yaml:"agent_key" json:"agent_key"`
	AppId           string `yaml:"app_id" json:"app_id"`
}

type ModelOption struct {
	Label    string `yaml:"label" json:"label"`
	Value    string `yaml:"value" json:"value"`
	Disabled bool   `yaml:"disabled" json:"disabled"`
}

type SessionSetting struct {
	ChatScope    int `yaml:"chat_scope" json:"chat_scope"`
	SessionScope int `yaml:"session_scope" json:"session_scope"`
	DayScope     int `yaml:"day_scope" json:"day_scope"`
}

type BigModel struct {
	Setting        Setting        `yaml:"setting"`         // 对话积分消耗
	ModelList      []ModelOption  `yaml:"model_list"`      // 会话的积分消耗
	SessionSetting SessionSetting `yaml:"session_setting"` // 每日赠送积分
}
