# About

[![Build](https://github.com/rgl/dumpargs/actions/workflows/build.yml/badge.svg)](https://github.com/rgl/dumpargs/actions/workflows/build.yml)

Dump all the process environment variables and cli arguments.

## Usage

Download a [release](https://github.com/rgl/dumpargs/releases), extract, and execute, e.g., in a Windows MSYS2 bash session:

```bash
curl -sL https://github.com/rgl/dumpargs/releases/download/v0.0.1/dumpargs_0.0.1_windows_amd64v3.tar.gz | tar xzf -
MSYS2_ARG_CONV_EXCL='-subj=' ./dumpargs.exe -subj=/CN=example
```
