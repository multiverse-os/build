<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: Go Build System 'build' *(not yet named)*
**URL** [multiverse-os.org](https://multiverse-os.org)

## Multiverse Go Compiler

## Dependency Management
Along with supporting existing dependency management, a very simple
`.build.deps` file will be supported. Design is heavily inspired by Ruby's
`RubyGems` but will support YAML, TOML and JSON data types; using whatever is
detected. It will define Go version being used by development environment,
everisons of dependent packages, repository (allowing specification of branch or
commit easily). 

### Design
The build tool will replace the [default Go command-line
tool](https://github.com/golang/go/tree/master/src/cmd/go) and extend its
functionality, add depdency management (design will be heavily inspired by
RubyGems) 

**Development Console**
A development console will be included to support testing, experimentation, and
to support teaching.

**Binary API**
Low-level API for defining compilation binary. 

### Useful Resources
The Go source will be the most important resource, we will include all
functionality found in `src/cmd/*` to ensure complete backwards compatability
and seamless upgrading to the Multiverse OS drop-in alterative build system. 

  * [buildid](https://github.com/golang/go/tree/master/src/cmd/internal/buildid)
    Can be customized to put our own data in the executable

  * [xcoff](https://github.com/golang/go/blob/master/src/cmd/internal/xcoff)
    // Package xcoff implements access to XCOFF (Extended Common Object File Format) files
