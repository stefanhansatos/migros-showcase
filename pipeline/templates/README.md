Create new triggers using templates here

e.g.

- Replace in `create-trigger.yaml` the trigger name `FuncToBeExecuted` with the name of the new trigger
- edit the trigger configuration in `cf-trigger-template.json` 
- upload the changes to the repository
- execute the trigger `gcloud alpha builds triggers run gcloud-build-triggers --branch=master`
- redo the changes in the files

Next copy `cf-trigger-template.yaml` to the YAML file specified in the trigger configuration and edit manually.


