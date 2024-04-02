### Solar System Article Viewer GUI Version
This little Go program provides a user with a handful of fictional
articles. There is a small story apart in the game.
The web backend is built with Gin.
I had a lot of fun with the articles, there might be some grammar mistakes primarly
because spacing once its read from a file at startup.

### Installation
You will need Go, podman/docker to run this program

close the repository into a project directory
```shell script
git clone https://github.com/cwdatlas/SolarSystemNews-Go-Web
```
Make sure your respective container engine is running docker/podman
Then start the database container
```shell script
podman run --detach --name SpaceWeb --env POSTGRES_PASSWORD=12345 -p 5432:5432 docker.io/library/postgres
```

Move into the newly created directory
```shell script
cd SolarSystemNews-Go-Web
```

Run the program by typing
```shell script
go run .
```

Go have fun looking at articles in the program! I hope you like my stories.