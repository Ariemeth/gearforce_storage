# Build Storage

## Running
Start local Arangodb
```
docker run -d --name arangodb-test1 -p 8529:8529 -e ARANGO_ROOT_PASSWORD=password -e ARANGODB_OVERRIDE_DETECTED_TOTAL_MEMORY=1G -e ARANGODB_OVERRIDE_DETECTED_NUMBER_OF_CORES=1  arangodb/arangodb:3.11.4
```