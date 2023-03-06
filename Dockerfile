# First stage: Build dependencies and install them
FROM python:3.9-slim-buster as builder

WORKDIR /build

# Install system dependencies needed for building Python packages
RUN apt-get update && \
    apt-get install -y \
    build-essential \
    libpq-dev \
    libssl-dev \
    libffi-dev \
    libxml2-dev \
    libxslt1-dev \
    zlib1g-dev

# Copy the requirements file and install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

# Second stage: Copy only the installed dependencies and the necessary source code
FROM python:3.9-slim-buster

WORKDIR /app

# Copy the installed dependencies from the builder stage
COPY --from=builder /install /usr/local

# Copy the source code
COPY . /app/

CMD ["celery", "-A", "tasks", "worker", "--loglevel=INFO"]
