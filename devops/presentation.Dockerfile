FROM node:20-alpine

COPY ./frontend/resources /app/resources

WORKDIR /app/resources

RUN npm install
RUN npm run build

COPY ./frontend/presentation /app/presentation

WORKDIR /app/presentation

RUN npm install
RUN npm run build

COPY ./devops/presentation.prod.json ./

RUN adduser -D nonroot
USER nonroot
ENV SERVICE_CONFIG_PATH=./presentation.prod.json

EXPOSE 8080

ENTRYPOINT ["node", "./dist/server/server.js"]
