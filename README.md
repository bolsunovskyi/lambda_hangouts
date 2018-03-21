# Hangouts bot lambda
1. Build: `GOOS=linux go build -o main`
2. Archive: `rm -rf deployment.zip && zip deployment.zip main && rm -rf main`
3. Upload...
3.1. Via aws cli: `aws lambda update-function-code --profile self-mike --function-name hangouts-bot --zip-file fileb://deployment.zip`