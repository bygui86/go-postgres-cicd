
# --- buider stage

# base image
FROM golang:1.16-buster AS builder

# set working dir
WORKDIR /src

# prepare go-modules
COPY go.mod go.mod
RUN go mod download

# compile code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app


# --- final stage

# base image
FROM alpine

# copy application
WORKDIR /bin/
COPY --from=builder /bin/app .

# expose ports
EXPOSE 8080
EXPOSE 9090

# set user
USER 1001

# run application
ENTRYPOINT "/bin/app"
