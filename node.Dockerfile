FROM node:21.6.2

WORKDIR /faucet

COPY ./faucet/package.json .
COPY ./faucet/package-lock.json .
RUN npm install

COPY ./faucet .
ENTRYPOINT [ "node" ]
