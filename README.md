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
usage: omnitool_windows_amd64.exe [<flags>] <command> [<args> ...]

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and
                 --help-man).
  -v, --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  account [<flags>] <accountuname> [<password>]
    account

  configdump
    Dump all entries in configuration                                    
```

##usage examples
#get full account (credits, account, tax rate)
```bash
./omnitool account 123456 1234567890
18.97,123456,24.00
```

#get full account (credits, account, tax rate) of stored account (in json config)
```bash
./omnitool account 123456
18.97,123456,24.00
```

#get just the credits
```bash
./omnitool account 123456 1234567890 -r credit
18.97
```

## Configuration
The tool retrieves its configuration from the omnitool.json file, you can place it in your $HOME or /etc/omnitool/ .

## omnitool.json  Example
```json
{                                    
    "appname": "omnitool",           
    "general": {},                   
    "accounts": {                    
        "999999": {                  
            "pass": "123456"       
        },                           
        "888888": {                  
            "pass": "123456"
        }                            
    }                                
}                                    
```
