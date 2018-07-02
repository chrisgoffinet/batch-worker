# Batch Worker
This simple go program will create [N] workers, add integers to a channel, batch those integers until reaching `batchSize` then `process()` the work.

I wrote this for cases where you might want to batch units of work when using an iterator without making the code complex. The real world example I used it for was when processing data and wanted to batch up my Solr documents into the `/update` endpoint.