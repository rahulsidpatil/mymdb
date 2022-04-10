# mymdb Entry point

## Directory structure:
```
cmd
│   ├── main.go
│   └── README.md

```

## Overview

The main.go file is an entry point for mymdb application. It invokes application handler.
```
package main

import "github.com/rahulsidpatil/mymdb/pkg/app"

func main() {
	a := app.MyMdb{}
	a.Initialize()
	a.Run()
}
```
