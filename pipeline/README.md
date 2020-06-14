# DevOps Pipelines

Our approach for DevOps pipelines uses Cloud Build triggers with configurations in YAML files to prepare, deploy, and 
test Cloud Functions. It uses Cloud Storage buckets with versioning and lifecycle management as deployment sources. 
Pros are immutable build-environments complete with all dependencies, for fast backup and recovery. 

Additionally, we provide a simple version functionality. I.e., the basename of the YAML file and the bash script are 
identical with an empty file's basename. This empty file has an ending showing the demanded version. 
Then, the bash script finds the corresponding directory containing the appropriate function definition. 
We prepare the directory, zip it, and store it in a bucket for another build trigger to deploy the function.

Currently, we have no trigger calling other triggers to present longer pipelines with parallel executions.

For convenience and reproducibility, we use trigger configurations in JSON format. These files have the same basename 
as its related files and the ending "trigger." 




