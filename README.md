# e-cheve
This is a program to index and an application to search for emails.

For the backend Golang is the main language used and Chi is for the routing.

To index all the data I use Zincsearch.

In the frontend I use TypeScript with VueJS and Tailwind.

Also, I implemented benchmarks to see the program performance and improve it.

Soon, this project will be on my portfolio for you to see it.

## Initialize project

### Backend

To initialize the backend api you need to execute `go run main.go`.

### Frontend

You need to setup the VueJs-Vite project executing `npm install`.
To initialize the frontend you need to execute `npm run dev`.

### Indexer

To initialize the indexer you need to execute `go run main.go`.
You need to pass an argument to index (a tar file). If you don't add this the program will not run.

### Profiling and Benchmark

Execute this command to generate the block.prof, cpu.prof and mem.prof files:
`go test -blockprofile block.prof -cpuprofile cpu.prof -memprofile mem.prof -bench=. -benchmem`.

Then you can execute this command to see in the browser the generated graph for the block prof:

`~/go/bin/pprof -http localhost:3333 ../e-cheve-indexer block.prof`

You can change the `block.prof` argument with the `mem.prof` or the `cpu.prof` files.

The port used in this example is just to show how to use it. You can use another port.
