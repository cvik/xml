# An alternative XML parser
This package contains a conceptually simple in-memory xml-parser. It creates
a tree of nodes with parent-links on each level for easy down- and up-traversal.

### Properties
* Doesn't use any reflection
* Traversal friendly (double linked element tree)
* Keeps track of namespaces at every element level

### TODO
* Handle CDATA nodes
* Better error messages
* Better test-coverage (currently at 80.1%)
* Look into value coercion (maybe)
* Add documentation (current at 0%)
* Add usage examples
* Look into memory usage (currently a bit high)
* Add SAX functionality


## Testing
 - `go test -coverprofile coverage.out`
 - `go tool cover -html=coverage.out`
