# Twitter-streaming-sentiment-analysis

Exploration of a data driven microserivce architecture that gathers sentiment analysis scores on the given tweets.


![alt text](https://github.com/yawlhead91/Twitter-streaming-sentiment-analysis/blob/master/ProjectLayout.svg)

The project is made up of four microservices, a data collection service, a sentiment analysis service, a data store service and the frontend service. Both the data collection service and sentiment analysis service are built using GO. The data collection service streams tweets received from the Twitter RESTfull API. The data is parsed and processed to a format desired by the sentiment analysis service before been streamed to the sentiment analysis service using GRPC's protocol buffers.

The sentiment analysis service receives data from a stream. As the data is received it is prepared to be used with the sentiment scoring library. The scores along with the tweet are the saved in a MongoDB service that allows the frontend to subscribe to it and gather the results. 
