# FritzGo

Retrieve basic information from your FRITZ!Box.

```shell
$ go run ./cmd/fritzgo/ help
NAME:
   FritzGo - CLI tool to access FRITZ!Box data

USAGE:
   FritzGo [global options] command [command options] [arguments...]

VERSION:
   <unknown> @ 1970-01-01T01:00:00+01:00

COMMANDS:
   info     
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

   Rendering style:

   --render.table.style value  The rendering table style. (default: "default") [$RENDER_TABLE_STYLE]
```

## Info

Display basic information.

```shell
$ go run ./cmd/fritzgo/ info

+---------------------------------+
| Info                            |
+-------------+-------------------+
| Name        | FRITZ!Box 7530    |
| Version     | 164.07.57         |
| Mac address | A3:7D:9B:C1:4E:2A |
| URL         | http://fritz.box  |
+-------------+-------------------+
```


## List users

Display fritz users.

```shell
$ go run ./cmd/fritzgo/ users list

+---------------------+
| List users          |
+-----------+---------+
| USER NAME | DEFAULT |
+-----------+---------+
| fritz5432 | true    |
+-----------+---------+
```

