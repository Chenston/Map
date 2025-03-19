@echo off
echo Starting MongoDB...
start cmd /k docker run --name mongo -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin123 mongo

echo Waiting for MongoDB to start...
timeout /t 5 >nul

echo Importing data into MongoDB...
docker cp states.json mongo:/states.json
docker exec -it mongo mongoimport --db us_states --collection states --file /states.json --jsonArray --username admin --password admin123 --authenticationDatabase admin

echo Starting backend...
start cmd /k "cd backend && go run main.go"

echo Starting frontend...
start cmd /k "cd frontend && npm install && npm run dev"

echo Waiting for frontend to start...
timeout /t 3 >nul

echo Opening browser...
start http://localhost:5173/

echo All services started successfully!
