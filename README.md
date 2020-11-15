<h1 align="center">iCanvas Analytics</h1>

<p align="center">
  <a href="https://travis-ci.com/abmid/icanvas-analytics.svg?branch=master"><img src="https://travis-ci.com/abmid/icanvas-analytics.svg?branch=master"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License"></a>
</p>
<p align="center"><img src="https://i.ibb.co/7VCrCtj/screencapture-localhost-8000-2020-11-14-10-40-09-copy.png"></p>
Web Apps for reporting Canvas LMS

- This project under development and part of UMM (University of Muhammadiyah Malang)

<h2>Todos</h2>

<h2>Installation / Use</h2>

iCanvas Analytics require database PostgreSQL to run.

### Docker

<ol>
  <li>Copy .env.example paste and give name .env</li>
  <li>Change information about database PostgreSQL (PG_HOST, PG_PORT, PG_DBNAME, PG_USER, PG_PASSWORD)</li>
  <li>Change <b>SECRET_KEY</b></li>
  <li>Open terminal change directory to this project and write command <code>make docker</code></li>
</ol>

After container successfull running, you can access default in port <code>:8000</code> / <code>http://localhost:8000</code>

To stop continer, use this command

```sh
make docker-stop
```

### Manual
<ol>
  <li>You must install Go lang first, check about installation in <a href="https://golang.org/doc/install" target="_blank">here</a>.</li>
  <li>Copy .env.example paste and give name .env</li>
  <li>Change information about database PostgreSQL (PG_HOST, PG_PORT, PG_DBNAME, PG_USER, PG_PASSWORD)</li>
  <li>Change <b>SECRET_KEY</b></li>
  <li>Open terminal change directory to this project and write command <code>make app</code></li>  
</ol>

After successfull generate, you can access default in port <code>:8000</code> / <code>http://localhost:8000</code>

License
----

MIT
