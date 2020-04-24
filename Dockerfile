FROM dockercloud/hello-world:latest

ARG IMAGE_TAG

ENV NAME ${IMAGE_TAG}