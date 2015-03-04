impressionist
=============

A flexible image server written in go, using chained image filters.

Status: dev, things might be broken or will be broken

Configuration
-------------

Example configuration file :

```json
{
	"http": {
		"port": 80,
		"root": "/impressionist"
	},
	"filters": [
		{
			"name": "small",
			"definition": "s,100x0"
		}
	],
	"jpeg": {
		"quality": 90
	},
	"storages": [
		{
			"name": "local",
			"type": "local",
			"path": "/tmp"
		}
	]
}
```

Storage
-------

One type of storage is currently supported : `local`.

URL structure
-------------

`/impressionist/:filter/:format/:storage/:path`

* `:filter` is a comma separated list of filters, as explained below,
* `:format` see below,
* `:storage` is the name of the storage, as defined in the configuration,
* `:path` is the path of the image, relative to the storage root.

Filters
-------

* crop : `c,[x]x[y]-[w]x[h]`
* resize : `s,[w]x[h]` if w or h is zero, image is resized to the other dimension, keeping the same ratio
* flip : `f,[h|v]` horizontal or vertical flip
* rotate: `r,[90|180|270]` rotate
* grayscale: `gs`

Predefined filters can be configured.

With the following configuration :

```json
{
	"filters": [
		{
			"name": "small",
			"definition": "s,100x0"
		}
	]
}
```

`/impressionist/small/:format/:storage/:path` is translated to `/impressionist/s,100x0/:format/:storage/:path`.

Formats
-------

* PNG: `png`,
* JPEG: `jpeg[,quality]` - if not specified, quality is set to the default value (`jpeg.quality` config variable, or 80 if not configured).

TODO
----

This is a first draft, everything is still pretty rough

* Tests, tests & tests
* Clean stuff
* Tuning
* Add context & nice logs
* Move image manipulation to separate goroutines
* Add more storage types (eg: AWS)
* Support environment variables for tokens
* Add more filters
* Cache resized images
* Secure URL using hashes
* Token support
* Optimize images using pngcrush/JPEGOptim/...
* Format negociation & webp
* Graceful restart & configuration reloading
* Metrics

LICENSE
-------

MIT, see LICENSE