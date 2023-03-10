package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Production bool          `envconfig:"PRODUCTION" default:"false"`
	LogLevel   zapcore.Level `envconfig:"LOG_LEVEL"  default:"info"`

	DatabaseURL         string `envconfig:"DATABASE_URL"          required:"true"`
	ListenAddr          string `envconfig:"LISTEN_ADDR"           default:"0.0.0.0:8000"`
	PublicWebAddress    string `envconfig:"PUBLIC_WEB_ADDRESS"    default:"http://localhost:3000"`
	HashKey             []byte `envconfig:"HASH_KEY"              required:"true"`
	BlockKey            []byte `envconfig:"BLOCK_KEY"             required:"true"`
	GithubClientID      string `envconfig:"GITHUB_CLIENT_ID"      required:"true"`
	GithubClientSecret  string `envconfig:"GITHUB_CLIENT_SECRET"  required:"true"`
	DiscordClientID     string `envconfig:"DISCORD_CLIENT_ID"     required:"true"`
	DiscordClientSecret string `envconfig:"DISCORD_CLIENT_SECRET" required:"true"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}
	return cfg, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
