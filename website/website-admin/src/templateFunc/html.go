package templateFunc

import (
	. "controllers"
	"helper"
	"strconv"
)

func init() {
	App.AddTemplateFunc("OptionsForSelect", func(m interface{}, def string) string {
		var html string

		switch m.(type) {
		case []string:
			for k, v := range m.([]string) {
				if v == "" {
					continue
				}

				ks := strconv.Itoa(k)
				html += `<option`
				if ks == def {
					html += ` selected `
				}

				html += ` value="` + ks + `">` + v + `</option>`
			}
		case []map[string]string:
			for _, v := range m.([]map[string]string) {
				html += `<option`
				if v["value"] == def {
					html += ` selected `
				}

				html += ` value="` + v["value"] + `">` + v["name"] + `</option>`
			}
		}

		return html
	})
	App.AddTemplateFunc("FilterHtmlTag", helper.FilterHtmlTag)
}
