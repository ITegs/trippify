FROM node:22-alpine3.19 AS build-stage

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine AS production-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

# COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
