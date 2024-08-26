package util

import (
	"context"

	"github.com/containerd/log"
)

func Test() {
	log.L.Error("This is a test error message")
	log.G(context.TODO())
}
