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


### Example:

![alt text](https://github.com/gtforge/rex_common/blob/master/gorm_log/readme_files/example.png "Example")
![alt text](https://github.com/gtforge/rex_common/blob/master/gorm_log/readme_files/example.png "Example")
![alt text](https://github.com/gtforge/rex_common/blob/master/gorm_log/readme_files/example.png "Example")


### What's going on:

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

