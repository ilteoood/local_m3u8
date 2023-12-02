FROM golang:1.21-alpine3.18 as builder
ADD . /local_m3u8
WORKDIR /local_m3u8
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod download && \
    go test ./... && \
    go build -o local_m3u8 ./main && \
    echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

###################################

FROM scratch
COPY --from=builder "/local_m3u8/local_m3u8" "./local_m3u8"
COPY --from=builder "/etc_passwd" "/etc/passwd"
USER nobody
ENTRYPOINT [ "./local_m3u8" ]
