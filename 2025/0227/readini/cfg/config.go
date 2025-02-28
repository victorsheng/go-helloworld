package cfg

type ConfigSt struct {
	Server       serverSt       `json:"server" ini:"server"`
	StreamEngine StreamEngineSt `json:"stream_engine" ini:"stream_engine"`
	OfflineCfgSt EnableCfgSt    `ini:"offline_config"`
	AccountCfgSt EnableCfgSt    `ini:"account_config"`
	// KafkaOutput                 kafkacfg.KafkaCfgSt  `json:"kafkaoutput" ini:"kafkaoutput"`
	FileOutputSensors           FileCfgSt `json:"file_output_sensors" ini:"file_output_sensors"`
	FileOutputPerformance       FileCfgSt `json:"file_output_performance" ini:"file_output_performance"`
	FileOutputAction            FileCfgSt `json:"file_output_action" ini:"file_output_action"`
	FileOutputZuiyouAction      FileCfgSt `json:"file_output_zuiyou_action" ini:"file_output_zuiyou_action"`
	FileOutputZuiyouSpeedAction FileCfgSt `json:"file_output_zuiyou_speed_action" ini:"file_output_zuiyou_speed_action"`
	FileOutputPipiAction        FileCfgSt `json:"file_output_pipi_action" ini:"file_output_pipi_action"`
	FileOutputNewPipiAction     FileCfgSt `json:"file_output_new_pipi_action" ini:"file_output_new_pipi_action"`
	FileOutputHanabiAction      FileCfgSt `json:"file_output_hanabi_action" ini:"file_output_hanabi_action"`
	FileOutputOmgAction         FileCfgSt `json:"file_output_omg_action" ini:"file_output_omg_action"`
	FileOutputTiantianAction    FileCfgSt `json:"file_output_tiantian_action" ini:"file_output_tiantian_action"`
	// StatCache                   rediscfg.RedisSt     `json:"stat_cache" ini:"stat_cache"`
	// UserProfileCache            rediscfg.RedisSt     `json:"user_profile_cache" ini:"user_profile_cache"`
	MySQL                MySQLCfgSt           `json:"my_sql" ini:"my_sql"`
	Presto               PrestoCfgSt          `json:"presto" ini:"presto"`
	Sensor               SensorCfgSt          `json:"sensor" ini:"sensor"`
	SyncUserProfileJobSt SyncUserProfileJobSt `ini:"sync_user_profile_job"`
	APK                  APKCfgSt             `json:"apk" ini:"apk"`

	// service discovery
	// ServiceDiscovery svcsd.ServiceDiscoverySt `json:"service-discovery" ini:"service-discovery"`
}

type EnableCfgSt struct {
	Enable bool `ini:"enable"`
}

type StreamSourceKafka struct {
	KafkaHostPorts string `ini:"source.kafka.hostports"`
	KafkaTopics    string `ini:"source.kafka.topics"`
	KafkaDecoder   string `ini:"source.kafka.decoder"`
	KafkaGroupID   string `ini:"source.kafka.groupid"`
}

type StreamSourceFile struct {
	FilePath     string `ini:"source.file.path"`
	FilePattern  string `ini:"source.file.pattern"`
	FileCategory string `ini:"source.file.category"`
}

type StreamEngineSt struct {
	Enable     bool   `ini:"enable"`
	SourceType string `ini:"source.type"`
	StreamSourceKafka
	StreamSourceFile
	Processors string `ini:"processors"`
}

type FileCfgSt struct {
	Filename string `json:"filename" ini:"filename"`
	Enable   bool   `json:"enable" ini:"enable"`
}

type PluginsSt struct {
	Start string `json:"start" ini:"start"`
}

type serverSt struct {
	Controllers   string `json:"controllers" ini:"controllers"`
	LogLevel      string `json:"log-level" ini:"log-level"`
	PrintReply    bool   `json:"print-reply" ini:"print-reply"`
	EnableDegrade bool   `json:"enable-degrade" ini:"enable-degrade"`

	UpstreamSetting string `json:"upstream-setting" ini:"upstream-setting"`
	SwitcherList    string `json:"switchers" ini:"switchers"`
}

type MySQLCfgSt struct {
	Host     string `json:"host" ini:"host"`
	User     string `json:"user" ini:"user"`
	Password string `json:"password" ini:"password"`
	DB       string `json:"db" ini:"db"`
}

type PrestoCfgSt struct {
	Host    string `json:"host" ini:"host"`
	User    string `json:"user" ini:"user"`
	Catalog string `json:"catalog" ini:"catalog"`
	Schema  string `json:"schema" ini:"schema"`
}

type SensorCfgSt struct {
	Host       string `json:"host" ini:"host"`
	Project    string `json:"project" ini:"project"`
	Token      string `json:"token" ini:"token"`
	SQLProject string `json:"sql_project" ini:"sql_project"`
	SQLToken   string `json:"sql_token" ini:"sql_token"`
}

type SyncUserProfileJobSt struct {
	Path string `ini:"path"`
}

type APKCfgSt struct {
	OutputDir string `json:"output_dir" ini:"output_dir"`
	JAR       string `json:"jar" ini:"jar"`
	OSSBin    string `json:"oss_bin" ini:"oss_bin"`
}
