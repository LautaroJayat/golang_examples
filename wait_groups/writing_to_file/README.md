# Writing to a file

Here we are using go rutines and wait groups to concurrently write to a single file.

In one example we just make a non atomic writing.

On the other we use the logger type to make sure that only one go rutine has access to the writer at a given moment see https://golang.org/pkg/log/#Logger