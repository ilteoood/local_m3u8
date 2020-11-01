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
        this.checkWalkResult(walkResult);
        return this.sendFile(response, walkResult);
    }

    private checkWalkResult(walkResult: string) {
        if (!walkResult) {
            throw new HttpException({
                status: HttpStatus.INTERNAL_SERVER_ERROR
            }, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @Get('/playlist')
    @Header('Content-Type', 'application/x-mpegURL')
    @Header('Content-Disposition', `attachment; filename=${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`)
    async getPlaylist(@Res() response) {
        return this.sendFile(response);
    }

    private sendFile(response, playlistPath = new PathWalker().playlistPath()) {
        response.download(playlistPath, `${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`);
        return response;
    }
}
