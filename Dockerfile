FROM library/golang:onbuild

WORKDIR /go/src/hs_replay_bot

RUN go get github.com/golang/dep/cmd/dep