# heyemoji 🏆 👏 🔥 
# 

[![Build](https://img.shields.io/docker/build/mmcdole/heyemoji)](https://dockerhub.com/mmcdole/heyemoji) [![Go Report Card](https://goreportcard.com/badge/github.com/mmcdole/heyemoji)](https://goreportcard.com/report/github.com/mmcdole/heyemoji) [![Doc](https://godoc.org/github.com/mmcdole/heyemoji?status.svg)](http://godoc.org/github.com/mmcdole/heyemoji) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

The `heyemoji` app is a slack reward system that allows team members to recognize eachother for anything awesome they may have done.  This is done by mentioning a user's slack **@username** in a channel along with a pre-configured reward **emoji** and a brief **description** of what they did.  The emoji points bestowed to users can be tracked via leaderboards.

## Table of Contents

- [Overview](#overview)
- [Configuration](#configuration)
- [Docker](#docker)

## Overview

// TODO

## Configuration

| ENV Var         | Default  | Required | Note                                                          |   |
|-----------------|----------|----------|---------------------------------------------------------------|---|
| HEY_BOT_NAME        | heyemoji | No       | The display name of the heyemoji bot                          |   |
| HEY_DATABASE_PATH   | ./data/  | No       | The directory that the database files should be written to    |   |
| HEY_SLACK_API_TOKEN |          | Yes      | The API tokens for the Slack API                              |   |
| HEY_SLACK_EMOJI     | star:1   | No       | Comma delimited set of emoji "name:value" pairs               |   |
| HEY_SLACK_DAILY_CAP | 5        | No       | The max number of emoji points that can be given out in a day |   |
| HEY_WEBSOCKET_PORT  | 3334     | No       | Port that the Slack RTM client will listen on                 |   |

## Docker


