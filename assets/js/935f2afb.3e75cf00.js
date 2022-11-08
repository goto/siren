"use strict";(self.webpackChunksiren=self.webpackChunksiren||[]).push([[53],{1109:e=>{e.exports=JSON.parse('{"pluginId":"default","version":"current","label":"Next","banner":null,"badge":false,"noIndex":false,"className":"docs-version-current","isLast":true,"docsSidebars":{"docsSidebar":[{"type":"link","label":"Introduction","href":"/siren/docs/introduction","docId":"introduction"},{"type":"link","label":"Installation","href":"/siren/docs/installation","docId":"installation"},{"type":"category","label":"Tour","items":[{"type":"link","label":"Introduction","href":"/siren/docs/tour/overview","docId":"tour/overview"},{"type":"link","label":"Start Server","href":"/siren/docs/tour/start_server","docId":"tour/start_server"},{"type":"link","label":"Register provider","href":"/siren/docs/tour/registering_provider","docId":"tour/registering_provider"},{"type":"link","label":"Register receivers","href":"/siren/docs/tour/registering_receivers","docId":"tour/registering_receivers"},{"type":"link","label":"4 - Sending Notification to Receiver","href":"/siren/docs/tour/sending_notifications_to_receiver","docId":"tour/sending_notifications_to_receiver"},{"type":"link","label":"5 - Configuring Provider Alerting Rules","href":"/siren/docs/tour/configuring_provider_alerting_rules","docId":"tour/configuring_provider_alerting_rules"},{"type":"link","label":"6 - Subscribing Notifications","href":"/siren/docs/tour/subscribing_notifications","docId":"tour/subscribing_notifications"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Concepts","items":[{"type":"link","label":"Overview","href":"/siren/docs/concepts/overview","docId":"concepts/overview"},{"type":"link","label":"Plugin","href":"/siren/docs/concepts/plugin","docId":"concepts/plugin"},{"type":"link","label":"Schema Design","href":"/siren/docs/concepts/schema","docId":"concepts/schema"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Guides","items":[{"type":"link","label":"Usage","href":"/siren/docs/guides/overview","docId":"guides/overview"},{"type":"link","label":"Provider and Namespace","href":"/siren/docs/guides/provider_and_namespace","docId":"guides/provider_and_namespace"},{"type":"link","label":"Receiver","href":"/siren/docs/guides/receiver","docId":"guides/receiver"},{"type":"link","label":"Subscription","href":"/siren/docs/guides/subscription","docId":"guides/subscription"},{"type":"link","label":"Rule","href":"/siren/docs/guides/rule","docId":"guides/rule"},{"type":"link","label":"Template","href":"/siren/docs/guides/template","docId":"guides/template"},{"type":"link","label":"Alert History","href":"/siren/docs/guides/alert_history","docId":"guides/alert_history"},{"type":"link","label":"Notification","href":"/siren/docs/guides/notification","docId":"guides/notification"},{"type":"link","label":"Deployment","href":"/siren/docs/guides/deployment","docId":"guides/deployment"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Contribute","items":[{"type":"link","label":"Contribution Process","href":"/siren/docs/contribute/contribution","docId":"contribute/contribution"},{"type":"link","label":"Add a New Receiver Plugin","href":"/siren/docs/contribute/receiver","docId":"contribute/receiver"},{"type":"link","label":"Add a New Provider Plugin","href":"/siren/docs/contribute/provider","docId":"contribute/provider"},{"type":"link","label":"Release","href":"/siren/docs/contribute/release","docId":"contribute/release"}],"collapsed":true,"collapsible":true},{"type":"category","label":"Reference","items":[{"type":"link","label":"Siren APIs","href":"/siren/docs/reference/api","docId":"reference/api"},{"type":"link","label":"Server Configuration","href":"/siren/docs/reference/server_configuration","docId":"reference/server_configuration"},{"type":"link","label":"Client Configuration","href":"/siren/docs/reference/client_configuration","docId":"reference/client_configuration"},{"type":"link","label":"Receiver","href":"/siren/docs/reference/receiver","docId":"reference/receiver"},{"type":"link","label":"CLI","href":"/siren/docs/reference/cli","docId":"reference/cli"}],"collapsed":true,"collapsible":true}]},"docs":{"concepts/glossary":{"id":"concepts/glossary","title":"Glossary","description":"Provider"},"concepts/notification":{"id":"concepts/notification","title":"Notification","description":"Notification is one of main features in Siren. Siren capables to send notification to various receivers (e.g. Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren routes notification to specific receivers by matching notification key-value labels with the provided label matchers."},"concepts/overview":{"id":"concepts/overview","title":"Overview","description":"The following contains all the details about architecture, database schema, code structure and other technical concepts of Siren.","sidebar":"docsSidebar"},"concepts/plugin":{"id":"concepts/plugin","title":"Plugin","description":"Siren decouples provider and receiver as a plugin. The purpose is to ease the extension of new plugin. We welcome all contributions to add new plugin.","sidebar":"docsSidebar"},"concepts/schema":{"id":"concepts/schema","title":"Schema Design","description":"Siren stores providers, namespaces, templates, rules and triggered alerts history, receivers and subscriptions in","sidebar":"docsSidebar"},"contribute/contribution":{"id":"contribute/contribution","title":"Contribution Process","description":"The following is a set of guidelines for contributing to Siren. These are mostly guidelines, not rules. Use your best","sidebar":"docsSidebar"},"contribute/provider":{"id":"contribute/provider","title":"Add a New Provider Plugin","description":"Provider plugin is being used to update rule configurations to the provider, setup some runtime provider configuration config if any, and transform incoming alert data in Siren\'s Hook API to list of *alert.Alert model.","sidebar":"docsSidebar"},"contribute/receiver":{"id":"contribute/receiver","title":"Add a New Receiver Plugin","description":"More details about the concept of receiver plugin can be found here.","sidebar":"docsSidebar"},"contribute/release":{"id":"contribute/release","title":"Release","description":"Siren release tags follow SEMVER convention.","sidebar":"docsSidebar"},"guides/alert_history":{"id":"guides/alert_history","title":"Alert History","description":"Siren can store the alerts triggered by provider e.g. Cortex Alertmanager. Provider needs to be configured to call Siren API using a webhook.","sidebar":"docsSidebar"},"guides/deployment":{"id":"guides/deployment","title":"Deployment","description":"Siren docker image can be found on Docker hub here. You can run the image with","sidebar":"docsSidebar"},"guides/notification":{"id":"guides/notification","title":"Notification","description":"Notification is one of main features in Siren. Siren capables to send notification to various receivers (e.g. Slack, PagerDuty). Notification in Siren could be sent directly to a receiver or user could subscribe notifications by providing key-value label matchers. For the latter, Siren routes notification to specific receivers by matching notification key-value labels with the provided label matchers.","sidebar":"docsSidebar"},"guides/overview":{"id":"guides/overview","title":"Usage","description":"The following topics will describe how to use Siren.","sidebar":"docsSidebar"},"guides/provider_and_namespace":{"id":"guides/provider_and_namespace","title":"Provider and Namespace","description":"Siren providers represent a monitoring server. We define alerts and alerts routing configuration inside providers. For example the below provider, describes how a cortex monitoring server info is stored inside Siren.","sidebar":"docsSidebar"},"guides/receiver":{"id":"guides/receiver","title":"Receiver","description":"You can use receivers to send notifications on demand as well as on certain matching conditions (API for this is in the roadmap). Subscriptions use receivers to define routing configuration in Siren. With Siren subscriptions, incoming alerts via webhook will be routed to the pre-registered receivers by matching the subscriptions label. More info about notification concept is here. The how-to sending notification can be found here.","sidebar":"docsSidebar"},"guides/rule":{"id":"guides/rule","title":"Rule","description":"Siren rules are generated from predefined templates by providing values of the variables of the template.","sidebar":"docsSidebar"},"guides/subscription":{"id":"guides/subscription","title":"Subscription","description":"Siren lets you subscribe to a notification when they are triggered. You can define custom matching conditions and use","sidebar":"docsSidebar"},"guides/template":{"id":"guides/template","title":"Template","description":"Templates concept in Siren is used for abstraction. The usage is versatile enough to be used to abstract out rules and notification format. It utilises go-templates to provide data-driven templates for generating textual output. The template delimiter used is [[ and ]].","sidebar":"docsSidebar"},"guides/worker":{"id":"guides/worker","title":"Workers","description":"Siren has a notification features that utilizes queue to publish notification messages. The architecture requires a detached worker running asynchronously and polling queue periodically to dequeue notification messages and publish them. By default, Siren server run this asynchronous worker inside it. However it is also possible to run the worker as a different process. Currently there are two possible workers to run"},"installation":{"id":"installation","title":"Installation","description":"There are several approaches to install Siren CLI","sidebar":"docsSidebar"},"introduction":{"id":"introduction","title":"Introduction","description":"Siren orchestrates alerting rules of your applications using a monitoring and alerting provider e.g. Cortex metrics and sending notifications in a simple DIY configuration. With Siren, you can define templates (using go templates standard), create/edit/enable/disable alerting rules on demand, and sending out notifications. It also gives flexibility to manage bulk of rules via YAML files. Siren can be integrated with any client such as CI/CD pipelines, Self-Serve UI, microservices etc.","sidebar":"docsSidebar"},"reference/api":{"id":"reference/api","title":"Siren APIs","description":"Documentation of our Siren API with gRPC and","sidebar":"docsSidebar"},"reference/cli":{"id":"reference/cli","title":"CLI","description":"siren alert","sidebar":"docsSidebar"},"reference/client_configuration":{"id":"reference/client_configuration","title":"Client Configuration","description":"When using siren client CLI, sometimes there are client-specifi flags that are required to be passed e.g. --host so you are calling Siren like this.","sidebar":"docsSidebar"},"reference/receiver":{"id":"reference/receiver","title":"Receiver","description":"Receiver represents a notification medium, which can be used to define notification routing configuration in Siren. Siren currently supports several type of receivers. Here is receiver schema. The type of stored receiver in the DB is immutable.","sidebar":"docsSidebar"},"reference/server_configuration":{"id":"reference/server_configuration","title":"Server Configuration","description":"Server configuration in siren is required to configure server, workers, and jobs. We can generate the default configuration with Siren CLI.","sidebar":"docsSidebar"},"support":{"id":"support","title":"Need help?","description":"Need a bit of help? We\'re here for you. Check out our current issues, GitHub discussions, or get support through Slack."},"tour/configuring_provider_alerting_rules":{"id":"tour/configuring_provider_alerting_rules","title":"5 - Configuring Provider Alerting Rules","description":"In this part we will create alerting rules for our Cortex monitoring provider. Rules in Siren relies on template for its abstraction. We need to create a rule\'s template first before uploading a rule.","sidebar":"docsSidebar"},"tour/introduction":{"id":"tour/introduction","title":"introduction","description":"Tour Introduction"},"tour/overview":{"id":"tour/overview","title":"Introduction","description":"This tour introduces you to how to use Siren to manage alerting rules and send notifications. We will integrate Siren with a monitoring provider, configuring alerting rules for the provider, accepting triggered alerts from the provider, sending notification from the triggered alerts, and sending notification on-demand. The tour takes approximately 20 minutes to complete.","sidebar":"docsSidebar"},"tour/registering_provider":{"id":"tour/registering_provider","title":"Register provider","description":"1. Register the provider","sidebar":"docsSidebar"},"tour/registering_receivers":{"id":"tour/registering_receivers","title":"Register receivers","description":"1. Register a receiver","sidebar":"docsSidebar"},"tour/sending_notifications_to_receiver":{"id":"tour/sending_notifications_to_receiver","title":"4 - Sending Notification to Receiver","description":"In previous part, we have already registered several receivers and got back the receiver IDs. We could send a notification to the receivers with /receivers/:receiverId/send API. We can use Siren CLI to do this.","sidebar":"docsSidebar"},"tour/start_server":{"id":"tour/start_server","title":"Start Server","description":"1. Build Siren","sidebar":"docsSidebar"},"tour/subscribing_notifications":{"id":"tour/subscribing_notifications","title":"6 - Subscribing Notifications","description":"Notifications can be subscribed and routed to the defined receivers by adding a subscription. In this part, we will simulate how Cortex Ruler trigger an alert to Cortex Alertmanager, and Cortex Alertmanager trigger webhook-notification and calling Siren alerts hook API. On Siren side, we expect a notification is published everytime the hook API is being called.","sidebar":"docsSidebar"}}}')}}]);