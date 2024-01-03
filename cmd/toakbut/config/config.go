package config

import "toakbut/pkg/common/db"

type Config struct {
	Name string `required:"true"`
	Host string `split_words:"true" default:"localhost"`
	Port int    `split_words:"true" default:"3000"`

	DB *db.Config `split_words:"true" required:"true"`

	Token     string `split_words:"true" required:"true"`
	ChannelID string `split_words:"true" required:"true"`
}
