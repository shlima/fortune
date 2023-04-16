ARG BUILD_BASE=golang:1.20-alpine
ARG ISO_BASE=alpine

FROM ${BUILD_BASE} as build
WORKDIR /go/src/github.com/github.com/shlima/fortune
ENV BUILD_DEPS make
COPY . .
RUN apk add --update --no-cache $BUILD_DEPS
RUN make linux

FROM ${ISO_BASE} as iso
WORKDIR /app
ENV PATH="/app:${PATH}"
COPY --from=build /go/src/github.com/github.com/shlima/fortune/build/linux ./fortune
COPY addresses/Bitcoin/2023 addresses/Bitcoin/2023
RUN addgroup --gid 2019 user && \
    adduser --disabled-password --uid 2019 --ingroup user --gecos user user
RUN chown -R user:user ./
USER user
SHELL ["/bin/bash", "-c"]
ENTRYPOINT ["fortune"]
