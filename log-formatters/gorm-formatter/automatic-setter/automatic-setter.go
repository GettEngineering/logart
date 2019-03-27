package automaticsetter

import (
	"github.com/gtforge/logart/log-formatters/gorm-formatter"
	"github.com/gtforge/services_common_go/gett-storages"
)

func init() {
	gettStorages.DB.SetLogger(gormlogformatter.DefaultLogger())
}
