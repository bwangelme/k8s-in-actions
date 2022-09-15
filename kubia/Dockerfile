# docker build -t bwangel/qae_kubia_$(date +%Y_%m_%d_%H%M) --build-arg app=kubia .

ARG build_base=golang:1.18
ARG base=golang:1.18

FROM ${build_base} as builder

RUN mkdir /code
WORKDIR /code
COPY . /code
RUN go build -o kubia .

FROM ${base}

COPY --from=builder /code/kubia /go/bin
ENTRYPOINT /go/bin/kubia