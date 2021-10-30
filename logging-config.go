package logrusPlusPlus

type LoggingConfig struct {
	LogFileActive        bool   `yaml:"logFileActive" env:"LOG_FILE_ACTIVE" env-default:"false"`
	LogFilePath          string `yaml:"logFilePath" env:"LOG_FILE_PATH" env-default:"/var/log/"`
	LogFileMaxSizeMb     int    `yaml:"logFileMaxSizeMb" env:"LOG_FILE_MAX_SIZE_MB" env-default:"50"`
	LogFileMaxBackups    int    `yaml:"logFileMaxBackups" env:"LOG_FILE_MAX_BACKUPS" env-default:"4"`
	LogFileMaxAgeDay     int    `yaml:"logFileMaxAgeDay" env:"LOG_FILE_MAX_AGE_DAY" env-default:"14"`
	LogFileCompress      bool   `yaml:"logFileCompress" env:"LOG_FILE_COMPRESS" env-default:"false"`
	LogStdoutActive      bool   `yaml:"logStdoutActive" env:"LOG_STDOUT_ACTIVE" env-default:"true"`
	LogLevel             int    `yaml:"logLevel" env:"LOG_LEVEL" env-default:"6"`
	LogPrettyPrintActive bool   `yaml:"logPrettyPrintActive" env:"LOG_PRETTY_PRINT_ACTIVE" env-default:"false"`
}
