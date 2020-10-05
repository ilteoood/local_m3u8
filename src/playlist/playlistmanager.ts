import * as path from "path";
import * as fs from "fs";

export default class Playlistmanager {

    private playlistContent: string = "#EXTM3U\n";

    constructor(private pathToScan: string,
                private playlistName: string = process.env.PLAYLIST_NAME || "Rclone") {
        this.addRow(`#PLAYLIST:${playlistName}`);
    }

    public addFile(filePath: string): Playlistmanager {
        const fileName = path.basename(filePath);
        const baseUrl = process.env.BASE_URL || 'localhost:8080';
        const relativePath = this.findRelativePath(filePath, this.pathToScan);
        return this.addRow(`#EXTINF:-1 tvg-name="${fileName}", ${fileName}`)
            .addRow(`#EXTGRP:${relativePath}`)
            .addRow(`${baseUrl}${relativePath}`);
    }

    private findRelativePath(filePath: string, pathToScan: string) {
        return filePath.replace(pathToScan, '');
    }

    public save() {
        fs.writeFileSync(this.playlistPath(), this.playlistContent, {flag: 'w'});
    }

    public playlistPath(): string {
        const fileName = `${this.playlistName}.m3u8`;
        return path.join(this.pathToScan, fileName);
    }

    private addRow(rowContent: string): Playlistmanager {
        this.playlistContent = this.playlistContent.concat(rowContent, "\n");
        return this;
    }


}