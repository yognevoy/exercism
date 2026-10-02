package sorting

import (
    "fmt"
    "strconv"
)

func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
    return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
    box, ok := fnb.(FancyNumber)
    if ok {
        i, _ := strconv.Atoi(box.Value())
        return i
    }
    return 0
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    return fmt.Sprintf(
        "This is a fancy box containing the number %.1f", 
        float64(ExtractFancyNumber(fnb)),
    )
}

func DescribeAnything(i any) string {
    switch v := i.(type) {
    case int:
        return DescribeNumber(float64(v))
    case float64:
        return DescribeNumber(v)
    case NumberBox:
        return DescribeNumberBox(v)
    case FancyNumberBox:
        return DescribeFancyNumberBox(v)
    default:
        return "Return to sender"
    }
}
