import {Controller, Get, Header, HttpException, HttpStatus, Res} from '@nestjs/common';
import {PathWalker} from "./playlist/pathWalker";
import * as fs from "fs";

@Controller()
export class AppController {
    constructor() {
    }

    @Get('/playlist/generate')
    @Header('Content-Type', 'application/x-mpegURL')
    @Header('Content-Disposition', `attachment; filename=${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`)
    async generatePlaylist(@Res() response) {
        const pathWalker = new PathWalker();
        const walkResult = await pathWalker.generate();
        if (!walkResult) {
            throw new HttpException({
                status: HttpStatus.INTERNAL_SERVER_ERROR
            }, HttpStatus.INTERNAL_SERVER_ERROR);
        }
        response.download(walkResult, `${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`);
        return response;
    }

    @Get('/playlist')
    @Header('Content-Type', 'application/x-mpegURL')
    @Header('Content-Disposition', `attachment; filename=${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`)
    async getPlaylist(@Res() response) {
        response.download(new PathWalker().playlistPath(), `${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`);
        return response;
    }
}
