package main

import (
	"os"
	"time"

	"github.com/m-awwab-khan/learning-go/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
