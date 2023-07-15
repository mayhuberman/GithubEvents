# GithubEvents

# The Task
- In this task you are required to periodically watch public Github events as they happen, through the relevant github events API, and process them in order to allow the functionality that's described below.
  1. For each event that you fetch, you are required to collect the following:
      - Event Type: You are required to store the event count per unique event type
      - Actor: Unique name of the actor. You are required to store the last 50 actors.
      - Repo URLs: you are required to store the last 20 repository URLs that you came across. 
      - Emails: You are required to store all of the unique email addresses that you come across
  - You should support retrieval of the data that you have collected either through API endpoints that your solution will expose, and that will return the collected data in JSON format, or by building a minimalistic web UI that will display it. You are free to choose between the two methods.
  2.In addition, you are required to support on-demand querying for the amount of stars per repository that you stored, and to return a list of repositories sorted by the amount of stars they have

# Bonus parts
- Use docker to wrap your application and make it ready to build and launch (no need to submit an actual image)
- Use a database of your choosing to allow persistency, rather than keeping things in memory

# How to run my app
On your terminal
1. "docker build -t myapp-image ."
2. "docker run -p 9090:8080 myapp-image"
   Now you have the periodic task running in the background

# API Documentation
1. GET "/events" -> json
fetching all the data required in section 1 of the requirements
On your terminal:  curl "http://localhost:9090/events"
response example:
{
   "actors":[
      "MrTuzki",
      "SanthoshK4494",
      "KrunalKevadiya18",
      "github-actions[bot]",
      "xyhtac",
      "dependabot[bot]",
      "AfricanGirl1",
      "josephperrott",
      "Muideen7",
      "davee8k"
   ],
   "emails":[
      "41898282+github-actions[bot]@users.noreply.github.com",
      "olayeyeayomide2000@gmail.com",
      "abhi247997@gmail.com",
      "90658725+gfourty@users.noreply.github.com",
      "119910615+happyfish2024@users.noreply.github.com",
      "moraab30@gmail.com",
      "2921768103@qq.com",
      "robert.debock@adfinis.com",
      "josephperrott@gmail.com",
      "dabreadman@users.noreply.github.com",
      "davee8k@users.noreply.github.com",
      "46420764+abhi2479@users.noreply.github.com",
      "117566446+sakshisahu612@users.noreply.github.com",
      "github-actions[bot]@users.noreply.github.com",
      "129385590+abhishekkumar-s@users.noreply.github.com",
      "root@traewelling.server.home"
   ],
   "eventTypes":{
      "CreateEvent":8,
      "DeleteEvent":1,
      "ForkEvent":1,
      "PullRequestEvent":1,
      "PullRequestReviewEvent":2,
      "PushEvent":15,
      "ReleaseEvent":1,
      "WatchEvent":1
   },
   "urls":[
      "https://api.github.com/repos/grakke/code",
      "https://api.github.com/repos/AfricanGirl1/alx-system_engineering-devops",
      "https://api.github.com/repos/SanthoshK4494/docker123",
      "https://api.github.com/repos/2921768103/1",
      "https://api.github.com/repos/robertdebock/ansible-role-facts",
      "https://api.github.com/repos/consumet/rapidclown",
      "https://api.github.com/repos/angular/dev-infra",
      "https://api.github.com/repos/Muideen7/alx-frontend",
      "https://api.github.com/repos/THIS-IS-NOT-A-BACKUP/jq",
      "https://api.github.com/repos/davee8k/autoloader"
   ]
}
3. GET "/repo-stars" -> json
fetching all the data required in section 2 of the requirements
On your terminal: curl "http://localhost:9090/repo-stars"
response example:
{
   "repoStars":[
      "https://api.github.com/repos/djprofessorkash/setlist: 0",
      "https://api.github.com/repos/AlaaTDD/astaze: 0",
      "https://api.github.com/repos/okraHQ/okra-react-native-official: 0",
      "https://api.github.com/repos/koeztu/blockchain-game-analysis: 0",
      "https://api.github.com/repos/dmihalcik-virtru/client-web: 0",
      "https://api.github.com/repos/DataDog/envoy-filter-example: 0",
      "https://api.github.com/repos/happyfish2024/mins: 1",
      "https://api.github.com/repos/liferay-headless/liferay-portal: 1",
      "https://api.github.com/repos/spencerwmiles/vscode-task-buttons: 9",
      "https://api.github.com/repos/joshka/junit-json-params: 49"
   ]
}

# Project Structure
- /clients/GithubClients: contains the method that interacts with Github's public API. It retrieves the events and returns a slice of the corresponding struct type.
- /services/GithubService: includes methods responsible for extracting the desired data from the events slice. It also contains methods that return the data in JSON format.
- /models: This package holds all the data types used in the project, defining the structure and properties of the data.
- /utils/Utils: contains utility methods used throughout the project, providing various helper functionalities.
- Dockerfile: This file contains the instructions for building a Docker image. It specifies the commands that users can run on the command line to assemble the image.
- go.mod: This file defines the module's path and specifies its dependencies.
- main: This file contains the main entry point of the application. It runs the periodic task and listens to the APIs, coordinating the different components of the project.

