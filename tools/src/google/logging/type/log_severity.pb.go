// Code generated by protoc-gen-go.
// source: google/logging/type/log_severity.proto
// DO NOT EDIT!

package _type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The severity of the event described in a log entry.  These guideline severity
// levels are ordered, with numerically smaller levels treated as less severe
// than numerically larger levels. If the source of the log entries uses a
// different set of severity levels, the client should select the closest
// corresponding `LogSeverity` value. For example, Java's FINE, FINER, and
// FINEST levels might all map to `LogSeverity.DEBUG`. If the original severity
// code must be preserved, it can be stored in the payload.
//
// (-- This list is intentionally not comprehensive; the intent is to provide a
// common set of "good enough" severity levels so that UIs can provide filtering
// and searching across log types. Users of the API are free not to use all
// severity levels in their log messages, and (if needed) to use a private field
// in their payload to store a more detailed severity code. --)
type LogSeverity int32

const (
	// The log entry has no assigned severity level.
	LogSeverity_DEFAULT LogSeverity = 0
	// Debug or trace information.
	LogSeverity_DEBUG LogSeverity = 100
	// Routine information, such as ongoing status or performance.
	LogSeverity_INFO LogSeverity = 200
	// Normal but significant events, such as start up, shut down, or
	// configuration.
	LogSeverity_NOTICE LogSeverity = 300
	// Warning events might cause problems.
	LogSeverity_WARNING LogSeverity = 400
	// Error events are likely to cause problems.
	LogSeverity_ERROR LogSeverity = 500
	// Critical events cause more severe problems or brief outages.
	LogSeverity_CRITICAL LogSeverity = 600
	// A person must take an action immediately.
	LogSeverity_ALERT LogSeverity = 700
	// One or more systems are unusable.
	LogSeverity_EMERGENCY LogSeverity = 800
)

var LogSeverity_name = map[int32]string{
	0:   "DEFAULT",
	100: "DEBUG",
	200: "INFO",
	300: "NOTICE",
	400: "WARNING",
	500: "ERROR",
	600: "CRITICAL",
	700: "ALERT",
	800: "EMERGENCY",
}
var LogSeverity_value = map[string]int32{
	"DEFAULT":   0,
	"DEBUG":     100,
	"INFO":      200,
	"NOTICE":    300,
	"WARNING":   400,
	"ERROR":     500,
	"CRITICAL":  600,
	"ALERT":     700,
	"EMERGENCY": 800,
}

func (x LogSeverity) String() string {
	return proto.EnumName(LogSeverity_name, int32(x))
}
func (LogSeverity) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterEnum("google.logging.type.LogSeverity", LogSeverity_name, LogSeverity_value)
}

func init() { proto.RegisterFile("google/logging/type/log_severity.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0xd7, 0x2f, 0xa9, 0x2c, 0x00, 0x73,
	0xe2, 0x8b, 0x53, 0xcb, 0x52, 0x8b, 0x32, 0x4b, 0x2a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x84, 0x21, 0xea, 0xf4, 0xa0, 0xea, 0xf4, 0x40, 0xea, 0xb4, 0x9a, 0x18, 0xb9, 0xb8, 0x7d, 0xf2,
	0xd3, 0x83, 0xa1, 0x4a, 0x85, 0xb8, 0xb9, 0xd8, 0x5d, 0x5c, 0xdd, 0x1c, 0x43, 0x7d, 0x42, 0x04,
	0x18, 0x84, 0x38, 0xb9, 0x58, 0x5d, 0x5c, 0x9d, 0x42, 0xdd, 0x05, 0x52, 0x84, 0x38, 0xb9, 0x58,
	0x3c, 0xfd, 0xdc, 0xfc, 0x05, 0x4e, 0x30, 0x0a, 0x71, 0x73, 0xb1, 0xf9, 0xf9, 0x87, 0x78, 0x3a,
	0xbb, 0x0a, 0xac, 0x61, 0x12, 0xe2, 0xe1, 0x62, 0x0f, 0x77, 0x0c, 0xf2, 0xf3, 0xf4, 0x73, 0x17,
	0x98, 0xc0, 0x2c, 0xc4, 0xc5, 0xc5, 0xea, 0x1a, 0x14, 0xe4, 0x1f, 0x24, 0xf0, 0x85, 0x59, 0x88,
	0x97, 0x8b, 0xc3, 0x39, 0xc8, 0x33, 0xc4, 0xd3, 0xd9, 0xd1, 0x47, 0xe0, 0x06, 0x0b, 0x48, 0xca,
	0xd1, 0xc7, 0x35, 0x28, 0x44, 0x60, 0x0f, 0xab, 0x10, 0x1f, 0x17, 0xa7, 0xab, 0xaf, 0x6b, 0x90,
	0xbb, 0xab, 0x9f, 0x73, 0xa4, 0xc0, 0x02, 0x36, 0x27, 0x5d, 0x2e, 0xf1, 0xe4, 0xfc, 0x5c, 0x3d,
	0x2c, 0xee, 0x73, 0x12, 0x40, 0x72, 0x5c, 0x00, 0xc8, 0x1b, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0xff,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf7, 0x59, 0x98, 0xf9, 0x00, 0x00, 0x00,
}