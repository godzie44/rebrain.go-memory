go build cgo.go

valgrind --error-limit=no --leak-check=full --leak-resolution=med ./cgo 2> vg_out