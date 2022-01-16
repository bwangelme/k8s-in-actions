# docker build -t bwangel/qae_kubia_$(date +%Y_%m_%d_%H%M) --build-arg app=kubia .

ARG build_base=golang:1.17
ARG base=golang:1.17

FROM ${build_base} as builder

ARG app

RUN mkdir /code
WORKDIR /code
COPY . /code
RUN go build -o ${app} .

FROM ${base}

ARG app

COPY --from=builder /code/${app} /go/bin
