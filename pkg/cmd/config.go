package cmd

// Config holds the configuration for folder and git repository creation
type Config struct {
	// Path is the absolute or relative path where the folder should be created. This is not populated directly from command line arguments.
	Path string
	// Name is the name of the folder to be created
	Name string
	// FullName is the full name of the folder to be created
	FullName string
	// ParentPath is the absolute or relative path where the folder should be created
	ParentPath string
}
