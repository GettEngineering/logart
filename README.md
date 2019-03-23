# logart 

### _"logs as an art"_

This package includes:

- nice wrapper for logrus [read more](./log)

- logrus formatter [read more](./formatters/logrus-formatter/)

- gorm formatter [read more](./formatters/gorm-formatter/)

Each one is fully independent and can be used separately. But to get
the best possible experience try to use formatters together with the log.
In addition, packages independence allows move from already used logging or/and
formatting tools in very gradual way.
