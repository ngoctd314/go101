# The better solution

We have decided to utilize a common pattern when using Go channels, in order to create a 2-tier channel system, one for queueing jobs and another to control how many workers operate on the JobQueue concurrently.

The idea was to parallelize the uploads to S3 to a somewhat sustainable rate.