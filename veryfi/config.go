package veryfi

import (
	"time"

	"github.com/creasty/defaults"
	"github.com/pkg/errors"
)

// Options is the root config options
type Options struct {
	// EnvironmentURL provided by Veryfi without trailing `/` or http scheme.
	EnvironmentURL string `default:"api.veryfi.com"`

	// ClientID provided by Veryfi.
	ClientID string `default:"-"`

	// Username provided by Veryfi.
	Username string `default:"-"`

	// APIKey provided by Veryfi.
	APIKey string `default:"-"`

	// HTTP specifies the options for http protocol, used by a http client.
	HTTP HTTPOptions
}

// HTTPOptions is the config options for http protocol,
type HTTPOptions struct {
	// Timeout specifies a time limit for a http request.
	Timeout time.Duration `default:"3s"`

	// Retry specifies the options for retry mechanism.
	Retry RetryOptions
}

// RetryOptions is the config options for backoff retry mechanism. Its strategy
// is to increase retry intervals after each failed attempt, until some maximum
// value.
type RetryOptions struct {
	// Count specifies the number of retry attempts. Zero value means no retry.
	Count uint `default:"3"`

	// WaitTime specifies the wait time before retrying request. It is
	// increased after each attempt.
	WaitTime time.Duration `default:"100ms"`

	// MaxWaitTime specifies the maximum wait time, the cap, of all retry
	// requests that are made.
	MaxWaitTime time.Duration `default:"3s"`
}

// setDefaults setups default options.
func setDefaults(opts *Options) error {
	if opts == nil {
		return errors.New("options can not be nil")
	}

	if err := defaults.Set(opts); err != nil {
		return errors.New("failed to set default configs")
	}

	return nil
}
