FROM node:20-alpine

WORKDIR /app/api

COPY api/package*.json ./

RUN npm install

COPY api/src ./src
COPY api/tsconfig.json ./

EXPOSE 3000

CMD ["npm", "run", "dev"]
