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


#### Security

We follow the principle of least privilege, providing roles to the service account 
for each active Cloud Function.

#### Versioning and A/B Testing

We use versioning for our microservices included in the messages to deploy different scenarios within the backend, not 
disturbing others, and useful for A/B testing, etc.
