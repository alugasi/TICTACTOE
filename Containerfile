FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
RUN go mod download
RUN go get github.com/alugasi/TICTACTOE/

# Copy source code

COPY main  ./
COPY logic ./
COPY view ./

# build
RUN cd main & go build -o /tictactoe

# Expose port
EXPOSE 5000

# Run
CMD [ "/tictactoe" ]