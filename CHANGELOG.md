# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [1.0.4](https://github.com/vlazic/gpt-project-context/compare/v1.0.1...v1.0.4) (2023-05-17)


### ⚠ BREAKING CHANGES

* The GetFilePaths function now requires an additional parameter. Existing calls to this function will need to be updated, either by providing a specific root directory or by passing an empty string to use the current working directory.

### Features

* Create separate .go files for flags, file operations, output creation, and clipboard operations ([6106436](https://github.com/vlazic/gpt-project-context/commit/6106436abc6cd4f4eac9e20b18cc4fa590acd139))
* Enhance GetFilePaths and FilterFiles functions for testing ([89ea603](https://github.com/vlazic/gpt-project-context/commit/89ea603ce83835ade9f79bd2fdf5fa1b845935d2))


### Bug Fixes

* Update Makefile to build all binaries before creating GitHub release ([0b84b06](https://github.com/vlazic/gpt-project-context/commit/0b84b0662979614a9a4673a9d22426992c138507))

### [1.0.3](https://github.com/vlazic/gpt-project-context/compare/v1.0.1...v1.0.3) (2023-05-17)


### Features

* Create separate .go files for flags, file operations, output creation, and clipboard operations ([6106436](https://github.com/vlazic/gpt-project-context/commit/6106436abc6cd4f4eac9e20b18cc4fa590acd139))


### Bug Fixes

* Update Makefile to build all binaries before creating GitHub release ([0b84b06](https://github.com/vlazic/gpt-project-context/commit/0b84b0662979614a9a4673a9d22426992c138507))

### [1.0.2](https://github.com/vlazic/gpt-project-context/compare/v1.0.1...v1.0.2) (2023-04-28)


### Bug Fixes

* Update Makefile to build all binaries before creating GitHub release ([0b84b06](https://github.com/vlazic/gpt-project-context/commit/0b84b0662979614a9a4673a9d22426992c138507))

### [1.0.1](https://github.com/vlazic/gpt-project-context/compare/v1.0.0...v1.0.1) (2023-04-28)


### Features

* Update Makefile to build for linux, update README.md with installation instructions ([92564f8](https://github.com/vlazic/gpt-project-context/commit/92564f841696dc6027ccccb60185ae789fc7c60e))

## 1.0.0 (2023-04-28)


### Features

* First commit ([f95d6a5](https://github.com/vlazic/gpt-project-context/commit/f95d6a5fc783b94aacff8b9dbb864322d163013e))
