ARG BASE_TAG=server-vulkan
FROM ghcr.io/ggml-org/llama.cpp:${BASE_TAG}
#FROM docker.io/rocm/llama.cpp:llama.cpp-b5997_rocm6.4.0_ubuntu24.04_server
RUN mkdir /config && chmod a+rw /config
WORKDIR /app
COPY build/llama-swap-linux-amd64 /app/llama-swap
HEALTHCHECK CMD curl -f http://localhost:8080/ || exit 1
ENTRYPOINT [ "/app/llama-swap", "-config", "/config/config.yaml" ]