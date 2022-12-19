# IBAN validator

## Requirements

* Docker
* Golang
* Make

## Usage

Start the service and send a GET request to `/validate` with a query param called `iban` holding the IBAN to validate. The IBAN must be in Standard format (no spaces).

Example:

    curl 'http://localhost:5000/validate?iban=SE4550000000058398257466'

* If the IBAN is valid, the response will have code *200*;
* if it is invalid, or if the country is not supported, it will have code *400*, together with an explanation.
* The response format is text.
* The port is hardcoded to `5000`.
* There is a healthcheck endpoint at `/health`.

## To build

    make build

## To run

    make run

## To test

    make test

Please note that the tests will run on the local machine, because the Docker image I chose seems to not support running tests(?).

## Notes

The validator does the general [mod 97 check](https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN).

The country specific checks are limited to checking for the country specific lengths. Two countries are supported: Albania and Sweden.
