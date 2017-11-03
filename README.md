# omnitools
==============

Tool to retrieve amount of credit on your account for omnivoice.

## OS & Arch Support
* Linux - amd64, 386, arm, arm64, mips, mipsle, ppc64, ppc64le, s390x
* Windows - 386, amd64
* OS X - 386, amd64
* Dragonfly - amd64
* FreeBSD - 386, amd64, arm
* NetBSD - 386, amd64, arm
* openBSD - 386, amd64
* Solaris - amd64

## Build
In most cases this is enough to build:
```
go get github.com/eservicesgreece/omnitools
```

## Usage
```bash
usage: omnitool [<flags>] <command> [<args> ...]

esgtools v

  Build:
  GIT:
  Copyright (c) 2016-8 eServices Greece - https://eservices-greece.com

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  account [<flags>] <accountuname> [<password>]
    account

  configdump
    Dump all entries in configuration                                          
```

## Configuration
The tool retrieves its configuration from the omnitool.json file, you can place it in your $HOME or /etc/omnitool/ .

