# syntax=docker/dockerfile:1
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
  libxkbcommon0 \
  ca-certificates \
  git \
  golang \
  unzip \
  libc++1 \
  vim \
  && apt-get clean autoclean

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /workspace

COPY . /workspace
