ARG base_image
FROM ${base_image} AS base

# Install packages that we want to make available at build time
USER root
RUN apt update && \
  apt install -y git wget jq && \
  rm -rf /var/lib/apt/lists/*

# Set user and group
ARG cnb_uid=1001
ARG cnb_gid=1001
USER ${cnb_uid}:${cnb_gid}

# Set required CNB information
ENV CNB_USER_ID=${cnb_uid}
ENV CNB_GROUP_ID=${cnb_gid}
