package model

type Options struct {
	Host         string `default:"127.0.0.1"`
	Port         int    `default:"5432"`
	User         string `default:"postgres"`
	Password     string `default:"postgres"`
	Database     string `default:"postgres"`
	SSLMode      string `default:"disable"`
	FilePerTable bool   `default:"false"`
	PackageName  string `default:"models"`
	Dir          string `default:"models"`
	OneFileName  string `default:"models_gen"`
}
