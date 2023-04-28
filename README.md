# TunnelHop Types

This repo contains shared types between TunnelHop projects compiled into multiple languages. 

# How to use this repo in other repos
In order to leverage this repo in other repos you must install the package via `NPM`.

https://www.npmjs.com/package/tunnelhop-common

Install with the following:
```bash
npm i tunnelhop-common
```

from here, you should be able to leverage types from the `common` package.

# Cutting a release

1. Update the version in `package.json`
2. Tag a new release `git tag -a v<SEMANTIC_VERSION> -m "<YOUR_MESSAGE>"`

You can follow semantic version rules stated from [semver.](semver.org)



# Development for this repo

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

## Generate code
### Typescript
```bash
$ json2ts server.json Server.ts 
```

### Go
```bash
$ json2ts server.json Server.ts 
```

