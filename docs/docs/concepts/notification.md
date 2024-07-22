# Notification

Notification is the main feature in Siren. Siren capables to send notification to various receivers (Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren able to routes notification directly to receivers or to specific receivers by matching notification key-value labels with the provided label matchers.

Sending bulk notifications is also supported in Siren. But this mode is only supporting subscription-based routing.

Below is how the notification is implemented in Siren

![Notification in Siren](../../static/img/siren_notification.svg)

Apart from [Alert Webhook](./alert.md), there are another entry points in Siren:
- [POST /notifications API](../apis/siren-service-post-notification.api.mdx)
    - To send a single notification through receiver based routing or subscription based routing
- [POST /bulk-notifications API](../apis/siren-service-post-bulk-notifications.api.mdx)
    - To send bulk notifications through subscription based routing

From this point, the notification is being dispatched by a `Notification Dispatcher` according to the strategy that it took (Bulk or Single Notification).

The dispatched notification will be evaluated based on the routing strategy too. If it is receiver based routing, the notification will be routed to the receiver. If it is subscription based routing, the notification will be evaluated by the subscription matcher.

Post this process, the `Meta Messages` will be generated. The generated meta messages could be reduced or grouped by a certain configuration. At the end, the meta messages will be rendered and this process produces the ![Notification Message](#notification-message) that will be enqueued and handled to be sent to the external notification services (e.g. Slack, Pagerduty, etc).

## Notification Meta Message
During dispatching Notification, the routed Notification will generate some `Meta Messages` from a single Notification. This is non-rendered version and the prototype of a ![Notification Message](#notification-message).
This model is useful for Siren to manipulate the meta messages like grouping or other transformation within the `Reducer` step.

## Notification Message

This is the rendered version of a ![Notification Meta Message](#notification-meta-message). In this state, the message would contain all information in the format that the receiver recognize or support. Each receiver might expect different payload. User needs to pass notification message payload in the same format as what receiver (notification vendor, e.g. slack, pagerduty, etc) expected. All message payload contracts could be found under [receivers](../receivers/slack.md).

### Templating Notification Message Payload

Message payload in notification could also be reused by defining template and passing some variables needed. See [template](../guides/template.md) for further details on how to use the template feature.

### Notification Message State

There are several possible states of notification message when Siren tries to send it.

- `enqueued` means notification message has been just enqueued. This is a fresh notification message that is ready to be dequeued.
- `pending` means the notification message is dequeued and the process of sending notification is still on-going.
- `failed` means there is something wrong when Siren tried to send the notification. The `last error` and `number of attempt` are recorded in the message itself.
- `published` means the message has successfully been sent.

### Notification Message Expiry Duration

It is possible to set notification message to be expired after some duration. If the expiry duration is not set, this assumes the notification message has no expiration time.


## Queue

Queue is used as a buffer to avoid pressure when notifications are being sent. Siren implements Queue as a plugin. Currently there are two kind of queue plugin supported: 
- in-memory (not for production usage)
- postgres

User could choose the which queue to use by mentioning it in the [config](../reference/server_configuration.md).

### In-memory Queue

In-memory queue simulates a queue with Go channel. The usage is intended to be used in development only.

### PostgreSQL Queue

Siren uses Postgres `SKIP LOCK` feature to implement queues with postgres.

## Notification Handlers

Notification handler responsibles to dequeue message and send notification to the receivers. There are two kind of Notification Handler in Siren, the main notification handler and the notification dlq handler. Both could be configured in the [server configuration](../reference/server_configuration.md).

Notification handlers could be run in the same process inside server or could be run separately as a [worker](../guides/workers.md). It is also possible for the handlers to only process specific receiver types. This can be done by specifying supported `receiver_types` in the [config](../reference/server_configuration.md). One could also configure the notification to be dequeued in batch by specifying `batch_size > 1` in the config.

### Notification Message Handler

The notification message handler is a main handler that will dequeue the supported `receiver_type` with `batch_size` number messages that are not expired from the main queue and try sending each message to the desired receivers. 

If there is an error, main notification handler will clasify whether the error is retryable or not (e.g. if bad request, it is non-retryable), mark the message as `failed` and queue it to DLQ.

### Notification DLQ handler

The notification dlq handler will dequeue the supported `receiver_type` with `batch_size` number messages that are not expired and have `failed` status from the dlq and try sending each message to the desired receivers. If there is an error, dlq handler will clasify whether the error is retryable or not. Notification handler will re-pick up the `failed` messages that are retryable and has been retried less than or equal max tries config.
