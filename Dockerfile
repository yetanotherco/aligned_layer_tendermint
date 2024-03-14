FROM ignitehq/cli as builder

USER root
WORKDIR /root

COPY . .

RUN ignite chain build 

ENTRYPOINT [ "alignedlayerd" ]
