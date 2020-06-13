# Serverless backend showcase

Using Cloud Functions in Go and Cloud Pub/Sub for fan-in/out of messages 
to connect the following services in a showcase:

- Cloud Translation
- Firebase Realtime Database
- BigQuery
- Cloud Storage

This schema shows the structure of the backend.
     
![Schema](schema.png)
    
#### DevOps Pipelines 

Additionally, we provide DevOps pipelines to build, test, and deploy the Cloud functions using Cloud Build. Therefore, 
we connect our GitHub repository with a GCP Source Repository. We use standard cloud-builder images, and build our own 
if needed stored in our GCP Container Registry. The creation of these build trigger is half-automated by templates.

#### Wireformat

We send the messages in JSON format in between the services. 

```bash
{
  "clientVersion": "0.0.1",
  "clientId": "beab10c6-deee-4843-9757-719566214526",
  "taskId": "31427acf-42e6-4981-82d1-0abd2c5c2abe",
  "text": "Today is Wednesday",
  "sourceLanguage": "en",
  "targetLanguage": "fr",
  "translatedText": "Aujourd'hui nous sommes mercredi"
}
```

Frontend messages are in JSON format as well.

Request:

```bash
{
  "clientVersion": "0.0.1",
  "clientId": "beab10c6-deee-4843-9757-719566214526",
  "text": "Today is Wednesday",
  "sourceLanguage": "en",
  "targetLanguage": "fr"
}
```

Response:

```bash
{
 "taskId": "4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5",
 "translatedText": "Aujourd'hui nous sommes mercredi",
 "loadCommands": [
  "firebase database:get --pretty --instance migros-showcase --project hybrid-cloud-22365 /translations_v0_0_1/beab10c6-deee-4843-9757-719566214526/4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5",
  "bq query 'SELECT * FROM migros_showcase.translations_v0_0_1 WHERE taskId = \"4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5\"'",
  "gsutil cat gs://hybrid-cloud-22365.appspot.com/0.0.1/beab10c6-deee-4843-9757-719566214526/4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5 | jq",
  "gcloud logging read 'resource.type=cloud_function resource.labels.region=europe-west1 textPayload=4d71c2a3-e6e1-4efd-bb1f-082227cfb0a5'"
 ]
}
```

The response contains a list of commands to access services via CLI.

#### Versioning and A/B Testing

We use versioning for our microservices included in the messages to deploy different scenarios within the backend, not 
disturbing others, and useful for A/B testing, etc.

#### Security

We follow the principle of least privilege, providing roles to the service account 
for each active Cloud Function.
