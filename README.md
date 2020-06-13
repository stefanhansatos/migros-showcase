# Serverless backend showcase around GCP, Firebase, and more

Using Cloud Functions in Go and Cloud Pub/Sub for fan-in/out of messages 
to connect the following services in a showcase:

- Cloud Translation
- Firebase Realtime Database
- BigQuery
- Cloud Storage

This schema shows the structure of the backend.
     
![Schema](schema.png)     

Additionally, we provide DevOps pipelines to build, test, and deploy the Cloud functions using Cloud Build. Therefore, 
we connect our GitHub repository with a GCP Source Repository. We use standard cloud-builder images, and build our own 
if needed stored in our GCP Container Registry. The creation of these build trigger is half-automated by templates.


#### Development

We develop in the working directory `functions` and push the `vendor` directory to the Github repository 
and to Cloud Source Repository, respectively.


