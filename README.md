<img src="docs/simple-iot-logo.png?raw=true" width="150">

Simple IoT is collection of building blocks and best practices
for building IoT applications, learned from experience building
real-world systems.

Demo is running at: https://portal.simpleiot.org/

There is not much here yet -- mostly just a build/deployment system at
this point.

[Detailed Documentation](docs/README.md)

## Example

In the following example, we are running the server process and a device
simulator on a local machine. The device simulator sensors simulated sensor
data to the server.

- make sure Go v1.11 or newer is installed
- `go get github.com/simpleiot/simpleiot/cmd/siot`
- start server: `~/go/bin/siot`
- start device simulator: `~/go/bin/siot -sim`
- open http://localhost:8080

(adjust above paths if you have your `GOPATH` somewhere else.)

## Vision

- built around collecting and visualizing data from devices
- provide a good base system to build IoT products that may support a number of devices, users, rules, etc.
- requires coding to customize. This is not a GUI for building IoT systems,
  but rather a code base software developers can use as a starting point.
- application technology is general, so you are not dependant on any one
  IoT company or cloud provider
- plugin architecture for extending the system with custom functionality
- simple deployment process (not a lot of pieces on the backend to manage)
  - Deployment/upgrade is as simple as copying one binary.
  - All assets are embedded.
  - For small deployments (< 1000 devices), application is completely self contained
    (no external databases or services are required).
- Storage (config and sensor data) supports multiple databases
  - embedded db for small deployments
  - (mongodb, Google Cloud Datastore, influxdb, etc) for larger deployments or other
    needs.

## Short term features

- edit/save device config
- esp32 client example
- user accounts
- store timeseries data
- graph timeseries data
- rules engine (conditions/consequences)

done:

- device management
- simple dashboard for each device showing collected parameters
- REST api for devices
- Embedded database using boldhold

## Long term features

- efficient protocols for cellular data connections (CoAP, etc.)
- Google Cloud Datastore
- App Engine Deployment
- edge computing features
- organization support

## Technology choices

Choices for the technology stack emphasize simplicity, not only in the
language, but just as important, in the deployment and tooling.

- **Backend**
  - [Go](https://golang.org/)
    - simple language and deployment model
    - nice balance of safety + productivity
    - excellent tooling and build system
- **Frontend**
  - Single Page Application (SPA) architecture
    - programming environment is much more powerful than server rendered
      pages (PHP, Rails, etc).
    - easier to transition to Progressive Web Apps (PWA)
  - [Elm](https://elm-lang.org/)
    - nice balance of safety + productivity
    - excellent compiler messages
    - reduces possibility for run time exceptions in browser
    - does not require a huge/fragile build system typical in
      Javascript frontends.
  - [Bootstrap](http://getbootstrap.com/)
    - mature CSS toolkit that handles browser differences and
      responsive design for mobile reasonably well.
    - widespread adoption and well understood by many developers
    - well supported [bindings in Elm](https://package.elm-lang.org/packages/rundis/elm-bootstrap/latest/)
- **Database**
  - Eventually support multiple databased backends depending on scaling/admin needs
  - Embedded db using [BoltHold](https://github.com/timshannon/bolthold)
    - no external services to configure/admin
- **Hosting**
  - Any server (Digital Ocean, Linode, etc)
  - [Google App Engine](https://cloud.google.com/appengine/)
    - is simple to deploy Go applications
    - handle high-availability, scaling, etc.
  - (any server/hosting environment that supports Go apps can be used)

In our experience, simplicity and good tooling matter. It is easy to add features
to a language, but creating a useful language/tooling that is simple is hard.
Since we are using Elm on the frontend, it might seem appropriate to select
a functional language like Elixir, Scala, Clojure, etc. for the backend. These
environments are likely excellent for many projects, but are also considerably more
complex to work in. The programming style (procedural, functional, etc.) is important,
but other factors such as simplicity/tooling/deployment are also important, especially
for small teams who don't have separate staff for backend/frontend/operations. Learning two
simple languages (Go and Elm) is a small task compared to dealing with huge
languages, fussy build tools, and complex deployment environments.

This is just a snapshot in time -- there will likely be other better technology choices in the
future. The backend and frontend are independent. If either needs
to be swapped out for a better technology in the future, that is possible.
