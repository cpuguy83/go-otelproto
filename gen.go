//go:generate protoc -I./proto/opentelemetry/proto --gofast_out=paths=source_relative:. ./proto/opentelemetry/proto/common/v1/common.proto
//go:generate protoc -I./proto/opentelemetry/proto -I./proto --gofast_out=paths=source_relative,Mopentelemetry/proto/common/v1/common.proto=github.com/cpuguy83/go-otelproto/common/v1:. ./proto/opentelemetry/proto/resource/v1/resource.proto
//go:generate protoc -I./proto/opentelemetry/proto -I./proto --gofast_out=paths=source_relative,Mopentelemetry/proto/resource/v1/resource.proto=github.com/cpuguy83/go-otelproto/resource/v1,Mopentelemetry/proto/common/v1/common.proto=github.com/cpuguy83/go-otelproto/common/v1:. ./proto/opentelemetry/proto/trace/v1/trace.proto

package otelproto
