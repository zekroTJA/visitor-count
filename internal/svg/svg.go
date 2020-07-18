package svg

import (
	"fmt"
	"strconv"
)

const (
	defaultTemplate = `<?xml version="1.0" encoding="UTF-8"?>
	<svg height="40px" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
		<title>Count</title>
	
		<g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
			<rect id="Rectangle" fill="#EEEEEE" x="0" y="0" width="110" height="40"></rect>
			<text id="0" font-family="sans-serif" font-weight="bold" font-size="18" fill="#455A64">
				<tspan x="13" y="26">VISITORS</tspan>
			</text>
	
			<rect id="Rectangle" fill="#03A9F4" x="110" y="0" width="%d" height="40"></rect>
			<text id="0" font-family="sans-serif" font-weight="bold" font-size="18" fill="#EEEEEE">
				<tspan x="123" y="26">%s</tspan>
			</text>
		</g>
	</svg>
	`
)

var (
	template = defaultTemplate
)

func GetFormattedSVG(count int) []byte {
	const minWidth, charWidth = 25, 11

	sCount := strconv.Itoa(count)
	width := minWidth + (len(sCount) * charWidth)
	res := fmt.Sprintf(template, width, sCount)

	return []byte(res)
}
