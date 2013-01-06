package templateFunc

import (
	. "controllers"
	"golanger.com/utils"
)

func init() {
	App.AddTemplateFunc("UnixToStr", func(tm int64) string {
		return utils.NewTime().UnixToStr(tm)
	})
}
