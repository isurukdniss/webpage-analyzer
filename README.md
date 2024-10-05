# Webpage Analyzer

The **Webpage Analyzer** is a Go based web application that allows users to analyze a webpage by entering a URL.

## Features
- Detects the HTML version used by the webpage.
- Retrieves the page title.
- Counts the number of headings at each level (`<h1>` to `<h6>`).
- Counts internal and external links, and detects any inaccessible links.
- Checks if the webpage contains a login form.
- Displays error messages for unreachable URLs or invalid responses from the server.

## Installation

#### Prerequisites
- Go (version 1.21.3 or higher)

#### Clone the Repository
```
git clone https://github.com/isurukdniss/webpage-analyzer
```

#### Install the Dependancies
```
go mod tidy
```

#### Run the Project
```
make run
```

The web server will start on `http://localhost:8080`. Please navigate to that URL using your web browser.

## Additional Commands

#### Build Binary
```
make build
```
The executable file can be found in the `/bin` directory.

#### Run Unit Tests
```
make test
```

#### Clean the Build Directory
```
make clean
```

## Continues Integration
The Continues Integration is achieved using the GitHub Actions. A workflow builds the project and run all unit test when a new commit is pushed to the `main` branch or a pull request is merged to the `main` branch.

Workflow runs can be found [here](https://github.com/isurukdniss/webpage-analyzer/actions). 

## Suggested Improvements
- Improving the UI with more advanced styling or using a front-end framework like React or Vue.
- Containerization of the application.
- Configure automated deployment to a cloud platform (eg. GCP CloudRun).




