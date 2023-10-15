# FritzGo

Retrieve basic information from your FRITZ!Box.

```
$ go run ./cmd/fritzgo/ --help

NAME:
   FritzGo - A new cli application

USAGE:
   FritzGo [global options] command [command options] [arguments...]

VERSION:
   <unknown> @ 1970-01-01T01:00:00+01:00

COMMANDS:
   users    
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h                           show help
   --log.ctx value [ --log.ctx value ]  A list of context field appended to every log. Format: key=value. [$LOG_CTX]
   --log.format value                   Specify the format of logs. Supported formats: 'logfmt', 'json', 'console' [$LOG_FORMAT]
   --log.level value                    Specify the log level. e.g. 'debug', 'info', 'error'. (default: "info") [$LOG_LEVEL]
   --version, -v                        print the version

   FRITZ!Box:

   --fritz.url value  The FRITZ!Box address. (default: "http://fritz.box") [$FRITZ_URL]
```

## List users

List fritz users.

```
$ go run ./cmd/fritzgo/ users list

+-----------+-------------+
| USER NAME | IS DEFAULT? |
+-----------+-------------+
| fritz7517 | true        |
+-----------+-------------+
```