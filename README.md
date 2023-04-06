
# WasaPhoto

This is an example of web application for [Web and Software Application](http://gamificationlab.uniroma1.it/en/wasa/)
course.

## Project structure
•	cmd/ contains all executables
    o	cmd/healthcheck a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
    o	cmd/webapi contains web API server daemon
•	demo/ contains a demo config file
•	doc/ contains the documentation (usually, for APIs, this means an OpenAPI file)
•	service/ has all packages for implementing project-specific functionalities
    o	service/api contains API server
    o	service/database contains code to manage the database
•	vendor/ is managed by Go, and contains a copy of all dependencies
•	webui/  frontend in Vue.js; it includes:
    o	Bootstrap JavaScript framework
    o	a customized version of "Bootstrap dashboard" template
    o	feather icons as SVG
    o	Go code for release embedding




## License

See [LICENSE](LICENSE).
