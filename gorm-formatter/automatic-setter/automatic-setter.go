package automaticsetter

import (
	"github.com/gtforge/logart/gorm-formatter"
	"github.com/gtforge/services_common_go/gett-storages"
)

func init() {
	gettStorages.DB.SetLogger(gormlogformatter.DefaultLogger())
}
