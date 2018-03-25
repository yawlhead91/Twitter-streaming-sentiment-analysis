FROM node:latest

# Create app directory
WORKDIR /app

COPY . /app
RUN npm install -g

CMD [ "npm", "start" ]

EXPOSE 3000