#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Function to log messages
log() {
  echo -e "\e[32m[INFO] $1\e[0m"
}

# Update the package list
log "Updating package list..."
sudo yum update -y

# Install Docker
log "Installing Docker..."
sudo amazon-linux-extras enable docker
sudo yum install -y docker

# Start Docker service
log "Starting Docker service..."
sudo systemctl start docker
sudo systemctl enable docker

# Add the current user to the Docker group
log "Adding user '$USER' to the Docker group..."
sudo usermod -aG docker $USER

# Verify Docker installation
log "Verifying Docker installation..."
docker --version

# Final instructions
log "Docker installation complete. Please log out and log back in for changes to take effect."
