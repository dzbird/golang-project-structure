## Usage

### Install dependencies with:

```
./installdeps
```

The dependencies are placed in `dependencies` directory.

This also adds a `.vscode` workspacesetting with current `GOPATH`. This allows lint and formatting tools to work correctly. 
Without it, you might have to use other workarounds/automations.


### Run with

```
./run
```

If you still want to do `go run ....`, make sure you do a `source ./setgopath` whenever the project terminal is opened.


### Build

```
./build
```

This will create the binary to `dist` directory.
