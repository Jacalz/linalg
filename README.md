# linalg
Easy-to-use abstractions around Linear Algebra calculations. Written in pure Go with no external dependencies.
The library provides three packages for working with vectors and points as well as one for working with matricies.

***NOTE:*** *This package is work in progress and might see large changes going forward. Use at your own risk.*

## Packages

### r2
An optimized abstraction layer for working with vectors and points on a plane, in two diemensions.

### r3
An optimized abstraction layer for working with vectors and points in a room, in three diemensions.

### rn
A generic abstraction layer for working with vectors and points in a space of dimension n.

***NOTE:*** *Only use this package for dimentions of 4 or higher. The r2 and r3 packages are better optimized and easier to use for their specific use cases.*

### matrix
Provides a generic abstraction layer around matrix calculations in a space of dimension n.

## Requirements
A [Go](https://golang.org/) compiler of a recent version should do. Official support will be for the two latest Go versions but older versions will most likely work without issues.

## Contributing
Contributions are strongly appreciated. Everything from creating bug reports to contributing code will help the project a lot, so please feel free to help in any way, shape or form that you feel comfortable with.

## License

Linalg is licensed under the `BSD 3-Clause License` and freely avaliable to all of those that wish to use it.
