"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[53],{1109:e=>{e.exports=JSON.parse('{"pluginId":"default","version":"current","label":"Next","banner":null,"badge":false,"noIndex":false,"className":"docs-version-current","isLast":true,"docsSidebars":{"docsSidebar":[{"type":"link","label":"Introduction","href":"/siren/docs/introduction","docId":"introduction"},{"type":"link","label":"Use Cases","href":"/siren/docs/use_cases","docId":"use_cases"},{"type":"link","label":"Installation","href":"/siren/docs/installation","docId":"installation"},{"type":"category","label":"Tour","items":[{"type":"link","label":"Introduction","href":"/siren/docs/tour/introduction","docId":"tour/introduction"},{"type":"link","label":"Setup Server","href":"/siren/docs/tour/setup_server","docId":"tour/setup_server"},{"type":"link","label":"1 Sending On-demand Notification","href":"/siren/docs/tour/1sending_notifications_overview","docId":"tour/1sending_notifications_overview"},{"type":"link","label":"2 Alerting Rules and Subscription","href":"/siren/docs/tour/2alerting_rules_subscriptions_overview","docId":"tour/2alerting_rules_subscriptions_overview"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Concepts","items":[{"type":"link","label":"Overview","href":"/siren/docs/concepts/overview","docId":"concepts/overview"},{"type":"link","label":"Plugin","href":"/siren/docs/concepts/plugin","docId":"concepts/plugin"},{"type":"link","label":"Notification","href":"/siren/docs/concepts/notification","docId":"concepts/notification"},{"type":"link","label":"Glossary","href":"/siren/docs/concepts/glossary","docId":"concepts/glossary"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Guides","items":[{"type":"link","label":"Overview","href":"/siren/docs/guides/overview","docId":"guides/overview"},{"type":"link","label":"Server Installation","href":"/siren/docs/guides/deployment","docId":"guides/deployment"},{"type":"link","label":"Provider and Namespace","href":"/siren/docs/guides/provider_and_namespace","docId":"guides/provider_and_namespace"},{"type":"link","label":"Receiver","href":"/siren/docs/guides/receiver","docId":"guides/receiver"},{"type":"link","label":"Subscription","href":"/siren/docs/guides/subscription","docId":"guides/subscription"},{"type":"link","label":"Rule","href":"/siren/docs/guides/rule","docId":"guides/rule"},{"type":"link","label":"Template","href":"/siren/docs/guides/template","docId":"guides/template"},{"type":"link","label":"Alert History","href":"/siren/docs/guides/alert_history","docId":"guides/alert_history"},{"type":"link","label":"Notification","href":"/siren/docs/guides/notification","docId":"guides/notification"},{"type":"link","label":"Workers","href":"/siren/docs/guides/workers","docId":"guides/workers"},{"type":"link","label":"Job","href":"/siren/docs/guides/job","docId":"guides/job"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Providers","items":[{"type":"link","label":"CortexMetrics","href":"/siren/docs/providers/cortexmetrics","docId":"providers/cortexmetrics"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Receivers","items":[{"type":"link","label":"Slack","href":"/siren/docs/receivers/slack","docId":"receivers/slack"},{"type":"link","label":"PagerDuty","href":"/siren/docs/receivers/pagerduty","docId":"receivers/pagerduty"},{"type":"link","label":"HTTP","href":"/siren/docs/receivers/http","docId":"receivers/http"},{"type":"link","label":"File","href":"/siren/docs/receivers/file","docId":"receivers/file"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Reference","items":[{"type":"link","label":"Siren APIs","href":"/siren/docs/reference/api","docId":"reference/api"},{"type":"link","label":"Server Configuration","href":"/siren/docs/reference/server_configuration","docId":"reference/server_configuration"},{"type":"link","label":"Client Configuration","href":"/siren/docs/reference/client_configuration","docId":"reference/client_configuration"},{"type":"link","label":"CLI","href":"/siren/docs/reference/cli","docId":"reference/cli"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Extend","items":[{"type":"link","label":"Add a New Provider Plugin","href":"/siren/docs/extend/adding_new_provider","docId":"extend/adding_new_provider"},{"type":"link","label":"Add a New Receiver Plugin","href":"/siren/docs/extend/adding_new_receiver","docId":"extend/adding_new_receiver"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Contribute","items":[{"type":"link","label":"Contribution Process","href":"/siren/docs/contribute/contribution","docId":"contribute/contribution"},{"type":"link","label":"Release","href":"/siren/docs/contribute/release","docId":"contribute/release"}],"collapsed":true,"collapsible":true}]},"docs":{"concepts/glossary":{"id":"concepts/glossary","title":"Glossary","description":"Provider","sidebar":"docsSidebar"},"concepts/notification":{"id":"concepts/notification","title":"Notification","description":"Notification is one of main features in Siren. Siren capables to send notification to various receivers (Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren routes notification to specific receivers by matching notification key-value labels with the provided label matchers.","sidebar":"docsSidebar"},"concepts/overview":{"id":"concepts/overview","title":"Overview","description":"The following contains all the details about architecture and other technical concepts of Siren.","sidebar":"docsSidebar"},"concepts/plugin":{"id":"concepts/plugin","title":"Plugin","description":"Siren decouples provider and receiver as a plugin. The purpose is to ease the extension of new plugin. We welcome all contributions to add new plugin.","sidebar":"docsSidebar"},"contribute/contribution":{"id":"contribute/contribution","title":"Contribution Process","description":"The following is a set of guidelines for contributing to Siren. These are mostly guidelines, not rules. Use your best","sidebar":"docsSidebar"},"contribute/release":{"id":"contribute/release","title":"Release","description":"Siren release tags follow SEMVER convention.","sidebar":"docsSidebar"},"extend/adding_new_provider":{"id":"extend/adding_new_provider","title":"Add a New Provider Plugin","description":"Provider plugin is being used to update rule configurations to the provider, setup some runtime provider configuration config if any, and transform incoming alert data in Siren\'s Hook API to a list of *alert.Alert model. More details about the concept of provider plugin can be found here.","sidebar":"docsSidebar"},"extend/adding_new_receiver":{"id":"extend/adding_new_receiver","title":"Add a New Receiver Plugin","description":"Receiver plugin is being used to send notifications to the receiver. More details about the concept of receiver plugin can be found here.","sidebar":"docsSidebar"},"guides/alert_history":{"id":"guides/alert_history","title":"Alert History","description":"Siren can store the alerts triggered by provider e.g. Cortex Alertmanager. Provider needs to be configured to call Siren API using a webhook.","sidebar":"docsSidebar"},"guides/deployment":{"id":"guides/deployment","title":"Server Installation","description":"There are several approaches to setup Siren Server","sidebar":"docsSidebar"},"guides/job":{"id":"guides/job","title":"Job","description":"Job is a task that only runs once and then the process is terminated. Job could be scheduled with Cron or triggered manually. Siren is currently only having a one very-specific job.","sidebar":"docsSidebar"},"guides/notification":{"id":"guides/notification","title":"Notification","description":"To understand more concepts of notification in Siren, you can visit this page.","sidebar":"docsSidebar"},"guides/overview":{"id":"guides/overview","title":"Overview","description":"The following topics will describe how to use Siren.","sidebar":"docsSidebar"},"guides/provider_and_namespace":{"id":"guides/provider_and_namespace","title":"Provider and Namespace","description":"Siren provider represents a monitoring server. We define alerts and alerts routing configuration inside providers. For example the below provider, describes how a cortex monitoring server info is stored inside Siren.","sidebar":"docsSidebar"},"guides/receiver":{"id":"guides/receiver","title":"Receiver","description":"You can use receivers to send notifications on demand as well as on certain matching conditions (API for this is in the roadmap). Subscriptions use receivers to define routing configuration in Siren. With Siren subscriptions, incoming alerts via webhook will be routed to the pre-registered receivers by matching the subscriptions label. More info about notification concept is here. The how-to sending notification can be found here.","sidebar":"docsSidebar"},"guides/rule":{"id":"guides/rule","title":"Rule","description":"Siren rules are generated from predefined templates by providing values of the variables of the template.","sidebar":"docsSidebar"},"guides/subscription":{"id":"guides/subscription","title":"Subscription","description":"Siren lets you subscribe to a notification when they are triggered. You can define custom matching conditions and use","sidebar":"docsSidebar"},"guides/template":{"id":"guides/template","title":"Template","description":"Templates concept in Siren is used for abstraction. The usage is versatile enough to be used to abstract out rules and notification format. It utilises go-templates to provide data-driven templates for generating textual output. The template delimiter used is [[ and ]].","sidebar":"docsSidebar"},"guides/workers":{"id":"guides/workers","title":"Workers","description":"Siren has a notification features that utilizes queue to publish notification messages. More concept about notification could be found here. The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run","sidebar":"docsSidebar"},"installation":{"id":"installation","title":"Installation","description":"There are several approaches to install Siren CLI","sidebar":"docsSidebar"},"introduction":{"id":"introduction","title":"Introduction","description":"Welcome to the introductory guide to Siren! This guide is the best place to start with Siren. We cover what Siren is, what problems it can solve, how it works, and how you can get started using it. If you are familiar with the basics of Siren, the guides provides a more detailed reference of available features.","sidebar":"docsSidebar"},"providers/cortexmetrics":{"id":"providers/cortexmetrics","title":"CortexMetrics","description":"|||","sidebar":"docsSidebar"},"receivers/file":{"id":"receivers/file","title":"File","description":"|||","sidebar":"docsSidebar"},"receivers/http":{"id":"receivers/http","title":"HTTP","description":"|||","sidebar":"docsSidebar"},"receivers/pagerduty":{"id":"receivers/pagerduty","title":"PagerDuty","description":"|||","sidebar":"docsSidebar"},"receivers/slack":{"id":"receivers/slack","title":"Slack","description":"|||","sidebar":"docsSidebar"},"reference/api":{"id":"reference/api","title":"Siren APIs","description":"Documentation of our Siren API with gRPC and","sidebar":"docsSidebar"},"reference/cli":{"id":"reference/cli","title":"CLI","description":"siren alert","sidebar":"docsSidebar"},"reference/client_configuration":{"id":"reference/client_configuration","title":"Client Configuration","description":"When using siren client CLI, sometimes there are client-specifi flags that are required to be passed e.g. --host so you are calling Siren like this.","sidebar":"docsSidebar"},"reference/server_configuration":{"id":"reference/server_configuration","title":"Server Configuration","description":"Server configuration in siren is required to configure server, workers, and jobs. We can generate the default configuration with Siren CLI.","sidebar":"docsSidebar"},"support":{"id":"support","title":"Need help?","description":"Need a bit of help? We\'re here for you. Check out our current issues, GitHub discussions, or get support through Slack."},"tour/1sending_notifications_overview":{"id":"tour/1sending_notifications_overview","title":"1 Sending On-demand Notification","description":"This tour shows you how to send a notification to a receiver. You need to pick to which receiver you want send the notification to. If the receiver is not added in Siren yet, you could add one using siren receiver create. See receiver guide to explore more on how to work with siren receiver command.","sidebar":"docsSidebar"},"tour/2alerting_rules_subscriptions_overview":{"id":"tour/2alerting_rules_subscriptions_overview","title":"2 Alerting Rules and Subscription","description":"This tour shows you how could we create alerting rules and we want to subscribe to a notification triggered by an alert. If you want to know how to send on-demand notification to a receiver, you could go to the first tour.","sidebar":"docsSidebar"},"tour/introduction":{"id":"tour/introduction","title":"Introduction","description":"This tour introduces you to Siren. Along the way you will learn how to manage alerting rules, notification receivers, and subscribing to alert notifications.","sidebar":"docsSidebar"},"tour/setup_server":{"id":"tour/setup_server","title":"Setup Server","description":"Siren binary contains both the CLI client and the server itself. Each has it\'s own configuration in order to run. Server configuration contains information such as database credentials, log severity, etc. while CLI client configuration only has configuration about which server to connect.","sidebar":"docsSidebar"},"use_cases":{"id":"use_cases","title":"Use Cases","description":"As an Incident Management Platform, Siren integrates with several monitoring and alerting providers (CortexMetrics, Prometheus, InfluxDB, etc) and orchestrates alerting rules in a simple DIY configuration. Siren capables to subscribe to alerts and send notifications based on the triggered alerts or sending on-demand notifications to the supported receivers (slack, pagerduty, etc).","sidebar":"docsSidebar"}}}')}}]);