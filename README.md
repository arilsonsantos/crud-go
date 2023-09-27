## CRUD-GO

An example of CRUD with Golang.  

### To run tests locally  

go test ./... -json > report.json &&  go test ./... -covermode=atomic -coverprofile=coverage.out

### To check on SonarQube  

docker run -d --name sonarqube -p 9000:9000 -p 9092:9092 sonarqube  

#### Download SonarScanner  
https://docs.sonarsource.com/sonarqube/9.9/analyzing-source-code/scanners/sonarscanner/  

Em ambiente dev:
/sonar-scanner -Dproject.settings=sonar-project-dev.properties  -Dsonar.organization=arilsonsantos -Dsonar.projectKey=arilsonsantos_crud-go -Dsonar.sources=. -Dsonar.host.url=http://localhost:9000
  

#### SonarCloud  
https://sonarcloud.io/summary/overall?id=arilsonsantos_crud-go