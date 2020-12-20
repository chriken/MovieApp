1. Clone the repo.  
``` git clone https://github.com/chriken/MovieApp.git ```
---
2. Change directory to root of git.
``` cd MovieApp ```
 ---
3. Build the docker image.
```  docker build --tag my-movies . ```
---
4. Launch the docker container.
``` docker run -e apiKey=[YourKey] -p 8080:8081 -it my-movies  ```
---
5. Open a browser window or curl to ->  http://localhost:8080/movies 
