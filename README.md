# heyemoji 🏆 👏 ⭐
# 

![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/mmcdole/heyemoji) [![Go Report Card](https://goreportcard.com/badge/github.com/mmcdole/heyemoji)](https://goreportcard.com/report/github.com/mmcdole/heyemoji) [![Doc](https://godoc.org/github.com/mmcdole/heyemoji?status.svg)](http://godoc.org/github.com/mmcdole/heyemoji) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

The `heyemoji` bot is a self-hosted slack reward system that allows team members to recognize eachother for anything awesome they may have done.  This is accomplished by mentioning a user's slack **@username** in a channel along with a pre-configured **reward emoji** and an optional **description** of what they did.  The emoji points bestowed to users can be tracked via leaderboards.

## Table of Contents

- [Usage](#basic-usage)
- [Setup](#setup)
- [Configuration](#configuration)

## Basic Usage

#### Give a single user emoji points 

`Great job filling out those TPS reports @michael.bolton! ⭐` 

#### Give multiple users emoji points

`Thanks @michael.bolton and @samir for coming in this weekend! ⭐`

#### Give multiple emojis to a user

`Score! @petergibbons found my red stapler! ⭐ 🏆 👏 👏 `

### Bot Commands

| Name                   | Description                                                |
|------------------------------------------|------------------------------------------------------------|
| `leaderboard <day\|week\|month\|year>`   | see the top 10 people on your leaderboard                  |
| `points`                                 | see how many emoji points you have left to give            |
| `help`                                   | get help with how to send recognition emoji                |

## Setup

1. Browse to the Slack App Console and [Create a Classic App](https://api.slack.com/apps?new_classic_app=1)
1. Assign a name and workspace to your new Slack Bot Application
1. `Basic Information` > Set display name and icon
1. `App Home` > Add Legacy Bot User
1. `OAuth & Permissions` > Install App to Workspace
1. Copy your **Bot User OAuth Access Token** for your `HEY_SLACK_TOKEN`
1. Create a `.env` file in the HeyEmoji root folder and add `HEY_SLACK_TOKEN=<Bot User OAuth Access Token>`
1. Run `docker-compose up`! 🎉

## Configuration

| ENV Var             | Default  | Required | Note                                                          |
|---------------------|----------|----------|---------------------------------------------------------------|
| HEY_BOT_NAME        | heyemoji | No       | The display name of the heyemoji bot                          |
| HEY_DATABASE_PATH   | ./data/  | No       | The directory that the database files should be written to    |
| HEY_SLACK_TOKEN     |          | Yes      | The API tokens for the Slack API                              |
| HEY_SLACK_EMOJI     | star:1   | No       | Comma delimited set of emoji "name:value" pairs               |
| HEY_SLACK_DAILY_CAP | 5        | No       | The max number of emoji points that can be given out in a day |
| HEY_WEBSOCKET_PORT  | 3334     | No       | Port that the Slack RTM client will listen on                 |


### Specifying Custom Reward Emoji

The `HEY_SLACK_EMOJI` setting lets you specify multiple different reward emoji as well as different point values for each. So, if you wanted the following emoji and reward values:

| Emoji         | Value  |
|---------------|--------|
| ⭐             | 1      |
| 👏             | 2      |
| 🏆             | 3      |

You would specify the `HEY_SLACK_EMOJI` as: `star:1,clap:2,trophy:3`

## Deploy to AWS
Deploying `heyemoji` to AWS is simple! The repo already contains the Github Action: [Deploy HeyEmoji](.github/workflows/deploy.yaml) which utilizes the [bitovi/github-actions-node-app-to-aws-vm](https://github.com/bitovi/github-actions-node-app-to-aws-vm) action to create an EC2 instance and deploy an application to it using [BitOps](https://github.com/bitovi/bitops).

To deploy to AWS;
1. Navigate to `Github Repo Settings` > Secrets > Actions
2. Within the `Repository secrets ` scope (minimum) add the following values;
    - AWS_ACCESS_KEY_ID > retrieved from the AWS IAM console 
    - AWS_SECRET_ACCESS_KEY > retrieved from the AWS IAM console 
    - DOT_ENV > Values passed to HeyEmoji. Use `key=value` separated by a newline


**Example DOT_ENV Value**
```
HEY_SLACK_TOKEN=xoxb-12345-12345-a1b2c3d4
HEY_SLACK_EMOJI=egg:1,hatching_chick:2,hatched_chick:3
```