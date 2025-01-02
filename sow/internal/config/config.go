package config

import (
	"log/slog"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Environment       string     `mapstructure:"environment" json:"environment,omitempty"`
	LogLevel          slog.Level `mapstructure:"log_level" json:"log_level,omitempty"`
	LogStructured     bool       `mapstructure:"log_structured" json:"log_structured,omitempty"`
	LogAddSource      bool       `mapstructure:"log_add_source" json:"log_add_source,omitempty"`
	LogRequestId      bool       `mapstructure:"log_request_id" json:"log_request_id,omitempty"`
	Host              string     `mapstructure:"host" json:"host,omitempty"`
	Port              int        `mapstructure:"port" json:"port,omitempty"`
	TlsEnabled        bool       `mapstructure:"tls_enabled" json:"tls_enabled,omitempty"`
	TlsCertificate    string     `mapstructure:"tls_certificate" json:"tls_certificate,omitempty"`
	TlsCertificateKey string     `mapstructure:"tls_certificate_key" json:"tls_certificate_key,omitempty"`
	DbInMemory        bool       `mapstructure:"db_in_memory" json:"db_in_memory,omitempty"`
	// No default will be set for these three
	DbHost     string `mapstructure:"db_host" json:"db_host,omitempty"`
	DbUser     string `mapstructure:"db_user" json:"db_user,omitempty"`
	DbPassword string `mapstructure:"db_password" json:"db_password,omitempty"`
}

func NewConfig(debug bool) (*Config, error) {
	// Setup
	base := &Config{}
	vip := viper.New()

	// Enable ENV var reading
	vip.AutomaticEnv()
	vip.SetEnvPrefix("SOW")
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Establish sane defaults before overriding
	vip.SetDefault("environment", "dev")
	vip.SetDefault("log_level", slog.LevelDebug)
	vip.SetDefault("log_structured", true)
	vip.SetDefault("log_add_source", true)
	vip.SetDefault("log_request_id", true)
	vip.SetDefault("host", "localhost")
	vip.SetDefault("port", 9099)
	vip.SetDefault("tls_enabled", false)
	vip.SetDefault("tls_certificate", "cert.crt")
	vip.SetDefault("tls_certificate_key", "cert.key")
	vip.SetDefault("db_in_memory", true)
	vip.BindEnv("db_host")
	vip.BindEnv("db_user")
	vip.BindEnv("db_password")

	// Magic to unamrshal viper into the config sturct. The decode hook is used to map things like the logging level
	// into the slog logging level type.
	if err := vip.Unmarshal(&base, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc())); err != nil {
		return &Config{}, err
	}

	// Forecefully override if we're on debug mode
	// Do this without viper/cobra buy-in and just do the simple
	// brute force
	if debug {
		base.LogLevel = slog.LevelDebug
	}

	return base, nil
}
