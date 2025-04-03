/*
This is a client for the Harper application platform.
It mirrors the HTTP API as of version 4.x and makes it
very easy to get up and running with Harper and your Go application.

For more information see: https://docs.harperdb.io/

# Basics

Instantiate a new client:

	client := harper.NewClient("http://localhost:9925", "username", "password")
*/
package harper
