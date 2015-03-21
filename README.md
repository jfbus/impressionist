impressionist
=============

[![wercker status](https://app.wercker.com/status/a7d76b3660d02a42306368445b74fcdf/s/master "wercker status")](https://app.wercker.com/project/bykey/a7d76b3660d02a42306368445b74fcdf)

A flexible image server written in go, using chained image filters.

Two backend are available : a pure Go backend, and a ImageMagick/GraphicsMagick backend.

Status: alpha

Building the server (pure Go):
```
go build -o impressionist server/pure/server.go
```

Building the server (ImageMagick):
```
go build -o impressionist server/magick/server.go
```

On ubuntu, you will need to install `libmagickcore-dev` and `libgif-dev`

Building the server (GraphicsMagick):
```
go build -tags gm -o impressionist server/magick/server.go
```

Running the server :
```
impressionist -cfg impressionist.json [-debug]
```

Configuration
-------------

Example configuration file :

```json
{
	"http": {
		"port": 80,
		"root": "/impressionist",
		"timeout": "30s",
		"workers": 10
	},
	"filters": [
		{
			"name": "small",
			"definition": "s,100x0"
		}
	],
	"image": {
		"quality": 90
	},
	"storages": [
		{
			"name": "local",
			"type": "local",
			"path": "/tmp"
		}
	],
	"caches": {
		"source": 100
	}
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
* resize : `s,[w]x[h][+|-]`
* flip : `f,[h|v]` h = horizontal, v = vertical 
* rotate: `r,[90|180|270]`
* grayscale: `gs`
* blur: `b,s` s = float (e.g: 1.0)

Note on resize :

* `s,500x100` resizes exactly to 500x100
* `s,500x0` and `s,0x100` resize to a width of 500px (height of 100px), keeping the ratio
* `s,500x100+` resizes to the maximum size within the specified dimensions, keeping the ratio
* `s,500x100-` resizes to the minimum dimention above the specified dimensions, keeping the ratio, and crops to 500x100

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

Filters supported by each backend :

|filter|pure Go|magick|
|------|-------|------|
|crop|yes|yes|
|resize|yes|yes|
|flip|yes|yes|
|rotate|yes|no|
|grayscale|yes|yes|
|blur|yes|no|


Formats
-------

* PNG: `png`,
* JPEG: `jpeg[,quality]` - if not specified, quality is set to the default value (`image.quality` config variable, or 75 if not configured).

Caching
-------

A in-memory cache is available. It stores source files to avoid accessing filesystem for each access.

By default, it caches 100 entries, and cache size can be changed by setting `caches.source`. A -1 value disables the cache.

Workers
-------

Image manipulation is delegated to a pool of workers. If a request is not finished within the configured timeout (`http.timeout`, by default, 30s), the request returns a 503 error.

TODO
----

This is a first draft, everything is still pretty rough

* Tests, tests & tests
* Clean stuff
* Tuning
* Add more storage types (eg: AWS, http)
* Support environment variables for configuration tokens
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