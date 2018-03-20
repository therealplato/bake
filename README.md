#### github.com/therealplato/bake

bake a file into a golang byte slice

like go-bindata but lighter weight

#### usage:

```
go install -v -u github.com/therealplato/bake
cd someproject
bake somefile
$EDITOR baked.go # change filename, package and variable name as needed
```

#### one way to manipulate the file:

```
var src bytes.Buffer
_, err := src.Write(baked)
if err != nil {
  log.Fatal(err)
}
// src is io.Reader
```
