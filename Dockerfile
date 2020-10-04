FROM node:slim
LABEL maintainer.name="Matteo Pietro Dazzi" \
    maintainer.email="matteopietro.dazzi@gmail.com" \
    version="1.0.0" \
    description="Scan all subdirectories and create a proper m3u8 file"
ENV PATH_TO_SCAN=
ENV BASE_URL=
ENV PATHS_TO_EXCLUDE=
ENV PLAYLIST_NAME=
ENV SUPPORTED_EXTENSIONS=
COPY . /local_m3u8
WORKDIR local_m3u8
RUN npm install
EXPOSE 3000
ENTRYPOINT npm run start
