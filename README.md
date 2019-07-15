# PicCollage take home quiz - Building Shorten URL service

PicCollage shorten url service written in Go and Node.js.

## This project contains three components:
* Shorten url service (written in Go)
  * Shorten url service
  * Counter dispatcher
* Migration scripts
* Test scripts (written in Node.js with mocha framework)

### Build shorten url services and test scripts docker images

1. cd to project's root directory

2. Build docker images:

`./docker/build.sh`

3. Start shorten url services + run migrations

`./docker/run_shorten_url.sh`

4. Run test scripts

`./docker/run_test_cases.sh`
