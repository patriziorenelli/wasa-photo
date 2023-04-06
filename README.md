
# WasaPhoto

This is an example of web application for [Web and Software Application](http://gamificationlab.uniroma1.it/en/wasa/)
course.

## Project structure
•	cmd/ contains all executables <br />
&emsp    o	cmd/healthcheck a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness
    probes (e.g., Docker engine) <br />
&emsp    o	cmd/webapi contains web API server daemon <br />
•	demo/ contains a demo config file <br />
•	doc/ contains the documentation (usually, for APIs, this means an OpenAPI file) <br />
•	service/ has all packages for implementing project-specific functionalities <br />
&emsp    o	service/api contains API server <br />
 &emsp   o	service/database contains code to manage the database <br />
•	vendor/ is managed by Go, and contains a copy of all dependencies <br />
•	webui/  frontend in Vue.js; it includes: <br />
&emsp    o	Bootstrap JavaScript framework <br />
&emsp    o	a customized version of "Bootstrap dashboard" template <br />
&emsp    o	feather icons as SVG <br />
&emsp    o	Go code for release embedding <br />




## License

See [LICENSE](LICENSE).
