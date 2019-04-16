package main

import (
        "context"

        "github.com/rghetia/gomodule/gcp"
)

func main () {
        gcp.DumpMetadata(context.Background())
}

