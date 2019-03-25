# gorm formatter


### Usage:


```
// Default:
gorm.DB.SetLogger(gormlog.DefaultFormated())

// Custom:
o := DefaultFormatOptions
o.LogLevelColor = 123
o.DontShowDate = false
gorm.DB.SetLogger(gormlog.CustomFormated(o))
```


Pay attention, if `*gorm.DB` is cloned, for example by using:
```
var Impl = impl{
    db: storages.DB.Table("table_name"), // <-- gorm.DB is cloned here
}

type impl struct {
    db *gorm.DB
}
```
This import should be done directly in package


### Options

- SourceFilePathDepth  `default = 3` - Show only last N parts of the source
file path

- DontShowDate         `default = true` - print only time. Generally, date is not needed
in stage/dev environments.

- LogLevelName         `default = "SQL"`

- color options (see available colors [here](https://github.com/artiomgiza/go-color-256)):
    - EnableColors         `default = true`
    - TimeColor            `default = 23`
    - DurationColor        `default = 23`
    - LogLevelColor        `default = 23`
    - RowsNoAffectedColor  `default = 23`
    - RowsYesAffectedColor `default = 35`
    - FilePathColor        `default = 23`
    - QueryColor           `default = 60`


### Example:

- Short and humble (as we are) formatted log (made by this formatter)

![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/gorm-formatter/readme_files/logart_gorm_formatter.png "Example")

- Default (long and too colored) gorm formatted log:

![alt text](https://github.com/gtforge/logart/blob/master/log-formatters/gorm-formatter/readme_files/default_gorm_formatter.png "Example")


### Technical details (advanced reading):

Currently gorm don't really has ability to change the log format.

We only can replace logger (struct that implements `Print(values ...interface{})`)
After logger is replaced (by calling: `gettStorages.DB.SetLogger(...)`)
Each time, gorm log is printed, the  Print(values...) is called.

The question is what are these `values`?

When SQL log is printed values are:
```
[]interface {}{
    "sql",
    "/Users/gizatullinartiom/go/src/.../services_common_go/gett-settings/settings.go:181",
    954153,
    "SELECT * FROM \"table_name\" WHERE $1 > 0",
    []interface {}{"field_name"},
    666,
}
```

We can see that there are 6 lines:

1. "sql" string
2. path to file
3. something not clear...
4. sql sequence with $1, $2, ...
5. this params that should replace $1, $2, ...
6. number

There is no too much we can do with this. So we would not work directly
with this values, but on formatted (by gorm formatter) values:

```
vals := gorm.LogFormatter(values...)
```

These `vals` look like this:
```
[]interface {}{
    "\x1b[35m(/Users/gizatullinartiom/go/src/.../services_common_go/gett-settings/settings.go:181)\x1b[0m",
    "\n\x1b[33m[2019-02-26 00:19:25]\x1b[0m",
    " \x1b[36;1m[2.70ms]\x1b[0m ",
    "SELECT * FROM \"table_name\" WHERE 'field_name' > 0 ",
    " \n\x1b[36;31m[666 rows affected or returned ]\x1b[0m ",
}
```


We can see that there are 5 lines:

1. path to file2
2. date/time
3. run duration
4. sql sequence
5. number of raws affected

Formatter converts this data to print nicer configurable log.

