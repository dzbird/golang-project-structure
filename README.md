## Usage

### Adding your dependencies

Edit the file `installdeps` and add `go get ...` statements. 
Example: <a href="https://github.com/dzbird/golang-project-structure/blob/master/installdeps#L7">Line 7</a>


### Install dependencies with:

```
./installdeps
```

The dependencies are placed in `dependencies` directory.

This also adds a `.vscode` workspace setting with current `GOPATH`. This allows lint and formatting tools to work correctly. 
Without it, you might have to use other workarounds.


### Run with

```
./run
```

If you still want to do `go run ....`, make sure you do a `source ./setgopath` whenever the project terminal is opened.


### Build

```
./build
```

This will create a binary in the `dist` directory.
