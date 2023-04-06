
# WasaPhoto

This is an example of web application for [Web and Software Application](http://gamificationlab.uniroma1.it/en/wasa/)
course.

## Project structure

* `cmd/` contains all executables
	* `cmd/healthcheck` daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains web API server daemon
* `demo/` contains config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains the API server (Go code)
	* `service/database` contains the code for interacting with the database (Go code interacting with sqlite3 databases)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` the web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding
	* Backend and frontend containers

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development


## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```
### My build works when I use `npm run dev`, however there is a Javascript crash in production/grading

Some errors in the code are somehow not shown in `vite` development mode. To preview the code that will be used in production/grading settings, use the following commands:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-prod
npm run preview
```



## License

See [LICENSE](LICENSE).
