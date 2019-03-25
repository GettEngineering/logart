# logart 

### _"logs as an art"_

This package includes:

- nice wrapper for logrus [read more](./log-art)

- logrus human readable formatter [read more](./log-formatters/logrus-human-formatter/)

- logrus JSON formatter [read more](./log-formatters/logrus-json-formatter/)

- gorm formatter [read more](./log-formatters/gorm-formatter/)

- errors handling package [read more](./err-art)

Each one is fully independent and can be used separately. But to get
the best possible experience try to use log-formatters together with the log.
In addition, packages independence allows move from already used logging or/and
formatting tools in very gradual way.
