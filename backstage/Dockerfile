FROM registry.access.redhat.com/ubi8/nodejs-18-minimal AS build

# Install yarn, SQLite, Python, and build tools.
USER 0
RUN npm install -g yarn
RUN microdnf install -y sqlite sqlite-devel python3 make gcc-c++ gzip zlib-devel brotli-devel openssl-devel

# Copy app code
COPY . .

# build the packages
RUN yarn install --frozen-lockfile --production --network-timeout 600000
RUN ls -al && yarn tsc
RUN yarn build:backend

# Extract the skeleton and bundle archives
RUN mkdir packages/backend/dist/skeleton packages/backend/dist/bundle && \
    tar xzf packages/backend/dist/skeleton.tar.gz -C packages/backend/dist/skeleton && \
    tar xzf packages/backend/dist/bundle.tar.gz -C packages/backend/dist/bundle

# Build the serving image and install production dependencies
FROM registry.access.redhat.com/ubi8/nodejs-18-minimal

# Install yarn, SQLite and build tools
USER 0
RUN npm install -g yarn
RUN microdnf install -y sqlite sqlite-devel python3 make gcc-c++ gzip zlib-devel brotli-devel openssl-devel && microdnf clean all

ENV NODE_ENV=production

# Copy & install dependencies prepared by the build stage
COPY --from=build /opt/app-root/src/yarn.lock \
                  /opt/app-root/src/package.json \
                  /opt/app-root/src/packages/backend/dist/skeleton/ ./ \
                  /opt/app-root/src/packages/backend/dist/bundle/ ./

# Adjust permissions on the .cache directory for non-root user
RUN mkdir -p /opt/app-root/src/.cache && \
    chown -R 1001:0 /opt/app-root/src/

USER 1001

# Fix permissions on all files to prevent issues with randomized user IDs
RUN fix-permissions ./

# Expose the Backstage port
EXPOSE 7007

# Launch the Backstage application
CMD ["node", "packages/backend", "--config", "app-config.yaml"]
