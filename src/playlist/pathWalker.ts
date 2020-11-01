import * as walk from "walk";
import Playlistmanager from "./playlistmanager";
import * as path from "path";

export class PathWalker {

    private readonly pathToScan: string = "/Users/a34p/Downloads";
    private readonly supportedExtensions: string [] = ['.avi', '.mkv', '.mp4'];

    constructor() {
        this.pathToScan = process.env.PATH_TO_SCAN || this.pathToScan;
        this.supportedExtensions = this.arrayFromEnv(process.env.SUPPORTED_EXTENSIONS, this.supportedExtensions);
    }

    public async generate(): Promise<string> {
        return new Promise((resolve, reject) => {
            const walker = walk.walk(this.pathToScan, {
                filters: this.arrayFromEnv(process.env.PATHS_TO_EXCLUDE, [])
            });
            const playlistManager = new Playlistmanager(this.pathToScan);
            walker.on("file", (root, fileStats, next) => {
                const filePath = `${root}/${fileStats.name}`;
                const fileExtension = path.extname(filePath);
                if (this.supportedExtensions.includes(fileExtension)) {
                    console.log(`New file: ${filePath}`);
                    playlistManager.addFile(filePath);
                }
                next();
            });
            walker.on("errors", () => {
                console.log("Error!");
                reject(undefined);
            });
            walker.on("end", () => {
                console.log("Finished!");
                playlistManager.save();
                resolve(playlistManager.playlistPath());
            });
        });
    }

    public playlistPath(): string {
        return new Playlistmanager(this.pathToScan).playlistPath();
    }

    private arrayFromEnv(envValue: string, fallbackValue: string[]): string[] {
        return envValue ? envValue.split(",") : fallbackValue;
    }
}