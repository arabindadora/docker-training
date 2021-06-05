FROM node:lts-alpine

# basic http server
RUN npm install -g http-server

# install dependencies
WORKDIR /app
COPY app/package*.json ./
RUN npm install

# copy and build app
COPY app ./
RUN npm run build

# serve
EXPOSE 8001
CMD [ "http-server", "-p", "8001" ]
