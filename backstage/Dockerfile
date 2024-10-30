ARG BASE=registry.access.redhat.com/ubi9/nodejs-20-minimal@sha256:cb43bbd5fad5b7e5123cf5262651265711bfc4740ccc105afd743f0109f95020
# STAGE 1 - Facilitate caching by filtering for package.json files
# ---
FROM $BASE AS packages

COPY --chown=1001:0 package.json yarn.lock ./
COPY --chown=1001:0 packages/ packages/
COPY --chown=1001:0 plugins/ plugins/

# Remove all files except package.json
RUN find packages \! -name "package.json" -mindepth 2 -maxdepth 2 -exec rm -rf {} \+

# STAGE 2 - Build packages
# ---
FROM $BASE AS build

# Install yarn, SQLite and build tools. You can skip SQLite & build tools when not needed
USER 0
RUN npm install -g yarn
RUN microdnf install -y sqlite sqlite-devel && microdnf clean all
RUN microdnf install -y python3 make gcc-c++ gzip zlib-devel make sqlite sqlite-devel brotli-devel openssl-devel && microdnf clean all
USER 1001

# Install dependencies
ENV CYPRESS_INSTALL_BINARY=0
COPY --from=packages /opt/app-root/src/ ./
RUN --mount=type=cache,target=/opt/app-root/src/.cache/yarn,uid=1001,gid=0 \
    yarn install --frozen-lockfile --network-timeout 600000

# Copy app code
COPY --chown=1001:0 . .

# Compile & build the packages
RUN --mount=type=cache,target=/opt/app-root/src/.cache/yarn,uid=1001,gid=0 \
    yarn tsc && \
    yarn --cwd packages/backend build
RUN mkdir packages/backend/dist/skeleton packages/backend/dist/bundle \
    && tar xzf packages/backend/dist/skeleton.tar.gz -C packages/backend/dist/skeleton \
    && tar xzf packages/backend/dist/bundle.tar.gz -C packages/backend/dist/bundle


# STAGE 3 - Build the serving image and install prod dependencies
# ---
FROM $BASE

# Install yarn, SQLite and build tools. You can skip SQLite when not needed
USER 0
RUN npm install -g yarn
RUN microdnf install -y python3 make gcc-c++ gzip zlib-devel make sqlite sqlite-devel brotli-devel openssl-devel && microdnf clean all
USER 1001

ENV NODE_ENV production

# Copy & install dependencies prepared by the build stage
COPY --from=build /opt/app-root/src/yarn.lock \
                  /opt/app-root/src/package.json \
                  /opt/app-root/src/packages/backend/dist/skeleton/ ./ \
                  /opt/app-root/src/plugins/knative-event-mesh-backend/dist ./
RUN --mount=type=cache,from=build,target=/opt/app-root/src/.cache/yarn,uid=1001,gid=0 \
    yarn install --frozen-lockfile --production --network-timeout 600000

# Copy the built packages from the build stage
COPY --from=build /opt/app-root/src/packages/backend/dist/bundle/ ./
COPY --from=build /opt/app-root/src/templates ./

EXPOSE 7007

# This disables node snapshot for Node 20 to work with the Scaffolder
ENV NODE_OPTIONS="--no-node-snapshot"

# Fix permissions on all files to prevent problems when running with randomized user IDs
RUN fix-permissions ./

# Launch backstage app
CMD ["node", "packages/backend", "--config", "app-config.yaml"]
