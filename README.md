# TunnelHop Types

This repo contains shared types between TunnelHop projects compiled into multiple languages. 

## Auto Generate Code
_See installation section below_

`make generate`

## Installation

Schema changes and additions must be manually generated and committed to the repository.

- Install Node/NPM
- Install Go

_go-jsonschema_
```
$ brew tap omissis/go-jsonschema
$ brew install go-jsonschema
```

_json2ts_
```
$ npm i -g json-schema-to-typescript
```

## Generate Code
### Typescript
```bash
$ json2ts server.json Server.ts 
```

### Go
```bash
$ json2ts server.json Server.ts 
```

