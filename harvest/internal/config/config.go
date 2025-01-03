package config

import (
	"log/slog"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Token             string     `mapstructure:"token" json:"token,omitempty"`
	Environment       string     `mapstructure:"environment" json:"environment,omitempty"`
	LogLevel          slog.Level `mapstructure:"log_level" json:"log_level,omitempty"`
	LogStructured     bool       `mapstructure:"log_structured" json:"log_structured,omitempty"`
	LogAddSource      bool       `mapstructure:"log_add_source" json:"log_add_source,omitempty"`
	LogRequestId      bool       `mapstructure:"log_request_id" json:"log_request_id,omitempty"`
	SowAddress        string     `mapstructure:"sow_address" json:"host,omitempty"`
	TlsEnabled        bool       `mapstructure:"tls_enabled" json:"tls_enabled,omitempty"`
	TlsCertificate    string     `mapstructure:"tls_certificate" json:"tls_certificate,omitempty"`
	TlsCertificateKey string     `mapstructure:"tls_certificate_key" json:"tls_certificate_key,omitempty"`
}

func NewConfig(debug bool) (*Config, error) {
	// Setup
	base := &Config{}
	vip := viper.New()

	// Enable ENV var reading
	vip.AutomaticEnv()
	vip.SetEnvPrefix("HARVEST")
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Establish sane defaults before overriding
	vip.SetDefault("environment", "dev")
	vip.SetDefault("log_level", slog.LevelDebug)
	vip.SetDefault("log_structured", true)
	vip.SetDefault("log_add_source", true)
	vip.SetDefault("log_request_id", true)
	vip.SetDefault("sow_address", "localhost:9099")
	vip.SetDefault("tls_enabled", false)
	vip.SetDefault("tls_certificate", "cert.crt")
	vip.SetDefault("tls_certificate_key", "cert.key")
	vip.BindEnv("token")

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
