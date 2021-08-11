FROM node:16-alpine

WORKDIR /app

COPY package-lock.json .
COPY package.json .
RUN npm install

COPY . .

CMD ["node", "index.js"]
#CMD ["sleep", "10000"]
