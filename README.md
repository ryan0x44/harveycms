# Harvey CMS

[Harvey](http://www.harveycms.org) is an Open Source CMS powered by [Go lang](https://golang.org) and [Ember.js](http://emberjs.com).

__Please Note:__ This project is currently pre-alpha stage and should not be used for production environments.
Please [join our mailing list](http://www.harveycms.org) if you're interested in keeping up-to-date with the latest project news.

This project aims to implement a simple, elegant and easy-to-use CMS which is secure and highly scalable. 
For more information, check out the [docs](docs/) subproject.

## Subprojects

Each folder in this repository represents a sub-project:

* __[api](api/)__
    Go lang backend for the CMS app, using the [JSON API](http://jsonapi.org/) spec.

* __[cli](cli/)__
    Go lang command-line interface app for running tasks.

* __[cms](cms/)__
    Ember.js control panel app for managing your content.

* __[core](core/)__
    Go lang core code library shared between api and web apps.

* __[demo](demo/)__
    Example Harvey CMS site directory containing demo config.

* __[docs](docs/)__
    Documentation website built with [Hugo](https://gohugo.io/) static site engine.

* __[modules](modules/)__
    Code modules shared between api and web apps.

* __[web](web/)__
    Go lang website frontend app, supporting custom themes.

Eventually these subprojects will be given their own repositories.

## License

Harvey CMS is distributed under the MIT License.
Copyright &copy; 2016 [Ryan D](http://ryan0x44.com). All Rights Reserved.
