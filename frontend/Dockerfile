FROM node:22-alpine3.19

RUN mkdir /app
WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci && npm cache clean --force

COPY ./src ./src

EXPOSE 5173

ENTRYPOINT ["npm", "run"]