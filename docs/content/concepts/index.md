---
date: 2016-04-11T19:12:46+01:00
title: Harvey CMS Concepts
---

## Architecture

Harvey consists of four apps:

* __CMS__
    [Ember.js](http://emberjs.com) control panel app for managing your content.
* __API__
    [Go lang](https://golang.org) backend for the CMS app, using the [JSON API](http://jsonapi.org/) spec.
* __WEB__
    [Go lang](https://golang.org) website frontend app, supporting custom themes.
* __CLI__
    [Go lang](https://golang.org) command-line interface app for running tasks.

### Modules

Harvey functionality is composed of modules, such as Pages, Blog, Forms, etc.
Modules can have several models, such as a Blog Post.

### Routing

Harvey has a single routes table. Any model can have an associated route.
This allows for complete flexibility in routing, while keeping the implementation simple and highly performant.

### HTTP Caching

Each model is responsible for implementing an interface which requires calculation of an Etag.
The HTTP Caching validation model allows websites to scale easily using tools such as [Varnish](https://www.varnish-cache.org).

### Themes

Harvey supports custom themes using the excellent [Go html/template](https://golang.org/pkg/html/template) library.

### Media

Assets such as Blog Post images are stored in media containers backed by services such as Amazon S3.

## Terminology

Key terminology used within Harvey CMS:

- **Route** a route is a fixed path to an entity

- **Rewrite** a rewrite is a dynamic path to an entity
