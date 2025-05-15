package cmd

// Config holds the configuration for folder and git repository creation
type Config struct {
	// Path is the absolute or relative path where the folder should be created
	Path     string
	Name     string
	FullName string
}
