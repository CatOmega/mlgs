//go:generate gen_static_data_go
//go:generate protoc --enum-go_out=. enum.proto global.proto
package sd
