import * as walk from "walk";
import Playlistmanager from "./playlistmanager";
import * as path from "path";

export class PathWalker {

    private readonly supportedExtensions: string [] = ['.avi', '.mkv', '.mp4'];

    constructor(private pathToScan: string = "/media/movies") {
    }

    async generate(): Promise<string> {
        return new Promise((resolve, reject) => {
            const walker = walk.walk(this.pathToScan, {});
            const playlistManager = new Playlistmanager("Rclone");
            walker.on("file", (root, fileStats, next) => {
                console.log(`New file: ${fileStats.name}`)
                const fileExtension = path.extname(fileStats.name);
                if (this.supportedExtensions.includes(fileExtension)) {
                    playlistManager.addFile(fileStats.name);
                }
                next();
            });
            walker.on("errors", () => {
                reject(undefined);
            });
            walker.on("end", () => {
                const playlistDirectory = playlistManager.save(this.pathToScan);
                resolve(playlistDirectory);
            });
        });
    }
}