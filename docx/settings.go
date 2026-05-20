package docx

import (
	"fmt"
	"strings"
)

const settingsPath = "word/settings.xml"

func (rd *RootDoc) SetUpdateFieldsOnOpen(value bool) {
	val := "false"
	if value {
		val = "true"
	}

	raw, ok := rd.FileMap.Load(settingsPath)
	if !ok {
		rd.FileMap.Store(settingsPath, []byte(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><w:settings xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:updateFields w:val="%s"/></w:settings>`, val)))
		return
	}

	settings := string(raw.([]byte))
	updateFields := fmt.Sprintf(`<w:updateFields w:val="%s"/>`, val)
	if strings.Contains(settings, "<w:updateFields") {
		start := strings.Index(settings, "<w:updateFields")
		end := strings.Index(settings[start:], "/>")
		if start >= 0 && end >= 0 {
			end += start + len("/>")
			settings = settings[:start] + updateFields + settings[end:]
		}
	} else if idx := strings.LastIndex(settings, "</w:settings>"); idx >= 0 {
		settings = settings[:idx] + updateFields + settings[idx:]
	}

	rd.FileMap.Store(settingsPath, []byte(settings))
}
