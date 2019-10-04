# Stage 1
FROM node:8 as react-build
WORKDIR /app
COPY ./react/ ./
RUN yarn
RUN yarn build

#Stage 2
FROM registry.access.redhat.com/ubi7/ubi
ADD ./battlefield-ui /
ADD ./resource /resource
COPY --from=react-build /app/build /static

EXPOSE 8080

ENTRYPOINT ["/battlefield-ui"]
