# Build stage
FROM node:22-alpine AS build

WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Copy package files and install deps
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# Copy source and build
COPY . .
RUN pnpm build

# Production stage
FROM nginx:alpine

# Copy built app
COPY --from=build /app/dist /usr/share/nginx/html

# Optional: copy custom nginx config (ensure it supports SPA routing)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]