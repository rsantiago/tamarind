---
title: Go Embed Magic
subtitle: Single binary deployment explained
date: 2025-11-22
tags: [golang, technical]
---

# The Power of `embed`

One of Tamarind's killer features is that it is a single binary. You don't zip up a folder of templates along with the executable. You just ship the executable.

This is made possible by the `embed` package introduced in Go 1.16.

## How it works

In `main.go`, we declare a variable to hold our assets:

```go
//go:embed assets/*
var embeddedAssets embed.FS
```

The compiler sees the directive and literally stuffs the contents of the `assets/` directory into the binary segments of the compiled code.

## Virtual Filesystem

At runtime, `embeddedAssets` acts like a read-only filesystem. We can traverse it, read files, and copy them.

When you run `tamarind init`, we execute a function that walks this virtual filesystem:

```go
fs.WalkDir(embeddedAssets, "assets/structure", func(path string, d fs.DirEntry, err error) error {
    // Logic to copy file from memory to disk
})
```

This allows us to carry a "skeleton key" inside the program—a perfect copy of the default project structure, ready to be replicated anywhere.
