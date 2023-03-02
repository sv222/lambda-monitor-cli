package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Config struct {
    AccessKey string   `json:"access_key"`
    SecretKey string   `json:"secret_key"`
    Region    string   `json:"region"`
    Functions []string `json:"functions"`
}

func main() {
    // Parse config file
    configFile, err := os.Open("config.json")
    if err != nil {
        fmt.Println("Error opening config file:", err)
        os.Exit(1)
    }
    defer configFile.Close()

    var config Config
    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        fmt.Println("Error parsing config file:", err)
        os.Exit(1)
    }

    // Create AWS session
    awsSession, err := session.NewSession(&aws.Config{
        Region:      aws.String(config.Region),
        Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
    })
    if err != nil {
        fmt.Println("Error creating AWS session:", err)
        os.Exit(1)
    }

    // Create CloudWatchLogs service client
    cloudWatchLogsClient := cloudwatchlogs.New(awsSession)

    // Monitor Lambda functions
    for _, functionName := range config.Functions {
        fmt.Println("Monitoring Lambda function:", functionName)

        // Get log group name for Lambda function
        logGroupName := "/aws/lambda/" + functionName

        // Get log streams for log group
        logStreams, err := cloudWatchLogsClient.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
            LogGroupName: aws.String(logGroupName),
        })
        if err != nil {
            fmt.Println("Error getting log streams for Lambda function:", functionName, err)
            continue
        }

        // Get latest log events from each log stream
        for _, logStream := range logStreams.LogStreams {
            latestLogEvent, err := cloudWatchLogsClient.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
                LogGroupName:  aws.String(logGroupName),
                LogStreamName: logStream.LogStreamName,
                Limit:         aws.Int64(1),
                StartFromHead: aws.Bool(true),
            })
            if err != nil {
                fmt.Println("Error getting latest log event for Lambda function:", functionName, "log stream:", *logStream.LogStreamName, err)
                continue
            }

            // Print latest log event timestamp and message
            if len(latestLogEvent.Events) > 0 {
                timestamp := time.Unix(*latestLogEvent.Events[0].Timestamp/1000, 0)
                message := *latestLogEvent.Events[0].Message
                fmt.Printf("\tLog stream: %s, Latest event: %v - %s\n", *logStream.LogStreamName, timestamp, message)
            }
        }
    }
}
