package templateFunc

import (
	. "controllers"
	"golanger.com/framework/utils"
)

func init() {
	App.AddTemplateFunc("UnixToStr", func(tm int64) string {
		return utils.NewTime().UnixToStr(tm)
	})
}
