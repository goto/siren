# Introduction

Welcome to the introductory guide to Siren! This guide is the best place to start with Siren. We cover what Siren is, what problems it can solve, how it works, and how you can get started using it. If you are familiar with the basics of Siren, the [guides](./guides/overview.md) provides a more detailed reference of available features.

## What is Siren?

Siren is an Incident Management Platform that orchestrates alerting rules of your applications using a monitoring and alerting provider ([Cortex metrics](https://cortexmetrics.io/), InfluxDB, Prometheus) and sending notifications in a simple DIY configuration.

![Siren Overview](../static/img/siren_arch.svg)

 With Siren, you can define templates (using go templates standard), create/edit/enable/disable alerting rules on demand, and sending out notifications to various sinks (slack, pagerduty, e-mail, etc). Siren also gives flexibility to manage bulk of rules via YAML files so it can be integrated with any client such as CI/CD pipelines, self-serve UI, microservices, etc.

## The problem we aim to solve

 System observability gives us visibility on how the system behaves and monitoring is one part of it. Not only system, feature and some business level metrics are sometimes also important to observe. Failures always happen and they are unavoidable, the best we can do is to reduce it and immediately take an action out of it. Each failure that happened is essential to be learned from so everyone in the team aware of it, hoping that it does not recur in the future.
 
 There are several monitoring and alerting solution out there but each has its own approach and semantics and have some limitations when dealing with several scenarios (creating on-demand alerting rules, extending alert's notifications, etc). Some of them are also not that compatible to be used in multi-tenant-fashioned way.

 Siren provides an easy way to do end-to-end multi-tenancy aware incident management solution orchestrating alerting rule creation to some supported monitoring providers (cortexmetrics, prometheus, and influxDB), sending out notifications, and managing incidents.

## How does Siren work?

Here are the steps to work with Siren.

1. **Registering Provider:** Monitoring and alerting provider need to be registered to Siren by an administrator. 

2. **Adding Namespaces:** For different tenant, different namespace of provider needs to be created in Siren too.

3. **Configuring Alerting Rules:** Alerting rules for a specific namespace could be created by a user and Siren will synchronize the configured rules in the selected monitoring providers.

4. **Registering Receivers:** User needs to add receivers (e.g. slack, pagerduty) for Siren to send notifications to.

5. **Sending Notifications:**
    - **Subscribing Alerts:** User could set up an alert subscription to be notified if an alert is triggered. User needs to create a subscription to subscribe to the alerts notification by label-matching.

    - **Sending On-demand Notification:** Siren could be used to directly send notification to the registered receivers.

## Key Features

- **Rule Templates:** Siren provides a way to define templates over alerting rule which can be reused to create multiple instances of the same rule with configurable thresholds.

- **Subscriptions:** Siren can be used to subscribe to notifications (with desired matching conditions) via the channel of your choice.

- **Multi-tenancy:** Rules created with Siren are by default multi-tenancy aware.

- **DIY Interface:** Siren can be used to easily create/edit alerting rules. It also provides soft-delete (disable) so that you can preserve thresholds in case you need to reuse the same alert.

- **Managing bulk rules:** Siren enables users to manage bulk alerting rules using YAML files in specified format with simple CLI.

- **Receivers:** Siren can be used to send out notifications to several channels (slack, pagerduty, email etc).

- **Alert History:** Siren can store alerts triggered by monitoring & alerting provider e.g. Cortex Alertmanager, which can be used for audit purposes.

## Using Siren

You can manage alerting rules, subscribe to, and sending a notification in any of the following ways:

### Siren Command Line Interface
You can use the Siren command line interface to issue commands and to perform the entire Siren features. Using the command line can be faster and more convenient than using API. For more information on using the Siren CLI, see the [CLI Reference](./reference/cli.md) page.

### HTTPS API
You can get hands on rule configuration, sending notification, notification subscription and much more by using the Siren HTTPS API, which lets you issue HTTPS requests directly to the service. For more information, see the [API Reference page](./apis/siren-apis.info.mdx).

## Where to go from here

See the [installation](./installation.md) page to install the Siren CLI. Next, we recommend completing the [guides](./guides/overview.md). The [tour](./tour/introduction.md) provides an overview of most of the existing functionality of Siren and takes approximately 20 minutes to complete.

After completing the tour, check out the remainder of the documentation in the [reference](./reference/server_configuration.md) and [concepts](./concepts/overview.md) sections for your specific areas of interest. We've aimed to provide as much documentation as we can for the various components of Siren to give you a full understanding of Siren's surface area. If you are interested to contribute, check out the [contribution](./contribute/contribution.md) page.
