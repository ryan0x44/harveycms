---
date: 2016-03-09T00:11:02+01:00
title: Getting started
---

{{< warning title="This documentation may be wrong" >}}
Some documentation may be written before actual functionality is implemented!
<br />As stated on the index of this site, this is pre-alpha software.
{{< /warning >}}

## Installation

### Before you start

We have a few dependencies:

- [Go lang 1.6](https://golang.org/doc/install)

- [Node.js 5.10 & NPM](https://nodejs.org/)

- [Docker Toolbox 1.10](https://docs.docker.com/toolbox/overview/)

### Installing Harvey CLI

The CLI tool will allow you to scaffold out and manage a new project.

```
go install github.com/ryan0x44/harveycms/cli/cmd/harvey
```

## Setup

### Creating a new Harvey project

Use the `new` command from the CLI tool to create a new project directory:

```
harvey new my-project
```

This should create a directory structure like this:

```
.
|-- my-project/
	|-- config.toml
```
