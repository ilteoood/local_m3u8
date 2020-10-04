import * as path from "path";
import * as fs from "fs";

export default class Playlistmanager {

    private playlistContent: string = "#EXTM3U\n";

    constructor(playlistName: string) {
        this.addRow(`#PLAYLIST:${playlistName}`);
    }

    public addFile(filePath: string): Playlistmanager {
        console.log(filePath);
        const directoryName = path.dirname(filePath);
        const fileName = path.basename(filePath);
        return this.addRow(`#EXTINF:-1 tvg-name="${fileName}", ${fileName}`)
            .addRow(`#EXTGRP:${directoryName}`)
            .addRow(`${process.env.BASE_URL || 'localhost:8080'}/${filePath}`);
    }

    public save(pathToSave: string): string {
        const fileName = `${Math.random().toString(36).substring(7)}.m3u8`;
        const filePath = path.join(pathToSave, fileName);
        fs.writeFileSync(filePath, this.playlistContent, {flag: 'w'});
        return filePath;
    }

    private addRow(rowContent: string): Playlistmanager {
        this.playlistContent = this.playlistContent.concat(rowContent, "\n");
        return this;
    }


}