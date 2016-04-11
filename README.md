# Harvey CMS

[Harvey](http://www.harveycms.org) is an Open Source CMS powered by [Go lang](https://golang.org) and [Ember.js](http://emberjs.com).

__Please Note:__ This project is currently pre-alpha stage and should not be used for production environments.
Please [join our mailing list](http://www.harveycms.org) if you're interested in keeping up-to-date with the latest project news.

This project aims to implement a simple, elegant and easy-to-use CMS which is secure and highly scalable.
For more information, check out our 'docs' subproject.

## Subprojects

Each folder in this repository represents a sub-project:

* __api__
    [Go lang](https://golang.org) backend for the CMS app, using the [JSON API](http://jsonapi.org/) spec.

* __cli__
    [Go lang](https://golang.org) command-line interface app for running tasks.

* __cms__
    [Ember.js](http://emberjs.com) control panel app for managing your content.

* __core__
    [Go lang](https://golang.org) core code library shared between api and web apps.

* __demo__
    Example Harvey CMS site directory containing demo config.

* __docs__
    Documentation website built with [Hugo](https://gohugo.io/) static site engine.

* __modules__
    Code modules shared between api and web apps.

* __web__
    [Go lang](https://golang.org) website frontend app, supporting custom themes.

Eventually these subprojects will be given their own repositories.

## License

Harvey CMS is distributed under the MIT License.
Copyright &copy; 2016 [Ryan D](http://ryan0x44.com). All Rights Reserved.
