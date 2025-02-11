#####################################
#   STEP 1 build executable binary  #
#####################################
#FROM golang:1.18 as builder
FROM golang:1.23.0-bullseye AS builder

# Create appuser.
ENV USER=apimodelo
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build the binary.
# RUN CGO_ENABLED=0 GOOS=linux go build -o apirest
# Por causa do Kafka/Oracle/Etc tem que usar CGO_ENABLED=1 e -tags musl
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o apimodelo src/api/server.go

#####################################
#   STEP 2 build a small image      #
#####################################
FROM instantclient:19.25
ENV TZ="America/Sao_Paulo"
WORKDIR /app
COPY --from=builder /app/apimodelo .
#COPY --from=builder /app/templates /app/templates
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Use an unprivileged user.
USER apimodelo:apimodelo

# Definição da variável de ambiente para o cliente Oracle
#ENV LD_LIBRARY_PATH=/usr/lib/oracle/19.15/client64/lib
#ENV ORACLE_LIB_DIR=/usr/lib/oracle/19.15/client64/lib
ENV ORACLE_LIB_DIR=/opt/oracle/instantclient_19/

CMD ["./apimodelo"]  
EXPOSE 8800