/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import (
	"github.com/mingderwnag/parse"
)

var src = `
/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import ()

func main() {
	parseTypeNameFromComment(src)
}
`

func ExampleParse() {
	parseTypeNameFromComment(src)
}
