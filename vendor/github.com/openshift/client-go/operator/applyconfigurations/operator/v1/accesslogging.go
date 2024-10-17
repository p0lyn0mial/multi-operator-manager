// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// AccessLoggingApplyConfiguration represents a declarative configuration of the AccessLogging type for use
// with apply.
type AccessLoggingApplyConfiguration struct {
	Destination        *LoggingDestinationApplyConfiguration                  `json:"destination,omitempty"`
	HttpLogFormat      *string                                                `json:"httpLogFormat,omitempty"`
	HTTPCaptureHeaders *IngressControllerCaptureHTTPHeadersApplyConfiguration `json:"httpCaptureHeaders,omitempty"`
	HTTPCaptureCookies []IngressControllerCaptureHTTPCookieApplyConfiguration `json:"httpCaptureCookies,omitempty"`
	LogEmptyRequests   *operatorv1.LoggingPolicy                              `json:"logEmptyRequests,omitempty"`
}

// AccessLoggingApplyConfiguration constructs a declarative configuration of the AccessLogging type for use with
// apply.
func AccessLogging() *AccessLoggingApplyConfiguration {
	return &AccessLoggingApplyConfiguration{}
}

// WithDestination sets the Destination field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Destination field is set to the value of the last call.
func (b *AccessLoggingApplyConfiguration) WithDestination(value *LoggingDestinationApplyConfiguration) *AccessLoggingApplyConfiguration {
	b.Destination = value
	return b
}

// WithHttpLogFormat sets the HttpLogFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HttpLogFormat field is set to the value of the last call.
func (b *AccessLoggingApplyConfiguration) WithHttpLogFormat(value string) *AccessLoggingApplyConfiguration {
	b.HttpLogFormat = &value
	return b
}

// WithHTTPCaptureHeaders sets the HTTPCaptureHeaders field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPCaptureHeaders field is set to the value of the last call.
func (b *AccessLoggingApplyConfiguration) WithHTTPCaptureHeaders(value *IngressControllerCaptureHTTPHeadersApplyConfiguration) *AccessLoggingApplyConfiguration {
	b.HTTPCaptureHeaders = value
	return b
}

// WithHTTPCaptureCookies adds the given value to the HTTPCaptureCookies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HTTPCaptureCookies field.
func (b *AccessLoggingApplyConfiguration) WithHTTPCaptureCookies(values ...*IngressControllerCaptureHTTPCookieApplyConfiguration) *AccessLoggingApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHTTPCaptureCookies")
		}
		b.HTTPCaptureCookies = append(b.HTTPCaptureCookies, *values[i])
	}
	return b
}

// WithLogEmptyRequests sets the LogEmptyRequests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogEmptyRequests field is set to the value of the last call.
func (b *AccessLoggingApplyConfiguration) WithLogEmptyRequests(value operatorv1.LoggingPolicy) *AccessLoggingApplyConfiguration {
	b.LogEmptyRequests = &value
	return b
}