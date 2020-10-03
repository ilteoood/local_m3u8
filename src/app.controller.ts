import {Controller, Get, HttpException, HttpStatus} from '@nestjs/common';
import {PathWalker} from "./playlist/pathWalker";
import * as fs from "fs";

@Controller()
export class AppController {
    constructor() {
    }

    @Get('/playlist')
    async getPlaylist() {
        const pathWalker = new PathWalker(process.env.PATH_TO_SCAN);
        const walkResult = await pathWalker.generate();
        if (walkResult) {
            return fs.createReadStream(walkResult);
        }
        throw new HttpException({
            status: HttpStatus.INTERNAL_SERVER_ERROR
        }, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
