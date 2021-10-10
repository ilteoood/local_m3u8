FROM golang:1.16-alpine3.14 as builder
ADD . /local_m3u8
WORKDIR /local_m3u8
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod download && \
    go test ./... && \
    go build -o local_m3u8 ./main

###################################

FROM scratch
USER scratchuser
COPY --from=builder "/local_m3u8/local_m3u8" "./local_m3u8"
ENTRYPOINT [ "./local_m3u8" ]