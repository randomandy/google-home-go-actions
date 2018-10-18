# google-home-go-actions
Some experiments with Google Home Actions in Go

## DialogFlow API v1

### Setup Google Cloud project

- Create a Google Cloud project (https://console.cloud.google.com)
- Install Google Cloud CLI (https://cloud.google.com/sdk/install)
- run `gcloud init` to link your google account and set the current project

### Deployment

You can deploy your code from inside your project directory with `gcloud app deploy`

### Actions on Google & Dialogflow

A very good basic introduction can be found here: https://codelabs.developers.google.com/codelabs/actions-1/index.html#0

#### Setup your project

- Open (https://console.actions.google.com) and import your Google Cloud project
- Skip the Category selection for now
- Click `Actions` in the `Build` section of the left side navigation
- On the `Create Action` dialog, select `Custom Intent` and click `Build`
- You're now being redirected to Dialogflow where you can confirm the project creation
