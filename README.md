# Cloud PubSub to Stackdriver Logging export
Export text or JSON data from Cloud PubSub to Stackdriver Logging using Cloud Functions. This is useful when message volume is relatively low and there is a need to keep the complexity to a minimum.

## Why Stackdriver Logging?

Having data in Stackdriver Logging allows you to:
* Perform basic queries on structured data using `Advanced Filter`
* Set up logging sinks to BigQuery, Cloud Storage and PubSub
* Use log data to create custom metrics for Stackdriver Monitoring

## Potential use cases
* Exporting data from Google services that do not support native export to BigQuery or GCS.
* Third party integration.

## Approximate cost
As mentioned earlier, this is cost-effective when data volume is low.

Below are some estimates data volume is 1 million 1 kB messages a month.
* Cloud Functions: 1 mil invocations x 5 seconds each - `$8.22`
* Cloud PubSub: 1GB of message data - `free` ($40.00/TB after free tier)
* Stackdriver Logging: 5GB of log data - `free` ($0.50/GB after free tier)