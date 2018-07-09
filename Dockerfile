FROM circleci/golang:1.9

ENV DISPLAY :99

RUN sudo apt-get update \
    && sudo apt-get install --no-install-recommends --yes \
    libsdl2-mixer-dev libsdl2-image-dev libsdl2-ttf-dev libsdl2-gfx-dev xvfb \
    && sudo apt-get clean \
    && sudo rm -rf /root/.cache \
    && sudo rm -rf /var/lib/apt/lists/*

RUN go get -v github.com/veandco/go-sdl2/sdl \
    && go get -v github.com/veandco/go-sdl2/img \
    && go get -v github.com/veandco/go-sdl2/mix \
    && go get -v github.com/veandco/go-sdl2/ttf \
    && go get github.com/mattn/goveralls
RUN mkdir -p /go/src/github.com/tommyblue/matrigo

COPY . /go/src/github.com/tommyblue/matrigo

WORKDIR /go/src/github.com/tommyblue/matrigo

CMD ["./script/docker-test"]
