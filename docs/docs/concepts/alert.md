# Alert

Alert is being one of the source of notification in Siren. Alert is ingested via ![Alert Webhook](../apis/siren-service-create-alerts.api.mdx) and the webhook is registered to a Monitoring System (e.g. Prometheus, Cortex).
Once ingested by Siren, alert will be stored as Alert History. Alerts that are being sent by provider will be translated to a `Notification` by `Notification Service`. From this point, the notification is being dispatched by a `Notification Dispatcher` according to the strategy that it took.

The next process could be found under ![Notification](./notification.md).