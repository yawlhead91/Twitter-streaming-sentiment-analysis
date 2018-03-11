var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {
  var db = req.db;
  var collection = db.get('tweet_sentiment');
  collection.find({}, {}, function (e, docs) {
    if(e){
      console.log(e)
    }
    res.render('index', {
      title: 'Tweet Sentiment',
      tweets: docs
    });
  });
});

module.exports = router;