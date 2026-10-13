// Package twofer implements a simple sharing phrase generator.
package twofer

import "fmt"

// ShareWith returns a string detailing how to share a cookie with the given name.
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}
