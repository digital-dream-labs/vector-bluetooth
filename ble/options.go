package ble

type options struct {
	outputDir string
}

// Option is the list of options
type Option func(*options)

// WithLogDirectory sets the logger output directory
func WithLogDirectory(l string) Option {
	return func(o *options) {
		o.outputDir = l
	}
}
