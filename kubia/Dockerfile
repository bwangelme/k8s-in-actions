# docker build -t bwangel/echo:latest .

ARG build_base=golang:1.18
ARG base=alpine

FROM ${build_base} as builder

RUN mkdir /code
WORKDIR /code
COPY . /code
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o http_echo .

FROM ${base}

COPY --from=builder /code/http_echo /bin