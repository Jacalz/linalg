# linalg

Fast and easy-to-use packages for Linear Algebra calculations. Written in pure Go with no external dependencies.
The goal is to provide vector and matrix implementations that are cross-platform, fast and easy-to-use. 

***NOTE:*** *This package is work in progress and might see large changes going forward. Use at your own risk.*

## Packages

Below is a list of the avaliable packages.

- **github.com/Jacalz/linalg/matrix**: Provides a package for working with matricies in any dimension.
  - **github.com/Jacalz/linalg/matrix/constants**: Predefined matricies for specific use cases.
- **github.com/Jacalz/linalg/vector**: Provides a package for working with vectors in any dimension.
  - **github.com/Jacalz/linalg/vector/r2**: Optimized package for working with vectors in two dimensions.
  - **github.com/Jacalz/linalg/vector/r3**: Optimized package for working with vectors in three dimensions.

## Requirements

A [Go](https://golang.org/) compiler of a recent version should do. Official support will be for the two latest stable Go versions, but older versions might still work without issues. For the currently recommended minimum version, see [go.mod](./go.mod).

## Contributing

Contributions are strongly appreciated. Everything from creating bug reports to contributing code will help the project a lot, so please feel free to help in any way, shape or form that you feel comfortable with.

## License

Linalg is licensed under the `BSD 3-Clause License` and freely avaliable to all of those that wish to use it.
