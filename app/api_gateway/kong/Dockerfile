# Build stage
FROM --platform=linux/amd64 kong/deck 
WORKDIR /usr/src

COPY app/api_gateway/kong .

ENV DECK_KONG_ADDR=http://kong:8001
CMD ["sync"]
