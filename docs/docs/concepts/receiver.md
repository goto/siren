# Receiver

## Parent vs Child Receiver

In Siren, one receiver could depends to other receiver and inherit its configuration. However it only supports one level parent-child relation.The receiver that depends on other receiver is called a `Child Receiver` but the receiver that other receiver depends on is called a `Parent Receiver`.

There is never an orphan `Child Receiver` in Siren. That means, a `Child Receiver` always has to have a `Parent Receiver`. However, not every receiver could be a `Child Receiver`. Only a specific receiver type that could be a `Child Receiver`. 

Currently the only receiver type that is a `Child Receiver` is:
- Slack Channel receiver, that depends on Slack receiver

The Slack Channel receiver has config
```
{
    "channel_name": "xxx"
}
```
While the Slack receiver has config
```
{
    "token": "1234",
    "workspace": "gotocompany"
}
```

When rendering a message, if a message is a child receiver, Siren would merge the configs of parent receiver to the child receiver.