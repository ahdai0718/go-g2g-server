#!/bin/bash
cd "$(dirname "$0")"

protoc --go_out=paths=source_relative:. *.proto
protoc --go_out=paths=source_relative:. gen_bq_schema/*.proto
protoc --go_out=paths=source_relative:. game/*.proto
protoc --go_out=paths=source_relative:. game/**/*.proto
# protoc --go_out=paths=source_relative:. game/**/**/*.proto

protoc --bq-schema_out=paths=source_relative:../logger/bigquery platform_log.proto
protoc --bq-schema_out=paths=source_relative:../logger/bigquery runtime_log.proto