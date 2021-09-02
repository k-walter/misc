# Kattis to Telgram
Created by me to monitor, aggregate and notify student assignment attempts on https://nus.kattis.com/. This has only been tested with the queue feature for NUS Kattis.

To deduplicate efforts, Huginn was chosen as a self-hosted FOSS IFTTT I use for many purposes. 

## Features
- [x] Authenticate with Kattis
- [x] Check for latest submissions
- [x] Filter interested users (my students)
- [x] Forward to Telegram bot
- [x] Schedule checks

### Huginn Agents
Each agent passes its output to the next agent instantly.
1. `Kattis Auth` agent runs every 1min to receive a cookie
1. `Kattis Submissions` crawls Kattis queue withe the cookie for new submissions
1. `Kattis Filter` filters interested users out of new updates
1. `Format Kattis --> Telegram` formats from JSON to Telegram readable HTML
1. `Telegram Bot` sends the message to you :)

## Getting Started
### 1. Install Huginns
1. Ensure you have `docker` and `docker-compose` installed. Alternatively, you may install Huginn locally and skip this section.
1. Run `./run.sh`.
1. If unable to run, please change permissions with `sudo chmod +x run.sh`.
1. If port 3000 is in use, please edit within `docker-compose.yml` the line `YOUR_PORT:3000`

### 2. Change configurations
1. Create a Telegram bot with https://t.me/BotFather. You will need the token.
1. Go to your bot and `/start`
1. In `user_credentials.json`, replace
	- `YOUR_USERNAME` with your kattis username/email
	- `YOUR_PASSWORD` with your kattis password
	- `YOUR_TOKEN` with your bot token
	- If you are not using NUS kattis, replace `https://nus.kattis.com/` with your Kattis domain
1. In `kattis-to-telegram.json`, find and replace `John Doe` and `Mary Jane` with the users' names you are interested in.

### 3. Setup huginn
1. Login to Huginn via `localhost:3000`, or your designated port.
	According to Huginn's GitHub, the default username and password are `admin` and `password` respectively. Please refer to the GitHub if this has changed.
1. Upload credentials via Credentials > Upload Credentials > browse > (`user_credentials.json` file)
1. Upload scenario via Scenarios > Import Scenario > browse > browse > (`kattis-to-telegram.json` file)
1. Configure your bot to message you via Agents > Telegram Bot > Actions > Edit agent > Chat > (Your Telegram name). Once done, click on save.

### 4. Verify setup
The agents should run every 1min.
1. Access `localhost:3000/jobs` and wait for 1 run to be complete.
1. Access `localhost:3000/events` to observe the outputs (`events`) passed between the agents.