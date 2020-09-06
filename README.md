# heyemoji
# 

[![Build Status](https://travis-ci.org/mmcdole/heyemoji.svg?branch=master)](https://travis-ci.org/mmcdole/heyemoji) [![Coverage Status](https://coveralls.io/repos/github/mmcdole/heyemoji/badge.svg?branch=master)](https://coveralls.io/github/mmcdole/heyemoji?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/mmcdole/heyemoji)](https://goreportcard.com/report/github.com/mmcdole/heyemoji) [![](https://godoc.org/github.com/mmcdole/heyemoji?status.svg)](http://godoc.org/github.com/mmcdole/heyemoji) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

The `heyemoji` app is a slack reward system that allows team members to recognize eachother for the work they do.  This is done by mentioning a user's slack name in a channel along with a pre-configured set of reward emoji.

## Table of Contents

- [Overview](#overview)
- [Configuration](#configuration)
- [Docker](#docker)

## Overview

// TODO

## Configuration

| ENV Var         | Default  | Required | Note                                                          |   |
|-----------------|----------|----------|---------------------------------------------------------------|---|
| BOT_NAME        | heyemoji | No       | The display name of the heyemoji bot                          |   |
| DATABASE_PATH   | ./data/  | No       | The directory that the database files should be written to    |   |
| SLACK_API_TOKEN |          | Yes      | The API tokens for the Slack API                              |   |
| SLACK_EMOJI     | star:1   | No       | Comma delimited set of emoji "name:value" pairs               |   |
| SLACK_DAILY_CAP | 5        | No       | The max number of emoji points that can be given out in a day |   |
| WEBSOCKET_PORT  | 3334     | No       | Port that the Slack RTM client will listen on                 |   |

## Docker


