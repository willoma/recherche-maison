package config

const (
	DBPath     = "recherche-maison.db"
	DBOptions  = "_pragma=foreign_keys(1)&_time_format=sqlite"
	UploadsDir = "uploads"

	Port = 8910

	MaxUploadSize = 100 * 1024 * 1024 // 100 MB
)
