<<<<<<< HEAD
# Use nginx:alpine as base image
FROM nginx:alpine

# Copy all files from the current directory into the Nginx serving directory
COPY . /usr/share/nginx/html

# Expose port 80 to outside traffic
EXPOSE 80

# Command to run Nginx server in foreground
CMD ["nginx", "-g", "daemon off;"]
=======
FROM nginx:alpine

# Copy static files and other content
COPY . /usr/share/nginx/html

# Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80 for Nginx
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
>>>>>>> main
