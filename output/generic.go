package output

import "time"

type OutputConfig struct {
	Destination string
	Format string
	Bulk_Support bool
	Bulk_Max_Size uint
	Bulk_Max_Duration time.Duration
	Http *HttpConfig
}

type HttpConfig struct {
	Method string
	Auth string
	Username string
	Password string
	Headers []*HttpHeader
}

type HttpHeader struct {
	K string
	V string
}