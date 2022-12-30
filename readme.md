## RUN GO
go run math.go

## RUN TEST
go mod init math.go
go test -run TestSoma

## COVERAGE SONAR
go test -coverprofile=coverage.out

## Usando sonar-project.properties
docker run --rm  -v "C:\Users\leonardo.freitas\Documents\Projetos FullCycle\continuous-integration:/usr/src" sonarsource/sonar-scanner-cli

## Sonar Scanner cli
docker run \
    --rm \
    -e SONAR_HOST_URL="http://localhost:9000" \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=go-app" \
    -e SONAR_LOGIN="sqp_6328d1c07ae2c954566cf232c67f041326c3d4b1" \
    -v "C:\Users\leonardo.freitas\Documents\Projetos FullCycle\continuous-integration:/usr/src" \
    sonarsource/sonar-scanner-cli