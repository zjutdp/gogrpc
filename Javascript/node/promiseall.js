const axios = require("axios");

function getFirst10TopicsIncludeFirstReplyAuthor() {
  return axios
    .get("https://cnodejs.org/api/v1/topics?limit=10")
    .then(function(response) {
      const json = response.data;
      const first10 = json.data.map(topic => {
        return {
          id: topic.id,
          title: topic.title,
          date: topic.create_at,
          author: topic.author.loginname
        };
      });

      const promises = first10.map(data => {
        return axios
          .get(`https://cnodejs.org/api/v1/topic/${data.id}`)
          .then(response => {
            const json = response.data;
            const firstReply = json.data.replies[0];
            return {
              id: json.data.id,
              firstReplyAuthor: firstReply && firstReply.author.loginname
            };
          });
      });
      return Promise.all(promises).then(rs => {
        const map = rs.reduce((acc, e) => {
          acc.set(e.id, e);
          return acc;
        }, new Map());
        for (let topic of first10) {
          topic.firstReplyAuthor = map.get(topic.id).firstReplyAuthor;
        }
        return first10;
      });
    })
    .catch(function(error) {
      console.log(error);
    });
}

getFirst10TopicsIncludeFirstReplyAuthor().then(data => console.log(data));
