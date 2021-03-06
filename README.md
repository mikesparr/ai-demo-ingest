[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mikesparr_ai-demo-ingest&metric=alert_status)](https://sonarcloud.io/dashboard?id=mikesparr_ai-demo-ingest)

This API accepts user submission of one or more bank notes for analysis. After submission, 
backend processors perform prediction and publish results. The results can then be viewed 
using the [predict API](https://github.com/mikesparr/ai-demo-predict). This is created for
demo purposes using [go-chi](https://github.com/go-chi/chi) HTTP framework for Golang.

# Demo (nothing too sexy)
![API Demo](./img_demo.gif)

# Architecture
![AI demo architecture](./img_arch.png)

# Components
- [Config](https://#) (pending)
- [Web App](https://github.com/mikesparr/ai-demo-web)
- [Ingest API](https://github.com/mikesparr/ai-demo-ingest) (this repo)
- [Predict API](https://github.com/mikesparr/ai-demo-predict)
- [Processors](https://github.com/mikesparr/ai-demo-functions)

# Prerequisites
You must be familiar with Google Cloud Platform and have the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) (`gcloud` CLI) installed. 
If you want to recreate the **AI Demo** then you will want an active project with billing enabled.

* NOTE: when you are done remember to **DELETE THE PROJECT** to avoid unneccessary billing.

# Install
The demo publishes a Docker image to container registry and deploys the app to Cloud Run. 
These are the steps to recreate this in your own environment.

```bash
export PROJECT_ID=$(gcloud config get-value project)
export TOPIC_ID="request"
export GCP_REGION="us-central1"

# create pubsub topics
gcloud pubsub topics create $TOPIC_ID

# create pubsub subscriptions
gcloud pubsub subscriptions create request-sub --topic $TOPIC_ID

# enable services
gcloud services enable compute.googleapis.com \
    run.googleapis.com \
    cloudbuild.googleapis.com

# clone repo and change to directory
git clone git@github.com:mikesparr/ai-demo-ingest.git
cd ai-demo-ingest

# build the api image
gcloud builds submit --tag gcr.io/${PROJECT_ID}/ai-demo-ingest

# deploy the api to cloud run
gcloud run deploy ai-demo-ingest \
    --image gcr.io/${PROJECT_ID}/ai-demo-ingest \
    --region $GCP_REGION \
    --allow-unauthenticated \
    --platform managed \
    --update-env-vars PROJECT_ID=$PROJECT_ID,TOPIC_ID=$TOPIC_ID
```

# Usage
Once deployed, you can fetch the `$INGEST_URL` from Cloud Run and `POST` data to the API. Since it just publishes data to a Pub/Sub topic, you will just receive either the submitted record, or error message.

```bash
# get URL of service
export INGEST_URL=$(gcloud run services describe ai-demo-ingest --format="value(status.url)" --platform managed --region $GCP_REGION)

# test the API
curl -XPOST -H "Content-type: application/json" \
    $INGEST_URL/notes \
    -d '{"subjects": ["test-record"], "features": [[0.2234,1.2342,-1.3243,-0.9383]]}'                       
```

# Spec
See the OAS2/Swagger specification `config.yaml` for more details

# Validation
In an attempt to minimize *"garbage in"* but make the API user-friendly, adding thorough input checks to the `{model}.Bind()` with useful error message responses.
![AI demo architecture](./img_validation.png)

# Other considerations
Although this is only a demo, a few additional features that should be added would be:
- automated tests
- standardized Request / Response
- retry logic with exponential backoff
- tracing, metrics using [opentelemetry](https://opentelemetry.io/)
- messages using [cloudevents](https://cloudevents.io/)

# Contributing
This is just a demo so fork and use at your own discretion.