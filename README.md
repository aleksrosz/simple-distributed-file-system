Distributed file system (DFS) based and inspired by HDFS.
Project for improving Go language, gRPC and distributed systems skills.

The system is composite of three main parts:
1. CLI client
2. DataNodes - on which we have chunks of files.
3. MetaDatanode - which has information about the metadata of files and controls DataNodes.

All communication between DataNodes and MetaDatanode is realized via gRPC.

Which features are implemented:

CLI client:
1. Upload, and delete files from DFS.

DataNodes:
1. Send report to MetaDatanode about current health status.
2. Send report to MetaDatanode about current storage status.

MetaDatanode:
1. In-memory database of working DataNodes in DFS.
2. In-memory database of blockreports containing metadata about files.
3. Query DataNodes about their health status and update database according to that.