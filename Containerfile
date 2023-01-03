FROM golang:1.18-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
RUN go mod download

# Copy source code

COPY main.go  ./
COPY logic/ logic/
COPY view/ view/

# build
RUN go build -o /tictactoe

# Expose port
EXPOSE 5000

# Run
CMD [ "/tictactoe" ]