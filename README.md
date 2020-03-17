
![](https://github.com/akiirobot/short-url/workflows/Deploy%20to%20Google%20Cloud%20Functions/badge.svg)

# Short Url on Google Cloud Function

## Databse

1. Enable Cloud Firestore on Google Cloud Function
2. Connect the project to the Firebase
3. Add the data to the database, and replace the word in [].
  - Collectons: **links**
  - Document: **[SHORT_NAME]**
  - Column:
    - **link**: **[REDIRECT_URL]**

## Setup Env

```shell
echo "PROJECT_ID: $PROJECT_ID" >> .env
```

## Deploy command

```sh
cd link
go mod vendor
cd ..
gcloud functions deploy link --source=link --entry-point=Link --runtime=go113 --trigger-http --quiet --env-vars-file .env
```

## Deploy firebase routes

Edit **firebase.json**

```shell
firebase use $PROJECT_ID
firebase deploy --only hosting
```

## Customism Domain Name

1. Go to Project in Firebase 
2. Select the **Hosting**
3. Create a new domain
  - Enter your domain
  - And Redirect your domain to current web site

