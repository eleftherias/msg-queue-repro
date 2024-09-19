### Run
```
docker compose up
```

### Errors
When running the code as-is, the message is repeatedly retried, and never added to the DLQ. Example logs:

<details>
  <summary>Expand logs using SQL</summary>

Note that the process was manually terminated. It would have likely continued indefinitely.

```
server-1  | 2024/09/19 11:16:16 sending message f0dd0fef-7875-4f08-ae5a-0a8d16707d01, correlation id: 390c9ebd-1789-4e43-8a37-b9c9e8e503e6
server-1  | 2024/09/19 11:16:16 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:27 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:27 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:38 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:16:38.593858 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.101677547s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:16:38 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:16:46.492239 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:16:47 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:49 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:16:49.599812 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.107720468s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:16:49 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:58 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:16:58 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:09 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:09.607599 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.103205296s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:17:09 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:17.503301 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:17:18 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:20 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:20.611498 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.106395302s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:17:20 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:29 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:29 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:40 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:40.629975 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.101326588s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:17:40 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:48.527435 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:17:49 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:17:51 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:17:51.634439 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.105707926s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:17:51 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:00 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:00 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:11 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:11.692690 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.105273505s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:18:11 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:19.573412 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:18:20 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:22 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:22.698556 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.11212926s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:18:23 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:31 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:31 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:42 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:42.691917 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.103809171s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:18:42 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:50.581968 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:18:51 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:18:53 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:18:53.696136 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.108232593s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:18:54 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:02 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:02 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:13 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:19:13.687156 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.101010422s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:19:13 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:19:21.586221 subscriber.go:455: 	level=INFO  msg="Discarding queued message, context canceled" consumer_group= msg_uuid=f0dd0fef-7875-4f08-ae5a-0a8d16707d01 subscriber_id=01J84ZYEFXAX7Y9YWF3BBR4CZE topic=messages_topic
server-1  | 2024/09/19 11:19:22 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:24 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | [watermill] 2024/09/19 11:19:24.687844 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.101586427s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:19:25 deadLetterHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:33 structHandler returning error for message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
server-1  | 2024/09/19 11:19:33 structHandler received message f0dd0fef-7875-4f08-ae5a-0a8d16707d01
```
</details>

When using the gochannel Pub/Sub, instead of SQL, the message is added to the DLQ after 3 reties. Example logs:
<details>
    <summary>Expand logs using gochannel</summary>
    
```
server-1  | 2024/09/19 11:09:34 sending message 4ad005cf-4d5d-4568-b006-0763a8167b80, correlation id: aac2d8e3-7f5e-458d-931e-5075b11fcd1c
server-1  | 2024/09/19 11:09:34 structHandler received message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | 2024/09/19 11:09:45 structHandler returning error for message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | 2024/09/19 11:09:45 structHandler received message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | 2024/09/19 11:09:56 structHandler returning error for message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | [watermill] 2024/09/19 11:09:56.947384 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=11.119478089s err="error from structHandler" max_retries=3 retry_no=1 wait_time=100ms
server-1  | 2024/09/19 11:09:56 structHandler received message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | 2024/09/19 11:10:07 structHandler returning error for message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | [watermill] 2024/09/19 11:10:07.962281 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=22.135531761s err="error from structHandler" max_retries=3 retry_no=2 wait_time=0s
server-1  | 2024/09/19 11:10:07 structHandler received message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | 2024/09/19 11:10:18 structHandler returning error for message 4ad005cf-4d5d-4568-b006-0763a8167b80
server-1  | [watermill] 2024/09/19 11:10:18.977979 retry.go:78: 	level=ERROR msg="Error occurred, retrying" elapsed_time=33.151471349s err="error from structHandler" max_retries=3 retry_no=3 wait_time=0s
server-1  | 2024/09/19 11:10:18 deadLetterHandler received message 4ad005cf-4d5d-4568-b006-0763a8167b80
```
</details>