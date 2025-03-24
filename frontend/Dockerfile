FROM node:22-alpine
ARG TILE_URL=https://tile.openstreetmap.de/{z}/{x}/{y}.png
ARG TILE_KEY=""
ARG MAP_ATTRIBUTION="&lt;a href=&quot;https://www.openstreetmap.org/copyright&quot; target=&quot;_blank&quot;&gt;&copy; OpenStreetMap&lt;/a&gt; contributors"

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

ENV VITE_TILE_URL ${TILE_URL}
ENV VITE_TILE_KEY ${TILE_KEY}
ENV VITE_MAP_ATTRIBUTION ${MAP_ATTRIBUTION}

RUN npm run build

EXPOSE 4173

ENTRYPOINT ["npm", "run", "preview"]