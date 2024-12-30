# Dockerfile

# Build stage
FROM node:18-alpine AS build

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy all files
COPY . .

# Create .env file from build arg
ARG VITE_API_URL
ENV VITE_API_URL=${VITE_API_URL}

# Build the app
RUN npm run build

# Serve stage
FROM nginx:alpine

# Copy nginx configuration
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

# Copy built files from build stage
COPY --from=build /app/dist /usr/share/nginx/html

# Add script to replace environment variables at runtime
RUN apk add --no-cache bash
COPY docker/env.sh /docker-entrypoint.d/40-env.sh
RUN chmod +x /docker-entrypoint.d/40-env.sh

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]