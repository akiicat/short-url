
![](https://github.com/akiicat/short-url/workflows/Deploy%20to%20Google%20Cloud%20Functions/badge.svg)

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
gcloud functions deploy api --entry-point=Link --runtime=go113 --trigger-http
```

## Customism Domain Name

1. Go to Project in Firebase 
2. Select the **Hosting**
3. Create a new domain
  - Enter your domain
  - And Redirect your domain to current web site

