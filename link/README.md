
```shell
gcloud functions deploy link --source=link --entry-point=Link --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding link --member=allUsers --role=roles/cloudfunctions.invoker
```
