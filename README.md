# LegacyLab

The LegacyLab project provides a set of tools for advanced technical debt
analysis and management.

The tools are build and tested on Mac OS.

Supported Languages: Java, Go

## Build

Builds and installs the LegacyLab binaries into the _bin_ directory of your
GOPATH.

```
make
```

## List of Tools

### Hotspot Analysis: hs

The hotspot analysis tool _hs_ analyses a git repository and reports a list of
files ordered by their commit frequency and complexity. The tool could be run
either on a local or remote git repository depending on the URL given as
command-line parameter:

```
# run hs on local repository
hs -url file:///Users/ralf/tmp/spring-data-jpa

# run hs on remote repository
hs -url https://github.com/spring-projects/spring-data-jpa.git
```

Remote repository URLs are cloned or pulled prior the analysis into _tmp_ of the
users home directory.

#### Usage

```
hs --help
```