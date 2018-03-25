var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {

  var db = req.db;
  var collection = db.get('tweet_sentiment');
  var translated, average;

  // Preform aggregate function to get average
  // sentiment score and total number of tweets
  collection.aggregate([{
    $group: {
      _id: null,
      avgscore: {
        $avg: "$score"
      }
    }
  }], (e, docs) => {

    if (e) next(e);
    average = docs.pop()['avgscore'];

    collection.find({}, {}, (e, docs) => {
      if (e) next(e);

      translated = docs.map(obj => {
        var sentiment = (obj['score'] > 0.5) ? 'Positive' : 'Negative';
        obj['score'] = sentiment
        return obj
      });

      res.render('index', {
        title: 'Tweet Sentiment',
        tweets: translated,
        avgscore: average
      });
      
    });
  });

 

});

module.exports = router;