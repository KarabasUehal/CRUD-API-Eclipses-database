# CRUD-Eclipses-database
A CRUD-application that contains information about fictional eclipses.

How to start CRUD-Eclipse-app from Docker-container:

Open Dockerfile and use command in terminal:

docker build -t myapp:latest .

Make sure you put dot after space. 
After app's has build you need use command:

docker run -it -d -p 3000:3000 --name=Myapp_cont  myapp:latest

If you need to save the data after stopping/removing container, instead use:

docker volume create myapp_volume

docker run -v myapp_volume:/app -p 3000:3000 --name=Myapp_cont myapp:latest

Then you can go to url: http://localhost:3000 and make GET requests:
/eclipse - to get all eclipses;
/eclipse/id - to get some eclipse by id;

To make POST, PUT and DELETE requests you can use curl, Postman or similar apps:
/eclipse/delete/id - use DELETE request to delete some eclipse by id;
/eclipse/update/id - use PUT request to update some eclipse by id or create new eclipse if such id will not found;
/eclipse/add - use POST request to add new eclipse.
