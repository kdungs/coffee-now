# CoffeeNow!

Experimenting with an app that allows you to have a coffee with people close by on short notice.

Using [PostgreSQL](http://www.postgresql.org/) for the database (especially for its ability to do [radius queries](http://datachomp.com/archives/radius-queries-in-postgres/)) and Websockets for communication between the server and the front-end.

At this stage there are two branches with separate experiments:
 * [experiments/postgres](https://github.com/kdungs/coffee-now-server/tree/experiments/postgres): Explore PostgreSQL bindings for Go and play around with radius queries.
 * [experiments/websocket](https://github.com/kdungs/coffee-now-server/tree/experiments/websocket): Have a look at websockets in Go with a demo HTML app.

_This repository is work in progress._
