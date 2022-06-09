# todolist

For the project to launch, you will need Docker installed, and preferably the Make tool too.

Having Make:
 Clone the repository, enter the root folder, then execute "make". To stop the app, execute "make stop".

Not having Make:
 Clone the repository, enter the root folder, then execute "docker build -t todolist ." and "docker run -p 8000:8000 --name todolist todolist". To stop the app, execute "docker stop todolist".
