# Building and Testing Stargate

This document describes how to set up your development environment to build and test Stargate.
It also explains the basic mechanics of using `git`, `node`, and `npm`.

* [Prerequisite Software](#prerequisite-software)
* [Getting the Sources](#getting-the-sources)
* [Installing NPM Modules](#installing-npm-modules)
* [Building](#building)
* [Running Tests Locally](#running-tests-locally)

See the [contribution guidelines](https://github.com/sfeir-open-source/stargate/blob/master/CONTRIBUTING.md)
if you'd like to contribute to Stargate.

## Prerequisite Software

Before you can build and test Stargate, you must install and configure the
following products on your development machine:

* [Git](http://git-scm.com) and/or the **GitHub app** (for [Mac](http://mac.github.com) or
  [Windows](http://windows.github.com)); [GitHub's Guide to Installing
  Git](https://help.github.com/articles/set-up-git) is a good source of information.

* [Node.js](http://nodejs.org), (version specified in the engines field of [`.nvmrc`](../.nvmrc)) which is used to run a development web server,
  run tests, and generate distributable files.

* [Java Development Kit](http://www.oracle.com/technetwork/es/java/javase/downloads/index.html) which is used
  to execute the selenium standalone server for e2e testing.

* [Bazel](https://bazel.build/), please follow instructions in [BAZEL.md](./BAZEL/)

## Getting the Sources

Fork and clone the Stargate repository:

1. Login to your GitHub account or create one by following the instructions given
   [here](https://github.com/signup/free).
2. [Fork](http://help.github.com/forking) the [main Stargate repository](https://github.com/sfeir-open-source/stargate).
3. Clone your fork of the Stargate repository and define an `upstream` remote pointing back to the Stargate repository that you forked in the first place.

```shell
# Clone your GitHub repository:
git clone git@github.com:<github username>/stargate.git

# Go to the Stargate directory:
cd stargate

# Add the main Stargate repository as an upstream remote to your repository:
git remote add upstream https://github.com/sfeir-open-source/stargate.git
```
## Installing project dependencies

@TODO

## Building

@TODO

## Running Tests Locally

@TODO

## <a name="clang-format"></a> Formatting your source code

Stargate uses [clang-format](http://clang.llvm.org/docs/ClangFormat.html) to format the source code. If the source code
is not properly formatted, the CI will fail and the PR can not be merged.

You can automatically format your code by running:

@TODO

## Linting/verifying your source code

You can check that your code is properly formatted and adheres to our coding style by running:

@TODO

## Publishing snapshot builds

@TODO