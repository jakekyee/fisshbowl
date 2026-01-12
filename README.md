# fiSSHbowl
Watch SSH probing attempts on your home network




See a live geographic maps of attempts on my network at jakeyee.com\maps
Note that the link may not work because I'm doing a server move right now.

For now you can see the results at 

https://api.jakeyee.com/attempts?from=2025-01-01T00:00:00Z&to=2027-01-10T23:59:59Z




There are 3 parts to this project:
- The API
- The database
- The logger

To run:


Docker
Go
Python

Start the database

cd dbstuff
docker compose -


Build and start the logger
cd fisshbowl
docker build -t my-app:latest .
docker run -d -p 22:22 my-app:latest

Start the api
cd apistuff
go run .



## Getting Started

### Clone the Repository

```sh
git clone https://github.com/jakekyee/fisshbowl.git
cd fisshbowl

```




