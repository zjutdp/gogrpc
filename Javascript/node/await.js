const axios = require("axios");

async function getFirst10TopicsIncludeFirstReplyAuthor() {
  const response = await axios.get(
    "https://cnodejs.org/api/v1/topics?limit=10"
  );
  const json = response.data;
  console.log('got 10')
  const first10 = json.data.map(topic => {
    return {
      id: topic.id,
      title: topic.title,
      date: topic.create_at,
      author: topic.author.loginname
    };
  });

  for (let topic of first10) {
    const response = await axios.get(
      `https://cnodejs.org/api/v1/topic/${topic.id}`
    );
    const json = response.data;
    console.log(topic.id)
    const firstReply = json.data.replies[0];
    topic.firstReplyAuthor = firstReply && firstReply.author.loginname;
  }

  return first10;
}

getFirst10TopicsIncludeFirstReplyAuthor().then(data => console.log(data));
