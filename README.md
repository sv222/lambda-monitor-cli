# AWS Lambda Monitor CLI

AWS Lambda Monitor CLI  is a tool that lets you monitor your Lambda functions from the command line. You can use this tool to monitor the health and performance of your Lambda functions, and to troubleshoot any issues that arise.

## Usage

1. Create a `config.json` file with your AWS access key, secret key, region, and a list of Lambda functions to monitor.

   ```json
   {
       "access_key": "<your_aws_access_key>",
       "secret_key": "<your_aws_secret_key>",
       "region": "<your_aws_region>",
       "functions": [
           "<lambda_function_1>",
           "<lambda_function_2>",
           "<lambda_function_3>"
       ]
   }
   ```

2. Run the aws-lambda-monitor command to monitor the Lambda functions specified in the config.json file.

```sh
        $ aws-lambda-monitor
Monitoring Lambda function: <lambda_function_1>
    Log stream: <log_stream_1>, Latest event: <timestamp> - <log_message>
    Log stream: <log_stream_2>, Latest event: <timestamp> - <log_message>
    ...
Monitoring Lambda function: <lambda_function_2>
    Log stream: <log_stream_1>, Latest event: <timestamp> - <log_message>
    Log stream: <log_stream_2>, Latest event: <timestamp> - <log_message>
    ...
Monitoring Lambda function: <lambda_function_3>
    Log stream: <log_stream_1>, Latest event: <timestamp> - <log_message>
    Log stream: <log_stream_2>, Latest event: <timestamp> - <log_message>
    ...
```

## Installation

1. Clone this repository:

```sh
    git clone https://github.com/sv222/lambda-monitor-cli.git
    cd lambda-monitor-cli
```

2. Build the binary:

```sh
    go build -o lambda-monitor main.go
```

3. Run the binary:

```sh
    ./lambda-monitor
```

## Contribution

If you would like to contribute to the AWS Lambda Monitor CLI, you can fork the project on GitHub and create a pull request with your changes. Contributions are always welcome, and we appreciate your help in making the program better.

## License

AWS Lambda Monitor CLI is licensed under the MIT License.
