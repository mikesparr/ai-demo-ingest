# AI Demo API (batch note review)
This API accepts user submission of one or more bank notes for analysis. After submission, 
backend processors perform prediction and publish results. The results can then be viewed 
using the [predict API](https://github.com/mikesparr/ai-demo-predict). This is created for
demo purposes using [go-chi](https://github.com/go-chi/chi) HTTP framework for Golang.

# Demo (nothing too sexy)
![API Demo](./demo.gif)

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
Once deployed, you can fetch the `$API_URL` from Cloud Run and `POST` data to the API. Since it just publishes data to a Pub/Sub topic, you will just receive either the submitted record, or error message.

```bash
# get URL of service
export API_URL=$(gcloud run services describe ai-demo-ingest --format="value(status.url)" --platform managed --region $GCP_REGION)

# test the API
curl -XPOST -H "Content-type: application/json" \
    $API_URL/notes \
    -d '{"subjects": ["abc123"], "features": [[0.2234,1.2342,-1.3243,-0.9383]]}'                           
```

# Spec
See the OAS2/Swagger specification `config.yaml` for more details

# Contributing
This is just a demo so fork and use at your own discretion.