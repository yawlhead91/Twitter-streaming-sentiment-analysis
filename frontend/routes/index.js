var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {
  var db = req.db;
  var collection = db.get('tweet_sentiment');
  collection.find({}, {}, function (e, docs) {
    if(e){
      console.log(e);
      return e;
    }

    translated = docs.map(obj => {
      var sentiment = (obj['score'] == 1) ? 'Positive' : 'Negative';
      obj['score'] = sentiment
      return obj
    });

    res.render('index', {
      title: 'Tweet Sentiment',
      tweets: translated
    });
  });
});

module.exports = router;