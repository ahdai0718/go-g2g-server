# G2G Server

## Reference

- Go

  - [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)

  - [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

- Protocol Buffer

  - [https://developers.google.com/protocol-buffers/](https://developers.google.com/protocol-buffers/)

  - [https://github.com/golang/protobuf](https://github.com/golang/protobuf)

  - [Code Style Guide](https://developers.google.com/protocol-buffers/docs/style)

## Code Style Guide

- [Code Style Guide](https://github.com/ahdai0718/code-style-guide)

## Architectures

- `cmd`

  - `bot`

  - `server`

    - `game`

    - `gateway`

- `internal`

  - `app`

    - `bot`

    - `game`

    - `gateway`

  - `pkg`

    - `common`

    - `constant`

    - `core`

      - `game`

        - `example`

        - `simple_factory.go`

      - `player`

      - `room`

      - `serverinfo`

    - `network`

    - `pb` protocol buffer

      - `game`

        - `example`

      - `pb.sh`

    - `platform`

    - `store`

    - `util`

## Todos

- Use `sync.Map` for concurrent safe map.

- Use `atomic` instead of channel in some cases.

- Use `gorm` for DB implement.
