package scheme

// Error describes an error response.
type Error struct {
	HTTPCode    int    `json:"http_code" yaml:"http_code" mapstructure:"http_code"`
	Description string `json:"description" yaml:"description" mapstructure:"description"`
	Timestamp   string `json:"timestamp" yaml:"timestamp" mapstructure:"timestamp"`
	Context     string `json:"context" yaml:"context" mapstructure:"context"`
}
