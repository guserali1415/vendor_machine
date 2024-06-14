# Vendor Machine application
This is PWA based application created via **vuejs** for frontend and used **golang** and **gin** for backend that shipped into **docker** image.


## Getting started

1. Clone the repository
2. In the linux terminal navigate to project root folder and run `docker run -it-rm --name machine machine` to build docker image
3. run `docker run -it -p 5000:5000 -p 8080:8080 --rm --name machine machine` to run instance of image as container
4. Open http://127.0.0.1:5000 in your browser

## Tips
- for developing vuejs application use `/web` folder and eventually run `npm run build` and copy `/web/app/dist` folder contents into `/server/web folder`
- Test file is in /server folder. use `go test`
- Screenshots added