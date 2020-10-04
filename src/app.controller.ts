import {Controller, Get, Header, HttpException, HttpStatus, Res} from '@nestjs/common';
import {PathWalker} from "./playlist/pathWalker";
import * as fs from "fs";

@Controller()
export class AppController {
    constructor() {
    }

    @Get('/playlist')
    @Header('Content-Type', 'application/x-mpegURL')
    @Header('Content-Disposition', `attachment; filename=${process.env.PLAYLIST_NAME || "Rclone"}.m3u8`)
    async getPlaylist(@Res() response) {
        const pathWalker = new PathWalker();
        const walkResult = await pathWalker.generate();
        if (walkResult) {
            return await fs.createReadStream(walkResult).pipe(response);
        }
        throw new HttpException({
            status: HttpStatus.INTERNAL_SERVER_ERROR
        }, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
