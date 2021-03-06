package main

const DefaultRoot = "/var/lib/verm"
const DirectoryPermission = 0777

const DefaultListenAddress = "0.0.0.0"
const DefaultPort = "3404"

const DefaultMimeTypesFile = "/etc/mime.types"
const DefaultDirectoryIfNotGivenByClient = "/default"
const UploadedFieldFieldForMultipart = "uploaded_file"

const ReplicationQueueSize = 1000000
const ReplicationMissingQueueSize = 10000
const ReplicationBackoffBaseDelay = 1
const ReplicationBackoffMaxDelay = 120
const ReplicationNetworkTimeout = 30
const ReplicationRequestTimeout = 3600 // want this to be large enough for a very large file over a relatively congested link; we really rely on TCP keepalives at both ends to ensure the connection goes away eventually if the network or other end has died
const ReplicationMissingFilesPath = "/_missing"
const ReplicationMissingFilesBatchSize = 256*1024 // bytes, but only approximate
const ReplicationMissingFilesBatchTime = 1 // seconds before we send even a small batch

const ReplicaProxyTimeout = 15

const ShutdownResponseTimeout = 15
