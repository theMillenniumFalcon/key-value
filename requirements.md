# key-value
a 1000 line non-proxying distributed key value store. Optimized for reading files between 1MB and 1GB.

### API
- GET /key
  - Supports range requests
  - 302 redirect to vaolume server
- {PUT, DELETE} /key
  - Blocks, 200 = writtten, anything else = nothing happened.
  
### Start Master Server (default port 3000)
```
./master.py -p 3000 /tmp/cachedb/
```

### Start Volume Server (default port 3001)
```
./volume.py -p 3001 /tmp/volume1/ localhost:3000
./volume.py -p 3002 /tmp/volume2/ localhost:3000
```