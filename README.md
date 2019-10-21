
![](https://github.com/akiicat/short-url/workflows/Deploy to Google Cloud Functions/badge.svg)

# Short Url on Google Cloud Function

## Databse

1. Enable Cloud Firestore on Google Cloud Function
2. Connect the project to the Firebase
3. Add the data to the database, and replace the word in [].
  - Collectons: **links**
  - Document: **[SHORT_NAME]**
  - Column:
    - **link**: **[REDIRECT_URL]**

## Deploy command

```sh
gcloud functions deploy api --entry-point=Link --runtime=go111 --trigger-http
```

