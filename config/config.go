package config

const (
	crdFile = ".cred.yaml"
	optFile = "option.yaml"
	colFile = "color.yaml"
)

// Config 設定構造体
type Config struct {
	Cred   *cred
	Option *option
	Color  *color
}

type cred struct {
	Token  string
	Secret string
}

type option struct {
	ConfigDir  string
	Counts     string
	DateFormat string
	TimeFormat string
}

type color struct {
	Accent1      int32
	Accent2      int32
	Accent3      int32
	Dim          int32
	BoxForground int32
	Username     int32
	Screenname   int32
	Reply        int32
	Hashtag      int32
	Favorite     int32
	Retweet      int32
	Verified     int32
	Protected    int32
	Follow       int32
	Block        int32
	Mute         int32
}

// NewConfig 設定構造体作成
func NewConfig() *Config {
	cfg := &Config{
		Cred: &cred{
			Token:  "",
			Secret: "",
		},
		Option: &option{
			ConfigDir:  getConfigDir(),
			Counts:     "25",
			DateFormat: "2006/01/02",
			TimeFormat: "15:04:05",
		},
		Color: &color{
			Accent1:      0xe06c75,
			Accent2:      0xc678dd,
			Accent3:      0x56b6c2,
			Dim:          0x343a44,
			BoxForground: 0x000000,
			Username:     0xffffff,
			Screenname:   0x9c9c9c,
			Reply:        0x56b6c2,
			Hashtag:      0x61afef,
			Favorite:     0xe887b9,
			Retweet:      0x98c379,
			Verified:     0x5685d1,
			Protected:    0x787878,
			Follow:       0x1877c9,
			Block:        0xe06c75,
			Mute:         0xe5C07b,
		},
	}
	return cfg
}

// Save 保存
func (cfg *Config) Save() {
	saveYaml(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	saveYaml(cfg.Option.ConfigDir, optFile, cfg.Option)
	saveYaml(cfg.Option.ConfigDir, colFile, cfg.Color)
}

// Load 読込
func (cfg *Config) Load() bool {
	if !configFileExists(cfg.Option.ConfigDir) {
		return false
	}
	loadYaml(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	loadYaml(cfg.Option.ConfigDir, optFile, cfg.Option)
	loadYaml(cfg.Option.ConfigDir, colFile, cfg.Color)
	return true
}
