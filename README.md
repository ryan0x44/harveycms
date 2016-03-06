# Harvey CMS

[Harvey](http://www.harveycms.org) is an Open Source CMS powered by [Go lang](https://golang.org) and [Ember.js](http://emberjs.com).

__Please Note:__ This project is currently pre-alpha stage and should not be used for production environments.
Please [join our mailing list](http://www.harveycms.org) if you're interested in keeping up-to-date with the latest project news.

This project aims to implement a simple, elegant and easy-to-use CMS which is secure and highly scalable.

## Architecture

Harvey consists of three apps:

* __CMS__
    An [Ember.js](http://emberjs.com) control panel app for managing your content.
* __API__
    A backend for the CMS app, using the [JSON API](http://jsonapi.org/) spec and [Go lang](https://golang.org).
* __WEB__
    A website frontend app, using [Go lang](https://golang.org) and supporting custom themes.

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

## License

The Harvey CMS is distributed under the MIT License.
Copyright &copy; 2016 [Ryan D](http://ryan0x44.com). All Rights Reserved.
