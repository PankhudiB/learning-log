## GFS - Google File System

```Started reading GFS paper.```

### Architecture

1. 1 Master - multiple chunk servers
2. Chunk server -> Linux machine
3. Client interacts with Master for metadata info
3. Client interacts with ChunkServers for actual data
4. Files -> divided into chunks
5. Neither client nor ChunkServers cache file