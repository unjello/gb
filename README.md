# gb: Great Builder, The

Is a yet another build generator for C++

## Why?

Convention-over-configuration seem to work well for Ruby and Maven. I wanted to have similar smooth experience with C++, without a need to write (or even worse, copy around) tons of CMake boilerplate code. I've decided to settle on structure similar to [pitchfork](https://github.com/vector-of-bool/pitchfork). Similar enough, that I've decided to adopt it fully, so maybe get out of `Status: DREAM`, and become a reality 😊

## Project Layout

### Header-only Library

If a project has `include` folder, but no `src` folder, it is considered to be _header-only library_.

### Library

If a project has both `include` and `src` folders, it is considered to be a _library_ - for now. All files under `src` folder (recursively) will be compiled in.

### Application

If a project has `src` folder, but no `include` folder, it is considered to be an application, and will compile to an executable. All files under `src` folder (recursively) will be compiled in.

### Unit-test

If `test` folder is present, each `.cpp` file inside is compiled independently. If project is a library, each test will automatically be linked with it, and will have library's public headers in include path. Otherwise the test is expected to include file under test (be it header or source file) manually.

The framework of choice for Unit Testing is [doctest](https://github.com/onqtam/doctest).

## License

- CCO (~Public Domain)
