# local_m3u8
Scan all subdirectories and create a proper m3u8 file.

## Env variables
- PLAYLIST_NAME: name used to generate the playlist, **without .m3u8** extension. Default: `Rclone`;
- PATH_TO_SCAN: path in which files will be searched. Default: `./`;
- BASE_URL: prefix url for playlist generation: Default: `localhost:8080`;
- PATHS_TO_EXCLUDE: paths to exclude from generation, separated by `,`.;
- SUPPORTED_EXTENSIONS: extension which will be used for the generation, separated by `,`. Default: `.avi,.mkv,.mp4`;
