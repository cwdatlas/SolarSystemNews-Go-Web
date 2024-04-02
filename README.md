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

It will ask for permissions if you are on windows.
### Web
Now you can connect at localhost:8080/
There you can create, Update, or search for an article!
The search looks through all article's title, author/date, and location, to find your keyword.
Be careful when making or updating an article, if there are any errors, blank entries or too many characters for an entry,
then your page will be refreshed, and you will be told the error. This means your work is erased. This is a major bug, save your work elsewhere before you write.
### update
If you want to update an article, then click the checkbox at the bottom of the page. 
To update an article you will need the exact title, so copy it from the article you want to edit.
Fill in the title entry with the title you want to edit then you can write your change to any other entry and it will replace the current
section of the article.

### Closing
To close the postgres database you will need to stop the container, then remove it.
```shell script
podman container stop SpaceWeb
```

```shell script
podman container rm SpaceWeb
```
Then you can close the webapp by ctr^C in the terminal!

#### Thank you
Anzhela for debugging my program and for everyone that reads my articles!
I hope you liked my stories.